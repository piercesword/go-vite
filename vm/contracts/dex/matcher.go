package dex

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/crypto"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/vm/contracts/dex/proto"
	"math/big"
	"time"
)

const maxTxsCountPerTaker = 1000
const timeoutMillisecond = 7 * 24 * 3600 * 1000

type matcher struct {
	contractAddress *types.Address
	storage         *BaseStorage
	protocol        *nodePayloadProtocol
	books           map[string]*skiplist
	settleActions   map[string]map[string]*proto.SettleAction // address->(asset->settleAction)
}

type OrderTx struct {
	proto.Transaction
	takerAddress string
	makerAddress string
	takerTradeAsset []byte
	takerQuoteAsset []byte
}

func NewMatcher(contractAddress *types.Address, storage *BaseStorage) *matcher {
	mc := &matcher{}
	mc.contractAddress = contractAddress
	mc.storage = storage
	var po nodePayloadProtocol = &OrderNodeProtocol{}
	mc.protocol = &po
	mc.books = make(map[string]*skiplist)
	return mc
}

func (mc *matcher) MatchOrder(order Order) {
	var (
		bookToTake *skiplist
		taker      Order
		ok         bool
	)
	if taker, ok = validateAndRenderOrder(order); !ok {
		mc.emitOrderRes(taker)
		return
	}
	bookToTake = mc.getBookByName(getBookNameToTake(taker))
	if err := mc.doMatchTaker(taker, bookToTake); err != nil {
		fmt.Printf("Failed to match taker for %s", err.Error())
	}
}

func (mc *matcher) GetSettleActions() map[string]map[string]*proto.SettleAction {
	return mc.settleActions
}

func (mc *matcher) GetOrderByIdAndBookName(orderId uint64, makerBookName string) (*Order, error) {
	book := mc.getBookByName(makerBookName)
	if pl, _, _, err := book.getByKey(orderKey{orderId}); err != nil {
		return nil, err
	} else {
		od, _ := (*pl).(Order)
		return &od, nil
	}
}

func (mc *matcher) CancelOrderByIdAndBookName(order *Order, makerBookName string) (err error) {
	book := mc.getBookByName(makerBookName)
	switch order.Status {
	case Pending:
		order.CancelReason = cancelledByUser
	case PartialExecuted:
		order.CancelReason = partialExecutedUserCancelled
	}
	order.Status = Cancelled
	mc.handleRefund(order)
	mc.emitOrderRes(*order)
	pl := nodePayload(order)
	if err = book.updatePayload(newOrderKey(order.Id), &pl); err != nil {
		return err
	}
	return book.delete(newOrderKey(order.Id))
}

func (mc *matcher) getBookByName(bookName string) *skiplist {
	if book, ok := mc.books[bookName]; !ok {
		book = newSkiplist(bookName, mc.contractAddress, mc.storage, mc.protocol)
		mc.books[bookName] = book
	}
	return mc.books[bookName]
}

func (mc *matcher) doMatchTaker(taker Order, makerBook *skiplist) error {
	if makerBook.length == 0 {
		mc.handleTakerRes(taker)
		return nil
	}
	modifiedMakers := make([]Order, 0, 20)
	txs := make([]OrderTx, 0, 20)
	if maker, nextOrderId, err := getMakerById(makerBook, makerBook.header); err != nil {
		return err
	} else {
		if err := mc.recursiveTakeOrder(&taker, maker, makerBook, &modifiedMakers, &txs, nextOrderId); err != nil {
			return err
		} else {
			mc.handleTakerRes(taker)
			mc.handleModifiedMakers(modifiedMakers, makerBook)
			mc.emitTxs(txs)
			return nil
		}
	}
}
//TODO add assertion for order calculation correctness
func (mc *matcher) recursiveTakeOrder(taker *Order, maker Order, makerBook *skiplist, modifiedMakers *[]Order, txs *[]OrderTx, nextOrderId nodeKeyType) error {
	if filterTimeout(&maker) {
		mc.handleRefund(&maker)
		*modifiedMakers = append(*modifiedMakers, maker)
	} else {
		matched, _ := matchPrice(*taker, maker)
		fmt.Printf("recursiveTakeOrder matched for taker.id %d is %t\n", taker.Id, matched)
		if matched {
			tx := calculateOrderAndTx(taker, &maker)
			mc.handleRefund(taker)
			mc.handleRefund(&maker)
			*modifiedMakers = append(*modifiedMakers, maker)
			*txs = append(*txs, tx)
		}
	}
	if taker.Status == PartialExecuted && len(*txs) >= maxTxsCountPerTaker {
		taker.Status = Cancelled
		taker.CancelReason = partialExecutedCancelledByMarket
	}
	if taker.Status == FullyExecuted || taker.Status == Cancelled {
		return nil
	}
	// current maker is the last item in the book
	if makerBook.tail.(orderKey).value == maker.Id {
		return nil
	}
	var err error
	if maker, nextOrderId, err = getMakerById(makerBook, nextOrderId); err != nil {
		return errors.New("Failed get order by nextOrderId")
	} else {
		return mc.recursiveTakeOrder(taker, maker, makerBook, modifiedMakers, txs, nextOrderId)
	}
}

func (mc *matcher) handleModifiedMakers(makers []Order, makerBook *skiplist) {
	for _, maker := range makers {
		pl := nodePayload(maker)
		if err := makerBook.updatePayload(newOrderKey(maker.Id), &pl); err != nil {
			fmt.Printf("failed update maker storage for err : %s\n", err.Error())
		}
		mc.emitOrderRes(maker)
	}
	size := len(makers)
	if size > 0 {
		if makers[size-1].Status == FullyExecuted || makers[size-1].Status == Cancelled {
			makerBook.truncateHeadTo(newOrderKey(makers[size-1].Id), size)
		} else if size >= 2 {
			makerBook.truncateHeadTo(newOrderKey(makers[size-2].Id), size-1)
		}
	}
}

func (mc *matcher) handleTakerRes(order Order) {
	if order.Status == PartialExecuted || order.Status == Pending {
		mc.saveTakerAsMaker(order)
	}
	mc.emitOrderRes(order)
}

func (mc *matcher) saveTakerAsMaker(maker Order) {
	bookToMake := mc.getBookByName(getBookNameToMakeForOrder(maker))
	pl := nodePayload(maker)
	bookToMake.insert(newOrderKey(maker.Id), &pl)
}

func (mc *matcher) emitOrderRes(orderRes Order) {
	event := OrderUpdateEvent{orderRes.Order}
	(*mc.storage).AddLog(newLog(event))
	fmt.Printf("order matched res %s : \n", orderRes.String())
}

func (mc *matcher) emitTxs(txs []OrderTx) {
	fmt.Printf("matched txs >>>>>>>>> %d\n", len(txs))
	for _, tx := range txs {
		mc.handleTxSettleAction(tx)
		txEvent := TransactionEvent{tx.Transaction}
		(*mc.storage).AddLog(newLog(txEvent))
		fmt.Printf("matched tx is : %s\n", tx.String())
	}
}

func (mc *matcher) handleRefund(order *Order) {
	if order.Status == FullyExecuted || order.Status == Cancelled {
		switch order.Side {
		case false: //buy
			order.RefundAsset = order.QuoteAsset
			order.RefundQuantity = order.Amount - order.ExecutedAmount
		case true:
			order.RefundAsset = order.TradeAsset
			order.RefundQuantity = order.Quantity - order.ExecutedQuantity
		}
		mc.updateSettleAction(proto.SettleAction{Address:order.Address, Asset:order.RefundAsset, ReleaseLocked:order.RefundQuantity})
	}
}

func (mc *matcher) handleTxSettleAction(tx OrderTx) {
	takerInSettle := proto.SettleAction{Address : tx.takerAddress}
	takerOutSettle := proto.SettleAction{Address : tx.takerAddress}
	makerInSettle := proto.SettleAction{Address : tx.makerAddress}
	makerOutSettle := proto.SettleAction{Address : tx.makerAddress}
	switch tx.TakerSide {
		case false: //buy
			takerInSettle.Asset = tx.takerTradeAsset
			takerInSettle.IncAvailable = tx.Quantity
			makerOutSettle.Asset = tx.takerTradeAsset
			makerOutSettle.DeduceLocked = tx.Quantity

			takerOutSettle.Asset = tx.takerQuoteAsset
			takerOutSettle.DeduceLocked = tx.Amount
			makerInSettle.Asset = tx.takerQuoteAsset
			makerInSettle.IncAvailable = tx.Amount

		case true: //sell
			takerInSettle.Asset = tx.takerQuoteAsset
			takerInSettle.IncAvailable = tx.Amount
			makerOutSettle.Asset = tx.takerQuoteAsset
			makerOutSettle.DeduceLocked = tx.Amount

			takerOutSettle.Asset = tx.takerTradeAsset
			takerOutSettle.DeduceLocked = tx.Quantity
			makerInSettle.Asset = tx.takerTradeAsset
			makerInSettle.IncAvailable = tx.Quantity
	}
	for _, ac := range []proto.SettleAction{takerInSettle, takerOutSettle, makerInSettle, makerOutSettle} {
		mc.updateSettleAction(ac)
	}
}

func (mc *matcher) updateSettleAction(action proto.SettleAction) {
	var (
		actionMap map[string]*proto.SettleAction
		ok bool
		ac *proto.SettleAction
		address = string(action.Address)
	)
	if actionMap, ok = mc.settleActions[address]; !ok {
		actionMap = make(map[string]*proto.SettleAction)
	}
	asset := string(action.Asset)
	if ac, ok = actionMap[string(asset)]; !ok {
		ac = &proto.SettleAction{Address:address, Asset:action.Asset}
	}
	ac.IncAvailable += action.IncAvailable
	ac.ReleaseLocked += action.ReleaseLocked
	ac.DeduceLocked += action.DeduceLocked
	actionMap[address] = ac
	mc.settleActions[address] = actionMap
}


func newLog(event OrderEvent) *ledger.VmLog {
	log := &ledger.VmLog{}
	log.Topics = append(log.Topics, event.getTopicId())
	log.Data = event.toDataBytes()
	return log
}

func matchPrice(taker Order, maker Order) (matched bool, executedPrice float64) {
	if taker.Type == Market || priceEqual(taker.Price, maker.Price) {
		return true, maker.Price
	} else {
		matched = false
		switch maker.Side {
		case false: // buy
			matched = maker.Price > taker.Price
		case true: // sell
			matched = maker.Price < taker.Price
		}
		return matched, maker.Price
	}
}

func filterTimeout(maker *Order) bool {
	if time.Now().Unix()*1000 > maker.Timestamp+timeoutMillisecond {
		switch maker.Status {
		case Pending:
			maker.CancelReason = cancelledOnTimeout
		case PartialExecuted:
			maker.CancelReason = partialExecutedCancelledOnTimeout
		default:
			maker.CancelReason = unknownCancelledOnTimeout
		}
		maker.Status = Cancelled
		return true
	} else {
		return false
	}
}

func calculateOrderAndTx(taker *Order, maker *Order) (tx OrderTx) {
	tx = OrderTx{}
	tx.Id = getTxId(taker.Id, maker.Id)
	tx.TakerSide = taker.Side
	tx.TakerId = taker.Id
	tx.MakerId = maker.Id
	tx.Price = maker.Price
	executeQuantity := minUint64(taker.Quantity-taker.ExecutedQuantity, maker.Quantity-maker.ExecutedQuantity)
	takerAmount := calculateOrderAmount(taker, executeQuantity, maker.Price)
	makerAmount := calculateOrderAmount(maker, executeQuantity, maker.Price)
	executeAmount := minUint64(takerAmount, makerAmount)
	executeQuantity = minUint64(executeQuantity, calculateQuantity(executeAmount, maker.Price))
	fmt.Printf("calculateOrderAndTx >>> taker.ExecutedQuantity : %d, maker.ExecutedQuantity : %d, maker.Price : %f, executeQuantity : %d, executeAmount : %d \n",
		taker.ExecutedQuantity,
		maker.ExecutedQuantity,
		maker.Price,
		executeQuantity,
		executeAmount)
	updateOrder(taker, executeQuantity, executeAmount)
	updateOrder(maker, executeQuantity, executeAmount)
	tx.Quantity = executeQuantity
	tx.Amount = executeAmount
	tx.Timestamp = time.Now().UnixNano() / 1000
	tx.takerAddress = taker.Address
	tx.makerAddress = maker.Address
	tx.takerTradeAsset = taker.TradeAsset
	tx.takerQuoteAsset = taker.QuoteAsset
	tx.makerAddress = maker.Address
	return tx
}

func calculateOrderAmount(order *Order, quantity uint64, price float64) uint64 {
	amount := CalculateAmount(quantity, price)
	if !order.Side && order.Amount < order.ExecutedAmount+amount {// side is buy
		amount = order.Amount - order.ExecutedAmount
	}
	return amount
}

func updateOrder(order *Order, quantity uint64, amount uint64) uint64 {
	order.ExecutedAmount += amount
	if order.Quantity-order.ExecutedQuantity == quantity || isDust(order, quantity) {
		order.Status = FullyExecuted
	} else {
		order.Status = PartialExecuted
	}
	order.ExecutedQuantity += quantity
	return amount
}
// leave quantity is too small for calculate precision
func isDust(order *Order, quantity uint64) bool {
	return CalculateAmount(order.Quantity-order.ExecutedQuantity - quantity, order.Price) < 1
}

func getMakerById(makerBook *skiplist, orderId nodeKeyType) (od Order, nextId orderKey, err error) {
	if pl, fwk, _, err := makerBook.getByKey(orderId); err != nil {
		return Order{}, (*makerBook.protocol).getNilKey().(orderKey), err
	} else {
		maker := (*pl).(Order)
		return maker, fwk.(orderKey), nil
	}
}

func validateAndRenderOrder(order Order) (orderRes Order, isValid bool) {
	if order.Id == 0 || !validPrice(order.Price) {
		order.Status = Cancelled
		order.CancelReason = cancelledByMarket
		return order, false
	} else {
		order.Status = Pending
		if order.Amount == 0 {
			order.Amount = CalculateAmount(order.Quantity, order.Price)
		}
		return order, true
	}
}

func validPrice(price float64) bool {
	if price < 0 || price < 0.00000001 {
		return false
	} else {
		return true
	}
}

func CalculateAmount(quantity uint64, price float64) uint64 {
	qtF := big.NewFloat(0).SetUint64(quantity)
	prF := big.NewFloat(0).SetFloat64(price)
	amountF := prF.Mul(prF, qtF)
	amount, _ := amountF.Add(amountF, big.NewFloat(0.5)).Uint64()
	return amount
}

func calculateQuantity(amount uint64, price float64) uint64 {
	amtF := big.NewFloat(0).SetUint64(amount)
	prF := big.NewFloat(0).SetFloat64(price)
	qtyF := amtF.Quo(amtF, prF)
	qty, _ := qtyF.Add(qtyF, big.NewFloat(0.5)).Uint64()
	return qty
}

func getBookNameToTake(order Order) string {
	return fmt.Sprintf("%s|%s|%d", string(order.TradeAsset), string(order.QuoteAsset), 1-toSideInt(order.Side))
}

func getBookNameToMakeForOrder(order Order) string {
	return GetBookNameToMake(order.TradeAsset, order.QuoteAsset, order.Side)
}

func GetBookNameToMake(tradeAsset []byte, quoteAsset []byte, side bool) string {
	return fmt.Sprintf("%s|%s|%d", string(tradeAsset), string(quoteAsset), toSideInt(side))
}

func toSideInt(side bool) int {
	sideInt := 0 // buy
	if side {
		sideInt = 1 // sell
	}
	return sideInt
}

func minUint64(a uint64, b uint64) uint64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func getTxId(takerId uint64, makerId uint64) uint64 {
	data := crypto.Hash(8, []byte(fmt.Sprintf("%d-%d", takerId, makerId)))
	var num uint64
	binary.Read(bytes.NewBuffer(data[:]), binary.LittleEndian, &num)
	return num
}

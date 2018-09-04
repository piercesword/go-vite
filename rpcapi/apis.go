package rpcapi

import (
	"github.com/vitelabs/go-vite/rpc"
	"github.com/vitelabs/go-vite/rpcapi/api"
	"github.com/vitelabs/go-vite/vite"
)

func GetAllApis(vite *vite.Vite) []rpc.API {
	ledgerApis := rpc.API{
		Namespace: "ledger",
		Version:   "1.0",
		Service:   api.NewLedgerApi(vite),
		Public:    true,
	}

	walletApis := rpc.API{
		Namespace: "wallet",
		Version:   "1.0",
		Service:   api.NewWalletApi(vite),
		Public:    false,
	}

	p2pApis := rpc.API{
		Namespace: "p2p",
		Version:   "1.0",
		Service:   api.NewP2PApi(vite.P2p()),
		Public:    true,
	}

	typesApis := rpc.API{
		Namespace: "types",
		Version:   "1.0",
		Service:   api.TypesApi{},
		Public:    true,
	}

	commonApis := rpc.API{
		Namespace: "common",
		Version:   "1.0",
		Service:   api.CommonApi{},
		Public:    true,
	}

	return []rpc.API{
		ledgerApis,
		walletApis,
		p2pApis,
		typesApis,
		commonApis,
	}
}

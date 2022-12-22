package precompile

import (
	"context"
	"testing"

	ethclient2 "github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/interfaces"
	"github.com/ethereum/go-ethereum/common"
)

func TestMulticall(t *testing.T) {
	url := "http://127.0.0.1:9654/ext/bc/jtd1894n4pQ2Ph3bdR22vWcEiBvbvWf6si67N9GDUr6hmHBTt/rpc"
	mcAddr := common.HexToAddress("0xA4cD3b0Eb6E5Ab5d8CE4065BcCD70040ADAB1F00")
	fnsig := CalculateFunctionSelector("getCurrentBlockTimestamp()")
	// input := common.Hex2Bytes("0x1111")
	// input := []byte{}

	calls := []IMulticallCall{{
		Target:   mcAddr,
		CallData: fnsig,
	}, {
		Target:   mcAddr,
		CallData: fnsig,
	}}
	t.Logf("%+v", calls)

	data, err := PackAggregate(calls)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)

	client, err := ethclient2.Dial(url)
	if err != nil {
		t.Fatal(err)
	}

	cm := interfaces.CallMsg{To: &mcAddr, Data: data}
	out, err := client.CallContract(context.Background(), cm, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := MulticallABI.Unpack("aggregate", out)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", res)
	t.Fatal()
}

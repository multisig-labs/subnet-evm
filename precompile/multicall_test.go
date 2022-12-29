package precompile

import (
	"context"
	"testing"

	ethclient2 "github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/interfaces"
	interfaces2 "github.com/ava-labs/coreth/interfaces"
	"github.com/ethereum/go-ethereum/common"
)

func TestFoo(t *testing.T) {
	input := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 224, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 111, 217, 2, 225, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 15, 40, 201, 125, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 36, 248, 178, 203, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 141, 185, 124, 124, 236, 226, 73, 194, 185, 139, 220, 2, 38, 204, 76, 42, 87, 191, 82, 252, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	url := "http://127.0.0.1:9650/ext/bc/C/rpc"
	client, err := ethclient2.Dial(url)
	if err != nil {
		t.Fatal(err)
	}
	cchainMulticallAddr := common.HexToAddress("0x5aa01B3b5877255cE50cc55e8986a7a5fe29C70e")
	m := MulticallABI.Methods["aggregate"]
	callData := append(m.ID, input...)
	t.Logf("callData %x", callData)
	cm := interfaces2.CallMsg{To: &cchainMulticallAddr, Data: callData}
	results, err := client.CallContract(context.Background(), cm, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("results %x", results)
	t.Fatal()
}

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

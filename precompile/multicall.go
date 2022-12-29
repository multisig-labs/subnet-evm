// Code generated
// This file is a generated precompile contract with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

// There are some must-be-done changes waiting in the file. Each area requiring you to add your code is marked with CUSTOM CODE to make them easy to find and modify.
// Additionally there are other files you need to edit to activate your precompile.
// These areas are highlighted with comments "ADD YOUR PRECOMPILE HERE".
// For testing take a look at other precompile tests in core/stateful_precompile_test.go

/* General guidelines for precompile development:
1- Read the comment and set a suitable contract address in precompile/params.go. E.g:
	MulticallAddress = common.HexToAddress("ASUITABLEHEXADDRESS")
2- Set gas costs here
3- It is recommended to only modify code in the highlighted areas marked with "CUSTOM CODE STARTS HERE". Modifying code outside of these areas should be done with caution and with a deep understanding of how these changes may impact the EVM.
Typically, custom codes are required in only those areas.
4- Add your upgradable config in params/precompile_config.go
5- Add your precompile upgrade in params/config.go
6- Add your solidity interface and test contract to contract-examples/contracts
7- Write solidity tests for your precompile in contract-examples/test
8- Create your genesis with your precompile enabled in tests/e2e/genesis/
9- Create e2e test for your solidity test in tests/e2e/solidity/suites.go
10- Run your e2e precompile Solidity tests with 'E2E=true ./scripts/run.sh'

*/

package precompile

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	ethclient2 "github.com/ava-labs/coreth/ethclient"
	interfaces2 "github.com/ava-labs/coreth/interfaces"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/vmerrs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

const (
	AggregateGasCost                uint64 = readGasCostPerSlot // Should be bigger but who cares?
	GetBalanceGasCost               uint64 = readGasCostPerSlot // SET A GAS COST HERE
	GetCurrentBlockNumberGasCost    uint64 = readGasCostPerSlot // SET A GAS COST HERE
	GetCurrentBlockTimestampGasCost uint64 = readGasCostPerSlot // SET A GAS COST HERE

	// MulticallRawABI contains the raw ABI of Multicall contract.
	MulticallRawABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
)

// Singleton StatefulPrecompiledContract and signatures.
var (
	_ StatefulPrecompileConfig = &MulticallConfig{}

	MulticallABI abi.ABI // will be initialized by init function

	MulticallPrecompile StatefulPrecompiledContract // will be initialized by init function
)

// MulticallConfig implements the StatefulPrecompileConfig
// interface while adding in the Multicall specific precompile address.
type MulticallConfig struct {
	UpgradeableConfig
}

// IMulticallCall is an auto generated low-level Go binding around an user-defined struct.
type IMulticallCall struct {
	Target   common.Address
	CallData []byte
}

type AggregateOutput struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}

func init() {
	parsed, err := abi.JSON(strings.NewReader(MulticallRawABI))
	if err != nil {
		panic(err)
	}
	MulticallABI = parsed

	MulticallPrecompile = createMulticallPrecompile(MulticallAddress)
}

// NewMulticallConfig returns a config for a network upgrade at [blockTimestamp] that enables
// Multicall .
func NewMulticallConfig(blockTimestamp *big.Int) *MulticallConfig {
	return &MulticallConfig{

		UpgradeableConfig: UpgradeableConfig{BlockTimestamp: blockTimestamp},
	}
}

// NewDisableMulticallConfig returns config for a network upgrade at [blockTimestamp]
// that disables Multicall.
func NewDisableMulticallConfig(blockTimestamp *big.Int) *MulticallConfig {
	return &MulticallConfig{
		UpgradeableConfig: UpgradeableConfig{
			BlockTimestamp: blockTimestamp,
			Disable:        true,
		},
	}
}

// Equal returns true if [s] is a [*MulticallConfig] and it has been configured identical to [c].
func (c *MulticallConfig) Equal(s StatefulPrecompileConfig) bool {
	// typecast before comparison
	other, ok := (s).(*MulticallConfig)
	if !ok {
		return false
	}
	// CUSTOM CODE STARTS HERE
	// modify this boolean accordingly with your custom MulticallConfig, to check if [other] and the current [c] are equal
	// if MulticallConfig contains only UpgradeableConfig  you can skip modifying it.
	equals := c.UpgradeableConfig.Equal(&other.UpgradeableConfig)
	return equals
}

// String returns a string representation of the MulticallConfig.
func (c *MulticallConfig) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}

// Address returns the address of the Multicall. Addresses reside under the precompile/params.go
// Select a non-conflicting address and set it in the params.go.
func (c *MulticallConfig) Address() common.Address {
	return MulticallAddress
}

// Configure configures [state] with the initial configuration.
func (c *MulticallConfig) Configure(_ ChainConfig, state StateDB, _ BlockContext) {

	// CUSTOM CODE STARTS HERE
}

// Contract returns the singleton stateful precompiled contract to be used for Multicall.
func (c *MulticallConfig) Contract() StatefulPrecompiledContract {
	return MulticallPrecompile
}

// Verify tries to verify MulticallConfig and returns an error accordingly.
func (c *MulticallConfig) Verify() error {

	// CUSTOM CODE STARTS HERE
	// Add your own custom verify code for MulticallConfig here
	// and return an error accordingly
	return nil
}

// UnpackAggregateInput attempts to unpack [input] into the []IMulticallCall type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackAggregateInput(input []byte) ([]IMulticallCall, error) {
	res, err := MulticallABI.UnpackInput("aggregate", input)
	if err != nil {
		return nil, err
	}
	unpacked := *abi.ConvertType(res[0], new([]IMulticallCall)).(*[]IMulticallCall)
	return unpacked, nil
}

// PackAggregate packs [calls] of type []IMulticallCall into the appropriate arguments for aggregate.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackAggregate(calls []IMulticallCall) ([]byte, error) {
	return MulticallABI.Pack("aggregate", calls)
}

// PackAggregateOutput attempts to pack given [outputStruct] of type AggregateOutput
// to conform the ABI outputs.
func PackAggregateOutput(outputStruct AggregateOutput) ([]byte, error) {
	return MulticallABI.PackOutput("aggregate",
		outputStruct.BlockNumber,
		outputStruct.ReturnData,
	)
}

// Multicall 0xA4cD3b0Eb6E5Ab5d8CE4065BcCD70040ADAB1F00
func aggregate(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	log.Error("aggregate", "input", input)
	if remainingGas, err = deductGas(suppliedGas, AggregateGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	// attempts to unpack [input] into the arguments to the AggregateInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	// inputStruct, err := UnpackAggregateInput(input)
	// if err != nil {
	// 	return nil, remainingGas, err
	// }

	// CUSTOM CODE STARTS HERE
	// _ = inputStruct // CUSTOM CODE OPERATES ON INPUT

	// // TODO how to get the URL port? Config?
	url := "http://127.0.0.1:9650/ext/bc/C/rpc"
	client, err := ethclient2.Dial(url)
	if err != nil {
		return nil, remainingGas, err
	}
	cchainMulticallAddr := common.HexToAddress("0x5aa01B3b5877255cE50cc55e8986a7a5fe29C70e")
	m := MulticallABI.Methods["aggregate"]
	callData := append(m.ID, input...)
	log.Error("aggregate", "callData", callData)
	cm := interfaces2.CallMsg{To: &cchainMulticallAddr, Data: callData}
	results, err := client.CallContract(context.Background(), cm, nil)
	if err != nil {
		return nil, remainingGas, err
	}
	log.Error("aggregate", "results", results)

	// var output AggregateOutput // CUSTOM CODE FOR AN OUTPUT
	// packedOutput, err := PackAggregateOutput(output)
	// if err != nil {
	// 	return nil, remainingGas, err
	// }

	// Return the packed output and the remaining gas
	// return packedOutput, remainingGas, nil
	return results, remainingGas, nil
}

// UnpackGetBalanceInput attempts to unpack [input] into the common.Address type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackGetBalanceInput(input []byte) (common.Address, error) {
	res, err := MulticallABI.UnpackInput("getBalance", input)
	if err != nil {
		return common.Address{}, err
	}
	unpacked := *abi.ConvertType(res[0], new(common.Address)).(*common.Address)
	return unpacked, nil
}

// PackGetBalance packs [addr] of type common.Address into the appropriate arguments for getBalance.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetBalance(addr common.Address) ([]byte, error) {
	return MulticallABI.Pack("getBalance", addr)
}

// PackGetBalanceOutput attempts to pack given balance of type *big.Int
// to conform the ABI outputs.
func PackGetBalanceOutput(balance *big.Int) ([]byte, error) {
	return MulticallABI.PackOutput("getBalance", balance)
}

func getBalance(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = deductGas(suppliedGas, GetBalanceGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the GetBalanceInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackGetBalanceInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE

	// TODO how to get the URL?
	url := "http://127.0.0.1:9650/ext/bc/C/rpc"
	client, err := ethclient2.Dial(url)
	if err != nil {
		return nil, remainingGas, err
	}

	bal, err := client.BalanceAt(context.Background(), inputStruct, nil)
	if err != nil {
		return nil, remainingGas, err
	}

	packedOutput, err := PackGetBalanceOutput(bal)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// PackGetCurrentBlockNumber packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetCurrentBlockNumber() ([]byte, error) {
	return MulticallABI.Pack("getCurrentBlockNumber")
}

// PackGetCurrentBlockNumberOutput attempts to pack given blockNumber of type *big.Int
// to conform the ABI outputs.
func PackGetCurrentBlockNumberOutput(blockNumber *big.Int) ([]byte, error) {
	return MulticallABI.PackOutput("getCurrentBlockNumber", blockNumber)
}

func getCurrentBlockNumber(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	methodAggregate := MulticallABI.Methods["aggregate"]
	log.Error("aggregate", "methodAggregate.ID", methodAggregate.ID)

	if remainingGas, err = deductGas(suppliedGas, GetCurrentBlockNumberGasCost); err != nil {
		return nil, 0, err
	}
	// no input provided for this function 0x252DBA42

	// CUSTOM CODE STARTS HERE

	// TODO how to get the URL port? config?
	url := "http://127.0.0.1:9650/ext/bc/C/rpc"
	client, err := ethclient2.Dial(url)
	if err != nil {
		return nil, remainingGas, err
	}

	num, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, remainingGas, err
	}

	output := big.NewInt(int64(num))
	packedOutput, err := PackGetCurrentBlockNumberOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// PackGetCurrentBlockTimestamp packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetCurrentBlockTimestamp() ([]byte, error) {
	return MulticallABI.Pack("getCurrentBlockTimestamp")
}

// PackGetCurrentBlockTimestampOutput attempts to pack given timestamp of type *big.Int
// to conform the ABI outputs.
func PackGetCurrentBlockTimestampOutput(timestamp *big.Int) ([]byte, error) {
	return MulticallABI.PackOutput("getCurrentBlockTimestamp", timestamp)
}

func getCurrentBlockTimestamp(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = deductGas(suppliedGas, GetCurrentBlockTimestampGasCost); err != nil {
		return nil, 0, err
	}
	// no input provided for this function

	// CUSTOM CODE STARTS HERE

	var output *big.Int // CUSTOM CODE FOR AN OUTPUT
	packedOutput, err := PackGetCurrentBlockTimestampOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createMulticallPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createMulticallPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	var functions []*statefulPrecompileFunction

	methodAggregate, ok := MulticallABI.Methods["aggregate"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodAggregate.ID, aggregate))

	methodGetBalance, ok := MulticallABI.Methods["getBalance"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodGetBalance.ID, getBalance))

	methodGetCurrentBlockNumber, ok := MulticallABI.Methods["getCurrentBlockNumber"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodGetCurrentBlockNumber.ID, getCurrentBlockNumber))

	methodGetCurrentBlockTimestamp, ok := MulticallABI.Methods["getCurrentBlockTimestamp"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodGetCurrentBlockTimestamp.ID, getCurrentBlockTimestamp))

	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, functions)
	return contract
}

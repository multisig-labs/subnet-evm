pragma solidity ^0.8.0;

import './IMulticall.sol';

contract ExampleMulticall {
	address constant MULTICALL_ADDRESS = 0x0300000000000000000000000000000000000000;
	IMulticall mc = IMulticall(MULTICALL_ADDRESS);

	constructor() {}

	function getCurrentBlockNumber() public view returns (uint256) {
		return mc.getCurrentBlockNumber();
	}
}

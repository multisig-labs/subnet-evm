pragma solidity ^0.8.0;

import './IMulticall.sol';
import './Multicall.sol';

contract ExampleMulticall {
	address constant MULTICALL_ADDRESS = 0x0300000000000000000000000000000000000000;
	IMulticall mc = IMulticall(MULTICALL_ADDRESS);

	constructor(address addr) {
		mc = IMulticall(addr);
	}

	function test() public returns (bytes[] memory returnData) {
		IMulticall.Call[] memory calls = new IMulticall.Call[](3);

		IMulticall.Call memory c1;
		c1.target = address(mc);
		c1.callData = abi.encodeWithSignature('getCurrentBlockNumber()');
		calls[0] = c1;

		IMulticall.Call memory c2;
		c2.target = address(mc);
		c2.callData = abi.encodeWithSignature('getCurrentBlockTimestamp()');
		calls[1] = c2;

		IMulticall.Call memory c3;
		c3.target = address(mc);
		c3.callData = abi.encodeWithSignature('getBalance(address)', 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC);
		calls[2] = c3;

		(uint256 bn, bytes[] memory d) = mc.aggregate(calls);
		return d;
	}

	function getCurrentBlockNumber() public view returns (uint256) {
		return mc.getCurrentBlockNumber();
	}

	function getBalance(address addr) public view returns (uint256) {
		return mc.getBalance(addr);
	}
}

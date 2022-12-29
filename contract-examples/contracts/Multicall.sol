// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Multicall {
	struct Call {
		address target;
		bytes callData;
	}

	function aggregate(Call[] memory calls) public returns (uint256 blockNumber, bytes[] memory returnData) {
		blockNumber = block.number;
		returnData = new bytes[](calls.length);
		for (uint256 i = 0; i < calls.length; i++) {
			// solhint-disable
			(bool success, bytes memory ret) = calls[i].target.call(calls[i].callData);
			require(success);
			returnData[i] = ret;
		}
	}

	function getBalance(address addr) public view returns (uint256 balance) {
		balance = addr.balance;
	}

	function getCurrentBlockTimestamp() public view returns (uint256 timestamp) {
		timestamp = block.timestamp;
	}

	function getCurrentBlockNumber() public view returns (uint256 blockNumber) {
		blockNumber = block.number;
	}
}

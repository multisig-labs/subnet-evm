// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface IMulticall {
	struct Call {
		address target;
		bytes callData;
	}

	function aggregate(Call[] memory calls) external returns (uint256 blockNumber, bytes[] memory returnData);

	function getBalance(address addr) external view returns (uint256 balance);

	function getCurrentBlockTimestamp() external view returns (uint256 timestamp);

	function getCurrentBlockNumber() external view returns (uint256 blockNumber);
}

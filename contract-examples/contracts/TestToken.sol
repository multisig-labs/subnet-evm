//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';

contract TestToken is ERC20 {
	string private constant TOKEN_NAME = 'TestToken';
	string private constant TOKEN_SYMBOL = 'TKN';

	constructor() ERC20(TOKEN_NAME, TOKEN_SYMBOL) {
		_mint(_msgSender(), 1000000 ether);
	}
}

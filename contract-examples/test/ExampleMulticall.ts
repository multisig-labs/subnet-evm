import { expect } from "chai";
import { ethers } from "hardhat";
import { Contract, ContractFactory } from "ethers";

describe("ExampleMulticall", function () {
  let multicallContract: Contract;
  const tokenAddr = "0x95CA0a568236fC7413Cd2b794A7da24422c2BBb6";
  const multicallAddr = "0x789a5FDac2b37FCD290fb2924382297A6AE65860";

  before(async function () {
    // Deploy Contract
    const ContractF: ContractFactory = await ethers.getContractFactory(
      "ExampleMulticall"
    );
    multicallContract = await ContractF.deploy();
    await multicallContract.deployed();
    const multicallContractAddress: string = multicallContract.address;
    console.log(`Contract deployed to: ${multicallContractAddress}`);
  });

  it("should getCurrentBlockNumber", async function () {
    let result = await multicallContract.callStatic.getCurrentBlockNumber();
    console.log(result);
  });

  it("should getBalance", async function () {
    let result = await multicallContract.callStatic.getBalance(
      "0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC"
    );
    console.log(result);
  });
});

import { expect } from "chai";
import { ethers } from "hardhat";
import { Contract, ContractFactory } from "ethers";

describe("ExampleMulticall", function () {
  let multicallContract: Contract;

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
    // expect(result).to.equal("foo");
  });
});

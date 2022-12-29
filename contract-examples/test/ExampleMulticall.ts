import { expect } from "chai";
import { ethers } from "hardhat";
import { Contract, ContractFactory } from "ethers";

describe("ExampleMulticall", function () {
  let multicallContract: Contract;
  let tokenContract: Contract;
  // const tokenAddr = "0x5aa01B3b5877255cE50cc55e8986a7a5fe29C70e";
  // const multicallAddr = "0x52C84043CD9c865236f11d9Fc9F56aa003c1f922";

  before(async function () {
    const ContractF: ContractFactory = await ethers.getContractFactory(
      "ExampleMulticall"
    );
    multicallContract = await ContractF.deploy(
      "0x0300000000000000000000000000000000000000"
    );
    await multicallContract.deployed();
    const multicallContractAddress: string = multicallContract.address;
    console.log(`Multicall Contract deployed to: ${multicallContractAddress}`);
  });

  it("should test", async function () {
    let result = await multicallContract.callStatic.test();
    console.log(result);
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

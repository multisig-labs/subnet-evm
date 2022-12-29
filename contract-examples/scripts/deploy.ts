import { writeFile } from "node:fs/promises";
import { ethers, network } from "hardhat";

type IFU = { [key: string]: any };

const addresses: IFU = {};
const instances: IFU = {};

// ContractName: [constructorArgs...]
const contracts: IFU = {
  Multicall: [],
  TestToken: [],
  ExampleMulticall: ["Multicall"],
};

const deploy = async () => {
  console.log(`Network: ${network.name}`);

  for (const contract in contracts) {
    const args = [];
    for (const name of contracts[contract]) {
      args.push(addresses[name]);
    }
    console.log(`Deploying ${contract} with args ${args}...`);
    const C = await ethers.getContractFactory(contract);
    const c = await C.deploy(...args);
    const inst = await c.deployed();
    instances[contract] = inst;
    addresses[contract] = c.address;
    console.log(`${contract} deployed to: ${c.address}`);
  }

  // Write out the deployed addresses to a format easily loaded by bash for use by cast
  let data = "declare -A addrs=(";
  for (const name in addresses) {
    data = data + `[${name}]="${addresses[name]}" `;
  }
  data = data + ")";
  await writeFile(`cache/deployed_addrs_${network.name}.bash`, data);

  // Write out the deployed addresses to a format easily loaded by javascript
  data = `module.exports = ${JSON.stringify(addresses, null, 2)}`;
  await writeFile(`cache/deployed_addrs_${network.name}.js`, data);

  // Write out the deployed addresses to json (used by Rialto during dev)
  data = JSON.stringify(addresses, null, 2);
  await writeFile(`cache/deployed_addrs_${network.name}.json`, data);
};

deploy()
  .then(() => {
    console.log("Done!");
    // eslint-disable-next-line no-process-exit
    process.exit(0);
  })
  .catch((error) => {
    console.error(error);
    throw error;
  });

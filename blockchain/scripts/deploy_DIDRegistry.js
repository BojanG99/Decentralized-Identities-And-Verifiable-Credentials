// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {
  const [dummy, signer] = await hre.ethers.getSigners();

  const network = await hre.network;

  console.log("Network Name:", network);

  const DIDRegistryContract = await hre.ethers.getContractFactory(
    "DIDRegistryContract"
  );

  did_registry_contract = await DIDRegistryContract.connect(signer).deploy();

  console.log(
    `DIDRegistryContract deployed to ${did_registry_contract.target}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

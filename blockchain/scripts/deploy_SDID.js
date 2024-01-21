// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {
  const [dummy, first, signer] = await hre.ethers.getSigners();

  const network = await hre.network;

  console.log("Network Name:", network);

  const SDIDContract = await hre.ethers.getContractFactory(
    "SingleDecentralizedIdentityContract"
  );

  s_did_contract = await SDIDContract.connect(signer).deploy(
    "QmVsUcm2skKLz8nrfAz8q3TnkSvEaG9my6bL8gjiSE1WLD"
  );

  console.log(
    `SingleDecentralizedIdentityContract deployed to ${s_did_contract.target}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

//0x40848Eb18591c3f44c1B73E8F814F49774734749

const hre = require("hardhat");

async function main() {
  const [dummy, first, signer] = await hre.ethers.getSigners();
  const contractAddress = "0x0f7773CE819dFDF650a6b646E8d34aF63d5E40C4";
  const didaddress = "0x1543e98924055A0dDB5fF6777be18CB486041094";
  const DIDRegistryContract = await hre.ethers.getContractFactory(
    "DIDRegistryContract"
  );
  const did_registry_contract = await DIDRegistryContract.attach(
    contractAddress
  );
  const cname = "ETF-Beograd-BU";
  const tx = await did_registry_contract
    .connect(signer)
    .addNewIdentity(cname, didaddress);

  console.log(tx);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

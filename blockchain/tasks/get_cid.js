const fs = require("fs");
//const hre = require("hardhat");

async function getCID(hre, contractAddress, prefixName, suffixName) {
  try {
    //const [u1, u2] = await hre.ethers.getSigners();
    // console.log(u1.address);
    // console.log(u2.address);
    // const args = process.argv.slice(2);
    // console.log(args);
    // if (args.length < 2) {
    //   console.error(
    //     "Usage: npx hardhat run --network <your network> <path_to>/getCID.js <registry address> <prefixname> (<suffixname>)"
    //   );
    //   process.exit(1);
    // }

    // Get the first argument (assuming it's the contract address)
    // const contractAddress = args[0];

    // Retrieve the deployed contract using its address and ABI

    const DIDRegistryContract = await hre.ethers.getContractFactory(
      "DIDRegistryContract"
    );
    const did_registry_contract = await DIDRegistryContract.attach(
      contractAddress
    );

    if (suffixName != undefined) {
      console.log("Error: Not yet supported");
      return;
    } else {
      //string memory prefix name, string memory cid, uint version, bool is revoked, bool isError
      const returnVal = await did_registry_contract.getIdentity(prefixName);
      if (returnVal[4]) {
        console.log(`Error: ${returnVal[0]}`);
        return;
      } else {
        console.log(`{
                "cid" : "${returnVal[1]}",
                "version" : ${returnVal[2]},
                "isRevoked" : ${returnVal[3]}
            }`);

        return;
      }
    }
  } catch (error) {
    console.log(`Error: ${error.message}`);
  }
}

task("get_cid", "Get cid of a DID document")
  .addParam("contractaddress")
  .addParam("prefixname")
  .addParam("suffixname", "", undefined, undefined, true)
  // .addParam("suffixname")
  .setAction(async (taskArgs, hre) => {
    //  console.log("taskArgs", taskArgs);
    if (taskArgs.suffixName != undefined && taskArgs.suffixName == "") {
      taskArgs.suffixName = undefined;
    }
    await getCID(
      hre,
      taskArgs.contractaddress,
      taskArgs.prefixname,
      taskArgs.suffixname
    );
  });

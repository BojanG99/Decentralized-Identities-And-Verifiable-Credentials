require("dotenv").config();
const fs = require("fs");
const pinataSDK = require("@pinata/sdk");
const pinata = new pinataSDK({ pinataJWTKey: process.env.PINATA_JWT_KEY });
const args = process.argv.slice(2);

async function uploadFileToIPFS(filePath) {
  try {
    const data = fs.readFileSync(filePath, "utf8");
    const jsonData = JSON.parse(data);
    const result = await pinata.pinJSONToIPFS(jsonData);
    console.log(`File ${filePath} is uploaded to IPFS`);
    console.log(result);
    return result;
  } catch (error) {
    console.log(error);
  }
}

async function main() {
  if (args.length == 0) {
    console.log(
      "You need to pass one or more JSON file path to upload to IPFS"
    );
  }

  if (args[0] === "--help") {
    console.log("pass one or more json file to upload to IPFS");
    return;
  }

  for (const arg of args) {
    await uploadFileToIPFS(arg);
  }
}

main();

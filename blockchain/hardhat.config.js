require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
require("./tasks/get_cid");
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${process.env.API_KEY}`,
      accounts: [
        `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80`, //dummy account for reading from blockchain
        `0x${process.env.SEPOLIA_PRIVATE_KEY_1}`,
        `0x${process.env.SEPOLIA_PRIVATE_KEY_2}`,
      ],
    },
  },
};

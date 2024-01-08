require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${process.env.API_KEY}`,
      accounts: [
        `0x${process.env.SEPOLIA_PRIVATE_KEY_1}` /*,`0x${SEPOLIA_PRIVATE_KEY_2}`,`0x${SEPOLIA_PRIVATE_KEY_3}`*/,
      ],
    },
  },
};

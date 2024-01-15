const axios = require("axios");
const fs = require("fs");
const contentType = require("content-type");
require("dotenv").config();

const fileUrl = `${process.env.PINATA_GATAWAY}/ipfs/${process.env.CID}`; // Replace PINATA_GATAWAY and CID in .env file

function generateRandomString(length) {
  const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
  let randomString = "";

  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    randomString += characters.charAt(randomIndex);
  }

  return randomString;
}

const destinationPath = `downloads/file_${generateRandomString(5)}`;

async function downloadFile() {
  try {
    const response = await axios.get(fileUrl, { responseType: "arraybuffer" });

    //   console.log(response);
    const type = response.headers["content-type"];
    console.log(type);
    // Extract the file extension from the MIME type
    const fileExtension = type.split("/")[1];

    fs.writeFileSync(`${destinationPath}.${fileExtension}`, response.data);

    // try {
    //   JSON.parse(response.data);
    //   console.log("JSON");
    // } catch (error) {
    //   console.log("NOT JSON");
    // }

    console.log("File downloaded successfully.");
  } catch (error) {
    console.error("Error downloading file:", error.message);
  }
}

// Call the function to download the file
downloadFile();

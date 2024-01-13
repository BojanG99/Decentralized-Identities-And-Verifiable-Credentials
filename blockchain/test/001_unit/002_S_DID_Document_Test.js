const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");

describe("SingleDecentralizedIdentityContract", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployContract() {
    const [owner, otherAccount] = await ethers.getSigners();

    const SingleDecentralizedIdentityContract = await ethers.getContractFactory(
      "SingleDecentralizedIdentityContract"
    );
    const did_doc_hash = "jafklr9f0avqj039t3063";
    const single_did_contract =
      await SingleDecentralizedIdentityContract.deploy(did_doc_hash);
    return { single_did_contract, owner, otherAccount, did_doc_hash };
  }

  async function deployVCStorage() {
    const [owner] = await ethers.getSigners();

    const VCStorageContract = await ethers.getContractFactory(
      "VCStorageContract"
    );
    const vc_storage_contract = await VCStorageContract.deploy(
      "0x0000000000000000000000000000000000000000"
    );
    return { vc_storage_contract, owner };
  }

  describe("Deployment", function () {
    it("Should set creator as owner", async function () {
      const { single_did_contract, owner } = await loadFixture(deployContract);

      expect(await single_did_contract.owner()).to.equal(owner.address);
    });

    it("Should set correct did document hash", async function () {
      const { single_did_contract, did_doc_hash } = await loadFixture(
        deployContract
      );

      const identity = await single_did_contract.getIdentity();

      //did document hash
      expect(identity[0]).to.equal(did_doc_hash);
      //version
      expect(identity[1]).to.equal(1);
      //dis revoked
      expect(identity[2]).to.equal(false);
    });
  });

  describe("Setting a VC Storage", function () {
    it("Should set a VC storage (already exists)", async function () {
      const { single_did_contract, owner } = await loadFixture(deployContract);
      const { vc_storage_contract } = await loadFixture(deployVCStorage);
      // console.log(vc_storage_contract.target);
      await single_did_contract.setVCStorage(vc_storage_contract.target);
      //  console.log(await single_did_contract.vc_storage());
      expect(await single_did_contract.vc_storage()).to.equal(
        vc_storage_contract.target
      );
    });

    it("Should set a VC storage (created from SDIDcontract)", async function () {
      const { single_did_contract } = await loadFixture(deployContract);

      await single_did_contract.createVCStorage();
      expect(await single_did_contract.vc_storage()).to.not.equal(
        "0x0000000000000000000000000000000000000000"
      );
    });

    it("Should not set an address for vc storage if that address is not vc contract", async function () {
      const { single_did_contract } = await loadFixture(deployContract);

      try {
        await single_did_contract.setVCStorage(single_did_contract.target);
      } catch (error) {
        expect(error.message.includes("NotVCStorageAddress")).to.equal(true);
        return;
      }

      throw new Error("Expected error didn't occure.");
    });

    it("Should not set vc storage from non owner", async function () {
      const { single_did_contract, owner, otherAccount } = await loadFixture(
        deployContract
      );
      const { vc_storage_contract } = await loadFixture(deployVCStorage);
      try {
        await single_did_contract
          .connect(otherAccount)
          .setVCStorage(vc_storage_contract.target);
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }

      throw new Error("Expected error didn't occure.");
    });

    it("Should not reassign vc storage address", async function () {
      const { single_did_contract } = await loadFixture(deployContract);

      await single_did_contract.createVCStorage();
      const { vc_storage_contract } = await loadFixture(deployVCStorage);
      try {
        await single_did_contract.setVCStorage(vc_storage_contract.target);
      } catch (error) {
        expect(error.message.includes("VCStorageIsSet")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });

    it("Should not reassign vc storage address", async function () {
      const { single_did_contract } = await loadFixture(deployContract);

      await single_did_contract.createVCStorage();
      try {
        await single_did_contract.createVCStorage();
      } catch (error) {
        expect(error.message.includes("VCStorageIsSet")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Setting new DID Document", function () {
    it("Should set new DID Document with incremented version", async function () {
      const { single_did_contract, owner } = await loadFixture(deployContract);
      const newDidDocument = "akgjn913mq0jf10-13f1";
      await single_did_contract.setNewDIDDocument(newDidDocument);
      did_document = await single_did_contract.getIdentity();
      expect(did_document[0]).to.equal(newDidDocument);
      expect(did_document[1]).to.equal(2);
      expect(did_document[2]).to.equal(false);
    });
    it("Should not allow non owner to set new DID Document", async function () {
      const { single_did_contract, otherAccount } = await loadFixture(
        deployContract
      );
      const newDidDocument = "akgjn913mq0jf10-13f1";

      try {
        await single_did_contract
          .connect(otherAccount)
          .setNewDIDDocument(newDidDocument);
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Revoking DID Document", function () {
    it("Should revoke did document", async function () {
      const { single_did_contract } = await loadFixture(deployContract);
      await single_did_contract.revokeDIDDocument();

      expect((await single_did_contract.getIdentity())[2]).to.equal(true);
    });
    it("Should not allow non owner to revoke did document", async function () {
      const { single_did_contract, otherAccount } = await loadFixture(
        deployContract
      );

      try {
        await single_did_contract.connect(otherAccount).revokeDIDDocument();
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });
});

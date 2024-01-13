const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");

describe("DecentralizedIdentityRegistryContract", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployContract() {
    const [owner, otherAccount] = await ethers.getSigners();

    const DIDRegistryContract = await ethers.getContractFactory(
      "DIDRegistryContract"
    );
    const did_registry_contract = await DIDRegistryContract.deploy();
    return { did_registry_contract, owner, otherAccount };
  }

  async function deploySDIDContract() {
    const [owner, otherAccount] = await ethers.getSigners();
    const cid = "0a986vqe56gq5ga96hq96ag";
    const SDIDContract = await ethers.getContractFactory(
      "SingleDecentralizedIdentityContract"
    );
    const s_did_contract = await SDIDContract.deploy(cid);
    return { s_did_contract, cid };
  }

  async function deploySDIDContract1() {
    const [owner, otherAccount] = await ethers.getSigners();
    const cid = "0a986vqe56gq5ga96hq96ag";
    const SDIDContract = await ethers.getContractFactory(
      "SingleDecentralizedIdentityContract"
    );
    const s_did_contract1 = await SDIDContract.deploy(cid);
    return { s_did_contract1 };
  }

  describe("Add new identity", function () {
    it("Should add new decentralized identity with unique name", async function () {
      const { did_registry_contract, owner } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";

      expect(await did_registry_contract.used_names(name)).to.equal(false);
      await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      expect(await did_registry_contract.used_names(name)).to.equal(true);
      expect(await did_registry_contract.owners(name)).to.equal(owner.address);
    });

    it("Should not add new decentralized identity if name is used", async function () {
      const { did_registry_contract, owner } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";

      await did_registry_contract.addNewIdentity(name, s_did_contract.target);

      try {
        await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      } catch (error) {
        expect(error.message.includes("NameIsUsed")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Delete identity", function () {
    it("Should allow owner to delete existing identity", async function () {
      const { did_registry_contract, owner } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";

      await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      await did_registry_contract.deleteIdentity(name);
      expect(await did_registry_contract.used_names(name)).to.equal(false);
      expect(await did_registry_contract.owners(name)).to.equal(
        "0x0000000000000000000000000000000000000000"
      );
    });

    it("Should not allow non owner to delete existing identity", async function () {
      const { did_registry_contract, owner, otherAccount } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";

      await did_registry_contract.addNewIdentity(name, s_did_contract.target);

      try {
        await did_registry_contract.connect(otherAccount).deleteIdentity(name);
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Change identity", function () {
    it("Should allow owner to change his identity", async function () {
      const { did_registry_contract, owner } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const { s_did_contract1 } = await loadFixture(deploySDIDContract1);
      const name = "Bojan Galic";

      await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      await did_registry_contract.changeIdentity(name, s_did_contract1.target);
      expect(await did_registry_contract.identities(name)).to.equal(
        s_did_contract1.target
      );
    });

    it("Should not allow non owner to change identity", async function () {
      const { did_registry_contract, owner, otherAccount } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";
      const { s_did_contract1 } = await loadFixture(deploySDIDContract1);
      await did_registry_contract.addNewIdentity(name, s_did_contract.target);

      try {
        await did_registry_contract
          .connect(otherAccount)
          .changeIdentity(name, s_did_contract1.target);
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Change name ownership", function () {
    it("Should allow owner to transfer name to other address", async function () {
      const { did_registry_contract, owner, otherAccount } = await loadFixture(
        deployContract
      );
      const { s_did_contract } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";

      await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      await did_registry_contract.changeOwnership(name, otherAccount.address);
      expect(await did_registry_contract.identities(name)).to.equal(
        "0x0000000000000000000000000000000000000000"
      );
      expect(await did_registry_contract.owners(name)).to.equal(
        otherAccount.address
      );
    });
  });

  describe("Get identity", function () {
    it("Should allow users to get identity for specific prefix name", async function () {
      const { did_registry_contract, owner, otherAccount } = await loadFixture(
        deployContract
      );
      const { s_did_contract, cid } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";

      await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      const identity = await did_registry_contract.getIdentity(name);

      expect(identity[4]).to.equal(false);
      expect(identity[0]).to.equal(name);
      expect(identity[1]).to.equal(cid);
      expect(identity[2]).to.equal(1);
      expect(identity[3]).to.equal(false);
    });

    it("Should not allow users to get identiti for unused prefix name", async function () {
      const { did_registry_contract, owner, otherAccount } = await loadFixture(
        deployContract
      );
      const { s_did_contract, cid } = await loadFixture(deploySDIDContract);
      const name = "Bojan Galic";
      const name1 = "Bojan Galic1";
      await did_registry_contract.addNewIdentity(name, s_did_contract.target);
      try {
        await did_registry_contract.getIdentity(name1);
      } catch (error) {
        expect(error.message.includes("NoIdentityWithPrefix")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });
});

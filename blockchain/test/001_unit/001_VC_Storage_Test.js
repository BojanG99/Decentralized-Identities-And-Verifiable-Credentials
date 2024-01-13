const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");

describe("VCStorageContract", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployAsUser() {
    const [owner, otherAccount] = await ethers.getSigners();

    const VCStorageContract = await ethers.getContractFactory(
      "VCStorageContract"
    );
    const vc_storage_contract = await VCStorageContract.deploy(
      "0x0000000000000000000000000000000000000000"
    );
    return { vc_storage_contract, owner, otherAccount };
  }

  async function deployAsContract() {
    // Contracts are deployed using the first signer/account by default
    const [owner, contract_addr] = await ethers.getSigners();
    const VCStorageContract = await ethers.getContractFactory(
      "VCStorageContract"
    );

    const constructorArguments = [owner.address];

    const vc_storage_contract = await VCStorageContract.connect(
      contract_addr
    ).deploy(...constructorArguments);
    return { vc_storage_contract, owner };
  }

  describe("Deployment", function () {
    it("Should set creator as owner", async function () {
      const { vc_storage_contract, owner } = await loadFixture(deployAsUser);

      expect(await vc_storage_contract.owner()).to.equal(owner.address);
    });

    it("Should set the owner from costructor", async function () {
      const { vc_storage_contract, owner } = await loadFixture(
        deployAsContract
      );

      expect(await vc_storage_contract.owner()).to.equal(owner.address);
    });
  });

  describe("Issuing Verifiable Credentials", function () {
    it("Should issue new verifiable credential without ipfs metadata", async function () {
      const { vc_storage_contract, owner } = await loadFixture(deployAsUser);
      const vc_hash = "myTestHash";
      //  const vc_hash1 = "myTestHash1";
      const is_revocable = true;

      await vc_storage_contract.issueNewCredential(vc_hash, is_revocable);

      const vc = await vc_storage_contract.getVCredentials(vc_hash);
      expect(vc[0]).to.equal("");
      expect(vc[1]).to.equal("");
      expect(vc[2]).to.equal(is_revocable);
      expect(vc[3]).to.equal(false);
    });

    it("Should issue new verifiable credential with ipfs metadata", async function () {
      const { vc_storage_contract, owner } = await loadFixture(deployAsUser);
      const vc_hash = "myTestHash";
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";
      //  const vc_hash1 = "myTestHash1";
      const is_revocable = true;

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );

      const vc = await vc_storage_contract.getVCredentials(vc_hash);
      expect(vc[0]).to.equal(ipfs_vc_hash);
      expect(vc[1]).to.equal(hashed_password);
      expect(vc[2]).to.equal(is_revocable);
      expect(vc[3]).to.equal(false);
    });

    it("Should be reverted if vc_hash is already issued", async function () {
      const { vc_storage_contract, owner } = await loadFixture(deployAsUser);
      const vc_hash = "myTestHash";
      //  const vc_hash1 = "myTestHash1";
      const is_revocable = true;

      await vc_storage_contract.issueNewCredential(vc_hash, is_revocable);

      const vc = await vc_storage_contract.getVCredentials(vc_hash);

      try {
        await vc_storage_contract.issueNewCredential(vc_hash, is_revocable);
      } catch (error) {
        expect(error.message.includes("VCAlreadyIssued")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });

    it("Should be reverted if vc_hash is already issued", async function () {
      const { vc_storage_contract, owner } = await loadFixture(deployAsUser);
      const vc_hash = "myTestHash";
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";
      //  const vc_hash1 = "myTestHash1";
      const is_revocable = true;

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );

      try {
        await vc_storage_contract.issueNewCredential(
          vc_hash,
          is_revocable,
          ipfs_vc_hash,
          hashed_password
        );
      } catch (error) {
        //  console.log(error.message);
        expect(error.message.includes("VCAlreadyIssued")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });

    it("Should not allow non-owner to issue vc without ipfs metadata", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = true;
      try {
        await vc_storage_contract
          .connect(otherAccount)
          .issueNewCredential(vc_hash, is_revocable);
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });

    it("Should not allow non-owner to issue vc with ipfs metadata", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = true;
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";
      try {
        await vc_storage_contract
          .connect(otherAccount)
          .issueNewCredential(
            vc_hash,
            is_revocable,
            ipfs_vc_hash,
            hashed_password
          );
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }
      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Revoking Verifiable Credentials", function () {
    it("Should revoke revocable vc", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = true;
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );

      await vc_storage_contract.revokeCredential(vc_hash);

      expect(await vc_storage_contract.isRevoked(vc_hash)).to.equal(true);
    });

    it("Should not revoke vc as non owner", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = true;
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );

      try {
        await vc_storage_contract
          .connect(otherAccount)
          .revokeCredential(vc_hash);
      } catch (error) {
        expect(error.message.includes("NotTheOwner")).to.equal(true);
        return;
      }

      throw new Error("Expected error didn't occure.");
    });

    it("Should not revoke nonrevocable vc", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = false;
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );

      try {
        await vc_storage_contract.revokeCredential(vc_hash);
      } catch (error) {
        expect(error.message.includes("UnrevocableVC")).to.equal(true);
        return;
      }

      throw new Error("Expected error didn't occure.");
    });
  });

  describe("Getting Verifiable Credentials", function () {
    it("Should get correct status of vc", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = true;
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );

      expect(await vc_storage_contract.isRevoked(vc_hash)).to.equal(false);

      await vc_storage_contract.revokeCredential(vc_hash);

      expect(await vc_storage_contract.isRevoked(vc_hash)).to.equal(true);
    });
    it("Should be reverted if vc_hash is not in the vc storage", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );
      const vc_hash = "myTestHash";
      const is_revocable = true;
      const ipfs_vc_hash = "cidtocryptedvc";
      const hashed_password = "hashedpassword";

      await vc_storage_contract.issueNewCredential(
        vc_hash,
        is_revocable,
        ipfs_vc_hash,
        hashed_password
      );
      const _vc_hash = "none";
      try {
        await vc_storage_contract.isRevoked(_vc_hash);
      } catch (error) {
        expect(error.message.includes("VCDoesNotExists")).to.equal(true);
        return;
      }

      throw new Error("Expected error didn't occure.");
    });
  });
  describe("Test functions", function () {
    it("Should return 1", async function () {
      const { vc_storage_contract, owner, otherAccount } = await loadFixture(
        deployAsUser
      );

      expect(await vc_storage_contract.testForAddress()).to.equal(1);
    });
  });
});

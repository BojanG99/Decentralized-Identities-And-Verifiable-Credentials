const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");

describe("Integration test", function () {
  async function deployRegistry() {
    const [owner] = await ethers.getSigners();

    const DIDRegistryContract = await ethers.getContractFactory(
      "DIDRegistryContract"
    );
    const did_registry_contract = await DIDRegistryContract.deploy();
    return did_registry_contract;
  }
  async function deploySDIDContract(owner, ipfs_did_doc_hash) {
    const SingleDecentralizedIdentityContract = await ethers.getContractFactory(
      "SingleDecentralizedIdentityContract"
    );
    const single_did_contract =
      await SingleDecentralizedIdentityContract.connect(owner).deploy(
        ipfs_did_doc_hash
      );
    return single_did_contract;
  }
  async function deployVCStorageContract(owner) {
    const VCStorageContract = await ethers.getContractFactory(
      "VCStorageContract"
    );
    const vc_storage_contract = await VCStorageContract.connect(owner).deploy(
      "0x0000000000000000000000000000000000000000"
    );
    return vc_storage_contract;
  }

  async function loadVCStorage(vc_storage_address) {
    const VCStorageContract = await ethers.getContractFactory(
      "VCStorageContract"
    );
    const vc_storage_contract = await VCStorageContract.attach(
      vc_storage_address
    );

    return vc_storage_contract;
  }

  it("Test 1", async function () {
    const did_registry_contract = await deployRegistry();
    const [user1, user2, user3, user4, user5, user6] =
      await ethers.getSigners();

    const user1_did_hash =
      "0a041b9462caa4a31bac3567e0b6e6fd9100787db2ab433d96f6d178cabfce90";
    const user2_did_hash =
      "6025d18fe48abd45168528f18a82e265dd98d421a7084aa09f61b341703901a3";
    const user3_did_hash =
      "5860faf02b6bc6222ba5aca523560f0e364ccd8b67bee486fe8bf7c01d492ccb";
    const user4_did_hash =
      "5269ef980de47819ba3d14340f4665262c41e933dc92c1a27dd5d01b047ac80e";
    const user5_did_hash =
      "5a39bead318f306939acb1d016647be2e38c6501c58367fdb3e9f52542aa2442";
    const user6_did_hash =
      "ecb48a1cc94f951252ec462fe9ecc55c3ef123fadfe935661396c26a45a5809d";

    sdid_contract1 = await deploySDIDContract(user1, user1_did_hash);
    sdid_contract2 = await deploySDIDContract(user2, user2_did_hash);
    sdid_contract3 = await deploySDIDContract(user3, user3_did_hash);
    sdid_contract4 = await deploySDIDContract(user4, user4_did_hash);
    sdid_contract5 = await deploySDIDContract(user5, user5_did_hash);
    sdid_contract6 = await deploySDIDContract(user6, user6_did_hash);

    const vc_storage1 = await deployVCStorageContract(user1);
    const vc_storage2 = await deployVCStorageContract(user2);
    const vc_storage3 = await deployVCStorageContract(user3);

    await sdid_contract1.setVCStorage(vc_storage1.target);
    await sdid_contract2.setVCStorage(vc_storage2.target);
    await sdid_contract3.setVCStorage(vc_storage3.target);
    await sdid_contract4.createVCStorage();
    await sdid_contract5.createVCStorage();
    await sdid_contract6.createVCStorage();

    const vc_storage4_address = await sdid_contract4.vc_storage();
    const vc_storage5_address = await sdid_contract5.vc_storage();
    const vc_storage6_address = await sdid_contract6.vc_storage();

    const tx1 = await did_registry_contract
      .connect(user1)
      .addNewIdentity("User 1", sdid_contract1.target);
    await did_registry_contract
      .connect(user2)
      .addNewIdentity("User 2", sdid_contract2.target);
    await did_registry_contract
      .connect(user3)
      .addNewIdentity("User 3", sdid_contract3.target);
    await did_registry_contract
      .connect(user4)
      .addNewIdentity("User 4", sdid_contract4.target);
    await did_registry_contract
      .connect(user5)
      .addNewIdentity("User 5", sdid_contract5.target);
    const tx6 = await did_registry_contract
      .connect(user6)
      .addNewIdentity("User 6", sdid_contract6.target);

    const identity1 = await did_registry_contract.getIdentity("User 1");
    const identity2 = await did_registry_contract.getIdentity("User 2");
    const identity3 = await did_registry_contract.getIdentity("User 3");
    const identity4 = await did_registry_contract.getIdentity("User 4");
    const identity5 = await did_registry_contract.getIdentity("User 5");
    const identity6 = await did_registry_contract.getIdentity("User 6");

    expect(identity1[1]).to.equal(user1_did_hash);
    expect(identity2[1]).to.equal(user2_did_hash);
    expect(identity3[1]).to.equal(user3_did_hash);
    expect(identity4[1]).to.equal(user4_did_hash);
    expect(identity5[1]).to.equal(user5_did_hash);
    expect(identity6[1]).to.equal(user6_did_hash);

    try {
      await did_registry_contract
        .connect(user5)
        .addNewIdentity("User 5", sdid_contract5.target);
      throw new Error("Expected error didn't occure.");
    } catch (error) {
      expect(error.message.includes("NameIsUsed")).to.equal(true);
    }

    expect(await did_registry_contract.identities("User 1")).to.equal(
      sdid_contract1.target
    );
    expect(await did_registry_contract.identities("User 2")).to.equal(
      sdid_contract2.target
    );
    expect(await did_registry_contract.identities("User 3")).to.equal(
      sdid_contract3.target
    );
    expect(await did_registry_contract.identities("User 4")).to.equal(
      sdid_contract4.target
    );
    expect(await did_registry_contract.identities("User 5")).to.equal(
      sdid_contract5.target
    );
    expect(await did_registry_contract.identities("User 6")).to.equal(
      sdid_contract6.target
    );
    expect(await did_registry_contract.identities("User 7")).to.equal(
      "0x0000000000000000000000000000000000000000"
    );
    try {
      await did_registry_contract
        .connect(user1)
        .changeIdentity("User 6", sdid_contract1.target);
      throw new Error("Expected error didn't occure.");
    } catch (error) {
      expect(error.message.includes("NotTheOwner()")).to.equal(true);
    }

    await did_registry_contract
      .connect(user6)
      .changeOwnership("User 6", user1.address);
    await did_registry_contract
      .connect(user1)
      .changeIdentity("User 6", sdid_contract1.target);
    expect(await did_registry_contract.identities("User 6")).to.equal(
      sdid_contract1.target
    );
    //ISUE VERIFIABLE CREDENTIAL
    //vc data
    const vc_hash = "";
    const ipfs_vc_hash = "";
    const ipfs_password = "";
    const revocable = true;

    vc_storage4 = await loadVCStorage(vc_storage4_address);
    await vc_storage4
      .connect(user4)
      .issueNewCredential(vc_hash, revocable, ipfs_vc_hash, ipfs_password);

    expect(await vc_storage4.isRevoked(vc_hash)).to.equal(false);
  });
});

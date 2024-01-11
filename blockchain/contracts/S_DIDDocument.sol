// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./VC_Storage.sol";

error VCStorageIsSet();
error NotVCStorageAddress();

contract SingleDecentralizedIdentityContract {
    string private did_document_cid;
    uint private version;
    bool private is_revoked;
    address public owner;

    address public vc_storage = 0x0000000000000000000000000000000000000000;
    bool private vcstorageSet = false;

    //   event DIDChanged(address indexed owner, string name, string did_cid, bool is_revoked);
    //   event IdentityCreated(address indexed owner, string name,string did_cid);

    modifier onlyOnce() {
        if (vcstorageSet) {
            revert VCStorageIsSet();
        }
        _;
        vcstorageSet = true;
    }

    modifier onlyOwner() {
        if (msg.sender != owner) {
            revert NotTheOwner();
        }
        _;
    }

    constructor(string memory _did_document_cid) {
        did_document_cid = _did_document_cid;
        version = 1;
        owner = msg.sender;
        is_revoked = false;
        //     emit IdentityCreated(msg.sender, _name,did_document_id);
    }

    function setVCStorage(
        address vc_storage_address
    ) external onlyOnce onlyOwner {
        if (isAddressOfVCStorageContract(vc_storage_address)) {
            vc_storage = vc_storage_address;
        } else {
            revert NotVCStorageAddress();
        }
    }

    function createVCStorage() external onlyOnce onlyOwner {
        vc_storage = address(new VCStorageContract(owner));
    }

    function setNewDIDDocument(
        string memory _did_document_cid
    ) external onlyOwner {
        //     emit DIDChanged(owner, name, did_document_cid, is_revoked);
        did_document_cid = _did_document_cid;
        version = version + 1;
        is_revoked = false;
    }

    function revokeDIDDocument() external onlyOwner {
        is_revoked = true;
    }

    function getIdentity() external view returns (string memory, uint, bool) {
        return (did_document_cid, version, is_revoked);
    }

    function isAddressOfVCStorageContract(
        address _addressToCheck
    ) internal pure returns (bool) {
        try VCStorageContract(_addressToCheck).testForAddress() returns (
            uint256
        ) {
            // If the function call is successful, it's an instance of ContractA
            return true;
        } catch {
            // If the function call fails, it's not an instance of ContractA
            return false;
        }
    }
}

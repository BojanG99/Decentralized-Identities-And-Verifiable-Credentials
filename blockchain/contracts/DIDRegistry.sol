// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./S_DIDDocument.sol";
//import "./M_DIDDocument.sol";

error NameIsUsed();
error NoIdentityWithPrefix();

contract DIDRegistryContract {
    mapping(string => address) public identities;
    mapping(string => bool) public used_names;
    mapping(string => address) owners;

    function addNewIdentity(
        string memory prefix_name,
        address did_document_contract_address
    ) external {
        if (used_names[prefix_name]) {
            revert NameIsUsed();
        }

        owners[prefix_name] = msg.sender;
        identities[prefix_name] = did_document_contract_address;
        used_names[prefix_name] = true;
    }

    function deleteIdentity(string memory prefix_name) external {
        if (owners[prefix_name] != msg.sender) {
            revert NotTheOwner();
        }

        used_names[prefix_name] = false;
        delete owners[prefix_name];
        delete identities[prefix_name];
    }

    function changeOwnership(
        string memory prefix_name,
        address newOwner
    ) public {
        if (owners[prefix_name] != msg.sender) {
            revert NotTheOwner();
        }

        owners[prefix_name] = newOwner;
        delete identities[prefix_name];
    }

    function changeIdentity(
        string memory prefix_name,
        address did_document_contract_address
    ) external {
        if (owners[prefix_name] != msg.sender) {
            revert NotTheOwner();
        }

        identities[prefix_name] = did_document_contract_address;
    }

    /*
function getIdentity(string memory prefix_name,string memory sufix_name) public view returns(string memory, string memory, uint, bool, bool) {
    
    if (used_names[prefix_name] == false) {

        revert NoIdentityWithPrefix();

    }

   address contractAddr = identities[prefix_name];

    if (isContract(contractAddr)) {

        try MultiDecentralizedIdentityContract(contractAddr).getIdentity(sufix_name) returns (string memory name,string memory cid,uint version,bool is_revoked) {
                return (name, cid, version, is_revoked, false);
            } catch {
                return ("Unknown Contract","",0,false, true);
        }
    } 

    return ("Not a contract address","",0,false, true);

}*/

    function getIdentity(
        string memory prefix_name
    ) public view returns (string memory, string memory, uint, bool, bool) {
        if (used_names[prefix_name] == false) {
            revert NoIdentityWithPrefix();
        }

        address contractAddr = identities[prefix_name];
        if (isContract(contractAddr)) {
            try
                SingleDecentralizedIdentityContract(contractAddr).getIdentity()
            returns (string memory cid, uint version, bool is_revoked) {
                return (prefix_name, cid, version, is_revoked, false);
            } catch {
                return ("Unknown Contract", "", 0, false, true);
            }
        }

        return ("Not a contract address", "", 0, false, true);
    }

    function isContract(address addr) internal view returns (bool) {
        uint256 size;
        assembly {
            size := extcodesize(addr)
        }
        return size > 0;
    }
}

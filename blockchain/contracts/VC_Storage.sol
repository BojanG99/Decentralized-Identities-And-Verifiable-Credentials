// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

error NotTheOwner();
error VCDoesNotExists();
error UnrevocableVC();

contract VCStorageContract {
    struct VerifiableCredential {
        string opt_vc_ipfs_hash; //optional
        string opt_locked_password; //optional
        bool revocable;
        bool is_revoked;
    }

    mapping(string => VerifiableCredential) vCredentials;
    mapping(string => bool) is_issued;
    string[] all_VCredentials;

    address public owner;

    constructor(address _owner) {
        if (_owner == 0x0000000000000000000000000000000000000000) {
            owner = msg.sender;
        } else {
            owner = _owner;
        }
    }

    // constructor(address _owner){
    //     owner = _owner;
    // }

    modifier onlyOwner() {
        if (owner != msg.sender) {
            revert NotTheOwner();
        }
        _;
    }

    function issueNewCredential(
        string memory _vc_hash,
        string memory _vc_ipfs_hash,
        string memory locked_password,
        bool is_revocable
    ) public onlyOwner {
        vCredentials[_vc_hash] = VerifiableCredential(
            _vc_ipfs_hash,
            locked_password,
            is_revocable,
            false
        );
        all_VCredentials.push(_vc_hash);
    }

    function revokeCredential(string memory _vc_hash) public onlyOwner {
        if (vCredentials[_vc_hash].revocable == false) {
            revert UnrevocableVC();
        }

        if (is_issued[_vc_hash] == false) {
            revert VCDoesNotExists();
        }

        vCredentials[_vc_hash].is_revoked = true;
    }

    function getVCredentials(
        string memory _vc_hash
    ) external view onlyOwner returns (string memory, string memory, bool) {
        if (is_issued[_vc_hash] == false) {
            revert VCDoesNotExists();
        }

        return (
            vCredentials[_vc_hash].opt_vc_ipfs_hash,
            vCredentials[_vc_hash].opt_locked_password,
            vCredentials[_vc_hash].is_revoked
        );
    }

    function isRevoked(string memory _vc_hash) external view returns (bool) {
        if (is_issued[_vc_hash] == false) {
            revert VCDoesNotExists();
        }

        return vCredentials[_vc_hash].is_revoked;
    }

    function testForAddress() public pure returns (uint) {
        return 100;
    }
}

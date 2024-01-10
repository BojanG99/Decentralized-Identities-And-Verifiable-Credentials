// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

error NotTheOwner();
error VCDoesNotExists();
error UnrevocableVC();
error VCAlreadyIssued();

contract VCStorageContract {
    struct VerifiableCredential {
        string opt_vc_ipfs_hash; //optional
        string opt_locked_password; //optional
        bool revocable;
        bool is_revoked;
    }

    mapping(string => VerifiableCredential) private vCredentials;
    mapping(string => bool) private is_issued;
    string[] private all_VCredentials;

    address public owner;

    constructor(address _owner) {
        if (_owner == 0x0000000000000000000000000000000000000000) {
            owner = msg.sender;
        } else {
            owner = _owner;
        }
    }

    modifier onlyOwner() {
        if (owner != msg.sender) {
            revert NotTheOwner();
        }
        _;
    }

    function issueNewCredential(
        string memory _vc_hash,
        bool is_revocable,
        string memory _vc_ipfs_hash,
        string memory locked_password
    ) public onlyOwner {
        if (is_issued[_vc_hash]) {
            revert VCAlreadyIssued();
        }
        vCredentials[_vc_hash] = VerifiableCredential(
            _vc_ipfs_hash,
            locked_password,
            is_revocable,
            false
        );
        is_issued[_vc_hash] = true;
        all_VCredentials.push(_vc_hash);
    }

    function issueNewCredential(
        string memory _vc_hash,
        bool is_revocable
    ) external onlyOwner {
        if (is_issued[_vc_hash]) {
            revert VCAlreadyIssued();
        }

        vCredentials[_vc_hash] = VerifiableCredential(
            "",
            "",
            is_revocable,
            false
        );
        is_issued[_vc_hash] = true;
        all_VCredentials.push(_vc_hash);
    }

    function revokeCredential(string memory _vc_hash) external onlyOwner {
        if (is_issued[_vc_hash] == false) {
            revert VCDoesNotExists();
        }
        if (vCredentials[_vc_hash].revocable == false) {
            revert UnrevocableVC();
        }

        vCredentials[_vc_hash].is_revoked = true;
    }

    function getVCredentials(
        string memory _vc_hash
    )
        external
        view
        onlyOwner
        returns (string memory, string memory, bool, bool)
    {
        if (is_issued[_vc_hash] == false) {
            revert VCDoesNotExists();
        }

        return (
            vCredentials[_vc_hash].opt_vc_ipfs_hash,
            vCredentials[_vc_hash].opt_locked_password,
            vCredentials[_vc_hash].revocable,
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
        return 1;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity >=0.4.16 <0.7.0;

contract LogLocationProof {
    modifier onlyBy(address _account) {
        require(
            msg.sender == _account,
            "Sender not authorized to update this mapping!"
        );

        _; // The "_;"! will be replaced by the actual function body when the modifier is used.
    }

    // DeviceID_LatestLocationHistoryCID is a mapping which persists the latest CIDs for each sensor identified by its own ID
    mapping(address => string) private deviceIDToLatestCID;

    // MappingUpdated is emitted when the mapping is udpated with a new ID by sensors
    event MappingUpdated(address deviceID, string latestCID);

    // SetLatestCID is a setter function to udpate the mapping with the latest CID
    function setLatestCID(address deviceID, string memory latestCID)
        public
        onlyBy(deviceID)
    {
        deviceIDToLatestCID[deviceID] = latestCID;
        emit MappingUpdated(deviceID, latestCID);
    }

    // GetLatestCID returns the latest CID of a given sensor
    function getLatestCID(address deviceID)
        public
        view
        returns (string memory latestCID)
    {
        return deviceIDToLatestCID[deviceID];
    }
}

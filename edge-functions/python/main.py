import requests
import os
import datetime
import calendar
import time

Tracking_ID = ""

url = 'https://www.googleapis.com/geolocation/v1/geolocate'
myobj = {'key': os.environ['GMAPS_API_KEY']}

moibit_url = 'https://kfs2.moibit.io/moibit/v0/storageused'
header_obj = {
    'api_key': os.environ['MOIBIT_API_KEY'],
    'api_secret': os.environ['MOIBIT_API_SECRET'],
    'content-type': "application/json"
}


def WriteDataTOIPFS():
    # TODO: Implement MoiBit updateable writes
    res = requests.get(moibit_url, headers=header_obj)
    print(res.json())


def getGeoCordinates():
    res = requests.post(url, data=myobj)
    geoCordinates = res.json()['location']
    accuracy = res.json()['accuracy']
    print("Latitude: %s" % geoCordinates['lat'])
    print("Longitude: %s" % geoCordinates['lng'])
    print("Accuracy of the location: %s (in metres)" % accuracy)
    return (geoCordinates['lat'], geoCordinates['lng'], accuracy)


def getCurrentTime():
    dt = datetime.datetime.utcnow()

    timestamp = time.mktime(dt.timetuple())
    print("Unix Timestamp: %d" % timestamp)
    return timestamp


def createID():
    Tracking_ID = "a"
    # TODO: Implement the creation of new ID using smart contract and persist it locally
    return Tracking_ID


def Marshal(id, lat, long, accuracy, timestamp):
    # TODO: Implement JSON Marshal to create a JSON object
    pass


def CommitTxn(id, cid):
    # TODO: Implement contract call
    pass


def main():
    # Fetching the Tracking ID locally, or generating a new one
    global Tracking_ID
    if Tracking_ID is "":
        Tracking_ID = createID()
    print("Generated a new tracking ID: ", Tracking_ID)

    # Getting the current geo-coordinates of the device
    (lat, long, accuracy) = getGeoCordinates()
    print("Received location data: ", lat, long, accuracy)

    # Get the UTC based Unix timestamp of the device
    timestamp = getCurrentTime()
    print("Got the current timestamp: ", timestamp)

    # Generate the JSON structure
    # jsonData = Marshal(Tracking_ID, lat, long, accuracy, timestamp)

    # Write the entry to IPFS
    # cid = WriteDataToIPFS(jsonData)

    # Publish the proof to Ethereum
    # txnHash = CommitTxn(Tracking_ID, cid)


if __name__ == "__main__":
    main()
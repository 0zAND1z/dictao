import os
import sys
import flask
from flask import request, jsonify
from web3 import Web3
import web3
import json
import http
from flask_cors import CORS
import pandas as pd
from sklearn.cluster import DBSCAN


app = flask.Flask(__name__)
CORS(app)
app.config["DEBUG"] = True

blockchain_url = 'https://kovan.infura.io/v3/' + \
    os.environ['WEB3_INFURA_PROJECT_ID']
abi = """[{"anonymous": false,"inputs": [{"indexed": false,"internalType": "address","name": "deviceID","type": "address"},{"indexed": false,"internalType": "string","name": "latestCID","type": "string"}],"name": "MappingUpdated","type": "event"},{"inputs": [{"internalType": "address","name": "deviceID","type": "address"},{"internalType": "string","name": "latestCID","type": "string"}],"name": "setLatestCID","outputs": [],"stateMutability": "nonpayable","type": "function"},{"inputs": [],"name": "getDeviceIDsLength","outputs": [{"internalType": "uint256","name": "","type": "uint256"}],"stateMutability": "view","type": "function"},{"inputs": [{"internalType": "uint256","name": "index","type": "uint256"}],"name": "getIDByIndex","outputs": [{"internalType": "address","name": "","type": "address"}],"stateMutability": "view","type": "function"},{"inputs": [{"internalType": "address","name": "deviceID","type": "address"}],"name": "getLatestCID","outputs": [{"internalType": "string","name": "latestCID","type": "string"}],"stateMutability": "view","type": "function"}]"""

conn = http.client.HTTPSConnection("kfs2.moibit.io")
moibit_url = 'https://kfs2.moibit.io/moibit/v0/'
moibit_header_obj = {
    'api_key': os.environ['MOIBIT_API_KEY'],
    'api_secret': os.environ['MOIBIT_API_SECRET'],
    'content-type': "application/json"
}


@app.route('/', methods=['GET'])
def home():
    return "<h1>DICTAO - Decentralized Intelligent Contact Tracing of Animals and Objects</h1><p>This is a simple demonstration of applying blockchain, decentralized storage and AI to solve the COVID-19 crisis.</p>"


@app.errorhandler(404)
def page_not_found(e):
    return "The given ID could not be found", 404


@app.errorhandler(500)
def internal_server_error(e):
    return e, 500


@app.route('/api/v0/get_infections', methods=['GET'])
def get_infections():
    masterDataSet = []
    query_parameters = request.args
    if 'id' in query_parameters:
        id = query_parameters.get('id')
        print("Received ID from the user: "+id)
        if getLatestCID(id) == "":
            return page_not_found(404)
        else:
            w3 = Web3(Web3.HTTPProvider(blockchain_url))
            contract = w3.eth.contract(
                os.environ['PROOF_SMART_CONTRACT_ADDRESS'], abi=abi)
            length = contract.functions.getDeviceIDsLength().call()
            print("Length of the deviceIDs: "+str(length))
            for i in range(length):
                tempId = contract.functions.getIDByIndex(i).call()
                # print(tempId)
                tempHash = contract.functions.getLatestCID(tempId).call()
                # print(tempHash)
                jsonData = getJsonDataFromMoiBit(tempHash)
                # print(jsonData)
                for location in jsonData:
                    masterDataSet.append(location)
            print("Generated live dataset of length: %d" % len(masterDataSet))
            try:
                with open('live_dataset.json', 'x') as outfile:
                    json.dump(masterDataSet, outfile, indent=2)
            except FileExistsError:
                os.remove('live_dataset.json')
                print("File Removed!")
                with open('live_dataset.json', 'x') as outfile:
                    json.dump(masterDataSet, outfile, indent=2)
            results = get_infected_ids(id)
            os.remove("live_dataset.json")
            response = {
                "id": id,
                "potential_infected_ids": results
            }
            return (jsonify(response))
    else:
        return "Error: Please specify an ID to identify potential infections."


def get_infected_ids(input_id):
    basePath = os.path.dirname(os.path.abspath('live_dataset.json'))
    dflive = pd.read_json(basePath + '/' + 'live_dataset.json')

    epsilon = 0.0018288 # a radial distance of 6 feet, which is medically presribed
    min_sample = 2
    model = DBSCAN(eps=epsilon, min_samples=min_sample, metric='haversine').fit(dflive[['latitude', 'longitude']])
    dflive['cluster'] = model.labels_.tolist()

    input_id_clusters = []
    for i in range(len(dflive)):
        if dflive['id'][i] == input_id:
            if dflive['cluster'][i] in input_id_clusters:
                pass
            else:
                input_id_clusters.append(dflive['cluster'][i])
#     print(input_id_clusters)
    
    infected_ids = []
    for cluster in input_id_clusters:
        if cluster != -1:
            ids_in_cluster = dflive.loc[dflive['cluster'] == cluster, 'id']
            for i in range(len(ids_in_cluster)):
                member_id = ids_in_cluster.iloc[i]
                if (member_id not in infected_ids) and (member_id != input_id):
                    infected_ids.append(member_id)
                else:
                    pass
    return infected_ids


def getJsonDataFromMoiBit(cid):
    pre_payload = {"hash": cid}
    payload = json.dumps(pre_payload)
    conn.request("POST", moibit_url+"readfilebyhash",
                 payload, moibit_header_obj)
    res = conn.getresponse()
    if res.status == 200:
        responseObject = json.loads(res.read())
        print(
            "updateLocationHistory(): Appending the captured data to historic data.")
        return responseObject


def getLatestCID(id):
    w3 = Web3(Web3.HTTPProvider(blockchain_url))
    contract = w3.eth.contract(
        os.environ['PROOF_SMART_CONTRACT_ADDRESS'], abi=abi)
    cid = ""
    try:
        cid = contract.functions.getLatestCID(id).call()
    except web3.exceptions.ValidationError:
        print("ID does not exist!")
        return ""
    except:
        print("Some other error occured!")
        return ""
    else:
        print(cid)
        return cid


app.run()

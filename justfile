export GMAPS_API_KEY := "<REPLACE_WITH_YOUR_KEY>"
export MOIBIT_API_KEY := "<REPLACE_WITH_YOUR_KEY>"
export MOIBIT_API_SECRET := "<REPLACE_WITH_YOUR_SECRET>"

run-python:
    python edge-functions/python/main.py

run-node:
    node edge-functions/nodejs/index.js

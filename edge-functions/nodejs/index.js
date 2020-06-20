const request = require('request');

const options = {
    url: 'https://www.googleapis.com/geolocation/v1/geolocate',
    form: {
        key: process.env.GMAPS_API_KEY
    }
};

function getGeoCordinates() {

    request.post(options, (err, res, body) => {
        if (err) {
            return console.log(err);
        }
        res = JSON.parse(body)

        console.log("Latitude: ", res.location.lat);
        console.log("Longitude:", res.location.lng);
        console.log("Accuracy of the location:", res.accuracy, "(in metres)");
    });

}

getGeoCordinates()

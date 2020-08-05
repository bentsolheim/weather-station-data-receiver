# weather-station-data-receiver
Simple endpoint for receiving data from the weather-station-mkrgsm1400

## Posting data to endpoint

Service will start on localhost:8081 by default.

In fish shell create a tmp function to allow piping long json without issues

    function tmp; set f (mktemp); echo $argv > "$f"; echo $f; end  
    
Readings-endpoint:
    
    http -v http://localhost:8081/api/v1/logger/bua/readings < (tmp '{ "readings": [ { "sensorName": "inne-temp", "value": 25.700001, "localTime": 1596663858, "unixTime": 1596656658 }, { "sensorName": "inne-humidity", "value": 68.000000, "localTime": 1596663858, "unixTime": 1596656658 } ] }')
    
Debug-endpoint:

    http -v http://localhost:8081/api/v1/logger/bua/debug < (tmp '{ "signalStrength": "6", "timeSpent": 17658, "iteration": 26, "connectionErrors": 1, "sensorErrors": 0, "millisSinceStart": 7891839, "battery": { "analogReading": 943, "voltage": 3.925273, "level": 67 } }')


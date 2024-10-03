#!/bin/bash
PORT=8080
HOST=localhost
URL=http://${HOST}:${PORT}/pack
DATA="'{\"Value\":123,\"Buckets\":[100,20,3]}'"
HEADER="\"Content-Type: application/json\""
CMD="curl -H $HEADER -d $DATA $URL"
echo -n Running command: $CMD " - result: "
eval $CMD | jq

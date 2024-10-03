# Paker API & frontend

## Description

The project shows how Fyne frontend may be used to trigger /pack endpoint of an REST HTTP server.
The packing algorhitm is in fact greedy search with merge phase to optimize bucket usage.

## Usage

### API

API and frontend are built as Golang executables to be started locally.
The make target build created them in ./bin path. The API can started with the target run.
The config for the HTTP REST server is stored in ./config/config.yaml file.

### Frontend

The hardcoded target values are localhost and port 8080. The exec can be started with command
./bin/packer-frontend. The packeages used un the run are loaded rom env variable BUCKETS.
The format of the value is csv without separators.

### Testing

The unit tests are to be started with make target test.

The run test uses curl and jq so that those apps shall be installed locally. The run test
make target is runtest.

### Building

The build of the executables is started wuth make target build.

TARGETS = packer-api packer-frontend

test:
	go test -v ./...

packer-api:
	go build -o bin/packer-api cmd/packer-api/main.go

packer-frontend:
	go build -o bin/packer-frontend cmd/packer-frontend/main.go

build: $(TARGETS)

clean:
	rm -f ./bin/*
	find . -name *~ -exec rm {} \;

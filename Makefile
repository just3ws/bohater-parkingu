GO=go
BUILD=$(GO) build
CLEAN=$(GO) clean
BINARY_NAME=bohater-parkingu

build:
	$(BUILD) -v

clean:
	$(CLEAN)
	rm -f $(BINARY_NAME)

start:
	docker-compose up --remove-orphans --force-recreate --build

stop:
	docker-compose down --remove-orphans --rmi all

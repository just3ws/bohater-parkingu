GO=go
BUILD=$(GO) build
CLEAN=$(GO) clean
BINARY_NAME=bohater-parkingu

build:
	$(BUILD) -v

clean:
	$(CLEAN)
	rm -f $(BINARY_NAME)

run:
	docker-compose up --remove-orphans --force-recreate --build --detach

stop:
	docker-compose down --remove-orphans --rmi all

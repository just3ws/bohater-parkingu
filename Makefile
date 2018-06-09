GO=go
BUILD=$(GO) build
CLEAN=$(GO) clean
BINARY_NAME=bohater-parkingu

build:
	$(BUILD) -v
clean:
	$(CLEAN)
	rm -f $(BINARY_NAME)

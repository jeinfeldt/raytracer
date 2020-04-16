# Makefile

# Constants
WIDTH=400
HEIGHT=200

# Targets
run:
	go run main.go

image:
	go run main.go image $(WIDTH) $(HEIGHT)

test:
	go test --count=1 --failfast -v ./...

clean:
	rm -rf *.png *.ppm

.PHONY: clean
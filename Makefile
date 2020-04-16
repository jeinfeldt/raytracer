# Makefile

# Constants
WIDTH=100
HEIGHT=50

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
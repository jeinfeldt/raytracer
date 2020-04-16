# Makefile

# Constants
WIDTH=600
HEIGHT=300

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
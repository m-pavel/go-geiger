all: mqtt

deps:
	go get -v -d ./...

mqtt: deps
	go build -o geiger-mqtt ./mqtt

clean:
	rm -f geiger-mqtt

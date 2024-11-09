build:
	cd chat && go build -o chat

run:
	cd chat && ./chat

.PHONY: build run

start:
	@go build \
	&& ./msgr-bot &

refresh:
	@go build \
	&& killall -1 msgr-bot

stop:
	@killall msgr-bot

build:
	@go build
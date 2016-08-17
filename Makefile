
all:
	go build 

test_parse:
	go test context.go rando.go parse.go parse_test.go -v

test_scanner:
	go test scanner_test.go -v

.PHONY: test_parse test_scanner

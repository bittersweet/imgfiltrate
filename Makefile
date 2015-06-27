all:
	@@go build imgfiltrate.go
	@@mv imgfiltrate /usr/local/bin/
	@@echo "Built and moved"

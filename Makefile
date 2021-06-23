build: tools
	go build .

tools:
	@# Install all tools in tools.go into ./bin, don't forget to add ./bin to your PATH!
	./helper.sh

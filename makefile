build:
	go build .
	chmod +x wprecon

install:
	sudo mv wprecon /usr/local/bin
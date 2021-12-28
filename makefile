build:
	go build cli/main.go
	chmod +x wprecon

install:
	sudo mv wprecon /usr/local/bin
	mkdir $$HOME/.wprecon
	mv internal/config/config.yaml $$HOME/.wprecon
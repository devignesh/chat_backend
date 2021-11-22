all: build run

build:
	@echo "buid Go"
	go build -o chat_backend

run:
	@echo "started server"
	./chat_backend

clean:
	@echo "remove build"
	rm -rf chat_backend

build:
	GOOS=windows GOARCH=amd64 go build -o bin/client.exe client/main.go

clean:
	rm -rf bin/client.exe

run: build
	./bin/client.exe

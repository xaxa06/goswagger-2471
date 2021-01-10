build: clean
	swagger generate client --target pkg/example --name example --spec swagger.yaml
	mkdir -p bin
	go build -o bin/bug
	
clean:
	rm -rf pkg/example/*


all: clean test.wasm run

test.wasm:
	GOOS=js GOARCH=wasm go build -o test.wasm main.go

run:
	node wasm_exec.js test.wasm

clean:
	rm test.wasm

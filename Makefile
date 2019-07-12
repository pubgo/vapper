
wasm:
	GOARCH=wasm GOOS=js go build -o example.wasm ./client
	mv example.wasm ./app/

run:  wasm
	cd server && go build && cd .. && ./server/server

#https://compile.jsgo.io/<path>
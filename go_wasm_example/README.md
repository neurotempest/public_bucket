Following tutorial from https://golangbot.com/webassembly-using-go/




### Build steps

0. Copy wasm js to assets:
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets
```


1. compile
```
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm
```

2. there is a fileserver in cmd/server which will serve the webpage

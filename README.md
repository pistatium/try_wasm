# try go wasm

## Environment
`go version go1.13.8 linux/amd64`

## Build wasm

```
GOOS=js GOARCH=wasm go build -o test.wasm
```

## Start nginx
```
 docker run -it -v $(pwd):/usr/share/nginx/html -v $(pwd)/mime.types:/etc/nginx/mime.types -p 8888:80 nginx:alpine
```

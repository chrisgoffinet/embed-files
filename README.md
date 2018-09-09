# embed-files
Simple golang http server that serves static files using fileb0x

## Requirements
You must have [fileb0x](https://github.com/UnnoTed/fileb0x) installed to bundle your static assets.

```
go get -u github.com/UnnoTed/fileb0x
```

## Build
```
$ make assets
```

## Run
```
$ go run main.go
2018/09/08 21:07:36 listening on :8080
```

Point your browser to http://localhost:8080/
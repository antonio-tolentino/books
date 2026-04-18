# Books API


## Gerarate Code coverage and html report
```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```


## Build image locally
```bash
## build image
docker build --rm -t localhost/books:v0.0.1 .
```
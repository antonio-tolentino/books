## Code static tests
go test . -coverprofile=coverage.out

## build image
docker build --rm -t tolentino/books:v1.0.0 .

## push image
docker push tolentino/books:v1.0.0 

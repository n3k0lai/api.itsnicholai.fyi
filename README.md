# My API

install:
```sh
# First initialize a Go module if you haven't already
go mod init api

# Then add dependencies
go get github.com/gorilla/mux
go get github.com/rs/cors
```

run:
```sh
go run main.go
```

Get Data:
```
curl -X GET http://localhost:6000/about | bat -l json
```

Build the Docker image:
```
docker build -t nicholai-api .
```

Run the container:
```
docker run -p 6000:6000 nicholai-api
```

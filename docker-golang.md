# docker golang

## pull images

```
docker pull golang:1.20-bullseye
```

## test images

```
docker run -it --rm --name glapp golang:1.20-bullseye
```

## Dockerfile

```
FROM golang:1.20
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./...
CMD ["app"]
```

## build images

```
docker build -t golang:mybuild .
docker run -it --rm --name my-running-app golang:mybuild
```

MORE INFO : https://hub.docker.com/_/golang

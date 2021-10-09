FROM golang:alpine AS build-stage

WORKDIR /usr/local/go/src/build

COPY src/* go.mod go.sum ./
RUN go build -o /usr/local/bin/myip-server .

FROM alpine

WORKDIR /usr/local/bin

COPY --from=build-stage /usr/local/bin/myip-server ./

ENTRYPOINT ["/usr/local/bin/myip-server"]

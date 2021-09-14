FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY pkg ./pkg
COPY cmd ./cmd

WORKDIR ./pkg/apidockertest
RUN go mod download
WORKDIR ../../cmd
RUN go build
CMD ["./cmd"]
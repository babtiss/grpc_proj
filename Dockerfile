FROM golang:1.17-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o grpc_proj ./cmd/server/main.go

CMD ["./grpc_proj"]
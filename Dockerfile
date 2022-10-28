FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avito-test-backend ./cmd/main.go
CMD ["./avito-test-backend"]
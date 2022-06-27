FROM golang:latest

ENV GOPATH=

WORKDIR /app
COPY . /app

RUN go mod init main \
RUN go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8080


CMD ["go", "run", "server.go"]

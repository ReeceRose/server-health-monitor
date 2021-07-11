FROM golang:1.16.5-alpine3.14

WORKDIR /src

ADD go.mod go.sum ./
COPY . .
RUN go mod download

RUN go build -o ./out/api cmd/api/main.go

ENTRYPOINT [ "./out/api" ]
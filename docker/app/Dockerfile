FROM golang:1.21

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["air"]
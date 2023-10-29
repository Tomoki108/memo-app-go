FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/onsi/ginkgo/v2/ginkgo

CMD ["air"]
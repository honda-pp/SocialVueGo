FROM golang:1.20.3

WORKDIR /app

COPY . .

RUN go install github.com/cosmtrek/air@v1.29.0

RUN go mod download

CMD ["go", "run", "main.go"]
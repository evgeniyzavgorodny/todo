FROM golang:1.23-alpine3.21

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .
ENV DATABASE_URL="postgresql://postgres:123456@host.docker.internal:5432/postgres"
ENV HTTP_ADDRESS="0.0.0.0:3000"

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]

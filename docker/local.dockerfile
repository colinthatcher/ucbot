# Start by building the application.
FROM golang:1.25 as builder

WORKDIR /app/
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@v1.64.5

EXPOSE 8080

CMD ["air"]

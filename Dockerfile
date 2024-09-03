
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /site-metadata-api-golang

EXPOSE 8080

CMD ["/site-metadata-api-golang"]
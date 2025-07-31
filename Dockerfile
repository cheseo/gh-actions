FROM golang:1.24.5

WORKDIR /app

COPY go.mod ./
RUN go mod download
EXPOSE 8000

COPY *.go ./

RUN go build

CMD ["./gh-actions"]


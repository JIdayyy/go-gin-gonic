FROM golang:1.21

WORKDIR /app

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

COPY . ./
COPY . ./
RUN go mod download

RUN  go build -o /go-http

EXPOSE 4010

# Run
CMD ["/go-http"]

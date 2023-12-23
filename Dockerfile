FROM golang:1.21

WORKDIR /app

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

COPY . .

RUN go mod download

RUN  go build -o /go-http

EXPOSE 4010

# Run
# CMD ["/go-http"]
CMD ["air"]

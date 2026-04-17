FROM golang:1.23-bullseye as builder

ENV GO111MODULE=on

WORKDIR /app

# Copy app and run go mod.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app .

# the binary need compilation tools for library kafka used
FROM golang:1.23-bullseye

WORKDIR /root/

COPY --from=0 /app .

ENTRYPOINT ["/root/app"]

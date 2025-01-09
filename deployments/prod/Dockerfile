# Step 1: Modules caching
FROM golang:alpine AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:alpine AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux  \
    go build  -o /bin/app ./cmd/main.go

# Step 3: Final
FROM alpine

COPY --from=builder bin/app /
COPY --from=builder app/api/bundle.yml /api/

RUN apk --no-cache add ca-certificates

CMD ["/app"]
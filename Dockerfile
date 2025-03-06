FROM golang:1.20-alpine AS builder
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /display
EXPOSE 8080
CMD ["/display"]
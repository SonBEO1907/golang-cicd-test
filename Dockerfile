FROM golang:1.20-alpine
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-world
EXPOSE 8080
CMD ["/hello-world"]
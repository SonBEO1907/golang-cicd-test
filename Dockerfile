FROM golang
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org
RUN go mod tidy && go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-world
EXPOSE 8080
CMD ["/hello-world"]
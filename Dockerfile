FROM golang:1.22-alpine as build
WORKDIR /app
COPY . .
RUN go mod init example.com/demo && go mod tidy
RUN go build -o demo

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/demo /app/demo
EXPOSE 8080
CMD ["/app/demo"]


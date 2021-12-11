# FROM golang:1.16-alpine as build
FROM golang:1.13.8-alpine3.11 as build-env

WORKDIR /app
# COPY go.mod /app
# COPY go.sum /app
# RUN go mod download
COPY *.go /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o docker-simple-app

FROM alpine:3.11.3 as final
# COPY —-from=dev /docker-simple-app /
# WORKDIR /app
COPY --from=build-env /app/docker-simple-app .
EXPOSE 8888

CMD ["./docker-simple-app"]


# RUN CGO_ENABLED=0 go build -ldflags=”-s -w” -o /my-awesome-app

# FROM scratch

# COPY — from=build /my-awesome-app /my-awesome-app/
# ENTRYPOINT [“/my-awesome-app”]
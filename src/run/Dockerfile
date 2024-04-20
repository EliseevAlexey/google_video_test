FROM golang:1.22 AS build
COPY go.mod *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /run

FROM alpine:3.19.1
COPY --from=build /run .
ENTRYPOINT ["./run"]

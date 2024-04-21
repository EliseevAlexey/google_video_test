FROM golang:1.22.2 AS build
COPY src/go.mod src/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /api

FROM alpine:3.19.1
COPY --from=build /api .
ENTRYPOINT ["./api"]

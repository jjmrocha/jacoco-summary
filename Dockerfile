# Build image
FROM golang:1.24.3 AS build

ENV BUILD_HOME "/build"
RUN mkdir $BUILD_HOME
WORKDIR $BUILD_HOME

COPY go.mod ./
COPY cmd cmd
COPY jacoco jacoco
COPY action action

RUN go build -o jsreport cmd/main.go

# Final image
FROM debian:12.10-slim

COPY --from=build /build/jsreport /

ENTRYPOINT ["/jsreport"]

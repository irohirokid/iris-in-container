FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /iris-example

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /iris-example /iris-example

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/iris-example"]

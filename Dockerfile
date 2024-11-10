FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY internal internal
COPY main.go .

RUN go build

FROM alpine 

RUN apk add gcompat

COPY --from=build /app/gone-broadcast /gone-broadcast

ENTRYPOINT ["/gone-broadcast"]





FROM golang:1.22.1-alpine as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /src/backend

FROM alpine:3.19

COPY --from=builder /src/backend /backend

EXPOSE 8090

ENTRYPOINT ["/backend"]
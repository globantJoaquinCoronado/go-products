FROM golang:1.18 AS GO_BUILD
ENV CGO_ENABLED 0
COPY . /products
WORKDIR /products
RUN go build -o server

FROM alpine:3.15
WORKDIR /products
COPY --from=GO_BUILD /products/server /products/server
EXPOSE 8000
CMD ["./server"]
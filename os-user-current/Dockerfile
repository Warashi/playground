FROM golang:1.19.3 as builder
WORKDIR /app
COPY main.go ./main.go
RUN go build -o main ./main.go

FROM debian
USER 1000
COPY --from=builder /app/main ./
CMD ["./main"]

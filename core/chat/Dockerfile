FROM golang:alpine3.17 as builder

WORKDIR /chat
COPY . .

WORKDIR /chat/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../app

FROM scratch
WORKDIR /chat

ENV PORT="8040"
EXPOSE 8040
EXPOSE 3011

COPY --from=builder /chat/app .
CMD ["./app"]

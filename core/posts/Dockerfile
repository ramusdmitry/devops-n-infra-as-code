FROM golang:1.19-buster as builder

WORKDIR /post-app
COPY . .

WORKDIR /post-app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../app

FROM alpine:3.15.4
WORKDIR /post-app

ENV PORT="8036"
ENV DBNAME="blog-db"
ENV DBUSERNAME="admin"
ENV DBPASSWORD="admin"
ENV DBHOST="localhost"
ENV DBPORT="5432"
ENV SSLMODE="disable"

EXPOSE 8036

COPY --from=builder /post-app/app .
CMD ["./app"]

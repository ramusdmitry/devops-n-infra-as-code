FROM golang:1.19-buster as builder

WORKDIR /group-app
COPY . .

WORKDIR /group-app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../app

FROM alpine:3.15.4
WORKDIR /group-app

ENV PORT="8050"
ENV DBNAME="blog-db"
ENV DBUSERNAME="admin"
ENV DBPASSWORD="admin"
ENV DBHOST="localhost"
ENV DBPORT="5432"
ENV SSLMODE="disabled"
EXPOSE 8050

COPY --from=builder /group-app/app .
CMD ["./app"]

FROM golang:1.19-buster as builder

WORKDIR /comms
COPY . .

WORKDIR /comms/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../app

FROM scratch
WORKDIR /comms

ENV PORT="8032"
ENV DBNAME="blog-db"
ENV DBUSERNAME="admin"
ENV DBPASSWORD="admin"
ENV DBHOST="localhost"
ENV DBPORT="5432"
ENV SSLMODE="disable"

EXPOSE 8032

COPY --from=builder /comms/app .
CMD ["./app"]
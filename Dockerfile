# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /usr/src/app
COPY . .
WORKDIR ./server
RUN go build -o /server
EXPOSE 8080
CMD ["/server"]
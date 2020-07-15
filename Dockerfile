FROM alpine:latest

WORKDIR /app

COPY ./dist/loment /app

EXPOSE 80

ENTRYPOINT ["./loment"]
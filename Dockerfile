FROM alpine:latest

WORKDIR /app

COPY ./dist/loment.exe /app

EXPOSE 80

ENTRYPOINT [ "./loment.exe" ]
#builder
FROM golang:alpine as builder
WORKDIR /home
COPY . .
RUN go build -o app

#final image
FROM alpine
RUN apk update && \
    apk add --no-cache tzdata && \
    apk add --no-cache curl
RUN rm -rf /var/cache/apk/* && date
COPY --from=builder /home/app .
EXPOSE 6060
CMD ./app
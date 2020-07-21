# Start from golang:1.12-alpine base image
# builder
FROM golang:1.14-alpine

RUN apk update && apk upgrade && \
    apk --update --no-cache add bash git make

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

RUN make engine

FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 9090

COPY --from=builder /app/engine /app

CMD /app/engine
# Start from golang:1.12-alpine base image
# builder
FROM golang:1.14-alpine as builder

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk --update --no-cache add bash git make

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependancies to vendor folder.
RUN go mod vendor

# Build the Go app
RUN go build -o example_service .

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8000

COPY --from=builder /app/example_service /app
COPY --from=builder /app/.env /app

# Run the executable
CMD /app/example_service
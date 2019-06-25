# build stage
FROM golang:1.12 AS builder
# working directory
WORKDIR /app/user-auth
COPY . .
# rebuilt built in libraries and disabled cgo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-auth .
# final stage
FROM alpine:latest
# working directory
WORKDIR /app/user-auth
# copy the binary file into working directory
COPY --from=builder /app/user-auth/user-auth .
# Run the docker_imgs command when the container starts.
CMD ["./user-auth"]
# http server listens on port 8080
EXPOSE 8000
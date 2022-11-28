# Start from golang base image
FROM golang:alpine as builder

WORKDIR /.
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /bin/numconverter

FROM gcr.io/distroless/static-debian11
COPY --from=builder /bin/numconverter /

# Run executable
CMD ["/numconverter"]

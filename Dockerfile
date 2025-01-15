FROM golang:1.23.4-alpine
COPY binary /app/binary
ENTRYPOINT ["/app/binary"]
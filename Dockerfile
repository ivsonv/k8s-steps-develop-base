FROM golang:1.22.1
COPY . .
RUN go env -w GO111MODULE=off
RUN go build -o server .
CMD ["./server"]
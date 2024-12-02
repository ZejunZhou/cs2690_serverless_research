
FROM golang:1.23

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH


WORKDIR /app


COPY . .

COPY grpc_shared_mem /app/grpc_shared_mem

RUN go build -o main .

ENTRYPOINT ["./main"]

FROM golang:alpine as builder
WORKDIR /app
RUN chmod 700 /app
COPY . .
ENV GO111MODULE=on CGO_ENABLED=0
#RUN go mod download
#RUN protoc --go_out=plugins=grpc:. api/proto/v1/proto-service.proto
#CMD [ "protoc" , "--go_out=plugins=grpc:." ,"api/proto/v1/proto-service.proto"]
EXPOSE 8080
RUN go build pkg/cmd/main.go

#final stage
FROM debian 
WORKDIR /app
COPY --from=builder /app/main /app/
RUN chmod 700 /app
ENTRYPOINT ["./main"]
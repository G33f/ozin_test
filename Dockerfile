FROM golang:latest AS builder
WORKDIR /build
COPY . ./
CMD ["make proto-gen"]
RUN CGO_ENABLED=0 go build -o short-url cmd/api/main.go

FROM ubuntu:latest
WORKDIR /root
COPY --from=builder build/short-url ./
COPY --from=builder build/config/config.yml ./config/
#CMD["./short-url"]
ARG STORAGE_TYPE
CMD ./short-url -storage_type=$STORAGE_TYPE
EXPOSE 8080
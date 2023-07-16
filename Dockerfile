FROM golang:1.20 as builder
ENV GOOS=linux \
    CGO_ENABLED=0 

WORKDIR /go/src/todolist
COPY go.mod go.sum ./
RUN go mod download -x

ADD internal internal
ADD pkg pkg
ADD cmd ./

RUN go build -o /app


FROM alpine:3.16 as production

RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 8000

ENTRYPOINT [ "./app" ]
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build


COPY . .

ADD go.mod .
ADD go.sum .
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/main main.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
ENV DOCKER_ENV=true

WORKDIR /app
COPY . .
COPY --from=builder /app/main /app/main

CMD ["./main"]
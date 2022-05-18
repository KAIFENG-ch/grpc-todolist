FROM golang:1.17 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN  CGO_ENABLE=0 GOOS=Linux GOARCH=amd64 go build -o main
RUN makedir /publish && cp main /publish && cp -r conf /publish


FROM busybox:1.28.4

WORKDIR /app

COPY --from=builder /app/publish .

EXPOSE 8000

ENTRYPOINT["./main"]
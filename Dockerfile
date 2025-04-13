FROM golang:1.21-alpine AS builder

WORKDIR /app

# 複製go.mod
COPY go.mod ./

# 複製源碼
COPY *.go ./

# 編譯應用程式
RUN go build -o /bin/ptt-login

FROM alpine:3.18

# 安裝運行時所需套件
RUN apk add --no-cache openssh-client expect ca-certificates

# 從builder複製編譯好的應用
COPY --from=builder /bin/ptt-login /bin/ptt-login

# 創建日誌目錄
RUN mkdir -p /var/log

# 設置入口點
ENTRYPOINT ["/bin/ptt-login"]
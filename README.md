# PTT 自動登入工具

這是一個使用 Golang 撰寫的 PTT 自動登入工具，並提供 Docker 封裝。透過此工具，您可以輕鬆實現對 PTT (批踢踢實業坊) 的自動化登入與登出。

## 功能特點

- ✅ 全自動登入 PTT
- ✅ 處理重複登入提示 (自動回答 "n")
- ✅ 登入後自動保持 10 秒
- ✅ 自動登出
- ✅ Docker 容器化部署
- ✅ 完整日誌記錄

## 系統需求

- Docker
- 有效的 PTT 帳號

## 快速開始

### 從源碼構建

1. 克隆此儲存庫
   ```bash
   git clone https://github.com/yourusername/ptt-login.git
   cd ptt-login
   ```

2. 構建 Docker 映像
   ```bash
   docker build -t ptt-login .
   ```

3. 運行容器
   ```bash
   docker run --rm -e PTT_ID=你的PTT帳號 -e PTT_PWD=你的PTT密碼 ptt-login
   ```

### 使用 Docker Compose

1. 創建 `.env` 檔案並設定您的 PTT 帳號密碼：
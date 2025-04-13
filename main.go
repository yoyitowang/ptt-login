package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// 檢查必要的環境變數是否設定
	pttID := os.Getenv("PTT_ID")
	pttPWD := os.Getenv("PTT_PWD")
	
	if pttID == "" || pttPWD == "" {
		log.Fatal("環境變數 PTT_ID 和 PTT_PWD 必須設定")
	}

	// 設置日誌
	logDir := "/var/log"
	os.MkdirAll(logDir, 0755)
	
	logFile, err := os.OpenFile(logDir+"/ptt.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("警告: 無法打開日誌文件: %v", err)
	} else {
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	// 創建expect腳本文件 - 修正語法錯誤
	scriptContent := `#!/usr/bin/expect -f
set timeout 60
log_file -noappend /var/log/ptt.log
spawn ssh -o "StrictHostKeyChecking no" bbsu@ptt.cc
expect "new"
send "$env(PTT_ID)\r"
expect ":"
send "$env(PTT_PWD)\r"
expect {
    "注意: 您有其它連線已登入此帳號" {
        # 使用部分字符串匹配，避免方括號問題
        expect "您想刪除其他重複登入的連線嗎"
        send "n\r"
        exp_continue
    }
    "您想刪除其他重複登入的連線嗎" {
        send "n\r"
        exp_continue
    }
    "請按任意鍵繼續" {
        send "\r"
        exp_continue
    }
    "批踢踢實業坊" {
        puts "成功登入PTT，等待10秒後自動登出..."
        # 登入成功，延遲10秒後登出
        sleep 10
        # 嘗試使用Ctrl+C來觸發登出流程
        send "\003"
        expect {
            "您確定要離開" {
        send "y\r"
}
            timeout {
                # 如果Ctrl+C沒反應，嘗試輸入q
                send "q\r"
                expect "您確定要離開"
                send "y\r"
}
        }
        expect eof
    }
}
`

	scriptPath := "/tmp/ptt_login.exp"
	err = os.WriteFile(scriptPath, []byte(scriptContent), 0755)
	if err != nil {
		log.Fatalf("無法寫入expect腳本: %v", err)
	}

	// 執行expect腳本
	cmd := exec.Command("expect", scriptPath)
	
	// 傳遞環境變數
	cmd.Env = os.Environ()
	
	// 連接命令輸出到日誌
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	log.Println("開始PTT登入")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("執行登入腳本失敗: %v", err)
	}
	
	log.Println("PTT登入/登出流程成功完成")
}
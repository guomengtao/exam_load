package main

import (
	"fmt"
	"os/exec"
	"time"
)

// macOS 通知
func notify(msg string) {
	exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "专注提醒"` , msg)).Run()
}

// macOS 锁屏
func lockScreen() {
	// 这条命令在 Mac 上锁屏
	exec.Command("osascript", "-e", `tell application "System Events" to keystroke "q" using {control down, command down}`).Run()
	// 或者使用以下命令（macOS Ventura 及以上支持）
	// exec.Command("pmset", "displaysleepnow").Run()
}

func main() {
	fmt.Println("⏳ 开始专注 60 分钟...")

	startTime := time.Now()

	// 提前 5 分钟提醒
	go func() {
		time.Sleep(1 * time.Minute)
		notify("⚠️ 距离锁屏还有 5 分钟，请收尾！")
		fmt.Println("⏰ 5 分钟后将锁屏")
	}()

	// 主倒计时
	time.Sleep(2 * time.Minute)
	notify("🔒 时间到，已锁屏")
	lockScreen()

	endTime := time.Now()
	fmt.Printf("✅ 开始于: %s\n", startTime.Format("15:04:05"))
	fmt.Printf("🔒 锁屏于: %s\n", endTime.Format("15:04:05"))
}
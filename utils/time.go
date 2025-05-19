package utils

import "time"

// NowTimeString 返回当前时间格式 MMddHHmmss 字符串，比如 0519101530 表示5月19日10:15:30
func NowTimeString() string {
    return time.Now().Format("0102150405") // 01月02日15时04分05秒，截取年月日时分秒
}
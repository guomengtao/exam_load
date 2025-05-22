package genlib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// writeFileWithPrompt 写文件，带交互覆盖提示。force为true时强制覆盖。
func writeFileWithPrompt(path string, content []byte, force bool) error {
	if _, err := os.Stat(path); err == nil && !force {
		// 文件存在，提示是否覆盖
		fmt.Printf("文件 %s 已存在，是否覆盖？(y/N): ", path)
		reader := bufio.NewReader(os.Stdin)
		resp, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("读取输入失败，默认不覆盖: %v", err)
			return nil
		}
		resp = strings.TrimSpace(resp)
		if !strings.EqualFold(resp, "y") {
			log.Printf("跳过文件生成: %s\n", path)
			return nil
		}
	}
	return os.WriteFile(path, content, 0644)
}
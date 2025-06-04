package utils

import (
	"os"
	"path/filepath"
)

// ProjectRoot 返回当前项目的根目录路径
func ProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			// 到达根目录还没找到
			break
		}
		dir = parent
	}

	return "." // fallback
}

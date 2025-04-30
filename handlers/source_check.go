package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Source struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Type     string `json:"type"`     // 前端 / 后端 / 通用
	Location string `json:"location"` // 国内 / 国外
	Status   string `json:"status,omitempty"`
	Latency  int64  `json:"latency_ms,omitempty"`
}

func loadSources(path string) ([]Source, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var sources []Source
	err = json.Unmarshal(data, &sources)
	return sources, err
}

func checkSourceAvailability(src *Source) {
	start := time.Now()
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(src.URL)
	latency := time.Since(start).Milliseconds()

	if err != nil || resp.StatusCode >= 400 {
		src.Status = "unreachable"
		src.Latency = latency
		return
	}
	src.Status = "ok"
	src.Latency = latency
}

func CheckAllSources(c *gin.Context) {
	sources, err := loadSources("static/sources.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加载源文件失败", "data": nil})
		return
	}
	for i := range sources {
		checkSourceAvailability(&sources[i])
	}

	// 可选：保存检查结果
	data, _ := json.MarshalIndent(sources, "", "  ")
	_ = ioutil.WriteFile("sources_checked.json", data, 0644)

	c.JSON(200, gin.H{"code": 200, "msg": "测试完成", "data": sources})
}
package myutils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os/exec"
	"strings"
)

type DataConsumed struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func GetStatistics(reset bool) (data []*DataConsumed, err error) {
	_v2ctl := "/usr/bin/v2ray/v2ctl"
	cmd := exec.Command(_v2ctl, "api", "--server", "127.0.0.1:8144", "StatsService.QueryStats", fmt.Sprintf("pattern: \"\" reset: %t", reset))
	logs.Debug(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Error(err, string(output))
		return
	}

	stats := strings.ReplaceAll(string(output), "\n>", "\n},")
	stats = strings.ReplaceAll(stats, "stat: <", "{")
	stats = strings.ReplaceAll(stats, "value", ",\"value\"")
	stats = strings.ReplaceAll(stats, "name", "\"name\"")
	stats = fmt.Sprint("[", stats[:len(stats)-3], "]")
	logs.Info(stats)

	err = json.Unmarshal([]byte(stats), &data)
	if err != nil {
		logs.Error(err)
	}
	return data, nil
}

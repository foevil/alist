package errs

import "fmt"

var (
	SearchNotAvailable  = fmt.Errorf("搜索未启用")
	BuildIndexIsRunning = fmt.Errorf("正在构建索引，请稍后重试")
)

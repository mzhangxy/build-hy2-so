package main

import "C"
import (
	"os"

	// 引入 Hysteria 2 官方的入口包
	"github.com/apernet/hysteria/app/cmd"
)

//export StartHysteria2
func StartHysteria2(configPath *C.char) C.int {
	// 将 Node.js 传过来的 C 字符串转换为 Go 字符串
	path := C.GoString(configPath)
	
	// 模拟命令行参数：相当于执行 `./hysteria server -c config.yaml`
	os.Args = []string{"hysteria", "server", "-c", path}
	
	// 使用 Goroutine 在后台异步启动 Hysteria，防止阻塞 Node.js 主线程
	go func() {
		cmd.Execute()
	}()
	
	return 0
}

// 编译为 c-shared 模式时，必须有一个空的 main 函数
func main() {}

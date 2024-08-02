package log

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func demo() {
	// 自定义配置，每个配置项的含义可以查看封装代码中的注释
	//c := NewConfig(
	//	SetBaseDirectoryName("alnklog"),
	//	SetInfoDirectoryName("alnkinfo"),
	//	SetWarnDirectoryName("alnkwarn"),
	//	SetErrorDirectoryName("alnkerr"),
	//	SetInfoFileName("alnk_info"),
	//	SetWarnFileName("alnk_warn"),
	//	SetErrorFileName("alnk_err"),
	//	SetLogFileMaxSize(1),
	//	SetLogFileMaxBackups(1),
	//	SetLogFileMaxAge(1),
	//	SetLogFileCompress(true),
	//	SetLogPrintTag(true),
	//)

	// 默认配置
	c := NewConfig()
	if err := Init(c); err != nil {
		fmt.Println(err)
	}

	testLogger()

}

func testLogger() {
	startTime := time.Now()

	for i := 0; i < 1000000; i++ {
		// 标准记录
		Logger.Info("Logger info")
		Logger.Warn("Logger warn")
		Logger.Error("Logger err")

		// 添加其他的key-value
		Logger.Info("Logger info", zap.Any("user", "alnk"))
		Logger.Warn("Logger Warn", zap.Any("user", "alnk"))
		Logger.Error("Logger Error", zap.Any("user", "alnk"))
	}

	fmt.Println("Logger执行时间: ", time.Since(startTime))
}

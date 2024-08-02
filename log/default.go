package log

import "fmt"

func InitDefaultLogger() {
	c := NewConfig()
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}

func InitConsoleDebugLogger(max, num, age int) {
	c := NewConfig(
		SetLogFileMaxSize(max),
		SetLogFileMaxBackups(num),
		SetLogFileMaxAge(age),
		SetLogFileCompress(true),
		SetLogPrintTag(true),
	)
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}

func InitFileDebugLogger(max, num, age int) {
	c := NewConfig(
		SetLogFileMaxSize(max),
		SetLogFileMaxBackups(num),
		SetLogFileMaxAge(age),
		SetLogFileCompress(true),
		SetLogDebugEnabled(true),
	)
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}

func InitAllDebugLogger(max, num, age int) {
	c := NewConfig(
		SetLogFileMaxSize(max),
		SetLogFileMaxBackups(num),
		SetLogFileMaxAge(age),
		SetLogFileCompress(true),
		SetLogPrintTag(true),
		SetLogDebugEnabled(true),
	)
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}

func InitLiteLogger(max, num, age int) {
	c := NewConfig(
		SetLogFileMaxSize(max),
		SetLogFileMaxBackups(num),
		SetLogFileMaxAge(age),
		SetLogFileCompress(true),
	)
	if err := Init(c); err != nil {
		fmt.Println(err)
	}
}

package trans

import (
	"trans/cfg"
	"trans/cmd"
	"trans/log"
	"trans/task"
)

func main() {
	cfg.InitConfig()
	log.InitWithViper()
	go cmd.Execute()
	task.RunAsyncQServer()
}

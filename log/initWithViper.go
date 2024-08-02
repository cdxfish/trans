package log

import "github.com/spf13/viper"

func InitWithViper() {
	logDebug := viper.GetBool("logger.debug")
	logConsole := viper.GetBool("logger.console")

	if logDebug && logConsole {
		InitAllDebugLogger(10, 10, 1)
	} else if logDebug && !logConsole {
		InitFileDebugLogger(10, 10, 1)
	} else if !logDebug && logConsole {
		InitConsoleDebugLogger(10, 10, 1)
	} else {
		InitLiteLogger(10, 10, 1)
	}
}

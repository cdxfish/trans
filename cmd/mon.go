/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"filemon/log"
	"filemon/task"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// monCmd represents the mon command
var monCmd = &cobra.Command{
	Use:   "mon",
	Short: "monitor file change and send event out",
	Long:  `monitor file change and send event out`,
	Run: func(cmd *cobra.Command, args []string) {
		monPath := viper.GetStringSlice("folder")
		log.Logger.Info("mon is running...", zap.Strings("path", monPath))

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			panic("failed to init file watcher.")
		}

		defer watcher.Close()

		done := make(chan bool)

		client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

		go func() {
			for {
				select {
				case evt, ok := <-watcher.Events:
					if !ok {
						return
					}
					if evt.Op&fsnotify.Create == fsnotify.Create {
						log.Logger.Debug("created file.", zap.String("filename", evt.Name))

						//task.AddTask(task.NewTask(map[string]interface{}{
						//	"paramA": "val",
						//}, []task.FacFunc{
						//	func(s string, m map[string]interface{}) (string, error) {
						//		sum := 1000
						//		for sum > 0 {
						//			time.Sleep(10)
						//			sum--
						//		}
						//		fmt.Println(s, evt.Name, windows.CurrentThread())
						//		return "ok", nil
						//	},
						//}, -1))

						task, err := task.NewFileChangedTask(evt.Name)
						if err != nil {
							log.Logger.Error(err.Error())
						}
						taskInfo, err := client.Enqueue(task)
						if err != nil {
							log.Logger.Error(err.Error())
						}
						log.Logger.Info("async task has been create.", zap.String("id", taskInfo.ID), zap.String("type", taskInfo.Type))

					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println(err)
					log.Logger.Error("failed to mon a file.", zap.String("msg", err.Error()))
				}
			}
		}()

		for _, val := range monPath {
			err = watcher.Add(val)
			if err != nil {
				fmt.Println("failed to add a file watcher.", val)
				log.Logger.Error("failed to add a file watcher.", zap.String("filename", val))
			}
		}
		<-done
	},
}

func init() {
	rootCmd.AddCommand(monCmd)

}

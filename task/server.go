package task

import "github.com/hibiken/asynq"

func RunAsyncQServer() {
	server := asynq.NewServer(asynq.RedisClientOpt{Addr: ":6379"}, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	mux := asynq.NewServeMux()
	mux.Handle(TypeFileChangeEvent, &FileChangeEvent{})

	if err := server.Run(mux); err != nil {
		panic("failed to start asyncQ server. please double check redis and filemon program.")
	}
}

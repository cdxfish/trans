package task

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"time"
	"trans/log"
)

// event type.
const (
	TypeFileChangeEvent = "file:changed"
)

// used by server
type FileChangeEvent struct {
	// used by server, the handler of the event
}

func (f *FileChangeEvent) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var p FileChangeEventPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		log.Logger.Error("Async Q failed on Unmarshal.", zap.String("error", err.Error()), zap.String("msg", asynq.SkipRetry.Error()))
		return err
	}
	log.Logger.Debug("Async Q is processing task.", zap.Any("payload", p))

	//handle logic

	return nil
}

// used by client
type FileChangeEventPayload struct {
	Name string
}

func NewFileChangedTask(filename string) (*asynq.Task, error) {
	payload, err := json.Marshal(FileChangeEventPayload{Name: filename})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeFileChangeEvent, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

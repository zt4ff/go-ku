package goku

import (
	"goku/config"
	"goku/job"
)

// multiple queues
// queues can have multiple job kinds
// append jobs

type Queue struct {
	config config.Config
}

func NewQueue() (*Queue, error) {
	// TODO set configs
	// - memory is set from the config

	return &Queue{}, nil
}

func (q *Queue) RegisterJob(name string, job job.Job) {
	// store job information into api
}

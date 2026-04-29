package job

import "context"

// return nil for job failure
type Job interface {
	Process(context.Context) error
}

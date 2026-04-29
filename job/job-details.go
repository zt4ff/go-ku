package job

type ID string

type Name string

type JobDetails struct {
	Job  Job
	ID   ID
	Name Name
}

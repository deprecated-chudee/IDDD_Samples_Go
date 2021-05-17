package task

type Status string

const (
	NotStarted Status = "NOT_STARTED"
	InProgress Status = "IN_PROGRESS"
	Impeded    Status = "IMPEDED"
	Done       Status = "DONE"
)

func (t Status) IsNotStarted() bool {
	return t == NotStarted
}

func (t Status) IsInProgress() bool {
	return t == InProgress
}

func (t Status) IsImpeded() bool {
	return t == Impeded
}

func (t Status) IsDone() bool {
	return t == Done
}
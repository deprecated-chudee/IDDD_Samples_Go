package backlogitem

type Status string

const (
	Planned   Status = "PLANNED"
	Scheduled Status = "SCHEDULED"
	Committed Status = "COMMITTED"
	Done      Status = "DONE"
	Removed   Status = "REMOVED"
)

func (s Status) IsPlanned() bool {
	return s == Planned
}

func (s Status) IsScheduled() bool {
	return s == Scheduled
}

func (s Status) IsCommitted() bool {
	return s == Committed
}

func (s Status) IsDone() bool {
	return s == Done
}

func (s Status) IsRemoved() bool {
	return s == Removed
}

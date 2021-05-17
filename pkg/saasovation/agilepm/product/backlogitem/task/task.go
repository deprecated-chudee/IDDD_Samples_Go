package task

import (
	//"iddd_samples_go/pkg/saasovation/agilepm/domain/product/backlogitem"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"
	backlogitem2 "iddd_samples_go/pkg/saasovation/agilepm/product/backlogitem"
)

type ID string

type Task struct {
	backlogItemID  backlogitem2.ID
	description    string
	estimationLog  []*backlogitem2.EstimationLogEntry
	hoursRemaining int
	name           string
	status         Status
	taskID         ID
	tenantID       tenant.ID
	volunteer      team.MemberID
}
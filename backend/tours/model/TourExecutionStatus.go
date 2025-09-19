package model

type TourExecutionStatus string

const (
	StatusActive    TourExecutionStatus = "active"
	StatusCompleted TourExecutionStatus = "completed"
	StatusAbandoned TourExecutionStatus = "abandoned"
)

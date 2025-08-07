package tasks

import "github.com/bogi-lyceya-44/task-tracker/internal/app/models"

func mapStorageTaskToDomain(task Task) models.Task {
	return models.Task{
		ID:           task.ID,
		Name:         task.Name,
		Description:  task.Description,
		Dependencies: task.Dependencies,
		Priority:     models.Priority(task.Priority),
		StartTime:    task.StartTime,
		FinishTime:   task.FinishTime,
		CreatedAt:    task.CreatedAt,
		UpdatedAt:    task.UpdatedAt,
	}
}

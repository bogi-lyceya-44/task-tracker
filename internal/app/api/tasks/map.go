package tasks

import (
	"fmt"

	"github.com/bogi-lyceya-44/task-tracker/internal/app/models"
	desc "github.com/bogi-lyceya-44/task-tracker/internal/pb/api/tasks"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO: make bidirectional map
var protoPriorityToDomain = map[desc.Priority]models.Priority{
	desc.Priority_PRIORITY_LOW:      models.PriorityLow,
	desc.Priority_PRIORITY_MEDIUM:   models.PriorityMedium,
	desc.Priority_PRIORITY_HIGH:     models.PriorityHigh,
	desc.Priority_PRIORITY_CRITICAL: models.PriorityCritical,
}

var domainPriorityToProto = map[models.Priority]desc.Priority{
	models.PriorityUndefined: desc.Priority_PRIORITY_UNSPECIFIED,
	models.PriorityLow:       desc.Priority_PRIORITY_LOW,
	models.PriorityMedium:    desc.Priority_PRIORITY_MEDIUM,
	models.PriorityHigh:      desc.Priority_PRIORITY_HIGH,
	models.PriorityCritical:  desc.Priority_PRIORITY_CRITICAL,
}

func MapTaskPriorityToDomain(priority desc.Priority) (models.Priority, error) {
	value, ok := protoPriorityToDomain[priority]
	if !ok {
		return models.PriorityUndefined, fmt.Errorf("incorrect priority value: %v", priority)
	}

	return value, nil
}

func MapDomainTaskPriorityToProto(priority models.Priority) (desc.Priority, error) {
	value, ok := domainPriorityToProto[priority]
	if !ok {
		return desc.Priority_PRIORITY_UNSPECIFIED, fmt.Errorf("undefined priority value: %v", priority)
	}

	return value, nil
}

func MapCreateTaskPrototypeToDomain(task *desc.CreateTasksRequest_TaskPrototype) (models.Task, error) {
	priority, err := MapTaskPriorityToDomain(task.GetPriority())
	if err != nil {
		return models.Task{}, errors.Wrap(err, "mapping create task prototype")
	}

	return models.Task{
		Name:         task.GetName(),
		Description:  task.GetDescription(),
		Dependencies: task.GetDeps(),
		Priority:     priority,
		StartTime:    task.GetStartTime().AsTime().UTC(),
		FinishTime:   task.GetFinishTime().AsTime().UTC(),
	}, nil
}

func MapUpdateTaskPrototypeToUpdatedTask(task *desc.UpdateTasksRequest_TaskPrototype) (models.UpdatedTask, error) {
	result := models.UpdatedTask{ID: task.GetId()}

	if task.Name != nil {
		result.Name = task.Name
	}

	if task.Description != nil {
		result.Description = task.Description
	}

	if task.Deps != nil {
		result.Dependencies = append(result.Dependencies, task.Deps...)
	}

	if task.Priority != nil {
		priority, err := MapTaskPriorityToDomain(task.GetPriority())
		if err != nil {
			return models.UpdatedTask{}, errors.Wrap(err, "mapping update task prototype")
		}

		result.Priority = &priority
	}

	if task.StartTime != nil {
		temp := task.StartTime.AsTime().UTC()
		result.StartTime = &temp
	}

	if task.FinishTime != nil {
		temp := task.FinishTime.AsTime().UTC()
		result.FinishTime = &temp
	}

	return result, nil
}

func MapDomainTaskToProto(task models.Task) (*desc.Task, error) {
	priority, err := MapDomainTaskPriorityToProto(task.Priority)
	if err != nil {
		return nil, errors.Wrap(err, "mapping domain task to proto")
	}

	return &desc.Task{
		Id:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Deps:        task.Dependencies,
		Priority:    priority,
		StartTime:   timestamppb.New(task.StartTime),
		FinishTime:  timestamppb.New(task.FinishTime),
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
	}, nil
}

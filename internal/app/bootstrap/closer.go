package bootstrap

import "github.com/bogi-lyceya-44/common/pkg/closer"

const (
	CloserGroupApp           = "app"
	CloserGroupConnections   = "connections"
	CloserGroupGlobalContext = "global context"
)

const (
	HighPriority = iota
	MediumPriority
	LowPriority
)

func InitCloser() {
	closer.AddGroups([]closer.Group{
		{
			Name:     CloserGroupApp,
			Priority: HighPriority,
		},
		{
			Name:     CloserGroupConnections,
			Priority: MediumPriority,
		},
		{
			Name:     CloserGroupGlobalContext,
			Priority: LowPriority,
		},
	}...)
}

package tasks

const (
	tableName = "tasks"

	columnID           = "id"
	columnName         = "name"
	columnDescription  = "description"
	columnDependencies = "dependencies"
	columnPriority     = "priority"
	columnStartTime    = "start_time"
	columnFinishTime   = "finish_time"
	columnCreatedAt    = "created_at"
	columnUpdatedAt    = "updated_at"
	columnIsDeleted    = "is_deleted"
)

var allColumns = []string{
	columnID,
	columnName,
	columnDescription,
	columnDependencies,
	columnPriority,
	columnStartTime,
	columnFinishTime,
	columnCreatedAt,
	columnUpdatedAt,
	columnIsDeleted,
}

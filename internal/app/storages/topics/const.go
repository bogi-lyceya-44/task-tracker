package topics

const (
	tableName = "topics"

	columnID        = "id"
	columnName      = "name"
	columnTaskIDs   = "task_ids"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnIsDeleted = "is_deleted"
)

var allColumns = []string{
	columnID,
	columnName,
	columnTaskIDs,
	columnCreatedAt,
	columnUpdatedAt,
	columnIsDeleted,
}

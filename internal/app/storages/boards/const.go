package boards

const (
	tableName = "boards"

	columnID        = "id"
	columnName      = "name"
	columnTopicIDs  = "topic_ids"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnIsDeleted = "is_deleted"
)

var allColumns = []string{
	columnID,
	columnName,
	columnTopicIDs,
	columnCreatedAt,
	columnUpdatedAt,
	columnIsDeleted,
}

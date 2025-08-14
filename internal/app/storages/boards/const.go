package boards

const (
	boardsTableName = "boards"
	orderTableName  = "board_order"

	columnID        = "id"
	columnName      = "name"
	columnTopicIDs  = "topic_ids"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnIsDeleted = "is_deleted"
	columnPlace     = "place"
)

var allColumns = []string{
	columnID,
	columnName,
	columnTopicIDs,
	columnCreatedAt,
	columnUpdatedAt,
	columnIsDeleted,
}

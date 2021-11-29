package models

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "users"
}

func (Todo) TableName() string {
	return "todos"
}
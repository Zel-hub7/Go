package main

// User represents a user model.
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100);unique"`
}

// TableName overrides the table name used by User to `users`.
func (User) TableName() string {
	return "users"
}

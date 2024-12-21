package user

// User represents the user entity.
type User struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

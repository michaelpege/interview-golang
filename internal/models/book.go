package models

type Book struct {
	ID     string `gorm:"primaryKey,default:uuid_generate_v3()" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	// CreatedAt time.Time `default:time.Now() son:"created_at"
	// UpdatedAt time.Time `default:time.Now() son:"updated_at"
}

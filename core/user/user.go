package core

type (
	BasicUser struct {
		// gorm.Model
		ID        uint `gorm:"primarykey"`
		CreatedAt uint64
		UserID    string `gorm:"size:36;unique_index"`
		Name      string `gorm:"size:64; unique_index" json:"name,omitempty"`
		Phone     string `gorm:"size:14;index" json:"phone,omitempty"`
		Email     string `gorm:"unique_index" json:"email,omitempty"`
		Password  string `gorm:"size:20" json:"password,omitempty"`
		Lang      string `gorm:"size:36;default:'zh'" json:"lang,omitempty"`
		Avatar    string `gorm:"size:255" json:"avatar,omitempty"`
		State     string `gorm:"size:24;default:'formal'" json:"state,omitempty"`
	}
)

package core

type (
	Message struct {
		// gorm.Model
		ID        uint   `gorm:"primarykey"`
		CreatedAt uint64 `gorm:"default:0"`
		Phone     string `gorm:"size:14;index" json:"phone,omitempty"`
		Type      uint8  `gorm:"type:tinyint(1);default:0" json:"type,omitempty"`
		Content   string `gorm:"size:512"`
		Code      string `gorm:"size:6"`
	}
)

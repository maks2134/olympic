package models

type Article struct {
	ID          uint      `gorm:"primary_key"`
	KnowledgeID uint      `gorm:"not null"`
	Knowledge   Knowledge `gorm:"foreignKey:KnowledgeID"`
	Category    string    `gorm:"not null"`
	ViewsCount  uint      `gorm:"default:0"`
}

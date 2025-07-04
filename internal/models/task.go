package models

type Task struct {
	ID          uint      `gorm:"primary_key"`
	KnowledgeID uint      `gorm:"not null"`
	Knowledge   Knowledge `gorm:"foreignKey:KnowledgeID"`
	Difficulty  string    `gorm:"not null"`
	TimeLimit   uint      `gorm:"not null"`
	MemoryLimit uint      `gorm:"not null"`
	TestCases   []TestCase
	Solutions   []Solution
	OlympiadID  uint
	Olympiad    Olympiad `gorm:"foreignKey:OlympiadID"`
}

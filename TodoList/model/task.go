package model

import "github.com/jinzhu/gorm"

// 任务模型
type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`	// 指定外键，指向用户的备忘录，使它的外键关联到user表
	Uid       uint  `gorm:"not null"`			// 备忘录的编号，因为需要外键关联，所以不能为空
	Title     string `gorm:"index;not null"`	// 备忘录的标题，因为需要外键关联，所以不能为空
	Status    int    `gorm:"default:0"`			// 备忘录的状态；0是未完成，1是已完成
	Content   string `gorm:"type:longtext"`		// 备忘录内容
	StartTime int64								// 备忘录开始时间，传过来是一个时间戳
	EndTime   int64 `gorm:"default:0"`			// 备忘录的结束时间
}
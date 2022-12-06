package model

func migration() {
	// 自动迁移模式，AutoMigrate将user和task表里的代码映射到数据库
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})

	// 添加外键关联,AddForeignKey函数将task表里的uid关联到了user表里的id,后面两个update和delete的时候进行级联更新或是级联删除
	DB.Model(&Task{}).AddForeignKey("uid","user(id)","CASCADE","CASCADE")
}

package init

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitDB to init database
func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/cloudwego?charset=utf8mb4&parseTime=True&loc=Local"
	// global mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}
	return db
}

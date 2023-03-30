package pkg

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Article struct {
	ID      int64  `gorm:"primarykey"`
	Title   string `gorm:"type:varchar(40)"`
	Content string `gorm:"type:varchar(255)"`
	Author  int64
}

// BeforeCreate uses snowflake to generate an ID.
func (a *Article) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if a.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(1)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	a.ID = sf.Generate().Int64()
	return nil
}

type ArticleManager struct {
	db *gorm.DB
}

// NewArticleManager creates a mysql manager.
func NewArticleManager(db *gorm.DB) *ArticleManager {
	m := db.Migrator()
	if !m.HasTable(&Article{}) {
		if err := m.CreateTable(&Article{}); err != nil {
			panic(err)
		}
	}
	return &ArticleManager{
		db: db,
	}
}

func (a *ArticleManager) Post(title, content string, uid int64) error {
	var atc Article
	atc.Title = title
	atc.Content = content
	atc.Author = uid
	err := a.db.Create(&atc).Error
	if err != nil {
		return err
	}
	return nil
}

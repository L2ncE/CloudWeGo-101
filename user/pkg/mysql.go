package mysql

import (
	"errors"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	ID         int64  `gorm:"primarykey"`
	Username   string `gorm:"type:varchar(40)"`
	Password   string `gorm:"type:varchar(40)"`
	ArticleNum int64
}

// BeforeCreate uses snowflake to generate an ID.
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if u.ID != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(0)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	u.ID = sf.Generate().Int64()
	return nil
}

type UserManager struct {
	db *gorm.DB
}

// NewUserManager creates a mysql manager.
func NewUserManager(db *gorm.DB) *UserManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	return &UserManager{
		db: db,
	}
}

func (m *UserManager) LoginCheck(username, password string) (bool, int64, error) {
	u, err := m.GetUserByName(username)
	if err != nil {
		return false, -1, err
	}
	if u.Password != password {
		return false, -1, nil
	}
	return true, u.ID, nil
}

func (m *UserManager) CreateUser(username, password string) error {
	var u User
	_, err := m.GetUserByName(username)
	if err == gorm.ErrRecordNotFound {
		u.Username = username
		u.Password = password
		err = m.db.Create(&u).Error
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	return errors.New("user exits")
}

func (m *UserManager) GetArticleNum(uid int64) (int64, error) {
	var count int64
	err := m.db.Model(User{}).Where(&User{ID: uid}).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (m *UserManager) ArticleNumPlusOne(uid int64) error {
	err := m.db.Model(User{}).Where(&User{ID: uid}).Update("article_num", gorm.Expr("article_num + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *UserManager) GetUserById(uid int64) (*User, error) {
	var user User
	err := m.db.Where(&User{ID: uid}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserManager) GetUserByName(name string) (*User, error) {
	var user User
	err := m.db.Where(&User{Username: name}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

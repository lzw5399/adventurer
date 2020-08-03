package model

import (
	"github.com/jinzhu/gorm"
)

// user model本体
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;comment:'用户登录名'"`
	Password string `json:"password" gorm:"not null;comment:'用户登录密码'"`
	Gender   int    `json:"gender" gorm:"comment:'性别男1女2'"`
	Phone    int    `json:"phone" gorm:"comment:'手机号'"`
	Email    string `json:"email" gorm:"comment:'邮箱'"`
}

//type SysUser struct {
//	gorm.Model
//	UUID        uuid.UUID    `json:"uuid" gorm:"comment:'用户UUID'"`
//	Username    string       `json:"userName" gorm:"comment:'用户登录名'"`
//	Password    string       `json:"-"  gorm:"comment:'用户登录密码'"`
//	NickName    string       `json:"nickName" gorm:"default:'系统用户';comment:'用户昵称'" `
//	HeaderImg   string       `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"`
//	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;comment:'用户角色'"`
//	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
//}

func (User) TableName() string {
	return "sys.user"
}

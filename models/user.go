package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
)

// User struct
type User struct {
	Id             int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Avatar         string `orm:"column(avatar);size(255)" description:"头像" json:"avatar"`
	Username       string `orm:"column(username);size(30)" description:"用户名" json:"username"`
	FullName       string `orm:"column(fullName);size(30)" description:"姓名" json:"fullName"`
	FullNamePinYin string `orm:"column(fullNamePinYin);size(30)" description:"姓名拼音" json:"fullNamePinYin"`
	Mobile         string `orm:"column(mobile);size(11)" description:"手机号" json:"mobile"`
	UserLevelId    int    `orm:"column(user_level_id);size(11);default(1)" description:"用户等级" json:"user_level_id"`
	Password       string `orm:"column(password);size(255)" description:"密码" json:"password"`
	Status         int8   `orm:"column(status);size(1);default(1)" description:"是否启用 0：否 1：是" json:"status"`
	Description    string `orm:"column(description);type(text)" description:"描述" json:"description"`
	CreateTime     int    `orm:"column(create_time);size(10);default(0)" description:"操作时间" json:"create_time"`
	UpdateTime     int    `orm:"column(update_time);size(10);default(0)" description:"更新时间" json:"update_time"`
	DeleteTime     int    `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

// TableName 自定义table 名称
func (*User) TableName() string {
	return "user"
}

// SearchField 定义模型的可搜索字段
func (*User) SearchField() []string {
	return []string{"username", "mobile", "nickname"}
}

// NoDeletionId 禁止删除的数据id
func (*User) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*User) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*User) TimeField() []string {
	return []string{}
}

// GetSignStr 获取加密字符串，用在登录的时候加密处理
func (u *User) GetSignStr(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", u.Id, u.Username, ua))))
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(User))
}

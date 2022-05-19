package models

import "github.com/beego/beego/v2/client/orm"

type Quota struct {
	Id           int `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	UniversityId int `orm:"column(university_id);size(11)" description:"高校ID" json:"universityId"`
	TrainPlanId  int `orm:"column(plan_id);size(11)" description:"培训计划ID" json:"trainPlanId"`
	Quota        int `orm:"column(quota);size(30)" description:"名额" json:"quota"`
	CreateTime   int `orm:"column(create_time);size(10);default(0)" description:"操作时间" json:"create_time"`
	UpdateTime   int `orm:"column(update_time);size(10);default(0)" description:"更新时间" json:"update_time"`
	DeleteTime   int `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

// TableName 自定义table 名称
func (*Quota) TableName() string {
	return "quota"
}

// SearchField 定义模型的可搜索字段
func (*Quota) SearchField() []string {
	return []string{"id", "university_id", "plan_id", "quota"}
}

// NoDeletionId 禁止删除的数据id
func (*Quota) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*Quota) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*Quota) TimeField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(Quota))
}

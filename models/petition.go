package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

// Petition struct 培训计划报名表
type Petition struct {
	Id           int       `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	CreatedAt    time.Time `orm:"column(createdAt);" description:"创建时间" json:"createdAt"`
	UpdatedAt    time.Time `orm:"column(updatedAt);" description:"创建时间" json:"updatedAt"`
	UserId       int       `orm:"column(user_id);size(11)" description:"用户ID" json:"userId"`
	PlanId       int       `orm:"column(plan_id);size(11)" description:"培训计划ID" json:"planId"`
	UniversityId int       `orm:"column(university_id);size(11)" description:"高校ID" json:"universityId"`
	Name         string    `orm:"column(name);size(100)" description:"姓名" json:"name"`
	NamePY       string    `orm:"column(namePY);size(100)" description:"姓名（拼音）" json:"namePY"`
	IdCardNum    string    `orm:"column(idCardNum);" description:"身份证号" json:"IdCardNum"`
	IdCardFront  string    `orm:"column(idCardFront);size(255);" description:"身份证正面" json:"idCardFront"`
	IdCardBack   string    `orm:"column(idCardBack);size(255);" description:"身份证反面" json:"idCardBack"`
	IdCardComb   string    `orm:"column(idCardComb);size(255);" description:"身份证和人合一" json:"idCardComb"`
	Tel          string    `orm:"column(tel);size(50)" description:"联系方式" json:"tel"`
	Email        string    `orm:"column(email);size(50)" description:"电子邮箱" json:"email"`
	Remark       string    `orm:"column(remark);type(text)" description:"备注" json:"remark"`
	Status       int8      `orm:"column(status);size(1)" description:"申请状态 0：待通过 1：通过 2：拒绝" json:"status"`
	DeletedAt    time.Time `orm:"column(deleteAt)" description:"删除时间" json:"delete_time"`
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(Petition))
}

// TableName 自定义table 名称
func (*Petition) TableName() string {
	return "petition"
}

// SearchField 定义模型的可搜索字段
func (*Petition) SearchField() []string {
	return []string{"name", "namePY", "idCardNum", "tel", "email", "remark"}
}

// NoDeletionId 禁止删除的数据id
func (*Petition) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*Petition) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*Petition) TimeField() []string {
	return []string{}
}

// GetSignStrByTrain 获取加密字符串，用在登录的时候加密处理
func (m *Petition) GetSignStrByTrain(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", m.Id, m.Name, ua))))
}

// GetAuthUrl 获取已授权url
func (m *Petition) GetAuthUrl() map[string]interface{} {
	//FIXME:var (
	//	urlArr orm.ParamsList
	//)
	authURL := make(map[string]interface{})

	//o := orm.NewOrm()
	//qs := o.QueryTable(new(AdminRole))

	//_, err := qs.Filter("id__in", strings.Split(Petition.Role, ",")).Filter("status", 1).ValuesFlat(&urlArr, "url")
	//if err == nil {
	//	urlIDStr := ""
	//	for k, row := range urlArr {
	//		urlStr, ok := row.(string)
	//		if ok {
	//			if k == 0 {
	//				urlIDStr = urlStr
	//			} else {
	//				urlIDStr += "," + urlStr
	//			}
	//		}
	//	}
	//	urlIDArr := strings.Split(urlIDStr, ",")
	//
	//	var authURLArr orm.ParamsList
	//
	//	if len(urlIDStr) > 0 {
	//		o = orm.NewOrm()
	//		qs = o.QueryTable(new(AdminMenu))
	//		_, err := qs.Filter("id__in", urlIDArr).ValuesFlat(&authURLArr, "url")
	//		if err == nil {
	//			for k, row := range authURLArr {
	//				val, ok := row.(string)
	//				if ok {
	//					authURL[val] = k
	//				}
	//			}
	//		}
	//	}
	//	return authURL
	//}
	return authURL
}

// GetShowMenu 获取当前用户已授权的显示菜单
func (m *Petition) GetShowMenu() map[int]orm.Params {
	var maps []orm.Params
	returnMaps := make(map[int]orm.Params)
	o := orm.NewOrm()

	if m.Id == 1 {
		_, err := o.QueryTable(new(AdminMenu)).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		if err == nil {
			for _, m := range maps {
				returnMaps[int(m["Id"].(int64))] = m
			}
			return returnMaps
		}
		return map[int]orm.Params{}
	}

	//FIXME:var list orm.ParamsList
	//_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(Petition.Role, ",")).Filter("status", 1).ValuesFlat(&list, "url")
	//if err == nil {
	//	var urlIDArr []string
	//	for _, m := range list {
	//		urlIDArr = append(urlIDArr, strings.Split(m.(string), ",")...)
	//	}
	//	_, err := o.QueryTable(new(AdminMenu)).Filter("id__in", urlIDArr).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
	//	if err == nil {
	//		for _, m := range maps {
	//			returnMaps[int(m["Id"].(int64))] = m
	//		}
	//		return returnMaps
	//	}
	//	return map[int]orm.Params{}
	//}
	return map[int]orm.Params{}

}

// GetRoleText 用户角色名称
func (m *Petition) GetRoleText() map[int]*AdminRole {
	//FIXME:roleIDArr := strings.Split(Petition.Role, ",")
	var adminRole []*AdminRole
	//_, err := orm.NewOrm().QueryTable(new(AdminRole)).Filter("id__in", roleIDArr, "id", "name").All(&adminRole)
	//if err != nil {
	//	return nil
	//}
	adminRoleMap := make(map[int]*AdminRole)
	for _, v := range adminRole {
		adminRoleMap[v.Id] = v
	}
	return adminRoleMap
}

// GetTrain 获取所有用户
func (*Petition) GetTrain() []*Petition {
	var Trains []*Petition
	_, err := orm.NewOrm().QueryTable(new(Petition)).All(&Trains)
	if err == nil {
		return Trains
	}
	return nil
}

package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

// TrainPlan struct
type TrainPlan struct {
	Id                    int       `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	CreatedAt             time.Time `orm:"column(createdAt);" description:"创建时间" json:"createdAt"`
	UpdatedAt             time.Time `orm:"column(updatedAt);" description:"创建时间" json:"updatedAt"`
	Title                 string    `orm:"column(title);size(255)" description:"用户名" json:"title"`
	Summary               string    `orm:"column(summary);type(text)" description:"密码" json:"summary"`
	RegistrationStartedAt time.Time `orm:"column(registrationStartedAt);" description:"报名开始时间" json:"registrationStartedAt"`
	RegistrationEndAt     time.Time `orm:"column(registrationEndAt);" description:"报名结束时间" json:"registrationEndAt"`
	PersonInCharge        string    `orm:"column(personInCharge);" description:"负责人" json:"personInCharge"`
	Status                int8      `orm:"column(status);size(1)" description:"是否启用 0：否 1：是" json:"status"`
	DeleteTime            int       `orm:"column(delete_time);size(10);default(0)" description:"删除时间" json:"delete_time"`
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(TrainPlan))
}

// TableName 自定义table 名称
func (*TrainPlan) TableName() string {
	return "train_plan"
}

// SearchField 定义模型的可搜索字段
func (*TrainPlan) SearchField() []string {
	return []string{"nickname", "username"}
}

// NoDeletionId 禁止删除的数据id
func (*TrainPlan) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*TrainPlan) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*TrainPlan) TimeField() []string {
	return []string{}
}

// GetSignStrByTrain 获取加密字符串，用在登录的时候加密处理
func (Train *TrainPlan) GetSignStrByTrain(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", Train.Id, Train.Title, ua))))
}

// GetAuthUrl 获取已授权url
func (Train *TrainPlan) GetAuthUrl() map[string]interface{} {
	//FIXME:var (
	//	urlArr orm.ParamsList
	//)
	authURL := make(map[string]interface{})

	//o := orm.NewOrm()
	//qs := o.QueryTable(new(AdminRole))

	//_, err := qs.Filter("id__in", strings.Split(TrainPlan.Role, ",")).Filter("status", 1).ValuesFlat(&urlArr, "url")
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
func (Train *TrainPlan) GetShowMenu() map[int]orm.Params {
	var maps []orm.Params
	returnMaps := make(map[int]orm.Params)
	o := orm.NewOrm()

	if Train.Id == 1 {
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
	//_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(TrainPlan.Role, ",")).Filter("status", 1).ValuesFlat(&list, "url")
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
func (Train *TrainPlan) GetRoleText() map[int]*AdminRole {
	//FIXME:roleIDArr := strings.Split(TrainPlan.Role, ",")
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
func (*TrainPlan) GetTrain() []*TrainPlan {
	var Trains []*TrainPlan
	_, err := orm.NewOrm().QueryTable(new(TrainPlan)).All(&Trains)
	if err == nil {
		return Trains
	}
	return nil
}

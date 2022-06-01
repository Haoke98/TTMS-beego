package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

// FavorPlan 培训计划收藏记录
type FavorPlan struct {
	Id        int       `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	CreatedAt time.Time `orm:"column(createdAt);" description:"创建时间" json:"createdAt"`
	UpdatedAt time.Time `orm:"column(updatedAt);" description:"创建时间" json:"updatedAt"`
	UserId    int       `orm:"column(user_id);size(11)" description:"用户ID" json:"UserId"`
	PlanId    int       `orm:"column(plan_id);size(11)" description:"培训计划ID" json:"planId"`
}

//在init义的model
func init() {
	orm.RegisterModel(new(FavorPlan))
}

// TableName 自定义table 名称
func (*FavorPlan) TableName() string {
	return "favor_plans"
}

// SearchField 定义模型的可搜索字段
func (*FavorPlan) SearchField() []string {
	return []string{"title", "personInCharge"}
}

// NoDeletionId 禁止删除的数据id
func (*FavorPlan) NoDeletionId() []int {
	return []int{}
}

// WhereField 定义模型可作为条件的字段
func (*FavorPlan) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*FavorPlan) TimeField() []string {
	return []string{}
}

// GetSignStrByTrain 获取加密字符串，用在登录的时候加密处理
func (af *FavorPlan) GetSignStrByTrain(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s", af.Id, ua))))
}

// GetAuthUrl 获取已授权url
func (af *FavorPlan) GetAuthUrl() map[string]interface{} {
	//FIXME:var (
	//	urlArr orm.ParamsList
	//)
	authURL := make(map[string]interface{})

	//o := orm.NewOrm()
	//qs := o.QueryTable(new(AdminRole))

	//_, err := qs.Filter("id__in", strings.Split(FavorPlan.Role, ",")).Filter("status", 1).ValuesFlat(&urlArr, "url")
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
func (af *FavorPlan) GetShowMenu() map[int]orm.Params {
	var maps []orm.Params
	returnMaps := make(map[int]orm.Params)
	o := orm.NewOrm()

	if af.Id == 1 {
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
	//_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(FavorPlan.Role, ",")).Filter("status", 1).ValuesFlat(&list, "url")
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
func (af *FavorPlan) GetRoleText() map[int]*AdminRole {
	//FIXME:roleIDArr := strings.Split(FavorPlan.Role, ",")
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
func (*FavorPlan) GetTrain() []*FavorPlan {
	var Trains []*FavorPlan
	_, err := orm.NewOrm().QueryTable(new(FavorPlan)).All(&Trains)
	if err == nil {
		return Trains
	}
	return nil
}

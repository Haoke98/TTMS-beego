package services

import (
	"beego-admin/formvalidate"
	"beego-admin/models"
	"beego-admin/utils"
	"beego-admin/utils/page"
	"encoding/base64"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"net/url"
	"time"
)

// TrainPlanService 培训计划服务
type TrainPlanService struct {
	BaseService
}

// GetById 根据id获取一条admin_user数据
func (*TrainPlanService) GetById(id int) *models.TrainPlan {
	o := orm.NewOrm()
	train := models.TrainPlan{Id: id}
	err := o.Read(&train)
	if err != nil {
		return nil
	}
	return &train
}

func (s *TrainPlanService) GetTempTrainPlan() *models.TrainPlan {
	var tempFrom = formvalidate.TrainPlanForm{
		Title:                 "",
		RegistrationStartedAt: time.Now(),
		RegistrationEndAt:     time.Now(),
		Status:                3,
	}
	o := orm.NewOrm()
	plan := models.TrainPlan{Status: 3}
	err := o.QueryTable(new(models.TrainPlan)).Filter("Status", 3).One(&plan)
	if err != nil {
		insertID, tempPlan := s.Create(&tempFrom)
		fmt.Printf("[indertID:%+v, tempPlan.ID:%+v]", insertID, tempPlan.Id)
		return &plan
	} else {
		plan.Title = tempFrom.Title
		plan.RegistrationStartedAt = tempFrom.RegistrationStartedAt
		plan.RegistrationEndAt = tempFrom.RegistrationEndAt
		return &plan
	}
}

// AuthCheck 权限检测
func (*TrainPlanService) AuthCheck(url string, authExcept map[string]interface{}, loginUser *models.AdminUser) bool {
	authURL := loginUser.GetAuthUrl()
	if utils.KeyInMap(url, authExcept) || utils.KeyInMap(url, authURL) {
		return true
	}
	return false
}

// GetCount 获取admin_user 总数
func (*TrainPlanService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllAdminUser 获取所有adminuser
func (*TrainPlanService) GetAllAdminUser() []*models.AdminUser {
	var adminUser []*models.AdminUser
	o := orm.NewOrm().QueryTable(new(models.AdminUser))
	_, err := o.All(&adminUser)
	if err != nil {
		return nil
	}
	return adminUser
}

// UpdateNickName 系统管理-个人资料-修改昵称
func (*TrainPlanService) UpdateNickName(id int, nickname string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"nickname": nickname,
	})
	if err != nil || num <= 0 {
		return 0
	}
	return int(num)
}

// UpdatePassword 修改密码
func (*TrainPlanService) UpdatePassword(id int, newPassword string) int {
	newPasswordForHash, err := utils.PasswordHash(newPassword)

	if err != nil {
		return 0
	}

	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"password": base64.StdEncoding.EncodeToString([]byte(newPasswordForHash)),
	})

	if err != nil || num <= 0 {
		return 0
	}

	return int(num)
}

// UpdateAvatar 系统管理-个人资料-修改头像
func (*TrainPlanService) UpdateAvatar(id int, avatar string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"avatar": avatar,
	})
	if err != nil || num <= 0 {
		return 0
	}
	return int(num)
}

// GetPaginateData 通过分页获取培训计划
func (ts *TrainPlanService) GetPaginateData(listRows int, params url.Values) ([]*models.TrainPlan, page.Pagination) {
	//搜索、查询字段赋值
	ts.SearchField = append(ts.SearchField, new(models.TrainPlan).SearchField()...)

	var trains []*models.TrainPlan
	o := orm.NewOrm().QueryTable(new(models.TrainPlan))
	_, err := ts.PaginateAndScopeWhere(o, listRows, params).Exclude("status", 3).All(&trains)
	if err != nil {
		return nil, ts.Pagination
	}
	return trains, ts.Pagination
}

// IsExistName 名称验重
func (*TrainPlanService) IsExistName(name string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.TrainPlan)).Filter("title", name).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.TrainPlan)).Filter("title", name).Exclude("id", id).Exist()
}

// Create 新增培训计划
func (*TrainPlanService) Create(form *formvalidate.TrainPlanForm) (int, models.TrainPlan) {
	trainPlan := models.TrainPlan{
		Title:                 form.Title,
		Summary:               form.Summary,
		RegistrationStartedAt: form.RegistrationStartedAt,
		RegistrationEndAt:     form.RegistrationEndAt,
		PersonInCharge:        form.PersonInCharge,
		Status:                int8(form.Status),
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
	id, err := orm.NewOrm().Insert(&trainPlan)

	if err == nil {
		return int(id), trainPlan
	} else {
		fmt.Println(err)
	}
	return 0, trainPlan
}

// Update 更新培训计划
func (*TrainPlanService) Update(form *formvalidate.TrainPlanForm) int {
	o := orm.NewOrm()
	train := models.TrainPlan{Id: form.Id}
	if o.Read(&train) == nil {
		train.Title = form.Title
		train.Summary = form.Summary
		train.PersonInCharge = form.PersonInCharge
		train.Status = int8(form.Status)
		train.RegistrationStartedAt = form.RegistrationStartedAt
		train.RegistrationEndAt = form.RegistrationEndAt
		num, err := o.Update(&train)
		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable 启用培训计划
func (*TrainPlanService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable 禁用培训计划
func (*TrainPlanService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del 删除培训计划
func (*TrainPlanService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

func (s *TrainPlanService) IsFavor(planId, userId int) bool {
	o := orm.NewOrm()
	var fp *models.FavorPlan
	err := o.QueryTable(new(models.FavorPlan)).Filter("plan_id", planId).Filter("user_id", userId).One(&fp)
	if err == nil && fp != nil {
		return true
	}
	return false
}

func (s *TrainPlanService) Favor(planId, userId int) int {
	fp := models.FavorPlan{UserId: userId, PlanId: planId}
	id, err := orm.NewOrm().Insert(&fp)
	if err == nil {
		return int(id)
	} else {
		fmt.Println(err)
		return -1
	}
}

func (s *TrainPlanService) Abolish(planId, userId int) int {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.FavorPlan)).Filter("plan_id", planId).Filter("user_id", userId).Delete()
	if err != nil {
		fmt.Println(err)
		return -1
	} else {
		return int(num)
	}
}

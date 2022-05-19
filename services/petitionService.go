package services

import (
	"beego-admin/formvalidate"
	"beego-admin/models"
	"beego-admin/utils/page"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"net/url"
	"time"
)

type PetitionService struct {
	BaseService
}
type PetitionDTO struct {
	models.Petition
	University models.University
}

// GetPaginateData 通过分页获取培训计划所关联的所有报名申请
func (s *PetitionService) GetPaginateData(listRows int, params url.Values) ([]*PetitionDTO, page.Pagination) {
	var (
		pdtos             []*PetitionDTO
		universityService UniversityService
		planIdStr         string
	)
	//搜索、查询字段赋值
	s.SearchField = append(s.SearchField, new(models.Petition).SearchField()...)

	var petitions []*models.Petition
	o := orm.NewOrm().QueryTable(new(models.Petition))
	if params.Has("id") && !params.Has("plan_id") {
		planIdStr = params.Get("id")
		params.Del("id")
		params.Add("plan_id", planIdStr)
	}
	_, err := s.PaginateAndScopeWhere(o, listRows, params).All(&petitions)
	if err != nil {
		return nil, s.Pagination
	}
	for _, p := range petitions {
		u := universityService.GetUniversityById(p.UniversityId)
		dto := PetitionDTO{Petition: *p, University: *u}
		pdtos = append(pdtos, &dto)
	}
	return pdtos, s.Pagination
}

// GetById 根据培训计划和高校ID获取Quota记录
func (*PetitionService) GetByUniversityIdAndPlanId(universityId, planId int) *models.Quota {
	o := orm.NewOrm()
	obj := models.Quota{}
	err := o.QueryTable(new(models.Quota)).Filter("university_id", universityId).Filter("plan_id", planId).One(&obj)
	if err != nil {
		return nil
	}
	return &obj
}

// Create 新增培训计划
func (*PetitionService) Create(form *formvalidate.QuotaForm) (int, models.Quota) {
	obj := models.Quota{
		TrainPlanId:  form.PlanId,
		UniversityId: form.UniversityId,
		Quota:        form.Quota,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	id, err := orm.NewOrm().Insert(&obj)

	if err == nil {
		return int(id), obj
	} else {
		fmt.Println(err)
	}
	return 0, obj
}

// Update 更新培训计划
func (*PetitionService) Update(form *formvalidate.QuotaForm) int {
	//o := orm.NewOrm()
	//obj := models.Quota{Id: form.Id}
	//if o.Read(&obj) == nil {
	//	obj.Name = form.Name
	//	obj.Summary = form.Summary
	//	obj.PersonInCharge = form.PersonInCharge
	//	obj.Status = int8(form.Status)
	//	obj.RegistrationStartedAt = form.RegistrationStartedAt
	//	obj.RegistrationEndAt = form.RegistrationEndAt
	//	num, err := o.Update(&obj)
	//	if err == nil {
	//		return int(num)
	//	}
	//	return 0
	//}
	return 0
}

// Update 更新或者新增Quota记录
func (s *PetitionService) CreateOrUpdate(form *formvalidate.QuotaForm) int {
	o := orm.NewOrm()
	var obj models.Quota
	err := o.QueryTable(new(models.Quota)).Filter("plan_id", form.PlanId).Filter("university_id", form.UniversityId).One(&obj)
	if err != nil {
		//TODO 新增
		num, _ := s.Create(form)
		return num
	} else {
		//TODO 更新
		obj.UpdatedAt = time.Now()
		obj.Quota = form.Quota
		num, err := o.Update(&obj)
		if err == nil {
			return int(num)
		}
		return 0
	}
}

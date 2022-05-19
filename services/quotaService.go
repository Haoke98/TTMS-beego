package services

import (
	"beego-admin/formvalidate"
	"beego-admin/models"
	"beego-admin/utils/page"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"net/url"
	"strconv"
	"time"
)

type QuotaService struct {
	BaseService
}
type QuotaDTO struct {
	models.Quota
	University models.University
}

// GetPaginateData 通过分页获取培训计划
func (s *QuotaService) GetPaginateData(listRows int, params url.Values) ([]*QuotaDTO, page.Pagination) {
	var (
		qdtos             []*QuotaDTO
		universityService UniversityService
		planIdStr         string
	)
	if params.Has("plan_id") {
		planIdStr = params.Get("plan_id")
	} else {
		planIdStr = params.Get("id")
		params.Del("id")
		params.Add("plan_id", planIdStr)
	}
	us, pagination := universityService.GetPaginateData(listRows, params)
	for _, u := range us {
		planId, err := strconv.Atoi(planIdStr)
		if err == nil {
			obj := s.GetByUniversityIdAndPlanId(u.Id, planId)
			qdto := QuotaDTO{Quota: models.Quota{UniversityId: u.Id}, University: *u}
			if obj != nil {
				qdto.Quota = *obj
			}
			qdtos = append(qdtos, &qdto)
		} else {
			fmt.Println(err)
		}
	}
	return qdtos, pagination
}

// GetById 根据培训计划和高校ID获取Quota记录
func (*QuotaService) GetByUniversityIdAndPlanId(universityId, planId int) *models.Quota {
	o := orm.NewOrm()
	obj := models.Quota{}
	err := o.QueryTable(new(models.Quota)).Filter("university_id", universityId).Filter("plan_id", planId).One(&obj)
	if err != nil {
		return nil
	}
	return &obj
}

// Create 新增培训计划
func (*QuotaService) Create(form *formvalidate.QuotaForm) (int, models.Quota) {
	obj := models.Quota{
		TrainPlanId:  form.TrainPlanId,
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
func (*QuotaService) Update(form *formvalidate.QuotaForm) int {
	//o := orm.NewOrm()
	//obj := models.Quota{Id: form.Id}
	//if o.Read(&obj) == nil {
	//	obj.Title = form.Title
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
func (s *QuotaService) CreateOrUpdate(form *formvalidate.QuotaForm) int {
	o := orm.NewOrm()
	var obj models.Quota
	err := o.QueryTable(new(models.Quota)).Filter("plan_id", form.TrainPlanId).Filter("university_id", form.UniversityId).One(&obj)
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

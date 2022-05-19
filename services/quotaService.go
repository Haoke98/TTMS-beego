package services

import (
	"beego-admin/models"
	"beego-admin/utils/page"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"net/url"
	"strconv"
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

// GetAdminUserById 根据id获取一条admin_user数据
func (*QuotaService) GetByUniversityIdAndPlanId(universityId, planId int) *models.Quota {
	o := orm.NewOrm()
	obj := models.Quota{}
	err := o.QueryTable(new(models.Quota)).Filter("university_id", universityId).Filter("plan_id", planId).One(&obj)
	if err != nil {
		return nil
	}
	return &obj
}

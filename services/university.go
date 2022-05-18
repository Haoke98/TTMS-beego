package services

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/utils"
	"beego-admin/utils/page"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"net/url"
	"strconv"
	"time"
)

// UniversityService 培训计划服务
type UniversityService struct {
	BaseService
}

// GetUniversityById 根据id获取一条admin_user数据
func (*UniversityService) GetUniversityById(id int) *models.University {
	o := orm.NewOrm()
	obj := models.University{Id: id}
	err := o.Read(&obj)
	if err != nil {
		return nil
	}
	return &obj
}

// AuthCheck 权限检测
func (*UniversityService) AuthCheck(url string, authExcept map[string]interface{}, loginUser *models.AdminUser) bool {
	authURL := loginUser.GetAuthUrl()
	if utils.KeyInMap(url, authExcept) || utils.KeyInMap(url, authURL) {
		return true
	}
	return false
}

// CheckLogin 用户登录验证
func (*UniversityService) CheckLogin(loginForm formvalidate.LoginForm, ctx *context.Context) (*models.AdminUser, error) {
	var adminUser models.AdminUser
	o := orm.NewOrm()
	err := o.QueryTable(new(models.AdminUser)).Filter("username", loginForm.Username).Limit(1).One(&adminUser)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	decodePasswdStr, err := base64.StdEncoding.DecodeString(adminUser.Password)

	if err != nil || !utils.PasswordVerify(loginForm.Password, string(decodePasswdStr)) {
		return nil, errors.New("密码错误")
	}

	if adminUser.Status != 1 {
		return nil, errors.New("用户被冻结")
	}

	ctx.Output.Session(global.LOGIN_USER, adminUser)

	if loginForm.Remember != "" {
		ctx.SetCookie(global.LOGIN_USER_ID, strconv.Itoa(adminUser.Id), 7200)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, adminUser.GetSignStrByAdminUser(ctx), 7200)
	} else {
		ctx.SetCookie(global.LOGIN_USER_ID, ctx.GetCookie(global.LOGIN_USER_ID), -1)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, ctx.GetCookie(global.LOGIN_USER_ID_SIGN), -1)
	}

	return &adminUser, nil

}

// GetCount 获取admin_user 总数
func (*UniversityService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllAdminUser 获取所有adminuser
func (*UniversityService) GetAllAdminUser() []*models.AdminUser {
	var adminUser []*models.AdminUser
	o := orm.NewOrm().QueryTable(new(models.AdminUser))
	_, err := o.All(&adminUser)
	if err != nil {
		return nil
	}
	return adminUser
}

// UpdateNickName 系统管理-个人资料-修改昵称
func (*UniversityService) UpdateNickName(id int, nickname string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"nickname": nickname,
	})
	if err != nil || num <= 0 {
		return 0
	}
	return int(num)
}

// UpdatePassword 修改密码
func (*UniversityService) UpdatePassword(id int, newPassword string) int {
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
func (*UniversityService) UpdateAvatar(id int, avatar string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"avatar": avatar,
	})
	if err != nil || num <= 0 {
		return 0
	}
	return int(num)
}

// GetPaginateData 通过分页获取培训计划
func (ts *UniversityService) GetPaginateData(listRows int, params url.Values) ([]*models.University, page.Pagination) {
	//搜索、查询字段赋值
	ts.SearchField = append(ts.SearchField, new(models.University).SearchField()...)

	var objs []*models.University
	o := orm.NewOrm().QueryTable(new(models.University))
	_, err := ts.PaginateAndScopeWhere(o, listRows, params).All(&objs)
	if err != nil {
		return nil, ts.Pagination
	}
	return objs, ts.Pagination
}

// IsExistName 名称验重
func (*UniversityService) IsExistName(title string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.University)).Filter("name", title).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.University)).Filter("name", title).Exclude("id", id).Exist()
}

// IsExistCode 编号验重
func (*UniversityService) IsExistCode(code string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.University)).Filter("code", code).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.University)).Filter("code", code).Exclude("id", id).Exist()
}

// Create 新增培训计划
func (*UniversityService) Create(form *formvalidate.UniversityForm) int {
	obj := models.University{
		Name:      form.Name,
		Code:      form.Code,
		Badge:     form.Badge,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := orm.NewOrm().Insert(&obj)

	if err == nil {
		return int(id)
	} else {
		fmt.Println(err)
	}
	return 0
}

// Update 更新高校信息
func (*UniversityService) Update(form *formvalidate.UniversityForm) int {
	o := orm.NewOrm()
	obj := models.University{Id: form.Id}
	if o.Read(&obj) == nil {
		obj.Name = form.Name
		obj.Code = form.Code
		obj.Badge = form.Badge
		num, err := o.Update(&obj)
		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable 启用培训计划
func (*UniversityService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable 禁用培训计划
func (*UniversityService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del 删除培训计划
func (*UniversityService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.University)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

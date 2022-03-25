package services

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/utils"
	"beego-admin/utils/page"
	"encoding/base64"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"net/url"
	"strconv"
	"time"
)

// TrainService 培训计划服务
type TrainService struct {
	BaseService
}

// GetAdminUserById 根据id获取一条admin_user数据
func (*TrainService) GetAdminUserById(id int) *models.AdminUser {
	o := orm.NewOrm()
	adminUser := models.AdminUser{Id: id}
	err := o.Read(&adminUser)
	if err != nil {
		return nil
	}
	return &adminUser
}

// AuthCheck 权限检测
func (*TrainService) AuthCheck(url string, authExcept map[string]interface{}, loginUser *models.AdminUser) bool {
	authURL := loginUser.GetAuthUrl()
	if utils.KeyInMap(url, authExcept) || utils.KeyInMap(url, authURL) {
		return true
	}
	return false
}

// CheckLogin 用户登录验证
func (*TrainService) CheckLogin(loginForm formvalidate.LoginForm, ctx *context.Context) (*models.AdminUser, error) {
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
func (*TrainService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetAllAdminUser 获取所有adminuser
func (*TrainService) GetAllAdminUser() []*models.AdminUser {
	var adminUser []*models.AdminUser
	o := orm.NewOrm().QueryTable(new(models.AdminUser))
	_, err := o.All(&adminUser)
	if err != nil {
		return nil
	}
	return adminUser
}

// UpdateNickName 系统管理-个人资料-修改昵称
func (*TrainService) UpdateNickName(id int, nickname string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"nickname": nickname,
	})
	if err != nil || num <= 0 {
		return 0
	}
	return int(num)
}

// UpdatePassword 修改密码
func (*TrainService) UpdatePassword(id int, newPassword string) int {
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
func (*TrainService) UpdateAvatar(id int, avatar string) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id", id).Update(orm.Params{
		"avatar": avatar,
	})
	if err != nil || num <= 0 {
		return 0
	}
	return int(num)
}

// GetPaginateData 通过分页获取adminuser
func (ts *TrainService) GetPaginateData(listRows int, params url.Values) ([]*models.Train, page.Pagination) {
	//搜索、查询字段赋值
	ts.SearchField = append(ts.SearchField, new(models.Train).SearchField()...)

	var trains []*models.Train
	o := orm.NewOrm().QueryTable(new(models.Train))
	_, err := ts.PaginateAndScopeWhere(o, listRows, params).All(&trains)
	if err != nil {
		return nil, ts.Pagination
	}
	return trains, ts.Pagination
}

// IsExistName 名称验重
func (*TrainService) IsExistName(username string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("username", username).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("username", username).Exclude("id", id).Exist()
}

// Create 新增admin user用户
func (*TrainService) Create(form *formvalidate.TrainForm) int {
	train := models.Train{
		Title:                 form.Title,
		Summary:               form.Summary,
		RegistrationStartedAt: form.RegistrationStartedAt,
		RegistrationEndAt:     form.RegistrationEndAt,
		PersonInCharge:        form.PersonInCharge,
		Status:                int8(form.Status),
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
	id, err := orm.NewOrm().Insert(&train)

	if err == nil {
		return int(id)
	}
	return 0
}

// Update 更新用户信息
func (*TrainService) Update(form *formvalidate.AdminUserForm) int {
	o := orm.NewOrm()
	adminUser := models.AdminUser{Id: form.Id}
	if o.Read(&adminUser) == nil {
		adminUser.Username = form.Username
		adminUser.Nickname = form.Nickname
		adminUser.Role = form.Role
		adminUser.Status = int8(form.Status)
		if adminUser.Password != form.Password {
			newPasswordForHash, err := utils.PasswordHash(form.Password)
			if err == nil {
				adminUser.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))
			}
		}
		num, err := o.Update(&adminUser)
		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable 启用用户
func (*TrainService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable 禁用用户
func (*TrainService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del 删除用户
func (*TrainService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

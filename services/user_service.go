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

// UserService struct
type UserService struct {
	BaseService
}

// GetPaginateData 通过分页获取user
func (us *UserService) GetPaginateData(listRows int, params url.Values) ([]*models.User, page.Pagination) {
	//搜索、查询字段赋值
	us.SearchField = append(us.SearchField, new(models.User).SearchField()...)

	var users []*models.User
	o := orm.NewOrm().QueryTable(new(models.User))
	_, err := us.PaginateAndScopeWhere(o, listRows, params).All(&users)
	if err != nil {
		return nil, us.Pagination
	}
	return users, us.Pagination
}

// Create 新增用户
func (*UserService) Create(form *formvalidate.UserForm) int {
	user := models.User{
		Username:       form.Username,
		FullName:       form.FullName,
		FullNamePinYin: form.FullNamePinyin,
		Mobile:         form.Mobile,
		Description:    form.Description,
		Status:         form.Status,
		CreateTime:     int(time.Now().Unix()),
		UpdateTime:     int(time.Now().Unix()),
	}
	if form.Avatar != "" {
		user.Avatar = form.Avatar
	}

	//密码加密
	newPasswordForHash, err := utils.PasswordHash(form.Password)
	if err != nil {
		return 0
	}
	user.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))

	id, err := orm.NewOrm().Insert(&user)

	if err == nil {
		return int(id)
	}
	return 0
}

// GetUserById 根据id获取一条user数据
func (*UserService) GetUserById(id int) *models.User {
	o := orm.NewOrm()
	user := models.User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return nil
	}
	return &user
}

// Update 更新用户
func (*UserService) Update(form *formvalidate.UserForm) int {
	o := orm.NewOrm()
	user := models.User{Id: form.Id}
	if o.Read(&user) == nil {

		//判断密码是否相等
		if user.Password != form.Password {
			newPasswordForHash, err := utils.PasswordHash(form.Password)
			if err == nil {
				user.Password = base64.StdEncoding.EncodeToString([]byte(newPasswordForHash))
			}
		}

		user.Username = form.Username
		user.FullName = form.FullName
		user.FullNamePinYin = form.FullNamePinyin
		user.Mobile = form.Mobile
		user.Description = form.Description
		user.Status = int8(form.Status)
		user.UpdateTime = int(time.Now().Unix())

		if form.Avatar != "" {
			user.Avatar = form.Avatar
		}
		num, err := o.Update(&user)

		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Enable 启用
func (*UserService) Enable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Disable 禁用
func (*UserService) Disable(ids []int) int {
	num, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Update(orm.Params{
		"status": 0,
	})
	if err == nil {
		return int(num)
	}
	return 0
}

// Del 删除
func (*UserService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData 获取导出数据
func (us *UserService) GetExportData(params url.Values) []*models.User {
	//搜索、查询字段赋值
	us.SearchField = append(us.SearchField, new(models.User).SearchField()...)
	var user []*models.User
	o := orm.NewOrm().QueryTable(new(models.User))
	_, err := us.ScopeWhere(o, params).All(&user)
	if err != nil {
		return nil
	}
	return user
}

// CheckLogin 用户登录验证
func (*UserService) CheckLogin(loginForm formvalidate.LoginForm, ctx *context.Context) (*models.User, error) {
	var user models.User
	o := orm.NewOrm()
	err := o.QueryTable(new(models.User)).Filter("username", loginForm.Username).Limit(1).One(&user)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	decodePasswdStr, err := base64.StdEncoding.DecodeString(user.Password)

	if err != nil || !utils.PasswordVerify(loginForm.Password, string(decodePasswdStr)) {
		return nil, errors.New("密码错误")
	}

	if user.Status != 1 {
		return nil, errors.New("用户被冻结")
	}

	ctx.Output.Session(global.LOGIN_USER, user)

	if loginForm.Remember == true {
		ctx.SetCookie(global.LOGIN_USER_ID, strconv.Itoa(user.Id), 7200)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, user.GetSignStr(ctx), 7200)
	} else {
		ctx.SetCookie(global.LOGIN_USER_ID, ctx.GetCookie(global.LOGIN_USER_ID), -1)
		ctx.SetCookie(global.LOGIN_USER_ID_SIGN, ctx.GetCookie(global.LOGIN_USER_ID_SIGN), -1)
	}

	return &user, nil

}

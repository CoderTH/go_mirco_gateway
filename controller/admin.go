package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/gin_scaffold/dao"
	"github.com/e421083458/gin_scaffold/dto"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.GET("/admin_info", adminLogin.AdminInfo)
	group.POST("/change_pwd", adminLogin.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (*AdminLoginController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessionInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), adminSessionInfo)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1.读取sessionKey对应的json转换为结构体
	//2. 取出数据然后封装输出结构体

	out := &dto.AdminInfoOutput{
		Id:           adminSessionInfo.Id,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "touxiang",
		Introduction: "adsdas",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (*AdminLoginController) ChangePwd(c *gin.Context) {

	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1.session中读取用户信息到结构体sessInfo
	//2.用sessInfo.ID读取数据库信息adminInfo
	//3.params.password + adminInfo.salt sha256 saltPassword
	//4.saltPassword ==> adminInfo.password 执行数据库保存

	sess := sessions.Default(c)
	sessionInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), adminSessionInfo)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	//从数据库中读取adminInfo
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, (&dao.Admin{UserName: adminSessionInfo.UserName}))
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	//生成新的加盐密码
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.PassWord = saltPassword
	//执行数据保存
	err = adminInfo.Save(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}

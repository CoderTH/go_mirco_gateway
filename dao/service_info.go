package dao

import (
	"github.com/e421083458/gin_scaffold/dto"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type ServiceInfo struct {
	Id        int64       `json:"id" gorm:"primary_key" description:"自增主键"`
	LoadType  int    `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName  string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc  string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否以删除 0：否 1：是"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (t *ServiceInfo)PageList(c *gin.Context,tx *gorm.DB,param *dto.ServiceListInput) ([]ServiceInfo,int64,error) {
	//初始化总数
	total :=int64(0)
	//初始化list
	list :=[]ServiceInfo{}
	//偏移量，从第几页开始
	offset := (param.PageNo-1)*param.PageSize
	//设置gin上下文
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query =query.Table(t.TableName()).Where("is_delete=0")
	if param.Info!="" {
		query =query.Where("(service_name like ? or service_desc like ?)","%"+param.Info+"%","%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error;err!=nil&&err!=gorm.ErrRecordNotFound{
		return nil ,0,err
	}
	query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Count(&total)
	return list,total,nil

}

func (t *ServiceInfo) ServiceDetail(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceDetail, error) {
	if search.ServiceName == "" {
			info,err :=t.Find(c,tx,search)
		if err != nil {
			return nil, err
		}
			search = info
	}
	httpRule := &HttpRule{ServiceID: search.Id}
	httpRule,err := httpRule.Find(c, tx, httpRule)
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil, err
	}
	tcpRule := &TcpRule{ServiceID: search.Id}
	tcpRule,err = tcpRule.Find(c, tx, tcpRule)
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	grpcRule := &GrpcRule{ServiceID: search.Id}
	grpcRule,err = grpcRule.Find(c, tx, grpcRule)
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	accessControl := &AccessControl{ServiceID: search.Id}
	accessControl,err = accessControl.Find(c, tx, accessControl)
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	loadBalance := &LoadBalance{ServiceID: search.Id}
	loadBalance,err = loadBalance.Find(c, tx, loadBalance)
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil, err
	}

	detail := &ServiceDetail{
		Info:          search,
		HTTPRule:      httpRule,
		TCPRule:       tcpRule,
		GRPCRule: grpcRule,
		LoadBalance:   loadBalance,
		AccessControl: accessControl,
	}
	return detail,nil
}

func (t *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
//更新保存
func (t *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}

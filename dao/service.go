package dao

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info"  description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http"  description:"http_url"`
	TCPRule       *TcpRule       `json:"tcp"  description:"tcp_url"`
	GRPCRule      *GrpcRule      `json:"grpc"  description:"grpc_url"`
	LoadBalance   *LoadBalance   `json:"load_balance"  description:"grpc_url"`
	AccessControl *AccessControl `json:"access_control"  description:"grpc_url"`
}

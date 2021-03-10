package dao

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info"  description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http"  description:"http_url"`
	TCPRule       *TcpRule       `json:"tcp"  description:"tcp_url"`
	GRPCRule      *GrpcRule      `json:"grpc"  description:"grpc_url"`
	LoadBalance   *LoadBalance   `json:"grpc"  description:"grpc_url"`
	AccessControl *AccessControl `json:"grpc"  description:"grpc_url"`
}

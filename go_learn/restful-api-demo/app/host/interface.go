package host

import "context"

type Service interface {
	//录入主机信息
	CreateHost(context.Context, *Host) (*Host, error)
	//查询主机列表
	QUeryHost(context.Context, *QueryHostRequest) (*Set, error)
	//主机详情查询
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	//主机信息修改
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	//删除主机
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

type Host struct{}

type QueryHostRequest struct {
	PageSize   int
	PageNumber int
}

type Set struct {
	Total int64
	Items []*Host
}

type DescribeHostRequest struct {
	Id string
}

type UpdateHostRequest struct {
	Id string
}

type DeleteHostRequest struct {
	Id string
}

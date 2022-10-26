package mrecord

import "builder/builder/pbs/recordpb"

type StatusType string

var (
	Init     StatusType = "Init"
	Building StatusType = "Building"
	Success  StatusType = "Success"
	Failed   StatusType = "Failed"
)

func (r StatusType) ToPBEnum() recordpb.StatusType {
	// TODO: StatusType 和 recordpb.StatusType 是一一对应的，
	// 所以这里不做检查，如果对应不上，程序会直接 Panic 退出
	status := recordpb.StatusType_value[string(r)]
	return recordpb.StatusType(status)
}

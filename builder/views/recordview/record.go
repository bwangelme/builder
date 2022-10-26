package recordview

import (
	"builder/builder/models/mrecord"
	"builder/builder/pbs/recordpb"
	"builder/builder/views/helper"

	"github.com/gin-gonic/gin"
)

func ListRecord(c *gin.Context) {
	req := &recordpb.CreateRecordReq{}
	err := helper.ReadJsonReq(req, c)
	if err != nil {
		return
	}

	resp := &recordpb.ListRecordResp{
		Error: "",
		Data:  &recordpb.ListRecordResp_RecordContainer{},
	}
	helper.WriteResp(resp, c)
}

func CreateRecord(c *gin.Context) {
	req := &recordpb.CreateRecordReq{}
	err := helper.ReadJsonReq(req, c)
	if err != nil {
		return
	}

	record, err := mrecord.Add(req.Giturl, req.Revision, req.Image, req.Tag)
	if err != nil {
		helper.AbortWith500(err, c)
		return
	}

	resp := &recordpb.CreateRecordResp{
		Error: "",
		Data: &recordpb.CreateRecordResp_RecordContainer{
			Record: record,
		},
	}

	helper.WriteResp(resp, c)
}

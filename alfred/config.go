package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	Language    = "QRCODE_SIZE"
	TimeFormat  = "TIME_FORMAT"
	TimeZone    = "TIME_ZONE"
	LiDefault   = "LI_DEFAULT"
	UuidDefault = "UUID_DEFAULT"
	Debug       = "DEBUG"
)

func GetQrcodeSize(wf *aw.Workflow) string {
	return wf.Config.Get(Language)
}

func GetTimeFormat(wf *aw.Workflow) string {
	return wf.Config.Get(TimeFormat)
}

func GetTimeZone(wf *aw.Workflow) string {
	return wf.Config.Get(TimeZone)
}

func GetLiDefault(wf *aw.Workflow) string {
	return wf.Config.Get(LiDefault)
}

func GetUuidDefault(wf *aw.Workflow) string {
	return wf.Config.Get(UuidDefault)
}

func GetDebug(wf *aw.Workflow) bool {
	return wf.Config.GetBool(Debug)
}

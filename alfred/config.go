package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	language   = "QRCODE_SIZE"
	timeFormat = "TIME_FORMAT"
	timeZone   = "TIME_ZONE"
)

func GetQrcodeSize(wf *aw.Workflow) string {
	return wf.Config.Get(language)
}

func GetTimeFormat(wf *aw.Workflow) string {
	return wf.Config.Get(timeFormat)
}

func GetTimeZone(wf *aw.Workflow) string {
	return wf.Config.Get(timeZone)
}

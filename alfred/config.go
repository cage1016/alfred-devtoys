package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	language = "QRCODE_SIZE"
)

func GetQrcodeSize(wf *aw.Workflow) string {
	return wf.Config.Get(language)
}

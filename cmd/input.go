package cmd

import (
	"fmt"

	aw "github.com/deanishe/awgo"
)

const maxInputSize = 64 * 1024

func inputTooLarge(query string, icon *aw.Icon) bool {
	if len(query) <= maxInputSize {
		return false
	}
	wf.NewItem("Input is too large").Subtitle(fmt.Sprintf("Try an input up to %d KB", maxInputSize/1024)).Icon(icon)
	wf.SendFeedback()
	return true
}

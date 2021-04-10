package logic

import (
	"log"

	"github.com/daheige/goapp/internal/pkg/helper"
)

// HomeLogic home logic
type HomeLogic struct {
	BaseLogic
}

// GetData get data
func (h *HomeLogic) GetData() []string {
	log.Println(helper.GetStringByCtx(h.Ctx, "current_uid"))
	return []string{
		"golang",
		"php",
	}
}

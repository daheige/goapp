package logic

import (
	"log"

	"github.com/daheige/goapp/pkg/helper"
)

type HomeLogic struct {
	BaseLogic
}

func (h *HomeLogic) GetData() []string {
	log.Println(helper.GetStringByCtx(h.Ctx, "current_uid"))

	return []string{
		"golang",
		"php",
	}
}

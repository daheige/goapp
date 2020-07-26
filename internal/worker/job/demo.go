package job

import (
	"log"

	"github.com/daheige/goapp/pkg/helper"
)

// TestJob test job.
type TestJob struct {
	BaseJob
}

// Info info test get key from ctx.
func (j *TestJob) Info() {
	id := helper.GetStringByCtx(j.ctx, "id")
	log.Println("current id: ", id)
}

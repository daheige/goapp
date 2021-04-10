package task

import "log"

// TestTask test task
type TestTask struct {
	BaseTask
}

// Hello hello task
func (t *TestTask) Hello() {
	log.Println("hello world")
}

package key

import "fmt"

func RdsTaskNum(taskId string) string {
	return fmt.Sprintf("task_id_%s",taskId)
}

package machinery

import (
	"github.com/RichardKnop/machinery/v2/tasks"
	"time"
)

func Sum(args []int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}

	return sum, tasks.NewErrRetryTaskLater("重试", 4 * time.Second)

}

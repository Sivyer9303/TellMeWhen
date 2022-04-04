package query

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReminderQuery_GetAllReminder(t *testing.T) {
	rq := GetRemindQuery()
	reminder := rq.GetAllReminder()
	for _, v := range reminder {
		marshal, err := json.Marshal(v)
		if err != nil {
			fmt.Println("err...")
			continue
		}
		fmt.Println(string(marshal))
	}
}

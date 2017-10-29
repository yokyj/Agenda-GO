package main

import (
	"fmt"
	"time"
	"tpisntgod/GO-Agenda/entity/meeting"
)

func main() {
	var parti []string
	parti = append(parti, "qqq")
	parti = append(parti, "www")
	parti = append(parti, "eee")
	t1, _ := time.Parse("2006-01-02 15:04:05", "2017-10-29 08:37:18")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2017-10-29 09:37:28")
	err := meeting.CreateMeeting("meet1", parti, t1, t2)
	fmt.Println(err)
	meeting.WriteMeetingInfo()
}

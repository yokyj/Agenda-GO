package meeting

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type meeting_records struct {
	Title        string
	Host         string
	Participator []string
	StartTime    time.Time
	EndTime      time.Time
}

var meetings []meeting_records
var curUser string
var isFileExist = true
var writeFilePath = "./MeetingInfo.json"

//Meeting comment
/*
type Meeting struct {
	CreateMeeting              func(title string, participator []string, startTime time.Time, endTime time.Time) error
	ChangeMeetingParticipators func(title string, participator []string, action int) error
	QueryMeeting               func(startTime time.Time, endTime time.Time) error
	CancelMeeting              func(title string) error
	QuitMeeting                func(title string) error
	ClearAllMeeting            func()
}
*/

//只是判断两个时间段是否overlap
func checkIfMeetingTimeOverlap(meetingStartTime, meetingEndTime, startTime, endTime time.Time) bool {
	if (meetingStartTime.Before(startTime) || meetingStartTime.Equal(startTime)) &&
		meetingEndTime.After(startTime) && (meetingEndTime.Before(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	if (meetingStartTime.Before(startTime) || meetingStartTime.Equal(startTime)) &&
		(meetingEndTime.After(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	if (meetingStartTime.After(startTime) || meetingStartTime.Equal(startTime)) &&
		meetingStartTime.Before(endTime) && (meetingEndTime.After(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	if (meetingStartTime.After(startTime) || meetingStartTime.Equal(startTime)) &&
		(meetingEndTime.Before(endTime) || meetingEndTime.Equal(endTime)) {
		return true
	}
	return false
}

func checkIfTwoMeetingTimeOverlap(title string, participator []string, startTime time.Time, endTime time.Time) int {
	for i := 0; i < len(meetings); i++ {
		//先判断两个会议是否时间overlap，然后判断有没有同时参加这两个会议的人
		if checkIfMeetingTimeOverlap(meetings[i].StartTime, meetings[i].EndTime, startTime, endTime) {
			PeopleToCheck := &meetings[i].Participator
			for j := 0; j < len(*PeopleToCheck); j++ {
				for k := 0; k < len(participator); k++ {
					if (*PeopleToCheck)[j] == participator[k] {
						return 2
					}
				}
			}
		}
	}
	return 0
}

//CreateMeeting 创建会议
func CreateMeeting(title string, participator []string, startTime time.Time, endTime time.Time) error {
	//先判断两个会议是否时间overlap，然后判断有没有同时参加这两个会议的人
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			return errors.New("会议title已经被注册")
		}
	}
	isOverlap := checkIfTwoMeetingTimeOverlap(title, participator, startTime, endTime)
	/*
		if isOverlap == 1 {
			return errors.New("会议title已经被注册")
		}
	*/
	if isOverlap == 2 {
		return errors.New("会议时间有冲突")
	}
	var meetingToAdd meeting_records
	meetingToAdd.Title = title
	meetingToAdd.Host = curUser
	//meetingToAdd.participator = make([]string, 5)
	for i := 0; i < len(participator); i++ {
		meetingToAdd.Participator = append(meetingToAdd.Participator, participator[i])
	}
	meetingToAdd.StartTime = startTime
	meetingToAdd.EndTime = endTime
	fmt.Println("0:meetingtoadd", meetingToAdd)
	meetings = append(meetings, meetingToAdd)
	fmt.Println("1:\n", meetings)
	return nil
}

//ChangeMeetingParticipators 增删会议参与者
func ChangeMeetingParticipators(title string, participator []string, action int) error {
	var isMeetingExist = 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			//添加会议参与者
			if action == 1 {
				for j := 0; j < len(meetings[i].Participator); j++ {
					for k := 0; k < len(participator); k++ {
						if meetings[i].Participator[j] == participator[k] {
							return errors.New("此title的会议已经有参与者")
						}
					}
				}
				//会议是否被注册处要修改
				//还要做时间重叠判断（允许仅有端点重叠的情况）
				isOverlap := checkIfTwoMeetingTimeOverlap(title, participator, meetings[i].StartTime, meetings[i].EndTime)
				if isOverlap == 2 {
					return errors.New("新增会议参与者和某个会议时间冲突")
				}
				for j := 0; j < len(participator); j++ {
					meetings[i].Participator = append(meetings[i].Participator, participator[j])
				}
			} else {
				//删除会议参与者
				var NumToDelete int
				NumToDelete = 0
				for j := 0; j < len(meetings[i].Participator); j++ {
					for k := 0; k < len(participator); k++ {
						if meetings[i].Participator[j] == participator[k] {
							NumToDelete++
						}
					}
				}
				if NumToDelete < len(participator) {
					return errors.New("将要删除的用户不存在")
				}
				//删除用户
				var partAfterDelete []string
				partAfterDelete = meetings[i].Participator
				for k := 0; k < len(participator); k++ {
					for j := 0; j < len(meetings[i].Participator); j++ {
						if meetings[i].Participator[j] == participator[k] {
							partAfterDelete = append(partAfterDelete[:j], partAfterDelete[j+1:]...)
						}
					}
				}
				//如果删除参与者后会议人数为0，删除该会议
				if len(partAfterDelete) == 0 {
					CancelMeeting(title)
				} else {
					meetings[i].Participator = partAfterDelete
				}
			}
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此title的会议不存在")
	}
	return nil
}

//QueryMeeting 查询会议
func QueryMeeting(startTime time.Time, endTime time.Time) error {
	isOverlap := 0
	for i := 0; i < len(meetings); i++ {
		if checkIfMeetingTimeOverlap(meetings[i].StartTime, meetings[i].EndTime, startTime, endTime) {
			isOverlap++
			if isOverlap == 1 {
				fmt.Println("指定时间范围内找到的所有会议安排")
				fmt.Println("会议主题：  起始时间：  终止时间：  发起者：  参与者：")
			}
			fmt.Printf("%v %v %v %v", meetings[i].Title, meetings[i].StartTime, meetings[i].EndTime, meetings[i].Host)
			for j := 0; j < len(meetings[i].Participator); j++ {
				fmt.Printf("%v ", meetings[i].Participator[j])
			}
			fmt.Println()
		}
	}
	if isOverlap == 0 {
		return errors.New("此时间段并没有你的会议安排")
	}
	return nil
}

//CancelMeeting 取消会议
func CancelMeeting(title string) error {
	isMeetingExist := 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			isMeetingExist = 1
			if meetings[i].Host == curUser {
				meetings = append(meetings[:i], meetings[i+1:]...)
				break
			} else {
				return errors.New("此会议的发起者并不是你")
			}
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此主题的会议不存在")
	}
	return nil
}

//QuitMeeting 退出会议
func QuitMeeting(title string) error {
	isMeetingExist := 0
	isAttendMeeting := 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			isMeetingExist = 1
			for j := 0; j < len(meetings[i].Participator); j++ {
				if meetings[i].Participator[j] == curUser {
					isAttendMeeting = 1
					meetings[i].Participator = append(meetings[i].Participator[:j], meetings[i].Participator[j+1:]...)
					if len(meetings[i].Participator) == 0 {
						CancelMeeting(title)
					}
					break
				}
			}
			break
		}
	}
	if isMeetingExist == 0 {
		return errors.New("此主题的会议不存在")
	}
	if isAttendMeeting == 0 {
		return errors.New("你并没有参加此会议")
	}
	return nil
}

//ClearAllMeeting 清空会议
func ClearAllMeeting() {
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Host == curUser {
			meetings = append(meetings[:i], meetings[i+1:]...)
			i--
		}
	}
}

//WriteMeetingInfo 将会议信息以JSON格式写入文件
func WriteMeetingInfo() {
	fmt.Println("2:\n", meetings)
	fmt.Println("len:", len(meetings), "cap:", cap(meetings))

	b, err := json.Marshal(meetings)
	if err != nil {
		panic(err)
	}
	/*
		if !isFileExist {
			_, err2 := os.Create("MeetingInfo")
			if err2 != nil {
				panic(err2)
			}
		}
	*/
	fmt.Println("b:\n", b)
	/*
		_, err = os.Open(writeFilePath)
		if err != nil {
			os.Create(writeFilePath)
		}
	*/
	err = ioutil.WriteFile(writeFilePath, b, 0644)
	if err != nil {
		panic(err)
	}
}

//SetCurrentUser 设置当前用户
func SetCurrentUser(currentUser string) {
	curUser = currentUser
}

func init() {
	/*
		meetings = make([]meeting_records, 5)
		for i := 0; i < len(meetings); i++ {
			meetings[i].participator = make([]string, 5)
		}
	*/
	curUser = "xxx"
	_, err2 := os.Stat(writeFilePath)
	if err2 == nil {
		data, err := ioutil.ReadFile(writeFilePath)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &meetings)
		if err != nil {
			panic(err)
		}
	} else {
		if os.IsNotExist(err2) {
			isFileExist = false
		} else {
			fmt.Println(errors.New("保存会议信息的文件打开失败"))
		}
	}
	fmt.Println("readfile:\n", meetings)
}

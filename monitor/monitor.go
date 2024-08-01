package monitor

import (
	"fmt"
	"net/http"
	"time"
)

// 要点
// 循环  time.sleep  time.Add
type webType struct {
	url  string
	code int
}
type Handle struct {
	sendEmial func(msg string)
}

type WebStatus struct {
	url         string
	status      string
	lastFailure time.Time
}

var websiteList = []webType{webType{url: "http://127.0.0.1:8080", code: 200}, webType{url: "https://www.zxzjhd.com/", code: 200}}

func Listen() {
	var webStatusList []WebStatus
	for {
		for _, web := range websiteList {
			body, err := http.Get(web.url)
			fmt.Println(body, err)
			if err != nil {
				fmt.Println("失败")
				alertUser(web.url, err, &webStatusList)
			} else if body.StatusCode != web.code {
				fmt.Println(body.StatusCode)
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func alertUser(web string, err error, webList *[]WebStatus) {
	downWebInfo := WebStatus{web, "down", time.Now()}
	if len(*webList) == 0 {
		*webList = append(*webList, downWebInfo)
		trigerEmial(web)
		return
	}
	isExitSend := checkForPreviousAlert(webList, web)
	if !isExitSend {
		*webList = append(*webList, downWebInfo)
		trigerEmial(web)
		return
	}
	if checkForReminderInterval(webList, web) {
		trigerEmial(web)
	}
}

func checkForPreviousAlert(webList *[]WebStatus, url string) bool {
	for _, item := range *webList {
		if item.url == url {
			return true
		}
	}
	return false
}
func checkForReminderInterval(webList *[]WebStatus, url string) bool {
	for i, item := range *webList {
		if item.url == url {
			// lastFailurePlusReminderMins := webStatusInfo.lastFailure.Add(time.Duration(reminderInterval) * time.Minute)
			lastKeepTime := item.lastFailure.Add(time.Duration(10) * time.Minute)
			if lastKeepTime.Before(time.Now()) {
				(*webList)[i] = WebStatus{url, "down", time.Now()}
				return true
			}
		}
	}
	return false
}
func trigerEmial(web string) {
	message := fmt.Sprintf("url is %s", web)
	fmt.Println(message)
}

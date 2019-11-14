package lang

import "time"

const second = 1000

//GetTonightZeroTime 获取今晚凌晨12点的时间
func GetTonightZeroTime() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, now.Location())
}

//GetTonightZeroTimestamp 获取今晚时间戳，毫秒为单位
func GetTonightZeroTimestamp() int64 {
	return GetTonightZeroTime().Unix() * second
}

//GetNowTimestamp 获取当前的timestamp 毫秒为单位
func GetNowTimestamp() int64 {
	return time.Now().Unix() * second
}

// GetYesterdayTimestamp 获取昨天的起始时间，单位毫秒
func GetYesterdayTimestamp() int64 {
	t := time.Now()
	date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	yesterday := date.AddDate(0, 0, -1)
	return yesterday.Unix() * second
}

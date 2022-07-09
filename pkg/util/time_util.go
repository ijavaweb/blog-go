package util

import "time"

func Time2Date(s int64) string  {
	var str string
	if s < 0 {
		return ""
	}
	tm:=time.Unix(s,0)
	str=tm.Format("2006-01-02")
	return str
}

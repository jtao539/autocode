package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

func TimeStamp2time(stamp int32) string {
	return time.Unix(int64(stamp), 0).Format("2006-01-02 15:04:05")
}

func Date2TimeStamp(date string) sql.NullInt32 {
	date = fmt.Sprintf("%s 00:00:00", date)
	tmp := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(tmp, date, time.Local)
	return IntToNullInt32(int(stamp.Unix()))
}

//JSONDecode ...
func JSONDecode(r io.Reader, obj interface{}) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return nil
}

func CheckStringNULL(args ...string) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			return true
		}
	}
	return false
}

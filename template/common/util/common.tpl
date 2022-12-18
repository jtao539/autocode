package util

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func TimeStamp2time(stamp int32) string {
	return time.Unix(int64(stamp), 0).Format("2006-01-02 15:04:05")
}

func TimeStamp2Date(stamp int32) string {
	return time.Unix(int64(stamp), 0).Format("2006-01-02 15:04:05")[:10]
}

func Time64Stamp2Date(stamp int64) string {
	return time.Unix(stamp, 0).Format("2006-01-02 15:04:05")[:10]
}

func Time64Stamp2time(stamp int64) string {
	return time.Unix(stamp, 0).Format("2006-01-02 15:04:05")
}

func Date2TimeStamp(date string) sql.NullInt32 {
	if strings.TrimSpace(date) == "" {
		return sql.NullInt32{Valid: false}
	}
	date = fmt.Sprintf("%s 00:00:00", date)
	tmp := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Asia/Shanghai")
	stamp, _ := time.ParseInLocation(tmp, date, local)
	return IntToNullInt32(int(stamp.Unix()))
}

func Date2Time64Stamp(date string) sql.NullInt64 {
	if strings.TrimSpace(date) == "" {
		return sql.NullInt64{Valid: false}
	}
	date = fmt.Sprintf("%s 00:00:00", date)
	tmp := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Asia/Shanghai")
	stamp, _ := time.ParseInLocation(tmp, date, local)
	return IntToNullInt64(stamp.Unix())
}

func Time2Time64Stamp(date string) sql.NullInt64 {
	if strings.TrimSpace(date) == "" {
		return sql.NullInt64{Valid: false}
	}
	date = fmt.Sprintf("%s", date)
	tmp := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Asia/Shanghai")
	stamp, _ := time.ParseInLocation(tmp, date, local)
	return IntToNullInt64(stamp.Unix())
}

func Date2TimeStampInt(date string) int {
	if strings.TrimSpace(date) == "" {
		return 0
	}
	date = fmt.Sprintf("%s 00:00:00", date)
	tmp := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Asia/Shanghai")
	stamp, _ := time.ParseInLocation(tmp, date, local)
	return int(stamp.Unix())
}

func Time2TimeStampInt(timeStr string) int {
	if strings.TrimSpace(timeStr) == "" {
		return 0
	}
	tmp := "2006-01-02 15:04:05"
	local, _ := time.LoadLocation("Asia/Shanghai")
	stamp, _ := time.ParseInLocation(tmp, timeStr, local)
	return int(stamp.Unix())
}

func TodayTimeStamp() int64 {
	year, month, day := time.Now().Date()
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 0, 0, 0, 0, location).Unix()
}

func TodayTime() time.Time {
	year, month, day := time.Now().Date()
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

func DayBetween(day, start, end time.Time) bool {
	return (day.Before(end) || day.Equal(end)) && (day.After(start) || day.Equal(start))
}

//JSONDecode ...
func JSONDecode(r io.Reader, obj interface{}) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return nil
}

func IntToNullInt32(a int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(a), Valid: true}
}

func IntToNullInt64(a int64) sql.NullInt64 {
	return sql.NullInt64{Int64: a, Valid: true}
}

func StringToNullString(a string) sql.NullString {
	return sql.NullString{String: a, Valid: true}
}

func FloatToNullFloat64(a float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: a, Valid: true}
}

func GetParentCode(districtCode string) string {
	var code string
	if strings.TrimSpace(districtCode) != "" {
		d, err := strconv.Atoi(districtCode)
		if err == nil {
			if d%10000 == 0 {
				code = strconv.Itoa(d / 10000)
			} else if d%100 == 0 {
				code = strconv.Itoa(d / 100)
			} else {
				code = strconv.Itoa(d)
			}
		} else {
			code = districtCode
		}
	}
	return code
}

// GenSalt 生产盐值
func GenSalt(n int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, n/2)
	rand.Read(result)
	return "$1$" + hex.EncodeToString(result)
}

// MD5Salt md5加盐加密
func MD5Salt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

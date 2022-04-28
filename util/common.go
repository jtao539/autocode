package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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

func test() {
	fmt.Println("getTmpDir（当前系统临时目录） = ", getTmpDir())
	fmt.Println("getCurrentAbPathByExecutable（仅支持go build） = ", getCurrentAbPathByExecutable())
	fmt.Println("getCurrentAbPathByCaller（仅支持go run） = ", getCurrentAbPathByCaller())
	fmt.Println("getCurrentAbPath（最终方案-全兼容） = ", getCurrentAbPath())
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

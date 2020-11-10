package Common

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func InitConfig() (err error) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	runPath, err := GetCurrentPath()

	Config.SetConfigType("json")

	Config.SetConfigName("config")
	Config.AddConfigPath(runPath)
	err = Config.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Println(runPath)
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}

	Config.WatchConfig()
	Config.OnConfigChange(func(in fsnotify.Event) {
		log.Println(os.Getpid(), "Config file changed:", in.Name, in.Op.String())
		err = Config.ReadInConfig()

		if err != nil { // Handle errors reading the config file
			log.Println(os.Getpid(), fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			log.Println(os.Getpid(), "reload config success")
		}
	})
	return
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	//if err != nil {
	//	return "", err
	//}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}

	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func FastSleepHour(sltime int) {
	log.Printf("sleep %d hours....", sltime)
	time.Sleep(time.Duration(sltime) * time.Hour)
}

func FastItoa(num int) string {
	return strconv.Itoa(num)
}

func FastAtoi(num string) int {
	ret, _ := strconv.Atoi(num)
	return ret
}

func FastJsonMarshal(_json interface{}) string {
	str, _ := Json.MarshalToString(_json)
	return str
}

func FastJsonMarshalIndent(_json interface{}) string {
	//str, _ := Json.MarshalToString(_json)
	str, err := Json.MarshalIndent(_json, "", " ")
	log.Println(err)
	return string(str)
}

func FastTimeParse(layout string, value string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(layout, value, loc)
	return t
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func SafeGetError(err error) string {
	if nil == err {
		return ""
	} else {
		return err.Error()
	}
}

// ArrayReverse array_reverse()
func SliceReverseString(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func SortPrintlnMapIntString(s map[int]string) {
	var keys []int
	for k, _ := range s {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, _k := range keys {

		CustomerLogPrintln(_k, ":::::", s[_k])
	}
}

func SortMapIntString(s map[int]string) map[int]string {
	var keys []int
	for k, _ := range s {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var newMap = make(map[int]string, len(s))
	for _, _k := range keys {
		newMap[_k] = s[_k]
	}

	return newMap
}

var IfTest = false

func CustomerLogPrintln(v ...interface{}) {
	if false == IfTest {
		log.Println(v)
	}
}

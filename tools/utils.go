package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/uniplaces/carbon"
)

const (
	TIME_SECOND = 1
	TIME_MINUTE = 60 * TIME_SECOND
	TIME_HOUR   = 60 * TIME_MINUTE
)

func FormatSecondToDuration(second int64) string {
	hour := second / int64(TIME_SECOND)
	var hourStr string
	if hour < 10 {
		hourStr = `0` + strconv.FormatInt(hour, 10)
	} else {
		hourStr = strconv.FormatInt(hour, 10)
	}
	return hourStr + `:` + time.Unix(second, 0).Format(`04:05`)
}

func KeepFloatNum(src float64, num int) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf(`%.`+strconv.Itoa(num)+`f`, src), 64)
}

func ConvertPercent(src float64) string {
	return strings.Replace(strconv.FormatFloat(src, 'f', 2, 64)+`%`, `0.`, ``, -1)
}

const myip = `http://myip.ipip.net/ip`

func GetLocalPublicIP() (string, error) {
	client := new(http.Client)
	resp, err := client.Get(myip)
	if err != nil {
		return ``, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var (
		ip = struct {
			IP string `json:"ip"`
		}{}
	)
	if err := json.Unmarshal(body, &ip); err != nil {
		return ``, err
	}
	return ip.IP, nil
}

func TimeNow() string {
	return time.Now().Format(carbon.DefaultFormat)
}

func Hang() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}

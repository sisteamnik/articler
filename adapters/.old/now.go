package adapters

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

//some code form https://github.com/jinzhu/now
var TimeFormats = []string{
	"1/2/2006",
	"1/2/2006 15:4:5",
	"2006-1-2 15:4:5",
	"2006-1-2 15:4",
	"2006-1-2",
	"02.01.2006",
	"02.01.06 15:04",
	"02.01.2006 15:04",
	"02 01 2006 15:04",
	"1-2",
	"15:4:5",
	"15:4",
	"15:04",
	"03:04",
	"15",
	"15:4:5 Jan 2, 2006 MST",
	"15:04, 02 01 2006",
	time.RFC3339,

	//smoldaily.ru
	"02.01.2006, 15:04",
}

func fixRussianDate(date []byte) []byte {
	alph := "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	if !strings.ContainsAny(string(date), alph) || (len(date) > 0 && string(date[1]) == "<") {
		return date
	}
	in := strings.ToLower(string(date))
	var months = map[string]string{
		"январь":   "01",
		"февраль":  "02",
		"март":     "03",
		"апрель":   "04",
		"май":      "05",
		"июнь":     "06",
		"июль":     "07",
		"август":   "08",
		"сентябрь": "09",
		"октябрь":  "10",
		"ноябрь":   "11",
		"декабрь":  "12",
		"января":   "01",
		"февраля":  "02",
		"марта":    "03",
		"апреля":   "04",
		"мая":      "05",
		"июня":     "06",
		"июля":     "07",
		"августа":  "08",
		"сентября": "09",
		"октября":  "10",
		"ноября":   "11",
		"декабря":  "12",

		"вчера": "yesterday",
	}
	var rep = []string{}
	for k, v := range months {
		rep = append(rep, k, v)
	}
	r := strings.NewReplacer(rep...)
	in = r.Replace(in)
	return []byte(in)
}

type fdur struct {
	Dur time.Duration
	Fmt string
}

var relativeAr = map[string]fdur{
	"yesterday": {-24 * time.Hour, "02.01.2006"},
}

func IsRelativeTime(in string) bool {
	in = string(fixRussianDate([]byte(in)))
	for k := range relativeAr {
		if strings.Contains(in, k) {
			return true
		}
	}
	return false
}

func FixRelativeTime(in string) string {
	in = string(fixRussianDate([]byte(in)))
	if !IsRelativeTime(in) {
		return in
	}
	n := time.Now()
	for k, v := range relativeAr {
		if strings.Contains(in, k) {
			n = n.Add(v.Dur)
			in = strings.Replace(in, k, n.Format(v.Fmt), 1)
		}
	}
	return in
}

type Now struct {
	time.Time
}

func NewNow(t time.Time) *Now {
	return &Now{t}
}

func ParseTime(strs ...string) (time.Time, error) {
	return NewNow(time.Now()).Parse(strs...)
}

func (now *Now) Parse(strs ...string) (t time.Time, err error) {
	var setCurrentTime bool
	parseTime := []int{}
	currentTime := []int{now.Second(), now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year()}
	currentLocation := now.Location()

	for _, str := range strs {
		str = string(fixRussianDate([]byte(str)))
		str = FixRelativeTime(str)
		str = strings.TrimSpace(str)
		str = strings.ToUpper(str)
		onlyTime := regexp.MustCompile(`^\s*\d+(:\d+)*\s*$`).MatchString(str) // match 15:04:05, 15

		t, err = parseWithFormat(str)
		location := t.Location()
		if location.String() == "UTC" {
			location = currentLocation
		}

		if err == nil {
			parseTime = []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
			onlyTime = onlyTime && (parseTime[3] == 1) && (parseTime[4] == 1)

			for i, v := range parseTime {
				// Fill up missed information with current time
				if v == 0 {
					if setCurrentTime {
						//	parseTime[i] = currentTime[i]
					}
				} else {
					//	setCurrentTime = true
				}

				// Default day and month is 1, fill up it if missing it
				if (i == 3 || i == 4) && onlyTime {
					parseTime[i] = currentTime[i]
				}
			}
		}

		if len(parseTime) > 0 {
			t = time.Date(parseTime[5], time.Month(parseTime[4]), parseTime[3], parseTime[2], parseTime[1], parseTime[0], 0, location)
			currentTime = []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
		}
	}
	return
}

func parseWithFormat(str string) (t time.Time, err error) {
	for _, format := range TimeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

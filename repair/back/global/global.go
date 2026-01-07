package global

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"repair/global/config"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/LoperLee/golang-hangul-toolkit/hangul"
	"github.com/google/uuid"
	"golang.org/x/exp/constraints"
)

type GeocodeAddress struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
	}
}

type GeocodeAddressInner struct {
	Code    string `json:"code"`
	Address string `json:"address"`
}

func ToMap(slice []string) map[string]int {
	m := map[string]int{}
	for i, x := range slice {
		m[x] = i
	}
	return m
}

func ReverseMap(inmap map[int]string) map[string]int {
	outmap := make(map[string]int)
	for k, v := range inmap {
		outmap[v] = k
	}
	return outmap
}

func ParseDatetime(str string) *time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err == nil {
		return &t
	}

	return nil
}

func GetTimestamp(str string) int64 {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err == nil {
		return t.Unix()
	}

	return 0
}

func Atoi(value string) int {
	value = strings.ReplaceAll(value, ",", "")
	value = strings.ReplaceAll(value, " ", "")
	i, _ := strconv.Atoi(value)
	return i
}

func Atol(value string) int64 {
	value = strings.ReplaceAll(value, ",", "")
	value = strings.ReplaceAll(value, " ", "")
	i, _ := strconv.ParseInt(value, 10, 64)
	return i
}

func Atof(value string) float64 {
	value = strings.ReplaceAll(value, ",", "")
	value = strings.ReplaceAll(value, " ", "")
	i, _ := strconv.ParseFloat(strings.Replace(value, ",", "", -1), 64)
	return i
}

func Itoa(value int) string {
	return fmt.Sprintf("%v", value)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetCurrentDatetime() string {
	return GetDatetime(time.Now())
}

func GetDatetime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func ArrayToString(A []int, delim string) string {
	var buffer bytes.Buffer
	for i := range A {
		buffer.WriteString(strconv.Itoa(A[i]))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

func StringToIntArray(value string) []int {
	values := strings.Split(value, ",")

	var items []int
	for _, item := range values {
		items = append(items, Atoi(item))
	}

	return items
}

func GetTempFilename() string {
	return filepath.Join("/tmp", uuid.New().String())
}

func Datetime(d string) string {
	if d == "" {
		return ""
	}

	return d
}

func Duration(seconds int) string {
	h := seconds / 60 / 60
	m := seconds / 60 % 60
	s := seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func GetMillisecond(t time.Time) int {
	return t.Nanosecond() / int(time.Millisecond)
}

func GetStringFromDatetime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetStringFromDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func GetDurationFromDate(t time.Time) (string, string) {
	return fmt.Sprintf("%04d-%02d-%02d 00:00:00", t.Year(), t.Month(), t.Day()), fmt.Sprintf("%04d-%02d-%02d 23:59:59", t.Year(), t.Month(), t.Day())
}

func Humandate(d string) string {
	target := ParseDatetime(GetStringFromDatetime(*ParseDatetime(d)))

	t := ParseDatetime(GetStringFromDatetime(time.Now()))
	diff := t.Sub(*target)

	if math.Floor(diff.Hours()/24) > 0 {
		if math.Floor(diff.Hours()/24) > 30 {
			return d[0:4] + "." + d[5:7] + "." + d[8:10]
		} else {
			return fmt.Sprintf("%v일전", math.Floor(diff.Hours()/24))
		}
	}

	if math.Floor(diff.Hours()/24) > 0 {
		return d[0:4] + "." + d[5:7] + "." + d[8:10]
	}

	if math.Floor(diff.Hours()) > 0 {
		return fmt.Sprintf("%v시간전", math.Floor(diff.Hours()))
	}

	m := math.Floor(diff.Minutes())

	if m == 0 {
		return "방금전"
	} else {
		return fmt.Sprintf("%v분전", m)
	}
}

func StripTags(content string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(content, "")
}

func FindImages(htm string) []string {
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

func FindImage(htm string) string {
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(htm, -1)

	if len(imgs) == 0 {
		return ""
	}

	return imgs[0][1]
}

func IsEmptyDate(date string) bool {
	if date == "" || date == "0000-00-00 00:00:00" || date == "1000-01-01 00:00:00" {
		return true
	} else {
		return false
	}
}

func GetSha256(str string) string {
	hash := sha256.New()

	hash.Write([]byte(str))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func SendEmail(email string, title string, content string) error {
	defer func() {
		recover()
	}()

	/*
		d := mail.NewDialer(config.MailHost, config.MailPort, config.MailUser, config.MailPasswd)
		//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		m := mail.NewMessage()
		m.SetHeader("From", config.MailSender)
		m.SetHeader("To", email)
		m.SetHeader("Subject", title)
		m.SetBody("text/html", content)
		err := d.DialAndSend(m)
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	*/
	resp, err := http.PostForm("http://linux.netb.co.kr:9999/mail", url.Values{"from": {config.Mail.Sender}, "to": {email}, "title": {title}, "content": {content}})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)

	return err
}

func SendSMS(tel string, content string) bool {
	defer func() {
		recover()
	}()

	str := fmt.Sprintf("user_id=%v&key=%v&sender=%v&receiver=%v&msg=%v", config.Sms.User, config.Sms.Key, config.Sms.Sender, tel, content)

	rqb := bytes.NewBufferString(str)
	rq, e := http.NewRequest("POST", "https://apis.aligo.in/send/", rqb)
	rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	hc := &http.Client{Timeout: 2 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	rs, e := hc.Do(rq)
	if e != nil {
		return false
	}

	defer rs.Body.Close()

	c, e := io.ReadAll(rs.Body)
	//	fmt.Printf("SMS:\n  body=%+v\n  err=%+v\n  msg=%+v\n", rs.Body, e, str)
	if e != nil {
		log.Println(e)
		return false
	}

	log.Println(string(c))

	return true
}

func WriteFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func ReadFile(filename string) string {
	dat, err := os.ReadFile(filename)

	if err != nil {
		return ""
	}

	return string(dat)
}

func Substr(str string, start int, end int) string {
	b := []byte(str)
	idx := 0
	length := 0
	for range start {
		_, size := utf8.DecodeRune(b[idx:])

		if size == 3 {
			length += 2
		} else {
			length++
		}

		if length >= start {
			break
		}
		idx += size
	}

	pos1 := idx
	idx = 0
	length = 0
	for range end {
		_, size := utf8.DecodeRune(b[idx:])

		if size == 3 {
			length += 2
		} else {
			length++
		}

		if length >= end {
			break
		}
		idx += size
	}

	return str[pos1:idx]
}

func Strlen(s string) int {
	length := len(s)
	r := utf8.RuneCountInString(s)

	return r + (length-r)/2
}

func DiffTime(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func CopyFile(src, dst string) bool {
	in, err := os.Open(src)
	if err != nil {
		return false
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return false
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		panic(err)
	}
	err = out.Sync()
	if err != nil {
		return false
	}

	return true
}

func Unique(s []int) []int {
	keys := make(map[int]struct{})
	res := make([]int, 0)
	for _, val := range s {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}

	sort.Ints(res)

	return res
}

func UniqueString(s []string) []string {
	keys := make(map[string]struct{})
	res := make([]string, 0)
	for _, val := range s {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}

	sort.Strings(res)

	return res
}

func UniqueStringWithoutSort(s []string) []string {
	keys := make(map[string]struct{})
	res := make([]string, 0)
	for _, val := range s {
		if val == "" {
			continue
		}
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}

	return res
}

func UniqueId() string {
	t := time.Now().UTC()
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d_%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixNano())
}

func GetExt(filename string) string {
	ext := strings.ToLower(strings.ReplaceAll(filepath.Ext(filename), ".", ""))
	if ext == "jpeg" {
		ext = "jpg"
	}

	return ext
}

func HumanMoney[T constraints.Integer](value T) string {
	hanA := []string{"", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구", "십"}
	danA := []string{"", "십", "백", "천", "", "십", "백", "천", "", "십", "백", "천", "", "십", "백", "천"}
	result := ""
	num := fmt.Sprintf("%v", value)

	for i := 0; i < len(num); i++ {

		pos := Atoi(num[len(num)-i-1 : len(num)-i])
		str := ""

		han := hanA[pos]
		if pos != 0 {
			str += han + danA[i]
		}
		if i == 4 {
			str += "만"
		}

		if i == 8 {
			str += "억"
		}
		if i == 12 {
			str += "조"
		}

		result = str + result
	}

	return result
}

func GetFile(url string, filename string) {
	file, err := os.Create(filename)

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	defer file.Close()
}

func GetJosa(str string, josa hangul.Josa) string {
	var re = regexp.MustCompile(`\([^)]+\)`)
	strip := re.ReplaceAllString(str, "")
	convert := hangul.GetJosa(strip, josa)

	remain := strings.ReplaceAll(convert, strip, "")
	return str + remain
}

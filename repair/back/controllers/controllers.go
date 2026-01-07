package controllers

import (
	"path"
	"repair/global"
	"repair/global/config"
	"repair/models"

	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/CloudyKit/jet/v3"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	humanize "github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type Controller struct {
	Context    *gin.Context
	Vars       jet.VarMap
	Result     gin.H
	Connection *models.Connection
	Current    string
	Code       int
	Session    *models.User

	Date string

	Page     int
	Pagesize int
}

func NewController(g *gin.Context) *Controller {
	var ctl Controller
	ctl.Init(g)
	return &ctl
}

func (c *Controller) Init(g *gin.Context) {
	c.Context = g
	c.Vars = make(jet.VarMap)
	c.Result = make(gin.H)
	c.Result["code"] = "ok"
	c.Connection = c.NewConnection()
	c.Code = http.StatusOK

	t := time.Now()
	c.Date = global.GetDate(t)

	if user, ok := g.Get("user"); ok {
		c.Session = user.(*models.User)
		c.Session.Passwd = ""
	} else {
		c.Session = nil
	}

	c.Set("_t", t.UnixNano())
}

func (c *Controller) Set(name string, value interface{}) {
	c.Result[name] = value
}

func (c *Controller) SetArray(value gin.H) {
	for k, v := range value {
		c.Result[k] = v
	}
}

func (c *Controller) JsonDisplay() string {
	str, err := json.Marshal(c.Result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str[:])
}

func (c *Controller) Display(filename string) {
	for k, v := range c.Result {
		c.Vars.Set(k, v)
	}

	var view = jet.NewHTMLSet("./views")
	view.SetDevelopmentMode(true)

	view.AddGlobal("isMobile", func() bool {
		return false
	})

	view.AddGlobal("get", func(str string) string {
		return c.Get(str)
	})

	view.AddGlobal("geti", func(str string) int {
		return c.Geti(str)
	})

	view.AddGlobal("getf", func(str string) float64 {
		return c.Getf(str)
	})

	view.AddGlobal("split", func(str string, sep string) []string {
		return strings.Split(str, sep)
	})

	view.AddGlobal("ismobile", func() bool {
		return strings.Contains(c.Context.GetHeader("User-Agent"), "Mobi")
	})

	view.AddGlobal("urlencode", func(str string) string {
		return url.QueryEscape(str)
	})

	view.AddGlobal("encode", func(str string) string {
		str = strings.Replace(str, "&#039;", "'", -1)
		str = strings.Replace(str, "&amp;", "&", -1)

		return str
	})

	view.AddGlobal("nl2br", func(str string) string {
		return strings.Replace(str, "\n", "<br />", -1)
	})

	view.AddGlobal("trim", func(str string) string {
		return strings.TrimSpace(str)
	})

	view.AddGlobal("cleartag", func(str string) string {
		return strings.Replace(str, "#", "", -1)
	})

	view.AddGlobal("striptags", func(str string) string {
		return global.StripTags(str)
	})

	view.AddGlobal("datetime", func(d string) string {
		return global.Datetime(d)
	})

	view.AddGlobal("humandate", func(d string) string {
		return global.Humandate(d)
	})

	view.AddGlobal("str", func(ii int64) string {
		return fmt.Sprintf("%v", ii)
	})

	view.AddGlobal("int", func(str string) int {
		if str == "" {
			return 0
		}

		val, _ := strconv.Atoi(str)
		return val
	})

	view.AddGlobal("comma", func(d int) string {
		return humanize.Comma(int64(d))
	})

	t, err := view.GetTemplate(filename)
	if err == nil {
		if err = t.Execute(c.Context.Writer, c.Vars, nil); err != nil {
			log.Println(err)
			// error when executing template
		}
	}
}

func (c *Controller) GetArray(name string) []string {
	if c.Context.Request.Method == "GET" {
		return c.Context.QueryArray(name)
	} else {
		return c.Context.PostFormArray(name)
	}
}

func (c *Controller) GetArrayComma(name string) []string {
	value := c.Get(name)

	return strings.Split(value, ",")
}

func (c *Controller) GetArrayCommai(name string) []int {
	value := c.Get(name)

	values := strings.Split(value, ",")

	var items []int
	for _, item := range values {
		items = append(items, global.Atoi(item))
	}

	return items
}

func (c *Controller) Get(name string) string {
	if c.Context.Request.Method == "GET" {
		return c.Query(name)
	} else {
		return c.Post(name)
	}
}

func (c *Controller) GetStartdate(name string) string {
	date := c.Get(name)

	if date != "" {
		date += ":00"
	}

	return date
}

func (c *Controller) GetEnddate(name string) string {
	date := c.Get(name)

	if date != "" {
		date += ":59"
	}

	return date
}

func (c *Controller) Geti(name string) int {
	if c.Context.Request.Method == "GET" {
		return c.Queryi(name)
	} else {
		return c.Posti(name)
	}
}

func (c *Controller) Geti64(name string) int64 {
	if c.Context.Request.Method == "GET" {
		return c.Queryi64(name)
	} else {
		return c.Posti64(name)
	}
}

func (c *Controller) Getf(name string) float64 {
	if c.Context.Request.Method == "GET" {
		return c.Queryf(name)
	} else {
		return c.Postf(name)
	}
}

func (c *Controller) Geti64Array(name string) []int64 {
	str := c.Get(name)

	ret := make([]int64, 0)

	if str == "" {
		return ret
	}

	items := strings.Split(str, ",")

	for _, v := range items {
		ret = append(ret, global.Atol(strings.TrimSpace(v)))
	}

	return ret
}

func (c *Controller) DefaultGet(name string, defaultValue string) string {
	if c.Context.Request.Method == "GET" {
		return c.DefaultQuery(name, defaultValue)
	} else {
		return c.DefaultPost(name, defaultValue)
	}
}

func (c *Controller) DefaultGeti(name string, defaultValue int) int {
	if c.Context.Request.Method == "GET" {
		return c.DefaultQueryi(name, defaultValue)
	} else {
		return c.DefaultPosti(name, defaultValue)
	}
}

func (c *Controller) DefaultGeti64(name string, defaultValue int64) int64 {
	if c.Context.Request.Method == "GET" {
		return c.DefaultQueryi64(name, defaultValue)
	} else {
		return c.DefaultPosti64(name, defaultValue)
	}
}

func (c *Controller) Post(name string) string {
	return c.Context.PostForm(name)
}

func (c *Controller) Posti(name string) int {
	value, _ := strconv.Atoi(c.Context.PostForm(name))
	return value
}

func (c *Controller) Posti64(name string) int64 {
	value, _ := strconv.ParseInt(c.Context.PostForm(name), 10, 64)
	return value
}

func (c *Controller) Postf(name string) float64 {
	value, _ := strconv.ParseFloat(c.Context.PostForm(name), 64)
	return value
}

func (c *Controller) DefaultPost(name string, defaultValue string) string {
	value := c.Context.PostForm(name)

	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultPosti(name string, defaultValue int) int {
	value, _ := strconv.Atoi(c.Context.PostForm(name))

	if value == 0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultPosti64(name string, defaultValue int64) int64 {
	value, _ := strconv.ParseInt(c.Context.PostForm(name), 10, 64)

	if value == 0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) Query(name string) string {
	return c.Context.Query(name)
}

func (c *Controller) Queryi(name string) int {
	value, _ := strconv.Atoi(c.Context.Query(name))
	return value
}

func (c *Controller) Queryi64(name string) int64 {
	value, _ := strconv.ParseInt(c.Context.Query(name), 10, 64)
	return value
}

func (c *Controller) Queryf(name string) float64 {
	value, _ := strconv.ParseFloat(c.Context.Query(name), 64)
	return value
}

func (c *Controller) DefaultQuery(name string, defaultValue string) string {
	value := c.Context.Query(name)

	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultQueryi(name string, defaultValue int) int {
	value, _ := strconv.Atoi(c.Context.Query(name))

	if value == 0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultQueryi64(name string, defaultValue int64) int64 {
	value, _ := strconv.ParseInt(c.Context.Query(name), 10, 64)

	if value == 0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultQueryf(name string, defaultValue float64) float64 {
	value, _ := strconv.ParseFloat(c.Context.Query(name), 64)

	if value == 0.0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) Refresh(url string) {
	str := "<script>location.href = '" + url + "';</script>"
	c.Context.Writer.WriteHeader(http.StatusOK)
	c.Context.Writer.Write([]byte(str))
}

func (c *Controller) Download(filename string, downloadFilename string) {
	filesize, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
	}
	log.Println("filesize", filesize.Size())
	c.Context.Header("Content-Type", "application/octet-stream")
	c.Context.Header("Content-Length", fmt.Sprintf("%v", filesize.Size()))
	c.Context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\";", filepath.Base(downloadFilename)))
	c.Context.Header("Content-Transfer-Encoding", "binary")
	c.Context.Header("Pragma", "no-cache")
	c.Context.Header("Expires", "0")

	c.Context.File(filename)
}

func (c *Controller) NewConnection() *models.Connection {
	if c.Connection != nil {
		return c.Connection
	}

	c.Connection = models.NewConnection()
	return c.Connection
}

func (c *Controller) Close() {
	if c.Connection != nil {
		c.Connection.Close()
		c.Connection = nil
	}
}

func (c *Controller) Bind(obj interface{}) error {
	return c.Context.ShouldBind(obj)
}

func (c *Controller) Paging(page int, totalRows int, pageSize int) {
	blockSize := 5

	totalPage := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	totalBlock := int(math.Ceil(float64(totalPage) / float64(blockSize)))
	currentBlock := int(math.Ceil(float64(page) / float64(blockSize)))

	startPage := (currentBlock-1)*blockSize + 1
	endPage := currentBlock * blockSize
	if endPage > totalPage {
		endPage = totalPage
	}

	s := make([]int, endPage-startPage+1)
	for i := range s {
		s[i] = startPage + i
	}

	c.Set("pages", s)
	c.Set("page", page)
	c.Set("blockSize", blockSize)
	c.Set("totalPage", totalPage)
	c.Set("totalBlock", totalBlock)
	c.Set("currentBlock", currentBlock)
}

func (c *Controller) GetUpload(uploadPath string, name string) (string, string) {
	log.Println("controller getupload")
	file, err := c.Context.FormFile(name)

	if err != nil {
		log.Println(err)
		return "", ""
	}

	t := time.Now()
	u2 := uuid.NewV4()

	filename := fmt.Sprintf("%v/%04d%02d%02d%02d%02d%02d_%v%v", uploadPath, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), strings.Replace(u2.String(), "-", "", -1), filepath.Ext(file.Filename))
	fullFilename := path.Join(config.UploadPath, filename)

	log.Println("path : ", config.UploadPath, fullFilename)

	c.Context.SaveUploadedFile(file, fullFilename)

	return file.Filename, filename
}

func (c *Controller) MultiUpload(uploadPath string, name string) []string {
	log.Println("controller multi upload")
	rets := make([]string, 0)

	form, _ := c.Context.MultipartForm()
	files := form.File[fmt.Sprintf("%v[]", name)]

	t := time.Now()

	for _, file := range files {
		log.Println(file.Filename)
		u2 := uuid.NewV4()

		filename := fmt.Sprintf("%v/%04d%02d%02d%02d%02d%02d_%v%v", uploadPath, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), strings.Replace(u2.String(), "-", "", -1), filepath.Ext(file.Filename))
		fullFilename := path.Join(config.UploadPath, filename)

		c.Context.SaveUploadedFile(file, fullFilename)
		log.Println(filename)
		rets = append(rets, filename)
	}

	return rets
}

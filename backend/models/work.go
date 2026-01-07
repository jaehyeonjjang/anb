package models

import (
    "anb/config"
    "database/sql"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"

    //_ "github.com/denisenkom/go-mssqldb"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
)

type Work struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Image                    int64 `json:"Image, int64" form:"image"`
    Type                     int `json:"Type, int" form:"type"`
    Content                  string `json:"Content, string" form:"content"`
    Pos                      int `json:"Pos, int" form:"pos"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type WorkManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewWorkManager(conn *sql.DB) *WorkManager {
    var item WorkManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "w"
    item.Index = ""

    return &item
}

func (p *WorkManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *WorkManager) GetLast(items *[]Work) *Work {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *WorkManager) SetIndex(index string) {
    p.Index = index
}

func (p *WorkManager) GetQuery() string {
    ret := ""

    tableName := "work_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".work_tb"
    }

    str := "select w_id, w_image, w_type, w_content, w_pos, w_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *WorkManager) GetQuerySelect() string {
    ret := ""

    tableName := "work_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".work_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *WorkManager) Insert(item *Work) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "work_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".work_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (w_id, w_image, w_type, w_content, w_pos, w_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Image, item.Type, item.Content, item.Pos, item.Date)
    } else {
        query = "insert into " + tableName + " (w_image, w_type, w_content, w_pos, w_date) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Image, item.Type, item.Content, item.Pos, item.Date)
    }

    if err == nil {
        p.Result = &res
    } else {
        log.Println(item)
        log.Println(err)
        p.Result = nil
    }

    return err
}
func (p *WorkManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "work_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".work_tb"
    }
    query := "delete from " + tableName + " where w_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *WorkManager) Update(item *Work) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "work_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".work_tb"
    }

	query := "update " + tableName + " set w_image = ?,w_type = ?,w_content = ?,w_pos = ?,w_date = ? where w_id = ?"
	_, err := p.Conn.Exec(query, item.Image, item.Type, item.Content, item.Pos, item.Date, item.Id)

    return err
}

func (p *WorkManager) GetIdentity() int64 {
    if p.Result == nil {
        return 0
    }

    id, err := (*p.Result).LastInsertId()

    if err != nil {
        return 0
    } else {
        return id
    }
}
    

func (p *WorkManager) ReadRow(rows *sql.Rows) *Work {
    var item Work
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Image, &item.Type, &item.Content, &item.Pos, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *WorkManager) ReadRows(rows *sql.Rows) *[]Work {
    var items []Work
    var err error

    for rows.Next() {
        var item Work
        err = rows.Scan(&item.Id, &item.Image, &item.Type, &item.Content, &item.Pos, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *WorkManager) Get(id int64) *Work {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where w_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *WorkManager) GetList(page int, pagesize int, order string) *[]Work {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "w_id desc"
        } else {
            order = "w_" + order
        }
        query += " order by " + order
        if config.Database == "mysql" {
            query += " limit ? offset ?"
            rows, err = p.Conn.Query(query, pagesize, startpage)
        } else if config.Database == "mssql" || config.Database == "sqlserver" {
            query += "OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
            rows, err = p.Conn.Query(query, startpage, pagesize)
        }
    } else {
        if order == "" {
            order = "w_id"
        } else {
            order = "w_" + order
        }
        query += " order by " + order
        rows, err = p.Conn.Query(query)
    }

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *WorkManager) GetCount() int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    rows, err := p.Conn.Query(query)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return 0
    }

    defer rows.Close()

    if !rows.Next() {
        return 0
    }

    cnt := 0
    err = rows.Scan(&cnt)

    if err != nil {
        return 0
    } else {
        return cnt
    }
}

func (p *WorkManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Work {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where w_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "w_id desc"
        } else {
            order = "w_" + order
        }
        query += " order by " + order
        if config.Database == "mysql" {
            query += " limit ? offset ?"
            rows, err = p.Conn.Query(query, pagesize, startpage)
        } else if config.Database == "mssql" || config.Database == "sqlserver" {
            query += "OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
            rows, err = p.Conn.Query(query, startpage, pagesize)
        }
    } else {
        if order == "" {
            order = "w_id"
        } else {
            order = "w_" + order
        }
        query += " order by " + order
        rows, err = p.Conn.Query(query)
    }

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *WorkManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where w_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    rows, err := p.Conn.Query(query)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return 0
    }

    defer rows.Close()

    if !rows.Next() {
        return 0
    }

    cnt := 0
    err = rows.Scan(&cnt)

    if err != nil {
        return 0
    } else {
        return cnt
    }
}

func (p *WorkManager) GetListByImage(image int64, page int, pagesize int, orderby string) *[]Work {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if image != 0 {
		query += " and w_image = ?"
		params = append(params, image)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "w_id desc"
        } else {
            orderby = "w_" + orderby
        }
        query += " order by " + orderby
        if config.Database == "mysql" {
            query += " limit ? offset ?"
            params = append(params, pagesize)
            params = append(params, startpage)
        } else if config.Database == "mssql" || config.Database == "sqlserver" {
            query += "OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
            params = append(params, startpage)
            params = append(params, pagesize)
        }
    } else {
        if orderby == "" {
            orderby = "w_id"
        } else {
            orderby = "w_" + orderby
        }
        query += " order by " + orderby
    }

    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRows(rows)
}

func (p *WorkManager) GetCountByImage(image int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if image != 0 {
		query += " and w_image = ?"
		params = append(params, image)
	}

    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return 0
    }

    defer rows.Close()

    if !rows.Next() {
        return 0
    }

    cnt := 0
    err = rows.Scan(&cnt)

    if err != nil {
        return 0
    } else {
        return cnt
    }
}

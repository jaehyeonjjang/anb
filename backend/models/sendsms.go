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

type Sendsms struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Level                    int `json:"Level, int" form:"level"`
    Content                  string `json:"Content, string" form:"content"`
    Status                   int `json:"Status, int" form:"status"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type SendsmsManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewSendsmsManager(conn *sql.DB) *SendsmsManager {
    var item SendsmsManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "ss"
    item.Index = ""

    return &item
}

func (p *SendsmsManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *SendsmsManager) GetLast(items *[]Sendsms) *Sendsms {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *SendsmsManager) SetIndex(index string) {
    p.Index = index
}

func (p *SendsmsManager) GetQuery() string {
    ret := ""

    tableName := "sendsms_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendsms_tb"
    }

    str := "select ss_id, ss_level, ss_content, ss_status, ss_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *SendsmsManager) GetQuerySelect() string {
    ret := ""

    tableName := "sendsms_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendsms_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *SendsmsManager) Insert(item *Sendsms) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "sendsms_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendsms_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (ss_id, ss_level, ss_content, ss_status, ss_date) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Level, item.Content, item.Status, item.Date)
    } else {
        query = "insert into " + tableName + " (ss_level, ss_content, ss_status, ss_date) values (?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Level, item.Content, item.Status, item.Date)
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
func (p *SendsmsManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "sendsms_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendsms_tb"
    }
    query := "delete from " + tableName + " where ss_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *SendsmsManager) Update(item *Sendsms) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "sendsms_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendsms_tb"
    }

	query := "update " + tableName + " set ss_level = ?,ss_content = ?,ss_status = ?,ss_date = ? where ss_id = ?"
	_, err := p.Conn.Exec(query, item.Level, item.Content, item.Status, item.Date, item.Id)

    return err
}

func (p *SendsmsManager) GetIdentity() int64 {
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
    

func (p *SendsmsManager) ReadRow(rows *sql.Rows) *Sendsms {
    var item Sendsms
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Level, &item.Content, &item.Status, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *SendsmsManager) ReadRows(rows *sql.Rows) *[]Sendsms {
    var items []Sendsms
    var err error

    for rows.Next() {
        var item Sendsms
        err = rows.Scan(&item.Id, &item.Level, &item.Content, &item.Status, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *SendsmsManager) Get(id int64) *Sendsms {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where ss_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *SendsmsManager) GetList(page int, pagesize int, order string) *[]Sendsms {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "ss_id desc"
        } else {
            order = "ss_" + order
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
            order = "ss_id"
        } else {
            order = "ss_" + order
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


func (p *SendsmsManager) GetCount() int {
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

func (p *SendsmsManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Sendsms {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where ss_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "ss_id desc"
        } else {
            order = "ss_" + order
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
            order = "ss_id"
        } else {
            order = "ss_" + order
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


func (p *SendsmsManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where ss_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *SendsmsManager) GetListByStatus(status int, page int, pagesize int, orderby string) *[]Sendsms {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if status != 0 {
		query += " and ss_status = ?"
		params = append(params, status)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "ss_id desc"
        } else {
            orderby = "ss_" + orderby
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
            orderby = "ss_id"
        } else {
            orderby = "ss_" + orderby
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

func (p *SendsmsManager) GetCountByStatus(status int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if status != 0 {
		query += " and ss_status = ?"
		params = append(params, status)
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

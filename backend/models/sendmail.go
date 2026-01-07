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

type Sendmail struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Level                    int `json:"Level, int" form:"level"`
    Title                    string `json:"Title, string" form:"title"`
    Content                  string `json:"Content, string" form:"content"`
    Status                   int `json:"Status, int" form:"status"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type SendmailManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewSendmailManager(conn *sql.DB) *SendmailManager {
    var item SendmailManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "sm"
    item.Index = ""

    return &item
}

func (p *SendmailManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *SendmailManager) GetLast(items *[]Sendmail) *Sendmail {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *SendmailManager) SetIndex(index string) {
    p.Index = index
}

func (p *SendmailManager) GetQuery() string {
    ret := ""

    tableName := "sendmail_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendmail_tb"
    }

    str := "select sm_id, sm_level, sm_title, sm_content, sm_status, sm_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *SendmailManager) GetQuerySelect() string {
    ret := ""

    tableName := "sendmail_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendmail_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *SendmailManager) Insert(item *Sendmail) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "sendmail_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendmail_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (sm_id, sm_level, sm_title, sm_content, sm_status, sm_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Level, item.Title, item.Content, item.Status, item.Date)
    } else {
        query = "insert into " + tableName + " (sm_level, sm_title, sm_content, sm_status, sm_date) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Level, item.Title, item.Content, item.Status, item.Date)
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
func (p *SendmailManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "sendmail_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendmail_tb"
    }
    query := "delete from " + tableName + " where sm_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *SendmailManager) Update(item *Sendmail) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "sendmail_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".sendmail_tb"
    }

	query := "update " + tableName + " set sm_level = ?,sm_title = ?,sm_content = ?,sm_status = ?,sm_date = ? where sm_id = ?"
	_, err := p.Conn.Exec(query, item.Level, item.Title, item.Content, item.Status, item.Date, item.Id)

    return err
}

func (p *SendmailManager) GetIdentity() int64 {
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
    

func (p *SendmailManager) ReadRow(rows *sql.Rows) *Sendmail {
    var item Sendmail
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Level, &item.Title, &item.Content, &item.Status, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *SendmailManager) ReadRows(rows *sql.Rows) *[]Sendmail {
    var items []Sendmail
    var err error

    for rows.Next() {
        var item Sendmail
        err = rows.Scan(&item.Id, &item.Level, &item.Title, &item.Content, &item.Status, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *SendmailManager) Get(id int64) *Sendmail {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where sm_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *SendmailManager) GetList(page int, pagesize int, order string) *[]Sendmail {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "sm_id desc"
        } else {
            order = "sm_" + order
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
            order = "sm_id"
        } else {
            order = "sm_" + order
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


func (p *SendmailManager) GetCount() int {
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

func (p *SendmailManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Sendmail {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where sm_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "sm_id desc"
        } else {
            order = "sm_" + order
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
            order = "sm_id"
        } else {
            order = "sm_" + order
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


func (p *SendmailManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where sm_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *SendmailManager) GetListByStatus(status int, page int, pagesize int, orderby string) *[]Sendmail {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if status != 0 {
		query += " and sm_status = ?"
		params = append(params, status)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "sm_id desc"
        } else {
            orderby = "sm_" + orderby
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
            orderby = "sm_id"
        } else {
            orderby = "sm_" + orderby
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

func (p *SendmailManager) GetCountByStatus(status int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if status != 0 {
		query += " and sm_status = ?"
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

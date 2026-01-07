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

type Findpasswd struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    User                     int64 `json:"User, int64" form:"user"`
    Link                     string `json:"Link, string" form:"link"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type FindpasswdManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewFindpasswdManager(conn *sql.DB) *FindpasswdManager {
    var item FindpasswdManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "fp"
    item.Index = ""

    return &item
}

func (p *FindpasswdManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *FindpasswdManager) GetLast(items *[]Findpasswd) *Findpasswd {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *FindpasswdManager) SetIndex(index string) {
    p.Index = index
}

func (p *FindpasswdManager) GetQuery() string {
    ret := ""

    tableName := "findpasswd_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".findpasswd_tb"
    }

    str := "select fp_id, fp_user, fp_link, fp_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *FindpasswdManager) GetQuerySelect() string {
    ret := ""

    tableName := "findpasswd_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".findpasswd_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *FindpasswdManager) Insert(item *Findpasswd) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "findpasswd_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".findpasswd_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (fp_id, fp_user, fp_link, fp_date) values (?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.User, item.Link, item.Date)
    } else {
        query = "insert into " + tableName + " (fp_user, fp_link, fp_date) values (?, ?, ?)"
        res, err = p.Conn.Exec(query, item.User, item.Link, item.Date)
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
func (p *FindpasswdManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "findpasswd_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".findpasswd_tb"
    }
    query := "delete from " + tableName + " where fp_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *FindpasswdManager) Update(item *Findpasswd) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "findpasswd_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".findpasswd_tb"
    }

	query := "update " + tableName + " set fp_user = ?,fp_link = ?,fp_date = ? where fp_id = ?"
	_, err := p.Conn.Exec(query, item.User, item.Link, item.Date, item.Id)

    return err
}

func (p *FindpasswdManager) GetIdentity() int64 {
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
    

func (p *FindpasswdManager) ReadRow(rows *sql.Rows) *Findpasswd {
    var item Findpasswd
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.User, &item.Link, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *FindpasswdManager) ReadRows(rows *sql.Rows) *[]Findpasswd {
    var items []Findpasswd
    var err error

    for rows.Next() {
        var item Findpasswd
        err = rows.Scan(&item.Id, &item.User, &item.Link, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *FindpasswdManager) Get(id int64) *Findpasswd {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where fp_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *FindpasswdManager) GetList(page int, pagesize int, order string) *[]Findpasswd {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "fp_id desc"
        } else {
            order = "fp_" + order
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
            order = "fp_id"
        } else {
            order = "fp_" + order
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


func (p *FindpasswdManager) GetCount() int {
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

func (p *FindpasswdManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Findpasswd {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where fp_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "fp_id desc"
        } else {
            order = "fp_" + order
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
            order = "fp_id"
        } else {
            order = "fp_" + order
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


func (p *FindpasswdManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where fp_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *FindpasswdManager) GetByLink(link string) *Findpasswd {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if link != "" {
		query += " and fp_link = ?"
		params = append(params, link)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

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

type Aptsubmaster struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Apt                      int64 `json:"Apt, int64" form:"apt"`
    User                     int64 `json:"User, int64" form:"user"`
    Level                    int `json:"Level, int" form:"level"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type AptsubmasterManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewAptsubmasterManager(conn *sql.DB) *AptsubmasterManager {
    var item AptsubmasterManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "as"
    item.Index = ""

    return &item
}

func (p *AptsubmasterManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptsubmasterManager) GetLast(items *[]Aptsubmaster) *Aptsubmaster {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *AptsubmasterManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptsubmasterManager) GetQuery() string {
    ret := ""

    tableName := "aptsubmaster_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptsubmaster_tb"
    }

    str := "select as_id, as_apt, as_user, as_level, as_company, as_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptsubmasterManager) GetQuerySelect() string {
    ret := ""

    tableName := "aptsubmaster_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptsubmaster_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptsubmasterManager) Insert(item *Aptsubmaster) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "aptsubmaster_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptsubmaster_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (as_id, as_apt, as_user, as_level, as_company, as_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Apt, item.User, item.Level, item.Company, item.Date)
    } else {
        query = "insert into " + tableName + " (as_apt, as_user, as_level, as_company, as_date) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Apt, item.User, item.Level, item.Company, item.Date)
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
func (p *AptsubmasterManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptsubmaster_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptsubmaster_tb"
    }
    query := "delete from " + tableName + " where as_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *AptsubmasterManager) Update(item *Aptsubmaster) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptsubmaster_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptsubmaster_tb"
    }

	query := "update " + tableName + " set as_apt = ?,as_user = ?,as_level = ?,as_company = ?,as_date = ? where as_id = ?"
	_, err := p.Conn.Exec(query, item.Apt, item.User, item.Level, item.Company, item.Date, item.Id)

    return err
}

func (p *AptsubmasterManager) GetIdentity() int64 {
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
    

func (p *AptsubmasterManager) ReadRow(rows *sql.Rows) *Aptsubmaster {
    var item Aptsubmaster
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Apt, &item.User, &item.Level, &item.Company, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *AptsubmasterManager) ReadRows(rows *sql.Rows) *[]Aptsubmaster {
    var items []Aptsubmaster
    var err error

    for rows.Next() {
        var item Aptsubmaster
        err = rows.Scan(&item.Id, &item.Apt, &item.User, &item.Level, &item.Company, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *AptsubmasterManager) Get(id int64) *Aptsubmaster {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where as_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *AptsubmasterManager) GetList(page int, pagesize int, order string) *[]Aptsubmaster {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "as_id desc"
        } else {
            order = "as_" + order
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
            order = "as_id"
        } else {
            order = "as_" + order
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


func (p *AptsubmasterManager) GetCount() int {
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

func (p *AptsubmasterManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Aptsubmaster {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where as_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "as_id desc"
        } else {
            order = "as_" + order
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
            order = "as_id"
        } else {
            order = "as_" + order
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


func (p *AptsubmasterManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where as_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

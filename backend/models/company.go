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

type Company struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Name                     string `json:"Name, string" form:"name"`
    Ceo                      string `json:"Ceo, string" form:"ceo"`
    Logo                     string `json:"Logo, string" form:"logo"`
    Stamp                    string `json:"Stamp, string" form:"stamp"`
    Contractstartdate        string `json:"Contractstartdate, string" form:"contractstartdate"`
    Contractenddate          string `json:"Contractenddate, string" form:"contractenddate"`
    Status                   int `json:"Status, int" form:"status"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type CompanyManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewCompanyManager(conn *sql.DB) *CompanyManager {
    var item CompanyManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "c"
    item.Index = ""

    return &item
}

func (p *CompanyManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CompanyManager) GetLast(items *[]Company) *Company {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *CompanyManager) SetIndex(index string) {
    p.Index = index
}

func (p *CompanyManager) GetQuery() string {
    ret := ""

    tableName := "company_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".company_tb"
    }

    str := "select c_id, c_name, c_ceo, c_logo, c_stamp, c_contractstartdate, c_contractenddate, c_status, c_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *CompanyManager) GetQuerySelect() string {
    ret := ""

    tableName := "company_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".company_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *CompanyManager) Insert(item *Company) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Contractstartdate == "" {
        item.Contractstartdate = "1000-01-01 00:00:00"
    }

    if item.Contractenddate == "" {
        item.Contractenddate = "1000-01-01 00:00:00"
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "company_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".company_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (c_id, c_name, c_ceo, c_logo, c_stamp, c_contractstartdate, c_contractenddate, c_status, c_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Name, item.Ceo, item.Logo, item.Stamp, item.Contractstartdate, item.Contractenddate, item.Status, item.Date)
    } else {
        query = "insert into " + tableName + " (c_name, c_ceo, c_logo, c_stamp, c_contractstartdate, c_contractenddate, c_status, c_date) values (?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Name, item.Ceo, item.Logo, item.Stamp, item.Contractstartdate, item.Contractenddate, item.Status, item.Date)
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
func (p *CompanyManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "company_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".company_tb"
    }
    query := "delete from " + tableName + " where c_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *CompanyManager) Update(item *Company) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Contractstartdate == "" {
        item.Contractstartdate = "1000-01-01 00:00:00"
    }

    if item.Contractenddate == "" {
        item.Contractenddate = "1000-01-01 00:00:00"
    }

    tableName := "company_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".company_tb"
    }

	query := "update " + tableName + " set c_name = ?,c_ceo = ?,c_logo = ?,c_stamp = ?,c_contractstartdate = ?,c_contractenddate = ?,c_status = ?,c_date = ? where c_id = ?"
	_, err := p.Conn.Exec(query, item.Name, item.Ceo, item.Logo, item.Stamp, item.Contractstartdate, item.Contractenddate, item.Status, item.Date, item.Id)

    return err
}

func (p *CompanyManager) GetIdentity() int64 {
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
    

func (p *CompanyManager) ReadRow(rows *sql.Rows) *Company {
    var item Company
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Ceo, &item.Logo, &item.Stamp, &item.Contractstartdate, &item.Contractenddate, &item.Status, &item.Date)

        if item.Contractstartdate == "1000-01-01 00:00:00" {
            item.Contractstartdate = ""
        }

        if item.Contractenddate == "1000-01-01 00:00:00" {
            item.Contractenddate = ""
        }
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *CompanyManager) ReadRows(rows *sql.Rows) *[]Company {
    var items []Company
    var err error

    for rows.Next() {
        var item Company
        err = rows.Scan(&item.Id, &item.Name, &item.Ceo, &item.Logo, &item.Stamp, &item.Contractstartdate, &item.Contractenddate, &item.Status, &item.Date)

        if item.Contractstartdate == "1000-01-01 00:00:00" {
            item.Contractstartdate = ""
        }

        if item.Contractenddate == "1000-01-01 00:00:00" {
            item.Contractenddate = ""
        }

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *CompanyManager) Get(id int64) *Company {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where c_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CompanyManager) GetList(page int, pagesize int, order string) *[]Company {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "c_id desc"
        } else {
            order = "c_" + order
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
            order = "c_id"
        } else {
            order = "c_" + order
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


func (p *CompanyManager) GetCount() int {
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

func (p *CompanyManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Company {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where c_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "c_id desc"
        } else {
            order = "c_" + order
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
            order = "c_id"
        } else {
            order = "c_" + order
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


func (p *CompanyManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where c_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

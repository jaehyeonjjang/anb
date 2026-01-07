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

type Companyuser struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Company                  int64 `json:"Company, int64" form:"company"`
    User                     int64 `json:"User, int64" form:"user"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type CompanyuserManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewCompanyuserManager(conn *sql.DB) *CompanyuserManager {
    var item CompanyuserManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "cu"
    item.Index = ""

    return &item
}

func (p *CompanyuserManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CompanyuserManager) GetLast(items *[]Companyuser) *Companyuser {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *CompanyuserManager) SetIndex(index string) {
    p.Index = index
}

func (p *CompanyuserManager) GetQuery() string {
    ret := ""

    tableName := "companyuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".companyuser_tb"
    }

    str := "select cu_id, cu_company, cu_user, cu_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *CompanyuserManager) GetQuerySelect() string {
    ret := ""

    tableName := "companyuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".companyuser_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *CompanyuserManager) Insert(item *Companyuser) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "companyuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".companyuser_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (cu_id, cu_company, cu_user, cu_date) values (?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Company, item.User, item.Date)
    } else {
        query = "insert into " + tableName + " (cu_company, cu_user, cu_date) values (?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Company, item.User, item.Date)
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
func (p *CompanyuserManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "companyuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".companyuser_tb"
    }
    query := "delete from " + tableName + " where cu_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *CompanyuserManager) Update(item *Companyuser) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "companyuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".companyuser_tb"
    }

	query := "update " + tableName + " set cu_company = ?,cu_user = ?,cu_date = ? where cu_id = ?"
	_, err := p.Conn.Exec(query, item.Company, item.User, item.Date, item.Id)

    return err
}

func (p *CompanyuserManager) GetIdentity() int64 {
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
    

func (p *CompanyuserManager) ReadRow(rows *sql.Rows) *Companyuser {
    var item Companyuser
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Company, &item.User, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *CompanyuserManager) ReadRows(rows *sql.Rows) *[]Companyuser {
    var items []Companyuser
    var err error

    for rows.Next() {
        var item Companyuser
        err = rows.Scan(&item.Id, &item.Company, &item.User, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *CompanyuserManager) Get(id int64) *Companyuser {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where cu_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CompanyuserManager) GetList(page int, pagesize int, order string) *[]Companyuser {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "cu_id desc"
        } else {
            order = "cu_" + order
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
            order = "cu_id"
        } else {
            order = "cu_" + order
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


func (p *CompanyuserManager) GetCount() int {
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

func (p *CompanyuserManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Companyuser {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where cu_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "cu_id desc"
        } else {
            order = "cu_" + order
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
            order = "cu_id"
        } else {
            order = "cu_" + order
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


func (p *CompanyuserManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where cu_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

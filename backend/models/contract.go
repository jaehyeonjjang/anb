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

type Contract struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Contractenddate          string `json:"Contractenddate, string" form:"contractenddate"`
    Contractstartdate        string `json:"Contractstartdate, string" form:"contractstartdate"`
    Date                     string `json:"Date, string" form:"date"`
    Status                   int `json:"Status, int" form:"status"`
    Extra                    interface{} `form:"extra"`
}

type ContractManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewContractManager(conn *sql.DB) *ContractManager {
    var item ContractManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "co"
    item.Index = ""

    return &item
}

func (p *ContractManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ContractManager) GetLast(items *[]Contract) *Contract {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *ContractManager) SetIndex(index string) {
    p.Index = index
}

func (p *ContractManager) GetQuery() string {
    ret := ""

    tableName := "contract_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".contract_tb"
    }

    str := "select co_id, co_company, co_contractenddate, co_contractstartdate, co_date, co_status from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *ContractManager) GetQuerySelect() string {
    ret := ""

    tableName := "contract_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".contract_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *ContractManager) Insert(item *Contract) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Contractenddate == "" {
        item.Contractenddate = "1000-01-01 00:00:00"
    }

    if item.Contractstartdate == "" {
        item.Contractstartdate = "1000-01-01 00:00:00"
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "contract_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".contract_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (co_id, co_company, co_contractenddate, co_contractstartdate, co_date, co_status) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Company, item.Contractenddate, item.Contractstartdate, item.Date, item.Status)
    } else {
        query = "insert into " + tableName + " (co_company, co_contractenddate, co_contractstartdate, co_date, co_status) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Company, item.Contractenddate, item.Contractstartdate, item.Date, item.Status)
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
func (p *ContractManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "contract_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".contract_tb"
    }
    query := "delete from " + tableName + " where co_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *ContractManager) Update(item *Contract) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Contractenddate == "" {
        item.Contractenddate = "1000-01-01 00:00:00"
    }

    if item.Contractstartdate == "" {
        item.Contractstartdate = "1000-01-01 00:00:00"
    }

    tableName := "contract_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".contract_tb"
    }

	query := "update " + tableName + " set co_company = ?,co_contractenddate = ?,co_contractstartdate = ?,co_date = ?,co_status = ? where co_id = ?"
	_, err := p.Conn.Exec(query, item.Company, item.Contractenddate, item.Contractstartdate, item.Date, item.Status, item.Id)

    return err
}

func (p *ContractManager) GetIdentity() int64 {
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
    

func (p *ContractManager) ReadRow(rows *sql.Rows) *Contract {
    var item Contract
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Company, &item.Contractenddate, &item.Contractstartdate, &item.Date, &item.Status)

        if item.Contractenddate == "1000-01-01 00:00:00" {
            item.Contractenddate = ""
        }

        if item.Contractstartdate == "1000-01-01 00:00:00" {
            item.Contractstartdate = ""
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

func (p *ContractManager) ReadRows(rows *sql.Rows) *[]Contract {
    var items []Contract
    var err error

    for rows.Next() {
        var item Contract
        err = rows.Scan(&item.Id, &item.Company, &item.Contractenddate, &item.Contractstartdate, &item.Date, &item.Status)

        if item.Contractenddate == "1000-01-01 00:00:00" {
            item.Contractenddate = ""
        }

        if item.Contractstartdate == "1000-01-01 00:00:00" {
            item.Contractstartdate = ""
        }

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *ContractManager) Get(id int64) *Contract {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where co_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ContractManager) GetList(page int, pagesize int, order string) *[]Contract {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "co_id desc"
        } else {
            order = "co_" + order
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
            order = "co_id"
        } else {
            order = "co_" + order
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


func (p *ContractManager) GetCount() int {
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

func (p *ContractManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Contract {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where co_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "co_id desc"
        } else {
            order = "co_" + order
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
            order = "co_id"
        } else {
            order = "co_" + order
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


func (p *ContractManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where co_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

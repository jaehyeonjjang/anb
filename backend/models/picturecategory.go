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

type Picturecategory struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Name                     string `json:"Name, string" form:"name"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Order                    int `json:"Order, int" form:"order"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type PicturecategoryManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewPicturecategoryManager(conn *sql.DB) *PicturecategoryManager {
    var item PicturecategoryManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "pc"
    item.Index = ""

    return &item
}

func (p *PicturecategoryManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PicturecategoryManager) GetLast(items *[]Picturecategory) *Picturecategory {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *PicturecategoryManager) SetIndex(index string) {
    p.Index = index
}

func (p *PicturecategoryManager) GetQuery() string {
    ret := ""

    tableName := "picturecategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picturecategory_tb"
    }

    str := "select pc_id, pc_name, pc_company, pc_order, pc_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *PicturecategoryManager) GetQuerySelect() string {
    ret := ""

    tableName := "picturecategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picturecategory_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *PicturecategoryManager) Insert(item *Picturecategory) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "picturecategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picturecategory_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (pc_id, pc_name, pc_company, pc_order, pc_date) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Name, item.Company, item.Order, item.Date)
    } else {
        query = "insert into " + tableName + " (pc_name, pc_company, pc_order, pc_date) values (?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Name, item.Company, item.Order, item.Date)
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
func (p *PicturecategoryManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "picturecategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picturecategory_tb"
    }
    query := "delete from " + tableName + " where pc_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *PicturecategoryManager) Update(item *Picturecategory) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "picturecategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picturecategory_tb"
    }

	query := "update " + tableName + " set pc_name = ?,pc_company = ?,pc_order = ?,pc_date = ? where pc_id = ?"
	_, err := p.Conn.Exec(query, item.Name, item.Company, item.Order, item.Date, item.Id)

    return err
}

func (p *PicturecategoryManager) GetIdentity() int64 {
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
    

func (p *PicturecategoryManager) ReadRow(rows *sql.Rows) *Picturecategory {
    var item Picturecategory
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Company, &item.Order, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *PicturecategoryManager) ReadRows(rows *sql.Rows) *[]Picturecategory {
    var items []Picturecategory
    var err error

    for rows.Next() {
        var item Picturecategory
        err = rows.Scan(&item.Id, &item.Name, &item.Company, &item.Order, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *PicturecategoryManager) Get(id int64) *Picturecategory {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where pc_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *PicturecategoryManager) GetList(page int, pagesize int, order string) *[]Picturecategory {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "pc_id desc"
        } else {
            order = "pc_" + order
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
            order = "pc_id"
        } else {
            order = "pc_" + order
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


func (p *PicturecategoryManager) GetCount() int {
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

func (p *PicturecategoryManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Picturecategory {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where pc_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "pc_id desc"
        } else {
            order = "pc_" + order
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
            order = "pc_id"
        } else {
            order = "pc_" + order
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


func (p *PicturecategoryManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where pc_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

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

type Statuscategory struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Name                     string `json:"Name, string" form:"name"`
    Type                     int `json:"Type, int" form:"type"`
    Floortype                int `json:"Floortype, int" form:"floortype"`
    Order                    int `json:"Order, int" form:"order"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type StatuscategoryManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewStatuscategoryManager(conn *sql.DB) *StatuscategoryManager {
    var item StatuscategoryManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "sc"
    item.Index = ""

    return &item
}

func (p *StatuscategoryManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *StatuscategoryManager) GetLast(items *[]Statuscategory) *Statuscategory {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *StatuscategoryManager) SetIndex(index string) {
    p.Index = index
}

func (p *StatuscategoryManager) GetQuery() string {
    ret := ""

    tableName := "statuscategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".statuscategory_tb"
    }

    str := "select sc_id, sc_name, sc_type, sc_floortype, sc_order, sc_company, sc_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *StatuscategoryManager) GetQuerySelect() string {
    ret := ""

    tableName := "statuscategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".statuscategory_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *StatuscategoryManager) Insert(item *Statuscategory) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "statuscategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".statuscategory_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (sc_id, sc_name, sc_type, sc_floortype, sc_order, sc_company, sc_date) values (?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Name, item.Type, item.Floortype, item.Order, item.Company, item.Date)
    } else {
        query = "insert into " + tableName + " (sc_name, sc_type, sc_floortype, sc_order, sc_company, sc_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Name, item.Type, item.Floortype, item.Order, item.Company, item.Date)
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
func (p *StatuscategoryManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "statuscategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".statuscategory_tb"
    }
    query := "delete from " + tableName + " where sc_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *StatuscategoryManager) Update(item *Statuscategory) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "statuscategory_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".statuscategory_tb"
    }

	query := "update " + tableName + " set sc_name = ?,sc_type = ?,sc_floortype = ?,sc_order = ?,sc_company = ?,sc_date = ? where sc_id = ?"
	_, err := p.Conn.Exec(query, item.Name, item.Type, item.Floortype, item.Order, item.Company, item.Date, item.Id)

    return err
}

func (p *StatuscategoryManager) GetIdentity() int64 {
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
    

func (p *StatuscategoryManager) ReadRow(rows *sql.Rows) *Statuscategory {
    var item Statuscategory
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Type, &item.Floortype, &item.Order, &item.Company, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *StatuscategoryManager) ReadRows(rows *sql.Rows) *[]Statuscategory {
    var items []Statuscategory
    var err error

    for rows.Next() {
        var item Statuscategory
        err = rows.Scan(&item.Id, &item.Name, &item.Type, &item.Floortype, &item.Order, &item.Company, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *StatuscategoryManager) Get(id int64) *Statuscategory {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where sc_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *StatuscategoryManager) GetList(page int, pagesize int, order string) *[]Statuscategory {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "sc_id desc"
        } else {
            order = "sc_" + order
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
            order = "sc_id"
        } else {
            order = "sc_" + order
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


func (p *StatuscategoryManager) GetCount() int {
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

func (p *StatuscategoryManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Statuscategory {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where sc_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "sc_id desc"
        } else {
            order = "sc_" + order
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
            order = "sc_id"
        } else {
            order = "sc_" + order
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


func (p *StatuscategoryManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where sc_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *StatuscategoryManager) GetListByType(typeid int, page int, pagesize int, orderby string) *[]Statuscategory {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if typeid != 0 {
		query += " and sc_type = ?"
		params = append(params, typeid)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "sc_id desc"
        } else {
            orderby = "sc_" + orderby
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
            orderby = "sc_id"
        } else {
            orderby = "sc_" + orderby
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

func (p *StatuscategoryManager) GetCountByType(typeid int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if typeid != 0 {
		query += " and sc_type = ?"
		params = append(params, typeid)
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

func (p *StatuscategoryManager) GetListByCompany(company int64, page int, pagesize int, orderby string) *[]Statuscategory {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if company != 0 {
		query += " and sc_company = ?"
		params = append(params, company)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "sc_id desc"
        } else {
            orderby = "sc_" + orderby
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
            orderby = "sc_id"
        } else {
            orderby = "sc_" + orderby
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

func (p *StatuscategoryManager) GetCountByCompany(company int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if company != 0 {
		query += " and sc_company = ?"
		params = append(params, company)
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

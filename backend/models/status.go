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

type Status struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Name                     string `json:"Name, string" form:"name"`
    Statuscategory           int64 `json:"Statuscategory, int64" form:"statuscategory"`
    Type                     int `json:"Type, int" form:"type"`
    Content                  string `json:"Content, string" form:"content"`
    Etc                      string `json:"Etc, string" form:"etc"`
    Order                    int `json:"Order, int" form:"order"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type StatusManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewStatusManager(conn *sql.DB) *StatusManager {
    var item StatusManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "s"
    item.Index = ""

    return &item
}

func (p *StatusManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *StatusManager) GetLast(items *[]Status) *Status {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *StatusManager) SetIndex(index string) {
    p.Index = index
}

func (p *StatusManager) GetQuery() string {
    ret := ""

    tableName := "status_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".status_tb"
    }

    str := "select s_id, s_name, s_statuscategory, s_type, s_content, s_etc, s_order, s_company, s_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *StatusManager) GetQuerySelect() string {
    ret := ""

    tableName := "status_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".status_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *StatusManager) Insert(item *Status) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "status_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".status_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (s_id, s_name, s_statuscategory, s_type, s_content, s_etc, s_order, s_company, s_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Name, item.Statuscategory, item.Type, item.Content, item.Etc, item.Order, item.Company, item.Date)
    } else {
        query = "insert into " + tableName + " (s_name, s_statuscategory, s_type, s_content, s_etc, s_order, s_company, s_date) values (?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Name, item.Statuscategory, item.Type, item.Content, item.Etc, item.Order, item.Company, item.Date)
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
func (p *StatusManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "status_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".status_tb"
    }
    query := "delete from " + tableName + " where s_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *StatusManager) Update(item *Status) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "status_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".status_tb"
    }

	query := "update " + tableName + " set s_name = ?,s_statuscategory = ?,s_type = ?,s_content = ?,s_etc = ?,s_order = ?,s_company = ?,s_date = ? where s_id = ?"
	_, err := p.Conn.Exec(query, item.Name, item.Statuscategory, item.Type, item.Content, item.Etc, item.Order, item.Company, item.Date, item.Id)

    return err
}

func (p *StatusManager) GetIdentity() int64 {
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
    

func (p *StatusManager) ReadRow(rows *sql.Rows) *Status {
    var item Status
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Statuscategory, &item.Type, &item.Content, &item.Etc, &item.Order, &item.Company, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *StatusManager) ReadRows(rows *sql.Rows) *[]Status {
    var items []Status
    var err error

    for rows.Next() {
        var item Status
        err = rows.Scan(&item.Id, &item.Name, &item.Statuscategory, &item.Type, &item.Content, &item.Etc, &item.Order, &item.Company, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *StatusManager) Get(id int64) *Status {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where s_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *StatusManager) GetList(page int, pagesize int, order string) *[]Status {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "s_id desc"
        } else {
            order = "s_" + order
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
            order = "s_id"
        } else {
            order = "s_" + order
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


func (p *StatusManager) GetCount() int {
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

func (p *StatusManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Status {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where s_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "s_id desc"
        } else {
            order = "s_" + order
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
            order = "s_id"
        } else {
            order = "s_" + order
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


func (p *StatusManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where s_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *StatusManager) GetListByType(typeid int, page int, pagesize int, orderby string) *[]Status {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if typeid != 0 {
		query += " and s_type = ?"
		params = append(params, typeid)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "s_id desc"
        } else {
            orderby = "s_" + orderby
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
            orderby = "s_id"
        } else {
            orderby = "s_" + orderby
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

func (p *StatusManager) GetCountByType(typeid int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if typeid != 0 {
		query += " and s_type = ?"
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

func (p *StatusManager) GetListByCompany(company int64, page int, pagesize int, orderby string) *[]Status {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if company != 0 {
		query += " and s_company = ?"
		params = append(params, company)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "s_id desc"
        } else {
            orderby = "s_" + orderby
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
            orderby = "s_id"
        } else {
            orderby = "s_" + orderby
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

func (p *StatusManager) GetCountByCompany(company int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if company != 0 {
		query += " and s_company = ?"
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

func (p *StatusManager) GetByCompanyTypeName(company int64, typeid int, name string) *Status {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if company != 0 {
		query += " and s_company = ?"
		params = append(params, company)
	}
	if typeid != 0 {
		query += " and s_type = ?"
		params = append(params, typeid)
	}
	if name != "" {
		query += " and s_name = ?"
		params = append(params, name)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *StatusManager) GetListByCompanyType(company int64, typeid int, page int, pagesize int, orderby string) *[]Status {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if company != 0 {
		query += " and s_company = ?"
		params = append(params, company)
	}
	if typeid != 0 {
		query += " and s_type = ?"
		params = append(params, typeid)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "s_id desc"
        } else {
            orderby = "s_" + orderby
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
            orderby = "s_id"
        } else {
            orderby = "s_" + orderby
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

func (p *StatusManager) GetCountByCompanyType(company int64, typeid int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if company != 0 {
		query += " and s_company = ?"
		params = append(params, company)
	}
	if typeid != 0 {
		query += " and s_type = ?"
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

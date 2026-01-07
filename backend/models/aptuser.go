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

type Aptuser struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Apt                      int64 `json:"Apt, int64" form:"apt"`
    User                     int64 `json:"User, int64" form:"user"`
    Level                    int `json:"Level, int" form:"level"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type AptuserManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewAptuserManager(conn *sql.DB) *AptuserManager {
    var item AptuserManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "au"
    item.Index = ""

    return &item
}

func (p *AptuserManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptuserManager) GetLast(items *[]Aptuser) *Aptuser {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *AptuserManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptuserManager) GetQuery() string {
    ret := ""

    tableName := "aptuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuser_tb"
    }

    str := "select au_id, au_apt, au_user, au_level, au_company, au_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptuserManager) GetQuerySelect() string {
    ret := ""

    tableName := "aptuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuser_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptuserManager) Insert(item *Aptuser) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "aptuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuser_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (au_id, au_apt, au_user, au_level, au_company, au_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Apt, item.User, item.Level, item.Company, item.Date)
    } else {
        query = "insert into " + tableName + " (au_apt, au_user, au_level, au_company, au_date) values (?, ?, ?, ?, ?)"
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
func (p *AptuserManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuser_tb"
    }
    query := "delete from " + tableName + " where au_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}

func (p *AptuserManager) DeleteByApt(apt int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuser_tb"
    }
    var params []interface{}
    query := "delete from " + tableName + " where au_apt = ?"
	params = append(params, apt)

    err := ExecArray(p.Conn, query, params)

    return err
}

func (p *AptuserManager) Update(item *Aptuser) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptuser_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuser_tb"
    }

	query := "update " + tableName + " set au_apt = ?,au_user = ?,au_level = ?,au_company = ?,au_date = ? where au_id = ?"
	_, err := p.Conn.Exec(query, item.Apt, item.User, item.Level, item.Company, item.Date, item.Id)

    return err
}

func (p *AptuserManager) GetIdentity() int64 {
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
    

func (p *AptuserManager) ReadRow(rows *sql.Rows) *Aptuser {
    var item Aptuser
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

func (p *AptuserManager) ReadRows(rows *sql.Rows) *[]Aptuser {
    var items []Aptuser
    var err error

    for rows.Next() {
        var item Aptuser
        err = rows.Scan(&item.Id, &item.Apt, &item.User, &item.Level, &item.Company, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *AptuserManager) Get(id int64) *Aptuser {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where au_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *AptuserManager) GetList(page int, pagesize int, order string) *[]Aptuser {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "au_id desc"
        } else {
            order = "au_" + order
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
            order = "au_id"
        } else {
            order = "au_" + order
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


func (p *AptuserManager) GetCount() int {
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

func (p *AptuserManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Aptuser {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where au_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "au_id desc"
        } else {
            order = "au_" + order
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
            order = "au_id"
        } else {
            order = "au_" + order
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


func (p *AptuserManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where au_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *AptuserManager) GetListByApt(apt int64, page int, pagesize int, orderby string) *[]Aptuser {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and au_apt = ?"
		params = append(params, apt)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "au_id desc"
        } else {
            orderby = "au_" + orderby
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
            orderby = "au_id"
        } else {
            orderby = "au_" + orderby
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

func (p *AptuserManager) GetCountByApt(apt int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and au_apt = ?"
		params = append(params, apt)
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

func (p *AptuserManager) GetListByAptLevel(apt int64, level int, page int, pagesize int, orderby string) *[]Aptuser {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and au_apt = ?"
		params = append(params, apt)
	}
	if level != 0 {
		query += " and au_level = ?"
		params = append(params, level)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "au_id desc"
        } else {
            orderby = "au_" + orderby
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
            orderby = "au_id"
        } else {
            orderby = "au_" + orderby
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

func (p *AptuserManager) GetCountByAptLevel(apt int64, level int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and au_apt = ?"
		params = append(params, apt)
	}
	if level != 0 {
		query += " and au_level = ?"
		params = append(params, level)
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

func (p *AptuserManager) GetByAptUser(apt int64, user int64) *Aptuser {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and au_apt = ?"
		params = append(params, apt)
	}
	if user != 0 {
		query += " and au_user = ?"
		params = append(params, user)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *AptuserManager) GetListByUser(user int64, page int, pagesize int, orderby string) *[]Aptuser {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if user != 0 {
		query += " and au_user = ?"
		params = append(params, user)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "au_id desc"
        } else {
            orderby = "au_" + orderby
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
            orderby = "au_id"
        } else {
            orderby = "au_" + orderby
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

func (p *AptuserManager) GetCountByUser(user int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if user != 0 {
		query += " and au_user = ?"
		params = append(params, user)
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

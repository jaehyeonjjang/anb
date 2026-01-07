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

type User struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Loginid                  string `json:"Loginid, string" form:"loginid"`
    Passwd                   string `json:"Passwd, string" form:"passwd"`
    Name                     string `json:"Name, string" form:"name"`
    Level                    int `json:"Level, int" form:"level"`
    Hp                       string `json:"Hp, string" form:"hp"`
    Email                    string `json:"Email, string" form:"email"`
    Grade                    int `json:"Grade, int" form:"grade"`
    Status                   int `json:"Status, int" form:"status"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type UserManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewUserManager(conn *sql.DB) *UserManager {
    var item UserManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "u"
    item.Index = ""

    return &item
}

func (p *UserManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *UserManager) GetLast(items *[]User) *User {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *UserManager) SetIndex(index string) {
    p.Index = index
}

func (p *UserManager) GetQuery() string {
    ret := ""

    tableName := "user_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".user_tb"
    }

    str := "select u_id, u_loginid, u_passwd, u_name, u_level, u_hp, u_email, u_grade, u_status, u_company, u_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *UserManager) GetQuerySelect() string {
    ret := ""

    tableName := "user_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".user_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *UserManager) Insert(item *User) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "user_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".user_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (u_id, u_loginid, u_passwd, u_name, u_level, u_hp, u_email, u_grade, u_status, u_company, u_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Loginid, item.Passwd, item.Name, item.Level, item.Hp, item.Email, item.Grade, item.Status, item.Company, item.Date)
    } else {
        query = "insert into " + tableName + " (u_loginid, u_passwd, u_name, u_level, u_hp, u_email, u_grade, u_status, u_company, u_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Loginid, item.Passwd, item.Name, item.Level, item.Hp, item.Email, item.Grade, item.Status, item.Company, item.Date)
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
func (p *UserManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "user_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".user_tb"
    }
    query := "delete from " + tableName + " where u_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *UserManager) Update(item *User) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "user_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".user_tb"
    }

	query := "update " + tableName + " set u_loginid = ?,u_passwd = ?,u_name = ?,u_level = ?,u_hp = ?,u_email = ?,u_grade = ?,u_status = ?,u_company = ?,u_date = ? where u_id = ?"
	_, err := p.Conn.Exec(query, item.Loginid, item.Passwd, item.Name, item.Level, item.Hp, item.Email, item.Grade, item.Status, item.Company, item.Date, item.Id)

    return err
}

func (p *UserManager) GetIdentity() int64 {
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
    

func (p *UserManager) ReadRow(rows *sql.Rows) *User {
    var item User
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Level, &item.Hp, &item.Email, &item.Grade, &item.Status, &item.Company, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *UserManager) ReadRows(rows *sql.Rows) *[]User {
    var items []User
    var err error

    for rows.Next() {
        var item User
        err = rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Level, &item.Hp, &item.Email, &item.Grade, &item.Status, &item.Company, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *UserManager) Get(id int64) *User {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where u_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *UserManager) GetList(page int, pagesize int, order string) *[]User {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "u_id desc"
        } else {
            order = "u_" + order
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
            order = "u_id"
        } else {
            order = "u_" + order
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


func (p *UserManager) GetCount() int {
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

func (p *UserManager) GetListInID(ids []int, page int, pagesize int, order string) *[]User {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where u_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "u_id desc"
        } else {
            order = "u_" + order
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
            order = "u_id"
        } else {
            order = "u_" + order
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


func (p *UserManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where u_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *UserManager) GetByLoginid(loginid string) *User {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if loginid != "" {
		query += " and u_loginid = ?"
		params = append(params, loginid)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *UserManager) GetByEmail(email string) *User {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if email != "" {
		query += " and u_email = ?"
		params = append(params, email)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *UserManager) GetListByLevelNameLoginidEmail(level int, name string, loginid string, email string, page int, pagesize int, orderby string) *[]User {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if level != 0 {
		query += " and u_level = ?"
		params = append(params, level)
	}
	if name != "" {
		query += " and u_name like ?"
		name_ := "%"+name+"%"
		params = append(params, name_)
	}
	if loginid != "" {
		query += " and u_loginid like ?"
		loginid_ := "%"+loginid+"%"
		params = append(params, loginid_)
	}
	if email != "" {
		query += " and u_email like ?"
		email_ := "%"+email+"%"
		params = append(params, email_)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "u_id desc"
        } else {
            orderby = "u_" + orderby
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
            orderby = "u_id"
        } else {
            orderby = "u_" + orderby
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

func (p *UserManager) GetCountByLevelNameLoginidEmail(level int, name string, loginid string, email string) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if level != 0 {
		query += " and u_level = ?"
		params = append(params, level)
	}
	if name != "" {
		query += " and u_name like ?"
		name_ := "%"+name+"%"
		params = append(params, name_)
	}
	if loginid != "" {
		query += " and u_loginid like ?"
		loginid_ := "%"+loginid+"%"
		params = append(params, loginid_)
	}
	if email != "" {
		query += " and u_email like ?"
		email_ := "%"+email+"%"
		params = append(params, email_)
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

func (p *UserManager) GetListByLevel(level int, page int, pagesize int, orderby string) *[]User {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if level != 0 {
		query += " and u_level = ?"
		params = append(params, level)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "u_id desc"
        } else {
            orderby = "u_" + orderby
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
            orderby = "u_id"
        } else {
            orderby = "u_" + orderby
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

func (p *UserManager) GetCountByLevel(level int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if level != 0 {
		query += " and u_level = ?"
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

func (p *UserManager) GetListByCompanyStatus(company int64, status int, page int, pagesize int, orderby string) *[]User {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if company != 0 {
		query += " and u_company = ?"
		params = append(params, company)
	}
	if status != 0 {
		query += " and u_status = ?"
		params = append(params, status)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "u_id desc"
        } else {
            orderby = "u_" + orderby
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
            orderby = "u_id"
        } else {
            orderby = "u_" + orderby
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

func (p *UserManager) GetCountByCompanyStatus(company int64, status int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if company != 0 {
		query += " and u_company = ?"
		params = append(params, company)
	}
	if status != 0 {
		query += " and u_status = ?"
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

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

type Aptuserlist struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Loginid                  string `json:"Loginid, string" form:"loginid"`
    Passwd                   string `json:"Passwd, string" form:"passwd"`
    Name                     string `json:"Name, string" form:"name"`
    Level                    int `json:"Level, int" form:"level"`
    Hp                       string `json:"Hp, string" form:"hp"`
    Email                    string `json:"Email, string" form:"email"`
    Status                   int `json:"Status, int" form:"status"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Date                     string `json:"Date, string" form:"date"`
    User                     int64 `json:"User, int64" form:"user"`
    Apt                      int64 `json:"Apt, int64" form:"apt"`
    Aptlevel                 int `json:"Aptlevel, int" form:"aptlevel"`
    Aptname                  string `json:"Aptname, string" form:"aptname"`
    Aptstatus                int `json:"Aptstatus, int" form:"aptstatus"`
    Extra                    interface{} `form:"extra"`
}

type AptuserlistManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewAptuserlistManager(conn *sql.DB) *AptuserlistManager {
    var item AptuserlistManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "u"
    item.Index = ""

    return &item
}

func (p *AptuserlistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptuserlistManager) GetLast(items *[]Aptuserlist) *Aptuserlist {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *AptuserlistManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptuserlistManager) GetQuery() string {
    ret := ""

    tableName := "aptuserlist_vw"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuserlist_vw"
    }

    str := "select u_id, u_loginid, u_passwd, u_name, u_level, u_hp, u_email, u_status, u_company, u_date, u_user, u_apt, u_aptlevel, u_aptname, u_aptstatus from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptuserlistManager) GetQuerySelect() string {
    ret := ""

    tableName := "aptuserlist_vw"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuserlist_vw"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptuserlistManager) Insert(item *Aptuserlist) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "aptuserlist_vw"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuserlist_vw"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (u_id, u_loginid, u_passwd, u_name, u_level, u_hp, u_email, u_status, u_company, u_date, u_user, u_apt, u_aptlevel, u_aptname, u_aptstatus) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Loginid, item.Passwd, item.Name, item.Level, item.Hp, item.Email, item.Status, item.Company, item.Date, item.User, item.Apt, item.Aptlevel, item.Aptname, item.Aptstatus)
    } else {
        query = "insert into " + tableName + " (u_loginid, u_passwd, u_name, u_level, u_hp, u_email, u_status, u_company, u_date, u_user, u_apt, u_aptlevel, u_aptname, u_aptstatus) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Loginid, item.Passwd, item.Name, item.Level, item.Hp, item.Email, item.Status, item.Company, item.Date, item.User, item.Apt, item.Aptlevel, item.Aptname, item.Aptstatus)
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
func (p *AptuserlistManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptuserlist_vw"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuserlist_vw"
    }
    query := "delete from " + tableName + " where u_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *AptuserlistManager) Update(item *Aptuserlist) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptuserlist_vw"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptuserlist_vw"
    }

	query := "update " + tableName + " set u_loginid = ?,u_passwd = ?,u_name = ?,u_level = ?,u_hp = ?,u_email = ?,u_status = ?,u_company = ?,u_date = ?,u_user = ?,u_apt = ?,u_aptlevel = ?,u_aptname = ?,u_aptstatus = ? where u_id = ?"
	_, err := p.Conn.Exec(query, item.Loginid, item.Passwd, item.Name, item.Level, item.Hp, item.Email, item.Status, item.Company, item.Date, item.User, item.Apt, item.Aptlevel, item.Aptname, item.Aptstatus, item.Id)

    return err
}

func (p *AptuserlistManager) GetIdentity() int64 {
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
    

func (p *AptuserlistManager) ReadRow(rows *sql.Rows) *Aptuserlist {
    var item Aptuserlist
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Level, &item.Hp, &item.Email, &item.Status, &item.Company, &item.Date, &item.User, &item.Apt, &item.Aptlevel, &item.Aptname, &item.Aptstatus)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *AptuserlistManager) ReadRows(rows *sql.Rows) *[]Aptuserlist {
    var items []Aptuserlist
    var err error

    for rows.Next() {
        var item Aptuserlist
        err = rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Level, &item.Hp, &item.Email, &item.Status, &item.Company, &item.Date, &item.User, &item.Apt, &item.Aptlevel, &item.Aptname, &item.Aptstatus)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *AptuserlistManager) Get(id int64) *Aptuserlist {
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

func (p *AptuserlistManager) GetList(page int, pagesize int, order string) *[]Aptuserlist {
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


func (p *AptuserlistManager) GetCount() int {
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

func (p *AptuserlistManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Aptuserlist {
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


func (p *AptuserlistManager) GetCountInID(ids []int) int {
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

func (p *AptuserlistManager) GetListByUser(user int64, page int, pagesize int, orderby string) *[]Aptuserlist {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if user != 0 {
		query += " and u_user = ?"
		params = append(params, user)
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

func (p *AptuserlistManager) GetCountByUser(user int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if user != 0 {
		query += " and u_user = ?"
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

func (p *AptuserlistManager) GetListByApt(apt int64, page int, pagesize int, orderby string) *[]Aptuserlist {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and u_apt = ?"
		params = append(params, apt)
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

func (p *AptuserlistManager) GetCountByApt(apt int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and u_apt = ?"
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

func (p *AptuserlistManager) GetListByAptAptlevel(apt int64, aptlevel int, page int, pagesize int, orderby string) *[]Aptuserlist {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and u_apt = ?"
		params = append(params, apt)
	}
	if aptlevel != 0 {
		query += " and u_aptlevel = ?"
		params = append(params, aptlevel)
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

func (p *AptuserlistManager) GetCountByAptAptlevel(apt int64, aptlevel int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and u_apt = ?"
		params = append(params, apt)
	}
	if aptlevel != 0 {
		query += " and u_aptlevel = ?"
		params = append(params, aptlevel)
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

func (p *AptuserlistManager) GetByAptUser(apt int64, user int64) *Aptuserlist {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and u_apt = ?"
		params = append(params, apt)
	}
	if user != 0 {
		query += " and u_user = ?"
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

func (p *AptuserlistManager) GetListByUserAptstatus(user int64, aptstatus int, page int, pagesize int, orderby string) *[]Aptuserlist {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if user != 0 {
		query += " and u_user = ?"
		params = append(params, user)
	}
	if aptstatus != 0 {
		query += " and u_aptstatus = ?"
		params = append(params, aptstatus)
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

func (p *AptuserlistManager) GetCountByUserAptstatus(user int64, aptstatus int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if user != 0 {
		query += " and u_user = ?"
		params = append(params, user)
	}
	if aptstatus != 0 {
		query += " and u_aptstatus = ?"
		params = append(params, aptstatus)
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

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

type Apt struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Aptgroup                 int64 `json:"Aptgroup, int64" form:"aptgroup"`
    Name                     string `json:"Name, string" form:"name"`
    Startdate                string `json:"Startdate, string" form:"startdate"`
    Enddate                  string `json:"Enddate, string" form:"enddate"`
    Type                     int `json:"Type, int" form:"type"`
    Master                   int64 `json:"Master, int64" form:"master"`
    Status                   int `json:"Status, int" form:"status"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Report                   int `json:"Report, int" form:"report"`
    Report1                  int `json:"Report1, int" form:"report1"`
    Report2                  int `json:"Report2, int" form:"report2"`
    Report3                  int `json:"Report3, int" form:"report3"`
    Report4                  int `json:"Report4, int" form:"report4"`
    Report5                  int `json:"Report5, int" form:"report5"`
    Report6                  int `json:"Report6, int" form:"report6"`
    Summarytype              int `json:"Summarytype, int" form:"summarytype"`
    Search                   string `json:"Search, string" form:"search"`
    User                     int64 `json:"User, int64" form:"user"`
    Updateuser               int64 `json:"Updateuser, int64" form:"updateuser"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type AptManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewAptManager(conn *sql.DB) *AptManager {
    var item AptManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "a"
    item.Index = ""

    return &item
}

func (p *AptManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptManager) GetLast(items *[]Apt) *Apt {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *AptManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptManager) GetQuery() string {
    ret := ""

    tableName := "apt_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".apt_tb"
    }

    str := "select a_id, a_aptgroup, a_name, a_startdate, a_enddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptManager) GetQuerySelect() string {
    ret := ""

    tableName := "apt_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".apt_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptManager) Insert(item *Apt) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "apt_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".apt_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (a_id, a_aptgroup, a_name, a_startdate, a_enddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Aptgroup, item.Name, item.Startdate, item.Enddate, item.Type, item.Master, item.Status, item.Company, item.Report, item.Report1, item.Report2, item.Report3, item.Report4, item.Report5, item.Report6, item.Summarytype, item.Search, item.User, item.Updateuser, item.Date)
    } else {
        query = "insert into " + tableName + " (a_aptgroup, a_name, a_startdate, a_enddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Aptgroup, item.Name, item.Startdate, item.Enddate, item.Type, item.Master, item.Status, item.Company, item.Report, item.Report1, item.Report2, item.Report3, item.Report4, item.Report5, item.Report6, item.Summarytype, item.Search, item.User, item.Updateuser, item.Date)
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
func (p *AptManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "apt_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".apt_tb"
    }
    query := "delete from " + tableName + " where a_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *AptManager) Update(item *Apt) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "apt_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".apt_tb"
    }

	query := "update " + tableName + " set a_aptgroup = ?,a_name = ?,a_startdate = ?,a_enddate = ?,a_type = ?,a_master = ?,a_status = ?,a_company = ?,a_report = ?,a_report1 = ?,a_report2 = ?,a_report3 = ?,a_report4 = ?,a_report5 = ?,a_report6 = ?,a_summarytype = ?,a_search = ?,a_user = ?,a_updateuser = ?,a_date = ? where a_id = ?"
	_, err := p.Conn.Exec(query, item.Aptgroup, item.Name, item.Startdate, item.Enddate, item.Type, item.Master, item.Status, item.Company, item.Report, item.Report1, item.Report2, item.Report3, item.Report4, item.Report5, item.Report6, item.Summarytype, item.Search, item.User, item.Updateuser, item.Date, item.Id)

    return err
}


func (p *AptManager) UpdateReportById(report int, id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "apt_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".apt_tb"
    }

	query := "update " + tableName + " set a_report = ? where a_id = ?"
	_, err := p.Conn.Exec(query, report, id)

    return err
}

func (p *AptManager) GetIdentity() int64 {
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
    

func (p *AptManager) ReadRow(rows *sql.Rows) *Apt {
    var item Apt
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Aptgroup, &item.Name, &item.Startdate, &item.Enddate, &item.Type, &item.Master, &item.Status, &item.Company, &item.Report, &item.Report1, &item.Report2, &item.Report3, &item.Report4, &item.Report5, &item.Report6, &item.Summarytype, &item.Search, &item.User, &item.Updateuser, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *AptManager) ReadRows(rows *sql.Rows) *[]Apt {
    var items []Apt
    var err error

    for rows.Next() {
        var item Apt
        err = rows.Scan(&item.Id, &item.Aptgroup, &item.Name, &item.Startdate, &item.Enddate, &item.Type, &item.Master, &item.Status, &item.Company, &item.Report, &item.Report1, &item.Report2, &item.Report3, &item.Report4, &item.Report5, &item.Report6, &item.Summarytype, &item.Search, &item.User, &item.Updateuser, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *AptManager) Get(id int64) *Apt {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where a_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *AptManager) GetList(page int, pagesize int, order string) *[]Apt {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "a_id desc"
        } else {
            order = "a_" + order
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
            order = "a_id"
        } else {
            order = "a_" + order
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


func (p *AptManager) GetCount() int {
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

func (p *AptManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Apt {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where a_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "a_id desc"
        } else {
            order = "a_" + order
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
            order = "a_id"
        } else {
            order = "a_" + order
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


func (p *AptManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where a_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *AptManager) GetListByName(name string, page int, pagesize int, orderby string) *[]Apt {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if name != "" {
		query += " and a_name like ?"
		name_ := "%"+name+"%"
		params = append(params, name_)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "a_id desc"
        } else {
            orderby = "a_" + orderby
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
            orderby = "a_id"
        } else {
            orderby = "a_" + orderby
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

func (p *AptManager) GetCountByName(name string) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if name != "" {
		query += " and a_name like ?"
		name_ := "%"+name+"%"
		params = append(params, name_)
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

func (p *AptManager) GetListByCompanyStatus(company int64, status int, page int, pagesize int, orderby string) *[]Apt {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if company != 0 {
		query += " and a_company = ?"
		params = append(params, company)
	}
	if status != 0 {
		query += " and a_status = ?"
		params = append(params, status)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "a_id desc"
        } else {
            orderby = "a_" + orderby
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
            orderby = "a_id"
        } else {
            orderby = "a_" + orderby
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

func (p *AptManager) GetCountByCompanyStatus(company int64, status int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if company != 0 {
		query += " and a_company = ?"
		params = append(params, company)
	}
	if status != 0 {
		query += " and a_status = ?"
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

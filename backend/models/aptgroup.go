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

type Aptgroup struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Company                  int64 `json:"Company, int64" form:"company"`
    Name                     string `json:"Name, string" form:"name"`
    Facility                 int `json:"Facility, int" form:"facility"`
    Type                     int `json:"Type, int" form:"type"`
    Status                   int `json:"Status, int" form:"status"`
    User                     int64 `json:"User, int64" form:"user"`
    Updateuser               int64 `json:"Updateuser, int64" form:"updateuser"`
    Date                     string `json:"Date, string" form:"date"`
    Imagecategory            string `json:"Imagecategory, string" form:"imagecategory"`
    Extra                    interface{} `form:"extra"`
}

type AptgroupManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewAptgroupManager(conn *sql.DB) *AptgroupManager {
    var item AptgroupManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "ag"
    item.Index = ""

    return &item
}

func (p *AptgroupManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptgroupManager) GetLast(items *[]Aptgroup) *Aptgroup {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *AptgroupManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptgroupManager) GetQuery() string {
    ret := ""

    tableName := "aptgroup_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptgroup_tb"
    }

    str := "select ag_id, ag_company, ag_name, ag_facility, ag_type, ag_status, ag_user, ag_updateuser, ag_date, ag_imagecategory from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptgroupManager) GetQuerySelect() string {
    ret := ""

    tableName := "aptgroup_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptgroup_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *AptgroupManager) Insert(item *Aptgroup) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "aptgroup_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptgroup_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (ag_id, ag_company, ag_name, ag_facility, ag_type, ag_status, ag_user, ag_updateuser, ag_date, ag_imagecategory) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Company, item.Name, item.Facility, item.Type, item.Status, item.User, item.Updateuser, item.Date, item.Imagecategory)
    } else {
        query = "insert into " + tableName + " (ag_company, ag_name, ag_facility, ag_type, ag_status, ag_user, ag_updateuser, ag_date, ag_imagecategory) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Company, item.Name, item.Facility, item.Type, item.Status, item.User, item.Updateuser, item.Date, item.Imagecategory)
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
func (p *AptgroupManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptgroup_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptgroup_tb"
    }
    query := "delete from " + tableName + " where ag_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *AptgroupManager) Update(item *Aptgroup) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "aptgroup_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".aptgroup_tb"
    }

	query := "update " + tableName + " set ag_company = ?,ag_name = ?,ag_facility = ?,ag_type = ?,ag_status = ?,ag_user = ?,ag_updateuser = ?,ag_date = ?,ag_imagecategory = ? where ag_id = ?"
	_, err := p.Conn.Exec(query, item.Company, item.Name, item.Facility, item.Type, item.Status, item.User, item.Updateuser, item.Date, item.Imagecategory, item.Id)

    return err
}

func (p *AptgroupManager) GetIdentity() int64 {
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
    

func (p *AptgroupManager) ReadRow(rows *sql.Rows) *Aptgroup {
    var item Aptgroup
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Company, &item.Name, &item.Facility, &item.Type, &item.Status, &item.User, &item.Updateuser, &item.Date, &item.Imagecategory)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *AptgroupManager) ReadRows(rows *sql.Rows) *[]Aptgroup {
    var items []Aptgroup
    var err error

    for rows.Next() {
        var item Aptgroup
        err = rows.Scan(&item.Id, &item.Company, &item.Name, &item.Facility, &item.Type, &item.Status, &item.User, &item.Updateuser, &item.Date, &item.Imagecategory)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *AptgroupManager) Get(id int64) *Aptgroup {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where ag_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *AptgroupManager) GetList(page int, pagesize int, order string) *[]Aptgroup {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "ag_id desc"
        } else {
            order = "ag_" + order
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
            order = "ag_id"
        } else {
            order = "ag_" + order
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


func (p *AptgroupManager) GetCount() int {
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

func (p *AptgroupManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Aptgroup {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where ag_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "ag_id desc"
        } else {
            order = "ag_" + order
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
            order = "ag_id"
        } else {
            order = "ag_" + order
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


func (p *AptgroupManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where ag_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

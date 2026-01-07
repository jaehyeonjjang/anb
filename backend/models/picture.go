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

type Picture struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Picturecategory          int64 `json:"Picturecategory, int64" form:"picturecategory"`
    Filename                 string `json:"Filename, string" form:"filename"`
    Category                 string `json:"Category, string" form:"category"`
    Content                  string `json:"Content, string" form:"content"`
    Apt                      int64 `json:"Apt, int64" form:"apt"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type PictureManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewPictureManager(conn *sql.DB) *PictureManager {
    var item PictureManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "p"
    item.Index = ""

    return &item
}

func (p *PictureManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PictureManager) GetLast(items *[]Picture) *Picture {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *PictureManager) SetIndex(index string) {
    p.Index = index
}

func (p *PictureManager) GetQuery() string {
    ret := ""

    tableName := "picture_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picture_tb"
    }

    str := "select p_id, p_picturecategory, p_filename, p_category, p_content, p_apt, p_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *PictureManager) GetQuerySelect() string {
    ret := ""

    tableName := "picture_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picture_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *PictureManager) Insert(item *Picture) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "picture_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picture_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (p_id, p_picturecategory, p_filename, p_category, p_content, p_apt, p_date) values (?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Picturecategory, item.Filename, item.Category, item.Content, item.Apt, item.Date)
    } else {
        query = "insert into " + tableName + " (p_picturecategory, p_filename, p_category, p_content, p_apt, p_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Picturecategory, item.Filename, item.Category, item.Content, item.Apt, item.Date)
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
func (p *PictureManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "picture_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picture_tb"
    }
    query := "delete from " + tableName + " where p_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *PictureManager) Update(item *Picture) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "picture_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".picture_tb"
    }

	query := "update " + tableName + " set p_picturecategory = ?,p_filename = ?,p_category = ?,p_content = ?,p_apt = ?,p_date = ? where p_id = ?"
	_, err := p.Conn.Exec(query, item.Picturecategory, item.Filename, item.Category, item.Content, item.Apt, item.Date, item.Id)

    return err
}

func (p *PictureManager) GetIdentity() int64 {
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
    

func (p *PictureManager) ReadRow(rows *sql.Rows) *Picture {
    var item Picture
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Picturecategory, &item.Filename, &item.Category, &item.Content, &item.Apt, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *PictureManager) ReadRows(rows *sql.Rows) *[]Picture {
    var items []Picture
    var err error

    for rows.Next() {
        var item Picture
        err = rows.Scan(&item.Id, &item.Picturecategory, &item.Filename, &item.Category, &item.Content, &item.Apt, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *PictureManager) Get(id int64) *Picture {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where p_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *PictureManager) GetList(page int, pagesize int, order string) *[]Picture {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "p_id desc"
        } else {
            order = "p_" + order
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
            order = "p_id"
        } else {
            order = "p_" + order
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


func (p *PictureManager) GetCount() int {
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

func (p *PictureManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Picture {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where p_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "p_id desc"
        } else {
            order = "p_" + order
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
            order = "p_id"
        } else {
            order = "p_" + order
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


func (p *PictureManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where p_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *PictureManager) GetListByApt(apt int64, page int, pagesize int, orderby string) *[]Picture {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and p_apt = ?"
		params = append(params, apt)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "p_id desc"
        } else {
            orderby = "p_" + orderby
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
            orderby = "p_id"
        } else {
            orderby = "p_" + orderby
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

func (p *PictureManager) GetCountByApt(apt int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and p_apt = ?"
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

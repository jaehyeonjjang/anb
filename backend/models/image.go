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

type Image struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Apt                      int64 `json:"Apt, int64" form:"apt"`
    Name                     string `json:"Name, string" form:"name"`
    Level                    int `json:"Level, int" form:"level"`
    Parent                   int64 `json:"Parent, int64" form:"parent"`
    Last                     int `json:"Last, int" form:"last"`
    Title                    string `json:"Title, string" form:"title"`
    Type                     int `json:"Type, int" form:"type"`
    Floortype                int `json:"Floortype, int" form:"floortype"`
    Filename                 string `json:"Filename, string" form:"filename"`
    Order                    int `json:"Order, int" form:"order"`
    Date                     string `json:"Date, string" form:"date"`
    Standard                 int `json:"Standard, int" form:"standard"`
    Extra                    interface{} `form:"extra"`
}

type ImageManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewImageManager(conn *sql.DB) *ImageManager {
    var item ImageManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "i"
    item.Index = ""

    return &item
}

func (p *ImageManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ImageManager) GetLast(items *[]Image) *Image {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *ImageManager) SetIndex(index string) {
    p.Index = index
}

func (p *ImageManager) GetQuery() string {
    ret := ""

    tableName := "image_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".image_tb"
    }

    str := "select i_id, i_apt, i_name, i_level, i_parent, i_last, i_title, i_type, i_floortype, i_filename, i_order, i_date, i_standard from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *ImageManager) GetQuerySelect() string {
    ret := ""

    tableName := "image_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".image_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *ImageManager) Insert(item *Image) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "image_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".image_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (i_id, i_apt, i_name, i_level, i_parent, i_last, i_title, i_type, i_floortype, i_filename, i_order, i_date, i_standard) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Apt, item.Name, item.Level, item.Parent, item.Last, item.Title, item.Type, item.Floortype, item.Filename, item.Order, item.Date, item.Standard)
    } else {
        query = "insert into " + tableName + " (i_apt, i_name, i_level, i_parent, i_last, i_title, i_type, i_floortype, i_filename, i_order, i_date, i_standard) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Apt, item.Name, item.Level, item.Parent, item.Last, item.Title, item.Type, item.Floortype, item.Filename, item.Order, item.Date, item.Standard)
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
func (p *ImageManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "image_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".image_tb"
    }
    query := "delete from " + tableName + " where i_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}

func (p *ImageManager) DeleteByApt(apt int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "image_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".image_tb"
    }
    var params []interface{}
    query := "delete from " + tableName + " where i_apt = ?"
	params = append(params, apt)

    err := ExecArray(p.Conn, query, params)

    return err
}

func (p *ImageManager) Update(item *Image) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "image_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".image_tb"
    }

	query := "update " + tableName + " set i_apt = ?,i_name = ?,i_level = ?,i_parent = ?,i_last = ?,i_title = ?,i_type = ?,i_floortype = ?,i_filename = ?,i_order = ?,i_date = ?,i_standard = ? where i_id = ?"
	_, err := p.Conn.Exec(query, item.Apt, item.Name, item.Level, item.Parent, item.Last, item.Title, item.Type, item.Floortype, item.Filename, item.Order, item.Date, item.Standard, item.Id)

    return err
}

func (p *ImageManager) GetIdentity() int64 {
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
    

func (p *ImageManager) ReadRow(rows *sql.Rows) *Image {
    var item Image
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Apt, &item.Name, &item.Level, &item.Parent, &item.Last, &item.Title, &item.Type, &item.Floortype, &item.Filename, &item.Order, &item.Date, &item.Standard)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *ImageManager) ReadRows(rows *sql.Rows) *[]Image {
    var items []Image
    var err error

    for rows.Next() {
        var item Image
        err = rows.Scan(&item.Id, &item.Apt, &item.Name, &item.Level, &item.Parent, &item.Last, &item.Title, &item.Type, &item.Floortype, &item.Filename, &item.Order, &item.Date, &item.Standard)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *ImageManager) Get(id int64) *Image {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where i_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ImageManager) GetList(page int, pagesize int, order string) *[]Image {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "i_id desc"
        } else {
            order = "i_" + order
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
            order = "i_id"
        } else {
            order = "i_" + order
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


func (p *ImageManager) GetCount() int {
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

func (p *ImageManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Image {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where i_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "i_id desc"
        } else {
            order = "i_" + order
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
            order = "i_id"
        } else {
            order = "i_" + order
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


func (p *ImageManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where i_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *ImageManager) GetListByApt(apt int64, page int, pagesize int, orderby string) *[]Image {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and i_apt = ?"
		params = append(params, apt)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "i_id desc"
        } else {
            orderby = "i_" + orderby
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
            orderby = "i_id"
        } else {
            orderby = "i_" + orderby
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

func (p *ImageManager) GetCountByApt(apt int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and i_apt = ?"
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

func (p *ImageManager) GetListByAptLevel(apt int64, level int, page int, pagesize int, orderby string) *[]Image {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and i_apt = ?"
		params = append(params, apt)
	}
	if level != 0 {
		query += " and i_level < ?"
		params = append(params, level)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "i_id desc"
        } else {
            orderby = "i_" + orderby
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
            orderby = "i_id"
        } else {
            orderby = "i_" + orderby
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

func (p *ImageManager) GetCountByAptLevel(apt int64, level int) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and i_apt = ?"
		params = append(params, apt)
	}
	if level != 0 {
		query += " and i_level < ?"
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

func (p *ImageManager) GetListByAptParent(apt int64, parent int64, page int, pagesize int, orderby string) *[]Image {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and i_apt = ?"
		params = append(params, apt)
	}
	if parent != 0 {
		query += " and i_parent = ?"
		params = append(params, parent)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "i_id desc"
        } else {
            orderby = "i_" + orderby
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
            orderby = "i_id"
        } else {
            orderby = "i_" + orderby
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

func (p *ImageManager) GetCountByAptParent(apt int64, parent int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and i_apt = ?"
		params = append(params, apt)
	}
	if parent != 0 {
		query += " and i_parent = ?"
		params = append(params, parent)
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

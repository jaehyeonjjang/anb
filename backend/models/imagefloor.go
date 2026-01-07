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

type Imagefloor struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Image                    int64 `json:"Image, int64" form:"image"`
    Name                     string `json:"Name, string" form:"name"`
    Imagename                string `json:"Imagename, string" form:"imagename"`
    Target                   int64 `json:"Target, int64" form:"target"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type ImagefloorManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewImagefloorManager(conn *sql.DB) *ImagefloorManager {
    var item ImagefloorManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "if"
    item.Index = ""

    return &item
}

func (p *ImagefloorManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ImagefloorManager) GetLast(items *[]Imagefloor) *Imagefloor {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *ImagefloorManager) SetIndex(index string) {
    p.Index = index
}

func (p *ImagefloorManager) GetQuery() string {
    ret := ""

    tableName := "imagefloor_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".imagefloor_tb"
    }

    str := "select if_id, if_image, if_name, if_imagename, if_target, if_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *ImagefloorManager) GetQuerySelect() string {
    ret := ""

    tableName := "imagefloor_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".imagefloor_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *ImagefloorManager) Insert(item *Imagefloor) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "imagefloor_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".imagefloor_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (if_id, if_image, if_name, if_imagename, if_target, if_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Image, item.Name, item.Imagename, item.Target, item.Date)
    } else {
        query = "insert into " + tableName + " (if_image, if_name, if_imagename, if_target, if_date) values (?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Image, item.Name, item.Imagename, item.Target, item.Date)
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
func (p *ImagefloorManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "imagefloor_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".imagefloor_tb"
    }
    query := "delete from " + tableName + " where if_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}
func (p *ImagefloorManager) Update(item *Imagefloor) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "imagefloor_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".imagefloor_tb"
    }

	query := "update " + tableName + " set if_image = ?,if_name = ?,if_imagename = ?,if_target = ?,if_date = ? where if_id = ?"
	_, err := p.Conn.Exec(query, item.Image, item.Name, item.Imagename, item.Target, item.Date, item.Id)

    return err
}

func (p *ImagefloorManager) GetIdentity() int64 {
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
    

func (p *ImagefloorManager) ReadRow(rows *sql.Rows) *Imagefloor {
    var item Imagefloor
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Image, &item.Name, &item.Imagename, &item.Target, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *ImagefloorManager) ReadRows(rows *sql.Rows) *[]Imagefloor {
    var items []Imagefloor
    var err error

    for rows.Next() {
        var item Imagefloor
        err = rows.Scan(&item.Id, &item.Image, &item.Name, &item.Imagename, &item.Target, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *ImagefloorManager) Get(id int64) *Imagefloor {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where if_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ImagefloorManager) GetList(page int, pagesize int, order string) *[]Imagefloor {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "if_id desc"
        } else {
            order = "if_" + order
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
            order = "if_id"
        } else {
            order = "if_" + order
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


func (p *ImagefloorManager) GetCount() int {
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

func (p *ImagefloorManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Imagefloor {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where if_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "if_id desc"
        } else {
            order = "if_" + order
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
            order = "if_id"
        } else {
            order = "if_" + order
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


func (p *ImagefloorManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where if_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *ImagefloorManager) GetByImageNameImagename(image int64, name string, imagename string) *Imagefloor {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if image != 0 {
		query += " and if_image = ?"
		params = append(params, image)
	}
	if name != "" {
		query += " and if_name = ?"
		params = append(params, name)
	}
	if imagename != "" {
		query += " and if_imagename = ?"
		params = append(params, imagename)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ImagefloorManager) GetByTarget(target int64) *Imagefloor {
        
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if target != 0 {
		query += " and if_target = ?"
		params = append(params, target)
	}


    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

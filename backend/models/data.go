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

type Data struct {
    Id                       int64 `json:"Id, int64" form:"id"`
    Apt                      int64 `json:"Apt, int64" form:"apt"`
    Image                    int64 `json:"Image, int64" form:"image"`
    Imagetype                int `json:"Imagetype, int" form:"imagetype"`
    User                     int64 `json:"User, int64" form:"user"`
    Type                     int `json:"Type, int" form:"type"`
    X                        float64 `json:"X, float64" form:"x"`
    Y                        float64 `json:"Y, float64" form:"y"`
    Point                    string `json:"Point, string" form:"point"`
    Number                   int `json:"Number, int" form:"number"`
    Group                    int `json:"Group, int" form:"group"`
    Name                     string `json:"Name, string" form:"name"`
    Fault                    string `json:"Fault, string" form:"fault"`
    Content                  string `json:"Content, string" form:"content"`
    Width                    float64 `json:"Width, float64" form:"width"`
    Length                   float64 `json:"Length, float64" form:"length"`
    Count                    string `json:"Count, string" form:"count"`
    Progress                 string `json:"Progress, string" form:"progress"`
    Remark                   string `json:"Remark, string" form:"remark"`
    Imagename                string `json:"Imagename, string" form:"imagename"`
    Filename                 string `json:"Filename, string" form:"filename"`
    Memo                     string `json:"Memo, string" form:"memo"`
    Report                   int `json:"Report, int" form:"report"`
    Usermemo                 string `json:"Usermemo, string" form:"usermemo"`
    Aptmemo                  string `json:"Aptmemo, string" form:"aptmemo"`
    Date                     string `json:"Date, string" form:"date"`
    Extra                    interface{} `form:"extra"`
}

type DataManager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}

func NewDataManager(conn *sql.DB) *DataManager {
    var item DataManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = "d"
    item.Index = ""

    return &item
}

func (p *DataManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *DataManager) GetLast(items *[]Data) *Data {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *DataManager) SetIndex(index string) {
    p.Index = index
}

func (p *DataManager) GetQuery() string {
    ret := ""

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }

    str := "select d_id, d_apt, d_image, d_imagetype, d_user, d_type, d_x, d_y, d_point, d_number, d_group, d_name, d_fault, d_content, d_width, d_length, d_count, d_progress, d_remark, d_imagename, d_filename, d_memo, d_report, d_usermemo, d_aptmemo, d_date from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *DataManager) GetQuerySelect() string {
    ret := ""

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }

    str := "select count(*) from " + tableName + " "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    return ret;
}

func (p *DataManager) Insert(item *Data) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }

    var err error
    var res sql.Result
    query := ""

    if item.Id > 0 {
        query = "insert into " + tableName + " (d_id, d_apt, d_image, d_imagetype, d_user, d_type, d_x, d_y, d_point, d_number, d_group, d_name, d_fault, d_content, d_width, d_length, d_count, d_progress, d_remark, d_imagename, d_filename, d_memo, d_report, d_usermemo, d_aptmemo, d_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Id, item.Apt, item.Image, item.Imagetype, item.User, item.Type, item.X, item.Y, item.Point, item.Number, item.Group, item.Name, item.Fault, item.Content, item.Width, item.Length, item.Count, item.Progress, item.Remark, item.Imagename, item.Filename, item.Memo, item.Report, item.Usermemo, item.Aptmemo, item.Date)
    } else {
        query = "insert into " + tableName + " (d_apt, d_image, d_imagetype, d_user, d_type, d_x, d_y, d_point, d_number, d_group, d_name, d_fault, d_content, d_width, d_length, d_count, d_progress, d_remark, d_imagename, d_filename, d_memo, d_report, d_usermemo, d_aptmemo, d_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Conn.Exec(query, item.Apt, item.Image, item.Imagetype, item.User, item.Type, item.X, item.Y, item.Point, item.Number, item.Group, item.Name, item.Fault, item.Content, item.Width, item.Length, item.Count, item.Progress, item.Remark, item.Imagename, item.Filename, item.Memo, item.Report, item.Usermemo, item.Aptmemo, item.Date)
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
func (p *DataManager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }
    query := "delete from " + tableName + " where d_id = ?"
    _, err := p.Conn.Exec(query, id)

    return err
}

func (p *DataManager) DeleteByImageUser(image int64, user int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }
    var params []interface{}
    query := "delete from " + tableName + " where d_image = ? and d_user = ?"
	params = append(params, image)
	params = append(params, user)

    err := ExecArray(p.Conn, query, params)

    return err
}


func (p *DataManager) DeleteByImage(image int64) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }
    var params []interface{}
    query := "delete from " + tableName + " where d_image = ?"
	params = append(params, image)

    err := ExecArray(p.Conn, query, params)

    return err
}

func (p *DataManager) Update(item *Data) error {
    if p.Conn == nil {
        return errors.New("Connection Error")
    }

    tableName := "data_tb"
    if config.Database == "mssql" || config.Database == "sqlserver" {
        tableName = config.Owner + ".data_tb"
    }

	query := "update " + tableName + " set d_apt = ?,d_image = ?,d_imagetype = ?,d_user = ?,d_type = ?,d_x = ?,d_y = ?,d_point = ?,d_number = ?,d_group = ?,d_name = ?,d_fault = ?,d_content = ?,d_width = ?,d_length = ?,d_count = ?,d_progress = ?,d_remark = ?,d_imagename = ?,d_filename = ?,d_memo = ?,d_report = ?,d_usermemo = ?,d_aptmemo = ?,d_date = ? where d_id = ?"
	_, err := p.Conn.Exec(query, item.Apt, item.Image, item.Imagetype, item.User, item.Type, item.X, item.Y, item.Point, item.Number, item.Group, item.Name, item.Fault, item.Content, item.Width, item.Length, item.Count, item.Progress, item.Remark, item.Imagename, item.Filename, item.Memo, item.Report, item.Usermemo, item.Aptmemo, item.Date, item.Id)

    return err
}

func (p *DataManager) GetIdentity() int64 {
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
    

func (p *DataManager) ReadRow(rows *sql.Rows) *Data {
    var item Data
    var err error

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Apt, &item.Image, &item.Imagetype, &item.User, &item.Type, &item.X, &item.Y, &item.Point, &item.Number, &item.Group, &item.Name, &item.Fault, &item.Content, &item.Width, &item.Length, &item.Count, &item.Progress, &item.Remark, &item.Imagename, &item.Filename, &item.Memo, &item.Report, &item.Usermemo, &item.Aptmemo, &item.Date)
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *DataManager) ReadRows(rows *sql.Rows) *[]Data {
    var items []Data
    var err error

    for rows.Next() {
        var item Data
        err = rows.Scan(&item.Id, &item.Apt, &item.Image, &item.Imagetype, &item.User, &item.Type, &item.X, &item.Y, &item.Point, &item.Number, &item.Group, &item.Name, &item.Fault, &item.Content, &item.Width, &item.Length, &item.Count, &item.Progress, &item.Remark, &item.Imagename, &item.Filename, &item.Memo, &item.Report, &item.Usermemo, &item.Aptmemo, &item.Date)

        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *DataManager) Get(id int64) *Data {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + " where d_id = ?"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *DataManager) GetList(page int, pagesize int, order string) *[]Data {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "d_id desc"
        } else {
            order = "d_" + order
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
            order = "d_id"
        } else {
            order = "d_" + order
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


func (p *DataManager) GetCount() int {
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

func (p *DataManager) GetListInID(ids []int, page int, pagesize int, order string) *[]Data {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + " where d_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

    if page > 0 && pagesize > 0 {
        if order == "" {
            order = "d_id desc"
        } else {
            order = "d_" + order
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
            order = "d_id"
        } else {
            order = "d_" + order
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


func (p *DataManager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + " where d_id in (" + strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ", ", -1), "[]") + ")"

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

func (p *DataManager) GetListByAptImage(apt int64, image int64, page int, pagesize int, orderby string) *[]Data {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and d_apt = ?"
		params = append(params, apt)
	}
	if image != 0 {
		query += " and d_image = ?"
		params = append(params, image)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "d_id desc"
        } else {
            orderby = "d_" + orderby
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
            orderby = "d_id"
        } else {
            orderby = "d_" + orderby
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

func (p *DataManager) GetCountByAptImage(apt int64, image int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and d_apt = ?"
		params = append(params, apt)
	}
	if image != 0 {
		query += " and d_image = ?"
		params = append(params, image)
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

func (p *DataManager) GetListByApt(apt int64, page int, pagesize int, orderby string) *[]Data {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if apt != 0 {
		query += " and d_apt = ?"
		params = append(params, apt)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "d_id desc"
        } else {
            orderby = "d_" + orderby
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
            orderby = "d_id"
        } else {
            orderby = "d_" + orderby
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

func (p *DataManager) GetCountByApt(apt int64) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if apt != 0 {
		query += " and d_apt = ?"
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

func (p *DataManager) GetListByImageNameImagename(image int64, name string, imagename string, page int, pagesize int, orderby string) *[]Data {
        
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + " where 1=1 "
    var params []interface{}

	if image != 0 {
		query += " and d_image = ?"
		params = append(params, image)
	}
	if name != "" {
		query += " and d_name = ?"
		params = append(params, name)
	}
	if imagename != "" {
		query += " and d_imagename = ?"
		params = append(params, imagename)
	}


    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "d_id desc"
        } else {
            orderby = "d_" + orderby
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
            orderby = "d_id"
        } else {
            orderby = "d_" + orderby
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

func (p *DataManager) GetCountByImageNameImagename(image int64, name string, imagename string) int {
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + " where 1=1 "
	if image != 0 {
		query += " and d_image = ?"
		params = append(params, image)
	}
	if name != "" {
		query += " and d_name = ?"
		params = append(params, name)
	}
	if imagename != "" {
		query += " and d_imagename = ?"
		params = append(params, imagename)
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

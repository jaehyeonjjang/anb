package models

import (
    "repair/models/repairarea"
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    "repair/global/config"
    log "repair/global/log"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"

    
)

type Repairarea struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Filename                string `json:"filename"`         
    Length                Double `json:"length"`         
    Standard                Double `json:"standard"`         
    Content                string `json:"content"`         
    Order                int `json:"order"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type RepairareaManager struct {
    Conn    *Connection
    Result  *sql.Result
    Index   string
    Isolation   bool
    SelectQuery  string
    JoinQuery string
    CountQuery   string
    GroupQuery string
    SelectLog bool
    Log bool
}



func (c *Repairarea) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewRepairareaManager(conn *Connection) *RepairareaManager {
    var item RepairareaManager


    if conn == nil {
        item.Conn = NewConnection()
        item.Isolation = false
    } else {
        item.Conn = conn 
        item.Isolation = conn.Isolation
    }

    item.Index = ""
    item.SelectLog = config.Log.Database
    item.Log = config.Log.Database

    return &item
}

func (p *RepairareaManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *RepairareaManager) SetIndex(index string) {
    p.Index = index
}

func (p *RepairareaManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *RepairareaManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *RepairareaManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *RepairareaManager) Query(query string, params ...any) (*sql.Rows, error) {
    if p.Isolation == true {
        query += " for update"
    }

    if p.SelectLog {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Query(query, params...)
}

func (p *RepairareaManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select ra_id, ra_name, ra_filename, ra_length, ra_standard, ra_content, ra_order, ra_apt, ra_date, a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date from repairarea_tb, apt_tb")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    if p.JoinQuery != "" {
        ret.WriteString(", ")
        ret.WriteString(p.JoinQuery)
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and ra_apt = a_id ")
    

    return ret.String()
}

func (p *RepairareaManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from repairarea_tb, apt_tb")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    if p.JoinQuery != "" {
        ret.WriteString(", ")
        ret.WriteString(p.JoinQuery)
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and ra_apt = a_id ")
    

    return ret.String()
}

func (p *RepairareaManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select ra_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from repairarea_tb, apt_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and ra_apt = a_id ")
    


    return ret.String()
}

func (p *RepairareaManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate repairarea_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *RepairareaManager) Insert(item *Repairarea) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into repairarea_tb (ra_id, ra_name, ra_filename, ra_length, ra_standard, ra_content, ra_order, ra_apt, ra_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
        } else {
          query = "insert into repairarea_tb (ra_id, ra_name, ra_filename, ra_length, ra_standard, ra_content, ra_order, ra_apt, ra_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Name, item.Filename, item.Length, item.Standard, item.Content, item.Order, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into repairarea_tb (ra_name, ra_filename, ra_length, ra_standard, ra_content, ra_order, ra_apt, ra_date) values ($1, $2, $3, $4, $5, $6, $7, $8)"
        } else {
          query = "insert into repairarea_tb (ra_name, ra_filename, ra_length, ra_standard, ra_content, ra_order, ra_apt, ra_date) values (?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Name, item.Filename, item.Length, item.Standard, item.Content, item.Order, item.Apt, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
        p.Result = nil
    }

    return err
}

func (p *RepairareaManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from repairarea_tb where ra_id = ")
    if config.Database.Type == config.Postgresql {
        query.WriteString("$1")
    } else {
        query.WriteString("?")
    }
    _, err := p.Exec(query.String(), id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *RepairareaManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from repairarea_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *RepairareaManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
    var params []any
    if initParams != nil {
        params = append(params, initParams...)
    }

    pos := 1

    var query strings.Builder
	query.WriteString(initQuery)

    for _, arg := range args {
        switch v := arg.(type) {        
        case Where:
            item := v

            if strings.Contains(item.Column, "_") {
                query.WriteString(" and ")
            } else {
                query.WriteString(" and ra_")
            }
            query.WriteString(item.Column)

            if item.Compare == "in" {
                query.WriteString(" in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "not in" {
                query.WriteString(" not in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "between" {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(fmt.Sprintf(" between $%v and $%v", pos, pos + 1))
                    pos += 2
                } else {
                    query.WriteString(" between ? and ?")
                }

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" $")
                    query.WriteString(fmt.Sprintf("%v", pos))
                    pos++
                } else {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" ?")
                }
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

            query.WriteString(" and ")
            query.WriteString(item.Query)
        }        
    }

	query.WriteString(postQuery)

    return query.String(), params
}

func (p *RepairareaManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from repairarea_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *RepairareaManager) Update(item *Repairarea) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update repairarea_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" ra_name = $1, ra_filename = $2, ra_length = $3, ra_standard = $4, ra_content = $5, ra_order = $6, ra_apt = $7, ra_date = $8 where ra_id = $9")
    } else {
        query.WriteString(" ra_name = ?, ra_filename = ?, ra_length = ?, ra_standard = ?, ra_content = ?, ra_order = ?, ra_apt = ?, ra_date = ? where ra_id = ?")
    }

	_, err := p.Exec(query.String() , item.Name, item.Filename, item.Length, item.Standard, item.Content, item.Order, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *RepairareaManager) UpdateWhere(columns []repairarea.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update repairarea_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == repairarea.ColumnId {
        initQuery.WriteString("ra_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnName {
        initQuery.WriteString("ra_name = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnFilename {
        initQuery.WriteString("ra_filename = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnLength {
        initQuery.WriteString("ra_length = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnStandard {
        initQuery.WriteString("ra_standard = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnContent {
        initQuery.WriteString("ra_content = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnOrder {
        initQuery.WriteString("ra_order = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnApt {
        initQuery.WriteString("ra_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repairarea.ColumnDate {
        initQuery.WriteString("ra_date = ?")
        initParams = append(initParams, v.Value)
        
        }
    }

    initQuery.WriteString(" where 1=1 ")

    query, params := p.MakeQuery(initQuery.String(), "", initParams, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}


/*

func (p *RepairareaManager) UpdateName(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_name = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) UpdateFilename(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_filename = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) UpdateLength(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_length = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) UpdateStandard(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_standard = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) UpdateContent(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_content = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) UpdateOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_order = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_apt = ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

/*

func (p *RepairareaManager) IncreaseLength(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_length = ra_length + ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) IncreaseStandard(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_standard = ra_standard + ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) IncreaseOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_order = ra_order + ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairareaManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairarea_tb set ra_apt = ra_apt + ? where ra_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *RepairareaManager) GetIdentity() int64 {
    if !p.Conn.IsConnect() {
        return 0
    }

    id, err := (*p.Result).LastInsertId()

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
        return 0
    } else {
        return id
    }
}

func (p *Repairarea) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *RepairareaManager) ReadRow(rows *sql.Rows) *Repairarea {
    var item Repairarea
    var err error

    var _apt Apt
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Filename, &item.Length, &item.Standard, &item.Content, &item.Order, &item.Apt, &item.Date, &_apt.Id, &_apt.Name, &_apt.Completeyear, &_apt.Flatcount, &_apt.Type, &_apt.Floor, &_apt.Familycount, &_apt.Familycount1, &_apt.Familycount2, &_apt.Familycount3, &_apt.Tel, &_apt.Fax, &_apt.Email, &_apt.Personalemail, &_apt.Personalname, &_apt.Personalhp, &_apt.Zip, &_apt.Address, &_apt.Address2, &_apt.Contracttype, &_apt.Contractprice, &_apt.Testdate, &_apt.Nexttestdate, &_apt.Repair, &_apt.Safety, &_apt.Fault, &_apt.Contractdate, &_apt.Contractduration, &_apt.Invoice, &_apt.Depositdate, &_apt.Fmsloginid, &_apt.Fmspasswd, &_apt.Facilitydivision, &_apt.Facilitycategory, &_apt.Position, &_apt.Area, &_apt.Groundfloor, &_apt.Undergroundfloor, &_apt.Useapproval, &_apt.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
        

    } else {
        return nil
    }

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
        return nil
    } else {

        item.InitExtra()
        _apt.InitExtra()
        item.AddExtra("apt",  _apt)

        return &item
    }
}

func (p *RepairareaManager) ReadRows(rows *sql.Rows) []Repairarea {
    items := make([]Repairarea, 0)

    for rows.Next() {
        var item Repairarea
        var _apt Apt
            
    
        err := rows.Scan(&item.Id, &item.Name, &item.Filename, &item.Length, &item.Standard, &item.Content, &item.Order, &item.Apt, &item.Date, &_apt.Id, &_apt.Name, &_apt.Completeyear, &_apt.Flatcount, &_apt.Type, &_apt.Floor, &_apt.Familycount, &_apt.Familycount1, &_apt.Familycount2, &_apt.Familycount3, &_apt.Tel, &_apt.Fax, &_apt.Email, &_apt.Personalemail, &_apt.Personalname, &_apt.Personalhp, &_apt.Zip, &_apt.Address, &_apt.Address2, &_apt.Contracttype, &_apt.Contractprice, &_apt.Testdate, &_apt.Nexttestdate, &_apt.Repair, &_apt.Safety, &_apt.Fault, &_apt.Contractdate, &_apt.Contractduration, &_apt.Invoice, &_apt.Depositdate, &_apt.Fmsloginid, &_apt.Fmspasswd, &_apt.Facilitydivision, &_apt.Facilitycategory, &_apt.Position, &_apt.Area, &_apt.Groundfloor, &_apt.Undergroundfloor, &_apt.Useapproval, &_apt.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
		
        
        item.InitExtra()        
        _apt.InitExtra()
        item.AddExtra("apt",  _apt)

        items = append(items, item)
    }


     return items
}

func (p *RepairareaManager) Get(id int64) *Repairarea {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and ra_id = ?")

    
    query.WriteString(" and ra_apt = a_id ")
    
    
    rows, err := p.Query(query.String(), id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *RepairareaManager) GetWhere(args []any) *Repairarea {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *RepairareaManager) Count(args []any) int {
    if !p.Conn.IsConnect() {
        return 0
    }

    query, params := p.MakeQuery(p.GetQuerySelect(), p.GroupQuery, nil, args)
    rows, err := p.Query(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
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

func (p *RepairareaManager) FindAll() []Repairarea {
    return p.Find(nil)
}

func (p *RepairareaManager) Find(args []any) []Repairarea {
    if !p.Conn.IsConnect() {
        items := make([]Repairarea, 0)
        return items
    }

    var params []any
    baseQuery := p.GetQuery()

    var query strings.Builder

    page := 0
    pagesize := 0
    orderby := ""

    pos := 1
    
    for _, arg := range args {
        switch v := arg.(type) {
        case PagingType:
            item := v
            page = item.Page
            pagesize = item.Pagesize            
        case OrderingType:
            item := v
            orderby = item.Order
        case LimitType:
            item := v
            page = 1
            pagesize = item.Limit
        case OptionType:
            item := v
            if item.Limit > 0 {
                page = 1
                pagesize = item.Limit
            } else {
                page = item.Page
                pagesize = item.Pagesize                
            }
            orderby = item.Order
        case Where:
            item := v

            if strings.Contains(item.Column, "_") {
                query.WriteString(" and ")
            } else {
                query.WriteString(" and ra_")
            }
            query.WriteString(item.Column)
            
            if item.Compare == "in" {
                query.WriteString(" in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "not in" {
                query.WriteString(" not in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "between" {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(fmt.Sprintf(" between $%v and $%v", pos, pos + 1))
                    pos += 2
                } else {
                    query.WriteString(" between ? and ?")
                }

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" $")
                    query.WriteString(fmt.Sprintf("%v", pos))
                    pos++
                } else {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" ?")
                }
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

            query.WriteString(" and ")
            query.WriteString(item.Query)
        case Base:
             item := v

             baseQuery = item.Query
        }
    }

    query.WriteString(p.GroupQuery)
    
    startpage := (page - 1) * pagesize
    
    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "ra_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "ra_" + orderby
                }
            }
            
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
        if config.Database.Type == config.Postgresql {
            query.WriteString(fmt.Sprintf(" limit $%v offset $%v", pos, pos + 1))
            params = append(params, pagesize)
            params = append(params, startpage)
        } else if config.Database.Type == config.Mysql {
            query.WriteString(" limit ? offset ?")
            params = append(params, pagesize)
            params = append(params, startpage)
        } else if config.Database.Type == config.Sqlserver {
            query.WriteString("OFFSET ? ROWS FETCH NEXT ? ROWS ONLY")
            params = append(params, startpage)
            params = append(params, pagesize)
        }
    } else {
        if orderby == "" {
            orderby = "ra_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "ra_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(baseQuery + query.String(), params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
        items := make([]Repairarea, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *RepairareaManager) GroupBy(name string, args []any) []Groupby {
    if !p.Conn.IsConnect() {
        var items []Groupby
        return items
    }

    var params []any
    baseQuery := p.GetQueryGroup(name)
    var query strings.Builder
    pos := 1

    for _, arg := range args {
        switch v := arg.(type) {
        case Where:
            item := v

            if strings.Contains(item.Column, "_") {
                query.WriteString(" and ")
            } else {
                query.WriteString(" and ra_")
            }
            query.WriteString(item.Column)
            
            if item.Compare == "in" {
                query.WriteString(" in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "not in" {
                query.WriteString(" not in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "between" {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(fmt.Sprintf(" between $%v and $%v", pos, pos + 1))
                    pos += 2
                } else {
                    query.WriteString(" between ? and ?")
                }

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" $")
                    query.WriteString(fmt.Sprintf("%v", pos))
                    pos++
                } else {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" ?")
                }
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

            query.WriteString(" and ")
            query.WriteString(item.Query)
        case Base:
             item := v

             baseQuery = item.Query
        }
    }
    
    query.WriteString(" group by ra_")
    query.WriteString(name)

    rows, err := p.Query(baseQuery + query.String(), params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
        var items []Groupby
        return items
    }

    defer rows.Close()

    var items []Groupby

    for rows.Next() {
        var item Groupby
        err := rows.Scan(&item.Value, &item.Count)
        if err != nil {
           if p.Log {
                log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        items = append(items, item)
    }

    return items
}



func (p *RepairareaManager) MakeMap(items []Repairarea) map[int64]Repairarea {
     ret := make(map[int64]Repairarea)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *RepairareaManager) FindToMap(args []any) map[int64]Repairarea {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *RepairareaManager) FindAllToMap() map[int64]Repairarea {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *RepairareaManager) MakeNameMap(items []Repairarea) map[string]Repairarea {
     ret := make(map[string]Repairarea)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *RepairareaManager) FindToNameMap(args []any) map[string]Repairarea {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *RepairareaManager) FindAllToNameMap() map[string]Repairarea {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}

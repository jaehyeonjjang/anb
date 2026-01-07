package models

import (
    "repair/models/periodictechnician"
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

type Periodictechnician struct {
            
    Id                int64 `json:"id"`         
    Type                int `json:"type"`         
    Part                string `json:"part"`         
    Signupstartdate                string `json:"signupstartdate"`         
    Signupenddate                string `json:"signupenddate"`         
    Remark                string `json:"remark"`         
    Order                int `json:"order"`         
    Technician                int64 `json:"technician"`         
    Periodic                int64 `json:"periodic"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodictechnicianManager struct {
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



func (c *Periodictechnician) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodictechnicianManager(conn *Connection) *PeriodictechnicianManager {
    var item PeriodictechnicianManager


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

func (p *PeriodictechnicianManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodictechnicianManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodictechnicianManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodictechnicianManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodictechnicianManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodictechnicianManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PeriodictechnicianManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select dt_id, dt_type, dt_part, dt_signupstartdate, dt_signupenddate, dt_remark, dt_order, dt_technician, dt_periodic, dt_date, te_id, te_name, te_grade, te_stamp, te_date from periodictechnician_tb, technician_tb")

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
    
    ret.WriteString(" and dt_technician = te_id ")
    

    return ret.String()
}

func (p *PeriodictechnicianManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodictechnician_tb, technician_tb")

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
    
    ret.WriteString(" and dt_technician = te_id ")
    

    return ret.String()
}

func (p *PeriodictechnicianManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select dt_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodictechnician_tb, technician_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and dt_technician = te_id ")
    


    return ret.String()
}

func (p *PeriodictechnicianManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodictechnician_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodictechnicianManager) Insert(item *Periodictechnician) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
	
    if item.Signupstartdate == "" {
       item.Signupstartdate = "1000-01-01"
    }
	
    if item.Signupenddate == "" {
       item.Signupenddate = "1000-01-01"
    }
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodictechnician_tb (dt_id, dt_type, dt_part, dt_signupstartdate, dt_signupenddate, dt_remark, dt_order, dt_technician, dt_periodic, dt_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
        } else {
          query = "insert into periodictechnician_tb (dt_id, dt_type, dt_part, dt_signupstartdate, dt_signupenddate, dt_remark, dt_order, dt_technician, dt_periodic, dt_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Type, item.Part, item.Signupstartdate, item.Signupenddate, item.Remark, item.Order, item.Technician, item.Periodic, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodictechnician_tb (dt_type, dt_part, dt_signupstartdate, dt_signupenddate, dt_remark, dt_order, dt_technician, dt_periodic, dt_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
        } else {
          query = "insert into periodictechnician_tb (dt_type, dt_part, dt_signupstartdate, dt_signupenddate, dt_remark, dt_order, dt_technician, dt_periodic, dt_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Type, item.Part, item.Signupstartdate, item.Signupenddate, item.Remark, item.Order, item.Technician, item.Periodic, item.Date)
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

func (p *PeriodictechnicianManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodictechnician_tb where dt_id = ")
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

func (p *PeriodictechnicianManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodictechnician_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodictechnicianManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and dt_")
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

func (p *PeriodictechnicianManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodictechnician_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodictechnicianManager) Update(item *Periodictechnician) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
    if item.Signupstartdate == "" {
       item.Signupstartdate = "1000-01-01"
    }
	
    if item.Signupenddate == "" {
       item.Signupenddate = "1000-01-01"
    }
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodictechnician_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" dt_type = $1, dt_part = $2, dt_signupstartdate = $3, dt_signupenddate = $4, dt_remark = $5, dt_order = $6, dt_technician = $7, dt_periodic = $8, dt_date = $9 where dt_id = $10")
    } else {
        query.WriteString(" dt_type = ?, dt_part = ?, dt_signupstartdate = ?, dt_signupenddate = ?, dt_remark = ?, dt_order = ?, dt_technician = ?, dt_periodic = ?, dt_date = ? where dt_id = ?")
    }

	_, err := p.Exec(query.String() , item.Type, item.Part, item.Signupstartdate, item.Signupenddate, item.Remark, item.Order, item.Technician, item.Periodic, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodictechnicianManager) UpdateWhere(columns []periodictechnician.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodictechnician_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodictechnician.ColumnId {
        initQuery.WriteString("dt_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnType {
        initQuery.WriteString("dt_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnPart {
        initQuery.WriteString("dt_part = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnSignupstartdate {
        initQuery.WriteString("dt_signupstartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnSignupenddate {
        initQuery.WriteString("dt_signupenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnRemark {
        initQuery.WriteString("dt_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnOrder {
        initQuery.WriteString("dt_order = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnTechnician {
        initQuery.WriteString("dt_technician = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnPeriodic {
        initQuery.WriteString("dt_periodic = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodictechnician.ColumnDate {
        initQuery.WriteString("dt_date = ?")
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

func (p *PeriodictechnicianManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_type = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdatePart(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_part = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdateSignupstartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_signupstartdate = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdateSignupenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_signupenddate = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_remark = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdateOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_order = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdateTechnician(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_technician = ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) UpdatePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_periodic = ? where dt_id = ?"
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

func (p *PeriodictechnicianManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_type = dt_type + ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) IncreaseOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_order = dt_order + ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) IncreaseTechnician(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_technician = dt_technician + ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodictechnicianManager) IncreasePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodictechnician_tb set dt_periodic = dt_periodic + ? where dt_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodictechnicianManager) GetIdentity() int64 {
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

func (p *Periodictechnician) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodictechnicianManager) ReadRow(rows *sql.Rows) *Periodictechnician {
    var item Periodictechnician
    var err error

    var _technician Technician
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Type, &item.Part, &item.Signupstartdate, &item.Signupenddate, &item.Remark, &item.Order, &item.Technician, &item.Periodic, &item.Date, &_technician.Id, &_technician.Name, &_technician.Grade, &_technician.Stamp, &_technician.Date)
        
        
        
        
        
        
        if item.Signupstartdate == "0000-00-00" || item.Signupstartdate == "1000-01-01" || item.Signupstartdate == "9999-01-01" {
            item.Signupstartdate = ""
        }
        
        if item.Signupenddate == "0000-00-00" || item.Signupenddate == "1000-01-01" || item.Signupenddate == "9999-01-01" {
            item.Signupenddate = ""
        }
        
        
        
        
        
        
        
        
        
        
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
        _technician.InitExtra()
        item.AddExtra("technician",  _technician)

        return &item
    }
}

func (p *PeriodictechnicianManager) ReadRows(rows *sql.Rows) []Periodictechnician {
    items := make([]Periodictechnician, 0)

    for rows.Next() {
        var item Periodictechnician
        var _technician Technician
            
    
        err := rows.Scan(&item.Id, &item.Type, &item.Part, &item.Signupstartdate, &item.Signupenddate, &item.Remark, &item.Order, &item.Technician, &item.Periodic, &item.Date, &_technician.Id, &_technician.Name, &_technician.Grade, &_technician.Stamp, &_technician.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		
        
		if item.Signupstartdate == "0000-00-00" || item.Signupstartdate == "1000-01-01" || item.Signupstartdate == "9999-01-01" {
            item.Signupstartdate = ""
        }
        
		if item.Signupenddate == "0000-00-00" || item.Signupenddate == "1000-01-01" || item.Signupenddate == "9999-01-01" {
            item.Signupenddate = ""
        }
        
		
        
		
        
		
        
		
        
		
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
		
        
        item.InitExtra()        
        _technician.InitExtra()
        item.AddExtra("technician",  _technician)

        items = append(items, item)
    }


     return items
}

func (p *PeriodictechnicianManager) Get(id int64) *Periodictechnician {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and dt_id = ?")

    
    query.WriteString(" and dt_technician = te_id ")
    
    
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

func (p *PeriodictechnicianManager) GetWhere(args []any) *Periodictechnician {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodictechnicianManager) Count(args []any) int {
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

func (p *PeriodictechnicianManager) FindAll() []Periodictechnician {
    return p.Find(nil)
}

func (p *PeriodictechnicianManager) Find(args []any) []Periodictechnician {
    if !p.Conn.IsConnect() {
        items := make([]Periodictechnician, 0)
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
                query.WriteString(" and dt_")
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
            orderby = "dt_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "dt_" + orderby
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
            orderby = "dt_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "dt_" + orderby
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
        items := make([]Periodictechnician, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *PeriodictechnicianManager) CountByPeriodic(periodic int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if periodic != 0 { 
        rets = append(rets, Where{Column:"periodic", Value:periodic, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *PeriodictechnicianManager) FindByPeriodic(periodic int64, args ...any) []Periodictechnician {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if periodic != 0 { 
        rets = append(rets, Where{Column:"periodic", Value:periodic, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *PeriodictechnicianManager) DeleteByPeriodic(periodic int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodictechnician_tb where dt_periodic = ?"
    _, err := p.Exec(query, periodic)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}




func (p *PeriodictechnicianManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and dt_")
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
    
    query.WriteString(" group by dt_")
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



func (p *PeriodictechnicianManager) MakeMap(items []Periodictechnician) map[int64]Periodictechnician {
     ret := make(map[int64]Periodictechnician)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodictechnicianManager) FindToMap(args []any) map[int64]Periodictechnician {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodictechnicianManager) FindAllToMap() map[int64]Periodictechnician {
     items := p.Find(nil)
     return p.MakeMap(items)
}



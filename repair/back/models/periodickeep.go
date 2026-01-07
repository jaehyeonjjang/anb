package models

import (
    "repair/models/periodickeep"
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

type Periodickeep struct {
            
    Id                int64 `json:"id"`         
    Status1                int `json:"status1"`         
    Status2                int `json:"status2"`         
    Status3                int `json:"status3"`         
    Status4                int `json:"status4"`         
    Status5                int `json:"status5"`         
    Status6                int `json:"status6"`         
    Content1                string `json:"content1"`         
    Content2                string `json:"content2"`         
    Content3                string `json:"content3"`         
    Content4                string `json:"content4"`         
    Content5                string `json:"content5"`         
    Content6                string `json:"content6"`         
    Remark1                string `json:"remark1"`         
    Remark2                string `json:"remark2"`         
    Remark3                string `json:"remark3"`         
    Remark4                string `json:"remark4"`         
    Remark5                string `json:"remark5"`         
    Remark6                string `json:"remark6"`         
    Periodic                int64 `json:"periodic"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodickeepManager struct {
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



func (c *Periodickeep) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodickeepManager(conn *Connection) *PeriodickeepManager {
    var item PeriodickeepManager


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

func (p *PeriodickeepManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodickeepManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodickeepManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodickeepManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodickeepManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodickeepManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PeriodickeepManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select pk_id, pk_status1, pk_status2, pk_status3, pk_status4, pk_status5, pk_status6, pk_content1, pk_content2, pk_content3, pk_content4, pk_content5, pk_content6, pk_remark1, pk_remark2, pk_remark3, pk_remark4, pk_remark5, pk_remark6, pk_periodic, pk_date from periodickeep_tb")

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
    

    return ret.String()
}

func (p *PeriodickeepManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodickeep_tb")

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
    

    return ret.String()
}

func (p *PeriodickeepManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select pk_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodickeep_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *PeriodickeepManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodickeep_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodickeepManager) Insert(item *Periodickeep) error {
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
          query = "insert into periodickeep_tb (pk_id, pk_status1, pk_status2, pk_status3, pk_status4, pk_status5, pk_status6, pk_content1, pk_content2, pk_content3, pk_content4, pk_content5, pk_content6, pk_remark1, pk_remark2, pk_remark3, pk_remark4, pk_remark5, pk_remark6, pk_periodic, pk_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)"
        } else {
          query = "insert into periodickeep_tb (pk_id, pk_status1, pk_status2, pk_status3, pk_status4, pk_status5, pk_status6, pk_content1, pk_content2, pk_content3, pk_content4, pk_content5, pk_content6, pk_remark1, pk_remark2, pk_remark3, pk_remark4, pk_remark5, pk_remark6, pk_periodic, pk_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Remark1, item.Remark2, item.Remark3, item.Remark4, item.Remark5, item.Remark6, item.Periodic, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodickeep_tb (pk_status1, pk_status2, pk_status3, pk_status4, pk_status5, pk_status6, pk_content1, pk_content2, pk_content3, pk_content4, pk_content5, pk_content6, pk_remark1, pk_remark2, pk_remark3, pk_remark4, pk_remark5, pk_remark6, pk_periodic, pk_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)"
        } else {
          query = "insert into periodickeep_tb (pk_status1, pk_status2, pk_status3, pk_status4, pk_status5, pk_status6, pk_content1, pk_content2, pk_content3, pk_content4, pk_content5, pk_content6, pk_remark1, pk_remark2, pk_remark3, pk_remark4, pk_remark5, pk_remark6, pk_periodic, pk_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Remark1, item.Remark2, item.Remark3, item.Remark4, item.Remark5, item.Remark6, item.Periodic, item.Date)
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

func (p *PeriodickeepManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodickeep_tb where pk_id = ")
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

func (p *PeriodickeepManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodickeep_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodickeepManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and pk_")
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

func (p *PeriodickeepManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodickeep_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodickeepManager) Update(item *Periodickeep) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodickeep_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" pk_status1 = $1, pk_status2 = $2, pk_status3 = $3, pk_status4 = $4, pk_status5 = $5, pk_status6 = $6, pk_content1 = $7, pk_content2 = $8, pk_content3 = $9, pk_content4 = $10, pk_content5 = $11, pk_content6 = $12, pk_remark1 = $13, pk_remark2 = $14, pk_remark3 = $15, pk_remark4 = $16, pk_remark5 = $17, pk_remark6 = $18, pk_periodic = $19, pk_date = $20 where pk_id = $21")
    } else {
        query.WriteString(" pk_status1 = ?, pk_status2 = ?, pk_status3 = ?, pk_status4 = ?, pk_status5 = ?, pk_status6 = ?, pk_content1 = ?, pk_content2 = ?, pk_content3 = ?, pk_content4 = ?, pk_content5 = ?, pk_content6 = ?, pk_remark1 = ?, pk_remark2 = ?, pk_remark3 = ?, pk_remark4 = ?, pk_remark5 = ?, pk_remark6 = ?, pk_periodic = ?, pk_date = ? where pk_id = ?")
    }

	_, err := p.Exec(query.String() , item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Remark1, item.Remark2, item.Remark3, item.Remark4, item.Remark5, item.Remark6, item.Periodic, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodickeepManager) UpdateWhere(columns []periodickeep.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodickeep_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodickeep.ColumnId {
        initQuery.WriteString("pk_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnStatus1 {
        initQuery.WriteString("pk_status1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnStatus2 {
        initQuery.WriteString("pk_status2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnStatus3 {
        initQuery.WriteString("pk_status3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnStatus4 {
        initQuery.WriteString("pk_status4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnStatus5 {
        initQuery.WriteString("pk_status5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnStatus6 {
        initQuery.WriteString("pk_status6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnContent1 {
        initQuery.WriteString("pk_content1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnContent2 {
        initQuery.WriteString("pk_content2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnContent3 {
        initQuery.WriteString("pk_content3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnContent4 {
        initQuery.WriteString("pk_content4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnContent5 {
        initQuery.WriteString("pk_content5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnContent6 {
        initQuery.WriteString("pk_content6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnRemark1 {
        initQuery.WriteString("pk_remark1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnRemark2 {
        initQuery.WriteString("pk_remark2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnRemark3 {
        initQuery.WriteString("pk_remark3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnRemark4 {
        initQuery.WriteString("pk_remark4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnRemark5 {
        initQuery.WriteString("pk_remark5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnRemark6 {
        initQuery.WriteString("pk_remark6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnPeriodic {
        initQuery.WriteString("pk_periodic = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodickeep.ColumnDate {
        initQuery.WriteString("pk_date = ?")
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

func (p *PeriodickeepManager) UpdateStatus1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status1 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateStatus2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status2 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateStatus3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status3 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateStatus4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status4 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateStatus5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status5 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateStatus6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status6 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateContent1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_content1 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateContent2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_content2 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateContent3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_content3 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateContent4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_content4 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateContent5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_content5 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateContent6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_content6 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateRemark1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_remark1 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateRemark2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_remark2 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateRemark3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_remark3 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateRemark4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_remark4 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateRemark5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_remark5 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdateRemark6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_remark6 = ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) UpdatePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_periodic = ? where pk_id = ?"
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

func (p *PeriodickeepManager) IncreaseStatus1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status1 = pk_status1 + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) IncreaseStatus2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status2 = pk_status2 + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) IncreaseStatus3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status3 = pk_status3 + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) IncreaseStatus4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status4 = pk_status4 + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) IncreaseStatus5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status5 = pk_status5 + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) IncreaseStatus6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_status6 = pk_status6 + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodickeepManager) IncreasePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodickeep_tb set pk_periodic = pk_periodic + ? where pk_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodickeepManager) GetIdentity() int64 {
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

func (p *Periodickeep) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodickeepManager) ReadRow(rows *sql.Rows) *Periodickeep {
    var item Periodickeep
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Status1, &item.Status2, &item.Status3, &item.Status4, &item.Status5, &item.Status6, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Content5, &item.Content6, &item.Remark1, &item.Remark2, &item.Remark3, &item.Remark4, &item.Remark5, &item.Remark6, &item.Periodic, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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
        
        return &item
    }
}

func (p *PeriodickeepManager) ReadRows(rows *sql.Rows) []Periodickeep {
    items := make([]Periodickeep, 0)

    for rows.Next() {
        var item Periodickeep
        
    
        err := rows.Scan(&item.Id, &item.Status1, &item.Status2, &item.Status3, &item.Status4, &item.Status5, &item.Status6, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Content5, &item.Content6, &item.Remark1, &item.Remark2, &item.Remark3, &item.Remark4, &item.Remark5, &item.Remark6, &item.Periodic, &item.Date)
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
        
        items = append(items, item)
    }


     return items
}

func (p *PeriodickeepManager) Get(id int64) *Periodickeep {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and pk_id = ?")

    
    
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

func (p *PeriodickeepManager) GetWhere(args []any) *Periodickeep {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodickeepManager) Count(args []any) int {
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

func (p *PeriodickeepManager) FindAll() []Periodickeep {
    return p.Find(nil)
}

func (p *PeriodickeepManager) Find(args []any) []Periodickeep {
    if !p.Conn.IsConnect() {
        items := make([]Periodickeep, 0)
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
                query.WriteString(" and pk_")
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
            orderby = "pk_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "pk_" + orderby
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
            orderby = "pk_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "pk_" + orderby
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
        items := make([]Periodickeep, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *PeriodickeepManager) GetByPeriodic(periodic int64, args ...any) *Periodickeep {
    rets := make([]any, 0)
    rets = append(rets, args...)
    if periodic != 0 {
        rets = append(rets, Where{Column:"periodic", Value:periodic, Compare:"="})        
    }
    
    items := p.Find(rets)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}




func (p *PeriodickeepManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and pk_")
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
    
    query.WriteString(" group by pk_")
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



func (p *PeriodickeepManager) MakeMap(items []Periodickeep) map[int64]Periodickeep {
     ret := make(map[int64]Periodickeep)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodickeepManager) FindToMap(args []any) map[int64]Periodickeep {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodickeepManager) FindAllToMap() map[int64]Periodickeep {
     items := p.Find(nil)
     return p.MakeMap(items)
}



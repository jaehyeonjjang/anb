package models

import (
    "repair/models/periodicouterwall"
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

type Periodicouterwall struct {
            
    Id                int64 `json:"id"`         
    Result1                string `json:"result1"`         
    Result2                string `json:"result2"`         
    Result3                string `json:"result3"`         
    Result4                string `json:"result4"`         
    Result5                string `json:"result5"`         
    Result6                string `json:"result6"`         
    Status1                string `json:"status1"`         
    Status2                string `json:"status2"`         
    Status3                string `json:"status3"`         
    Status4                string `json:"status4"`         
    Status5                string `json:"status5"`         
    Status6                string `json:"status6"`         
    Position1                string `json:"position1"`         
    Position2                string `json:"position2"`         
    Position3                string `json:"position3"`         
    Position4                string `json:"position4"`         
    Position5                string `json:"position5"`         
    Position6                string `json:"position6"`         
    Content                string `json:"content"`         
    Periodic                int64 `json:"periodic"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodicouterwallManager struct {
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



func (c *Periodicouterwall) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodicouterwallManager(conn *Connection) *PeriodicouterwallManager {
    var item PeriodicouterwallManager


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

func (p *PeriodicouterwallManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodicouterwallManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodicouterwallManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodicouterwallManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodicouterwallManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodicouterwallManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PeriodicouterwallManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select po_id, po_result1, po_result2, po_result3, po_result4, po_result5, po_result6, po_status1, po_status2, po_status3, po_status4, po_status5, po_status6, po_position1, po_position2, po_position3, po_position4, po_position5, po_position6, po_content, po_periodic, po_date from periodicouterwall_tb")

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

func (p *PeriodicouterwallManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodicouterwall_tb")

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

func (p *PeriodicouterwallManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select po_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodicouterwall_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *PeriodicouterwallManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodicouterwall_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodicouterwallManager) Insert(item *Periodicouterwall) error {
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
          query = "insert into periodicouterwall_tb (po_id, po_result1, po_result2, po_result3, po_result4, po_result5, po_result6, po_status1, po_status2, po_status3, po_status4, po_status5, po_status6, po_position1, po_position2, po_position3, po_position4, po_position5, po_position6, po_content, po_periodic, po_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)"
        } else {
          query = "insert into periodicouterwall_tb (po_id, po_result1, po_result2, po_result3, po_result4, po_result5, po_result6, po_status1, po_status2, po_status3, po_status4, po_status5, po_status6, po_position1, po_position2, po_position3, po_position4, po_position5, po_position6, po_content, po_periodic, po_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Result6, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Position1, item.Position2, item.Position3, item.Position4, item.Position5, item.Position6, item.Content, item.Periodic, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodicouterwall_tb (po_result1, po_result2, po_result3, po_result4, po_result5, po_result6, po_status1, po_status2, po_status3, po_status4, po_status5, po_status6, po_position1, po_position2, po_position3, po_position4, po_position5, po_position6, po_content, po_periodic, po_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)"
        } else {
          query = "insert into periodicouterwall_tb (po_result1, po_result2, po_result3, po_result4, po_result5, po_result6, po_status1, po_status2, po_status3, po_status4, po_status5, po_status6, po_position1, po_position2, po_position3, po_position4, po_position5, po_position6, po_content, po_periodic, po_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Result6, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Position1, item.Position2, item.Position3, item.Position4, item.Position5, item.Position6, item.Content, item.Periodic, item.Date)
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

func (p *PeriodicouterwallManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodicouterwall_tb where po_id = ")
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

func (p *PeriodicouterwallManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodicouterwall_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodicouterwallManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and po_")
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

func (p *PeriodicouterwallManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodicouterwall_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodicouterwallManager) Update(item *Periodicouterwall) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodicouterwall_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" po_result1 = $1, po_result2 = $2, po_result3 = $3, po_result4 = $4, po_result5 = $5, po_result6 = $6, po_status1 = $7, po_status2 = $8, po_status3 = $9, po_status4 = $10, po_status5 = $11, po_status6 = $12, po_position1 = $13, po_position2 = $14, po_position3 = $15, po_position4 = $16, po_position5 = $17, po_position6 = $18, po_content = $19, po_periodic = $20, po_date = $21 where po_id = $22")
    } else {
        query.WriteString(" po_result1 = ?, po_result2 = ?, po_result3 = ?, po_result4 = ?, po_result5 = ?, po_result6 = ?, po_status1 = ?, po_status2 = ?, po_status3 = ?, po_status4 = ?, po_status5 = ?, po_status6 = ?, po_position1 = ?, po_position2 = ?, po_position3 = ?, po_position4 = ?, po_position5 = ?, po_position6 = ?, po_content = ?, po_periodic = ?, po_date = ? where po_id = ?")
    }

	_, err := p.Exec(query.String() , item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Result6, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Position1, item.Position2, item.Position3, item.Position4, item.Position5, item.Position6, item.Content, item.Periodic, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodicouterwallManager) UpdateWhere(columns []periodicouterwall.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodicouterwall_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodicouterwall.ColumnId {
        initQuery.WriteString("po_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnResult1 {
        initQuery.WriteString("po_result1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnResult2 {
        initQuery.WriteString("po_result2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnResult3 {
        initQuery.WriteString("po_result3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnResult4 {
        initQuery.WriteString("po_result4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnResult5 {
        initQuery.WriteString("po_result5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnResult6 {
        initQuery.WriteString("po_result6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnStatus1 {
        initQuery.WriteString("po_status1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnStatus2 {
        initQuery.WriteString("po_status2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnStatus3 {
        initQuery.WriteString("po_status3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnStatus4 {
        initQuery.WriteString("po_status4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnStatus5 {
        initQuery.WriteString("po_status5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnStatus6 {
        initQuery.WriteString("po_status6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPosition1 {
        initQuery.WriteString("po_position1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPosition2 {
        initQuery.WriteString("po_position2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPosition3 {
        initQuery.WriteString("po_position3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPosition4 {
        initQuery.WriteString("po_position4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPosition5 {
        initQuery.WriteString("po_position5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPosition6 {
        initQuery.WriteString("po_position6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnContent {
        initQuery.WriteString("po_content = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnPeriodic {
        initQuery.WriteString("po_periodic = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicouterwall.ColumnDate {
        initQuery.WriteString("po_date = ?")
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

func (p *PeriodicouterwallManager) UpdateResult1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_result1 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateResult2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_result2 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateResult3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_result3 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateResult4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_result4 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateResult5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_result5 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateResult6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_result6 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateStatus1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_status1 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateStatus2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_status2 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateStatus3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_status3 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateStatus4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_status4 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateStatus5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_status5 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateStatus6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_status6 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePosition1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_position1 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePosition2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_position2 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePosition3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_position3 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePosition4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_position4 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePosition5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_position5 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePosition6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_position6 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdateContent(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_content = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicouterwallManager) UpdatePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_periodic = ? where po_id = ?"
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

func (p *PeriodicouterwallManager) IncreasePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicouterwall_tb set po_periodic = po_periodic + ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodicouterwallManager) GetIdentity() int64 {
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

func (p *Periodicouterwall) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodicouterwallManager) ReadRow(rows *sql.Rows) *Periodicouterwall {
    var item Periodicouterwall
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Result1, &item.Result2, &item.Result3, &item.Result4, &item.Result5, &item.Result6, &item.Status1, &item.Status2, &item.Status3, &item.Status4, &item.Status5, &item.Status6, &item.Position1, &item.Position2, &item.Position3, &item.Position4, &item.Position5, &item.Position6, &item.Content, &item.Periodic, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *PeriodicouterwallManager) ReadRows(rows *sql.Rows) []Periodicouterwall {
    items := make([]Periodicouterwall, 0)

    for rows.Next() {
        var item Periodicouterwall
        
    
        err := rows.Scan(&item.Id, &item.Result1, &item.Result2, &item.Result3, &item.Result4, &item.Result5, &item.Result6, &item.Status1, &item.Status2, &item.Status3, &item.Status4, &item.Status5, &item.Status6, &item.Position1, &item.Position2, &item.Position3, &item.Position4, &item.Position5, &item.Position6, &item.Content, &item.Periodic, &item.Date)
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

func (p *PeriodicouterwallManager) Get(id int64) *Periodicouterwall {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and po_id = ?")

    
    
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

func (p *PeriodicouterwallManager) GetWhere(args []any) *Periodicouterwall {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodicouterwallManager) Count(args []any) int {
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

func (p *PeriodicouterwallManager) FindAll() []Periodicouterwall {
    return p.Find(nil)
}

func (p *PeriodicouterwallManager) Find(args []any) []Periodicouterwall {
    if !p.Conn.IsConnect() {
        items := make([]Periodicouterwall, 0)
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
                query.WriteString(" and po_")
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
            orderby = "po_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "po_" + orderby
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
            orderby = "po_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "po_" + orderby
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
        items := make([]Periodicouterwall, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *PeriodicouterwallManager) GetByPeriodic(periodic int64, args ...any) *Periodicouterwall {
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




func (p *PeriodicouterwallManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and po_")
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
    
    query.WriteString(" group by po_")
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



func (p *PeriodicouterwallManager) MakeMap(items []Periodicouterwall) map[int64]Periodicouterwall {
     ret := make(map[int64]Periodicouterwall)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodicouterwallManager) FindToMap(args []any) map[int64]Periodicouterwall {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodicouterwallManager) FindAllToMap() map[int64]Periodicouterwall {
     items := p.Find(nil)
     return p.MakeMap(items)
}



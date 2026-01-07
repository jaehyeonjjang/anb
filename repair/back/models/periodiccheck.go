package models

import (
    "repair/models/periodiccheck"
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

type Periodiccheck struct {
            
    Id                int64 `json:"id"`         
    Content1                string `json:"content1"`         
    Content2                string `json:"content2"`         
    Content3                string `json:"content3"`         
    Content4                string `json:"content4"`         
    Content5                string `json:"content5"`         
    Content6                string `json:"content6"`         
    Content7                string `json:"content7"`         
    Content8                string `json:"content8"`         
    Content9                string `json:"content9"`         
    Content10                string `json:"content10"`         
    Content11                string `json:"content11"`         
    Content12                string `json:"content12"`         
    Content13                string `json:"content13"`         
    Content14                string `json:"content14"`         
    Content15                string `json:"content15"`         
    Content16                string `json:"content16"`         
    Use1                int `json:"use1"`         
    Use2                int `json:"use2"`         
    Use3                int `json:"use3"`         
    Use4                int `json:"use4"`         
    Need1                int `json:"need1"`         
    Need2                int `json:"need2"`         
    Need3                int `json:"need3"`         
    Need4                int `json:"need4"`         
    Aptdong                int64 `json:"aptdong"`         
    Periodic                int64 `json:"periodic"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodiccheckManager struct {
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



func (c *Periodiccheck) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodiccheckManager(conn *Connection) *PeriodiccheckManager {
    var item PeriodiccheckManager


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

func (p *PeriodiccheckManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodiccheckManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodiccheckManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodiccheckManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodiccheckManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodiccheckManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PeriodiccheckManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select pc_id, pc_content1, pc_content2, pc_content3, pc_content4, pc_content5, pc_content6, pc_content7, pc_content8, pc_content9, pc_content10, pc_content11, pc_content12, pc_content13, pc_content14, pc_content15, pc_content16, pc_use1, pc_use2, pc_use3, pc_use4, pc_need1, pc_need2, pc_need3, pc_need4, pc_aptdong, pc_periodic, pc_date from periodiccheck_tb")

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

func (p *PeriodiccheckManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodiccheck_tb")

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

func (p *PeriodiccheckManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select pc_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodiccheck_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *PeriodiccheckManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodiccheck_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodiccheckManager) Insert(item *Periodiccheck) error {
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
          query = "insert into periodiccheck_tb (pc_id, pc_content1, pc_content2, pc_content3, pc_content4, pc_content5, pc_content6, pc_content7, pc_content8, pc_content9, pc_content10, pc_content11, pc_content12, pc_content13, pc_content14, pc_content15, pc_content16, pc_use1, pc_use2, pc_use3, pc_use4, pc_need1, pc_need2, pc_need3, pc_need4, pc_aptdong, pc_periodic, pc_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)"
        } else {
          query = "insert into periodiccheck_tb (pc_id, pc_content1, pc_content2, pc_content3, pc_content4, pc_content5, pc_content6, pc_content7, pc_content8, pc_content9, pc_content10, pc_content11, pc_content12, pc_content13, pc_content14, pc_content15, pc_content16, pc_use1, pc_use2, pc_use3, pc_use4, pc_need1, pc_need2, pc_need3, pc_need4, pc_aptdong, pc_periodic, pc_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Content7, item.Content8, item.Content9, item.Content10, item.Content11, item.Content12, item.Content13, item.Content14, item.Content15, item.Content16, item.Use1, item.Use2, item.Use3, item.Use4, item.Need1, item.Need2, item.Need3, item.Need4, item.Aptdong, item.Periodic, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodiccheck_tb (pc_content1, pc_content2, pc_content3, pc_content4, pc_content5, pc_content6, pc_content7, pc_content8, pc_content9, pc_content10, pc_content11, pc_content12, pc_content13, pc_content14, pc_content15, pc_content16, pc_use1, pc_use2, pc_use3, pc_use4, pc_need1, pc_need2, pc_need3, pc_need4, pc_aptdong, pc_periodic, pc_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)"
        } else {
          query = "insert into periodiccheck_tb (pc_content1, pc_content2, pc_content3, pc_content4, pc_content5, pc_content6, pc_content7, pc_content8, pc_content9, pc_content10, pc_content11, pc_content12, pc_content13, pc_content14, pc_content15, pc_content16, pc_use1, pc_use2, pc_use3, pc_use4, pc_need1, pc_need2, pc_need3, pc_need4, pc_aptdong, pc_periodic, pc_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Content7, item.Content8, item.Content9, item.Content10, item.Content11, item.Content12, item.Content13, item.Content14, item.Content15, item.Content16, item.Use1, item.Use2, item.Use3, item.Use4, item.Need1, item.Need2, item.Need3, item.Need4, item.Aptdong, item.Periodic, item.Date)
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

func (p *PeriodiccheckManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodiccheck_tb where pc_id = ")
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

func (p *PeriodiccheckManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodiccheck_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodiccheckManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and pc_")
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

func (p *PeriodiccheckManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodiccheck_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodiccheckManager) Update(item *Periodiccheck) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodiccheck_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" pc_content1 = $1, pc_content2 = $2, pc_content3 = $3, pc_content4 = $4, pc_content5 = $5, pc_content6 = $6, pc_content7 = $7, pc_content8 = $8, pc_content9 = $9, pc_content10 = $10, pc_content11 = $11, pc_content12 = $12, pc_content13 = $13, pc_content14 = $14, pc_content15 = $15, pc_content16 = $16, pc_use1 = $17, pc_use2 = $18, pc_use3 = $19, pc_use4 = $20, pc_need1 = $21, pc_need2 = $22, pc_need3 = $23, pc_need4 = $24, pc_aptdong = $25, pc_periodic = $26, pc_date = $27 where pc_id = $28")
    } else {
        query.WriteString(" pc_content1 = ?, pc_content2 = ?, pc_content3 = ?, pc_content4 = ?, pc_content5 = ?, pc_content6 = ?, pc_content7 = ?, pc_content8 = ?, pc_content9 = ?, pc_content10 = ?, pc_content11 = ?, pc_content12 = ?, pc_content13 = ?, pc_content14 = ?, pc_content15 = ?, pc_content16 = ?, pc_use1 = ?, pc_use2 = ?, pc_use3 = ?, pc_use4 = ?, pc_need1 = ?, pc_need2 = ?, pc_need3 = ?, pc_need4 = ?, pc_aptdong = ?, pc_periodic = ?, pc_date = ? where pc_id = ?")
    }

	_, err := p.Exec(query.String() , item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Content7, item.Content8, item.Content9, item.Content10, item.Content11, item.Content12, item.Content13, item.Content14, item.Content15, item.Content16, item.Use1, item.Use2, item.Use3, item.Use4, item.Need1, item.Need2, item.Need3, item.Need4, item.Aptdong, item.Periodic, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodiccheckManager) UpdateWhere(columns []periodiccheck.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodiccheck_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodiccheck.ColumnId {
        initQuery.WriteString("pc_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent1 {
        initQuery.WriteString("pc_content1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent2 {
        initQuery.WriteString("pc_content2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent3 {
        initQuery.WriteString("pc_content3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent4 {
        initQuery.WriteString("pc_content4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent5 {
        initQuery.WriteString("pc_content5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent6 {
        initQuery.WriteString("pc_content6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent7 {
        initQuery.WriteString("pc_content7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent8 {
        initQuery.WriteString("pc_content8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent9 {
        initQuery.WriteString("pc_content9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent10 {
        initQuery.WriteString("pc_content10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent11 {
        initQuery.WriteString("pc_content11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent12 {
        initQuery.WriteString("pc_content12 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent13 {
        initQuery.WriteString("pc_content13 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent14 {
        initQuery.WriteString("pc_content14 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent15 {
        initQuery.WriteString("pc_content15 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnContent16 {
        initQuery.WriteString("pc_content16 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnUse1 {
        initQuery.WriteString("pc_use1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnUse2 {
        initQuery.WriteString("pc_use2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnUse3 {
        initQuery.WriteString("pc_use3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnUse4 {
        initQuery.WriteString("pc_use4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnNeed1 {
        initQuery.WriteString("pc_need1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnNeed2 {
        initQuery.WriteString("pc_need2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnNeed3 {
        initQuery.WriteString("pc_need3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnNeed4 {
        initQuery.WriteString("pc_need4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnAptdong {
        initQuery.WriteString("pc_aptdong = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnPeriodic {
        initQuery.WriteString("pc_periodic = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodiccheck.ColumnDate {
        initQuery.WriteString("pc_date = ?")
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

func (p *PeriodiccheckManager) UpdateContent1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content1 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content2 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content3 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content4 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content5 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content6 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content7 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content8 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content9 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content10 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content11 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent12(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content12 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent13(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content13 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent14(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content14 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent15(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content15 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateContent16(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_content16 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateUse1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use1 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateUse2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use2 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateUse3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use3 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateUse4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use4 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateNeed1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need1 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateNeed2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need2 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateNeed3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need3 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateNeed4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need4 = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdateAptdong(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_aptdong = ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) UpdatePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_periodic = ? where pc_id = ?"
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

func (p *PeriodiccheckManager) IncreaseUse1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use1 = pc_use1 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseUse2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use2 = pc_use2 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseUse3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use3 = pc_use3 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseUse4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_use4 = pc_use4 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseNeed1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need1 = pc_need1 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseNeed2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need2 = pc_need2 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseNeed3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need3 = pc_need3 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseNeed4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_need4 = pc_need4 + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreaseAptdong(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_aptdong = pc_aptdong + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodiccheckManager) IncreasePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodiccheck_tb set pc_periodic = pc_periodic + ? where pc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodiccheckManager) GetIdentity() int64 {
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

func (p *Periodiccheck) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodiccheckManager) ReadRow(rows *sql.Rows) *Periodiccheck {
    var item Periodiccheck
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Content5, &item.Content6, &item.Content7, &item.Content8, &item.Content9, &item.Content10, &item.Content11, &item.Content12, &item.Content13, &item.Content14, &item.Content15, &item.Content16, &item.Use1, &item.Use2, &item.Use3, &item.Use4, &item.Need1, &item.Need2, &item.Need3, &item.Need4, &item.Aptdong, &item.Periodic, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *PeriodiccheckManager) ReadRows(rows *sql.Rows) []Periodiccheck {
    items := make([]Periodiccheck, 0)

    for rows.Next() {
        var item Periodiccheck
        
    
        err := rows.Scan(&item.Id, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Content5, &item.Content6, &item.Content7, &item.Content8, &item.Content9, &item.Content10, &item.Content11, &item.Content12, &item.Content13, &item.Content14, &item.Content15, &item.Content16, &item.Use1, &item.Use2, &item.Use3, &item.Use4, &item.Need1, &item.Need2, &item.Need3, &item.Need4, &item.Aptdong, &item.Periodic, &item.Date)
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

func (p *PeriodiccheckManager) Get(id int64) *Periodiccheck {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and pc_id = ?")

    
    
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

func (p *PeriodiccheckManager) GetWhere(args []any) *Periodiccheck {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodiccheckManager) Count(args []any) int {
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

func (p *PeriodiccheckManager) FindAll() []Periodiccheck {
    return p.Find(nil)
}

func (p *PeriodiccheckManager) Find(args []any) []Periodiccheck {
    if !p.Conn.IsConnect() {
        items := make([]Periodiccheck, 0)
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
                query.WriteString(" and pc_")
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
            orderby = "pc_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "pc_" + orderby
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
            orderby = "pc_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "pc_" + orderby
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
        items := make([]Periodiccheck, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *PeriodiccheckManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and pc_")
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
    
    query.WriteString(" group by pc_")
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



func (p *PeriodiccheckManager) MakeMap(items []Periodiccheck) map[int64]Periodiccheck {
     ret := make(map[int64]Periodiccheck)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodiccheckManager) FindToMap(args []any) map[int64]Periodiccheck {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodiccheckManager) FindAllToMap() map[int64]Periodiccheck {
     items := p.Find(nil)
     return p.MakeMap(items)
}



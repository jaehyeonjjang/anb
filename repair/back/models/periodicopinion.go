package models

import (
    "repair/models/periodicopinion"
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

type Periodicopinion struct {
            
    Id                int64 `json:"id"`         
    Grade                int `json:"grade"`         
    Content1                string `json:"content1"`         
    Content2                string `json:"content2"`         
    Content3                string `json:"content3"`         
    Content4                string `json:"content4"`         
    Cause1                string `json:"cause1"`         
    Cause2                string `json:"cause2"`         
    Cause3                string `json:"cause3"`         
    Cause4                string `json:"cause4"`         
    Cause5                string `json:"cause5"`         
    Cause6                string `json:"cause6"`         
    Periodic                int64 `json:"periodic"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodicopinionManager struct {
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



func (c *Periodicopinion) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodicopinionManager(conn *Connection) *PeriodicopinionManager {
    var item PeriodicopinionManager


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

func (p *PeriodicopinionManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodicopinionManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodicopinionManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodicopinionManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodicopinionManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodicopinionManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PeriodicopinionManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select po_id, po_grade, po_content1, po_content2, po_content3, po_content4, po_cause1, po_cause2, po_cause3, po_cause4, po_cause5, po_cause6, po_periodic, po_date from periodicopinion_tb")

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

func (p *PeriodicopinionManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodicopinion_tb")

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

func (p *PeriodicopinionManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select po_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodicopinion_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *PeriodicopinionManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodicopinion_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodicopinionManager) Insert(item *Periodicopinion) error {
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
          query = "insert into periodicopinion_tb (po_id, po_grade, po_content1, po_content2, po_content3, po_content4, po_cause1, po_cause2, po_cause3, po_cause4, po_cause5, po_cause6, po_periodic, po_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
        } else {
          query = "insert into periodicopinion_tb (po_id, po_grade, po_content1, po_content2, po_content3, po_content4, po_cause1, po_cause2, po_cause3, po_cause4, po_cause5, po_cause6, po_periodic, po_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Grade, item.Content1, item.Content2, item.Content3, item.Content4, item.Cause1, item.Cause2, item.Cause3, item.Cause4, item.Cause5, item.Cause6, item.Periodic, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodicopinion_tb (po_grade, po_content1, po_content2, po_content3, po_content4, po_cause1, po_cause2, po_cause3, po_cause4, po_cause5, po_cause6, po_periodic, po_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
        } else {
          query = "insert into periodicopinion_tb (po_grade, po_content1, po_content2, po_content3, po_content4, po_cause1, po_cause2, po_cause3, po_cause4, po_cause5, po_cause6, po_periodic, po_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Grade, item.Content1, item.Content2, item.Content3, item.Content4, item.Cause1, item.Cause2, item.Cause3, item.Cause4, item.Cause5, item.Cause6, item.Periodic, item.Date)
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

func (p *PeriodicopinionManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodicopinion_tb where po_id = ")
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

func (p *PeriodicopinionManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodicopinion_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodicopinionManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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

func (p *PeriodicopinionManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodicopinion_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodicopinionManager) Update(item *Periodicopinion) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodicopinion_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" po_grade = $1, po_content1 = $2, po_content2 = $3, po_content3 = $4, po_content4 = $5, po_cause1 = $6, po_cause2 = $7, po_cause3 = $8, po_cause4 = $9, po_cause5 = $10, po_cause6 = $11, po_periodic = $12, po_date = $13 where po_id = $14")
    } else {
        query.WriteString(" po_grade = ?, po_content1 = ?, po_content2 = ?, po_content3 = ?, po_content4 = ?, po_cause1 = ?, po_cause2 = ?, po_cause3 = ?, po_cause4 = ?, po_cause5 = ?, po_cause6 = ?, po_periodic = ?, po_date = ? where po_id = ?")
    }

	_, err := p.Exec(query.String() , item.Grade, item.Content1, item.Content2, item.Content3, item.Content4, item.Cause1, item.Cause2, item.Cause3, item.Cause4, item.Cause5, item.Cause6, item.Periodic, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodicopinionManager) UpdateWhere(columns []periodicopinion.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodicopinion_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodicopinion.ColumnId {
        initQuery.WriteString("po_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnGrade {
        initQuery.WriteString("po_grade = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnContent1 {
        initQuery.WriteString("po_content1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnContent2 {
        initQuery.WriteString("po_content2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnContent3 {
        initQuery.WriteString("po_content3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnContent4 {
        initQuery.WriteString("po_content4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnCause1 {
        initQuery.WriteString("po_cause1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnCause2 {
        initQuery.WriteString("po_cause2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnCause3 {
        initQuery.WriteString("po_cause3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnCause4 {
        initQuery.WriteString("po_cause4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnCause5 {
        initQuery.WriteString("po_cause5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnCause6 {
        initQuery.WriteString("po_cause6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnPeriodic {
        initQuery.WriteString("po_periodic = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicopinion.ColumnDate {
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

func (p *PeriodicopinionManager) UpdateGrade(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_grade = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateContent1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_content1 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateContent2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_content2 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateContent3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_content3 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateContent4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_content4 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateCause1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_cause1 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateCause2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_cause2 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateCause3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_cause3 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateCause4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_cause4 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateCause5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_cause5 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdateCause6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_cause6 = ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) UpdatePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_periodic = ? where po_id = ?"
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

func (p *PeriodicopinionManager) IncreaseGrade(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_grade = po_grade + ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicopinionManager) IncreasePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicopinion_tb set po_periodic = po_periodic + ? where po_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodicopinionManager) GetIdentity() int64 {
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

func (p *Periodicopinion) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodicopinionManager) ReadRow(rows *sql.Rows) *Periodicopinion {
    var item Periodicopinion
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Grade, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Cause1, &item.Cause2, &item.Cause3, &item.Cause4, &item.Cause5, &item.Cause6, &item.Periodic, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *PeriodicopinionManager) ReadRows(rows *sql.Rows) []Periodicopinion {
    items := make([]Periodicopinion, 0)

    for rows.Next() {
        var item Periodicopinion
        
    
        err := rows.Scan(&item.Id, &item.Grade, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Cause1, &item.Cause2, &item.Cause3, &item.Cause4, &item.Cause5, &item.Cause6, &item.Periodic, &item.Date)
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

func (p *PeriodicopinionManager) Get(id int64) *Periodicopinion {
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

func (p *PeriodicopinionManager) GetWhere(args []any) *Periodicopinion {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodicopinionManager) Count(args []any) int {
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

func (p *PeriodicopinionManager) FindAll() []Periodicopinion {
    return p.Find(nil)
}

func (p *PeriodicopinionManager) Find(args []any) []Periodicopinion {
    if !p.Conn.IsConnect() {
        items := make([]Periodicopinion, 0)
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
        items := make([]Periodicopinion, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *PeriodicopinionManager) GetByPeriodic(periodic int64, args ...any) *Periodicopinion {
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




func (p *PeriodicopinionManager) GroupBy(name string, args []any) []Groupby {
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



func (p *PeriodicopinionManager) MakeMap(items []Periodicopinion) map[int64]Periodicopinion {
     ret := make(map[int64]Periodicopinion)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodicopinionManager) FindToMap(args []any) map[int64]Periodicopinion {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodicopinionManager) FindAllToMap() map[int64]Periodicopinion {
     items := p.Find(nil)
     return p.MakeMap(items)
}



package models

import (
    "repair/models/standardbackup"
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

type Standardbackup struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Direct                int64 `json:"direct"`         
    Labor                int `json:"labor"`         
    Cost                int `json:"cost"`         
    Unit                string `json:"unit"`         
    Order                int `json:"order"`         
    Original                int64 `json:"original"`         
    Category                int64 `json:"category"`         
    Apt                int64 `json:"apt"`         
    Standard                int64 `json:"standard"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type StandardbackupManager struct {
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



func (c *Standardbackup) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewStandardbackupManager(conn *Connection) *StandardbackupManager {
    var item StandardbackupManager


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

func (p *StandardbackupManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *StandardbackupManager) SetIndex(index string) {
    p.Index = index
}

func (p *StandardbackupManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *StandardbackupManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *StandardbackupManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *StandardbackupManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *StandardbackupManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select sb_id, sb_name, sb_direct, sb_labor, sb_cost, sb_unit, sb_order, sb_original, sb_category, sb_apt, sb_standard, sb_date from standardbackup_tb")

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

func (p *StandardbackupManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from standardbackup_tb")

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

func (p *StandardbackupManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select sb_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from standardbackup_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *StandardbackupManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate standardbackup_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *StandardbackupManager) Insert(item *Standardbackup) error {
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
          query = "insert into standardbackup_tb (sb_id, sb_name, sb_direct, sb_labor, sb_cost, sb_unit, sb_order, sb_original, sb_category, sb_apt, sb_standard, sb_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
        } else {
          query = "insert into standardbackup_tb (sb_id, sb_name, sb_direct, sb_labor, sb_cost, sb_unit, sb_order, sb_original, sb_category, sb_apt, sb_standard, sb_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Name, item.Direct, item.Labor, item.Cost, item.Unit, item.Order, item.Original, item.Category, item.Apt, item.Standard, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into standardbackup_tb (sb_name, sb_direct, sb_labor, sb_cost, sb_unit, sb_order, sb_original, sb_category, sb_apt, sb_standard, sb_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
        } else {
          query = "insert into standardbackup_tb (sb_name, sb_direct, sb_labor, sb_cost, sb_unit, sb_order, sb_original, sb_category, sb_apt, sb_standard, sb_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Name, item.Direct, item.Labor, item.Cost, item.Unit, item.Order, item.Original, item.Category, item.Apt, item.Standard, item.Date)
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

func (p *StandardbackupManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from standardbackup_tb where sb_id = ")
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

func (p *StandardbackupManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from standardbackup_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *StandardbackupManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and sb_")
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

func (p *StandardbackupManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from standardbackup_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *StandardbackupManager) Update(item *Standardbackup) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update standardbackup_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" sb_name = $1, sb_direct = $2, sb_labor = $3, sb_cost = $4, sb_unit = $5, sb_order = $6, sb_original = $7, sb_category = $8, sb_apt = $9, sb_standard = $10, sb_date = $11 where sb_id = $12")
    } else {
        query.WriteString(" sb_name = ?, sb_direct = ?, sb_labor = ?, sb_cost = ?, sb_unit = ?, sb_order = ?, sb_original = ?, sb_category = ?, sb_apt = ?, sb_standard = ?, sb_date = ? where sb_id = ?")
    }

	_, err := p.Exec(query.String() , item.Name, item.Direct, item.Labor, item.Cost, item.Unit, item.Order, item.Original, item.Category, item.Apt, item.Standard, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *StandardbackupManager) UpdateWhere(columns []standardbackup.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update standardbackup_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == standardbackup.ColumnId {
        initQuery.WriteString("sb_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnName {
        initQuery.WriteString("sb_name = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnDirect {
        initQuery.WriteString("sb_direct = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnLabor {
        initQuery.WriteString("sb_labor = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnCost {
        initQuery.WriteString("sb_cost = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnUnit {
        initQuery.WriteString("sb_unit = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnOrder {
        initQuery.WriteString("sb_order = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnOriginal {
        initQuery.WriteString("sb_original = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnCategory {
        initQuery.WriteString("sb_category = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnApt {
        initQuery.WriteString("sb_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnStandard {
        initQuery.WriteString("sb_standard = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardbackup.ColumnDate {
        initQuery.WriteString("sb_date = ?")
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

func (p *StandardbackupManager) UpdateName(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_name = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateDirect(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_direct = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateLabor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_labor = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateCost(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_cost = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateUnit(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_unit = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_order = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateOriginal(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_original = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_category = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_apt = ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) UpdateStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_standard = ? where sb_id = ?"
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

func (p *StandardbackupManager) IncreaseDirect(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_direct = sb_direct + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseLabor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_labor = sb_labor + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseCost(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_cost = sb_cost + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_order = sb_order + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseOriginal(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_original = sb_original + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_category = sb_category + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_apt = sb_apt + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardbackupManager) IncreaseStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardbackup_tb set sb_standard = sb_standard + ? where sb_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *StandardbackupManager) GetIdentity() int64 {
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

func (p *Standardbackup) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *StandardbackupManager) ReadRow(rows *sql.Rows) *Standardbackup {
    var item Standardbackup
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Direct, &item.Labor, &item.Cost, &item.Unit, &item.Order, &item.Original, &item.Category, &item.Apt, &item.Standard, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *StandardbackupManager) ReadRows(rows *sql.Rows) []Standardbackup {
    items := make([]Standardbackup, 0)

    for rows.Next() {
        var item Standardbackup
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Direct, &item.Labor, &item.Cost, &item.Unit, &item.Order, &item.Original, &item.Category, &item.Apt, &item.Standard, &item.Date)
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

func (p *StandardbackupManager) Get(id int64) *Standardbackup {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and sb_id = ?")

    
    
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

func (p *StandardbackupManager) GetWhere(args []any) *Standardbackup {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *StandardbackupManager) Count(args []any) int {
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

func (p *StandardbackupManager) FindAll() []Standardbackup {
    return p.Find(nil)
}

func (p *StandardbackupManager) Find(args []any) []Standardbackup {
    if !p.Conn.IsConnect() {
        items := make([]Standardbackup, 0)
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
                query.WriteString(" and sb_")
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
            orderby = "sb_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "sb_" + orderby
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
            orderby = "sb_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "sb_" + orderby
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
        items := make([]Standardbackup, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *StandardbackupManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and sb_")
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
    
    query.WriteString(" group by sb_")
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



func (p *StandardbackupManager) MakeMap(items []Standardbackup) map[int64]Standardbackup {
     ret := make(map[int64]Standardbackup)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *StandardbackupManager) FindToMap(args []any) map[int64]Standardbackup {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *StandardbackupManager) FindAllToMap() map[int64]Standardbackup {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *StandardbackupManager) MakeNameMap(items []Standardbackup) map[string]Standardbackup {
     ret := make(map[string]Standardbackup)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *StandardbackupManager) FindToNameMap(args []any) map[string]Standardbackup {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *StandardbackupManager) FindAllToNameMap() map[string]Standardbackup {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}

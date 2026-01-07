package models

import (
    "repair/models/breakdown"
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

type Breakdown struct {
            
    Id                int64 `json:"id"`         
    Topcategory                int64 `json:"topcategory"`         
    Subcategory                int64 `json:"subcategory"`         
    Category                int64 `json:"category"`         
    Method                int64 `json:"method"`         
    Count                int `json:"count"`         
    Lastdate                int `json:"lastdate"`         
    Duedate                int `json:"duedate"`         
    Remark                string `json:"remark"`         
    Elevator                int `json:"elevator"`         
    Percent                Double `json:"percent"`         
    Rate                Double `json:"rate"`         
    Type                int `json:"type"`         
    Dong                int64 `json:"dong"`         
    Standard                int64 `json:"standard"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type BreakdownManager struct {
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



func (c *Breakdown) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewBreakdownManager(conn *Connection) *BreakdownManager {
    var item BreakdownManager


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

func (p *BreakdownManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BreakdownManager) SetIndex(index string) {
    p.Index = index
}

func (p *BreakdownManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *BreakdownManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *BreakdownManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *BreakdownManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *BreakdownManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select b_id, b_topcategory, b_subcategory, b_category, b_method, b_count, b_lastdate, b_duedate, b_remark, b_elevator, b_percent, b_rate, b_type, b_dong, b_standard, b_apt, b_date, s_id, s_name, s_direct, s_labor, s_cost, s_unit, s_order, s_original, s_category, s_apt, s_date, c_id, c_name, c_level, c_parent, c_cycle, c_percent, c_unit, c_elevator, c_remark, c_order, c_apt, c_date, d_id, d_name, d_ground, d_underground, d_familycount, d_parking, d_elevator, d_basic, d_remark, d_order, d_apt, d_date from breakdown_tb, standard_tb, category_tb, dong_tb")

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
    
    ret.WriteString(" and b_standard = s_id ")
    
    ret.WriteString(" and b_method = c_id ")
    
    ret.WriteString(" and b_dong = d_id ")
    

    return ret.String()
}

func (p *BreakdownManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from breakdown_tb, standard_tb, category_tb, dong_tb")

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
    
    ret.WriteString(" and b_standard = s_id ")
    
    ret.WriteString(" and b_method = c_id ")
    
    ret.WriteString(" and b_dong = d_id ")
    

    return ret.String()
}

func (p *BreakdownManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select b_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from breakdown_tb, standard_tb, category_tb, dong_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and b_standard = s_id ")
    
    ret.WriteString(" and b_method = c_id ")
    
    ret.WriteString(" and b_dong = d_id ")
    


    return ret.String()
}

func (p *BreakdownManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate breakdown_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *BreakdownManager) Insert(item *Breakdown) error {
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
          query = "insert into breakdown_tb (b_id, b_topcategory, b_subcategory, b_category, b_method, b_count, b_lastdate, b_duedate, b_remark, b_elevator, b_percent, b_rate, b_type, b_dong, b_standard, b_apt, b_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)"
        } else {
          query = "insert into breakdown_tb (b_id, b_topcategory, b_subcategory, b_category, b_method, b_count, b_lastdate, b_duedate, b_remark, b_elevator, b_percent, b_rate, b_type, b_dong, b_standard, b_apt, b_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Topcategory, item.Subcategory, item.Category, item.Method, item.Count, item.Lastdate, item.Duedate, item.Remark, item.Elevator, item.Percent, item.Rate, item.Type, item.Dong, item.Standard, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into breakdown_tb (b_topcategory, b_subcategory, b_category, b_method, b_count, b_lastdate, b_duedate, b_remark, b_elevator, b_percent, b_rate, b_type, b_dong, b_standard, b_apt, b_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)"
        } else {
          query = "insert into breakdown_tb (b_topcategory, b_subcategory, b_category, b_method, b_count, b_lastdate, b_duedate, b_remark, b_elevator, b_percent, b_rate, b_type, b_dong, b_standard, b_apt, b_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Topcategory, item.Subcategory, item.Category, item.Method, item.Count, item.Lastdate, item.Duedate, item.Remark, item.Elevator, item.Percent, item.Rate, item.Type, item.Dong, item.Standard, item.Apt, item.Date)
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

func (p *BreakdownManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from breakdown_tb where b_id = ")
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

func (p *BreakdownManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from breakdown_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *BreakdownManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and b_")
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

func (p *BreakdownManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from breakdown_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *BreakdownManager) Update(item *Breakdown) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update breakdown_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" b_topcategory = $1, b_subcategory = $2, b_category = $3, b_method = $4, b_count = $5, b_lastdate = $6, b_duedate = $7, b_remark = $8, b_elevator = $9, b_percent = $10, b_rate = $11, b_type = $12, b_dong = $13, b_standard = $14, b_apt = $15, b_date = $16 where b_id = $17")
    } else {
        query.WriteString(" b_topcategory = ?, b_subcategory = ?, b_category = ?, b_method = ?, b_count = ?, b_lastdate = ?, b_duedate = ?, b_remark = ?, b_elevator = ?, b_percent = ?, b_rate = ?, b_type = ?, b_dong = ?, b_standard = ?, b_apt = ?, b_date = ? where b_id = ?")
    }

	_, err := p.Exec(query.String() , item.Topcategory, item.Subcategory, item.Category, item.Method, item.Count, item.Lastdate, item.Duedate, item.Remark, item.Elevator, item.Percent, item.Rate, item.Type, item.Dong, item.Standard, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *BreakdownManager) UpdateWhere(columns []breakdown.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update breakdown_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == breakdown.ColumnId {
        initQuery.WriteString("b_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnTopcategory {
        initQuery.WriteString("b_topcategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnSubcategory {
        initQuery.WriteString("b_subcategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnCategory {
        initQuery.WriteString("b_category = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnMethod {
        initQuery.WriteString("b_method = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnCount {
        initQuery.WriteString("b_count = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnLastdate {
        initQuery.WriteString("b_lastdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnDuedate {
        initQuery.WriteString("b_duedate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnRemark {
        initQuery.WriteString("b_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnElevator {
        initQuery.WriteString("b_elevator = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnPercent {
        initQuery.WriteString("b_percent = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnRate {
        initQuery.WriteString("b_rate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnType {
        initQuery.WriteString("b_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnDong {
        initQuery.WriteString("b_dong = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnStandard {
        initQuery.WriteString("b_standard = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnApt {
        initQuery.WriteString("b_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdown.ColumnDate {
        initQuery.WriteString("b_date = ?")
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

func (p *BreakdownManager) UpdateTopcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_topcategory = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateSubcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_subcategory = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_category = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateMethod(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_method = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateCount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_count = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateLastdate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_lastdate = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateDuedate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_duedate = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_remark = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateElevator(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_elevator = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdatePercent(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_percent = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateRate(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_rate = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_type = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateDong(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_dong = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_standard = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_apt = ? where b_id = ?"
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

func (p *BreakdownManager) IncreaseTopcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_topcategory = b_topcategory + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseSubcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_subcategory = b_subcategory + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_category = b_category + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseMethod(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_method = b_method + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseCount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_count = b_count + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseLastdate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_lastdate = b_lastdate + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseDuedate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_duedate = b_duedate + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseElevator(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_elevator = b_elevator + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreasePercent(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_percent = b_percent + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseRate(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_rate = b_rate + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_type = b_type + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseDong(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_dong = b_dong + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_standard = b_standard + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdown_tb set b_apt = b_apt + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *BreakdownManager) GetIdentity() int64 {
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

func (p *Breakdown) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *BreakdownManager) ReadRow(rows *sql.Rows) *Breakdown {
    var item Breakdown
    var err error

    var _standard Standard
    var _category Category
    var _dong Dong
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Topcategory, &item.Subcategory, &item.Category, &item.Method, &item.Count, &item.Lastdate, &item.Duedate, &item.Remark, &item.Elevator, &item.Percent, &item.Rate, &item.Type, &item.Dong, &item.Standard, &item.Apt, &item.Date, &_standard.Id, &_standard.Name, &_standard.Direct, &_standard.Labor, &_standard.Cost, &_standard.Unit, &_standard.Order, &_standard.Original, &_standard.Category, &_standard.Apt, &_standard.Date, &_category.Id, &_category.Name, &_category.Level, &_category.Parent, &_category.Cycle, &_category.Percent, &_category.Unit, &_category.Elevator, &_category.Remark, &_category.Order, &_category.Apt, &_category.Date, &_dong.Id, &_dong.Name, &_dong.Ground, &_dong.Underground, &_dong.Familycount, &_dong.Parking, &_dong.Elevator, &_dong.Basic, &_dong.Remark, &_dong.Order, &_dong.Apt, &_dong.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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
        _standard.InitExtra()
        item.AddExtra("standard",  _standard)
_category.InitExtra()
        item.AddExtra("category",  _category)
_dong.InitExtra()
        item.AddExtra("dong",  _dong)

        return &item
    }
}

func (p *BreakdownManager) ReadRows(rows *sql.Rows) []Breakdown {
    items := make([]Breakdown, 0)

    for rows.Next() {
        var item Breakdown
        var _standard Standard
            var _category Category
            var _dong Dong
            
    
        err := rows.Scan(&item.Id, &item.Topcategory, &item.Subcategory, &item.Category, &item.Method, &item.Count, &item.Lastdate, &item.Duedate, &item.Remark, &item.Elevator, &item.Percent, &item.Rate, &item.Type, &item.Dong, &item.Standard, &item.Apt, &item.Date, &_standard.Id, &_standard.Name, &_standard.Direct, &_standard.Labor, &_standard.Cost, &_standard.Unit, &_standard.Order, &_standard.Original, &_standard.Category, &_standard.Apt, &_standard.Date, &_category.Id, &_category.Name, &_category.Level, &_category.Parent, &_category.Cycle, &_category.Percent, &_category.Unit, &_category.Elevator, &_category.Remark, &_category.Order, &_category.Apt, &_category.Date, &_dong.Id, &_dong.Name, &_dong.Ground, &_dong.Underground, &_dong.Familycount, &_dong.Parking, &_dong.Elevator, &_dong.Basic, &_dong.Remark, &_dong.Order, &_dong.Apt, &_dong.Date)
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
        _standard.InitExtra()
        item.AddExtra("standard",  _standard)
_category.InitExtra()
        item.AddExtra("category",  _category)
_dong.InitExtra()
        item.AddExtra("dong",  _dong)

        items = append(items, item)
    }


     return items
}

func (p *BreakdownManager) Get(id int64) *Breakdown {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and b_id = ?")

    
    query.WriteString(" and b_standard = s_id ")
    
    query.WriteString(" and b_method = c_id ")
    
    query.WriteString(" and b_dong = d_id ")
    
    
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

func (p *BreakdownManager) GetWhere(args []any) *Breakdown {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *BreakdownManager) Count(args []any) int {
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

func (p *BreakdownManager) FindAll() []Breakdown {
    return p.Find(nil)
}

func (p *BreakdownManager) Find(args []any) []Breakdown {
    if !p.Conn.IsConnect() {
        items := make([]Breakdown, 0)
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
                query.WriteString(" and b_")
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
            orderby = "b_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "b_" + orderby
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
            orderby = "b_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "b_" + orderby
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
        items := make([]Breakdown, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *BreakdownManager) CountByApt(apt int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *BreakdownManager) FindByApt(apt int64, args ...any) []Breakdown {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *BreakdownManager) DeleteByApt(apt int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from breakdown_tb where b_apt = ?"
    _, err := p.Exec(query, apt)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *BreakdownManager) CountByAptDong(apt int64, dong int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    if dong != 0 { 
        rets = append(rets, Where{Column:"dong", Value:dong, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *BreakdownManager) CountByAptStandard(apt int64, standard int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    if standard != 0 { 
        rets = append(rets, Where{Column:"standard", Value:standard, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *BreakdownManager) UpdateDuedateById(duedate int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "update breakdown_tb set b_duedate = ? where 1=1 and b_id = ?"
	_, err := p.Exec(query, duedate, id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err    
}

func (p *BreakdownManager) UpdateLastdateById(lastdate int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "update breakdown_tb set b_lastdate = ? where 1=1 and b_id = ?"
	_, err := p.Exec(query, lastdate, id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err    
}

func (p *BreakdownManager) FindByCategory(category int64, args ...any) []Breakdown {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if category != 0 { 
        rets = append(rets, Where{Column:"category", Value:category, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *BreakdownManager) DeleteByCategory(category int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from breakdown_tb where b_category = ?"
    _, err := p.Exec(query, category)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *BreakdownManager) FindByMethod(method int64, args ...any) []Breakdown {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if method != 0 { 
        rets = append(rets, Where{Column:"method", Value:method, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *BreakdownManager) UpdateSubcategoryByCategory(subcategory int64, category int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "update breakdown_tb set b_subcategory = ? where 1=1 and b_category = ?"
	_, err := p.Exec(query, subcategory, category)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err    
}


func (p *BreakdownManager) Sum(args []any) *Breakdown {
    if !p.Conn.IsConnect() {
        var item Breakdown
        return &item
    }

    var params []any

    
    var query strings.Builder
    query.WriteString("select sum(b_count) from breakdown_tb")

    if p.Index != "" {
        query.WriteString(" use index(")
        query.WriteString(p.Index)
        query.WriteString(") ")
    }

    query.WriteString("where 1=1 ")

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
                query.WriteString(" and b_")
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
    
    startpage := (page - 1) * pagesize
    
    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "b_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "b_" + orderby
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
            orderby = "b_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "b_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(query.String(), params...)

    var item Breakdown
    
    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        err := rows.Scan(&item.Count)        
        if err != nil {
            if p.Log {
                log.Error().Str("error", err.Error()).Msg("SQL")
            }

            return &item
        }
    }

    return &item        
}

func (p *BreakdownManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and b_")
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
    
    query.WriteString(" group by b_")
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



func (p *BreakdownManager) MakeMap(items []Breakdown) map[int64]Breakdown {
     ret := make(map[int64]Breakdown)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *BreakdownManager) FindToMap(args []any) map[int64]Breakdown {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *BreakdownManager) FindAllToMap() map[int64]Breakdown {
     items := p.Find(nil)
     return p.MakeMap(items)
}



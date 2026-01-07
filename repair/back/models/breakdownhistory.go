package models

import (
    "repair/models/breakdownhistory"
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

type Breakdownhistory struct {
            
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
    Originalcount                int `json:"originalcount"`         
    Originalprice                int `json:"originalprice"`         
    Originalduedate                int `json:"originalduedate"`         
    Totalcount                int `json:"totalcount"`         
    Totalprice                int64 `json:"totalprice"`         
    Dong                int64 `json:"dong"`         
    Standard                int64 `json:"standard"`         
    Breakdown                int64 `json:"breakdown"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type BreakdownhistoryManager struct {
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



func (c *Breakdownhistory) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewBreakdownhistoryManager(conn *Connection) *BreakdownhistoryManager {
    var item BreakdownhistoryManager


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

func (p *BreakdownhistoryManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BreakdownhistoryManager) SetIndex(index string) {
    p.Index = index
}

func (p *BreakdownhistoryManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *BreakdownhistoryManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *BreakdownhistoryManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *BreakdownhistoryManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *BreakdownhistoryManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select bh_id, bh_topcategory, bh_subcategory, bh_category, bh_method, bh_count, bh_lastdate, bh_duedate, bh_remark, bh_elevator, bh_percent, bh_rate, bh_type, bh_originalcount, bh_originalprice, bh_originalduedate, bh_totalcount, bh_totalprice, bh_dong, bh_standard, bh_breakdown, bh_apt, bh_date, c_id, c_name, c_level, c_parent, c_cycle, c_percent, c_unit, c_elevator, c_remark, c_order, c_apt, c_date, s_id, s_name, s_direct, s_labor, s_cost, s_unit, s_order, s_original, s_category, s_apt, s_date from breakdownhistory_tb, category_tb, standard_tb")

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
    
    ret.WriteString(" and bh_method = c_id ")
    
    ret.WriteString(" and bh_standard = s_id ")
    

    return ret.String()
}

func (p *BreakdownhistoryManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from breakdownhistory_tb, category_tb, standard_tb")

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
    
    ret.WriteString(" and bh_method = c_id ")
    
    ret.WriteString(" and bh_standard = s_id ")
    

    return ret.String()
}

func (p *BreakdownhistoryManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select bh_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from breakdownhistory_tb, category_tb, standard_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and bh_method = c_id ")
    
    ret.WriteString(" and bh_standard = s_id ")
    


    return ret.String()
}

func (p *BreakdownhistoryManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate breakdownhistory_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *BreakdownhistoryManager) Insert(item *Breakdownhistory) error {
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
          query = "insert into breakdownhistory_tb (bh_id, bh_topcategory, bh_subcategory, bh_category, bh_method, bh_count, bh_lastdate, bh_duedate, bh_remark, bh_elevator, bh_percent, bh_rate, bh_type, bh_originalcount, bh_originalprice, bh_originalduedate, bh_totalcount, bh_totalprice, bh_dong, bh_standard, bh_breakdown, bh_apt, bh_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)"
        } else {
          query = "insert into breakdownhistory_tb (bh_id, bh_topcategory, bh_subcategory, bh_category, bh_method, bh_count, bh_lastdate, bh_duedate, bh_remark, bh_elevator, bh_percent, bh_rate, bh_type, bh_originalcount, bh_originalprice, bh_originalduedate, bh_totalcount, bh_totalprice, bh_dong, bh_standard, bh_breakdown, bh_apt, bh_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Topcategory, item.Subcategory, item.Category, item.Method, item.Count, item.Lastdate, item.Duedate, item.Remark, item.Elevator, item.Percent, item.Rate, item.Type, item.Originalcount, item.Originalprice, item.Originalduedate, item.Totalcount, item.Totalprice, item.Dong, item.Standard, item.Breakdown, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into breakdownhistory_tb (bh_topcategory, bh_subcategory, bh_category, bh_method, bh_count, bh_lastdate, bh_duedate, bh_remark, bh_elevator, bh_percent, bh_rate, bh_type, bh_originalcount, bh_originalprice, bh_originalduedate, bh_totalcount, bh_totalprice, bh_dong, bh_standard, bh_breakdown, bh_apt, bh_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)"
        } else {
          query = "insert into breakdownhistory_tb (bh_topcategory, bh_subcategory, bh_category, bh_method, bh_count, bh_lastdate, bh_duedate, bh_remark, bh_elevator, bh_percent, bh_rate, bh_type, bh_originalcount, bh_originalprice, bh_originalduedate, bh_totalcount, bh_totalprice, bh_dong, bh_standard, bh_breakdown, bh_apt, bh_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Topcategory, item.Subcategory, item.Category, item.Method, item.Count, item.Lastdate, item.Duedate, item.Remark, item.Elevator, item.Percent, item.Rate, item.Type, item.Originalcount, item.Originalprice, item.Originalduedate, item.Totalcount, item.Totalprice, item.Dong, item.Standard, item.Breakdown, item.Apt, item.Date)
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

func (p *BreakdownhistoryManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from breakdownhistory_tb where bh_id = ")
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

func (p *BreakdownhistoryManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from breakdownhistory_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *BreakdownhistoryManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and bh_")
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

func (p *BreakdownhistoryManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from breakdownhistory_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *BreakdownhistoryManager) Update(item *Breakdownhistory) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update breakdownhistory_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" bh_topcategory = $1, bh_subcategory = $2, bh_category = $3, bh_method = $4, bh_count = $5, bh_lastdate = $6, bh_duedate = $7, bh_remark = $8, bh_elevator = $9, bh_percent = $10, bh_rate = $11, bh_type = $12, bh_originalcount = $13, bh_originalprice = $14, bh_originalduedate = $15, bh_totalcount = $16, bh_totalprice = $17, bh_dong = $18, bh_standard = $19, bh_breakdown = $20, bh_apt = $21, bh_date = $22 where bh_id = $23")
    } else {
        query.WriteString(" bh_topcategory = ?, bh_subcategory = ?, bh_category = ?, bh_method = ?, bh_count = ?, bh_lastdate = ?, bh_duedate = ?, bh_remark = ?, bh_elevator = ?, bh_percent = ?, bh_rate = ?, bh_type = ?, bh_originalcount = ?, bh_originalprice = ?, bh_originalduedate = ?, bh_totalcount = ?, bh_totalprice = ?, bh_dong = ?, bh_standard = ?, bh_breakdown = ?, bh_apt = ?, bh_date = ? where bh_id = ?")
    }

	_, err := p.Exec(query.String() , item.Topcategory, item.Subcategory, item.Category, item.Method, item.Count, item.Lastdate, item.Duedate, item.Remark, item.Elevator, item.Percent, item.Rate, item.Type, item.Originalcount, item.Originalprice, item.Originalduedate, item.Totalcount, item.Totalprice, item.Dong, item.Standard, item.Breakdown, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *BreakdownhistoryManager) UpdateWhere(columns []breakdownhistory.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update breakdownhistory_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == breakdownhistory.ColumnId {
        initQuery.WriteString("bh_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnTopcategory {
        initQuery.WriteString("bh_topcategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnSubcategory {
        initQuery.WriteString("bh_subcategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnCategory {
        initQuery.WriteString("bh_category = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnMethod {
        initQuery.WriteString("bh_method = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnCount {
        initQuery.WriteString("bh_count = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnLastdate {
        initQuery.WriteString("bh_lastdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnDuedate {
        initQuery.WriteString("bh_duedate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnRemark {
        initQuery.WriteString("bh_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnElevator {
        initQuery.WriteString("bh_elevator = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnPercent {
        initQuery.WriteString("bh_percent = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnRate {
        initQuery.WriteString("bh_rate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnType {
        initQuery.WriteString("bh_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnOriginalcount {
        initQuery.WriteString("bh_originalcount = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnOriginalprice {
        initQuery.WriteString("bh_originalprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnOriginalduedate {
        initQuery.WriteString("bh_originalduedate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnTotalcount {
        initQuery.WriteString("bh_totalcount = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnTotalprice {
        initQuery.WriteString("bh_totalprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnDong {
        initQuery.WriteString("bh_dong = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnStandard {
        initQuery.WriteString("bh_standard = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnBreakdown {
        initQuery.WriteString("bh_breakdown = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnApt {
        initQuery.WriteString("bh_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == breakdownhistory.ColumnDate {
        initQuery.WriteString("bh_date = ?")
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

func (p *BreakdownhistoryManager) UpdateTopcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_topcategory = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateSubcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_subcategory = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_category = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateMethod(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_method = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateCount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_count = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateLastdate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_lastdate = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateDuedate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_duedate = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_remark = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateElevator(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_elevator = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdatePercent(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_percent = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateRate(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_rate = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_type = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateOriginalcount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_originalcount = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateOriginalprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_originalprice = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateOriginalduedate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_originalduedate = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateTotalcount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_totalcount = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateTotalprice(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_totalprice = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateDong(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_dong = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_standard = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateBreakdown(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_breakdown = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_apt = ? where bh_id = ?"
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

func (p *BreakdownhistoryManager) IncreaseTopcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_topcategory = bh_topcategory + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseSubcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_subcategory = bh_subcategory + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_category = bh_category + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseMethod(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_method = bh_method + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseCount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_count = bh_count + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseLastdate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_lastdate = bh_lastdate + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseDuedate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_duedate = bh_duedate + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseElevator(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_elevator = bh_elevator + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreasePercent(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_percent = bh_percent + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseRate(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_rate = bh_rate + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_type = bh_type + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseOriginalcount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_originalcount = bh_originalcount + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseOriginalprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_originalprice = bh_originalprice + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseOriginalduedate(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_originalduedate = bh_originalduedate + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseTotalcount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_totalcount = bh_totalcount + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseTotalprice(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_totalprice = bh_totalprice + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseDong(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_dong = bh_dong + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_standard = bh_standard + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseBreakdown(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_breakdown = bh_breakdown + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *BreakdownhistoryManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update breakdownhistory_tb set bh_apt = bh_apt + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *BreakdownhistoryManager) GetIdentity() int64 {
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

func (p *Breakdownhistory) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *BreakdownhistoryManager) ReadRow(rows *sql.Rows) *Breakdownhistory {
    var item Breakdownhistory
    var err error

    var _category Category
    var _standard Standard
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Topcategory, &item.Subcategory, &item.Category, &item.Method, &item.Count, &item.Lastdate, &item.Duedate, &item.Remark, &item.Elevator, &item.Percent, &item.Rate, &item.Type, &item.Originalcount, &item.Originalprice, &item.Originalduedate, &item.Totalcount, &item.Totalprice, &item.Dong, &item.Standard, &item.Breakdown, &item.Apt, &item.Date, &_category.Id, &_category.Name, &_category.Level, &_category.Parent, &_category.Cycle, &_category.Percent, &_category.Unit, &_category.Elevator, &_category.Remark, &_category.Order, &_category.Apt, &_category.Date, &_standard.Id, &_standard.Name, &_standard.Direct, &_standard.Labor, &_standard.Cost, &_standard.Unit, &_standard.Order, &_standard.Original, &_standard.Category, &_standard.Apt, &_standard.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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
        _category.InitExtra()
        item.AddExtra("category",  _category)
_standard.InitExtra()
        item.AddExtra("standard",  _standard)

        return &item
    }
}

func (p *BreakdownhistoryManager) ReadRows(rows *sql.Rows) []Breakdownhistory {
    items := make([]Breakdownhistory, 0)

    for rows.Next() {
        var item Breakdownhistory
        var _category Category
            var _standard Standard
            
    
        err := rows.Scan(&item.Id, &item.Topcategory, &item.Subcategory, &item.Category, &item.Method, &item.Count, &item.Lastdate, &item.Duedate, &item.Remark, &item.Elevator, &item.Percent, &item.Rate, &item.Type, &item.Originalcount, &item.Originalprice, &item.Originalduedate, &item.Totalcount, &item.Totalprice, &item.Dong, &item.Standard, &item.Breakdown, &item.Apt, &item.Date, &_category.Id, &_category.Name, &_category.Level, &_category.Parent, &_category.Cycle, &_category.Percent, &_category.Unit, &_category.Elevator, &_category.Remark, &_category.Order, &_category.Apt, &_category.Date, &_standard.Id, &_standard.Name, &_standard.Direct, &_standard.Labor, &_standard.Cost, &_standard.Unit, &_standard.Order, &_standard.Original, &_standard.Category, &_standard.Apt, &_standard.Date)
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
        _category.InitExtra()
        item.AddExtra("category",  _category)
_standard.InitExtra()
        item.AddExtra("standard",  _standard)

        items = append(items, item)
    }


     return items
}

func (p *BreakdownhistoryManager) Get(id int64) *Breakdownhistory {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and bh_id = ?")

    
    query.WriteString(" and bh_method = c_id ")
    
    query.WriteString(" and bh_standard = s_id ")
    
    
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

func (p *BreakdownhistoryManager) GetWhere(args []any) *Breakdownhistory {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *BreakdownhistoryManager) Count(args []any) int {
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

func (p *BreakdownhistoryManager) FindAll() []Breakdownhistory {
    return p.Find(nil)
}

func (p *BreakdownhistoryManager) Find(args []any) []Breakdownhistory {
    if !p.Conn.IsConnect() {
        items := make([]Breakdownhistory, 0)
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
                query.WriteString(" and bh_")
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
            orderby = "bh_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "bh_" + orderby
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
            orderby = "bh_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "bh_" + orderby
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
        items := make([]Breakdownhistory, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *BreakdownhistoryManager) CountByApt(apt int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *BreakdownhistoryManager) FindByApt(apt int64, args ...any) []Breakdownhistory {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *BreakdownhistoryManager) DeleteByApt(apt int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from breakdownhistory_tb where bh_apt = ?"
    _, err := p.Exec(query, apt)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *BreakdownhistoryManager) CountByBreakdown(breakdown int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if breakdown != 0 { 
        rets = append(rets, Where{Column:"breakdown", Value:breakdown, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *BreakdownhistoryManager) GetByBreakdown(breakdown int64, args ...any) *Breakdownhistory {
    rets := make([]any, 0)
    rets = append(rets, args...)
    if breakdown != 0 {
        rets = append(rets, Where{Column:"breakdown", Value:breakdown, Compare:"="})        
    }
    
    items := p.Find(rets)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}

func (p *BreakdownhistoryManager) FindByCategory(category int64, args ...any) []Breakdownhistory {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if category != 0 { 
        rets = append(rets, Where{Column:"category", Value:category, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *BreakdownhistoryManager) DeleteByCategory(category int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from breakdownhistory_tb where bh_category = ?"
    _, err := p.Exec(query, category)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *BreakdownhistoryManager) FindByMethod(method int64, args ...any) []Breakdownhistory {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if method != 0 { 
        rets = append(rets, Where{Column:"method", Value:method, Compare:"="})
     }
    
    
    return p.Find(rets)
}


func (p *BreakdownhistoryManager) Sum(args []any) *Breakdownhistory {
    if !p.Conn.IsConnect() {
        var item Breakdownhistory
        return &item
    }

    var params []any

    
    var query strings.Builder
    query.WriteString("select sum(bh_count,totalprice) from breakdownhistory_tb")

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
                query.WriteString(" and bh_")
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
            orderby = "bh_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "bh_" + orderby
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
            orderby = "bh_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "bh_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(query.String(), params...)

    var item Breakdownhistory
    
    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        err := rows.Scan(&item.Count,&item.Totalprice)        
        if err != nil {
            if p.Log {
                log.Error().Str("error", err.Error()).Msg("SQL")
            }

            return &item
        }
    }

    return &item        
}

func (p *BreakdownhistoryManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and bh_")
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
    
    query.WriteString(" group by bh_")
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



func (p *BreakdownhistoryManager) MakeMap(items []Breakdownhistory) map[int64]Breakdownhistory {
     ret := make(map[int64]Breakdownhistory)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *BreakdownhistoryManager) FindToMap(args []any) map[int64]Breakdownhistory {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *BreakdownhistoryManager) FindAllToMap() map[int64]Breakdownhistory {
     items := p.Find(nil)
     return p.MakeMap(items)
}



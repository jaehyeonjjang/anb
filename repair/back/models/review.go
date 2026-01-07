package models

import (
    "repair/models/review"
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

type Review struct {
            
    Id                int64 `json:"id"`         
    Topcategory                int64 `json:"topcategory"`         
    Subcategory                int64 `json:"subcategory"`         
    Category                int64 `json:"category"`         
    Standard                int64 `json:"standard"`         
    Method                int64 `json:"method"`         
    Cycle                string `json:"cycle"`         
    Percent                int `json:"percent"`         
    Count                int `json:"count"`         
    Price                int64 `json:"price"`         
    Content                string `json:"content"`         
    Adjust                string `json:"adjust"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type ReviewManager struct {
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



func (c *Review) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewReviewManager(conn *Connection) *ReviewManager {
    var item ReviewManager


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

func (p *ReviewManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ReviewManager) SetIndex(index string) {
    p.Index = index
}

func (p *ReviewManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *ReviewManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *ReviewManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *ReviewManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *ReviewManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select re_id, re_topcategory, re_subcategory, re_category, re_standard, re_method, re_cycle, re_percent, re_count, re_price, re_content, re_adjust, re_apt, re_date, c_id, c_name, c_level, c_parent, c_cycle, c_percent, c_unit, c_elevator, c_remark, c_order, c_apt, c_date, s_id, s_name, s_direct, s_labor, s_cost, s_unit, s_order, s_original, s_category, s_apt, s_date from review_tb, category_tb, standard_tb")

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
    
    ret.WriteString(" and re_category = c_id ")
    
    ret.WriteString(" and re_standard = s_id ")
    

    return ret.String()
}

func (p *ReviewManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from review_tb, category_tb, standard_tb")

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
    
    ret.WriteString(" and re_category = c_id ")
    
    ret.WriteString(" and re_standard = s_id ")
    

    return ret.String()
}

func (p *ReviewManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select re_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from review_tb, category_tb, standard_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and re_category = c_id ")
    
    ret.WriteString(" and re_standard = s_id ")
    


    return ret.String()
}

func (p *ReviewManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate review_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *ReviewManager) Insert(item *Review) error {
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
          query = "insert into review_tb (re_id, re_topcategory, re_subcategory, re_category, re_standard, re_method, re_cycle, re_percent, re_count, re_price, re_content, re_adjust, re_apt, re_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
        } else {
          query = "insert into review_tb (re_id, re_topcategory, re_subcategory, re_category, re_standard, re_method, re_cycle, re_percent, re_count, re_price, re_content, re_adjust, re_apt, re_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Topcategory, item.Subcategory, item.Category, item.Standard, item.Method, item.Cycle, item.Percent, item.Count, item.Price, item.Content, item.Adjust, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into review_tb (re_topcategory, re_subcategory, re_category, re_standard, re_method, re_cycle, re_percent, re_count, re_price, re_content, re_adjust, re_apt, re_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
        } else {
          query = "insert into review_tb (re_topcategory, re_subcategory, re_category, re_standard, re_method, re_cycle, re_percent, re_count, re_price, re_content, re_adjust, re_apt, re_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Topcategory, item.Subcategory, item.Category, item.Standard, item.Method, item.Cycle, item.Percent, item.Count, item.Price, item.Content, item.Adjust, item.Apt, item.Date)
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

func (p *ReviewManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from review_tb where re_id = ")
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

func (p *ReviewManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from review_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *ReviewManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and re_")
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

func (p *ReviewManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from review_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *ReviewManager) Update(item *Review) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update review_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" re_topcategory = $1, re_subcategory = $2, re_category = $3, re_standard = $4, re_method = $5, re_cycle = $6, re_percent = $7, re_count = $8, re_price = $9, re_content = $10, re_adjust = $11, re_apt = $12, re_date = $13 where re_id = $14")
    } else {
        query.WriteString(" re_topcategory = ?, re_subcategory = ?, re_category = ?, re_standard = ?, re_method = ?, re_cycle = ?, re_percent = ?, re_count = ?, re_price = ?, re_content = ?, re_adjust = ?, re_apt = ?, re_date = ? where re_id = ?")
    }

	_, err := p.Exec(query.String() , item.Topcategory, item.Subcategory, item.Category, item.Standard, item.Method, item.Cycle, item.Percent, item.Count, item.Price, item.Content, item.Adjust, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *ReviewManager) UpdateWhere(columns []review.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update review_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == review.ColumnId {
        initQuery.WriteString("re_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnTopcategory {
        initQuery.WriteString("re_topcategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnSubcategory {
        initQuery.WriteString("re_subcategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnCategory {
        initQuery.WriteString("re_category = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnStandard {
        initQuery.WriteString("re_standard = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnMethod {
        initQuery.WriteString("re_method = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnCycle {
        initQuery.WriteString("re_cycle = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnPercent {
        initQuery.WriteString("re_percent = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnCount {
        initQuery.WriteString("re_count = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnPrice {
        initQuery.WriteString("re_price = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnContent {
        initQuery.WriteString("re_content = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnAdjust {
        initQuery.WriteString("re_adjust = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnApt {
        initQuery.WriteString("re_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == review.ColumnDate {
        initQuery.WriteString("re_date = ?")
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

func (p *ReviewManager) UpdateTopcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_topcategory = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateSubcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_subcategory = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_category = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_standard = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateMethod(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_method = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateCycle(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_cycle = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdatePercent(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_percent = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateCount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_count = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdatePrice(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_price = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateContent(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_content = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateAdjust(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_adjust = ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_apt = ? where re_id = ?"
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

func (p *ReviewManager) IncreaseTopcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_topcategory = re_topcategory + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreaseSubcategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_subcategory = re_subcategory + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreaseCategory(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_category = re_category + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreaseStandard(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_standard = re_standard + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreaseMethod(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_method = re_method + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreasePercent(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_percent = re_percent + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreaseCount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_count = re_count + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreasePrice(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_price = re_price + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ReviewManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update review_tb set re_apt = re_apt + ? where re_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *ReviewManager) GetIdentity() int64 {
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

func (p *Review) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *ReviewManager) ReadRow(rows *sql.Rows) *Review {
    var item Review
    var err error

    var _category Category
    var _standard Standard
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Topcategory, &item.Subcategory, &item.Category, &item.Standard, &item.Method, &item.Cycle, &item.Percent, &item.Count, &item.Price, &item.Content, &item.Adjust, &item.Apt, &item.Date, &_category.Id, &_category.Name, &_category.Level, &_category.Parent, &_category.Cycle, &_category.Percent, &_category.Unit, &_category.Elevator, &_category.Remark, &_category.Order, &_category.Apt, &_category.Date, &_standard.Id, &_standard.Name, &_standard.Direct, &_standard.Labor, &_standard.Cost, &_standard.Unit, &_standard.Order, &_standard.Original, &_standard.Category, &_standard.Apt, &_standard.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *ReviewManager) ReadRows(rows *sql.Rows) []Review {
    items := make([]Review, 0)

    for rows.Next() {
        var item Review
        var _category Category
            var _standard Standard
            
    
        err := rows.Scan(&item.Id, &item.Topcategory, &item.Subcategory, &item.Category, &item.Standard, &item.Method, &item.Cycle, &item.Percent, &item.Count, &item.Price, &item.Content, &item.Adjust, &item.Apt, &item.Date, &_category.Id, &_category.Name, &_category.Level, &_category.Parent, &_category.Cycle, &_category.Percent, &_category.Unit, &_category.Elevator, &_category.Remark, &_category.Order, &_category.Apt, &_category.Date, &_standard.Id, &_standard.Name, &_standard.Direct, &_standard.Labor, &_standard.Cost, &_standard.Unit, &_standard.Order, &_standard.Original, &_standard.Category, &_standard.Apt, &_standard.Date)
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

func (p *ReviewManager) Get(id int64) *Review {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and re_id = ?")

    
    query.WriteString(" and re_category = c_id ")
    
    query.WriteString(" and re_standard = s_id ")
    
    
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

func (p *ReviewManager) GetWhere(args []any) *Review {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *ReviewManager) Count(args []any) int {
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

func (p *ReviewManager) FindAll() []Review {
    return p.Find(nil)
}

func (p *ReviewManager) Find(args []any) []Review {
    if !p.Conn.IsConnect() {
        items := make([]Review, 0)
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
                query.WriteString(" and re_")
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
            orderby = "re_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "re_" + orderby
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
            orderby = "re_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "re_" + orderby
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
        items := make([]Review, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *ReviewManager) CountByApt(apt int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *ReviewManager) FindByApt(apt int64, args ...any) []Review {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *ReviewManager) FindByCategory(category int64, args ...any) []Review {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if category != 0 { 
        rets = append(rets, Where{Column:"category", Value:category, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *ReviewManager) DeleteByCategory(category int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from review_tb where re_category = ?"
    _, err := p.Exec(query, category)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *ReviewManager) FindByMethod(method int64, args ...any) []Review {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if method != 0 { 
        rets = append(rets, Where{Column:"method", Value:method, Compare:"="})
     }
    
    
    return p.Find(rets)
}


func (p *ReviewManager) Sum(args []any) *Review {
    if !p.Conn.IsConnect() {
        var item Review
        return &item
    }

    var params []any

    
    var query strings.Builder
    query.WriteString("select sum(re_count,price) from review_tb")

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
                query.WriteString(" and re_")
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
            orderby = "re_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "re_" + orderby
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
            orderby = "re_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "re_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(query.String(), params...)

    var item Review
    
    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        err := rows.Scan(&item.Count,&item.Price)        
        if err != nil {
            if p.Log {
                log.Error().Str("error", err.Error()).Msg("SQL")
            }

            return &item
        }
    }

    return &item        
}

func (p *ReviewManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and re_")
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
    
    query.WriteString(" group by re_")
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



func (p *ReviewManager) MakeMap(items []Review) map[int64]Review {
     ret := make(map[int64]Review)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *ReviewManager) FindToMap(args []any) map[int64]Review {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *ReviewManager) FindAllToMap() map[int64]Review {
     items := p.Find(nil)
     return p.MakeMap(items)
}



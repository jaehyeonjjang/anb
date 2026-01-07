package models

import (
    "repair/models/aptperiodic"
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

type Aptperiodic struct {
            
    Id                int64 `json:"id"`         
    Current1                string `json:"current1"`         
    Current2                string `json:"current2"`         
    Current3                string `json:"current3"`         
    Current4                string `json:"current4"`         
    Current5                string `json:"current5"`         
    Current6                int `json:"current6"`         
    Current7                int `json:"current7"`         
    Current8                int `json:"current8"`         
    Current9                string `json:"current9"`         
    Current10                string `json:"current10"`         
    Current11                int `json:"current11"`         
    Current12                int `json:"current12"`         
    Current13                int `json:"current13"`         
    Current14                int `json:"current14"`         
    Current15                string `json:"current15"`         
    Current16                string `json:"current16"`         
    Current17                string `json:"current17"`         
    Current18                string `json:"current18"`         
    Current19                string `json:"current19"`         
    Current20                string `json:"current20"`         
    Current21                string `json:"current21"`         
    Current22                int `json:"current22"`         
    Current23                int `json:"current23"`         
    Outline1                Double `json:"outline1"`         
    Outline2                Double `json:"outline2"`         
    Outline3                int `json:"outline3"`         
    Outline4                int `json:"outline4"`         
    Outline5                int `json:"outline5"`         
    Outline6                Double `json:"outline6"`         
    Outline7                int `json:"outline7"`         
    Outline7content                string `json:"outline7content"`         
    Outline8                Double `json:"outline8"`         
    Outline9                int `json:"outline9"`         
    Outline9content                string `json:"outline9content"`         
    Record1                string `json:"record1"`         
    Record2                string `json:"record2"`         
    Record3                string `json:"record3"`         
    Record4                string `json:"record4"`         
    Record5                string `json:"record5"`         
    Deligate                string `json:"deligate"`         
    Facilitydivision                int `json:"facilitydivision"`         
    Facilitytype                int `json:"facilitytype"`         
    Facilitycategory                int `json:"facilitycategory"`         
    Struct1                string `json:"struct1"`         
    Struct2                string `json:"struct2"`         
    Struct3                string `json:"struct3"`         
    Struct4                string `json:"struct4"`         
    Struct5                int `json:"struct5"`         
    Struct6                string `json:"struct6"`         
    Struct7                string `json:"struct7"`         
    Struct8                string `json:"struct8"`         
    Struct9                string `json:"struct9"`         
    Struct10                string `json:"struct10"`         
    Struct11                string `json:"struct11"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type AptperiodicManager struct {
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



func (c *Aptperiodic) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewAptperiodicManager(conn *Connection) *AptperiodicManager {
    var item AptperiodicManager


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

func (p *AptperiodicManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptperiodicManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptperiodicManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *AptperiodicManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *AptperiodicManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *AptperiodicManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *AptperiodicManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select ad_id, ad_current1, ad_current2, ad_current3, ad_current4, ad_current5, ad_current6, ad_current7, ad_current8, ad_current9, ad_current10, ad_current11, ad_current12, ad_current13, ad_current14, ad_current15, ad_current16, ad_current17, ad_current18, ad_current19, ad_current20, ad_current21, ad_current22, ad_current23, ad_outline1, ad_outline2, ad_outline3, ad_outline4, ad_outline5, ad_outline6, ad_outline7, ad_outline7content, ad_outline8, ad_outline9, ad_outline9content, ad_record1, ad_record2, ad_record3, ad_record4, ad_record5, ad_deligate, ad_facilitydivision, ad_facilitytype, ad_facilitycategory, ad_struct1, ad_struct2, ad_struct3, ad_struct4, ad_struct5, ad_struct6, ad_struct7, ad_struct8, ad_struct9, ad_struct10, ad_struct11, ad_date from aptperiodic_tb")

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

func (p *AptperiodicManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from aptperiodic_tb")

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

func (p *AptperiodicManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select ad_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from aptperiodic_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *AptperiodicManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate aptperiodic_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *AptperiodicManager) Insert(item *Aptperiodic) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Record4 == "" {
       item.Record4 = "1000-01-01"
    }
	
    if item.Record5 == "" {
       item.Record5 = "1000-01-01"
    }
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into aptperiodic_tb (ad_id, ad_current1, ad_current2, ad_current3, ad_current4, ad_current5, ad_current6, ad_current7, ad_current8, ad_current9, ad_current10, ad_current11, ad_current12, ad_current13, ad_current14, ad_current15, ad_current16, ad_current17, ad_current18, ad_current19, ad_current20, ad_current21, ad_current22, ad_current23, ad_outline1, ad_outline2, ad_outline3, ad_outline4, ad_outline5, ad_outline6, ad_outline7, ad_outline7content, ad_outline8, ad_outline9, ad_outline9content, ad_record1, ad_record2, ad_record3, ad_record4, ad_record5, ad_deligate, ad_facilitydivision, ad_facilitytype, ad_facilitycategory, ad_struct1, ad_struct2, ad_struct3, ad_struct4, ad_struct5, ad_struct6, ad_struct7, ad_struct8, ad_struct9, ad_struct10, ad_struct11, ad_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56)"
        } else {
          query = "insert into aptperiodic_tb (ad_id, ad_current1, ad_current2, ad_current3, ad_current4, ad_current5, ad_current6, ad_current7, ad_current8, ad_current9, ad_current10, ad_current11, ad_current12, ad_current13, ad_current14, ad_current15, ad_current16, ad_current17, ad_current18, ad_current19, ad_current20, ad_current21, ad_current22, ad_current23, ad_outline1, ad_outline2, ad_outline3, ad_outline4, ad_outline5, ad_outline6, ad_outline7, ad_outline7content, ad_outline8, ad_outline9, ad_outline9content, ad_record1, ad_record2, ad_record3, ad_record4, ad_record5, ad_deligate, ad_facilitydivision, ad_facilitytype, ad_facilitycategory, ad_struct1, ad_struct2, ad_struct3, ad_struct4, ad_struct5, ad_struct6, ad_struct7, ad_struct8, ad_struct9, ad_struct10, ad_struct11, ad_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Current1, item.Current2, item.Current3, item.Current4, item.Current5, item.Current6, item.Current7, item.Current8, item.Current9, item.Current10, item.Current11, item.Current12, item.Current13, item.Current14, item.Current15, item.Current16, item.Current17, item.Current18, item.Current19, item.Current20, item.Current21, item.Current22, item.Current23, item.Outline1, item.Outline2, item.Outline3, item.Outline4, item.Outline5, item.Outline6, item.Outline7, item.Outline7content, item.Outline8, item.Outline9, item.Outline9content, item.Record1, item.Record2, item.Record3, item.Record4, item.Record5, item.Deligate, item.Facilitydivision, item.Facilitytype, item.Facilitycategory, item.Struct1, item.Struct2, item.Struct3, item.Struct4, item.Struct5, item.Struct6, item.Struct7, item.Struct8, item.Struct9, item.Struct10, item.Struct11, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into aptperiodic_tb (ad_current1, ad_current2, ad_current3, ad_current4, ad_current5, ad_current6, ad_current7, ad_current8, ad_current9, ad_current10, ad_current11, ad_current12, ad_current13, ad_current14, ad_current15, ad_current16, ad_current17, ad_current18, ad_current19, ad_current20, ad_current21, ad_current22, ad_current23, ad_outline1, ad_outline2, ad_outline3, ad_outline4, ad_outline5, ad_outline6, ad_outline7, ad_outline7content, ad_outline8, ad_outline9, ad_outline9content, ad_record1, ad_record2, ad_record3, ad_record4, ad_record5, ad_deligate, ad_facilitydivision, ad_facilitytype, ad_facilitycategory, ad_struct1, ad_struct2, ad_struct3, ad_struct4, ad_struct5, ad_struct6, ad_struct7, ad_struct8, ad_struct9, ad_struct10, ad_struct11, ad_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)"
        } else {
          query = "insert into aptperiodic_tb (ad_current1, ad_current2, ad_current3, ad_current4, ad_current5, ad_current6, ad_current7, ad_current8, ad_current9, ad_current10, ad_current11, ad_current12, ad_current13, ad_current14, ad_current15, ad_current16, ad_current17, ad_current18, ad_current19, ad_current20, ad_current21, ad_current22, ad_current23, ad_outline1, ad_outline2, ad_outline3, ad_outline4, ad_outline5, ad_outline6, ad_outline7, ad_outline7content, ad_outline8, ad_outline9, ad_outline9content, ad_record1, ad_record2, ad_record3, ad_record4, ad_record5, ad_deligate, ad_facilitydivision, ad_facilitytype, ad_facilitycategory, ad_struct1, ad_struct2, ad_struct3, ad_struct4, ad_struct5, ad_struct6, ad_struct7, ad_struct8, ad_struct9, ad_struct10, ad_struct11, ad_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Current1, item.Current2, item.Current3, item.Current4, item.Current5, item.Current6, item.Current7, item.Current8, item.Current9, item.Current10, item.Current11, item.Current12, item.Current13, item.Current14, item.Current15, item.Current16, item.Current17, item.Current18, item.Current19, item.Current20, item.Current21, item.Current22, item.Current23, item.Outline1, item.Outline2, item.Outline3, item.Outline4, item.Outline5, item.Outline6, item.Outline7, item.Outline7content, item.Outline8, item.Outline9, item.Outline9content, item.Record1, item.Record2, item.Record3, item.Record4, item.Record5, item.Deligate, item.Facilitydivision, item.Facilitytype, item.Facilitycategory, item.Struct1, item.Struct2, item.Struct3, item.Struct4, item.Struct5, item.Struct6, item.Struct7, item.Struct8, item.Struct9, item.Struct10, item.Struct11, item.Date)
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

func (p *AptperiodicManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from aptperiodic_tb where ad_id = ")
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

func (p *AptperiodicManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from aptperiodic_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *AptperiodicManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and ad_")
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

func (p *AptperiodicManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from aptperiodic_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *AptperiodicManager) Update(item *Aptperiodic) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Record4 == "" {
       item.Record4 = "1000-01-01"
    }
	
    if item.Record5 == "" {
       item.Record5 = "1000-01-01"
    }
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update aptperiodic_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" ad_current1 = $1, ad_current2 = $2, ad_current3 = $3, ad_current4 = $4, ad_current5 = $5, ad_current6 = $6, ad_current7 = $7, ad_current8 = $8, ad_current9 = $9, ad_current10 = $10, ad_current11 = $11, ad_current12 = $12, ad_current13 = $13, ad_current14 = $14, ad_current15 = $15, ad_current16 = $16, ad_current17 = $17, ad_current18 = $18, ad_current19 = $19, ad_current20 = $20, ad_current21 = $21, ad_current22 = $22, ad_current23 = $23, ad_outline1 = $24, ad_outline2 = $25, ad_outline3 = $26, ad_outline4 = $27, ad_outline5 = $28, ad_outline6 = $29, ad_outline7 = $30, ad_outline7content = $31, ad_outline8 = $32, ad_outline9 = $33, ad_outline9content = $34, ad_record1 = $35, ad_record2 = $36, ad_record3 = $37, ad_record4 = $38, ad_record5 = $39, ad_deligate = $40, ad_facilitydivision = $41, ad_facilitytype = $42, ad_facilitycategory = $43, ad_struct1 = $44, ad_struct2 = $45, ad_struct3 = $46, ad_struct4 = $47, ad_struct5 = $48, ad_struct6 = $49, ad_struct7 = $50, ad_struct8 = $51, ad_struct9 = $52, ad_struct10 = $53, ad_struct11 = $54, ad_date = $55 where ad_id = $56")
    } else {
        query.WriteString(" ad_current1 = ?, ad_current2 = ?, ad_current3 = ?, ad_current4 = ?, ad_current5 = ?, ad_current6 = ?, ad_current7 = ?, ad_current8 = ?, ad_current9 = ?, ad_current10 = ?, ad_current11 = ?, ad_current12 = ?, ad_current13 = ?, ad_current14 = ?, ad_current15 = ?, ad_current16 = ?, ad_current17 = ?, ad_current18 = ?, ad_current19 = ?, ad_current20 = ?, ad_current21 = ?, ad_current22 = ?, ad_current23 = ?, ad_outline1 = ?, ad_outline2 = ?, ad_outline3 = ?, ad_outline4 = ?, ad_outline5 = ?, ad_outline6 = ?, ad_outline7 = ?, ad_outline7content = ?, ad_outline8 = ?, ad_outline9 = ?, ad_outline9content = ?, ad_record1 = ?, ad_record2 = ?, ad_record3 = ?, ad_record4 = ?, ad_record5 = ?, ad_deligate = ?, ad_facilitydivision = ?, ad_facilitytype = ?, ad_facilitycategory = ?, ad_struct1 = ?, ad_struct2 = ?, ad_struct3 = ?, ad_struct4 = ?, ad_struct5 = ?, ad_struct6 = ?, ad_struct7 = ?, ad_struct8 = ?, ad_struct9 = ?, ad_struct10 = ?, ad_struct11 = ?, ad_date = ? where ad_id = ?")
    }

	_, err := p.Exec(query.String() , item.Current1, item.Current2, item.Current3, item.Current4, item.Current5, item.Current6, item.Current7, item.Current8, item.Current9, item.Current10, item.Current11, item.Current12, item.Current13, item.Current14, item.Current15, item.Current16, item.Current17, item.Current18, item.Current19, item.Current20, item.Current21, item.Current22, item.Current23, item.Outline1, item.Outline2, item.Outline3, item.Outline4, item.Outline5, item.Outline6, item.Outline7, item.Outline7content, item.Outline8, item.Outline9, item.Outline9content, item.Record1, item.Record2, item.Record3, item.Record4, item.Record5, item.Deligate, item.Facilitydivision, item.Facilitytype, item.Facilitycategory, item.Struct1, item.Struct2, item.Struct3, item.Struct4, item.Struct5, item.Struct6, item.Struct7, item.Struct8, item.Struct9, item.Struct10, item.Struct11, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *AptperiodicManager) UpdateWhere(columns []aptperiodic.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update aptperiodic_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == aptperiodic.ColumnId {
        initQuery.WriteString("ad_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent1 {
        initQuery.WriteString("ad_current1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent2 {
        initQuery.WriteString("ad_current2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent3 {
        initQuery.WriteString("ad_current3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent4 {
        initQuery.WriteString("ad_current4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent5 {
        initQuery.WriteString("ad_current5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent6 {
        initQuery.WriteString("ad_current6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent7 {
        initQuery.WriteString("ad_current7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent8 {
        initQuery.WriteString("ad_current8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent9 {
        initQuery.WriteString("ad_current9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent10 {
        initQuery.WriteString("ad_current10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent11 {
        initQuery.WriteString("ad_current11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent12 {
        initQuery.WriteString("ad_current12 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent13 {
        initQuery.WriteString("ad_current13 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent14 {
        initQuery.WriteString("ad_current14 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent15 {
        initQuery.WriteString("ad_current15 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent16 {
        initQuery.WriteString("ad_current16 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent17 {
        initQuery.WriteString("ad_current17 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent18 {
        initQuery.WriteString("ad_current18 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent19 {
        initQuery.WriteString("ad_current19 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent20 {
        initQuery.WriteString("ad_current20 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent21 {
        initQuery.WriteString("ad_current21 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent22 {
        initQuery.WriteString("ad_current22 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnCurrent23 {
        initQuery.WriteString("ad_current23 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline1 {
        initQuery.WriteString("ad_outline1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline2 {
        initQuery.WriteString("ad_outline2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline3 {
        initQuery.WriteString("ad_outline3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline4 {
        initQuery.WriteString("ad_outline4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline5 {
        initQuery.WriteString("ad_outline5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline6 {
        initQuery.WriteString("ad_outline6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline7 {
        initQuery.WriteString("ad_outline7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline7content {
        initQuery.WriteString("ad_outline7content = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline8 {
        initQuery.WriteString("ad_outline8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline9 {
        initQuery.WriteString("ad_outline9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnOutline9content {
        initQuery.WriteString("ad_outline9content = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnRecord1 {
        initQuery.WriteString("ad_record1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnRecord2 {
        initQuery.WriteString("ad_record2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnRecord3 {
        initQuery.WriteString("ad_record3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnRecord4 {
        initQuery.WriteString("ad_record4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnRecord5 {
        initQuery.WriteString("ad_record5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnDeligate {
        initQuery.WriteString("ad_deligate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnFacilitydivision {
        initQuery.WriteString("ad_facilitydivision = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnFacilitytype {
        initQuery.WriteString("ad_facilitytype = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnFacilitycategory {
        initQuery.WriteString("ad_facilitycategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct1 {
        initQuery.WriteString("ad_struct1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct2 {
        initQuery.WriteString("ad_struct2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct3 {
        initQuery.WriteString("ad_struct3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct4 {
        initQuery.WriteString("ad_struct4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct5 {
        initQuery.WriteString("ad_struct5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct6 {
        initQuery.WriteString("ad_struct6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct7 {
        initQuery.WriteString("ad_struct7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct8 {
        initQuery.WriteString("ad_struct8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct9 {
        initQuery.WriteString("ad_struct9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct10 {
        initQuery.WriteString("ad_struct10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnStruct11 {
        initQuery.WriteString("ad_struct11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == aptperiodic.ColumnDate {
        initQuery.WriteString("ad_date = ?")
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

func (p *AptperiodicManager) UpdateCurrent1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current1 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current2 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current3 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current4 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current5 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current6 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current7 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current8 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current9 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current10 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent11(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current11 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent12(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current12 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent13(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current13 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent14(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current14 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent15(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current15 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent16(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current16 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent17(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current17 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent18(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current18 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent19(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current19 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent20(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current20 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent21(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current21 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent22(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current22 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateCurrent23(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current23 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline1(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline1 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline2(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline2 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline3 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline4 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline5 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline6(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline6 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline7 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline7content(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline7content = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline8(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline8 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline9 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateOutline9content(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline9content = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateRecord1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_record1 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateRecord2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_record2 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateRecord3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_record3 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateRecord4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_record4 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateRecord5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_record5 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateDeligate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_deligate = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateFacilitydivision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_facilitydivision = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateFacilitytype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_facilitytype = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateFacilitycategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_facilitycategory = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct1 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct2 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct3 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct4 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct5 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct6 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct7 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct8 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct9 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct10 = ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) UpdateStruct11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct11 = ? where ad_id = ?"
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

func (p *AptperiodicManager) IncreaseCurrent6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current6 = ad_current6 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current7 = ad_current7 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current8 = ad_current8 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent11(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current11 = ad_current11 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent12(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current12 = ad_current12 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent13(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current13 = ad_current13 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent14(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current14 = ad_current14 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent22(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current22 = ad_current22 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseCurrent23(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_current23 = ad_current23 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline1(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline1 = ad_outline1 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline2(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline2 = ad_outline2 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline3 = ad_outline3 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline4 = ad_outline4 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline5 = ad_outline5 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline6(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline6 = ad_outline6 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline7 = ad_outline7 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline8(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline8 = ad_outline8 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseOutline9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_outline9 = ad_outline9 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseFacilitydivision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_facilitydivision = ad_facilitydivision + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseFacilitytype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_facilitytype = ad_facilitytype + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseFacilitycategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_facilitycategory = ad_facilitycategory + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptperiodicManager) IncreaseStruct5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update aptperiodic_tb set ad_struct5 = ad_struct5 + ? where ad_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *AptperiodicManager) GetIdentity() int64 {
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

func (p *Aptperiodic) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *AptperiodicManager) ReadRow(rows *sql.Rows) *Aptperiodic {
    var item Aptperiodic
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Current1, &item.Current2, &item.Current3, &item.Current4, &item.Current5, &item.Current6, &item.Current7, &item.Current8, &item.Current9, &item.Current10, &item.Current11, &item.Current12, &item.Current13, &item.Current14, &item.Current15, &item.Current16, &item.Current17, &item.Current18, &item.Current19, &item.Current20, &item.Current21, &item.Current22, &item.Current23, &item.Outline1, &item.Outline2, &item.Outline3, &item.Outline4, &item.Outline5, &item.Outline6, &item.Outline7, &item.Outline7content, &item.Outline8, &item.Outline9, &item.Outline9content, &item.Record1, &item.Record2, &item.Record3, &item.Record4, &item.Record5, &item.Deligate, &item.Facilitydivision, &item.Facilitytype, &item.Facilitycategory, &item.Struct1, &item.Struct2, &item.Struct3, &item.Struct4, &item.Struct5, &item.Struct6, &item.Struct7, &item.Struct8, &item.Struct9, &item.Struct10, &item.Struct11, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Record4 == "0000-00-00" || item.Record4 == "1000-01-01" || item.Record4 == "9999-01-01" {
            item.Record4 = ""
        }
        
        if item.Record5 == "0000-00-00" || item.Record5 == "1000-01-01" || item.Record5 == "9999-01-01" {
            item.Record5 = ""
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
        
        return &item
    }
}

func (p *AptperiodicManager) ReadRows(rows *sql.Rows) []Aptperiodic {
    items := make([]Aptperiodic, 0)

    for rows.Next() {
        var item Aptperiodic
        
    
        err := rows.Scan(&item.Id, &item.Current1, &item.Current2, &item.Current3, &item.Current4, &item.Current5, &item.Current6, &item.Current7, &item.Current8, &item.Current9, &item.Current10, &item.Current11, &item.Current12, &item.Current13, &item.Current14, &item.Current15, &item.Current16, &item.Current17, &item.Current18, &item.Current19, &item.Current20, &item.Current21, &item.Current22, &item.Current23, &item.Outline1, &item.Outline2, &item.Outline3, &item.Outline4, &item.Outline5, &item.Outline6, &item.Outline7, &item.Outline7content, &item.Outline8, &item.Outline9, &item.Outline9content, &item.Record1, &item.Record2, &item.Record3, &item.Record4, &item.Record5, &item.Deligate, &item.Facilitydivision, &item.Facilitytype, &item.Facilitycategory, &item.Struct1, &item.Struct2, &item.Struct3, &item.Struct4, &item.Struct5, &item.Struct6, &item.Struct7, &item.Struct8, &item.Struct9, &item.Struct10, &item.Struct11, &item.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		if item.Record4 == "0000-00-00" || item.Record4 == "1000-01-01" || item.Record4 == "9999-01-01" {
            item.Record4 = ""
        }
        
		if item.Record5 == "0000-00-00" || item.Record5 == "1000-01-01" || item.Record5 == "9999-01-01" {
            item.Record5 = ""
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

func (p *AptperiodicManager) Get(id int64) *Aptperiodic {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and ad_id = ?")

    
    
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

func (p *AptperiodicManager) GetWhere(args []any) *Aptperiodic {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *AptperiodicManager) Count(args []any) int {
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

func (p *AptperiodicManager) FindAll() []Aptperiodic {
    return p.Find(nil)
}

func (p *AptperiodicManager) Find(args []any) []Aptperiodic {
    if !p.Conn.IsConnect() {
        items := make([]Aptperiodic, 0)
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
                query.WriteString(" and ad_")
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
            orderby = "ad_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "ad_" + orderby
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
            orderby = "ad_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "ad_" + orderby
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
        items := make([]Aptperiodic, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *AptperiodicManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and ad_")
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
    
    query.WriteString(" group by ad_")
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



func (p *AptperiodicManager) MakeMap(items []Aptperiodic) map[int64]Aptperiodic {
     ret := make(map[int64]Aptperiodic)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *AptperiodicManager) FindToMap(args []any) map[int64]Aptperiodic {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *AptperiodicManager) FindAllToMap() map[int64]Aptperiodic {
     items := p.Find(nil)
     return p.MakeMap(items)
}



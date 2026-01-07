package models

import (
    "repair/models/standardwage"
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

type Standardwage struct {
            
    Id                int64 `json:"id"`         
    Person1                int `json:"person1"`         
    Person2                int `json:"person2"`         
    Person3                int `json:"person3"`         
    Person4                int `json:"person4"`         
    Person5                int `json:"person5"`         
    Person6                int `json:"person6"`         
    Person7                int `json:"person7"`         
    Person8                int `json:"person8"`         
    Person9                int `json:"person9"`         
    Person10                int `json:"person10"`         
    Techprice1                int `json:"techprice1"`         
    Techprice2                int `json:"techprice2"`         
    Techprice3                int `json:"techprice3"`         
    Techprice4                int `json:"techprice4"`         
    Financialprice1                int `json:"financialprice1"`         
    Financialprice2                int `json:"financialprice2"`         
    Financialprice3                int `json:"financialprice3"`         
    Financialprice4                int `json:"financialprice4"`         
    Directprice                int `json:"directprice"`         
    Printprice1                int `json:"printprice1"`         
    Printprice2                int `json:"printprice2"`         
    Lossprice                int `json:"lossprice"`         
    Gasprice                int `json:"gasprice"`         
    Travelprice                int `json:"travelprice"`         
    Travel                int `json:"travel"`         
    Loss                int `json:"loss"`         
    Gas                int `json:"gas"`         
    Etc                int `json:"etc"`         
    Danger                int `json:"danger"`         
    Machine                int `json:"machine"`         
    Print                int `json:"print"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type StandardwageManager struct {
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



func (c *Standardwage) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewStandardwageManager(conn *Connection) *StandardwageManager {
    var item StandardwageManager


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

func (p *StandardwageManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *StandardwageManager) SetIndex(index string) {
    p.Index = index
}

func (p *StandardwageManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *StandardwageManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *StandardwageManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *StandardwageManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *StandardwageManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select sw_id, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5, sw_person6, sw_person7, sw_person8, sw_person9, sw_person10, sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4, sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4, sw_directprice, sw_printprice1, sw_printprice2, sw_lossprice, sw_gasprice, sw_travelprice, sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print, sw_date from standardwage_tb")

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

func (p *StandardwageManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from standardwage_tb")

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

func (p *StandardwageManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select sw_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from standardwage_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *StandardwageManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate standardwage_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *StandardwageManager) Insert(item *Standardwage) error {
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
          query = "insert into standardwage_tb (sw_id, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5, sw_person6, sw_person7, sw_person8, sw_person9, sw_person10, sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4, sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4, sw_directprice, sw_printprice1, sw_printprice2, sw_lossprice, sw_gasprice, sw_travelprice, sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print, sw_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33)"
        } else {
          query = "insert into standardwage_tb (sw_id, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5, sw_person6, sw_person7, sw_person8, sw_person9, sw_person10, sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4, sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4, sw_directprice, sw_printprice1, sw_printprice2, sw_lossprice, sw_gasprice, sw_travelprice, sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print, sw_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Person1, item.Person2, item.Person3, item.Person4, item.Person5, item.Person6, item.Person7, item.Person8, item.Person9, item.Person10, item.Techprice1, item.Techprice2, item.Techprice3, item.Techprice4, item.Financialprice1, item.Financialprice2, item.Financialprice3, item.Financialprice4, item.Directprice, item.Printprice1, item.Printprice2, item.Lossprice, item.Gasprice, item.Travelprice, item.Travel, item.Loss, item.Gas, item.Etc, item.Danger, item.Machine, item.Print, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into standardwage_tb (sw_person1, sw_person2, sw_person3, sw_person4, sw_person5, sw_person6, sw_person7, sw_person8, sw_person9, sw_person10, sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4, sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4, sw_directprice, sw_printprice1, sw_printprice2, sw_lossprice, sw_gasprice, sw_travelprice, sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print, sw_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32)"
        } else {
          query = "insert into standardwage_tb (sw_person1, sw_person2, sw_person3, sw_person4, sw_person5, sw_person6, sw_person7, sw_person8, sw_person9, sw_person10, sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4, sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4, sw_directprice, sw_printprice1, sw_printprice2, sw_lossprice, sw_gasprice, sw_travelprice, sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print, sw_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Person1, item.Person2, item.Person3, item.Person4, item.Person5, item.Person6, item.Person7, item.Person8, item.Person9, item.Person10, item.Techprice1, item.Techprice2, item.Techprice3, item.Techprice4, item.Financialprice1, item.Financialprice2, item.Financialprice3, item.Financialprice4, item.Directprice, item.Printprice1, item.Printprice2, item.Lossprice, item.Gasprice, item.Travelprice, item.Travel, item.Loss, item.Gas, item.Etc, item.Danger, item.Machine, item.Print, item.Date)
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

func (p *StandardwageManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from standardwage_tb where sw_id = ")
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

func (p *StandardwageManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from standardwage_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *StandardwageManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and sw_")
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

func (p *StandardwageManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from standardwage_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *StandardwageManager) Update(item *Standardwage) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update standardwage_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" sw_person1 = $1, sw_person2 = $2, sw_person3 = $3, sw_person4 = $4, sw_person5 = $5, sw_person6 = $6, sw_person7 = $7, sw_person8 = $8, sw_person9 = $9, sw_person10 = $10, sw_techprice1 = $11, sw_techprice2 = $12, sw_techprice3 = $13, sw_techprice4 = $14, sw_financialprice1 = $15, sw_financialprice2 = $16, sw_financialprice3 = $17, sw_financialprice4 = $18, sw_directprice = $19, sw_printprice1 = $20, sw_printprice2 = $21, sw_lossprice = $22, sw_gasprice = $23, sw_travelprice = $24, sw_travel = $25, sw_loss = $26, sw_gas = $27, sw_etc = $28, sw_danger = $29, sw_machine = $30, sw_print = $31, sw_date = $32 where sw_id = $33")
    } else {
        query.WriteString(" sw_person1 = ?, sw_person2 = ?, sw_person3 = ?, sw_person4 = ?, sw_person5 = ?, sw_person6 = ?, sw_person7 = ?, sw_person8 = ?, sw_person9 = ?, sw_person10 = ?, sw_techprice1 = ?, sw_techprice2 = ?, sw_techprice3 = ?, sw_techprice4 = ?, sw_financialprice1 = ?, sw_financialprice2 = ?, sw_financialprice3 = ?, sw_financialprice4 = ?, sw_directprice = ?, sw_printprice1 = ?, sw_printprice2 = ?, sw_lossprice = ?, sw_gasprice = ?, sw_travelprice = ?, sw_travel = ?, sw_loss = ?, sw_gas = ?, sw_etc = ?, sw_danger = ?, sw_machine = ?, sw_print = ?, sw_date = ? where sw_id = ?")
    }

	_, err := p.Exec(query.String() , item.Person1, item.Person2, item.Person3, item.Person4, item.Person5, item.Person6, item.Person7, item.Person8, item.Person9, item.Person10, item.Techprice1, item.Techprice2, item.Techprice3, item.Techprice4, item.Financialprice1, item.Financialprice2, item.Financialprice3, item.Financialprice4, item.Directprice, item.Printprice1, item.Printprice2, item.Lossprice, item.Gasprice, item.Travelprice, item.Travel, item.Loss, item.Gas, item.Etc, item.Danger, item.Machine, item.Print, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *StandardwageManager) UpdateWhere(columns []standardwage.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update standardwage_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == standardwage.ColumnId {
        initQuery.WriteString("sw_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson1 {
        initQuery.WriteString("sw_person1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson2 {
        initQuery.WriteString("sw_person2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson3 {
        initQuery.WriteString("sw_person3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson4 {
        initQuery.WriteString("sw_person4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson5 {
        initQuery.WriteString("sw_person5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson6 {
        initQuery.WriteString("sw_person6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson7 {
        initQuery.WriteString("sw_person7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson8 {
        initQuery.WriteString("sw_person8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson9 {
        initQuery.WriteString("sw_person9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPerson10 {
        initQuery.WriteString("sw_person10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnTechprice1 {
        initQuery.WriteString("sw_techprice1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnTechprice2 {
        initQuery.WriteString("sw_techprice2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnTechprice3 {
        initQuery.WriteString("sw_techprice3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnTechprice4 {
        initQuery.WriteString("sw_techprice4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnFinancialprice1 {
        initQuery.WriteString("sw_financialprice1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnFinancialprice2 {
        initQuery.WriteString("sw_financialprice2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnFinancialprice3 {
        initQuery.WriteString("sw_financialprice3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnFinancialprice4 {
        initQuery.WriteString("sw_financialprice4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnDirectprice {
        initQuery.WriteString("sw_directprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPrintprice1 {
        initQuery.WriteString("sw_printprice1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPrintprice2 {
        initQuery.WriteString("sw_printprice2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnLossprice {
        initQuery.WriteString("sw_lossprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnGasprice {
        initQuery.WriteString("sw_gasprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnTravelprice {
        initQuery.WriteString("sw_travelprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnTravel {
        initQuery.WriteString("sw_travel = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnLoss {
        initQuery.WriteString("sw_loss = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnGas {
        initQuery.WriteString("sw_gas = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnEtc {
        initQuery.WriteString("sw_etc = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnDanger {
        initQuery.WriteString("sw_danger = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnMachine {
        initQuery.WriteString("sw_machine = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnPrint {
        initQuery.WriteString("sw_print = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == standardwage.ColumnDate {
        initQuery.WriteString("sw_date = ?")
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

func (p *StandardwageManager) UpdatePerson1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person1 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person2 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person3 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person4 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person5 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person6 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person7 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person8 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person9 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePerson10(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person10 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateTechprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice1 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateTechprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice2 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateTechprice3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice3 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateTechprice4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice4 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateFinancialprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice1 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateFinancialprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice2 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateFinancialprice3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice3 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateFinancialprice4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice4 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateDirectprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_directprice = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePrintprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_printprice1 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePrintprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_printprice2 = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateLossprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_lossprice = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateGasprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_gasprice = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateTravelprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_travelprice = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateTravel(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_travel = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateLoss(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_loss = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateGas(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_gas = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateEtc(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_etc = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateDanger(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_danger = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdateMachine(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_machine = ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) UpdatePrint(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_print = ? where sw_id = ?"
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

func (p *StandardwageManager) IncreasePerson1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person1 = sw_person1 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person2 = sw_person2 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person3 = sw_person3 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person4 = sw_person4 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person5 = sw_person5 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person6 = sw_person6 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person7 = sw_person7 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person8 = sw_person8 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person9 = sw_person9 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePerson10(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_person10 = sw_person10 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseTechprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice1 = sw_techprice1 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseTechprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice2 = sw_techprice2 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseTechprice3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice3 = sw_techprice3 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseTechprice4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_techprice4 = sw_techprice4 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseFinancialprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice1 = sw_financialprice1 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseFinancialprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice2 = sw_financialprice2 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseFinancialprice3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice3 = sw_financialprice3 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseFinancialprice4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_financialprice4 = sw_financialprice4 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseDirectprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_directprice = sw_directprice + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePrintprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_printprice1 = sw_printprice1 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePrintprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_printprice2 = sw_printprice2 + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseLossprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_lossprice = sw_lossprice + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseGasprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_gasprice = sw_gasprice + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseTravelprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_travelprice = sw_travelprice + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseTravel(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_travel = sw_travel + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseLoss(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_loss = sw_loss + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseGas(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_gas = sw_gas + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseEtc(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_etc = sw_etc + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseDanger(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_danger = sw_danger + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreaseMachine(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_machine = sw_machine + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *StandardwageManager) IncreasePrint(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update standardwage_tb set sw_print = sw_print + ? where sw_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *StandardwageManager) GetIdentity() int64 {
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

func (p *Standardwage) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *StandardwageManager) ReadRow(rows *sql.Rows) *Standardwage {
    var item Standardwage
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Person1, &item.Person2, &item.Person3, &item.Person4, &item.Person5, &item.Person6, &item.Person7, &item.Person8, &item.Person9, &item.Person10, &item.Techprice1, &item.Techprice2, &item.Techprice3, &item.Techprice4, &item.Financialprice1, &item.Financialprice2, &item.Financialprice3, &item.Financialprice4, &item.Directprice, &item.Printprice1, &item.Printprice2, &item.Lossprice, &item.Gasprice, &item.Travelprice, &item.Travel, &item.Loss, &item.Gas, &item.Etc, &item.Danger, &item.Machine, &item.Print, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *StandardwageManager) ReadRows(rows *sql.Rows) []Standardwage {
    items := make([]Standardwage, 0)

    for rows.Next() {
        var item Standardwage
        
    
        err := rows.Scan(&item.Id, &item.Person1, &item.Person2, &item.Person3, &item.Person4, &item.Person5, &item.Person6, &item.Person7, &item.Person8, &item.Person9, &item.Person10, &item.Techprice1, &item.Techprice2, &item.Techprice3, &item.Techprice4, &item.Financialprice1, &item.Financialprice2, &item.Financialprice3, &item.Financialprice4, &item.Directprice, &item.Printprice1, &item.Printprice2, &item.Lossprice, &item.Gasprice, &item.Travelprice, &item.Travel, &item.Loss, &item.Gas, &item.Etc, &item.Danger, &item.Machine, &item.Print, &item.Date)
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

func (p *StandardwageManager) Get(id int64) *Standardwage {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and sw_id = ?")

    
    
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

func (p *StandardwageManager) GetWhere(args []any) *Standardwage {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *StandardwageManager) Count(args []any) int {
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

func (p *StandardwageManager) FindAll() []Standardwage {
    return p.Find(nil)
}

func (p *StandardwageManager) Find(args []any) []Standardwage {
    if !p.Conn.IsConnect() {
        items := make([]Standardwage, 0)
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
                query.WriteString(" and sw_")
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
            orderby = "sw_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "sw_" + orderby
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
            orderby = "sw_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "sw_" + orderby
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
        items := make([]Standardwage, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *StandardwageManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and sw_")
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
    
    query.WriteString(" group by sw_")
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



func (p *StandardwageManager) MakeMap(items []Standardwage) map[int64]Standardwage {
     ret := make(map[int64]Standardwage)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *StandardwageManager) FindToMap(args []any) map[int64]Standardwage {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *StandardwageManager) FindAllToMap() map[int64]Standardwage {
     items := p.Find(nil)
     return p.MakeMap(items)
}



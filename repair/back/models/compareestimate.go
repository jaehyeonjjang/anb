package models

import (
    "repair/models/compareestimate"
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

type Compareestimate struct {
            
    Id                int64 `json:"id"`         
    Type                int `json:"type"`         
    Subtype                int `json:"subtype"`         
    Originalprice                int `json:"originalprice"`         
    Saleprice                int `json:"saleprice"`         
    Price                int `json:"price"`         
    Financialprice                int `json:"financialprice"`         
    Techprice                int `json:"techprice"`         
    Directprice                int `json:"directprice"`         
    Printprice                int `json:"printprice"`         
    Extraprice                int `json:"extraprice"`         
    Travelprice                int `json:"travelprice"`         
    Lossprice                int `json:"lossprice"`         
    Gasprice                int `json:"gasprice"`         
    Etcprice                int `json:"etcprice"`         
    Dangerprice                int `json:"dangerprice"`         
    Machineprice                int `json:"machineprice"`         
    Carprice                int `json:"carprice"`         
    Discount                int `json:"discount"`         
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
    Personprice1                int `json:"personprice1"`         
    Personprice2                int `json:"personprice2"`         
    Personprice3                int `json:"personprice3"`         
    Personprice4                int `json:"personprice4"`         
    Personprice5                int `json:"personprice5"`         
    Personprice6                int `json:"personprice6"`         
    Personprice7                int `json:"personprice7"`         
    Personprice8                int `json:"personprice8"`         
    Personprice9                int `json:"personprice9"`         
    Personprice10                int `json:"personprice10"`         
    Travel                int `json:"travel"`         
    Loss                int `json:"loss"`         
    Gas                int `json:"gas"`         
    Etc                int `json:"etc"`         
    Danger                int `json:"danger"`         
    Machine                int `json:"machine"`         
    Car                int `json:"car"`         
    Print                int `json:"print"`         
    Writedate                string `json:"writedate"`         
    Start                string `json:"start"`         
    Remark                string `json:"remark"`         
    Adjust                int `json:"adjust"`         
    User                int64 `json:"user"`         
    Comparecompany                int64 `json:"comparecompany"`         
    Estimate                int64 `json:"estimate"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type CompareestimateManager struct {
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



func (c *Compareestimate) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewCompareestimateManager(conn *Connection) *CompareestimateManager {
    var item CompareestimateManager


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

func (p *CompareestimateManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CompareestimateManager) SetIndex(index string) {
    p.Index = index
}

func (p *CompareestimateManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *CompareestimateManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *CompareestimateManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *CompareestimateManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *CompareestimateManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select e_id, e_type, e_subtype, e_originalprice, e_saleprice, e_price, e_financialprice, e_techprice, e_directprice, e_printprice, e_extraprice, e_travelprice, e_lossprice, e_gasprice, e_etcprice, e_dangerprice, e_machineprice, e_carprice, e_discount, e_person1, e_person2, e_person3, e_person4, e_person5, e_person6, e_person7, e_person8, e_person9, e_person10, e_personprice1, e_personprice2, e_personprice3, e_personprice4, e_personprice5, e_personprice6, e_personprice7, e_personprice8, e_personprice9, e_personprice10, e_travel, e_loss, e_gas, e_etc, e_danger, e_machine, e_car, e_print, e_writedate, e_start, e_remark, e_adjust, e_user, e_comparecompany, e_estimate, e_apt, e_date, cc_id, cc_name, cc_address, cc_addressetc, cc_tel, cc_fax, cc_ceo, cc_format, cc_image, cc_image2, cc_adjust, cc_financialprice, cc_techprice, cc_directprice, cc_printprice, cc_extraprice, cc_travelprice, cc_gasprice, cc_dangerprice, cc_machineprice, cc_remark, cc_type, cc_default, cc_order, cc_date from compareestimate_tb, comparecompany_tb")

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
    
    ret.WriteString(" and e_comparecompany = cc_id ")
    

    return ret.String()
}

func (p *CompareestimateManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from compareestimate_tb, comparecompany_tb")

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
    
    ret.WriteString(" and e_comparecompany = cc_id ")
    

    return ret.String()
}

func (p *CompareestimateManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select e_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from compareestimate_tb, comparecompany_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and e_comparecompany = cc_id ")
    


    return ret.String()
}

func (p *CompareestimateManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate compareestimate_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *CompareestimateManager) Insert(item *Compareestimate) error {
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
          query = "insert into compareestimate_tb (e_id, e_type, e_subtype, e_originalprice, e_saleprice, e_price, e_financialprice, e_techprice, e_directprice, e_printprice, e_extraprice, e_travelprice, e_lossprice, e_gasprice, e_etcprice, e_dangerprice, e_machineprice, e_carprice, e_discount, e_person1, e_person2, e_person3, e_person4, e_person5, e_person6, e_person7, e_person8, e_person9, e_person10, e_personprice1, e_personprice2, e_personprice3, e_personprice4, e_personprice5, e_personprice6, e_personprice7, e_personprice8, e_personprice9, e_personprice10, e_travel, e_loss, e_gas, e_etc, e_danger, e_machine, e_car, e_print, e_writedate, e_start, e_remark, e_adjust, e_user, e_comparecompany, e_estimate, e_apt, e_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56)"
        } else {
          query = "insert into compareestimate_tb (e_id, e_type, e_subtype, e_originalprice, e_saleprice, e_price, e_financialprice, e_techprice, e_directprice, e_printprice, e_extraprice, e_travelprice, e_lossprice, e_gasprice, e_etcprice, e_dangerprice, e_machineprice, e_carprice, e_discount, e_person1, e_person2, e_person3, e_person4, e_person5, e_person6, e_person7, e_person8, e_person9, e_person10, e_personprice1, e_personprice2, e_personprice3, e_personprice4, e_personprice5, e_personprice6, e_personprice7, e_personprice8, e_personprice9, e_personprice10, e_travel, e_loss, e_gas, e_etc, e_danger, e_machine, e_car, e_print, e_writedate, e_start, e_remark, e_adjust, e_user, e_comparecompany, e_estimate, e_apt, e_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Type, item.Subtype, item.Originalprice, item.Saleprice, item.Price, item.Financialprice, item.Techprice, item.Directprice, item.Printprice, item.Extraprice, item.Travelprice, item.Lossprice, item.Gasprice, item.Etcprice, item.Dangerprice, item.Machineprice, item.Carprice, item.Discount, item.Person1, item.Person2, item.Person3, item.Person4, item.Person5, item.Person6, item.Person7, item.Person8, item.Person9, item.Person10, item.Personprice1, item.Personprice2, item.Personprice3, item.Personprice4, item.Personprice5, item.Personprice6, item.Personprice7, item.Personprice8, item.Personprice9, item.Personprice10, item.Travel, item.Loss, item.Gas, item.Etc, item.Danger, item.Machine, item.Car, item.Print, item.Writedate, item.Start, item.Remark, item.Adjust, item.User, item.Comparecompany, item.Estimate, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into compareestimate_tb (e_type, e_subtype, e_originalprice, e_saleprice, e_price, e_financialprice, e_techprice, e_directprice, e_printprice, e_extraprice, e_travelprice, e_lossprice, e_gasprice, e_etcprice, e_dangerprice, e_machineprice, e_carprice, e_discount, e_person1, e_person2, e_person3, e_person4, e_person5, e_person6, e_person7, e_person8, e_person9, e_person10, e_personprice1, e_personprice2, e_personprice3, e_personprice4, e_personprice5, e_personprice6, e_personprice7, e_personprice8, e_personprice9, e_personprice10, e_travel, e_loss, e_gas, e_etc, e_danger, e_machine, e_car, e_print, e_writedate, e_start, e_remark, e_adjust, e_user, e_comparecompany, e_estimate, e_apt, e_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)"
        } else {
          query = "insert into compareestimate_tb (e_type, e_subtype, e_originalprice, e_saleprice, e_price, e_financialprice, e_techprice, e_directprice, e_printprice, e_extraprice, e_travelprice, e_lossprice, e_gasprice, e_etcprice, e_dangerprice, e_machineprice, e_carprice, e_discount, e_person1, e_person2, e_person3, e_person4, e_person5, e_person6, e_person7, e_person8, e_person9, e_person10, e_personprice1, e_personprice2, e_personprice3, e_personprice4, e_personprice5, e_personprice6, e_personprice7, e_personprice8, e_personprice9, e_personprice10, e_travel, e_loss, e_gas, e_etc, e_danger, e_machine, e_car, e_print, e_writedate, e_start, e_remark, e_adjust, e_user, e_comparecompany, e_estimate, e_apt, e_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Type, item.Subtype, item.Originalprice, item.Saleprice, item.Price, item.Financialprice, item.Techprice, item.Directprice, item.Printprice, item.Extraprice, item.Travelprice, item.Lossprice, item.Gasprice, item.Etcprice, item.Dangerprice, item.Machineprice, item.Carprice, item.Discount, item.Person1, item.Person2, item.Person3, item.Person4, item.Person5, item.Person6, item.Person7, item.Person8, item.Person9, item.Person10, item.Personprice1, item.Personprice2, item.Personprice3, item.Personprice4, item.Personprice5, item.Personprice6, item.Personprice7, item.Personprice8, item.Personprice9, item.Personprice10, item.Travel, item.Loss, item.Gas, item.Etc, item.Danger, item.Machine, item.Car, item.Print, item.Writedate, item.Start, item.Remark, item.Adjust, item.User, item.Comparecompany, item.Estimate, item.Apt, item.Date)
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

func (p *CompareestimateManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from compareestimate_tb where e_id = ")
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

func (p *CompareestimateManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from compareestimate_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *CompareestimateManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and e_")
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

func (p *CompareestimateManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from compareestimate_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *CompareestimateManager) Update(item *Compareestimate) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update compareestimate_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" e_type = $1, e_subtype = $2, e_originalprice = $3, e_saleprice = $4, e_price = $5, e_financialprice = $6, e_techprice = $7, e_directprice = $8, e_printprice = $9, e_extraprice = $10, e_travelprice = $11, e_lossprice = $12, e_gasprice = $13, e_etcprice = $14, e_dangerprice = $15, e_machineprice = $16, e_carprice = $17, e_discount = $18, e_person1 = $19, e_person2 = $20, e_person3 = $21, e_person4 = $22, e_person5 = $23, e_person6 = $24, e_person7 = $25, e_person8 = $26, e_person9 = $27, e_person10 = $28, e_personprice1 = $29, e_personprice2 = $30, e_personprice3 = $31, e_personprice4 = $32, e_personprice5 = $33, e_personprice6 = $34, e_personprice7 = $35, e_personprice8 = $36, e_personprice9 = $37, e_personprice10 = $38, e_travel = $39, e_loss = $40, e_gas = $41, e_etc = $42, e_danger = $43, e_machine = $44, e_car = $45, e_print = $46, e_writedate = $47, e_start = $48, e_remark = $49, e_adjust = $50, e_user = $51, e_comparecompany = $52, e_estimate = $53, e_apt = $54, e_date = $55 where e_id = $56")
    } else {
        query.WriteString(" e_type = ?, e_subtype = ?, e_originalprice = ?, e_saleprice = ?, e_price = ?, e_financialprice = ?, e_techprice = ?, e_directprice = ?, e_printprice = ?, e_extraprice = ?, e_travelprice = ?, e_lossprice = ?, e_gasprice = ?, e_etcprice = ?, e_dangerprice = ?, e_machineprice = ?, e_carprice = ?, e_discount = ?, e_person1 = ?, e_person2 = ?, e_person3 = ?, e_person4 = ?, e_person5 = ?, e_person6 = ?, e_person7 = ?, e_person8 = ?, e_person9 = ?, e_person10 = ?, e_personprice1 = ?, e_personprice2 = ?, e_personprice3 = ?, e_personprice4 = ?, e_personprice5 = ?, e_personprice6 = ?, e_personprice7 = ?, e_personprice8 = ?, e_personprice9 = ?, e_personprice10 = ?, e_travel = ?, e_loss = ?, e_gas = ?, e_etc = ?, e_danger = ?, e_machine = ?, e_car = ?, e_print = ?, e_writedate = ?, e_start = ?, e_remark = ?, e_adjust = ?, e_user = ?, e_comparecompany = ?, e_estimate = ?, e_apt = ?, e_date = ? where e_id = ?")
    }

	_, err := p.Exec(query.String() , item.Type, item.Subtype, item.Originalprice, item.Saleprice, item.Price, item.Financialprice, item.Techprice, item.Directprice, item.Printprice, item.Extraprice, item.Travelprice, item.Lossprice, item.Gasprice, item.Etcprice, item.Dangerprice, item.Machineprice, item.Carprice, item.Discount, item.Person1, item.Person2, item.Person3, item.Person4, item.Person5, item.Person6, item.Person7, item.Person8, item.Person9, item.Person10, item.Personprice1, item.Personprice2, item.Personprice3, item.Personprice4, item.Personprice5, item.Personprice6, item.Personprice7, item.Personprice8, item.Personprice9, item.Personprice10, item.Travel, item.Loss, item.Gas, item.Etc, item.Danger, item.Machine, item.Car, item.Print, item.Writedate, item.Start, item.Remark, item.Adjust, item.User, item.Comparecompany, item.Estimate, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *CompareestimateManager) UpdateWhere(columns []compareestimate.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update compareestimate_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == compareestimate.ColumnId {
        initQuery.WriteString("e_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnType {
        initQuery.WriteString("e_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnSubtype {
        initQuery.WriteString("e_subtype = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnOriginalprice {
        initQuery.WriteString("e_originalprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnSaleprice {
        initQuery.WriteString("e_saleprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPrice {
        initQuery.WriteString("e_price = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnFinancialprice {
        initQuery.WriteString("e_financialprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnTechprice {
        initQuery.WriteString("e_techprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnDirectprice {
        initQuery.WriteString("e_directprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPrintprice {
        initQuery.WriteString("e_printprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnExtraprice {
        initQuery.WriteString("e_extraprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnTravelprice {
        initQuery.WriteString("e_travelprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnLossprice {
        initQuery.WriteString("e_lossprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnGasprice {
        initQuery.WriteString("e_gasprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnEtcprice {
        initQuery.WriteString("e_etcprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnDangerprice {
        initQuery.WriteString("e_dangerprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnMachineprice {
        initQuery.WriteString("e_machineprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnCarprice {
        initQuery.WriteString("e_carprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnDiscount {
        initQuery.WriteString("e_discount = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson1 {
        initQuery.WriteString("e_person1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson2 {
        initQuery.WriteString("e_person2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson3 {
        initQuery.WriteString("e_person3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson4 {
        initQuery.WriteString("e_person4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson5 {
        initQuery.WriteString("e_person5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson6 {
        initQuery.WriteString("e_person6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson7 {
        initQuery.WriteString("e_person7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson8 {
        initQuery.WriteString("e_person8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson9 {
        initQuery.WriteString("e_person9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPerson10 {
        initQuery.WriteString("e_person10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice1 {
        initQuery.WriteString("e_personprice1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice2 {
        initQuery.WriteString("e_personprice2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice3 {
        initQuery.WriteString("e_personprice3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice4 {
        initQuery.WriteString("e_personprice4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice5 {
        initQuery.WriteString("e_personprice5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice6 {
        initQuery.WriteString("e_personprice6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice7 {
        initQuery.WriteString("e_personprice7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice8 {
        initQuery.WriteString("e_personprice8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice9 {
        initQuery.WriteString("e_personprice9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPersonprice10 {
        initQuery.WriteString("e_personprice10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnTravel {
        initQuery.WriteString("e_travel = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnLoss {
        initQuery.WriteString("e_loss = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnGas {
        initQuery.WriteString("e_gas = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnEtc {
        initQuery.WriteString("e_etc = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnDanger {
        initQuery.WriteString("e_danger = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnMachine {
        initQuery.WriteString("e_machine = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnCar {
        initQuery.WriteString("e_car = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnPrint {
        initQuery.WriteString("e_print = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnWritedate {
        initQuery.WriteString("e_writedate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnStart {
        initQuery.WriteString("e_start = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnRemark {
        initQuery.WriteString("e_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnAdjust {
        initQuery.WriteString("e_adjust = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnUser {
        initQuery.WriteString("e_user = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnComparecompany {
        initQuery.WriteString("e_comparecompany = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnEstimate {
        initQuery.WriteString("e_estimate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnApt {
        initQuery.WriteString("e_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == compareestimate.ColumnDate {
        initQuery.WriteString("e_date = ?")
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

func (p *CompareestimateManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_type = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateSubtype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_subtype = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateOriginalprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_originalprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateSaleprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_saleprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePrice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_price = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateFinancialprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_financialprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateTechprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_techprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateDirectprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_directprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePrintprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_printprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateExtraprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_extraprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateTravelprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_travelprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateLossprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_lossprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateGasprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_gasprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateEtcprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_etcprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateDangerprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_dangerprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateMachineprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_machineprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateCarprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_carprice = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateDiscount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_discount = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person1 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person2 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person3 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person4 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person5 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person6 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person7 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person8 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person9 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePerson10(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person10 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice1 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice2 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice3 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice4 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice5 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice6 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice7 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice8 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice9 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePersonprice10(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice10 = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateTravel(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_travel = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateLoss(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_loss = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateGas(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_gas = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateEtc(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_etc = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateDanger(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_danger = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateMachine(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_machine = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateCar(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_car = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdatePrint(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_print = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateWritedate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_writedate = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateStart(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_start = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_remark = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateAdjust(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_adjust = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_user = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateComparecompany(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_comparecompany = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateEstimate(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_estimate = ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_apt = ? where e_id = ?"
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

func (p *CompareestimateManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_type = e_type + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseSubtype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_subtype = e_subtype + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseOriginalprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_originalprice = e_originalprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseSaleprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_saleprice = e_saleprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePrice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_price = e_price + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseFinancialprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_financialprice = e_financialprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseTechprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_techprice = e_techprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseDirectprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_directprice = e_directprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePrintprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_printprice = e_printprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseExtraprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_extraprice = e_extraprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseTravelprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_travelprice = e_travelprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseLossprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_lossprice = e_lossprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseGasprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_gasprice = e_gasprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseEtcprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_etcprice = e_etcprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseDangerprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_dangerprice = e_dangerprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseMachineprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_machineprice = e_machineprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseCarprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_carprice = e_carprice + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseDiscount(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_discount = e_discount + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person1 = e_person1 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person2 = e_person2 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person3 = e_person3 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person4 = e_person4 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person5 = e_person5 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person6 = e_person6 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person7 = e_person7 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person8 = e_person8 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person9 = e_person9 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePerson10(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_person10 = e_person10 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice1 = e_personprice1 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice2 = e_personprice2 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice3 = e_personprice3 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice4 = e_personprice4 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice5 = e_personprice5 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice6 = e_personprice6 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice7 = e_personprice7 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice8 = e_personprice8 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice9 = e_personprice9 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePersonprice10(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_personprice10 = e_personprice10 + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseTravel(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_travel = e_travel + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseLoss(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_loss = e_loss + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseGas(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_gas = e_gas + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseEtc(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_etc = e_etc + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseDanger(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_danger = e_danger + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseMachine(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_machine = e_machine + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseCar(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_car = e_car + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreasePrint(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_print = e_print + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseAdjust(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_adjust = e_adjust + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_user = e_user + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseComparecompany(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_comparecompany = e_comparecompany + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseEstimate(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_estimate = e_estimate + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *CompareestimateManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update compareestimate_tb set e_apt = e_apt + ? where e_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *CompareestimateManager) GetIdentity() int64 {
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

func (p *Compareestimate) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *CompareestimateManager) ReadRow(rows *sql.Rows) *Compareestimate {
    var item Compareestimate
    var err error

    var _comparecompany Comparecompany
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Type, &item.Subtype, &item.Originalprice, &item.Saleprice, &item.Price, &item.Financialprice, &item.Techprice, &item.Directprice, &item.Printprice, &item.Extraprice, &item.Travelprice, &item.Lossprice, &item.Gasprice, &item.Etcprice, &item.Dangerprice, &item.Machineprice, &item.Carprice, &item.Discount, &item.Person1, &item.Person2, &item.Person3, &item.Person4, &item.Person5, &item.Person6, &item.Person7, &item.Person8, &item.Person9, &item.Person10, &item.Personprice1, &item.Personprice2, &item.Personprice3, &item.Personprice4, &item.Personprice5, &item.Personprice6, &item.Personprice7, &item.Personprice8, &item.Personprice9, &item.Personprice10, &item.Travel, &item.Loss, &item.Gas, &item.Etc, &item.Danger, &item.Machine, &item.Car, &item.Print, &item.Writedate, &item.Start, &item.Remark, &item.Adjust, &item.User, &item.Comparecompany, &item.Estimate, &item.Apt, &item.Date, &_comparecompany.Id, &_comparecompany.Name, &_comparecompany.Address, &_comparecompany.Addressetc, &_comparecompany.Tel, &_comparecompany.Fax, &_comparecompany.Ceo, &_comparecompany.Format, &_comparecompany.Image, &_comparecompany.Image2, &_comparecompany.Adjust, &_comparecompany.Financialprice, &_comparecompany.Techprice, &_comparecompany.Directprice, &_comparecompany.Printprice, &_comparecompany.Extraprice, &_comparecompany.Travelprice, &_comparecompany.Gasprice, &_comparecompany.Dangerprice, &_comparecompany.Machineprice, &_comparecompany.Remark, &_comparecompany.Type, &_comparecompany.Default, &_comparecompany.Order, &_comparecompany.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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
        _comparecompany.InitExtra()
        item.AddExtra("comparecompany",  _comparecompany)

        return &item
    }
}

func (p *CompareestimateManager) ReadRows(rows *sql.Rows) []Compareestimate {
    items := make([]Compareestimate, 0)

    for rows.Next() {
        var item Compareestimate
        var _comparecompany Comparecompany
            
    
        err := rows.Scan(&item.Id, &item.Type, &item.Subtype, &item.Originalprice, &item.Saleprice, &item.Price, &item.Financialprice, &item.Techprice, &item.Directprice, &item.Printprice, &item.Extraprice, &item.Travelprice, &item.Lossprice, &item.Gasprice, &item.Etcprice, &item.Dangerprice, &item.Machineprice, &item.Carprice, &item.Discount, &item.Person1, &item.Person2, &item.Person3, &item.Person4, &item.Person5, &item.Person6, &item.Person7, &item.Person8, &item.Person9, &item.Person10, &item.Personprice1, &item.Personprice2, &item.Personprice3, &item.Personprice4, &item.Personprice5, &item.Personprice6, &item.Personprice7, &item.Personprice8, &item.Personprice9, &item.Personprice10, &item.Travel, &item.Loss, &item.Gas, &item.Etc, &item.Danger, &item.Machine, &item.Car, &item.Print, &item.Writedate, &item.Start, &item.Remark, &item.Adjust, &item.User, &item.Comparecompany, &item.Estimate, &item.Apt, &item.Date, &_comparecompany.Id, &_comparecompany.Name, &_comparecompany.Address, &_comparecompany.Addressetc, &_comparecompany.Tel, &_comparecompany.Fax, &_comparecompany.Ceo, &_comparecompany.Format, &_comparecompany.Image, &_comparecompany.Image2, &_comparecompany.Adjust, &_comparecompany.Financialprice, &_comparecompany.Techprice, &_comparecompany.Directprice, &_comparecompany.Printprice, &_comparecompany.Extraprice, &_comparecompany.Travelprice, &_comparecompany.Gasprice, &_comparecompany.Dangerprice, &_comparecompany.Machineprice, &_comparecompany.Remark, &_comparecompany.Type, &_comparecompany.Default, &_comparecompany.Order, &_comparecompany.Date)
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
        _comparecompany.InitExtra()
        item.AddExtra("comparecompany",  _comparecompany)

        items = append(items, item)
    }


     return items
}

func (p *CompareestimateManager) Get(id int64) *Compareestimate {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and e_id = ?")

    
    query.WriteString(" and e_comparecompany = cc_id ")
    
    
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

func (p *CompareestimateManager) GetWhere(args []any) *Compareestimate {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *CompareestimateManager) Count(args []any) int {
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

func (p *CompareestimateManager) FindAll() []Compareestimate {
    return p.Find(nil)
}

func (p *CompareestimateManager) Find(args []any) []Compareestimate {
    if !p.Conn.IsConnect() {
        items := make([]Compareestimate, 0)
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
                query.WriteString(" and e_")
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
            orderby = "e_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "e_" + orderby
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
            orderby = "e_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "e_" + orderby
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
        items := make([]Compareestimate, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *CompareestimateManager) FindByEstimate(estimate int64, args ...any) []Compareestimate {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if estimate != 0 { 
        rets = append(rets, Where{Column:"estimate", Value:estimate, Compare:"="})
     }
    
    
    return p.Find(rets)
}

func (p *CompareestimateManager) DeleteByEstimate(estimate int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from compareestimate_tb where e_estimate = ?"
    _, err := p.Exec(query, estimate)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}


func (p *CompareestimateManager) Sum(args []any) *Compareestimate {
    if !p.Conn.IsConnect() {
        var item Compareestimate
        return &item
    }

    var params []any

    
    var query strings.Builder
    query.WriteString("select sum(e_price) from compareestimate_tb")

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
                query.WriteString(" and e_")
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
            orderby = "e_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "e_" + orderby
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
            orderby = "e_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "e_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(query.String(), params...)

    var item Compareestimate
    
    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        err := rows.Scan(&item.Price)        
        if err != nil {
            if p.Log {
                log.Error().Str("error", err.Error()).Msg("SQL")
            }

            return &item
        }
    }

    return &item        
}

func (p *CompareestimateManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and e_")
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
    
    query.WriteString(" group by e_")
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



func (p *CompareestimateManager) MakeMap(items []Compareestimate) map[int64]Compareestimate {
     ret := make(map[int64]Compareestimate)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *CompareestimateManager) FindToMap(args []any) map[int64]Compareestimate {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *CompareestimateManager) FindAllToMap() map[int64]Compareestimate {
     items := p.Find(nil)
     return p.MakeMap(items)
}



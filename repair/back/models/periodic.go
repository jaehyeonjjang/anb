package models

import (
    "repair/models/periodic"
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

type Periodic struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Aptname                string `json:"aptname"`         
    Taskrange                string `json:"taskrange"`         
    Reportdate                string `json:"reportdate"`         
    Startdate                string `json:"startdate"`         
    Enddate                string `json:"enddate"`         
    Supply                string `json:"supply"`         
    Contract                string `json:"contract"`         
    Price                string `json:"price"`         
    Safetygrade                string `json:"safetygrade"`         
    Status                int `json:"status"`         
    Prestartdate                string `json:"prestartdate"`         
    Preenddate                string `json:"preenddate"`         
    Researchstartdate                string `json:"researchstartdate"`         
    Researchenddate                string `json:"researchenddate"`         
    Analyzestartdate                string `json:"analyzestartdate"`         
    Analyzeenddate                string `json:"analyzeenddate"`         
    Ratingstartdate                string `json:"ratingstartdate"`         
    Ratingenddate                string `json:"ratingenddate"`         
    Writestartdate                string `json:"writestartdate"`         
    Writeenddate                string `json:"writeenddate"`         
    Printstartdate                string `json:"printstartdate"`         
    Printenddate                string `json:"printenddate"`         
    Blueprint1                int `json:"blueprint1"`         
    Blueprint2                int `json:"blueprint2"`         
    Blueprint3                int `json:"blueprint3"`         
    Blueprint4                int `json:"blueprint4"`         
    Blueprint5                int `json:"blueprint5"`         
    Blueprint6                int `json:"blueprint6"`         
    Blueprint7                int `json:"blueprint7"`         
    Blueprint8                int `json:"blueprint8"`         
    Blueprint9                int `json:"blueprint9"`         
    Blueprint10                string `json:"blueprint10"`         
    Blueprint11                int `json:"blueprint11"`         
    Blueprint1save                int `json:"blueprint1save"`         
    Owner                string `json:"owner"`         
    Manager                string `json:"manager"`         
    Agent                string `json:"agent"`         
    Result1                int `json:"result1"`         
    Result2                int `json:"result2"`         
    Result3                int `json:"result3"`         
    Result4                int `json:"result4"`         
    Result5                int `json:"result5"`         
    Resulttext1                string `json:"resulttext1"`         
    Resulttext2                string `json:"resulttext2"`         
    Resulttext3                string `json:"resulttext3"`         
    Resulttext4                string `json:"resulttext4"`         
    Resulttext5                string `json:"resulttext5"`         
    Past                string `json:"past"`         
    Category                int `json:"category"`         
    User                int64 `json:"user"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodicManager struct {
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



func (c *Periodic) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodicManager(conn *Connection) *PeriodicManager {
    var item PeriodicManager


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

func (p *PeriodicManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodicManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodicManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodicManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodicManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodicManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PeriodicManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select d_id, d_name, d_aptname, d_taskrange, d_reportdate, d_startdate, d_enddate, d_supply, d_contract, d_price, d_safetygrade, d_status, d_prestartdate, d_preenddate, d_researchstartdate, d_researchenddate, d_analyzestartdate, d_analyzeenddate, d_ratingstartdate, d_ratingenddate, d_writestartdate, d_writeenddate, d_printstartdate, d_printenddate, d_blueprint1, d_blueprint2, d_blueprint3, d_blueprint4, d_blueprint5, d_blueprint6, d_blueprint7, d_blueprint8, d_blueprint9, d_blueprint10, d_blueprint11, d_blueprint1save, d_owner, d_manager, d_agent, d_result1, d_result2, d_result3, d_result4, d_result5, d_resulttext1, d_resulttext2, d_resulttext3, d_resulttext4, d_resulttext5, d_past, d_category, d_user, d_apt, d_date, a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date from periodic_tb, apt_tb")

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
    
    ret.WriteString(" and d_apt = a_id ")
    

    return ret.String()
}

func (p *PeriodicManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodic_tb, apt_tb")

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
    
    ret.WriteString(" and d_apt = a_id ")
    

    return ret.String()
}

func (p *PeriodicManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select d_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodic_tb, apt_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and d_apt = a_id ")
    


    return ret.String()
}

func (p *PeriodicManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodic_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodicManager) Insert(item *Periodic) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
	
	
    if item.Reportdate == "" {
       item.Reportdate = "1000-01-01"
    }
	
    if item.Startdate == "" {
       item.Startdate = "1000-01-01"
    }
	
    if item.Enddate == "" {
       item.Enddate = "1000-01-01"
    }
	
	
	
	
	
	
    if item.Prestartdate == "" {
       item.Prestartdate = "1000-01-01"
    }
	
    if item.Preenddate == "" {
       item.Preenddate = "1000-01-01"
    }
	
    if item.Researchstartdate == "" {
       item.Researchstartdate = "1000-01-01"
    }
	
    if item.Researchenddate == "" {
       item.Researchenddate = "1000-01-01"
    }
	
    if item.Analyzestartdate == "" {
       item.Analyzestartdate = "1000-01-01"
    }
	
    if item.Analyzeenddate == "" {
       item.Analyzeenddate = "1000-01-01"
    }
	
    if item.Ratingstartdate == "" {
       item.Ratingstartdate = "1000-01-01"
    }
	
    if item.Ratingenddate == "" {
       item.Ratingenddate = "1000-01-01"
    }
	
    if item.Writestartdate == "" {
       item.Writestartdate = "1000-01-01"
    }
	
    if item.Writeenddate == "" {
       item.Writeenddate = "1000-01-01"
    }
	
    if item.Printstartdate == "" {
       item.Printstartdate = "1000-01-01"
    }
	
    if item.Printenddate == "" {
       item.Printenddate = "1000-01-01"
    }
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodic_tb (d_id, d_name, d_aptname, d_taskrange, d_reportdate, d_startdate, d_enddate, d_supply, d_contract, d_price, d_safetygrade, d_status, d_prestartdate, d_preenddate, d_researchstartdate, d_researchenddate, d_analyzestartdate, d_analyzeenddate, d_ratingstartdate, d_ratingenddate, d_writestartdate, d_writeenddate, d_printstartdate, d_printenddate, d_blueprint1, d_blueprint2, d_blueprint3, d_blueprint4, d_blueprint5, d_blueprint6, d_blueprint7, d_blueprint8, d_blueprint9, d_blueprint10, d_blueprint11, d_blueprint1save, d_owner, d_manager, d_agent, d_result1, d_result2, d_result3, d_result4, d_result5, d_resulttext1, d_resulttext2, d_resulttext3, d_resulttext4, d_resulttext5, d_past, d_category, d_user, d_apt, d_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54)"
        } else {
          query = "insert into periodic_tb (d_id, d_name, d_aptname, d_taskrange, d_reportdate, d_startdate, d_enddate, d_supply, d_contract, d_price, d_safetygrade, d_status, d_prestartdate, d_preenddate, d_researchstartdate, d_researchenddate, d_analyzestartdate, d_analyzeenddate, d_ratingstartdate, d_ratingenddate, d_writestartdate, d_writeenddate, d_printstartdate, d_printenddate, d_blueprint1, d_blueprint2, d_blueprint3, d_blueprint4, d_blueprint5, d_blueprint6, d_blueprint7, d_blueprint8, d_blueprint9, d_blueprint10, d_blueprint11, d_blueprint1save, d_owner, d_manager, d_agent, d_result1, d_result2, d_result3, d_result4, d_result5, d_resulttext1, d_resulttext2, d_resulttext3, d_resulttext4, d_resulttext5, d_past, d_category, d_user, d_apt, d_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Name, item.Aptname, item.Taskrange, item.Reportdate, item.Startdate, item.Enddate, item.Supply, item.Contract, item.Price, item.Safetygrade, item.Status, item.Prestartdate, item.Preenddate, item.Researchstartdate, item.Researchenddate, item.Analyzestartdate, item.Analyzeenddate, item.Ratingstartdate, item.Ratingenddate, item.Writestartdate, item.Writeenddate, item.Printstartdate, item.Printenddate, item.Blueprint1, item.Blueprint2, item.Blueprint3, item.Blueprint4, item.Blueprint5, item.Blueprint6, item.Blueprint7, item.Blueprint8, item.Blueprint9, item.Blueprint10, item.Blueprint11, item.Blueprint1save, item.Owner, item.Manager, item.Agent, item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Resulttext1, item.Resulttext2, item.Resulttext3, item.Resulttext4, item.Resulttext5, item.Past, item.Category, item.User, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodic_tb (d_name, d_aptname, d_taskrange, d_reportdate, d_startdate, d_enddate, d_supply, d_contract, d_price, d_safetygrade, d_status, d_prestartdate, d_preenddate, d_researchstartdate, d_researchenddate, d_analyzestartdate, d_analyzeenddate, d_ratingstartdate, d_ratingenddate, d_writestartdate, d_writeenddate, d_printstartdate, d_printenddate, d_blueprint1, d_blueprint2, d_blueprint3, d_blueprint4, d_blueprint5, d_blueprint6, d_blueprint7, d_blueprint8, d_blueprint9, d_blueprint10, d_blueprint11, d_blueprint1save, d_owner, d_manager, d_agent, d_result1, d_result2, d_result3, d_result4, d_result5, d_resulttext1, d_resulttext2, d_resulttext3, d_resulttext4, d_resulttext5, d_past, d_category, d_user, d_apt, d_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53)"
        } else {
          query = "insert into periodic_tb (d_name, d_aptname, d_taskrange, d_reportdate, d_startdate, d_enddate, d_supply, d_contract, d_price, d_safetygrade, d_status, d_prestartdate, d_preenddate, d_researchstartdate, d_researchenddate, d_analyzestartdate, d_analyzeenddate, d_ratingstartdate, d_ratingenddate, d_writestartdate, d_writeenddate, d_printstartdate, d_printenddate, d_blueprint1, d_blueprint2, d_blueprint3, d_blueprint4, d_blueprint5, d_blueprint6, d_blueprint7, d_blueprint8, d_blueprint9, d_blueprint10, d_blueprint11, d_blueprint1save, d_owner, d_manager, d_agent, d_result1, d_result2, d_result3, d_result4, d_result5, d_resulttext1, d_resulttext2, d_resulttext3, d_resulttext4, d_resulttext5, d_past, d_category, d_user, d_apt, d_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Name, item.Aptname, item.Taskrange, item.Reportdate, item.Startdate, item.Enddate, item.Supply, item.Contract, item.Price, item.Safetygrade, item.Status, item.Prestartdate, item.Preenddate, item.Researchstartdate, item.Researchenddate, item.Analyzestartdate, item.Analyzeenddate, item.Ratingstartdate, item.Ratingenddate, item.Writestartdate, item.Writeenddate, item.Printstartdate, item.Printenddate, item.Blueprint1, item.Blueprint2, item.Blueprint3, item.Blueprint4, item.Blueprint5, item.Blueprint6, item.Blueprint7, item.Blueprint8, item.Blueprint9, item.Blueprint10, item.Blueprint11, item.Blueprint1save, item.Owner, item.Manager, item.Agent, item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Resulttext1, item.Resulttext2, item.Resulttext3, item.Resulttext4, item.Resulttext5, item.Past, item.Category, item.User, item.Apt, item.Date)
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

func (p *PeriodicManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodic_tb where d_id = ")
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

func (p *PeriodicManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodic_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodicManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and d_")
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

func (p *PeriodicManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodic_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodicManager) Update(item *Periodic) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
    if item.Reportdate == "" {
       item.Reportdate = "1000-01-01"
    }
	
    if item.Startdate == "" {
       item.Startdate = "1000-01-01"
    }
	
    if item.Enddate == "" {
       item.Enddate = "1000-01-01"
    }
	
	
	
	
	
	
    if item.Prestartdate == "" {
       item.Prestartdate = "1000-01-01"
    }
	
    if item.Preenddate == "" {
       item.Preenddate = "1000-01-01"
    }
	
    if item.Researchstartdate == "" {
       item.Researchstartdate = "1000-01-01"
    }
	
    if item.Researchenddate == "" {
       item.Researchenddate = "1000-01-01"
    }
	
    if item.Analyzestartdate == "" {
       item.Analyzestartdate = "1000-01-01"
    }
	
    if item.Analyzeenddate == "" {
       item.Analyzeenddate = "1000-01-01"
    }
	
    if item.Ratingstartdate == "" {
       item.Ratingstartdate = "1000-01-01"
    }
	
    if item.Ratingenddate == "" {
       item.Ratingenddate = "1000-01-01"
    }
	
    if item.Writestartdate == "" {
       item.Writestartdate = "1000-01-01"
    }
	
    if item.Writeenddate == "" {
       item.Writeenddate = "1000-01-01"
    }
	
    if item.Printstartdate == "" {
       item.Printstartdate = "1000-01-01"
    }
	
    if item.Printenddate == "" {
       item.Printenddate = "1000-01-01"
    }
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodic_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" d_name = $1, d_aptname = $2, d_taskrange = $3, d_reportdate = $4, d_startdate = $5, d_enddate = $6, d_supply = $7, d_contract = $8, d_price = $9, d_safetygrade = $10, d_status = $11, d_prestartdate = $12, d_preenddate = $13, d_researchstartdate = $14, d_researchenddate = $15, d_analyzestartdate = $16, d_analyzeenddate = $17, d_ratingstartdate = $18, d_ratingenddate = $19, d_writestartdate = $20, d_writeenddate = $21, d_printstartdate = $22, d_printenddate = $23, d_blueprint1 = $24, d_blueprint2 = $25, d_blueprint3 = $26, d_blueprint4 = $27, d_blueprint5 = $28, d_blueprint6 = $29, d_blueprint7 = $30, d_blueprint8 = $31, d_blueprint9 = $32, d_blueprint10 = $33, d_blueprint11 = $34, d_blueprint1save = $35, d_owner = $36, d_manager = $37, d_agent = $38, d_result1 = $39, d_result2 = $40, d_result3 = $41, d_result4 = $42, d_result5 = $43, d_resulttext1 = $44, d_resulttext2 = $45, d_resulttext3 = $46, d_resulttext4 = $47, d_resulttext5 = $48, d_past = $49, d_category = $50, d_user = $51, d_apt = $52, d_date = $53 where d_id = $54")
    } else {
        query.WriteString(" d_name = ?, d_aptname = ?, d_taskrange = ?, d_reportdate = ?, d_startdate = ?, d_enddate = ?, d_supply = ?, d_contract = ?, d_price = ?, d_safetygrade = ?, d_status = ?, d_prestartdate = ?, d_preenddate = ?, d_researchstartdate = ?, d_researchenddate = ?, d_analyzestartdate = ?, d_analyzeenddate = ?, d_ratingstartdate = ?, d_ratingenddate = ?, d_writestartdate = ?, d_writeenddate = ?, d_printstartdate = ?, d_printenddate = ?, d_blueprint1 = ?, d_blueprint2 = ?, d_blueprint3 = ?, d_blueprint4 = ?, d_blueprint5 = ?, d_blueprint6 = ?, d_blueprint7 = ?, d_blueprint8 = ?, d_blueprint9 = ?, d_blueprint10 = ?, d_blueprint11 = ?, d_blueprint1save = ?, d_owner = ?, d_manager = ?, d_agent = ?, d_result1 = ?, d_result2 = ?, d_result3 = ?, d_result4 = ?, d_result5 = ?, d_resulttext1 = ?, d_resulttext2 = ?, d_resulttext3 = ?, d_resulttext4 = ?, d_resulttext5 = ?, d_past = ?, d_category = ?, d_user = ?, d_apt = ?, d_date = ? where d_id = ?")
    }

	_, err := p.Exec(query.String() , item.Name, item.Aptname, item.Taskrange, item.Reportdate, item.Startdate, item.Enddate, item.Supply, item.Contract, item.Price, item.Safetygrade, item.Status, item.Prestartdate, item.Preenddate, item.Researchstartdate, item.Researchenddate, item.Analyzestartdate, item.Analyzeenddate, item.Ratingstartdate, item.Ratingenddate, item.Writestartdate, item.Writeenddate, item.Printstartdate, item.Printenddate, item.Blueprint1, item.Blueprint2, item.Blueprint3, item.Blueprint4, item.Blueprint5, item.Blueprint6, item.Blueprint7, item.Blueprint8, item.Blueprint9, item.Blueprint10, item.Blueprint11, item.Blueprint1save, item.Owner, item.Manager, item.Agent, item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Resulttext1, item.Resulttext2, item.Resulttext3, item.Resulttext4, item.Resulttext5, item.Past, item.Category, item.User, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodicManager) UpdateWhere(columns []periodic.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodic_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodic.ColumnId {
        initQuery.WriteString("d_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnName {
        initQuery.WriteString("d_name = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnAptname {
        initQuery.WriteString("d_aptname = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnTaskrange {
        initQuery.WriteString("d_taskrange = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnReportdate {
        initQuery.WriteString("d_reportdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnStartdate {
        initQuery.WriteString("d_startdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnEnddate {
        initQuery.WriteString("d_enddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnSupply {
        initQuery.WriteString("d_supply = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnContract {
        initQuery.WriteString("d_contract = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnPrice {
        initQuery.WriteString("d_price = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnSafetygrade {
        initQuery.WriteString("d_safetygrade = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnStatus {
        initQuery.WriteString("d_status = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnPrestartdate {
        initQuery.WriteString("d_prestartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnPreenddate {
        initQuery.WriteString("d_preenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResearchstartdate {
        initQuery.WriteString("d_researchstartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResearchenddate {
        initQuery.WriteString("d_researchenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnAnalyzestartdate {
        initQuery.WriteString("d_analyzestartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnAnalyzeenddate {
        initQuery.WriteString("d_analyzeenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnRatingstartdate {
        initQuery.WriteString("d_ratingstartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnRatingenddate {
        initQuery.WriteString("d_ratingenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnWritestartdate {
        initQuery.WriteString("d_writestartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnWriteenddate {
        initQuery.WriteString("d_writeenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnPrintstartdate {
        initQuery.WriteString("d_printstartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnPrintenddate {
        initQuery.WriteString("d_printenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint1 {
        initQuery.WriteString("d_blueprint1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint2 {
        initQuery.WriteString("d_blueprint2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint3 {
        initQuery.WriteString("d_blueprint3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint4 {
        initQuery.WriteString("d_blueprint4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint5 {
        initQuery.WriteString("d_blueprint5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint6 {
        initQuery.WriteString("d_blueprint6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint7 {
        initQuery.WriteString("d_blueprint7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint8 {
        initQuery.WriteString("d_blueprint8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint9 {
        initQuery.WriteString("d_blueprint9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint10 {
        initQuery.WriteString("d_blueprint10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint11 {
        initQuery.WriteString("d_blueprint11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnBlueprint1save {
        initQuery.WriteString("d_blueprint1save = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnOwner {
        initQuery.WriteString("d_owner = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnManager {
        initQuery.WriteString("d_manager = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnAgent {
        initQuery.WriteString("d_agent = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResult1 {
        initQuery.WriteString("d_result1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResult2 {
        initQuery.WriteString("d_result2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResult3 {
        initQuery.WriteString("d_result3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResult4 {
        initQuery.WriteString("d_result4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResult5 {
        initQuery.WriteString("d_result5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResulttext1 {
        initQuery.WriteString("d_resulttext1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResulttext2 {
        initQuery.WriteString("d_resulttext2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResulttext3 {
        initQuery.WriteString("d_resulttext3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResulttext4 {
        initQuery.WriteString("d_resulttext4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnResulttext5 {
        initQuery.WriteString("d_resulttext5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnPast {
        initQuery.WriteString("d_past = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnCategory {
        initQuery.WriteString("d_category = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnUser {
        initQuery.WriteString("d_user = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnApt {
        initQuery.WriteString("d_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodic.ColumnDate {
        initQuery.WriteString("d_date = ?")
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

func (p *PeriodicManager) UpdateName(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_name = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateAptname(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_aptname = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateTaskrange(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_taskrange = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateReportdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_reportdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateStartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_startdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateEnddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_enddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateSupply(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_supply = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateContract(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_contract = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdatePrice(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_price = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateSafetygrade(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_safetygrade = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_status = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdatePrestartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_prestartdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdatePreenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_preenddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResearchstartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_researchstartdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResearchenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_researchenddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateAnalyzestartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_analyzestartdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateAnalyzeenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_analyzeenddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateRatingstartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_ratingstartdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateRatingenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_ratingenddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateWritestartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_writestartdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateWriteenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_writeenddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdatePrintstartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_printstartdate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdatePrintenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_printenddate = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint1 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint2 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint3 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint4 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint5 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint6 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint7 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint8 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint9 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint10 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint11(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint11 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateBlueprint1save(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint1save = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateOwner(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_owner = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateManager(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_manager = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateAgent(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_agent = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResult1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result1 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResult2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result2 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResult3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result3 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResult4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result4 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResult5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result5 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResulttext1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_resulttext1 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResulttext2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_resulttext2 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResulttext3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_resulttext3 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResulttext4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_resulttext4 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateResulttext5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_resulttext5 = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdatePast(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_past = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateCategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_category = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_user = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_apt = ? where d_id = ?"
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

func (p *PeriodicManager) IncreaseStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_status = d_status + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint1 = d_blueprint1 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint2 = d_blueprint2 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint3 = d_blueprint3 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint4 = d_blueprint4 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint5 = d_blueprint5 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint6 = d_blueprint6 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint7(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint7 = d_blueprint7 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint8(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint8 = d_blueprint8 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint9(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint9 = d_blueprint9 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint11(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint11 = d_blueprint11 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseBlueprint1save(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_blueprint1save = d_blueprint1save + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseResult1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result1 = d_result1 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseResult2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result2 = d_result2 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseResult3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result3 = d_result3 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseResult4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result4 = d_result4 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseResult5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_result5 = d_result5 + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseCategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_category = d_category + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_user = d_user + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodic_tb set d_apt = d_apt + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodicManager) GetIdentity() int64 {
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

func (p *Periodic) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodicManager) ReadRow(rows *sql.Rows) *Periodic {
    var item Periodic
    var err error

    var _apt Apt
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Aptname, &item.Taskrange, &item.Reportdate, &item.Startdate, &item.Enddate, &item.Supply, &item.Contract, &item.Price, &item.Safetygrade, &item.Status, &item.Prestartdate, &item.Preenddate, &item.Researchstartdate, &item.Researchenddate, &item.Analyzestartdate, &item.Analyzeenddate, &item.Ratingstartdate, &item.Ratingenddate, &item.Writestartdate, &item.Writeenddate, &item.Printstartdate, &item.Printenddate, &item.Blueprint1, &item.Blueprint2, &item.Blueprint3, &item.Blueprint4, &item.Blueprint5, &item.Blueprint6, &item.Blueprint7, &item.Blueprint8, &item.Blueprint9, &item.Blueprint10, &item.Blueprint11, &item.Blueprint1save, &item.Owner, &item.Manager, &item.Agent, &item.Result1, &item.Result2, &item.Result3, &item.Result4, &item.Result5, &item.Resulttext1, &item.Resulttext2, &item.Resulttext3, &item.Resulttext4, &item.Resulttext5, &item.Past, &item.Category, &item.User, &item.Apt, &item.Date, &_apt.Id, &_apt.Name, &_apt.Completeyear, &_apt.Flatcount, &_apt.Type, &_apt.Floor, &_apt.Familycount, &_apt.Familycount1, &_apt.Familycount2, &_apt.Familycount3, &_apt.Tel, &_apt.Fax, &_apt.Email, &_apt.Personalemail, &_apt.Personalname, &_apt.Personalhp, &_apt.Zip, &_apt.Address, &_apt.Address2, &_apt.Contracttype, &_apt.Contractprice, &_apt.Testdate, &_apt.Nexttestdate, &_apt.Repair, &_apt.Safety, &_apt.Fault, &_apt.Contractdate, &_apt.Contractduration, &_apt.Invoice, &_apt.Depositdate, &_apt.Fmsloginid, &_apt.Fmspasswd, &_apt.Facilitydivision, &_apt.Facilitycategory, &_apt.Position, &_apt.Area, &_apt.Groundfloor, &_apt.Undergroundfloor, &_apt.Useapproval, &_apt.Date)
        
        
        
        
        
        
        
        
        if item.Reportdate == "0000-00-00" || item.Reportdate == "1000-01-01" || item.Reportdate == "9999-01-01" {
            item.Reportdate = ""
        }
        
        if item.Startdate == "0000-00-00" || item.Startdate == "1000-01-01" || item.Startdate == "9999-01-01" {
            item.Startdate = ""
        }
        
        if item.Enddate == "0000-00-00" || item.Enddate == "1000-01-01" || item.Enddate == "9999-01-01" {
            item.Enddate = ""
        }
        
        
        
        
        
        
        
        
        
        
        
        if item.Prestartdate == "0000-00-00" || item.Prestartdate == "1000-01-01" || item.Prestartdate == "9999-01-01" {
            item.Prestartdate = ""
        }
        
        if item.Preenddate == "0000-00-00" || item.Preenddate == "1000-01-01" || item.Preenddate == "9999-01-01" {
            item.Preenddate = ""
        }
        
        if item.Researchstartdate == "0000-00-00" || item.Researchstartdate == "1000-01-01" || item.Researchstartdate == "9999-01-01" {
            item.Researchstartdate = ""
        }
        
        if item.Researchenddate == "0000-00-00" || item.Researchenddate == "1000-01-01" || item.Researchenddate == "9999-01-01" {
            item.Researchenddate = ""
        }
        
        if item.Analyzestartdate == "0000-00-00" || item.Analyzestartdate == "1000-01-01" || item.Analyzestartdate == "9999-01-01" {
            item.Analyzestartdate = ""
        }
        
        if item.Analyzeenddate == "0000-00-00" || item.Analyzeenddate == "1000-01-01" || item.Analyzeenddate == "9999-01-01" {
            item.Analyzeenddate = ""
        }
        
        if item.Ratingstartdate == "0000-00-00" || item.Ratingstartdate == "1000-01-01" || item.Ratingstartdate == "9999-01-01" {
            item.Ratingstartdate = ""
        }
        
        if item.Ratingenddate == "0000-00-00" || item.Ratingenddate == "1000-01-01" || item.Ratingenddate == "9999-01-01" {
            item.Ratingenddate = ""
        }
        
        if item.Writestartdate == "0000-00-00" || item.Writestartdate == "1000-01-01" || item.Writestartdate == "9999-01-01" {
            item.Writestartdate = ""
        }
        
        if item.Writeenddate == "0000-00-00" || item.Writeenddate == "1000-01-01" || item.Writeenddate == "9999-01-01" {
            item.Writeenddate = ""
        }
        
        if item.Printstartdate == "0000-00-00" || item.Printstartdate == "1000-01-01" || item.Printstartdate == "9999-01-01" {
            item.Printstartdate = ""
        }
        
        if item.Printenddate == "0000-00-00" || item.Printenddate == "1000-01-01" || item.Printenddate == "9999-01-01" {
            item.Printenddate = ""
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
        _apt.InitExtra()
        item.AddExtra("apt",  _apt)

        return &item
    }
}

func (p *PeriodicManager) ReadRows(rows *sql.Rows) []Periodic {
    items := make([]Periodic, 0)

    for rows.Next() {
        var item Periodic
        var _apt Apt
            
    
        err := rows.Scan(&item.Id, &item.Name, &item.Aptname, &item.Taskrange, &item.Reportdate, &item.Startdate, &item.Enddate, &item.Supply, &item.Contract, &item.Price, &item.Safetygrade, &item.Status, &item.Prestartdate, &item.Preenddate, &item.Researchstartdate, &item.Researchenddate, &item.Analyzestartdate, &item.Analyzeenddate, &item.Ratingstartdate, &item.Ratingenddate, &item.Writestartdate, &item.Writeenddate, &item.Printstartdate, &item.Printenddate, &item.Blueprint1, &item.Blueprint2, &item.Blueprint3, &item.Blueprint4, &item.Blueprint5, &item.Blueprint6, &item.Blueprint7, &item.Blueprint8, &item.Blueprint9, &item.Blueprint10, &item.Blueprint11, &item.Blueprint1save, &item.Owner, &item.Manager, &item.Agent, &item.Result1, &item.Result2, &item.Result3, &item.Result4, &item.Result5, &item.Resulttext1, &item.Resulttext2, &item.Resulttext3, &item.Resulttext4, &item.Resulttext5, &item.Past, &item.Category, &item.User, &item.Apt, &item.Date, &_apt.Id, &_apt.Name, &_apt.Completeyear, &_apt.Flatcount, &_apt.Type, &_apt.Floor, &_apt.Familycount, &_apt.Familycount1, &_apt.Familycount2, &_apt.Familycount3, &_apt.Tel, &_apt.Fax, &_apt.Email, &_apt.Personalemail, &_apt.Personalname, &_apt.Personalhp, &_apt.Zip, &_apt.Address, &_apt.Address2, &_apt.Contracttype, &_apt.Contractprice, &_apt.Testdate, &_apt.Nexttestdate, &_apt.Repair, &_apt.Safety, &_apt.Fault, &_apt.Contractdate, &_apt.Contractduration, &_apt.Invoice, &_apt.Depositdate, &_apt.Fmsloginid, &_apt.Fmspasswd, &_apt.Facilitydivision, &_apt.Facilitycategory, &_apt.Position, &_apt.Area, &_apt.Groundfloor, &_apt.Undergroundfloor, &_apt.Useapproval, &_apt.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		
        
		
        
		if item.Reportdate == "0000-00-00" || item.Reportdate == "1000-01-01" || item.Reportdate == "9999-01-01" {
            item.Reportdate = ""
        }
        
		if item.Startdate == "0000-00-00" || item.Startdate == "1000-01-01" || item.Startdate == "9999-01-01" {
            item.Startdate = ""
        }
        
		if item.Enddate == "0000-00-00" || item.Enddate == "1000-01-01" || item.Enddate == "9999-01-01" {
            item.Enddate = ""
        }
        
		
        
		
        
		
        
		
        
		
        
		if item.Prestartdate == "0000-00-00" || item.Prestartdate == "1000-01-01" || item.Prestartdate == "9999-01-01" {
            item.Prestartdate = ""
        }
        
		if item.Preenddate == "0000-00-00" || item.Preenddate == "1000-01-01" || item.Preenddate == "9999-01-01" {
            item.Preenddate = ""
        }
        
		if item.Researchstartdate == "0000-00-00" || item.Researchstartdate == "1000-01-01" || item.Researchstartdate == "9999-01-01" {
            item.Researchstartdate = ""
        }
        
		if item.Researchenddate == "0000-00-00" || item.Researchenddate == "1000-01-01" || item.Researchenddate == "9999-01-01" {
            item.Researchenddate = ""
        }
        
		if item.Analyzestartdate == "0000-00-00" || item.Analyzestartdate == "1000-01-01" || item.Analyzestartdate == "9999-01-01" {
            item.Analyzestartdate = ""
        }
        
		if item.Analyzeenddate == "0000-00-00" || item.Analyzeenddate == "1000-01-01" || item.Analyzeenddate == "9999-01-01" {
            item.Analyzeenddate = ""
        }
        
		if item.Ratingstartdate == "0000-00-00" || item.Ratingstartdate == "1000-01-01" || item.Ratingstartdate == "9999-01-01" {
            item.Ratingstartdate = ""
        }
        
		if item.Ratingenddate == "0000-00-00" || item.Ratingenddate == "1000-01-01" || item.Ratingenddate == "9999-01-01" {
            item.Ratingenddate = ""
        }
        
		if item.Writestartdate == "0000-00-00" || item.Writestartdate == "1000-01-01" || item.Writestartdate == "9999-01-01" {
            item.Writestartdate = ""
        }
        
		if item.Writeenddate == "0000-00-00" || item.Writeenddate == "1000-01-01" || item.Writeenddate == "9999-01-01" {
            item.Writeenddate = ""
        }
        
		if item.Printstartdate == "0000-00-00" || item.Printstartdate == "1000-01-01" || item.Printstartdate == "9999-01-01" {
            item.Printstartdate = ""
        }
        
		if item.Printenddate == "0000-00-00" || item.Printenddate == "1000-01-01" || item.Printenddate == "9999-01-01" {
            item.Printenddate = ""
        }
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
		
        
        item.InitExtra()        
        _apt.InitExtra()
        item.AddExtra("apt",  _apt)

        items = append(items, item)
    }


     return items
}

func (p *PeriodicManager) Get(id int64) *Periodic {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and d_id = ?")

    
    query.WriteString(" and d_apt = a_id ")
    
    
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

func (p *PeriodicManager) GetWhere(args []any) *Periodic {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodicManager) Count(args []any) int {
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

func (p *PeriodicManager) FindAll() []Periodic {
    return p.Find(nil)
}

func (p *PeriodicManager) Find(args []any) []Periodic {
    if !p.Conn.IsConnect() {
        items := make([]Periodic, 0)
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
                query.WriteString(" and d_")
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
            orderby = "d_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "d_" + orderby
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
            orderby = "d_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "d_" + orderby
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
        items := make([]Periodic, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *PeriodicManager) Sum(args []any) *Periodic {
    if !p.Conn.IsConnect() {
        var item Periodic
        return &item
    }

    var params []any

    
    var query strings.Builder
    query.WriteString("select sum(d_price) from periodic_tb")

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
                query.WriteString(" and d_")
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
            orderby = "d_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "d_" + orderby
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
            orderby = "d_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "d_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(query.String(), params...)

    var item Periodic
    
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

func (p *PeriodicManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and d_")
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
    
    query.WriteString(" group by d_")
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



func (p *PeriodicManager) MakeMap(items []Periodic) map[int64]Periodic {
     ret := make(map[int64]Periodic)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodicManager) FindToMap(args []any) map[int64]Periodic {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodicManager) FindAllToMap() map[int64]Periodic {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *PeriodicManager) MakeNameMap(items []Periodic) map[string]Periodic {
     ret := make(map[string]Periodic)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *PeriodicManager) FindToNameMap(args []any) map[string]Periodic {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *PeriodicManager) FindAllToNameMap() map[string]Periodic {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}

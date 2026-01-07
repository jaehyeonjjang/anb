package models

import (
    "repair/models/apt"
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

type Apt struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Completeyear                string `json:"completeyear"`         
    Flatcount                string `json:"flatcount"`         
    Type                string `json:"type"`         
    Floor                string `json:"floor"`         
    Familycount                string `json:"familycount"`         
    Familycount1                int `json:"familycount1"`         
    Familycount2                int `json:"familycount2"`         
    Familycount3                int `json:"familycount3"`         
    Tel                string `json:"tel"`         
    Fax                string `json:"fax"`         
    Email                string `json:"email"`         
    Personalemail                string `json:"personalemail"`         
    Personalname                string `json:"personalname"`         
    Personalhp                string `json:"personalhp"`         
    Zip                string `json:"zip"`         
    Address                string `json:"address"`         
    Address2                string `json:"address2"`         
    Contracttype                int `json:"contracttype"`         
    Contractprice                string `json:"contractprice"`         
    Testdate                string `json:"testdate"`         
    Nexttestdate                string `json:"nexttestdate"`         
    Repair                string `json:"repair"`         
    Safety                string `json:"safety"`         
    Fault                string `json:"fault"`         
    Contractdate                string `json:"contractdate"`         
    Contractduration                string `json:"contractduration"`         
    Invoice                string `json:"invoice"`         
    Depositdate                string `json:"depositdate"`         
    Fmsloginid                string `json:"fmsloginid"`         
    Fmspasswd                string `json:"fmspasswd"`         
    Facilitydivision                int `json:"facilitydivision"`         
    Facilitycategory                int `json:"facilitycategory"`         
    Position                string `json:"position"`         
    Area                string `json:"area"`         
    Groundfloor                int `json:"groundfloor"`         
    Undergroundfloor                int `json:"undergroundfloor"`         
    Useapproval                string `json:"useapproval"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type AptManager struct {
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



func (c *Apt) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewAptManager(conn *Connection) *AptManager {
    var item AptManager


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

func (p *AptManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *AptManager) SetIndex(index string) {
    p.Index = index
}

func (p *AptManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *AptManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *AptManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *AptManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *AptManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date from apt_tb")

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

func (p *AptManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from apt_tb")

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

func (p *AptManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select a_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from apt_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *AptManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate apt_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *AptManager) Insert(item *Apt) error {
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
          query = "insert into apt_tb (a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40)"
        } else {
          query = "insert into apt_tb (a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Name, item.Completeyear, item.Flatcount, item.Type, item.Floor, item.Familycount, item.Familycount1, item.Familycount2, item.Familycount3, item.Tel, item.Fax, item.Email, item.Personalemail, item.Personalname, item.Personalhp, item.Zip, item.Address, item.Address2, item.Contracttype, item.Contractprice, item.Testdate, item.Nexttestdate, item.Repair, item.Safety, item.Fault, item.Contractdate, item.Contractduration, item.Invoice, item.Depositdate, item.Fmsloginid, item.Fmspasswd, item.Facilitydivision, item.Facilitycategory, item.Position, item.Area, item.Groundfloor, item.Undergroundfloor, item.Useapproval, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into apt_tb (a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39)"
        } else {
          query = "insert into apt_tb (a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Name, item.Completeyear, item.Flatcount, item.Type, item.Floor, item.Familycount, item.Familycount1, item.Familycount2, item.Familycount3, item.Tel, item.Fax, item.Email, item.Personalemail, item.Personalname, item.Personalhp, item.Zip, item.Address, item.Address2, item.Contracttype, item.Contractprice, item.Testdate, item.Nexttestdate, item.Repair, item.Safety, item.Fault, item.Contractdate, item.Contractduration, item.Invoice, item.Depositdate, item.Fmsloginid, item.Fmspasswd, item.Facilitydivision, item.Facilitycategory, item.Position, item.Area, item.Groundfloor, item.Undergroundfloor, item.Useapproval, item.Date)
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

func (p *AptManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from apt_tb where a_id = ")
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

func (p *AptManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from apt_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *AptManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and a_")
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

func (p *AptManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from apt_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *AptManager) Update(item *Apt) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update apt_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" a_name = $1, a_completeyear = $2, a_flatcount = $3, a_type = $4, a_floor = $5, a_familycount = $6, a_familycount1 = $7, a_familycount2 = $8, a_familycount3 = $9, a_tel = $10, a_fax = $11, a_email = $12, a_personalemail = $13, a_personalname = $14, a_personalhp = $15, a_zip = $16, a_address = $17, a_address2 = $18, a_contracttype = $19, a_contractprice = $20, a_testdate = $21, a_nexttestdate = $22, a_repair = $23, a_safety = $24, a_fault = $25, a_contractdate = $26, a_contractduration = $27, a_invoice = $28, a_depositdate = $29, a_fmsloginid = $30, a_fmspasswd = $31, a_facilitydivision = $32, a_facilitycategory = $33, a_position = $34, a_area = $35, a_groundfloor = $36, a_undergroundfloor = $37, a_useapproval = $38, a_date = $39 where a_id = $40")
    } else {
        query.WriteString(" a_name = ?, a_completeyear = ?, a_flatcount = ?, a_type = ?, a_floor = ?, a_familycount = ?, a_familycount1 = ?, a_familycount2 = ?, a_familycount3 = ?, a_tel = ?, a_fax = ?, a_email = ?, a_personalemail = ?, a_personalname = ?, a_personalhp = ?, a_zip = ?, a_address = ?, a_address2 = ?, a_contracttype = ?, a_contractprice = ?, a_testdate = ?, a_nexttestdate = ?, a_repair = ?, a_safety = ?, a_fault = ?, a_contractdate = ?, a_contractduration = ?, a_invoice = ?, a_depositdate = ?, a_fmsloginid = ?, a_fmspasswd = ?, a_facilitydivision = ?, a_facilitycategory = ?, a_position = ?, a_area = ?, a_groundfloor = ?, a_undergroundfloor = ?, a_useapproval = ?, a_date = ? where a_id = ?")
    }

	_, err := p.Exec(query.String() , item.Name, item.Completeyear, item.Flatcount, item.Type, item.Floor, item.Familycount, item.Familycount1, item.Familycount2, item.Familycount3, item.Tel, item.Fax, item.Email, item.Personalemail, item.Personalname, item.Personalhp, item.Zip, item.Address, item.Address2, item.Contracttype, item.Contractprice, item.Testdate, item.Nexttestdate, item.Repair, item.Safety, item.Fault, item.Contractdate, item.Contractduration, item.Invoice, item.Depositdate, item.Fmsloginid, item.Fmspasswd, item.Facilitydivision, item.Facilitycategory, item.Position, item.Area, item.Groundfloor, item.Undergroundfloor, item.Useapproval, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *AptManager) UpdateWhere(columns []apt.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update apt_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == apt.ColumnId {
        initQuery.WriteString("a_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnName {
        initQuery.WriteString("a_name = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnCompleteyear {
        initQuery.WriteString("a_completeyear = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFlatcount {
        initQuery.WriteString("a_flatcount = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnType {
        initQuery.WriteString("a_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFloor {
        initQuery.WriteString("a_floor = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFamilycount {
        initQuery.WriteString("a_familycount = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFamilycount1 {
        initQuery.WriteString("a_familycount1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFamilycount2 {
        initQuery.WriteString("a_familycount2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFamilycount3 {
        initQuery.WriteString("a_familycount3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnTel {
        initQuery.WriteString("a_tel = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFax {
        initQuery.WriteString("a_fax = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnEmail {
        initQuery.WriteString("a_email = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnPersonalemail {
        initQuery.WriteString("a_personalemail = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnPersonalname {
        initQuery.WriteString("a_personalname = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnPersonalhp {
        initQuery.WriteString("a_personalhp = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnZip {
        initQuery.WriteString("a_zip = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnAddress {
        initQuery.WriteString("a_address = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnAddress2 {
        initQuery.WriteString("a_address2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnContracttype {
        initQuery.WriteString("a_contracttype = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnContractprice {
        initQuery.WriteString("a_contractprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnTestdate {
        initQuery.WriteString("a_testdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnNexttestdate {
        initQuery.WriteString("a_nexttestdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnRepair {
        initQuery.WriteString("a_repair = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnSafety {
        initQuery.WriteString("a_safety = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFault {
        initQuery.WriteString("a_fault = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnContractdate {
        initQuery.WriteString("a_contractdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnContractduration {
        initQuery.WriteString("a_contractduration = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnInvoice {
        initQuery.WriteString("a_invoice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnDepositdate {
        initQuery.WriteString("a_depositdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFmsloginid {
        initQuery.WriteString("a_fmsloginid = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFmspasswd {
        initQuery.WriteString("a_fmspasswd = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFacilitydivision {
        initQuery.WriteString("a_facilitydivision = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnFacilitycategory {
        initQuery.WriteString("a_facilitycategory = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnPosition {
        initQuery.WriteString("a_position = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnArea {
        initQuery.WriteString("a_area = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnGroundfloor {
        initQuery.WriteString("a_groundfloor = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnUndergroundfloor {
        initQuery.WriteString("a_undergroundfloor = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnUseapproval {
        initQuery.WriteString("a_useapproval = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == apt.ColumnDate {
        initQuery.WriteString("a_date = ?")
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

func (p *AptManager) UpdateName(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_name = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateCompleteyear(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_completeyear = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFlatcount(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_flatcount = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateType(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_type = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFloor(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_floor = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFamilycount(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFamilycount1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount1 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFamilycount2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount2 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFamilycount3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount3 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateTel(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_tel = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFax(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_fax = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateEmail(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_email = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdatePersonalemail(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_personalemail = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdatePersonalname(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_personalname = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdatePersonalhp(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_personalhp = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateZip(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_zip = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateAddress(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_address = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateAddress2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_address2 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateContracttype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_contracttype = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateContractprice(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_contractprice = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateTestdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_testdate = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateNexttestdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_nexttestdate = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateRepair(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_repair = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateSafety(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_safety = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFault(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_fault = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateContractdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_contractdate = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateContractduration(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_contractduration = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateInvoice(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_invoice = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateDepositdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_depositdate = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFmsloginid(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_fmsloginid = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFmspasswd(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_fmspasswd = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFacilitydivision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_facilitydivision = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateFacilitycategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_facilitycategory = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdatePosition(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_position = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateArea(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_area = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateGroundfloor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_groundfloor = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateUndergroundfloor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_undergroundfloor = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) UpdateUseapproval(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_useapproval = ? where a_id = ?"
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

func (p *AptManager) IncreaseFamilycount1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount1 = a_familycount1 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseFamilycount2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount2 = a_familycount2 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseFamilycount3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_familycount3 = a_familycount3 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseContracttype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_contracttype = a_contracttype + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseFacilitydivision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_facilitydivision = a_facilitydivision + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseFacilitycategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_facilitycategory = a_facilitycategory + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseGroundfloor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_groundfloor = a_groundfloor + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *AptManager) IncreaseUndergroundfloor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update apt_tb set a_undergroundfloor = a_undergroundfloor + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *AptManager) GetIdentity() int64 {
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

func (p *Apt) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *AptManager) ReadRow(rows *sql.Rows) *Apt {
    var item Apt
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Completeyear, &item.Flatcount, &item.Type, &item.Floor, &item.Familycount, &item.Familycount1, &item.Familycount2, &item.Familycount3, &item.Tel, &item.Fax, &item.Email, &item.Personalemail, &item.Personalname, &item.Personalhp, &item.Zip, &item.Address, &item.Address2, &item.Contracttype, &item.Contractprice, &item.Testdate, &item.Nexttestdate, &item.Repair, &item.Safety, &item.Fault, &item.Contractdate, &item.Contractduration, &item.Invoice, &item.Depositdate, &item.Fmsloginid, &item.Fmspasswd, &item.Facilitydivision, &item.Facilitycategory, &item.Position, &item.Area, &item.Groundfloor, &item.Undergroundfloor, &item.Useapproval, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *AptManager) ReadRows(rows *sql.Rows) []Apt {
    items := make([]Apt, 0)

    for rows.Next() {
        var item Apt
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Completeyear, &item.Flatcount, &item.Type, &item.Floor, &item.Familycount, &item.Familycount1, &item.Familycount2, &item.Familycount3, &item.Tel, &item.Fax, &item.Email, &item.Personalemail, &item.Personalname, &item.Personalhp, &item.Zip, &item.Address, &item.Address2, &item.Contracttype, &item.Contractprice, &item.Testdate, &item.Nexttestdate, &item.Repair, &item.Safety, &item.Fault, &item.Contractdate, &item.Contractduration, &item.Invoice, &item.Depositdate, &item.Fmsloginid, &item.Fmspasswd, &item.Facilitydivision, &item.Facilitycategory, &item.Position, &item.Area, &item.Groundfloor, &item.Undergroundfloor, &item.Useapproval, &item.Date)
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

func (p *AptManager) Get(id int64) *Apt {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and a_id = ?")

    
    
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

func (p *AptManager) GetWhere(args []any) *Apt {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *AptManager) Count(args []any) int {
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

func (p *AptManager) FindAll() []Apt {
    return p.Find(nil)
}

func (p *AptManager) Find(args []any) []Apt {
    if !p.Conn.IsConnect() {
        items := make([]Apt, 0)
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
                query.WriteString(" and a_")
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
            orderby = "a_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "a_" + orderby
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
            orderby = "a_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "a_" + orderby
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
        items := make([]Apt, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *AptManager) CountByNamelike(name string, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if name != "" { 
        rets = append(rets, Where{Column:"name", Value:name, Compare:"like"})
     }
    
    return p.Count(rets)
}

func (p *AptManager) FindByNamelike(name string, args ...any) []Apt {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if name != "" { 
        rets = append(rets, Where{Column:"name", Value:name, Compare:"like"})
     }
    
    
    return p.Find(rets)
}

func (p *AptManager) CountByEmaillike(email string, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if email != "" { 
        rets = append(rets, Where{Column:"email", Value:email, Compare:"like"})
     }
    
    return p.Count(rets)
}

func (p *AptManager) FindByEmaillike(email string, args ...any) []Apt {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if email != "" { 
        rets = append(rets, Where{Column:"email", Value:email, Compare:"like"})
     }
    
    
    return p.Find(rets)
}




func (p *AptManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and a_")
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
    
    query.WriteString(" group by a_")
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



func (p *AptManager) MakeMap(items []Apt) map[int64]Apt {
     ret := make(map[int64]Apt)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *AptManager) FindToMap(args []any) map[int64]Apt {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *AptManager) FindAllToMap() map[int64]Apt {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *AptManager) MakeNameMap(items []Apt) map[string]Apt {
     ret := make(map[string]Apt)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *AptManager) FindToNameMap(args []any) map[string]Apt {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *AptManager) FindAllToNameMap() map[string]Apt {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}

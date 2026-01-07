package models

import (
    "repair/models/contract"
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

type Contract struct {
            
    Id                int64 `json:"id"`         
    Type                int `json:"type"`         
    Contractdate                string `json:"contractdate"`         
    Contractstartdate                string `json:"contractstartdate"`         
    Contractenddate                string `json:"contractenddate"`         
    Price                int `json:"price"`         
    Vat                int `json:"vat"`         
    Invoice                string `json:"invoice"`         
    Depositdate                string `json:"depositdate"`         
    Remark                string `json:"remark"`         
    User                int64 `json:"user"`         
    Estimate                int64 `json:"estimate"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type ContractManager struct {
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



func (c *Contract) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewContractManager(conn *Connection) *ContractManager {
    var item ContractManager


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

func (p *ContractManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ContractManager) SetIndex(index string) {
    p.Index = index
}

func (p *ContractManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *ContractManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *ContractManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *ContractManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *ContractManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select co_id, co_type, co_contractdate, co_contractstartdate, co_contractenddate, co_price, co_vat, co_invoice, co_depositdate, co_remark, co_user, co_estimate, co_apt, co_date, u_id, u_loginid, u_passwd, u_name, u_email, u_level, u_apt, u_date from contract_tb, user_tb")

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
    
    ret.WriteString(" and co_user = u_id ")
    

    return ret.String()
}

func (p *ContractManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from contract_tb, user_tb")

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
    
    ret.WriteString(" and co_user = u_id ")
    

    return ret.String()
}

func (p *ContractManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select co_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from contract_tb, user_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and co_user = u_id ")
    


    return ret.String()
}

func (p *ContractManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate contract_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *ContractManager) Insert(item *Contract) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
    if item.Contractdate == "" {
       item.Contractdate = "1000-01-01"
    }
	
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
	
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
	
	
	
    if item.Invoice == "" {
       item.Invoice = "1000-01-01"
    }
	
    if item.Depositdate == "" {
       item.Depositdate = "1000-01-01"
    }
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into contract_tb (co_id, co_type, co_contractdate, co_contractstartdate, co_contractenddate, co_price, co_vat, co_invoice, co_depositdate, co_remark, co_user, co_estimate, co_apt, co_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"
        } else {
          query = "insert into contract_tb (co_id, co_type, co_contractdate, co_contractstartdate, co_contractenddate, co_price, co_vat, co_invoice, co_depositdate, co_remark, co_user, co_estimate, co_apt, co_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Type, item.Contractdate, item.Contractstartdate, item.Contractenddate, item.Price, item.Vat, item.Invoice, item.Depositdate, item.Remark, item.User, item.Estimate, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into contract_tb (co_type, co_contractdate, co_contractstartdate, co_contractenddate, co_price, co_vat, co_invoice, co_depositdate, co_remark, co_user, co_estimate, co_apt, co_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
        } else {
          query = "insert into contract_tb (co_type, co_contractdate, co_contractstartdate, co_contractenddate, co_price, co_vat, co_invoice, co_depositdate, co_remark, co_user, co_estimate, co_apt, co_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Type, item.Contractdate, item.Contractstartdate, item.Contractenddate, item.Price, item.Vat, item.Invoice, item.Depositdate, item.Remark, item.User, item.Estimate, item.Apt, item.Date)
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

func (p *ContractManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from contract_tb where co_id = ")
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

func (p *ContractManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from contract_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *ContractManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and co_")
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

func (p *ContractManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from contract_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *ContractManager) Update(item *Contract) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
    if item.Contractdate == "" {
       item.Contractdate = "1000-01-01"
    }
	
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
	
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
	
	
	
    if item.Invoice == "" {
       item.Invoice = "1000-01-01"
    }
	
    if item.Depositdate == "" {
       item.Depositdate = "1000-01-01"
    }
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update contract_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" co_type = $1, co_contractdate = $2, co_contractstartdate = $3, co_contractenddate = $4, co_price = $5, co_vat = $6, co_invoice = $7, co_depositdate = $8, co_remark = $9, co_user = $10, co_estimate = $11, co_apt = $12, co_date = $13 where co_id = $14")
    } else {
        query.WriteString(" co_type = ?, co_contractdate = ?, co_contractstartdate = ?, co_contractenddate = ?, co_price = ?, co_vat = ?, co_invoice = ?, co_depositdate = ?, co_remark = ?, co_user = ?, co_estimate = ?, co_apt = ?, co_date = ? where co_id = ?")
    }

	_, err := p.Exec(query.String() , item.Type, item.Contractdate, item.Contractstartdate, item.Contractenddate, item.Price, item.Vat, item.Invoice, item.Depositdate, item.Remark, item.User, item.Estimate, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *ContractManager) UpdateWhere(columns []contract.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update contract_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == contract.ColumnId {
        initQuery.WriteString("co_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnType {
        initQuery.WriteString("co_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnContractdate {
        initQuery.WriteString("co_contractdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnContractstartdate {
        initQuery.WriteString("co_contractstartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnContractenddate {
        initQuery.WriteString("co_contractenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnPrice {
        initQuery.WriteString("co_price = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnVat {
        initQuery.WriteString("co_vat = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnInvoice {
        initQuery.WriteString("co_invoice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnDepositdate {
        initQuery.WriteString("co_depositdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnRemark {
        initQuery.WriteString("co_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnUser {
        initQuery.WriteString("co_user = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnEstimate {
        initQuery.WriteString("co_estimate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnApt {
        initQuery.WriteString("co_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == contract.ColumnDate {
        initQuery.WriteString("co_date = ?")
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

func (p *ContractManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_type = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateContractdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_contractdate = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateContractstartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_contractstartdate = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateContractenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_contractenddate = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdatePrice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_price = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateVat(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_vat = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateInvoice(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_invoice = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateDepositdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_depositdate = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_remark = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_user = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateEstimate(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_estimate = ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_apt = ? where co_id = ?"
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

func (p *ContractManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_type = co_type + ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) IncreasePrice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_price = co_price + ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) IncreaseVat(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_vat = co_vat + ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) IncreaseUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_user = co_user + ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) IncreaseEstimate(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_estimate = co_estimate + ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ContractManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update contract_tb set co_apt = co_apt + ? where co_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *ContractManager) GetIdentity() int64 {
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

func (p *Contract) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *ContractManager) ReadRow(rows *sql.Rows) *Contract {
    var item Contract
    var err error

    var _user User
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Type, &item.Contractdate, &item.Contractstartdate, &item.Contractenddate, &item.Price, &item.Vat, &item.Invoice, &item.Depositdate, &item.Remark, &item.User, &item.Estimate, &item.Apt, &item.Date, &_user.Id, &_user.Loginid, &_user.Passwd, &_user.Name, &_user.Email, &_user.Level, &_user.Apt, &_user.Date)
        
        
        
        
        if item.Contractdate == "0000-00-00" || item.Contractdate == "1000-01-01" || item.Contractdate == "9999-01-01" {
            item.Contractdate = ""
        }
        
        if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" || item.Contractstartdate == "9999-01-01" {
            item.Contractstartdate = ""
        }
        
        if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" || item.Contractenddate == "9999-01-01" {
            item.Contractenddate = ""
        }
        
        
        
        
        
        if item.Invoice == "0000-00-00" || item.Invoice == "1000-01-01" || item.Invoice == "9999-01-01" {
            item.Invoice = ""
        }
        
        if item.Depositdate == "0000-00-00" || item.Depositdate == "1000-01-01" || item.Depositdate == "9999-01-01" {
            item.Depositdate = ""
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
        _user.InitExtra()
        item.AddExtra("user",  _user)

        return &item
    }
}

func (p *ContractManager) ReadRows(rows *sql.Rows) []Contract {
    items := make([]Contract, 0)

    for rows.Next() {
        var item Contract
        var _user User
            
    
        err := rows.Scan(&item.Id, &item.Type, &item.Contractdate, &item.Contractstartdate, &item.Contractenddate, &item.Price, &item.Vat, &item.Invoice, &item.Depositdate, &item.Remark, &item.User, &item.Estimate, &item.Apt, &item.Date, &_user.Id, &_user.Loginid, &_user.Passwd, &_user.Name, &_user.Email, &_user.Level, &_user.Apt, &_user.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		if item.Contractdate == "0000-00-00" || item.Contractdate == "1000-01-01" || item.Contractdate == "9999-01-01" {
            item.Contractdate = ""
        }
        
		if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" || item.Contractstartdate == "9999-01-01" {
            item.Contractstartdate = ""
        }
        
		if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" || item.Contractenddate == "9999-01-01" {
            item.Contractenddate = ""
        }
        
		
        
		
        
		if item.Invoice == "0000-00-00" || item.Invoice == "1000-01-01" || item.Invoice == "9999-01-01" {
            item.Invoice = ""
        }
        
		if item.Depositdate == "0000-00-00" || item.Depositdate == "1000-01-01" || item.Depositdate == "9999-01-01" {
            item.Depositdate = ""
        }
        
		
        
		
        
		
        
		
        
		
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
		
        
        item.InitExtra()        
        _user.InitExtra()
        item.AddExtra("user",  _user)

        items = append(items, item)
    }


     return items
}

func (p *ContractManager) Get(id int64) *Contract {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and co_id = ?")

    
    query.WriteString(" and co_user = u_id ")
    
    
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

func (p *ContractManager) GetWhere(args []any) *Contract {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *ContractManager) Count(args []any) int {
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

func (p *ContractManager) FindAll() []Contract {
    return p.Find(nil)
}

func (p *ContractManager) Find(args []any) []Contract {
    if !p.Conn.IsConnect() {
        items := make([]Contract, 0)
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
                query.WriteString(" and co_")
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
            orderby = "co_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "co_" + orderby
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
            orderby = "co_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "co_" + orderby
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
        items := make([]Contract, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *ContractManager) GetByEstimate(estimate int64, args ...any) *Contract {
    rets := make([]any, 0)
    rets = append(rets, args...)
    if estimate != 0 {
        rets = append(rets, Where{Column:"estimate", Value:estimate, Compare:"="})        
    }
    
    items := p.Find(rets)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}


func (p *ContractManager) Sum(args []any) *Contract {
    if !p.Conn.IsConnect() {
        var item Contract
        return &item
    }

    var params []any

    
    var query strings.Builder
    query.WriteString("select sum(co_price) from contract_tb")

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
                query.WriteString(" and co_")
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
            orderby = "co_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "co_" + orderby
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
            orderby = "co_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                   orderby = "co_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(query.String(), params...)

    var item Contract
    
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

func (p *ContractManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and co_")
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
    
    query.WriteString(" group by co_")
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



func (p *ContractManager) MakeMap(items []Contract) map[int64]Contract {
     ret := make(map[int64]Contract)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *ContractManager) FindToMap(args []any) map[int64]Contract {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *ContractManager) FindAllToMap() map[int64]Contract {
     items := p.Find(nil)
     return p.MakeMap(items)
}



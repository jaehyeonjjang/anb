package models

import (
    "repair/models/comparecompany"
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

type Comparecompany struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Tel                string `json:"tel"`         
    Fax                string `json:"fax"`         
    Ceo                string `json:"ceo"`         
    Format                string `json:"format"`         
    Image                string `json:"image"`         
    Image2                string `json:"image2"`         
    Adjust                int `json:"adjust"`         
    Financialprice                int `json:"financialprice"`         
    Techprice                int `json:"techprice"`         
    Directprice                int `json:"directprice"`         
    Printprice                int `json:"printprice"`         
    Extraprice                int `json:"extraprice"`         
    Travelprice                int `json:"travelprice"`         
    Gasprice                int `json:"gasprice"`         
    Dangerprice                int `json:"dangerprice"`         
    Machineprice                int `json:"machineprice"`         
    Remark                string `json:"remark"`         
    Type                int `json:"type"`         
    Default                int `json:"default"`         
    Order                int `json:"order"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type ComparecompanyManager struct {
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



func (c *Comparecompany) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewComparecompanyManager(conn *Connection) *ComparecompanyManager {
    var item ComparecompanyManager


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

func (p *ComparecompanyManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ComparecompanyManager) SetIndex(index string) {
    p.Index = index
}

func (p *ComparecompanyManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *ComparecompanyManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *ComparecompanyManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *ComparecompanyManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *ComparecompanyManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select cc_id, cc_name, cc_address, cc_addressetc, cc_tel, cc_fax, cc_ceo, cc_format, cc_image, cc_image2, cc_adjust, cc_financialprice, cc_techprice, cc_directprice, cc_printprice, cc_extraprice, cc_travelprice, cc_gasprice, cc_dangerprice, cc_machineprice, cc_remark, cc_type, cc_default, cc_order, cc_date from comparecompany_tb")

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

func (p *ComparecompanyManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from comparecompany_tb")

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

func (p *ComparecompanyManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select cc_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from comparecompany_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *ComparecompanyManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate comparecompany_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *ComparecompanyManager) Insert(item *Comparecompany) error {
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
          query = "insert into comparecompany_tb (cc_id, cc_name, cc_address, cc_addressetc, cc_tel, cc_fax, cc_ceo, cc_format, cc_image, cc_image2, cc_adjust, cc_financialprice, cc_techprice, cc_directprice, cc_printprice, cc_extraprice, cc_travelprice, cc_gasprice, cc_dangerprice, cc_machineprice, cc_remark, cc_type, cc_default, cc_order, cc_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)"
        } else {
          query = "insert into comparecompany_tb (cc_id, cc_name, cc_address, cc_addressetc, cc_tel, cc_fax, cc_ceo, cc_format, cc_image, cc_image2, cc_adjust, cc_financialprice, cc_techprice, cc_directprice, cc_printprice, cc_extraprice, cc_travelprice, cc_gasprice, cc_dangerprice, cc_machineprice, cc_remark, cc_type, cc_default, cc_order, cc_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Name, item.Address, item.Addressetc, item.Tel, item.Fax, item.Ceo, item.Format, item.Image, item.Image2, item.Adjust, item.Financialprice, item.Techprice, item.Directprice, item.Printprice, item.Extraprice, item.Travelprice, item.Gasprice, item.Dangerprice, item.Machineprice, item.Remark, item.Type, item.Default, item.Order, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into comparecompany_tb (cc_name, cc_address, cc_addressetc, cc_tel, cc_fax, cc_ceo, cc_format, cc_image, cc_image2, cc_adjust, cc_financialprice, cc_techprice, cc_directprice, cc_printprice, cc_extraprice, cc_travelprice, cc_gasprice, cc_dangerprice, cc_machineprice, cc_remark, cc_type, cc_default, cc_order, cc_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)"
        } else {
          query = "insert into comparecompany_tb (cc_name, cc_address, cc_addressetc, cc_tel, cc_fax, cc_ceo, cc_format, cc_image, cc_image2, cc_adjust, cc_financialprice, cc_techprice, cc_directprice, cc_printprice, cc_extraprice, cc_travelprice, cc_gasprice, cc_dangerprice, cc_machineprice, cc_remark, cc_type, cc_default, cc_order, cc_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Name, item.Address, item.Addressetc, item.Tel, item.Fax, item.Ceo, item.Format, item.Image, item.Image2, item.Adjust, item.Financialprice, item.Techprice, item.Directprice, item.Printprice, item.Extraprice, item.Travelprice, item.Gasprice, item.Dangerprice, item.Machineprice, item.Remark, item.Type, item.Default, item.Order, item.Date)
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

func (p *ComparecompanyManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from comparecompany_tb where cc_id = ")
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

func (p *ComparecompanyManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from comparecompany_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *ComparecompanyManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and cc_")
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

func (p *ComparecompanyManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from comparecompany_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *ComparecompanyManager) Update(item *Comparecompany) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update comparecompany_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" cc_name = $1, cc_address = $2, cc_addressetc = $3, cc_tel = $4, cc_fax = $5, cc_ceo = $6, cc_format = $7, cc_image = $8, cc_image2 = $9, cc_adjust = $10, cc_financialprice = $11, cc_techprice = $12, cc_directprice = $13, cc_printprice = $14, cc_extraprice = $15, cc_travelprice = $16, cc_gasprice = $17, cc_dangerprice = $18, cc_machineprice = $19, cc_remark = $20, cc_type = $21, cc_default = $22, cc_order = $23, cc_date = $24 where cc_id = $25")
    } else {
        query.WriteString(" cc_name = ?, cc_address = ?, cc_addressetc = ?, cc_tel = ?, cc_fax = ?, cc_ceo = ?, cc_format = ?, cc_image = ?, cc_image2 = ?, cc_adjust = ?, cc_financialprice = ?, cc_techprice = ?, cc_directprice = ?, cc_printprice = ?, cc_extraprice = ?, cc_travelprice = ?, cc_gasprice = ?, cc_dangerprice = ?, cc_machineprice = ?, cc_remark = ?, cc_type = ?, cc_default = ?, cc_order = ?, cc_date = ? where cc_id = ?")
    }

	_, err := p.Exec(query.String() , item.Name, item.Address, item.Addressetc, item.Tel, item.Fax, item.Ceo, item.Format, item.Image, item.Image2, item.Adjust, item.Financialprice, item.Techprice, item.Directprice, item.Printprice, item.Extraprice, item.Travelprice, item.Gasprice, item.Dangerprice, item.Machineprice, item.Remark, item.Type, item.Default, item.Order, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *ComparecompanyManager) UpdateWhere(columns []comparecompany.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update comparecompany_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == comparecompany.ColumnId {
        initQuery.WriteString("cc_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnName {
        initQuery.WriteString("cc_name = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnAddress {
        initQuery.WriteString("cc_address = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnAddressetc {
        initQuery.WriteString("cc_addressetc = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnTel {
        initQuery.WriteString("cc_tel = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnFax {
        initQuery.WriteString("cc_fax = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnCeo {
        initQuery.WriteString("cc_ceo = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnFormat {
        initQuery.WriteString("cc_format = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnImage {
        initQuery.WriteString("cc_image = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnImage2 {
        initQuery.WriteString("cc_image2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnAdjust {
        initQuery.WriteString("cc_adjust = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnFinancialprice {
        initQuery.WriteString("cc_financialprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnTechprice {
        initQuery.WriteString("cc_techprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnDirectprice {
        initQuery.WriteString("cc_directprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnPrintprice {
        initQuery.WriteString("cc_printprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnExtraprice {
        initQuery.WriteString("cc_extraprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnTravelprice {
        initQuery.WriteString("cc_travelprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnGasprice {
        initQuery.WriteString("cc_gasprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnDangerprice {
        initQuery.WriteString("cc_dangerprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnMachineprice {
        initQuery.WriteString("cc_machineprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnRemark {
        initQuery.WriteString("cc_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnType {
        initQuery.WriteString("cc_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnDefault {
        initQuery.WriteString("cc_default = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnOrder {
        initQuery.WriteString("cc_order = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == comparecompany.ColumnDate {
        initQuery.WriteString("cc_date = ?")
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

func (p *ComparecompanyManager) UpdateName(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_name = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateAddress(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_address = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateAddressetc(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_addressetc = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateTel(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_tel = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateFax(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_fax = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateCeo(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_ceo = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateFormat(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_format = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateImage(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_image = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateImage2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_image2 = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateAdjust(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_adjust = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateFinancialprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_financialprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateTechprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_techprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateDirectprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_directprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdatePrintprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_printprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateExtraprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_extraprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateTravelprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_travelprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateGasprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_gasprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateDangerprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_dangerprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateMachineprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_machineprice = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_remark = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_type = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateDefault(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_default = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) UpdateOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_order = ? where cc_id = ?"
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

func (p *ComparecompanyManager) IncreaseAdjust(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_adjust = cc_adjust + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseFinancialprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_financialprice = cc_financialprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseTechprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_techprice = cc_techprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseDirectprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_directprice = cc_directprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreasePrintprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_printprice = cc_printprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseExtraprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_extraprice = cc_extraprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseTravelprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_travelprice = cc_travelprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseGasprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_gasprice = cc_gasprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseDangerprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_dangerprice = cc_dangerprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseMachineprice(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_machineprice = cc_machineprice + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_type = cc_type + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseDefault(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_default = cc_default + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *ComparecompanyManager) IncreaseOrder(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update comparecompany_tb set cc_order = cc_order + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *ComparecompanyManager) GetIdentity() int64 {
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

func (p *Comparecompany) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *ComparecompanyManager) ReadRow(rows *sql.Rows) *Comparecompany {
    var item Comparecompany
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Address, &item.Addressetc, &item.Tel, &item.Fax, &item.Ceo, &item.Format, &item.Image, &item.Image2, &item.Adjust, &item.Financialprice, &item.Techprice, &item.Directprice, &item.Printprice, &item.Extraprice, &item.Travelprice, &item.Gasprice, &item.Dangerprice, &item.Machineprice, &item.Remark, &item.Type, &item.Default, &item.Order, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *ComparecompanyManager) ReadRows(rows *sql.Rows) []Comparecompany {
    items := make([]Comparecompany, 0)

    for rows.Next() {
        var item Comparecompany
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Address, &item.Addressetc, &item.Tel, &item.Fax, &item.Ceo, &item.Format, &item.Image, &item.Image2, &item.Adjust, &item.Financialprice, &item.Techprice, &item.Directprice, &item.Printprice, &item.Extraprice, &item.Travelprice, &item.Gasprice, &item.Dangerprice, &item.Machineprice, &item.Remark, &item.Type, &item.Default, &item.Order, &item.Date)
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

func (p *ComparecompanyManager) Get(id int64) *Comparecompany {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and cc_id = ?")

    
    
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

func (p *ComparecompanyManager) GetWhere(args []any) *Comparecompany {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *ComparecompanyManager) Count(args []any) int {
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

func (p *ComparecompanyManager) FindAll() []Comparecompany {
    return p.Find(nil)
}

func (p *ComparecompanyManager) Find(args []any) []Comparecompany {
    if !p.Conn.IsConnect() {
        items := make([]Comparecompany, 0)
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
                query.WriteString(" and cc_")
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
            orderby = "cc_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "cc_" + orderby
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
            orderby = "cc_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "cc_" + orderby
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
        items := make([]Comparecompany, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *ComparecompanyManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and cc_")
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
    
    query.WriteString(" group by cc_")
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



func (p *ComparecompanyManager) MakeMap(items []Comparecompany) map[int64]Comparecompany {
     ret := make(map[int64]Comparecompany)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *ComparecompanyManager) FindToMap(args []any) map[int64]Comparecompany {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *ComparecompanyManager) FindAllToMap() map[int64]Comparecompany {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *ComparecompanyManager) MakeNameMap(items []Comparecompany) map[string]Comparecompany {
     ret := make(map[string]Comparecompany)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *ComparecompanyManager) FindToNameMap(args []any) map[string]Comparecompany {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *ComparecompanyManager) FindAllToNameMap() map[string]Comparecompany {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}

package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type EstimateController struct {
	controllers.Controller
}



func (c *EstimateController) Insert(item *models.Estimate) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewEstimateManager(conn)
	err := manager.Insert(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *EstimateController) Insertbatch(item *[]models.Estimate) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewEstimateManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *EstimateController) Update(item *models.Estimate) {
    
    
	conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *EstimateController) Delete(item *models.Estimate) {
    
    
    conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *EstimateController) Deletebatch(item *[]models.Estimate) {
    
    
    conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *EstimateController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _subtype := c.Geti("subtype")
    if _subtype != 0 {
        args = append(args, models.Where{Column:"subtype", Value:_subtype, Compare:"="})    
    }
    _originalprice := c.Geti("originalprice")
    if _originalprice != 0 {
        args = append(args, models.Where{Column:"originalprice", Value:_originalprice, Compare:"="})    
    }
    _saleprice := c.Geti("saleprice")
    if _saleprice != 0 {
        args = append(args, models.Where{Column:"saleprice", Value:_saleprice, Compare:"="})    
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _financialprice := c.Geti("financialprice")
    if _financialprice != 0 {
        args = append(args, models.Where{Column:"financialprice", Value:_financialprice, Compare:"="})    
    }
    _techprice := c.Geti("techprice")
    if _techprice != 0 {
        args = append(args, models.Where{Column:"techprice", Value:_techprice, Compare:"="})    
    }
    _directprice := c.Geti("directprice")
    if _directprice != 0 {
        args = append(args, models.Where{Column:"directprice", Value:_directprice, Compare:"="})    
    }
    _printprice := c.Geti("printprice")
    if _printprice != 0 {
        args = append(args, models.Where{Column:"printprice", Value:_printprice, Compare:"="})    
    }
    _extraprice := c.Geti("extraprice")
    if _extraprice != 0 {
        args = append(args, models.Where{Column:"extraprice", Value:_extraprice, Compare:"="})    
    }
    _travelprice := c.Geti("travelprice")
    if _travelprice != 0 {
        args = append(args, models.Where{Column:"travelprice", Value:_travelprice, Compare:"="})    
    }
    _lossprice := c.Geti("lossprice")
    if _lossprice != 0 {
        args = append(args, models.Where{Column:"lossprice", Value:_lossprice, Compare:"="})    
    }
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _etcprice := c.Geti("etcprice")
    if _etcprice != 0 {
        args = append(args, models.Where{Column:"etcprice", Value:_etcprice, Compare:"="})    
    }
    _dangerprice := c.Geti("dangerprice")
    if _dangerprice != 0 {
        args = append(args, models.Where{Column:"dangerprice", Value:_dangerprice, Compare:"="})    
    }
    _machineprice := c.Geti("machineprice")
    if _machineprice != 0 {
        args = append(args, models.Where{Column:"machineprice", Value:_machineprice, Compare:"="})    
    }
    _carprice := c.Geti("carprice")
    if _carprice != 0 {
        args = append(args, models.Where{Column:"carprice", Value:_carprice, Compare:"="})    
    }
    _discount := c.Geti("discount")
    if _discount != 0 {
        args = append(args, models.Where{Column:"discount", Value:_discount, Compare:"="})    
    }
    _person1 := c.Geti("person1")
    if _person1 != 0 {
        args = append(args, models.Where{Column:"person1", Value:_person1, Compare:"="})    
    }
    _person2 := c.Geti("person2")
    if _person2 != 0 {
        args = append(args, models.Where{Column:"person2", Value:_person2, Compare:"="})    
    }
    _person3 := c.Geti("person3")
    if _person3 != 0 {
        args = append(args, models.Where{Column:"person3", Value:_person3, Compare:"="})    
    }
    _person4 := c.Geti("person4")
    if _person4 != 0 {
        args = append(args, models.Where{Column:"person4", Value:_person4, Compare:"="})    
    }
    _person5 := c.Geti("person5")
    if _person5 != 0 {
        args = append(args, models.Where{Column:"person5", Value:_person5, Compare:"="})    
    }
    _person6 := c.Geti("person6")
    if _person6 != 0 {
        args = append(args, models.Where{Column:"person6", Value:_person6, Compare:"="})    
    }
    _person7 := c.Geti("person7")
    if _person7 != 0 {
        args = append(args, models.Where{Column:"person7", Value:_person7, Compare:"="})    
    }
    _person8 := c.Geti("person8")
    if _person8 != 0 {
        args = append(args, models.Where{Column:"person8", Value:_person8, Compare:"="})    
    }
    _person9 := c.Geti("person9")
    if _person9 != 0 {
        args = append(args, models.Where{Column:"person9", Value:_person9, Compare:"="})    
    }
    _person10 := c.Geti("person10")
    if _person10 != 0 {
        args = append(args, models.Where{Column:"person10", Value:_person10, Compare:"="})    
    }
    _personprice1 := c.Geti("personprice1")
    if _personprice1 != 0 {
        args = append(args, models.Where{Column:"personprice1", Value:_personprice1, Compare:"="})    
    }
    _personprice2 := c.Geti("personprice2")
    if _personprice2 != 0 {
        args = append(args, models.Where{Column:"personprice2", Value:_personprice2, Compare:"="})    
    }
    _personprice3 := c.Geti("personprice3")
    if _personprice3 != 0 {
        args = append(args, models.Where{Column:"personprice3", Value:_personprice3, Compare:"="})    
    }
    _personprice4 := c.Geti("personprice4")
    if _personprice4 != 0 {
        args = append(args, models.Where{Column:"personprice4", Value:_personprice4, Compare:"="})    
    }
    _personprice5 := c.Geti("personprice5")
    if _personprice5 != 0 {
        args = append(args, models.Where{Column:"personprice5", Value:_personprice5, Compare:"="})    
    }
    _personprice6 := c.Geti("personprice6")
    if _personprice6 != 0 {
        args = append(args, models.Where{Column:"personprice6", Value:_personprice6, Compare:"="})    
    }
    _personprice7 := c.Geti("personprice7")
    if _personprice7 != 0 {
        args = append(args, models.Where{Column:"personprice7", Value:_personprice7, Compare:"="})    
    }
    _personprice8 := c.Geti("personprice8")
    if _personprice8 != 0 {
        args = append(args, models.Where{Column:"personprice8", Value:_personprice8, Compare:"="})    
    }
    _personprice9 := c.Geti("personprice9")
    if _personprice9 != 0 {
        args = append(args, models.Where{Column:"personprice9", Value:_personprice9, Compare:"="})    
    }
    _personprice10 := c.Geti("personprice10")
    if _personprice10 != 0 {
        args = append(args, models.Where{Column:"personprice10", Value:_personprice10, Compare:"="})    
    }
    _days := c.Geti("days")
    if _days != 0 {
        args = append(args, models.Where{Column:"days", Value:_days, Compare:"="})    
    }
    _travel := c.Geti("travel")
    if _travel != 0 {
        args = append(args, models.Where{Column:"travel", Value:_travel, Compare:"="})    
    }
    _loss := c.Geti("loss")
    if _loss != 0 {
        args = append(args, models.Where{Column:"loss", Value:_loss, Compare:"="})    
    }
    _gas := c.Geti("gas")
    if _gas != 0 {
        args = append(args, models.Where{Column:"gas", Value:_gas, Compare:"="})    
    }
    _comparecount := c.Geti("comparecount")
    if _comparecount != 0 {
        args = append(args, models.Where{Column:"comparecount", Value:_comparecount, Compare:"="})    
    }
    _etc := c.Geti("etc")
    if _etc != 0 {
        args = append(args, models.Where{Column:"etc", Value:_etc, Compare:"="})    
    }
    _danger := c.Geti("danger")
    if _danger != 0 {
        args = append(args, models.Where{Column:"danger", Value:_danger, Compare:"="})    
    }
    _machine := c.Geti("machine")
    if _machine != 0 {
        args = append(args, models.Where{Column:"machine", Value:_machine, Compare:"="})    
    }
    _car := c.Geti("car")
    if _car != 0 {
        args = append(args, models.Where{Column:"car", Value:_car, Compare:"="})    
    }
    _print := c.Geti("print")
    if _print != 0 {
        args = append(args, models.Where{Column:"print", Value:_print, Compare:"="})    
    }
    _stability := c.Geti("stability")
    if _stability != 0 {
        args = append(args, models.Where{Column:"stability", Value:_stability, Compare:"="})    
    }
    _earthquake := c.Geti("earthquake")
    if _earthquake != 0 {
        args = append(args, models.Where{Column:"earthquake", Value:_earthquake, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _writedate := c.Get("writedate")
    if _writedate != "" {
        args = append(args, models.Where{Column:"writedate", Value:_writedate, Compare:"="})
    }
    _start := c.Get("start")
    if _start != "" {
        args = append(args, models.Where{Column:"start", Value:_start, Compare:"="})
    }
    _event := c.Geti("event")
    if _event != 0 {
        args = append(args, models.Where{Column:"event", Value:_event, Compare:"="})    
    }
    _parcel := c.Geti("parcel")
    if _parcel != 0 {
        args = append(args, models.Where{Column:"parcel", Value:_parcel, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")

    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *EstimateController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *EstimateController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _subtype := c.Geti("subtype")
    if _subtype != 0 {
        args = append(args, models.Where{Column:"subtype", Value:_subtype, Compare:"="})    
    }
    _originalprice := c.Geti("originalprice")
    if _originalprice != 0 {
        args = append(args, models.Where{Column:"originalprice", Value:_originalprice, Compare:"="})    
    }
    _saleprice := c.Geti("saleprice")
    if _saleprice != 0 {
        args = append(args, models.Where{Column:"saleprice", Value:_saleprice, Compare:"="})    
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _financialprice := c.Geti("financialprice")
    if _financialprice != 0 {
        args = append(args, models.Where{Column:"financialprice", Value:_financialprice, Compare:"="})    
    }
    _techprice := c.Geti("techprice")
    if _techprice != 0 {
        args = append(args, models.Where{Column:"techprice", Value:_techprice, Compare:"="})    
    }
    _directprice := c.Geti("directprice")
    if _directprice != 0 {
        args = append(args, models.Where{Column:"directprice", Value:_directprice, Compare:"="})    
    }
    _printprice := c.Geti("printprice")
    if _printprice != 0 {
        args = append(args, models.Where{Column:"printprice", Value:_printprice, Compare:"="})    
    }
    _extraprice := c.Geti("extraprice")
    if _extraprice != 0 {
        args = append(args, models.Where{Column:"extraprice", Value:_extraprice, Compare:"="})    
    }
    _travelprice := c.Geti("travelprice")
    if _travelprice != 0 {
        args = append(args, models.Where{Column:"travelprice", Value:_travelprice, Compare:"="})    
    }
    _lossprice := c.Geti("lossprice")
    if _lossprice != 0 {
        args = append(args, models.Where{Column:"lossprice", Value:_lossprice, Compare:"="})    
    }
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _etcprice := c.Geti("etcprice")
    if _etcprice != 0 {
        args = append(args, models.Where{Column:"etcprice", Value:_etcprice, Compare:"="})    
    }
    _dangerprice := c.Geti("dangerprice")
    if _dangerprice != 0 {
        args = append(args, models.Where{Column:"dangerprice", Value:_dangerprice, Compare:"="})    
    }
    _machineprice := c.Geti("machineprice")
    if _machineprice != 0 {
        args = append(args, models.Where{Column:"machineprice", Value:_machineprice, Compare:"="})    
    }
    _carprice := c.Geti("carprice")
    if _carprice != 0 {
        args = append(args, models.Where{Column:"carprice", Value:_carprice, Compare:"="})    
    }
    _discount := c.Geti("discount")
    if _discount != 0 {
        args = append(args, models.Where{Column:"discount", Value:_discount, Compare:"="})    
    }
    _person1 := c.Geti("person1")
    if _person1 != 0 {
        args = append(args, models.Where{Column:"person1", Value:_person1, Compare:"="})    
    }
    _person2 := c.Geti("person2")
    if _person2 != 0 {
        args = append(args, models.Where{Column:"person2", Value:_person2, Compare:"="})    
    }
    _person3 := c.Geti("person3")
    if _person3 != 0 {
        args = append(args, models.Where{Column:"person3", Value:_person3, Compare:"="})    
    }
    _person4 := c.Geti("person4")
    if _person4 != 0 {
        args = append(args, models.Where{Column:"person4", Value:_person4, Compare:"="})    
    }
    _person5 := c.Geti("person5")
    if _person5 != 0 {
        args = append(args, models.Where{Column:"person5", Value:_person5, Compare:"="})    
    }
    _person6 := c.Geti("person6")
    if _person6 != 0 {
        args = append(args, models.Where{Column:"person6", Value:_person6, Compare:"="})    
    }
    _person7 := c.Geti("person7")
    if _person7 != 0 {
        args = append(args, models.Where{Column:"person7", Value:_person7, Compare:"="})    
    }
    _person8 := c.Geti("person8")
    if _person8 != 0 {
        args = append(args, models.Where{Column:"person8", Value:_person8, Compare:"="})    
    }
    _person9 := c.Geti("person9")
    if _person9 != 0 {
        args = append(args, models.Where{Column:"person9", Value:_person9, Compare:"="})    
    }
    _person10 := c.Geti("person10")
    if _person10 != 0 {
        args = append(args, models.Where{Column:"person10", Value:_person10, Compare:"="})    
    }
    _personprice1 := c.Geti("personprice1")
    if _personprice1 != 0 {
        args = append(args, models.Where{Column:"personprice1", Value:_personprice1, Compare:"="})    
    }
    _personprice2 := c.Geti("personprice2")
    if _personprice2 != 0 {
        args = append(args, models.Where{Column:"personprice2", Value:_personprice2, Compare:"="})    
    }
    _personprice3 := c.Geti("personprice3")
    if _personprice3 != 0 {
        args = append(args, models.Where{Column:"personprice3", Value:_personprice3, Compare:"="})    
    }
    _personprice4 := c.Geti("personprice4")
    if _personprice4 != 0 {
        args = append(args, models.Where{Column:"personprice4", Value:_personprice4, Compare:"="})    
    }
    _personprice5 := c.Geti("personprice5")
    if _personprice5 != 0 {
        args = append(args, models.Where{Column:"personprice5", Value:_personprice5, Compare:"="})    
    }
    _personprice6 := c.Geti("personprice6")
    if _personprice6 != 0 {
        args = append(args, models.Where{Column:"personprice6", Value:_personprice6, Compare:"="})    
    }
    _personprice7 := c.Geti("personprice7")
    if _personprice7 != 0 {
        args = append(args, models.Where{Column:"personprice7", Value:_personprice7, Compare:"="})    
    }
    _personprice8 := c.Geti("personprice8")
    if _personprice8 != 0 {
        args = append(args, models.Where{Column:"personprice8", Value:_personprice8, Compare:"="})    
    }
    _personprice9 := c.Geti("personprice9")
    if _personprice9 != 0 {
        args = append(args, models.Where{Column:"personprice9", Value:_personprice9, Compare:"="})    
    }
    _personprice10 := c.Geti("personprice10")
    if _personprice10 != 0 {
        args = append(args, models.Where{Column:"personprice10", Value:_personprice10, Compare:"="})    
    }
    _days := c.Geti("days")
    if _days != 0 {
        args = append(args, models.Where{Column:"days", Value:_days, Compare:"="})    
    }
    _travel := c.Geti("travel")
    if _travel != 0 {
        args = append(args, models.Where{Column:"travel", Value:_travel, Compare:"="})    
    }
    _loss := c.Geti("loss")
    if _loss != 0 {
        args = append(args, models.Where{Column:"loss", Value:_loss, Compare:"="})    
    }
    _gas := c.Geti("gas")
    if _gas != 0 {
        args = append(args, models.Where{Column:"gas", Value:_gas, Compare:"="})    
    }
    _comparecount := c.Geti("comparecount")
    if _comparecount != 0 {
        args = append(args, models.Where{Column:"comparecount", Value:_comparecount, Compare:"="})    
    }
    _etc := c.Geti("etc")
    if _etc != 0 {
        args = append(args, models.Where{Column:"etc", Value:_etc, Compare:"="})    
    }
    _danger := c.Geti("danger")
    if _danger != 0 {
        args = append(args, models.Where{Column:"danger", Value:_danger, Compare:"="})    
    }
    _machine := c.Geti("machine")
    if _machine != 0 {
        args = append(args, models.Where{Column:"machine", Value:_machine, Compare:"="})    
    }
    _car := c.Geti("car")
    if _car != 0 {
        args = append(args, models.Where{Column:"car", Value:_car, Compare:"="})    
    }
    _print := c.Geti("print")
    if _print != 0 {
        args = append(args, models.Where{Column:"print", Value:_print, Compare:"="})    
    }
    _stability := c.Geti("stability")
    if _stability != 0 {
        args = append(args, models.Where{Column:"stability", Value:_stability, Compare:"="})    
    }
    _earthquake := c.Geti("earthquake")
    if _earthquake != 0 {
        args = append(args, models.Where{Column:"earthquake", Value:_earthquake, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _writedate := c.Get("writedate")
    if _writedate != "" {
        args = append(args, models.Where{Column:"writedate", Value:_writedate, Compare:"="})
    }
    _start := c.Get("start")
    if _start != "" {
        args = append(args, models.Where{Column:"start", Value:_start, Compare:"="})
    }
    _event := c.Geti("event")
    if _event != 0 {
        args = append(args, models.Where{Column:"event", Value:_event, Compare:"="})    
    }
    _parcel := c.Geti("parcel")
    if _parcel != 0 {
        args = append(args, models.Where{Column:"parcel", Value:_parcel, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    if page != 0 && pagesize != 0 {
        args = append(args, models.Paging(page, pagesize))
    }
    
    orderby := c.Get("orderby")
    if orderby == "" {
        if page != 0 && pagesize != 0 {
            orderby = "id desc"
            args = append(args, models.Ordering(orderby))
        }
    } else {
        orderbys := strings.Split(orderby, ",")

        str := ""
        for i, v := range orderbys {
            if i == 0 {
                str += v
            } else {
                if strings.Contains(v, "_") {                   
                    str += ", " + strings.Trim(v, " ")
                } else {
                    str += ", e_" + strings.Trim(v, " ")                
                }
            }
        }
        
        args = append(args, models.Ordering(str))
    }
    
	items := manager.Find(args)
	c.Set("items", items)

    if page == 1 {
       total := manager.Count(args)
	   c.Set("total", total)
    }
}





func (c *EstimateController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewEstimateManager(conn)

    var args []interface{}
    
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _subtype := c.Geti("subtype")
    if _subtype != 0 {
        args = append(args, models.Where{Column:"subtype", Value:_subtype, Compare:"="})    
    }
    _originalprice := c.Geti("originalprice")
    if _originalprice != 0 {
        args = append(args, models.Where{Column:"originalprice", Value:_originalprice, Compare:"="})    
    }
    _saleprice := c.Geti("saleprice")
    if _saleprice != 0 {
        args = append(args, models.Where{Column:"saleprice", Value:_saleprice, Compare:"="})    
    }
    _price := c.Geti("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
    }
    _financialprice := c.Geti("financialprice")
    if _financialprice != 0 {
        args = append(args, models.Where{Column:"financialprice", Value:_financialprice, Compare:"="})    
    }
    _techprice := c.Geti("techprice")
    if _techprice != 0 {
        args = append(args, models.Where{Column:"techprice", Value:_techprice, Compare:"="})    
    }
    _directprice := c.Geti("directprice")
    if _directprice != 0 {
        args = append(args, models.Where{Column:"directprice", Value:_directprice, Compare:"="})    
    }
    _printprice := c.Geti("printprice")
    if _printprice != 0 {
        args = append(args, models.Where{Column:"printprice", Value:_printprice, Compare:"="})    
    }
    _extraprice := c.Geti("extraprice")
    if _extraprice != 0 {
        args = append(args, models.Where{Column:"extraprice", Value:_extraprice, Compare:"="})    
    }
    _travelprice := c.Geti("travelprice")
    if _travelprice != 0 {
        args = append(args, models.Where{Column:"travelprice", Value:_travelprice, Compare:"="})    
    }
    _lossprice := c.Geti("lossprice")
    if _lossprice != 0 {
        args = append(args, models.Where{Column:"lossprice", Value:_lossprice, Compare:"="})    
    }
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _etcprice := c.Geti("etcprice")
    if _etcprice != 0 {
        args = append(args, models.Where{Column:"etcprice", Value:_etcprice, Compare:"="})    
    }
    _dangerprice := c.Geti("dangerprice")
    if _dangerprice != 0 {
        args = append(args, models.Where{Column:"dangerprice", Value:_dangerprice, Compare:"="})    
    }
    _machineprice := c.Geti("machineprice")
    if _machineprice != 0 {
        args = append(args, models.Where{Column:"machineprice", Value:_machineprice, Compare:"="})    
    }
    _carprice := c.Geti("carprice")
    if _carprice != 0 {
        args = append(args, models.Where{Column:"carprice", Value:_carprice, Compare:"="})    
    }
    _discount := c.Geti("discount")
    if _discount != 0 {
        args = append(args, models.Where{Column:"discount", Value:_discount, Compare:"="})    
    }
    _person1 := c.Geti("person1")
    if _person1 != 0 {
        args = append(args, models.Where{Column:"person1", Value:_person1, Compare:"="})    
    }
    _person2 := c.Geti("person2")
    if _person2 != 0 {
        args = append(args, models.Where{Column:"person2", Value:_person2, Compare:"="})    
    }
    _person3 := c.Geti("person3")
    if _person3 != 0 {
        args = append(args, models.Where{Column:"person3", Value:_person3, Compare:"="})    
    }
    _person4 := c.Geti("person4")
    if _person4 != 0 {
        args = append(args, models.Where{Column:"person4", Value:_person4, Compare:"="})    
    }
    _person5 := c.Geti("person5")
    if _person5 != 0 {
        args = append(args, models.Where{Column:"person5", Value:_person5, Compare:"="})    
    }
    _person6 := c.Geti("person6")
    if _person6 != 0 {
        args = append(args, models.Where{Column:"person6", Value:_person6, Compare:"="})    
    }
    _person7 := c.Geti("person7")
    if _person7 != 0 {
        args = append(args, models.Where{Column:"person7", Value:_person7, Compare:"="})    
    }
    _person8 := c.Geti("person8")
    if _person8 != 0 {
        args = append(args, models.Where{Column:"person8", Value:_person8, Compare:"="})    
    }
    _person9 := c.Geti("person9")
    if _person9 != 0 {
        args = append(args, models.Where{Column:"person9", Value:_person9, Compare:"="})    
    }
    _person10 := c.Geti("person10")
    if _person10 != 0 {
        args = append(args, models.Where{Column:"person10", Value:_person10, Compare:"="})    
    }
    _personprice1 := c.Geti("personprice1")
    if _personprice1 != 0 {
        args = append(args, models.Where{Column:"personprice1", Value:_personprice1, Compare:"="})    
    }
    _personprice2 := c.Geti("personprice2")
    if _personprice2 != 0 {
        args = append(args, models.Where{Column:"personprice2", Value:_personprice2, Compare:"="})    
    }
    _personprice3 := c.Geti("personprice3")
    if _personprice3 != 0 {
        args = append(args, models.Where{Column:"personprice3", Value:_personprice3, Compare:"="})    
    }
    _personprice4 := c.Geti("personprice4")
    if _personprice4 != 0 {
        args = append(args, models.Where{Column:"personprice4", Value:_personprice4, Compare:"="})    
    }
    _personprice5 := c.Geti("personprice5")
    if _personprice5 != 0 {
        args = append(args, models.Where{Column:"personprice5", Value:_personprice5, Compare:"="})    
    }
    _personprice6 := c.Geti("personprice6")
    if _personprice6 != 0 {
        args = append(args, models.Where{Column:"personprice6", Value:_personprice6, Compare:"="})    
    }
    _personprice7 := c.Geti("personprice7")
    if _personprice7 != 0 {
        args = append(args, models.Where{Column:"personprice7", Value:_personprice7, Compare:"="})    
    }
    _personprice8 := c.Geti("personprice8")
    if _personprice8 != 0 {
        args = append(args, models.Where{Column:"personprice8", Value:_personprice8, Compare:"="})    
    }
    _personprice9 := c.Geti("personprice9")
    if _personprice9 != 0 {
        args = append(args, models.Where{Column:"personprice9", Value:_personprice9, Compare:"="})    
    }
    _personprice10 := c.Geti("personprice10")
    if _personprice10 != 0 {
        args = append(args, models.Where{Column:"personprice10", Value:_personprice10, Compare:"="})    
    }
    _days := c.Geti("days")
    if _days != 0 {
        args = append(args, models.Where{Column:"days", Value:_days, Compare:"="})    
    }
    _travel := c.Geti("travel")
    if _travel != 0 {
        args = append(args, models.Where{Column:"travel", Value:_travel, Compare:"="})    
    }
    _loss := c.Geti("loss")
    if _loss != 0 {
        args = append(args, models.Where{Column:"loss", Value:_loss, Compare:"="})    
    }
    _gas := c.Geti("gas")
    if _gas != 0 {
        args = append(args, models.Where{Column:"gas", Value:_gas, Compare:"="})    
    }
    _comparecount := c.Geti("comparecount")
    if _comparecount != 0 {
        args = append(args, models.Where{Column:"comparecount", Value:_comparecount, Compare:"="})    
    }
    _etc := c.Geti("etc")
    if _etc != 0 {
        args = append(args, models.Where{Column:"etc", Value:_etc, Compare:"="})    
    }
    _danger := c.Geti("danger")
    if _danger != 0 {
        args = append(args, models.Where{Column:"danger", Value:_danger, Compare:"="})    
    }
    _machine := c.Geti("machine")
    if _machine != 0 {
        args = append(args, models.Where{Column:"machine", Value:_machine, Compare:"="})    
    }
    _car := c.Geti("car")
    if _car != 0 {
        args = append(args, models.Where{Column:"car", Value:_car, Compare:"="})    
    }
    _print := c.Geti("print")
    if _print != 0 {
        args = append(args, models.Where{Column:"print", Value:_print, Compare:"="})    
    }
    _stability := c.Geti("stability")
    if _stability != 0 {
        args = append(args, models.Where{Column:"stability", Value:_stability, Compare:"="})    
    }
    _earthquake := c.Geti("earthquake")
    if _earthquake != 0 {
        args = append(args, models.Where{Column:"earthquake", Value:_earthquake, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _writedate := c.Get("writedate")
    if _writedate != "" {
        args = append(args, models.Where{Column:"writedate", Value:_writedate, Compare:"like"})
    }
    _start := c.Get("start")
    if _start != "" {
        args = append(args, models.Where{Column:"start", Value:_start, Compare:"like"})
    }
    _event := c.Geti("event")
    if _event != 0 {
        args = append(args, models.Where{Column:"event", Value:_event, Compare:"="})    
    }
    _parcel := c.Geti("parcel")
    if _parcel != 0 {
        args = append(args, models.Where{Column:"parcel", Value:_parcel, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}


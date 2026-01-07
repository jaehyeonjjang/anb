package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type StandardwageController struct {
	controllers.Controller
}



func (c *StandardwageController) Insert(item *models.Standardwage) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewStandardwageManager(conn)
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

func (c *StandardwageController) Insertbatch(item *[]models.Standardwage) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewStandardwageManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *StandardwageController) Update(item *models.Standardwage) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardwageManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *StandardwageController) Delete(item *models.Standardwage) {
    
    
    conn := c.NewConnection()

	manager := models.NewStandardwageManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *StandardwageController) Deletebatch(item *[]models.Standardwage) {
    
    
    conn := c.NewConnection()

	manager := models.NewStandardwageManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *StandardwageController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardwageManager(conn)

    var args []interface{}
    
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
    _techprice1 := c.Geti("techprice1")
    if _techprice1 != 0 {
        args = append(args, models.Where{Column:"techprice1", Value:_techprice1, Compare:"="})    
    }
    _techprice2 := c.Geti("techprice2")
    if _techprice2 != 0 {
        args = append(args, models.Where{Column:"techprice2", Value:_techprice2, Compare:"="})    
    }
    _techprice3 := c.Geti("techprice3")
    if _techprice3 != 0 {
        args = append(args, models.Where{Column:"techprice3", Value:_techprice3, Compare:"="})    
    }
    _techprice4 := c.Geti("techprice4")
    if _techprice4 != 0 {
        args = append(args, models.Where{Column:"techprice4", Value:_techprice4, Compare:"="})    
    }
    _financialprice1 := c.Geti("financialprice1")
    if _financialprice1 != 0 {
        args = append(args, models.Where{Column:"financialprice1", Value:_financialprice1, Compare:"="})    
    }
    _financialprice2 := c.Geti("financialprice2")
    if _financialprice2 != 0 {
        args = append(args, models.Where{Column:"financialprice2", Value:_financialprice2, Compare:"="})    
    }
    _financialprice3 := c.Geti("financialprice3")
    if _financialprice3 != 0 {
        args = append(args, models.Where{Column:"financialprice3", Value:_financialprice3, Compare:"="})    
    }
    _financialprice4 := c.Geti("financialprice4")
    if _financialprice4 != 0 {
        args = append(args, models.Where{Column:"financialprice4", Value:_financialprice4, Compare:"="})    
    }
    _directprice := c.Geti("directprice")
    if _directprice != 0 {
        args = append(args, models.Where{Column:"directprice", Value:_directprice, Compare:"="})    
    }
    _printprice1 := c.Geti("printprice1")
    if _printprice1 != 0 {
        args = append(args, models.Where{Column:"printprice1", Value:_printprice1, Compare:"="})    
    }
    _printprice2 := c.Geti("printprice2")
    if _printprice2 != 0 {
        args = append(args, models.Where{Column:"printprice2", Value:_printprice2, Compare:"="})    
    }
    _lossprice := c.Geti("lossprice")
    if _lossprice != 0 {
        args = append(args, models.Where{Column:"lossprice", Value:_lossprice, Compare:"="})    
    }
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _travelprice := c.Geti("travelprice")
    if _travelprice != 0 {
        args = append(args, models.Where{Column:"travelprice", Value:_travelprice, Compare:"="})    
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
    _print := c.Geti("print")
    if _print != 0 {
        args = append(args, models.Where{Column:"print", Value:_print, Compare:"="})    
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


func (c *StandardwageController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardwageManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *StandardwageController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewStandardwageManager(conn)

    var args []interface{}
    
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
    _techprice1 := c.Geti("techprice1")
    if _techprice1 != 0 {
        args = append(args, models.Where{Column:"techprice1", Value:_techprice1, Compare:"="})    
    }
    _techprice2 := c.Geti("techprice2")
    if _techprice2 != 0 {
        args = append(args, models.Where{Column:"techprice2", Value:_techprice2, Compare:"="})    
    }
    _techprice3 := c.Geti("techprice3")
    if _techprice3 != 0 {
        args = append(args, models.Where{Column:"techprice3", Value:_techprice3, Compare:"="})    
    }
    _techprice4 := c.Geti("techprice4")
    if _techprice4 != 0 {
        args = append(args, models.Where{Column:"techprice4", Value:_techprice4, Compare:"="})    
    }
    _financialprice1 := c.Geti("financialprice1")
    if _financialprice1 != 0 {
        args = append(args, models.Where{Column:"financialprice1", Value:_financialprice1, Compare:"="})    
    }
    _financialprice2 := c.Geti("financialprice2")
    if _financialprice2 != 0 {
        args = append(args, models.Where{Column:"financialprice2", Value:_financialprice2, Compare:"="})    
    }
    _financialprice3 := c.Geti("financialprice3")
    if _financialprice3 != 0 {
        args = append(args, models.Where{Column:"financialprice3", Value:_financialprice3, Compare:"="})    
    }
    _financialprice4 := c.Geti("financialprice4")
    if _financialprice4 != 0 {
        args = append(args, models.Where{Column:"financialprice4", Value:_financialprice4, Compare:"="})    
    }
    _directprice := c.Geti("directprice")
    if _directprice != 0 {
        args = append(args, models.Where{Column:"directprice", Value:_directprice, Compare:"="})    
    }
    _printprice1 := c.Geti("printprice1")
    if _printprice1 != 0 {
        args = append(args, models.Where{Column:"printprice1", Value:_printprice1, Compare:"="})    
    }
    _printprice2 := c.Geti("printprice2")
    if _printprice2 != 0 {
        args = append(args, models.Where{Column:"printprice2", Value:_printprice2, Compare:"="})    
    }
    _lossprice := c.Geti("lossprice")
    if _lossprice != 0 {
        args = append(args, models.Where{Column:"lossprice", Value:_lossprice, Compare:"="})    
    }
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _travelprice := c.Geti("travelprice")
    if _travelprice != 0 {
        args = append(args, models.Where{Column:"travelprice", Value:_travelprice, Compare:"="})    
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
    _print := c.Geti("print")
    if _print != 0 {
        args = append(args, models.Where{Column:"print", Value:_print, Compare:"="})    
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
                    str += ", sw_" + strings.Trim(v, " ")                
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






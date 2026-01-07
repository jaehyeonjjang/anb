package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type ComparecompanyController struct {
	controllers.Controller
}



func (c *ComparecompanyController) Insert(item *models.Comparecompany) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewComparecompanyManager(conn)
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

func (c *ComparecompanyController) Insertbatch(item *[]models.Comparecompany) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewComparecompanyManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *ComparecompanyController) Update(item *models.Comparecompany) {
    
    
	conn := c.NewConnection()

	manager := models.NewComparecompanyManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *ComparecompanyController) Delete(item *models.Comparecompany) {
    
    
    conn := c.NewConnection()

	manager := models.NewComparecompanyManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *ComparecompanyController) Deletebatch(item *[]models.Comparecompany) {
    
    
    conn := c.NewConnection()

	manager := models.NewComparecompanyManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *ComparecompanyController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewComparecompanyManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"="})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"="})
    }
    _ceo := c.Get("ceo")
    if _ceo != "" {
        args = append(args, models.Where{Column:"ceo", Value:_ceo, Compare:"="})
    }
    _format := c.Get("format")
    if _format != "" {
        args = append(args, models.Where{Column:"format", Value:_format, Compare:"="})
    }
    _image := c.Get("image")
    if _image != "" {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})
    }
    _image2 := c.Get("image2")
    if _image2 != "" {
        args = append(args, models.Where{Column:"image2", Value:_image2, Compare:"="})
    }
    _adjust := c.Geti("adjust")
    if _adjust != 0 {
        args = append(args, models.Where{Column:"adjust", Value:_adjust, Compare:"="})    
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
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _dangerprice := c.Geti("dangerprice")
    if _dangerprice != 0 {
        args = append(args, models.Where{Column:"dangerprice", Value:_dangerprice, Compare:"="})    
    }
    _machineprice := c.Geti("machineprice")
    if _machineprice != 0 {
        args = append(args, models.Where{Column:"machineprice", Value:_machineprice, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _default := c.Geti("default")
    if _default != 0 {
        args = append(args, models.Where{Column:"default", Value:_default, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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


func (c *ComparecompanyController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewComparecompanyManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *ComparecompanyController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewComparecompanyManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _addressetc := c.Get("addressetc")
    if _addressetc != "" {
        args = append(args, models.Where{Column:"addressetc", Value:_addressetc, Compare:"="})
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"="})
    }
    _ceo := c.Get("ceo")
    if _ceo != "" {
        args = append(args, models.Where{Column:"ceo", Value:_ceo, Compare:"="})
    }
    _format := c.Get("format")
    if _format != "" {
        args = append(args, models.Where{Column:"format", Value:_format, Compare:"="})
    }
    _image := c.Get("image")
    if _image != "" {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"="})
    }
    _image2 := c.Get("image2")
    if _image2 != "" {
        args = append(args, models.Where{Column:"image2", Value:_image2, Compare:"="})
    }
    _adjust := c.Geti("adjust")
    if _adjust != 0 {
        args = append(args, models.Where{Column:"adjust", Value:_adjust, Compare:"="})    
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
    _gasprice := c.Geti("gasprice")
    if _gasprice != 0 {
        args = append(args, models.Where{Column:"gasprice", Value:_gasprice, Compare:"="})    
    }
    _dangerprice := c.Geti("dangerprice")
    if _dangerprice != 0 {
        args = append(args, models.Where{Column:"dangerprice", Value:_dangerprice, Compare:"="})    
    }
    _machineprice := c.Geti("machineprice")
    if _machineprice != 0 {
        args = append(args, models.Where{Column:"machineprice", Value:_machineprice, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _default := c.Geti("default")
    if _default != 0 {
        args = append(args, models.Where{Column:"default", Value:_default, Compare:"="})    
    }
    _order := c.Geti("order")
    if _order != 0 {
        args = append(args, models.Where{Column:"order", Value:_order, Compare:"="})    
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
                    str += ", cc_" + strings.Trim(v, " ")                
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






package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type AptController struct {
	controllers.Controller
}



func (c *AptController) CountByNamelike(name string) int {
    
    conn := c.NewConnection()

	_manager := models.NewAptManager(conn)
    
    item := _manager.CountByNamelike(name)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *AptController) FindByNamelike(name string) []models.Apt {
    
    conn := c.NewConnection()

	_manager := models.NewAptManager(conn)
    
    item := _manager.FindByNamelike(name)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *AptController) CountByEmaillike(email string) int {
    
    conn := c.NewConnection()

	_manager := models.NewAptManager(conn)
    
    item := _manager.CountByEmaillike(email)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *AptController) FindByEmaillike(email string) []models.Apt {
    
    conn := c.NewConnection()

	_manager := models.NewAptManager(conn)
    
    item := _manager.FindByEmaillike(email)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *AptController) Insert(item *models.Apt) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptManager(conn)
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

func (c *AptController) Insertbatch(item *[]models.Apt) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewAptManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptController) Update(item *models.Apt) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *AptController) Delete(item *models.Apt) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *AptController) Deletebatch(item *[]models.Apt) {
    
    
    conn := c.NewConnection()

	manager := models.NewAptManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *AptController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewAptManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _completeyear := c.Get("completeyear")
    if _completeyear != "" {
        args = append(args, models.Where{Column:"completeyear", Value:_completeyear, Compare:"="})
    }
    _flatcount := c.Get("flatcount")
    if _flatcount != "" {
        args = append(args, models.Where{Column:"flatcount", Value:_flatcount, Compare:"="})
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
    }
    _floor := c.Get("floor")
    if _floor != "" {
        args = append(args, models.Where{Column:"floor", Value:_floor, Compare:"="})
    }
    _familycount := c.Get("familycount")
    if _familycount != "" {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})
    }
    _familycount1 := c.Geti("familycount1")
    if _familycount1 != 0 {
        args = append(args, models.Where{Column:"familycount1", Value:_familycount1, Compare:"="})    
    }
    _familycount2 := c.Geti("familycount2")
    if _familycount2 != 0 {
        args = append(args, models.Where{Column:"familycount2", Value:_familycount2, Compare:"="})    
    }
    _familycount3 := c.Geti("familycount3")
    if _familycount3 != 0 {
        args = append(args, models.Where{Column:"familycount3", Value:_familycount3, Compare:"="})    
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"="})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _personalemail := c.Get("personalemail")
    if _personalemail != "" {
        args = append(args, models.Where{Column:"personalemail", Value:_personalemail, Compare:"="})
    }
    _personalname := c.Get("personalname")
    if _personalname != "" {
        args = append(args, models.Where{Column:"personalname", Value:_personalname, Compare:"="})
    }
    _personalhp := c.Get("personalhp")
    if _personalhp != "" {
        args = append(args, models.Where{Column:"personalhp", Value:_personalhp, Compare:"="})
    }
    _zip := c.Get("zip")
    if _zip != "" {
        args = append(args, models.Where{Column:"zip", Value:_zip, Compare:"="})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _address2 := c.Get("address2")
    if _address2 != "" {
        args = append(args, models.Where{Column:"address2", Value:_address2, Compare:"="})
    }
    _contracttype := c.Geti("contracttype")
    if _contracttype != 0 {
        args = append(args, models.Where{Column:"contracttype", Value:_contracttype, Compare:"="})    
    }
    _contractprice := c.Get("contractprice")
    if _contractprice != "" {
        args = append(args, models.Where{Column:"contractprice", Value:_contractprice, Compare:"="})
    }
    _testdate := c.Get("testdate")
    if _testdate != "" {
        args = append(args, models.Where{Column:"testdate", Value:_testdate, Compare:"="})
    }
    _nexttestdate := c.Get("nexttestdate")
    if _nexttestdate != "" {
        args = append(args, models.Where{Column:"nexttestdate", Value:_nexttestdate, Compare:"="})
    }
    _repair := c.Get("repair")
    if _repair != "" {
        args = append(args, models.Where{Column:"repair", Value:_repair, Compare:"="})
    }
    _safety := c.Get("safety")
    if _safety != "" {
        args = append(args, models.Where{Column:"safety", Value:_safety, Compare:"="})
    }
    _fault := c.Get("fault")
    if _fault != "" {
        args = append(args, models.Where{Column:"fault", Value:_fault, Compare:"="})
    }
    _contractdate := c.Get("contractdate")
    if _contractdate != "" {
        args = append(args, models.Where{Column:"contractdate", Value:_contractdate, Compare:"="})
    }
    _contractduration := c.Get("contractduration")
    if _contractduration != "" {
        args = append(args, models.Where{Column:"contractduration", Value:_contractduration, Compare:"="})
    }
    _invoice := c.Get("invoice")
    if _invoice != "" {
        args = append(args, models.Where{Column:"invoice", Value:_invoice, Compare:"="})
    }
    _depositdate := c.Get("depositdate")
    if _depositdate != "" {
        args = append(args, models.Where{Column:"depositdate", Value:_depositdate, Compare:"="})
    }
    _fmsloginid := c.Get("fmsloginid")
    if _fmsloginid != "" {
        args = append(args, models.Where{Column:"fmsloginid", Value:_fmsloginid, Compare:"="})
    }
    _fmspasswd := c.Get("fmspasswd")
    if _fmspasswd != "" {
        args = append(args, models.Where{Column:"fmspasswd", Value:_fmspasswd, Compare:"="})
    }
    _facilitydivision := c.Geti("facilitydivision")
    if _facilitydivision != 0 {
        args = append(args, models.Where{Column:"facilitydivision", Value:_facilitydivision, Compare:"="})    
    }
    _facilitycategory := c.Geti("facilitycategory")
    if _facilitycategory != 0 {
        args = append(args, models.Where{Column:"facilitycategory", Value:_facilitycategory, Compare:"="})    
    }
    _position := c.Get("position")
    if _position != "" {
        args = append(args, models.Where{Column:"position", Value:_position, Compare:"="})
    }
    _area := c.Get("area")
    if _area != "" {
        args = append(args, models.Where{Column:"area", Value:_area, Compare:"="})
    }
    _groundfloor := c.Geti("groundfloor")
    if _groundfloor != 0 {
        args = append(args, models.Where{Column:"groundfloor", Value:_groundfloor, Compare:"="})    
    }
    _undergroundfloor := c.Geti("undergroundfloor")
    if _undergroundfloor != 0 {
        args = append(args, models.Where{Column:"undergroundfloor", Value:_undergroundfloor, Compare:"="})    
    }
    _useapproval := c.Get("useapproval")
    if _useapproval != "" {
        args = append(args, models.Where{Column:"useapproval", Value:_useapproval, Compare:"="})
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


func (c *AptController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *AptController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewAptManager(conn)

    var args []interface{}
    
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _completeyear := c.Get("completeyear")
    if _completeyear != "" {
        args = append(args, models.Where{Column:"completeyear", Value:_completeyear, Compare:"="})
    }
    _flatcount := c.Get("flatcount")
    if _flatcount != "" {
        args = append(args, models.Where{Column:"flatcount", Value:_flatcount, Compare:"="})
    }
    _type := c.Get("type")
    if _type != "" {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})
    }
    _floor := c.Get("floor")
    if _floor != "" {
        args = append(args, models.Where{Column:"floor", Value:_floor, Compare:"="})
    }
    _familycount := c.Get("familycount")
    if _familycount != "" {
        args = append(args, models.Where{Column:"familycount", Value:_familycount, Compare:"="})
    }
    _familycount1 := c.Geti("familycount1")
    if _familycount1 != 0 {
        args = append(args, models.Where{Column:"familycount1", Value:_familycount1, Compare:"="})    
    }
    _familycount2 := c.Geti("familycount2")
    if _familycount2 != 0 {
        args = append(args, models.Where{Column:"familycount2", Value:_familycount2, Compare:"="})    
    }
    _familycount3 := c.Geti("familycount3")
    if _familycount3 != 0 {
        args = append(args, models.Where{Column:"familycount3", Value:_familycount3, Compare:"="})    
    }
    _tel := c.Get("tel")
    if _tel != "" {
        args = append(args, models.Where{Column:"tel", Value:_tel, Compare:"="})
    }
    _fax := c.Get("fax")
    if _fax != "" {
        args = append(args, models.Where{Column:"fax", Value:_fax, Compare:"="})
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _personalemail := c.Get("personalemail")
    if _personalemail != "" {
        args = append(args, models.Where{Column:"personalemail", Value:_personalemail, Compare:"="})
    }
    _personalname := c.Get("personalname")
    if _personalname != "" {
        args = append(args, models.Where{Column:"personalname", Value:_personalname, Compare:"="})
    }
    _personalhp := c.Get("personalhp")
    if _personalhp != "" {
        args = append(args, models.Where{Column:"personalhp", Value:_personalhp, Compare:"="})
    }
    _zip := c.Get("zip")
    if _zip != "" {
        args = append(args, models.Where{Column:"zip", Value:_zip, Compare:"="})
    }
    _address := c.Get("address")
    if _address != "" {
        args = append(args, models.Where{Column:"address", Value:_address, Compare:"="})
    }
    _address2 := c.Get("address2")
    if _address2 != "" {
        args = append(args, models.Where{Column:"address2", Value:_address2, Compare:"="})
    }
    _contracttype := c.Geti("contracttype")
    if _contracttype != 0 {
        args = append(args, models.Where{Column:"contracttype", Value:_contracttype, Compare:"="})    
    }
    _contractprice := c.Get("contractprice")
    if _contractprice != "" {
        args = append(args, models.Where{Column:"contractprice", Value:_contractprice, Compare:"="})
    }
    _testdate := c.Get("testdate")
    if _testdate != "" {
        args = append(args, models.Where{Column:"testdate", Value:_testdate, Compare:"="})
    }
    _nexttestdate := c.Get("nexttestdate")
    if _nexttestdate != "" {
        args = append(args, models.Where{Column:"nexttestdate", Value:_nexttestdate, Compare:"="})
    }
    _repair := c.Get("repair")
    if _repair != "" {
        args = append(args, models.Where{Column:"repair", Value:_repair, Compare:"="})
    }
    _safety := c.Get("safety")
    if _safety != "" {
        args = append(args, models.Where{Column:"safety", Value:_safety, Compare:"="})
    }
    _fault := c.Get("fault")
    if _fault != "" {
        args = append(args, models.Where{Column:"fault", Value:_fault, Compare:"="})
    }
    _contractdate := c.Get("contractdate")
    if _contractdate != "" {
        args = append(args, models.Where{Column:"contractdate", Value:_contractdate, Compare:"="})
    }
    _contractduration := c.Get("contractduration")
    if _contractduration != "" {
        args = append(args, models.Where{Column:"contractduration", Value:_contractduration, Compare:"="})
    }
    _invoice := c.Get("invoice")
    if _invoice != "" {
        args = append(args, models.Where{Column:"invoice", Value:_invoice, Compare:"="})
    }
    _depositdate := c.Get("depositdate")
    if _depositdate != "" {
        args = append(args, models.Where{Column:"depositdate", Value:_depositdate, Compare:"="})
    }
    _fmsloginid := c.Get("fmsloginid")
    if _fmsloginid != "" {
        args = append(args, models.Where{Column:"fmsloginid", Value:_fmsloginid, Compare:"="})
    }
    _fmspasswd := c.Get("fmspasswd")
    if _fmspasswd != "" {
        args = append(args, models.Where{Column:"fmspasswd", Value:_fmspasswd, Compare:"="})
    }
    _facilitydivision := c.Geti("facilitydivision")
    if _facilitydivision != 0 {
        args = append(args, models.Where{Column:"facilitydivision", Value:_facilitydivision, Compare:"="})    
    }
    _facilitycategory := c.Geti("facilitycategory")
    if _facilitycategory != 0 {
        args = append(args, models.Where{Column:"facilitycategory", Value:_facilitycategory, Compare:"="})    
    }
    _position := c.Get("position")
    if _position != "" {
        args = append(args, models.Where{Column:"position", Value:_position, Compare:"="})
    }
    _area := c.Get("area")
    if _area != "" {
        args = append(args, models.Where{Column:"area", Value:_area, Compare:"="})
    }
    _groundfloor := c.Geti("groundfloor")
    if _groundfloor != 0 {
        args = append(args, models.Where{Column:"groundfloor", Value:_groundfloor, Compare:"="})    
    }
    _undergroundfloor := c.Geti("undergroundfloor")
    if _undergroundfloor != 0 {
        args = append(args, models.Where{Column:"undergroundfloor", Value:_undergroundfloor, Compare:"="})    
    }
    _useapproval := c.Get("useapproval")
    if _useapproval != "" {
        args = append(args, models.Where{Column:"useapproval", Value:_useapproval, Compare:"="})
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
                    str += ", a_" + strings.Trim(v, " ")                
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






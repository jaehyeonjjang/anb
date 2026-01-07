package api

import (
	"repair/controllers"
	"repair/models"
)

type ContractController struct {
	controllers.Controller
}

func (c *ContractController) Post_Insert(item *models.Contract) {
	updateContract(item)
}

func (c *ContractController) Post_Update(item *models.Contract) {
	updateContract(item)
}

func (c *ContractController) Post_Delete(item *models.Contract) {
	updateContract(item)
}

func (c *ContractController) Post_Deletebatch(item *[]models.Contract) {
	if len(*item) == 0 {
		return
	}

	updateContract(&(*item)[0])
}
func updateContract(item *models.Contract) {
	conn := models.NewConnection()
	defer conn.Close()

	aptManager := models.NewAptManager(conn)
	contractManager := models.NewContractManager(conn)

	items := contractManager.Find([]interface{}{
		models.Where{Column: "apt", Value: item.Apt, Compare: "="},
	})

	value := 0

	for _, v := range items {
		value = value | v.Type
	}

	apt := aptManager.Get(item.Apt)
	apt.Contracttype = value
	aptManager.Update(apt)
}

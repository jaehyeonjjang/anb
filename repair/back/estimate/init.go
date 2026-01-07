package estimate

import (
	"fmt"
	"repair/global/time"
	"repair/models"
)

func GetEstimateNo(typeid int, t *time.Time, date string, conn *models.Connection) string {
	estimateManager := models.NewEstimateManager(conn)
	duration := time.Now().GetDurationArray()
	count := estimateManager.Count([]interface{}{
		// models.Where{Column: "type", Value: 1, Compare: "="},
		models.Where{Column: "date", Value: duration[0], Compare: ">="},
		models.Where{Column: "date", Value: date, Compare: "<"},
	})

	no := fmt.Sprintf("ANB-%v-%v", t.DateAsOnlyNumber(), count+1)

	return no
}

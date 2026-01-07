package estimate

import (
	"repair/models"
)

func Estimate(id int64, typeid int, conn *models.Connection) string {
	estimateManager := models.NewEstimateManager(conn)
	compareestimateManager := models.NewCompareestimateManager(conn)

	estimate := estimateManager.Get(id)

	if estimate == nil {
		return ""
	}

	aptManager := models.NewAptManager(conn)
	apt := aptManager.Get(estimate.Apt)

	if apt == nil {
		return ""
	}

	compareestimates := compareestimateManager.FindByEstimate(id, []any{models.Ordering("e_id")})

	switch estimate.Type {
	case 1:
		return Repair(id, typeid, conn, estimate, compareestimates, apt)
	case 3:
		return Periodic(id, typeid, conn, estimate, compareestimates, apt)
	case 7:
		return Supervision(id, typeid, conn, estimate, compareestimates, apt)
	case 10:
		return Program(id, typeid, conn, estimate, compareestimates, apt)
	default:
		return Detail(estimate.Type, id, typeid, conn, estimate, compareestimates, apt)
	}
}

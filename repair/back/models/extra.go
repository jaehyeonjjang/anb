package models

import log "github.com/sirupsen/logrus"

type EstimateExtra struct {
	Estimate Estimate          `json:"estimate"`
	Compares []Compareestimate `json:"compares"`
}

func (p *PeriodicdataManager) FakeDelete(periodic int64, blueprint int64) {
	query := "update periodicdata_tb set pd_blueprint = ?, pd_periodic = ? where pd_blueprint = ? and pd_periodic = ?"
	_, err := p.Exec(query, blueprint*-1, periodic*-1, blueprint, periodic)

	log.Println(err)
}

func (p *PeriodicdataimageManager) FakeDelete(periodic int64, periodicdata int64) {
	query := "update periodicdataimage_tb set pi_periodicdata = ?, pi_periodic = ? where pi_periodicdata = ?"
	_, err := p.Exec(query, periodicdata*-1, periodic*-1, periodicdata)

	log.Println(err)
}

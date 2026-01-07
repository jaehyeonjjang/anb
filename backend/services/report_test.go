package services_test
 
import (
    "testing"
	"anb/services"
)
 
func TestReport(t *testing.T) {
	services.MakeSummary()
	//t.Error("Wrong result")
}

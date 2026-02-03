package main

import (
	"fmt"
	"log"
	"repair/models"
)

func main() {
	// 데이터베이스 연결
	conn := models.GetConnection()
	if conn == nil {
		log.Fatal("데이터베이스 연결 실패")
	}
	defer conn.Close()

	// standardwage_tb 테이블 구조 확인
	fmt.Println("=== standardwage_tb 테이블 구조 ===")
	rows, err := conn.Query("DESCRIBE standardwage_tb")
	if err != nil {
		log.Fatal("테이블 구조 조회 실패:", err)
	}
	defer rows.Close()

	fmt.Printf("%-25s %-20s %-10s %-10s %-20s %-10s\n", "Field", "Type", "Null", "Key", "Default", "Extra")
	fmt.Println("-------------------------------------------------------------------------------------------")

	for rows.Next() {
		var field, typ, null, key, extra string
		var defaultVal *string

		if err := rows.Scan(&field, &typ, &null, &key, &defaultVal, &extra); err != nil {
			log.Fatal("Row 스캔 실패:", err)
		}

		def := "NULL"
		if defaultVal != nil {
			def = *defaultVal
		}

		fmt.Printf("%-25s %-20s %-10s %-10s %-20s %-10s\n", field, typ, null, key, def, extra)
	}

	// 실제 데이터 조회
	fmt.Println("\n=== standardwage_tb 데이터 (ID=1) ===")

	// 직접 SQL로 모든 컬럼 조회
	var id int64
	var person1, person2, person3, person4, person5, person6, person7, person8, person9, person10 int
	var techprice1, techprice2, techprice3, techprice4 int
	var financialprice1, financialprice2, financialprice3, financialprice4 int
	var directprice, printprice1, printprice2 int
	var lossprice, gasprice, travelprice int
	var travel, loss, gas, etc, danger, machine, print int
	var date string

	query := `SELECT sw_id, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5, 
	                 sw_person6, sw_person7, sw_person8, sw_person9, sw_person10,
	                 sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4,
	                 sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4,
	                 sw_directprice, sw_printprice1, sw_printprice2,
	                 sw_lossprice, sw_gasprice, sw_travelprice,
	                 sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print,
	                 sw_date
	          FROM standardwage_tb WHERE sw_id = 1`

	row := conn.Conn.QueryRow(query)
	err = row.Scan(&id, &person1, &person2, &person3, &person4, &person5,
		&person6, &person7, &person8, &person9, &person10,
		&techprice1, &techprice2, &techprice3, &techprice4,
		&financialprice1, &financialprice2, &financialprice3, &financialprice4,
		&directprice, &printprice1, &printprice2,
		&lossprice, &gasprice, &travelprice,
		&travel, &loss, &gas, &etc, &danger, &machine, &print,
		&date)

	if err != nil {
		log.Fatal("데이터 조회 실패:", err)
	}

	fmt.Printf("ID: %d\n", id)
	fmt.Printf("등록일: %s\n\n", date)

	fmt.Println("[ 인건비 단가 ]")
	fmt.Printf("  기술사 (person1): %d원\n", person1)
	fmt.Printf("  특급기술자 (person2): %d원\n", person2)
	fmt.Printf("  고급기술자 (person3): %d원\n", person3)
	fmt.Printf("  중급기술자 (person4): %d원\n", person4)
	fmt.Printf("  초급기술자 (person5): %d원 ★\n", person5)
	if person6 != 0 {
		fmt.Printf("  person6: %d원\n", person6)
	}
	if person7 != 0 {
		fmt.Printf("  person7: %d원\n", person7)
	}
	if person8 != 0 {
		fmt.Printf("  person8: %d원\n", person8)
	}
	if person9 != 0 {
		fmt.Printf("  person9: %d원\n", person9)
	}
	if person10 != 0 {
		fmt.Printf("  person10: %d원\n", person10)
	}

	fmt.Println("\n[ 기술료 ]")
	fmt.Printf("  techprice1: %d%%\n", techprice1)
	fmt.Printf("  techprice2: %d%%\n", techprice2)
	fmt.Printf("  techprice3: %d%%\n", techprice3)
	fmt.Printf("  techprice4: %d%%\n", techprice4)

	fmt.Println("\n[ 제경비 ]")
	fmt.Printf("  financialprice1: %d%%\n", financialprice1)
	fmt.Printf("  financialprice2: %d%%\n", financialprice2)
	fmt.Printf("  financialprice3: %d%%\n", financialprice3)
	fmt.Printf("  financialprice4: %d%%\n", financialprice4)

	fmt.Println("\n[ 기타 단가 ]")
	fmt.Printf("  직접경비 (directprice): %d원\n", directprice)
	fmt.Printf("  인쇄비1 (printprice1): %d원\n", printprice1)
	fmt.Printf("  인쇄비2 (printprice2): %d원\n", printprice2)
	fmt.Printf("  손실단가 (lossprice): %d원\n", lossprice)
	fmt.Printf("  가스단가 (gasprice): %d원\n", gasprice)
	fmt.Printf("  출장단가 (travelprice): %d원\n", travelprice)

	fmt.Println("\n[ 비율/계수 ]")
	fmt.Printf("  출장 (travel): %d\n", travel)
	fmt.Printf("  손실 (loss): %d\n", loss)
	fmt.Printf("  가스 (gas): %d\n", gas)
	fmt.Printf("  기타 (etc): %d%%\n", etc)
	fmt.Printf("  위험 (danger): %d%%\n", danger)
	fmt.Printf("  기계 (machine): %d%%\n", machine)
	fmt.Printf("  인쇄 (print): %d\n", print)

	// 웹사이트 표시 금액과 비교
	webAmount := 235459
	dbAmount := person5
	if webAmount != dbAmount {
		diff := webAmount - dbAmount
		percentage := float64(diff) / float64(dbAmount) * 100
		fmt.Printf("\n[ 웹사이트 금액 차이 분석 ]\n")
		fmt.Printf("  DB 금액: %d원\n", dbAmount)
		fmt.Printf("  웹 금액: %d원\n", webAmount)
		fmt.Printf("  차이: %d원 (%.2f%%)\n", diff, percentage)
	}

	fmt.Println("\n완료!")
}

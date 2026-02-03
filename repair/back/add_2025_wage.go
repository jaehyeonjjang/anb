package main

import (
	"fmt"
	"log"
	"repair/models"
)

func main() {
	conn := models.GetConnection()
	if conn == nil {
		log.Fatal("데이터베이스 연결 실패")
	}
	defer conn.Close()
	
	// 먼저 ID=2가 이미 있는지 확인
	var count int
	err := conn.Conn.QueryRow("SELECT COUNT(*) FROM standardwage_tb WHERE sw_id = 2").Scan(&count)
	if err != nil {
		log.Fatal("확인 실패:", err)
	}
	
	if count > 0 {
		fmt.Println("ID=2 단가가 이미 존재합니다. 업데이트를 진행합니다.")
		
		query := `UPDATE standardwage_tb SET
			sw_person1 = 455276,
			sw_person2 = 377190,
			sw_person3 = 316791,
			sw_person4 = 298943,
			sw_person5 = 235459,
			sw_person6 = 455276,
			sw_person7 = 377190,
			sw_person8 = 316791,
			sw_person9 = 298943,
			sw_person10 = 235459,
			sw_date = '2025-01-01 00:00:00'
		WHERE sw_id = 2`
		
		_, err = conn.Exec(query)
		if err != nil {
			log.Fatal("업데이트 실패:", err)
		}
		
		fmt.Println("✅ 25년도 단가 업데이트 완료!")
	} else {
		fmt.Println("25년도 단가를 새로 추가합니다...")
		
		// 24년도 단가를 복사해서 5.28% 인상
		query := `INSERT INTO standardwage_tb (
			sw_id, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5,
			sw_person6, sw_person7, sw_person8, sw_person9, sw_person10,
			sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4,
			sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4,
			sw_directprice, sw_printprice1, sw_printprice2,
			sw_lossprice, sw_gasprice, sw_travelprice,
			sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print,
			sw_date
		) VALUES (
			2, 
			455276, 377190, 316791, 298943, 235459,
			455276, 377190, 316791, 298943, 235459,
			20, 20, 5, 5,
			90, 110, 5, 5,
			100000, 100000, 30000,
			3651, 1580, 20000,
			1, 1, 1, 10, 8, 3, 1,
			'2025-01-01 00:00:00'
		)`
		
		_, err = conn.Exec(query)
		if err != nil {
			log.Fatal("추가 실패:", err)
		}
		
		fmt.Println("✅ 25년도 단가 추가 완료!")
	}
	
	// ID=1의 날짜도 업데이트
	_, err = conn.Exec("UPDATE standardwage_tb SET sw_date = '2024-01-01 00:00:00' WHERE sw_id = 1")
	if err != nil {
		log.Fatal("24년도 날짜 업데이트 실패:", err)
	}
	
	fmt.Println("✅ 24년도 단가 날짜 업데이트 완료!")
	
	// 결과 확인
	fmt.Println("\n=== 등록된 단가 목록 ===")
	rows, err := conn.Query(`
		SELECT sw_id, sw_person5, sw_date 
		FROM standardwage_tb 
		ORDER BY sw_id
	`)
	if err != nil {
		log.Fatal("조회 실패:", err)
	}
	defer rows.Close()
	
	fmt.Printf("%-10s %-20s %-25s\n", "ID", "초급기술자 단가", "등록일")
	fmt.Println("--------------------------------------------------------")
	
	for rows.Next() {
		var id int64
		var person5 int
		var date string
		
		rows.Scan(&id, &person5, &date)
		
		year := ""
		if len(date) >= 4 {
			year = date[0:4] + "년도"
		}
		
		fmt.Printf("%-10d %-20d %-25s (%s)\n", id, person5, date, year)
	}
	
	fmt.Println("\n완료! 이제 웹 화면에서 견적 작성 시 단가를 선택할 수 있습니다.")
}

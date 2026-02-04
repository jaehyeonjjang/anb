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

	fmt.Println("=== 기존 단가 데이터 ===")
	rows, err := conn.Query("SELECT sw_id, sw_date, sw_person5 FROM standardwage_tb ORDER BY sw_date")
	if err != nil {
		log.Fatal("조회 실패:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var date string
		var person5 int
		rows.Scan(&id, &date, &person5)
		fmt.Printf("ID=%d, 날짜=%s, 초급기술자=%d\n", id, date, person5)
	}

	// ID=1을 2024년도로 설정
	fmt.Println("\n=== ID=1을 2024년도로 설정 ===")
	_, err = conn.Exec("UPDATE standardwage_tb SET sw_date = '2024-01-01 00:00:00' WHERE sw_id = 1")
	if err != nil {
		log.Fatal("ID=1 업데이트 실패:", err)
	}
	fmt.Println("✅ ID=1 날짜 업데이트 완료")

	// ID=2 존재 여부 확인
	var count int
	err = conn.Conn.QueryRow("SELECT COUNT(*) FROM standardwage_tb WHERE sw_id = 2").Scan(&count)
	if err != nil {
		log.Fatal("확인 실패:", err)
	}

	if count > 0 {
		fmt.Println("\n=== ID=2 존재, 2025년도 단가로 업데이트 ===")
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
		fmt.Println("✅ 2025년도 단가 업데이트 완료!")
	} else {
		fmt.Println("\n=== ID=2 없음, 2025년도 단가 신규 추가 ===")
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
			2, 455276, 377190, 316791, 298943, 235459,
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
		fmt.Println("✅ 2025년도 단가 추가 완료!")
	}

	// 결과 확인
	fmt.Println("\n=== 최종 단가 데이터 ===")
	rows, err = conn.Query(`
		SELECT sw_id, sw_date, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5 
		FROM standardwage_tb 
		ORDER BY sw_date
	`)
	if err != nil {
		log.Fatal("조회 실패:", err)
	}
	defer rows.Close()

	fmt.Printf("%-5s %-20s %-10s %-10s %-10s %-10s %-10s\n", "ID", "날짜", "기술사", "특급", "고급", "중급", "초급")
	fmt.Println("--------------------------------------------------------------------------------")
	for rows.Next() {
		var id int
		var date string
		var p1, p2, p3, p4, p5 int
		rows.Scan(&id, &date, &p1, &p2, &p3, &p4, &p5)
		fmt.Printf("%-5d %-20s %-10d %-10d %-10d %-10d %-10d\n", id, date, p1, p2, p3, p4, p5)
	}
}

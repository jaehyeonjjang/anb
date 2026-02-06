package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 프로덕션 DB 연결 정보
	connectionString := "repair:repairdb@tcp(10.34.96.4:3306)/repair?charset=utf8mb4&parseTime=True&loc=Local"
	
	fmt.Println("⚠️  프로덕션 DB에 2024년도 단가 데이터를 추가합니다.")
	fmt.Println("연결 정보를 확인해주세요:")
	fmt.Println(connectionString)
	fmt.Print("\n계속하려면 'yes'를 입력하세요: ")
	
	var confirm string
	fmt.Scanln(&confirm)
	
	if confirm != "yes" {
		fmt.Println("작업이 취소되었습니다.")
		return
	}

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}
	defer db.Close()

	// 연결 테스트
	err = db.Ping()
	if err != nil {
		log.Fatal("DB Ping 실패:", err)
	}
	fmt.Println("✅ DB 연결 성공")

	// 현재 데이터 확인
	fmt.Println("\n📊 현재 standardwage_tb 데이터:")
	rows, err := db.Query("SELECT sw_id, sw_date, sw_person5 FROM standardwage_tb ORDER BY sw_date DESC")
	if err != nil {
		log.Fatal("조회 실패:", err)
	}
	
	var count int
	for rows.Next() {
		var id int
		var date string
		var person5 int
		rows.Scan(&id, &date, &person5)
		fmt.Printf("  ID: %d, 날짜: %s, 초급기술자: %d원\n", id, date, person5)
		count++
	}
	rows.Close()
	
	if count >= 2 {
		fmt.Println("\n⚠️  이미 2개 이상의 레코드가 존재합니다. 작업을 중단합니다.")
		return
	}

	// 2024년도 데이터 INSERT
	fmt.Println("\n📝 2024년도 단가 데이터 삽입 중...")
	
	query := `INSERT INTO standardwage_tb (
		sw_date, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5,
		sw_person6, sw_person7, sw_person8, sw_person9, sw_person10,
		sw_techprice1, sw_techprice2, sw_techprice3, sw_techprice4,
		sw_financialprice1, sw_financialprice2, sw_financialprice3, sw_financialprice4,
		sw_directprice, sw_printprice1, sw_printprice2, sw_lossprice, sw_gasprice, sw_travelprice,
		sw_travel, sw_loss, sw_gas, sw_etc, sw_danger, sw_machine, sw_print
	) VALUES (
		'2024-01-01 00:00:00',
		452718, 358273, 300980, 284046, 223644,
		452718, 358273, 300980, 284046, 223644,
		20, 25, 30, 35,
		110, 120, 130, 140,
		1000000, 500000, 300000, 50000, 100000, 150000,
		0, 0, 0, 0, 0, 0, 0
	)`

	result, err := db.Exec(query)
	if err != nil {
		log.Fatal("INSERT 실패:", err)
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	
	fmt.Printf("✅ 삽입 완료! (영향받은 행: %d, 삽입된 ID: %d)\n", rowsAffected, lastInsertId)

	// 삽입 결과 확인
	fmt.Println("\n📊 삽입 후 standardwage_tb 데이터:")
	rows, err = db.Query("SELECT sw_id, sw_date, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5 FROM standardwage_tb ORDER BY sw_date DESC")
	if err != nil {
		log.Fatal("조회 실패:", err)
	}
	
	for rows.Next() {
		var id int
		var date string
		var person1, person2, person3, person4, person5 int
		rows.Scan(&id, &date, &person1, &person2, &person3, &person4, &person5)
		fmt.Printf("  ID: %d, 날짜: %s\n", id, date)
		fmt.Printf("    기술사: %d원, 특급: %d원, 고급: %d원, 중급: %d원, 초급: %d원\n", 
			person1, person2, person3, person4, person5)
	}
	rows.Close()

	fmt.Println("\n✅ 모든 작업이 완료되었습니다!")
	fmt.Println("이제 프론트엔드를 배포하면 전년도 단가 기능을 사용할 수 있습니다.")
}

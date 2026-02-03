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

	// standardhistory_tb 테이블 구조 확인
	fmt.Println("=== standardhistory_tb 테이블 구조 ===")
	rows, err := conn.Query("DESCRIBE standardhistory_tb")
	if err != nil {
		log.Fatal("테이블 구조 조회 실패:", err)
	}

	fmt.Printf("%-30s %-20s %-10s %-10s\n", "Field", "Type", "Null", "Key")
	fmt.Println("-----------------------------------------------------------------------")

	for rows.Next() {
		var field, typ, null, key, extra string
		var defaultVal *string
		rows.Scan(&field, &typ, &null, &key, &defaultVal, &extra)
		fmt.Printf("%-30s %-20s %-10s %-10s\n", field, typ, null, key)
	}
	rows.Close()

	// 데이터 조회 - 최근 변경 내역
	fmt.Println("\n=== standardhistory_tb 최근 데이터 (최근 10개) ===")
	rows, err = conn.Query("SELECT * FROM standardhistory_tb ORDER BY sh_id DESC LIMIT 10")
	if err != nil {
		log.Fatal("데이터 조회 실패:", err)
	}

	// 컬럼 정보 가져오기
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("컬럼 정보 조회 실패:", err)
	}

	// 값을 저장할 슬라이스 생성
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	count := 0
	for rows.Next() {
		err = rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal("Row 스캔 실패:", err)
		}

		if count == 0 {
			// 헤더 출력
			for _, col := range columns {
				fmt.Printf("%-15s ", col)
			}
			fmt.Println()
			fmt.Println(string(make([]byte, len(columns)*16)))
		}

		// 값 출력
		for _, val := range values {
			switch v := val.(type) {
			case nil:
				fmt.Printf("%-15s ", "NULL")
			case []byte:
				fmt.Printf("%-15s ", string(v))
			default:
				fmt.Printf("%-15v ", v)
			}
		}
		fmt.Println()
		count++
	}
	rows.Close()

	if count == 0 {
		fmt.Println("데이터가 없습니다.")
	}

	// 24년도, 25년도 단가 찾기
	fmt.Println("\n=== 년도별 초급기술자(person5) 단가 변경 이력 ===")
	query := `SELECT sh_id, sh_person5, sh_date, sh_year 
	          FROM standardhistory_tb 
	          ORDER BY sh_date DESC, sh_id DESC 
	          LIMIT 20`

	rows, err = conn.Query(query)
	if err != nil {
		// sh_year 컬럼이 없을 수도 있음
		query = `SELECT sh_id, sh_person5, sh_date 
		         FROM standardhistory_tb 
		         ORDER BY sh_date DESC, sh_id DESC 
		         LIMIT 20`
		rows, err = conn.Query(query)
		if err != nil {
			fmt.Println("조회 실패 또는 테이블이 비어있습니다:", err)
			return
		}
	}

	fmt.Printf("%-10s %-20s %-25s\n", "ID", "초급기술자 단가", "등록/수정일")
	fmt.Println("--------------------------------------------------------")

	for rows.Next() {
		var id int64
		var person5 int
		var date string

		err = rows.Scan(&id, &person5, &date)
		if err != nil {
			log.Fatal("Row 스캔 실패:", err)
		}

		fmt.Printf("%-10d %-20d %-25s\n", id, person5, date)

		if person5 == 235459 {
			fmt.Printf("  → 25년도 단가 발견!\n")
		} else if person5 == 223644 {
			fmt.Printf("  → 24년도 단가 발견!\n")
		}
	}
	rows.Close()

	fmt.Println("\n완료!")
}

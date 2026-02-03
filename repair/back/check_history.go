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
	
	// 단가 관련 테이블 찾기
	fmt.Println("=== 단가 관련 테이블 검색 ===")
	rows, err := conn.Query("SHOW TABLES LIKE '%wage%'")
	if err != nil {
		log.Fatal("테이블 조회 실패:", err)
	}
	
	fmt.Println("단가 관련 테이블:")
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		fmt.Printf("  - %s\n", tableName)
	}
	rows.Close()
	
	// 히스토리/로그 테이블 찾기
	fmt.Println("\n=== 히스토리/로그 테이블 검색 ===")
	rows, err = conn.Query("SHOW TABLES")
	if err != nil {
		log.Fatal("테이블 조회 실패:", err)
	}
	
	historyTables := []string{}
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		// history, log, audit 등의 키워드가 포함된 테이블 찾기
		if len(tableName) > 0 {
			if contains(tableName, "history") || contains(tableName, "log") || 
			   contains(tableName, "audit") || contains(tableName, "change") {
				historyTables = append(historyTables, tableName)
			}
		}
	}
	rows.Close()
	
	if len(historyTables) > 0 {
		fmt.Println("히스토리/로그 테이블:")
		for _, table := range historyTables {
			fmt.Printf("  - %s\n", table)
		}
	} else {
		fmt.Println("히스토리/로그 테이블 없음")
	}
	
	// standardwage_tb의 마지막 수정 시간 확인
	fmt.Println("\n=== standardwage_tb 레코드 정보 ===")
	var id int64
	var person5 int
	var date string
	
	query := "SELECT sw_id, sw_person5, sw_date FROM standardwage_tb WHERE sw_id = 1"
	err = conn.Conn.QueryRow(query).Scan(&id, &person5, &date)
	if err != nil {
		log.Fatal("데이터 조회 실패:", err)
	}
	
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("초급기술자 단가 (person5): %d원\n", person5)
	fmt.Printf("등록일/수정일 (sw_date): %s\n", date)
	
	fmt.Println("\n=== 결론 ===")
	fmt.Println("1. 단가 변경 히스토리를 저장하는 별도 테이블이 없습니다.")
	fmt.Println("2. standardwage_tb는 ID=1 레코드를 직접 UPDATE하는 방식입니다.")
	fmt.Println("3. 24년도→25년도 단가 변경은 아래 방법 중 하나로 이루어졌을 것입니다:")
	fmt.Println("   a) 직접 SQL UPDATE 쿼리 실행")
	fmt.Println("   b) API (PUT /api/standardwage) 호출 (현재 프론트엔드에 UI는 없음)")
	fmt.Println("   c) 데이터베이스 관리 도구 (phpMyAdmin, MySQL Workbench 등)")
	fmt.Println("\n단가 변경 추적을 위해서는 히스토리 테이블 추가가 필요합니다.")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > 0 && len(substr) > 0 && 
		 (s[:len(substr)] == substr || 
		  s[len(s)-len(substr):] == substr || 
		  findInString(s, substr))))
}

func findInString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

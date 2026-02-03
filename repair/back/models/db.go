// ============================================
// Package models - 데이터베이스 연결 및 관리
// ============================================
// 이 파일은 데이터베이스 연결, 트랜잭션 관리, 공통 타입 정의를 담당합니다.
// PostgreSQL/MySQL 데이터베이스를 지원하며, 연결 풀링과 재시도 로직을 포함합니다.
// ============================================

package models

import (
	"fmt"
	"repair/global/config"
	"repair/global/log"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버
	_ "github.com/lib/pq"              // PostgreSQL 드라이버
)

// ============================================
// 공통 쿼리 옵션 타입들
// ============================================

// PagingType - 페이징 옵션 (페이지 번호, 페이지 크기)
type PagingType struct {
	Page     int // 현재 페이지 번호 (1부터 시작)
	Pagesize int // 페이지당 항목 수
}

// OrderingType - 정렬 옵션
type OrderingType struct {
	Order string // 정렬 조건 (예: "id DESC", "name ASC")
}

// LimitType - 결과 개수 제한
type LimitType struct {
	Limit int // 최대 조회 개수
}

// OptionType - 통합 쿼리 옵션
type OptionType struct {
	Page     int
	Pagesize int
	Order    string
	Limit    int
}

// Where - WHERE 절 조건
type Where struct {
	Column  string      // 컬럼 이름
	Value   interface{} // 비교 값
	Compare string      // 비교 연산자 (=, >, <, LIKE 등)
}

// Custom - 사용자 정의 쿼리
type Custom struct {
	Query string
}

// Base - 기본 쿼리
type Base struct {
	Query string
}

// Groupby - GROUP BY 결과
type Groupby struct {
	Value int `json:"value"` // 그룹 값
	Count int `json:"count"` // 그룹 항목 개수
}

// ============================================
// 헬퍼 함수들
// ============================================

// Paging - 페이징 옵션 생성
// 사용: Paging(1, 20) -> 1페이지, 페이지당 20개
func Paging(page int, pagesize int) PagingType {
	return PagingType{Page: page, Pagesize: pagesize}
}

// Ordering - 정렬 옵션 생성
// 사용: Ordering("id DESC") -> id 내림차순 정렬
func Ordering(order string) OrderingType {
	return OrderingType{Order: order}
}

// Limit - 결과 개수 제한 생성
// 사용: Limit(100) -> 최대 100개 조회
func Limit(limit int) LimitType {
	return LimitType{Limit: limit}
}

// ============================================
// Connection - 데이터베이스 연결 관리
// ============================================
type Connection struct {
	Conn        *sql.DB // 데이터베이스 연결
	Tx          *sql.Tx // 트랜잭션 객체
	Transaction bool    // 트랜잭션 활성화 여부
	Isolation   bool    // 격리 수준 설정 여부
}

// Close - 데이터베이스 연결 종료
func (c *Connection) Close() {
	c.Conn.Close()
}

// IsConnect - 연결 상태 확인
func (c *Connection) IsConnect() bool {
	return c.Conn != nil
}

// Exec - SQL 실행 (INSERT, UPDATE, DELETE)
// 트랜잭션 활성화 시 트랜잭션 내에서 실행
func (c *Connection) Exec(query string, params ...interface{}) (sql.Result, error) {
	if c.Transaction {
		return c.Tx.Exec(query, params...)
	} else {
		return c.Conn.Exec(query, params...)
	}
}

// Query - SQL 조회 (SELECT)
// 트랜잭션 활성화 시 트랜잭션 내에서 실행
func (c *Connection) Query(query string, params ...interface{}) (*sql.Rows, error) {
	if c.Transaction {
		return c.Tx.Query(query, params...)
	} else {
		return c.Conn.Query(query, params...)
	}
}

// Begin - 트랜잭션 시작
// 여러 쿼리를 하나의 단위로 묶어서 실행 (All or Nothing)
func (c *Connection) Begin() {
	if c.Transaction {
		return
	}

	c.Tx, _ = c.Conn.Begin()
	c.Transaction = true
	c.Isolation = true
}

// Commit - 트랜잭션 커밋 (변경사항 확정)
func (c *Connection) Commit() error {
	c.Transaction = false
	return c.Tx.Commit()
}

// Rollback - 트랜잭션 롤백 (변경사항 취소)
// 오류 발생 시 모든 변경사항을 되돌립니다.
func (c *Connection) Rollback() {
	if !c.Transaction {
		return
	}

	err := c.Tx.Rollback()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	c.Transaction = false
}

// ============================================
// GetConnection - 데이터베이스 연결 생성
// ============================================
// config에 설정된 DB 정보로 연결을 생성합니다.
// 연결 풀 설정:
// - 최대 열린 연결: 100개
// - 최대 유휴 연결: 10개
// - 연결 최대 수명: 5분
func GetConnection() *Connection {
	conn, err := sql.Open(config.Database.TypeString, config.Database.ConnectionString)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}

	conn.SetMaxOpenConns(100)                // 동시에 열 수 있는 최대 연결 수
	conn.SetMaxIdleConns(10)                 // 유휴 상태로 유지할 최대 연결 수
	conn.SetConnMaxLifetime(5 * time.Minute) // 연결 재사용 최대 시간

	return &Connection{
		Conn:        conn,
		Tx:          nil,
		Transaction: false,
	}
}

// ============================================
// NewConnection - 재시도 로직이 포함된 연결 생성
// ============================================
// 연결 실패 시 자동으로 재시도합니다.
// 재시도 간격: 100ms -> 500ms -> 1s -> 2s
// 최대 5회 시도 후 실패하면 nil 반환
func NewConnection() *Connection {
	db := GetConnection()

	if db != nil {
		return db
	}

	// 1차 재시도 (100ms 대기)
	time.Sleep(100 * time.Millisecond)
	db = GetConnection()

	if db != nil {
		return db
	}

	// 2차 재시도 (500ms 대기)
	time.Sleep(500 * time.Millisecond)
	db = GetConnection()

	if db != nil {
		return db
	}

	// 3차 재시도 (1초 대기)
	time.Sleep(1 * time.Second)
	db = GetConnection()

	if db != nil {
		return db
	}

	// 4차 재시도 (2초 대기)
	time.Sleep(2 * time.Second)
	db = GetConnection()

	return db
}

// ============================================
// 레거시 쿼리 함수들
// ============================================

// QueryArray - 배열 파라미터로 쿼리 실행
func QueryArray(db *Connection, query string, items []interface{}) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	rows, err = db.Conn.Query(query, items...)
	return rows, err
}

// ExecArray - 배열 파라미터로 SQL 실행
func ExecArray(db *Connection, query string, items []interface{}) error {
	var err error

	_, err = db.Conn.Exec(query, items...)
	return err
}

// ============================================
// 유틸리티 함수들
// ============================================

// InitDate - 초기 날짜 값 반환
// 사용: 날짜가 설정되지 않은 경우의 기본값
func InitDate() string {
	return "1000-01-01 00:00:00"
}

// ============================================
// Double - JSON 직렬화를 위한 커스텀 float64 타입
// ============================================
// 정수형 값도 소수점 표기 (예: 10 -> 10.0)
type Double float64

// MarshalJSON - JSON 변환 시 소수점 표기 보장
// 예: 10 -> "10.0", 10.5 -> "10.5"
func (c Double) MarshalJSON() ([]byte, error) {
	if float64(c) == float64(int(c)) {
		return []byte(fmt.Sprintf("%v.0", int64(c))), nil
	}

	return []byte(fmt.Sprintf("%v", float64(c))), nil
}

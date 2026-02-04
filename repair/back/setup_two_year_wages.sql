-- 2개년도 단가 설정 SQL 스크립트
-- 사용법: mysql -h dev.netb.co.kr -u yuhki -p'yuhki123!@#' -D yuhki < setup_two_year_wages.sql

-- 기존 데이터 확인
SELECT '=== 기존 단가 데이터 ===' AS info;
SELECT sw_id, sw_date, sw_person5 AS '초급기술자' FROM standardwage_tb ORDER BY sw_date;

-- ID=1을 2024년도 단가로 업데이트 (기존 값 유지, 날짜만 명확히)
UPDATE standardwage_tb 
SET sw_date = '2024-01-01 00:00:00' 
WHERE sw_id = 1;

-- ID=2가 이미 있는지 확인하고 없으면 추가
INSERT INTO standardwage_tb (
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
  -- 2025년도 단가 (2024년도 대비 5.28% 인상)
  455276,  -- person1: 기술사
  377190,  -- person2: 특급기술자
  316791,  -- person3: 고급기술자
  298943,  -- person4: 중급기술자
  235459,  -- person5: 초급기술자
  455276,  -- person6
  377190,  -- person7
  316791,  -- person8
  298943,  -- person9
  235459,  -- person10
  20, 20, 5, 5,  -- techprice
  90, 110, 5, 5,  -- financialprice
  100000, 100000, 30000,  -- directprice, printprice
  3651, 1580, 20000,  -- lossprice, gasprice, travelprice
  1, 1, 1, 10, 8, 3, 1,  -- travel, loss, gas, etc, danger, machine, print
  '2025-01-01 00:00:00'
)
ON DUPLICATE KEY UPDATE
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
  sw_date = '2025-01-01 00:00:00';

-- 결과 확인
SELECT '=== 업데이트 후 단가 데이터 ===' AS info;
SELECT 
  sw_id AS 'ID',
  sw_date AS '연도',
  sw_person1 AS '기술사',
  sw_person2 AS '특급',
  sw_person3 AS '고급',
  sw_person4 AS '중급',
  sw_person5 AS '초급'
FROM standardwage_tb 
ORDER BY sw_date;

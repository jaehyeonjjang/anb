-- ====================================================================
-- 프로덕션 DB에 2024년도 단가 추가
-- ====================================================================
-- 데이터베이스: repair (10.34.96.4:3306)
-- 사용자: repair
-- 실행 전 반드시 현재 데이터 백업 권장!
-- ====================================================================

USE repair;

-- 1. 현재 데이터 확인
SELECT '=== 현재 standardwage_tb 데이터 ===' AS '';
SELECT sw_id, sw_date, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5 
FROM standardwage_tb 
ORDER BY sw_date DESC;

-- 2. 2024년도 데이터가 이미 있는지 확인
SELECT '=== 2024년도 데이터 존재 여부 확인 ===' AS '';
SELECT COUNT(*) as count_2024 
FROM standardwage_tb 
WHERE sw_date LIKE '2024%';

-- 3. 2024년도 데이터 INSERT
SELECT '=== 2024년도 단가 데이터 삽입 시작 ===' AS '';
INSERT INTO standardwage_tb (
    sw_date,
    sw_person1,  -- 기술사
    sw_person2,  -- 특급기술자
    sw_person3,  -- 고급기술자
    sw_person4,  -- 중급기술자
    sw_person5,  -- 초급기술자
    sw_person6,
    sw_person7,
    sw_person8,
    sw_person9,
    sw_person10,
    sw_techprice1,
    sw_techprice2,
    sw_techprice3,
    sw_techprice4,
    sw_financialprice1,
    sw_financialprice2,
    sw_financialprice3,
    sw_financialprice4,
    sw_directprice,
    sw_printprice1,
    sw_printprice2,
    sw_lossprice,
    sw_gasprice,
    sw_travelprice,
    sw_travel,
    sw_loss,
    sw_gas,
    sw_etc,
    sw_danger,
    sw_machine,
    sw_print
) VALUES (
    '2024-01-01 00:00:00',  -- 2024년도
    452718,   -- sw_person1 (기술사)
    358273,   -- sw_person2 (특급기술자)
    300980,   -- sw_person3 (고급기술자)
    284046,   -- sw_person4 (중급기술자)
    223644,   -- sw_person5 (초급기술자)
    452718,   -- sw_person6
    358273,   -- sw_person7
    300980,   -- sw_person8
    284046,   -- sw_person9
    223644,   -- sw_person10
    20,       -- sw_techprice1
    25,       -- sw_techprice2
    30,       -- sw_techprice3
    35,       -- sw_techprice4
    110,      -- sw_financialprice1
    120,      -- sw_financialprice2
    130,      -- sw_financialprice3
    140,      -- sw_financialprice4
    1000000,  -- sw_directprice
    500000,   -- sw_printprice1
    300000,   -- sw_printprice2
    50000,    -- sw_lossprice
    100000,   -- sw_gasprice
    150000,   -- sw_travelprice
    0,        -- sw_travel
    0,        -- sw_loss
   4. 삽입 결과 확인
SELECT '=== 삽입 완료! 전체 데이터 확인 ===' AS '';
SELECT sw_id, sw_date, 
       sw_person1 as '기술사', 
       sw_person2 as '특급', 
       sw_person3 as '고급', 
       sw_person4 as '중급', 
       sw_person5 as '초급'
FROM standardwage_tb 
ORDER BY sw_date DESC;

-- 5. 최종 확인: 총 레코드 수
SELECT '=== 최종 확인 ===' AS '';
SELECT COUNT(*) as '총 레코드 수' FROM standardwage_tb;
SELECT 
  SUM(CASE WHEN sw_date LIKE '2024%' THEN 1 ELSE 0 END) as '2024년도 레코드',
  SUM(CASE WHEN sw_date LIKE '2025%' THEN 1 ELSE 0 END) as '2025년도 레코드'
FROM standardwage_tb;

SELECT '=== 작업 완료! ===' AS ''
-- 3. 삽입 결과 확인
SELECT sw_id, sw_date, sw_person1, sw_person2, sw_person3, sw_person4, sw_person5 
FROM standardwage_tb 
ORDER BY sw_date DESC;

-- 4. 2개 레코드가 있는지 확인
SELECT COUNT(*) as total_records FROM standardwage_tb;

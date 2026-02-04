-- 테스트용 SQLite 데이터베이스 초기화

CREATE TABLE IF NOT EXISTS standardwage_tb (
    sw_id INTEGER PRIMARY KEY,
    sw_person1 INTEGER,
    sw_person2 INTEGER,
    sw_person3 INTEGER,
    sw_person4 INTEGER,
    sw_person5 INTEGER,
    sw_person6 INTEGER,
    sw_person7 INTEGER,
    sw_person8 INTEGER,
    sw_person9 INTEGER,
    sw_person10 INTEGER,
    sw_techprice1 INTEGER,
    sw_techprice2 INTEGER,
    sw_techprice3 INTEGER,
    sw_techprice4 INTEGER,
    sw_financialprice1 INTEGER,
    sw_financialprice2 INTEGER,
    sw_financialprice3 INTEGER,
    sw_financialprice4 INTEGER,
    sw_directprice INTEGER,
    sw_printprice1 INTEGER,
    sw_printprice2 INTEGER,
    sw_lossprice INTEGER,
    sw_gasprice INTEGER,
    sw_travelprice INTEGER,
    sw_travel INTEGER,
    sw_loss INTEGER,
    sw_gas INTEGER,
    sw_etc INTEGER,
    sw_danger INTEGER,
    sw_machine INTEGER,
    sw_print INTEGER,
    sw_date TEXT
);

-- 2024년도 단가 (ID=1)
INSERT OR REPLACE INTO standardwage_tb VALUES (
    1,
    432440, 358273, 300980, 284046, 223644,
    432440, 358273, 300980, 284046, 223644,
    20, 20, 5, 5,
    90, 110, 5, 5,
    100000, 100000, 30000,
    3468, 1501, 19000,
    1, 1, 1, 10, 8, 3, 1,
    '2024-01-01 00:00:00'
);

-- 2025년도 단가 (ID=2)
INSERT OR REPLACE INTO standardwage_tb VALUES (
    2,
    455276, 377190, 316791, 298943, 235459,
    455276, 377190, 316791, 298943, 235459,
    20, 20, 5, 5,
    90, 110, 5, 5,
    100000, 100000, 30000,
    3651, 1580, 20000,
    1, 1, 1, 10, 8, 3, 1,
    '2025-01-01 00:00:00'
);

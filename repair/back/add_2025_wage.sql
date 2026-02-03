-- 25년도 단가를 ID=2로 추가
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
  -- 25년도 단가 (24년도 대비 5.28% 인상 적용)
  455276,  -- person1: 기술사 (432440 * 1.0528)
  377190,  -- person2: 특급기술자 (358273 * 1.0528)
  316791,  -- person3: 고급기술자 (300980 * 1.0528)
  298943,  -- person4: 중급기술자 (284046 * 1.0528)
  235459,  -- person5: 초급기술자 (223644 * 1.0528)
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
);

-- ID=1의 날짜를 24년도로 명확히 표시
UPDATE standardwage_tb SET sw_date = '2024-01-01 00:00:00' WHERE sw_id = 1;

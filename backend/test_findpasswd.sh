#!/bin/bash

# 비밀번호 찾기 시뮬레이션

echo "=== 1단계: 비밀번호 찾기 요청 ==="
echo "사용자: admin (u_id=1)"
echo "이메일: admin@test.com"
echo ""

# 랜덤 토큰 생성
TOKEN=$(echo -n "admin$(date +%s)" | shasum -a 256 | awk '{print $1}')
RESET_LINK="http://localhost:9108/reset-password?token=$TOKEN"

echo "생성된 토큰: $TOKEN"
echo "재설정 링크: $RESET_LINK"
echo ""

# DB에 저장
sqlite3 anb.db "DELETE FROM findpasswd_tb WHERE fp_user = 1"
sqlite3 anb.db "INSERT INTO findpasswd_tb (fp_user, fp_link, fp_date) VALUES (1, '$RESET_LINK', datetime('now'))"

echo "=== 2단계: DB에 재설정 링크 저장 완료 ==="
sqlite3 anb.db "SELECT fp_id, fp_user, fp_link, fp_date FROM findpasswd_tb WHERE fp_user = 1"
echo ""

echo "=== 3단계: 이메일 발송 (시뮬레이션) ==="
echo "To: admin@test.com"
echo "Subject: 비밀번호 재설정 요청"
echo "Body:"
echo "안녕하세요,"
echo "비밀번호 재설정을 요청하셨습니다."
echo "아래 링크를 클릭하여 비밀번호를 재설정하세요:"
echo "$RESET_LINK"
echo ""

echo "=== 4단계: 사용자가 링크 클릭 후 새 비밀번호 설정 ==="
NEW_PASSWORD="newpass123"
NEW_PASSWORD_HASH=$(echo -n "$NEW_PASSWORD" | shasum -a 256 | awk '{print $1}')
echo "새 비밀번호: $NEW_PASSWORD"
echo "해시: $NEW_PASSWORD_HASH"
echo ""

# 비밀번호 업데이트
sqlite3 anb.db "UPDATE user_tb SET u_passwd = '$NEW_PASSWORD_HASH' WHERE u_id = 1"
sqlite3 anb.db "DELETE FROM findpasswd_tb WHERE fp_user = 1"

echo "=== 5단계: 비밀번호 변경 완료 ==="
sqlite3 anb.db "SELECT u_id, u_loginid, u_name, u_passwd FROM user_tb WHERE u_id = 1"
echo ""

echo "=== 6단계: 새 비밀번호로 로그인 테스트 ==="
echo "curl 'http://localhost:9108/api/jwt?loginid=admin&passwd=$NEW_PASSWORD'"

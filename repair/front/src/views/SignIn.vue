<!-- ============================================ -->
<!-- 로그인 화면 컴포넌트 (SignIn.vue) -->
<!-- ============================================ -->
<!-- 기능: -->
<!-- 1. 사용자 로그인 (아이디/비밀번호 입력) -->
<!-- 2. JWT 토큰 발급 및 저장 -->
<!-- 3. 모바일 앱(정기점검, 순찰) 다운로드 제공 -->
<!-- ============================================ -->

<template>
  <!-- ============================================ -->
  <!-- 로그인 폼 영역 -->
  <!-- ============================================ -->
  <div class="flex-container">
    <el-card class="box-card">
      <el-form label-width="120px">
        <!-- 로그인 아이디 입력 -->
        <el-form-item label="Loginid">
          <el-input v-model="item.loginid" />
        </el-form-item>
        
        <!-- 비밀번호 입력 (암호화 표시 + 엔터키로 로그인) -->
        <el-form-item label="Password">
          <el-input v-model="item.passwd" show-password @keypress.enter.native="clickSignin" />
        </el-form-item>
        
        <!-- 로그인 버튼 -->
        <el-button type="primary" @click="clickSignin">Sign In</el-button>
      </el-form>
    </el-card>    
  </div>

  <!-- ============================================ -->
  <!-- 프로그램 다운로드 링크 -->
  <!-- ============================================ -->
  <div style="margin-top:-80px;" @click="clickDownload">
    프로그램 다운로드
  </div>


  <!-- ============================================ -->
  <!-- 프로그램 다운로드 다이얼로그 -->
  <!-- ============================================ -->
  <!-- 정기점검 프로그램과 순찰 프로그램 APK 다운로드 제공 -->
  <el-dialog v-model="data.visibleDownload" width="800px" title="프로그램 다운로드">

    <y-table>
      <y-tr>
        <y-th style="text-align:center;">구분</y-th>
        <y-th style="text-align:center;">Version</y-th>
        <y-th style="text-align:center;">다운로드</y-th>
      </y-tr>
      
      <!-- 정기점검 프로그램 다운로드 -->
      <y-tr>
        <y-td style="text-align:center;">정기점검 프로그램</y-td>
        <y-td style="text-align:center;">V{{data.periodicProgram}}</y-td>
        <y-td style="text-align:center;"><el-button size="small" @click="clickDownloadPeriodic">다운로드</el-button></y-td>
      </y-tr>
      
      <!-- 순찰 프로그램 다운로드 -->
      <y-tr>
        <y-td style="text-align:center;">순찰 프로그램</y-td>
        <y-td style="text-align:center;">V{{data.patrolProgram}}</y-td>
        <y-td style="text-align:center;"><el-button size="small" @click="clickDownloadPatrol">다운로드</el-button></y-td>
      </y-tr>
    </y-table>
    
    <template #footer>
      <el-button size="small" @click="data.visibleDownload = false">닫기</el-button>
    </template>
  </el-dialog>
  
</template>

<script setup lang="ts">
// ============================================
// Import 영역
// ============================================
import { reactive } from 'vue'           // Vue 3 반응형 데이터
import { useStore } from 'vuex'          // Vuex 스토어 (전역 상태 관리)
import { util } from '~/global'          // 공통 유틸리티 (에러 메시지, 다운로드 등)
import { Login, Program } from "~/models" // API 모델 (로그인, 프로그램 정보)
import router from '~/router'            // Vue Router (페이지 이동)

// Vuex 스토어 인스턴스 가져오기
const store = useStore()

// ============================================
// 반응형 데이터: 로그인 폼 데이터
// ============================================
const item = reactive({
  loginid: '',  // 로그인 아이디
  passwd: ''    // 비밀번호
})

// ============================================
// 반응형 데이터: 다이얼로그 및 프로그램 버전
// ============================================
const data = reactive({
  visibleDownload: false  // 다운로드 다이얼로그 표시 여부
})



// ============================================
// 로그인 처리 함수
// ============================================
// 1. 입력 검증 (아이디, 비밀번호 필수)
// 2. 서버에 로그인 요청 (JWT 토큰 발급)
// 3. 로그인 성공 시:
//    - Vuex 스토어에 토큰과 사용자 정보 저장
//    - 메인 페이지(/)로 이동
// 4. 로그인 실패 시: 에러 메시지 표시
async function clickSignin() {
  // 아이디 입력 검증
  if (item.loginid === '') {
    util.error('로그인 아이디를 입력하세요')
    return
  }

  // 비밀번호 입력 검증
  if (item.passwd === '') {
    util.error('패스워드를 입력하세요')
    return
  }  

  // 서버에 로그인 요청 (API 호출)
  const res = await Login.login(item)
  
  if (res.code === 'ok') {
    // 로그인 성공
    store.commit('setRepair', null)        // 기존 수리 정보 초기화
    util.login(store, res)                 // 토큰과 사용자 정보 저장
    router.push('/')                       // 메인 페이지로 이동
  } else {
    // 로그인 실패
    console.log(res);
    util.error('로그인 정보가 정확하지 않습니다')
  }
}

// ============================================
// 프로그램 다운로드 다이얼로그 열기
// ============================================
// 1. 서버에서 최신 프로그램 버전 정보 조회
// 2. 정기점검 프로그램(type=1), 순찰 프로그램(type=3) 버전 저장
// 3. 다운로드 다이얼로그 표시
async function clickDownload() {
  // 최신 프로그램 정보 조회 (내림차순 정렬)
  let res = await Program.find({orderby: 'p_id desc'})

  let items = res.items

  // 프로그램 타입별 버전 정보 추출
  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    if (item.type == 1) {
      data.periodicProgram = item.version  // 정기점검 프로그램 버전
    }

    if (item.type == 3) {
      data.patrolProgram = item.version    // 순찰 프로그램 버전
    }
  }

  // 다운로드 다이얼로그 표시
  data.visibleDownload = true
}

// ============================================
// 정기점검 프로그램 APK 다운로드
// ============================================
// 예: /webdata/apk/periodic-V1.0.0.apk
function clickDownloadPeriodic() {
  let version = data.periodicProgram
  const url = `/webdata/apk/periodic-V${version}.apk`
  const filename = `ANB-정기점검프로그램-V${version}.apk`

  util.download(store, url, filename)  
}

// ============================================
// 순찰 프로그램 APK 다운로드
// ============================================
// 예: /webdata/apk/patrol-V1.0.0.apk
function clickDownloadPatrol() {
  let version = data.patrolProgram
  const url = `/webdata/apk/patrol-V${version}.apk`
  const filename = `ANB-순찰프로그램-V${version}.apk`

  util.download(store, url, filename)  
}
</script>

// ============================================
// Vue 3 애플리케이션 진입점 (Entry Point)
// ============================================

// === 핵심 Vue 모듈 ===
import { createApp } from 'vue'              // Vue 앱 생성 함수
import App from './App.vue'                  // 루트 컴포넌트
import router from './router'                // Vue Router - 페이지 라우팅 관리
import store from './store'                  // Vuex - 전역 상태 관리 (로그인 정보, 토큰 등)
import { setupGlobalDirectives } from './directives'  // 커스텀 디렉티브 설정

// === 서드파티 라이브러리 ===
import v3ImgPreview from 'v3-img-preview'    // 이미지 미리보기 라이브러리

// === Element Plus UI 프레임워크 ===
import ElementPlus from 'element-plus'       // Element Plus 메인 모듈
import 'element-plus/dist/index.css'         // Element Plus 전체 CSS 스타일
import 'element-plus/theme-chalk/src/message.scss'  // 메시지 컴포넌트 스타일
import koKR from 'element-plus/es/locale/lang/ko'   // 한국어 언어팩
import * as ElementPlusIconsVue from '@element-plus/icons-vue'  // Element Plus 아이콘 전체

// === 커스텀 스타일 ===
import '~/styles/index.scss'                 // 프로젝트 전역 스타일

// ============================================
// 애플리케이션 초기화 및 설정
// ============================================

// Vue 앱 인스턴스 생성
const app = createApp(App)

// 전역 커스텀 디렉티브 등록 (v-permission 등)
setupGlobalDirectives(app)

// Element Plus 아이콘을 전역 컴포넌트로 등록
// 사용 예: <Edit />, <Delete />, <Search /> 등
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// ============================================
// 플러그인 등록
// ============================================

app.use(router)                              // 라우터 활성화 - 페이지 이동 처리
app.use(store)                               // Vuex 스토어 활성화 - 전역 상태 관리
app.use(ElementPlus, {                       // Element Plus UI 프레임워크 활성화
  locale: koKR,                              // 한국어로 설정
})
app.use(v3ImgPreview, {                      // 이미지 미리보기 플러그인 활성화
})

// ============================================
// 앱을 DOM에 마운트 (실제 화면에 렌더링)
// ============================================
app.mount('#app')  // index.html의 <div id="app"></div>에 Vue 앱 연결

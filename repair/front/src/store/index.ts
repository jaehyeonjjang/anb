// ============================================
// Vuex 스토어 - 전역 상태 관리
// ============================================
// 이 파일은 애플리케이션의 전역 상태(로그인 정보, 토큰 등)를
// 관리하고, 상태 변경 메서드(mutations)와 조회 메서드(getters)를 제공합니다.
// vuex-persistedstate 플러그인으로 브라우저 localStorage에 자동 저장되어
// 페이지를 새로고침해도 로그인 상태가 유지됩니다.
// ============================================

import { createStore } from 'vuex'
import createPersistedState from 'vuex-persistedstate'

export default createStore({
    // ============================================
    // State: 전역 상태 정의
    // ============================================
    state: {
        token: '',        // JWT 인증 토큰 (로그인 시 서버에서 발급)
        user: null,       // 로그인한 사용자 정보 (아이디, 이름, 권한 등)
        repair: null      // 현재 선택된 수리 항목 정보
    },
    
    // ============================================
    // Mutations: 상태를 변경하는 메서드들
    // ============================================
    // 주의: mutations는 동기적으로 실행되어야 합니다.
    // state를 직접 수정하지 말고 반드시 mutations를 통해 변경해야 합니다.
    mutations: {
        // 토큰만 업데이트
        setToken(state, value) {
            state.token = value
        },
        
        // 로그인 성공 시 토큰과 사용자 정보 저장
        // 사용: store.commit('setLogin', { token: 'xxx', user: {...} })
        setLogin(state, { token, user }) {
            state.token = token
            state.user = user
        },
        
        // 로그아웃 시 모든 상태 초기화
        // 사용: store.commit('setLogout')
        setLogout(state) {
            state.token = ''
            state.repair = null
            state.user = null
        },
        
        // 현재 작업 중인 수리 항목 설정
        setRepair(state, value) {
            state.repair = value
        },
    },
    
    // ============================================
    // Getters: 상태를 조회하는 계산된 속성들
    // ============================================
    // computed와 유사하게 동작하며, 상태를 가공하여 반환합니다.
    getters: {
        // 로그인 여부 확인
        // 사용: store.getters.isLogin
        // 반환: true (로그인됨) / false (로그인 안됨)
        isLogin(state) {
            if (state.token == undefined || state.token == null || state.token == '') {
                return false
            }

            return true;
        },
        
        // JWT 토큰 조회
        // 사용: store.getters.getToken
        getToken(state) {
            return state.token
        },
        
        // 사용자 정보 조회
        // 사용: store.getters.getUser
        getUser(state) {
            return state.user
        },
        
        // 현재 수리 항목 조회
        // 사용: store.getters.getRepair
        getRepair(state) {
            return state.repair
        },
        
        // 사용자 권한 레벨 조회
        // 사용: store.getters.getLevel
        // 반환: 'none' | 'normal' | 'manager' | 'admin'
        // - none: 로그인 안됨 또는 권한 없음
        // - normal: 일반 사용자 (level 1)
        // - manager: 관리자 (level 2)
        // - admin: 최고 관리자 (level 3, 4)
        getLevel(state) {
            if (state == null) {
                return 'none'
            }

            if (state.user == null) {
                return 'none'
            }

            if (state.user.level < 1 || state.user.level > 4) {
                return 'none'
            }

            const levels = ['none', 'normal', 'manager', 'admin', 'admin']

            return levels[state.user.level]
        }
    },
    
    // ============================================
    // Plugins: Vuex 플러그인
    // ============================================
    // createPersistedState(): 상태를 localStorage에 자동 저장/복원
    // - 브라우저를 닫았다가 다시 열어도 로그인 상태 유지
    // - 기본 저장 위치: localStorage의 'vuex' 키
    plugins: [createPersistedState()]
})

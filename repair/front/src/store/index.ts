import { createStore } from 'vuex'
import createPersistedState from 'vuex-persistedstate'

export default createStore({
    state: {
        token: '',
        user: null,
        repair: null
    },
    mutations: {
        setToken(state, value) {
            state.token = value
        },
        setLogin(state, { token, user }) {
            state.token = token
            state.user = user
        },
        setLogout(state) {
            state.token = ''
            state.repair = null
            state.user = null
        },
        setRepair(state, value) {
            state.repair = value
        },
    },
    getters: {
        isLogin(state) {
            if (state.token == undefined || state.token == null || state.token == '') {
                return false
            }

            return true;
        },
        getToken(state) {
            return state.token
        },
        getUser(state) {
            return state.user
        },
        getRepair(state) {
            return state.repair
        },
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
    plugins: [createPersistedState()]
})

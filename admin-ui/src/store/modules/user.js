import { getToken, setToken, removeToken } from '@/utils/auth'
import router, { resetRouter } from '@/router'
import request from '@/utils/request'

const state = {
  token: getToken(),
  id: 0,
  company: 0,
  name: '',
  avatar: '',
  introduction: '',
  roles: [],
  isAdmin: false
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_ID: (state, id) => {
    state.id = id
  },
  SET_COMPANY: (state, company) => {
    state.company = company
  },
  SET_ISADMIN: (state, value) => {
    state.isAdmin = value
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      const params = 'loginid=' + username + '&passwd=' + password
      request.get('/api/login?' + params).then((response) => {
        const data = response.data

        if (data.id == null) {
          reject(data.code)
        } else {
          commit('SET_TOKEN', data.id)
          setToken(data.id)
          resolve()
        }
      }).catch((error) => {
        console.log('errror ----------------')
        console.log(error)
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      request.get('/api/user/' + state.token).then((response) => {
        const data = response.data

        let roles = ['editor']

        if (data.level === 1) {
          roles = ['editor']
        } else if (data.level === 2) {
          roles = ['manager']
        } else if (data.level === 3) {
          roles = ['admin']
        } else if (data.level === 4) {
          roles = ['superadmin']
        }

        commit('SET_ROLES', roles)
        commit('SET_ID', data.id)
        commit('SET_NAME', data.name)
        commit('SET_COMPANY', data.company)
        commit('SET_AVATAR', '')
        commit('SET_INTRODUCTION', '')

        if (data.level === 4) {
          commit('SET_ISADMIN', true)
        } else {
          commit('SET_ISADMIN', false)
        }

        data.roles = roles

        resolve(data)
      }).catch((error) => {
        console.log('errror ----------------')
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      request.get('/api/login/logout').then((response) => {
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        removeToken()
        resetRouter()

        // reset visited views and cached views
        // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
        dispatch('tagsView/delAllViews', null, { root: true })

        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      resolve()
    })
  },

  // dynamically modify permissions
  async changeRoles({ commit, dispatch }, roles) {
    resetRouter()

    const accessRoutes = await dispatch('permission/generateRoutes', roles, { root: true })

    router.addRoutes(accessRoutes)

    console.log(accessRoutes)

    // reset visited views and cached views
    dispatch('tagsView/delAllViews', null, { root: true })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

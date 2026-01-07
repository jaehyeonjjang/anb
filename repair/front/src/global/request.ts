import axios from 'axios'
import store from '../store'
import router from '../router'

// create an axios instance
const request = axios.create({
  baseURL: import.meta.env.VITE_API_URL, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 1000 * 60 * 10, // request timeout
})

request.interceptors.request.use(
  function (config: any) {
    config.headers.Authorization = 'Bearer ' + store.state.token
    return config
  },
  function (error) {
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
request.interceptors.response.use(
  function (response) {
    const res = response.data;

    // if the custom code is not 20000, it is judged as an error.
    /*
    if (response.status !== 200) {
      Message({
        message: res.message || 'Error',
        type: 'error',
        duration: 5 * 1000
      })

      // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (response.status === 50008 || response.status === 50012 || response.status === 50014) {
        // to re-login
        MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
          confirmButtonText: 'Re-Login',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          store.dispatch('user/resetToken').then(() => {
            location.reload()
          })
        })
      }
      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      return res
    }
    */

    return res;
  },
  function (error) {
    if (error.response && error.response.status === 401) {
      store.commit('clear')
      router.push({ name: 'SignIn' })
    }
    console.log('err' + error)
    // eslint-disable-next-line no-undef
    /*
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000,
    });
    return Promise.reject(error);
    */

    return Promise.reject(error)
  }
);

export default request

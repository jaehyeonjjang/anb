import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/* Router Modules */
/*
import componentsRouter from './modules/components'
import chartsRouter from './modules/charts'
import tableRouter from './modules/table'
import nestedRouter from './modules/nested'
*/

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
 roles: ['admin','editor']    control the page roles (you can set multiple roles)
 title: 'title'               the name show in sidebar and breadcrumb (recommend set)
 icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
 noCache: true                if set true, the page will no be cached(default is false)
 affix: true                  if set true, the tag will affix in the tags-view
 breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
 activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
 }
*/

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'ashboard',
        meta: { title: 'Home', icon: 'dashboard', affix: true }
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/aptgroup',
    component: Layout,
    children: [
      {
        path: 'aptgroup',
        component: () => import('@/views/aptgroup/aptgroup'),
        name: 'AptgroupList',
        meta: { title: '현장 관리', icon: 'list', roles: ['admin'] }
      },
      {
        path: 'aptgroup/insert',
        component: () => import('@/views/aptgroup/aptgroupInsert'),
        name: 'AptgroupInsert',
        meta: { title: '현장 등록', noCache: true, activeMenu: '/aptgroup/aptgroup', roles: ['admin'] },
        hidden: true
      },
      {
        path: 'aptgroup/update/:id(\\d+)',
        component: () => import('@/views/aptgroup/aptgroupUpdate'),
        name: 'AptgroupUpdate',
        meta: { title: '현장 수정', noCache: true, activeMenu: '/aptgroup/aptgroup', roles: ['admin'] },
        hidden: true
      }
    ]
  },
  {
    path: '/apt',
    component: Layout,
    hidden: true,
    children: [
      {
        path: 'apt',
        component: () => import('@/views/apt/apt'),
        name: 'AptList',
        meta: { title: '작업 관리', icon: 'list', roles: ['admin'] }
      },
      {
        path: 'apt/insert',
        component: () => import('@/views/apt/aptInsert'),
        name: 'AptInsert',
        meta: { title: '작업 등록', noCache: true, activeMenu: '/apt/apt', roles: ['admin'] },
        hidden: true
      },
      {
        path: 'apt/update/:id(\\d+)',
        component: () => import('@/views/apt/aptUpdate'),
        name: 'AptUpdate',
        meta: { title: '작업 수정', noCache: true, activeMenu: '/apt/apt', roles: ['admin'] },
        hidden: true
      }
    ]
  },
  {
    path: '/report',
    component: Layout,
    children: [
      {
        path: 'report',
        component: () => import('@/views/report/report'),
        name: 'Report',
        meta: { title: '보고서 관리', icon: 'list', roles: ['admin'] }
      }
    ]
  },
  {
    path: '/management',
    component: Layout,
    name: 'Management',
    meta: {
      title: '관리자',
      icon: 'el-icon-s-help',
      roles: ['superadmin', 'admin']
    },
    children: [
      {
        path: 'company/contract',
        component: () => import('@/views/management/companyContract'),
        name: 'CompanyContract',
        meta: { title: '계약 현황', icon: 'list', roles: ['admin'] }
      },
      {
        path: 'company/basic',
        component: () => import('@/views/management/companyBasic'),
        name: 'CompanyBasic',
        meta: { title: '기본정보 관리', icon: 'list', roles: ['admin'] }
      },
      {
        path: 'statuscategory',
        component: () => import('@/views/management/statuscategory'),
        name: 'StatuscategoryList',
        meta: { title: '유형분류 관리', icon: 'list', roles: ['superadmin', 'admin'] }
      },
      {
        path: 'statuscategory/insert',
        component: () => import('@/views/management/statuscategoryInsert'),
        name: 'StatuscategoryInsert',
        meta: { title: '유형분류 등록', noCache: true, activeMenu: '/management/statuscategory', roles: ['superadmin', 'admin'] },
        hidden: true
      },
      {
        path: 'statuscategory/update/:id(\\d+)',
        component: () => import('@/views/management/statuscategoryUpdate'),
        name: 'StatuscategoryUpdate',
        meta: { title: '유형분류 수정', noCache: true, activeMenu: '/management/statuscategory', roles: ['superadmin', 'admin'] },
        hidden: true
      },
      {
        path: 'status',
        component: () => import('@/views/management/status'),
        name: 'StatusList',
        meta: { title: '유형 관리', icon: 'list', roles: ['superadmin', 'admin'] }
      },
      {
        path: 'status/insert',
        component: () => import('@/views/management/statusInsert'),
        name: 'StatusInsert',
        meta: { title: '유형 등록', noCache: true, activeMenu: '/management/status', roles: ['superadmin', 'admin'] },
        hidden: true
      },
      {
        path: 'status/update/:id(\\d+)',
        component: () => import('@/views/management/statusUpdate'),
        name: 'StatusUpdate',
        meta: { title: '유형 수정', noCache: true, activeMenu: '/management/status', roles: ['superadmin', 'admin'] },
        hidden: true
      },
      {
        path: 'company',
        component: () => import('@/views/management/company'),
        name: 'CompanyList',
        meta: { title: '업체 관리', icon: 'list', roles: ['superadmin'] }
      },
      {
        path: 'company/insert',
        component: () => import('@/views/management/companyInsert'),
        name: 'CompanyInsert',
        meta: { title: '업체 등록', noCache: true, activeMenu: '/management/company', roles: ['superadmin'] },
        hidden: true
      },
      {
        path: 'company/update/:id(\\d+)',
        component: () => import('@/views/management/companyUpdate'),
        name: 'CompanyUpdate',
        meta: { title: '업체 수정', noCache: true, activeMenu: '/management/company', roles: ['superadmin'] },
        hidden: true
      },
      {
        path: 'user',
        component: () => import('@/views/management/user'),
        name: 'UserList',
        meta: { title: '사용자 관리', icon: 'list', roles: ['superadmin', 'admin'] }
      },
      {
        path: 'user/insert',
        component: () => import('@/views/management/userInsert'),
        name: 'UserInsert',
        meta: { title: '사용자 등록', noCache: true, activeMenu: '/management/user', roles: ['superadmin', 'admin'] },
        hidden: true
      },
      {
        path: 'user/update/:id(\\d+)',
        component: () => import('@/views/management/userUpdate'),
        name: 'UserUpdate',
        meta: { title: '사용자 수정', noCache: true, activeMenu: '/management/user', roles: ['superadmin', 'admin'] },
        hidden: true
      },
      {
        path: 'contract',
        component: () => import('@/views/management/contract'),
        name: 'ContractList',
        meta: { title: '사용기한 관리', icon: 'list', roles: ['superadmin'] }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router

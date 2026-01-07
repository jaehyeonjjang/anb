// Vuex 때 처럼 create* 함수를 제공한다.
import { createRouter, createWebHistory } from 'vue-router'
import store from '~/store'
import AptApt from '~/views/apt/Apt.vue'
import AptArea from '~/views/apt/Area.vue'
import AptBlueprint from '~/views/apt/Blueprint.vue'
import AptDong from '~/views/apt/Dong.vue'
import AptPatrolblueprint from '~/views/apt/Patrolblueprint.vue'
import AptUser from '~/views/apt/User.vue'
import DetailBlueprint from '~/views/detail/Blueprint.vue'
import DetailCurrent from '~/views/detail/Current.vue'
import DetailDetail from '~/views/detail/Detail.vue'
import DetailNormal from '~/views/detail/Normal.vue'
import DetailOutline from '~/views/detail/Outline.vue'
import DetailSchedule from '~/views/detail/Schedule.vue'
import DetailStruct from '~/views/detail/Struct.vue'
import DetailTechnician from '~/views/detail/Technician.vue'
import DetailUsagefloor from '~/views/detail/Usagefloor.vue'
import ManagementApt from '~/views/management/Apt.vue'
import ManagementDetailDetail from '~/views/management/detail/Detail.vue'
import ManagementPatrolPatrol from '~/views/management/patrol/Patrol.vue'
import ManagementPeriodicPeriodic from '~/views/management/periodic/Periodic.vue'
import ManagementRepairArea from '~/views/management/repair/Area.vue'
import ManagementRepairCategory from '~/views/management/repair/Category.vue'
import ManagementRepairRepair from '~/views/management/repair/Repair.vue'
import ManagementRepairReviewbasic from '~/views/management/repair/Reviewbasic.vue'
import ManagementRepairReviewcontentbasic from '~/views/management/repair/Reviewcontentbasic.vue'
import ManagementRepairStandard from '~/views/management/repair/Standard.vue'
import ManagementSales from '~/views/management/Sales.vue'
import ManagementSettingComparecompany from '~/views/management/setting/Comparecompany.vue'
import ManagementSettingDatacategory from '~/views/management/setting/Datacategory.vue'
import ManagementSettingEstimate from '~/views/management/setting/Estimate.vue'
import ManagementSettingTechnician from '~/views/management/setting/Technician.vue'
import ManagementSettingUser from '~/views/management/setting/User.vue'
import PatrolPatrol from '~/views/patrol/Patrol.vue'
import PeriodicCause from '~/views/periodic/Cause.vue'
import PeriodicCheck from '~/views/periodic/Check.vue'
import PeriodicConvert from '~/views/periodic/Convert.vue'
import PeriodicData from '~/views/periodic/Data.vue'
import PeriodicEtc from '~/views/periodic/Etc.vue'
import PeriodicImage from '~/views/periodic/Image.vue'
import PeriodicIncidental from '~/views/periodic/Incidental.vue'
import PeriodicManagebook from '~/views/periodic/Managebook.vue'
import PeriodicOpinion from '~/views/periodic/Opinion.vue'
import PeriodicOuterwall from '~/views/periodic/Outerwall.vue'
import PeriodicPeriodic from '~/views/periodic/Periodic.vue'
import PeriodicPublic from '~/views/periodic/Public.vue'
import PeriodicResult from '~/views/periodic/Result.vue'
import PeriodicSchedule from '~/views/periodic/Schedule.vue'
import PeriodicStruct from '~/views/periodic/Struct.vue'
import PeriodicUsagefloor from '~/views/periodic/Usagefloor.vue'
import PeriodicVent from '~/views/periodic/Vent.vue'
import RepairAdjust from '~/views/repair/Adjust.vue'
import RepairAdvice from '~/views/repair/Advice.vue'
import RepairApproval from '~/views/repair/Approval.vue'
import RepairArea from '~/views/repair/Area.vue'
import RepairBreakdown from '~/views/repair/Breakdown.vue'
import RepairCategory from '~/views/repair/Category.vue'
import RepairDong from '~/views/repair/Dong.vue'
import RepairException from '~/views/repair/Exception.vue'
import RepairFile from '~/views/repair/File.vue'
import RepairHistory from '~/views/repair/History.vue'
import RepairOutline from '~/views/repair/Outline.vue'
import RepairRepair from '~/views/repair/Repair.vue'
import RepairReport from '~/views/repair/Report.vue'
import RepairReview from '~/views/repair/Review.vue'
import RepairReviewcontent from '~/views/repair/Reviewcontent.vue'
import RepairSaving from '~/views/repair/Saving.vue'
import RepairStandard from '~/views/repair/Standard.vue'
import SignIn from '~/views/SignIn.vue'
import SignUp from '~/views/SignUp.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: SignIn
  },
  {
    path: '/signin',
    name: 'SignIn',
    component: SignIn
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: '/management/apt',
    name: 'ManagementApt',
    meta: { authorization: ['admin'] },
    component: ManagementApt
  },
  {
    path: '/management/sales',
    name: 'ManagementSales',
    meta: { authorization: ['admin'] },
    component: ManagementSales
  },
  {
    path: '/management/setting/user',
    name: 'ManagementSettingUser',
    meta: { authorization: ['admin'] },
    component: ManagementSettingUser
  },
  {
    path: '/management/setting/technician',
    name: 'ManagementSettingTechnician',
    meta: { authorization: ['admin'] },
    component: ManagementSettingTechnician
  },
  {
    path: '/management/setting/comparecompany',
    name: 'ManagementSettingComprecompany',
    meta: { authorization: ['admin'] },
    component: ManagementSettingComparecompany
  },
  {
    path: '/management/setting/datacategory',
    name: 'ManagementSettingDatacategory',
    meta: { authorization: ['admin'] },
    component: ManagementSettingDatacategory
  },
  {
    path: '/management/setting/estimate',
    name: 'ManagementSettingEstimate',
    meta: { authorization: ['admin'] },
    component: ManagementSettingEstimate
  },
  {
    path: '/management/repair/reviewcontentbasic',
    name: 'ManagementRepairReviewcontentbasic',
    meta: { authorization: ['admin'] },
    component: ManagementRepairReviewcontentbasic
  },
  {
    path: '/management/repair/category',
    name: 'ManagementRepairCategory',
    meta: { authorization: ['admin'] },
    component: ManagementRepairCategory
  },
  {
    path: '/management/repair/standard',
    name: 'ManagementRepairStandard',
    meta: { authorization: ['admin'] },
    component: ManagementRepairStandard
  },
  {
    path: '/management/repair/reviewbasic',
    name: 'ManagementRepairReviewbasic',
    meta: { authorization: ['admin'] },
    component: ManagementRepairReviewbasic
  },
  {
    path: '/management/repair/repair',
    name: 'ManagementRepairRepair',
    meta: { authorization: ['admin'] },
    component: ManagementRepairRepair
  },
  {
    path: '/management/repair/area',
    name: 'ManagementRepairArea',
    meta: { authorization: ['admin'] },
    component: ManagementRepairArea
  },
  {
    path: '/management/patrol/patrol',
    name: 'ManagementPatrolPatrol',
    meta: { authorization: ['admin'] },
    component: ManagementPatrolPatrol
  },
  {
    path: '/management/detail/detail',
    name: 'ManagementDetailDetail',
    meta: { authorization: ['admin'] },
    component: ManagementDetailDetail
  },
  {
    path: '/management/periodic/periodic',
    name: 'ManagementPeriodicPeriodic',
    meta: { authorization: ['admin'] },
    component: ManagementPeriodicPeriodic
  },
  {
    path: '/:apt/apt/apt',
    name: 'AptApt',
    component: AptApt
  },
  {
    path: '/:apt/apt/blueprint',
    name: 'AptBlueprint',
    component: AptBlueprint
  },
  {
    path: '/:apt/apt/patrolblueprint',
    name: 'AptPatrolblueprint',
    component: AptPatrolblueprint
  },
  {
    path: '/:apt/apt/user',
    name: 'AptUser',
    component: AptUser
  },
  {
    path: '/:apt/apt/dong',
    name: 'AptDong',
    component: AptDong
  },
  {
    path: '/:apt/apt/area',
    name: 'AptArea',
    component: AptArea
  },
  {
    path: '/:apt/repair/:id/repair',
    name: 'RepairRepair',
    component: RepairRepair
  },
  {
    path: '/:apt/repair/:id/category',
    name: 'RepairCategory',
    component: RepairCategory
  },
  {
    path: '/:apt/repair/:id/standard',
    name: 'RepairStandard',
    component: RepairStandard
  },
  {
    path: '/:apt/repair/:id/dong',
    name: 'RepairDong',
    component: RepairDong
  },
  {
    path: '/:apt/repair/:id/history',
    name: 'RepairHistory',
    component: RepairHistory
  },
  {
    path: '/:apt/repair/:id/breakdown',
    name: 'RepairBreakdown',
    component: RepairBreakdown
  },
  {
    path: '/:apt/repair/:id/area',
    name: 'RepairArea',
    component: RepairArea
  },
  {
    path: '/:apt/repair/:id/advice',
    name: 'RepairAdvice',
    component: RepairAdvice
  },
  {
    path: '/:apt/repair/:id/reviewcontent',
    name: 'RepairReviewcontent',
    component: RepairReviewcontent
  },
  {
    path: '/:apt/repair/:id/review',
    name: 'RepairReview',
    component: RepairReview
  },
  {
    path: '/:apt/repair/:id/report',
    name: 'RepairReport',
    component: RepairReport
  },
  {
    path: '/:apt/repair/:id/saving',
    name: 'RepairSaving',
    component: RepairSaving
  },
  {
    path: '/:apt/repair/:id/outline',
    name: 'RepairOutline',
    component: RepairOutline
  },
  {
    path: '/:apt/repair/:id/file',
    name: 'RepairFile',
    component: RepairFile
  },
  {
    path: '/:apt/repair/:id/exception',
    name: 'RepairException',
    component: RepairException
  },
  {
    path: '/:apt/repair/:id/approval',
    name: 'RepairApproval',
    component: RepairApproval
  },
  {
    path: '/:apt/repair/:id/adjust',
    name: 'RepairAdjust',
    component: RepairAdjust
  },
  {
    path: '/:apt/patrol/patrol',
    name: 'PatrolPatrol',
    component: PatrolPatrol
  },
  {
    path: '/:apt/detail/:id/detail',
    name: 'DetailDetail',
    component: DetailDetail
  },
  {
    path: '/:apt/detail/:id/outline',
    name: 'DetailOutline',
    component: DetailOutline
  },
  {
    path: '/:apt/detail/:id/current',
    name: 'DetailCurrent',
    component: DetailCurrent
  },
  {
    path: '/:apt/detail/:id/normal',
    name: 'DetailNormal',
    component: DetailNormal
  },
  {
    path: '/:apt/detail/:id/technician',
    name: 'DetailTechnician',
    component: DetailTechnician
  },
  {
    path: '/:apt/detail/:id/schedule',
    name: 'DetailSchedule',
    component: DetailSchedule
  },
  {
    path: '/:apt/detail/:id/blueprint',
    name: 'DetailBlueprint',
    component: DetailBlueprint
  },
  {
    path: '/:apt/detail/:id/struct',
    name: 'DetailStruct',
    component: DetailStruct
  },
  {
    path: '/:apt/detail/:id/usagefloor',
    name: 'DetailUsagefloor',
    component: DetailUsagefloor
  },
  {
    path: '/:apt/periodic/:id/periodic',
    name: 'PeriodicPeriodic',
    component: PeriodicPeriodic
  },
  {
    path: '/:apt/periodic/:id/data',
    name: 'PeriodicData',
    component: PeriodicData
  },
  {
    path: '/:apt/periodic/:id/check',
    name: 'PeriodicCheck',
    component: PeriodicCheck
  },
  {
    path: '/:apt/periodic/:id/result',
    name: 'PeriodicResult',
    component: PeriodicResult
  },
  {
    path: '/:apt/periodic/:id/incidental',
    name: 'PeriodicIncidental',
    component: PeriodicIncidental
  },
  {
    path: '/:apt/periodic/:id/image',
    name: 'PeriodicImage',
    component: PeriodicImage
  },
  {
    path: '/:apt/periodic/:id/opinion',
    name: 'PeriodicOpinion',
    component: PeriodicOpinion
  },
  {
    path: '/:apt/periodic/:id/cause',
    name: 'PeriodicCause',
    component: PeriodicCause
  },
  {
    path: '/:apt/periodic/:id/managebook',
    name: 'PeriodicManagebook',
    component: PeriodicManagebook
  },
  {
    path: '/:apt/periodic/:id/outerwall',
    name: 'PeriodicOuterwall',
    component: PeriodicOuterwall
  },
  {
    path: '/:apt/periodic/:id/vent',
    name: 'PeriodicVent',
    component: PeriodicVent
  },
  {
    path: '/:apt/periodic/:id/usagefloor',
    name: 'PeriodicUsagefloor',
    component: PeriodicUsagefloor
  },
  {
    path: '/:apt/periodic/:id/struct',
    name: 'PeriodicStrcut',
    component: PeriodicStruct
  },
  {
    path: '/:apt/periodic/:id/schedule',
    name: 'PeriodicSchedule',
    component: PeriodicSchedule
  },
  {
    path: '/:apt/periodic/:id/etc',
    name: 'PeriodicEtc',
    component: PeriodicEtc
  },
  {
    path: '/:apt/periodic/:id/public',
    name: 'PeriodicPublic',
    component: PeriodicPublic
  },
  {
    path: '/:apt/periodic/:id/convert',
    name: 'PeriodicConvert',
    component: PeriodicConvert
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


router.beforeEach(function (to, from, next) {
  const { authorization } = to.meta

  if (store.state.token == '') {
    if (to.path === '/signin') {
      next()
    } else {
      next('/signin')
    }

    return
  }

  /*
  store.commit('setLogout')
  router.push('/')
  */

  const level = store.getters['getLevel']

  if (authorization != undefined && !authorization.includes(level)) {
    next('/not-found')

    return
  }

  if (to.path === '/signin' || to.path == '/') {
    if (level == 'normal' || level == 'manager') {
      const apt = store.getters['getUser'].apt
      if (apt == undefined) {
        next('/signin')
      } else {
        next(`/${apt}/apt/apt`)
      }
    } else {
      next('/management/apt')
    }
  } else {
    next()
  }
})

export default router

<template>

  <el-row style="border-bottom: 1px solid #ccc;height:">
    <el-col :span="12" style="text-align:left;">
      <el-menu mode="horizontal" style="border:none;" v-if="data.apt.id == 0 && store.getters['getLevel'] == 'admin'">
        <el-menu-item index="7" @click="clickBase('/management/sales')" style="font-size:14px;font-weight:bold;">주소 관리</el-menu-item>
        <el-menu-item index="1" @click="clickBase('/management/apt')" style="font-size:14px;font-weight:bold;">작업 관리</el-menu-item>
        <el-menu-item index="3" @click="clickBase('/management/repair/repair')" style="font-size:14px;font-weight:bold;">장기수선</el-menu-item>
        <!-- <el-menu-item index="4" @click="clickBase('/management/detail/detail')" style="font-size:14px;font-weight:bold;">정밀점검</el-menu-item> -->
        <el-menu-item index="5" @click="clickBase('/management/periodic/periodic')" style="font-size:14px;font-weight:bold;">안전점검</el-menu-item>
        <el-menu-item index="6" @click="clickBase('/management/patrol/patrol')" style="font-size:14px;font-weight:bold;">순찰</el-menu-item>
        <el-menu-item index="2" @click="clickBase('/management/setting/user')" style="font-size:14px;font-weight:bold;">설정</el-menu-item>
      </el-menu>
      <div style="height:43px;font-size:20px;font-weight:bold;margin: 5px 0px 0px 5px;display:flex;padding-top:10px;" v-if="data.apt.id != 0">
        <div @click="clickHome" style="margin-left:10px;">
          <el-icon :size="25"><HomeFilled /></el-icon>
        </div>
        <div style="margin-left:20px;color:#666;">
          
          <el-popover :width="400" v-if="store.getters['getLevel'] == 'admin'"
                                    popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
          >
            <template #reference>
              <div @click="clickAptView">{{data.apt.name}}</div>              
            </template>
            <template #default>

              <el-table :data="[1, 2, 3, 4, 5]" :show-header="false" border>
                <el-table-column width="70" property="name" label="date">
                  <template #default="scope">
                    <div v-if="scope.$index == 0">전화번호</div>
                    <div v-if="scope.$index == 1">Fax</div>
                    <div v-if="scope.$index == 2">공용 이메일</div>
                    <div v-if="scope.$index == 3">개인 이메일</div>
                    <div v-if="scope.$index == 4">주소</div>
                  </template>
                </el-table-column>
                <el-table-column property="value" label="name">
                  <template #default="scope">
                    <div v-if="scope.$index == 0">{{data.apt.tel}}</div>
                    <div v-if="scope.$index == 1">{{data.apt.fax}}</div>
                    <div v-if="scope.$index == 2">{{data.apt.email}}</div>
                    <div v-if="scope.$index == 3">{{data.apt.personalemail}}</div>
                    <div v-if="scope.$index == 4">{{data.apt.address}}</div>
                  </template>
                </el-table-column>
              </el-table>
              
            </template>
          </el-popover>
          <div v-else>{{data.apt.name}}</div>
        </div>
        <div style="font-size:14px;color:#666;font-weight:bold;margin: 5px 0px 0px 10px;" v-if="store.getters['getLevel'] == 'admin'">
          <span v-if="data.apt.tel != ''">({{data.apt.tel}}, 준공년도 : {{data.apt.completeyear}})</span>
          <span v-else style="font-size:14px;">(준공년도 : {{data.apt.completeyear}})</span>
        </div>
        <div style="font-size:14px;color:#409eff;font-weight:bold;margin: 5px 0px 0px 10px;">          
          <div v-if="data.menu == 'repair'">            
            장기수선계획
            <span v-if="data.repair.type == 1">재수립</span>
            <span v-else>검토조정</span>
            ({{data.repair.reportdate}})
          </div>
          <div v-if="data.menu == 'apt'">
            기본관리
          </div>
          <div v-if="data.menu == 'patrol'">
            순찰
          </div>
          <!-- <div v-if="data.menu == 'detail'">
               정밀점검
               </div> -->
          <div v-if="data.menu == 'periodic'">
            <span v-if="data.periodic.category == 1">정기점검</span>
            <span v-if="data.periodic.category == 2">정밀점검</span>
          </div>
        </div>


      </div>
    </el-col>

    <el-col :span="12">
      <div style="float:right;margin-top:20px;text-align:right;font-size:12px;margin-right:20px;">
        <el-dropdown>
          <span style="font-size:12px;">            
            {{store.getters['getUser'].name}}
            <el-icon class="el-icon--right">
              <arrow-down />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item style="font-size:12px;">회원정보 수정</el-dropdown-item>
              <el-dropdown-item style="font-size:12px;" @click="clickDownload">프로그램 다운로드</el-dropdown-item>
              <el-dropdown-item style="font-size:12px;" divided @click="clickLogout" >로그아웃</el-dropdown-item>
            </el-dropdown-menu>
          </template>
      </el-dropdown>
      </div>

      <div style="width:500px;float:right;">
        <el-menu mode="horizontal" style="width:100%;" v-if="data.apt.id > 0" :ellipsis="false">
          <el-menu-item index="1" @click="clickAptApt" style="font-size:14px;font-weight:bold;">기본관리</el-menu-item>
          <el-menu-item index="2" @click="clickAptRepair" v-if="data.apt.contracttype & 1" style="font-size:14px;font-weight:bold;">장기수선</el-menu-item>
          <!-- <el-menu-item index="3" @click="clickAptDetail" v-if="data.apt.contracttype & 2" style="font-size:14px;font-weight:bold;">정밀점검</el-menu-item> -->
          <el-menu-item index="4" @click="clickAptPeriodic" v-if="data.apt.contracttype & 4" style="font-size:14px;font-weight:bold;">안전점검</el-menu-item>
          <el-menu-item index="5" @click="clickAptPatrol" v-if="data.apt.contracttype & 256" style="font-size:14px;font-weight:bold;">순찰</el-menu-item>
        </el-menu>
      </div>      
      <div style="clear:both;"></div>
    </el-col>
  </el-row>

    
  <AptInsert :id="data.apt.id" ref="apt" />

  <el-dialog v-model="data.visibleRepair" width="950px" :before-close="handleCloseRepair" >
    <RepairInsert :id="data.apt.id" ref="repairInsert" :close="closeRepair" />
  </el-dialog>

  <el-dialog v-model="data.visibleDetail" width="950px" :before-close="handleCloseDetail" >
    <DetailInsert :id="data.apt.id" ref="detailInsert" :close="closeDetail" />
  </el-dialog>

  <el-dialog v-model="data.visiblePeriodic" width="950px" :before-close="handleClosePeriodic" >
    <PeriodicInsert :id="data.apt.id" ref="periodicInsert" :close="closePeriodic" />
  </el-dialog>


  <el-dialog v-model="data.visibleDownload" width="950px" title="프로그램 다운로드">

    <y-table>
      <y-tr>
        <y-th style="text-align:center;">구분</y-th>
        <y-th style="text-align:center;">Version</y-th>
        <y-th style="text-align:center;">다운로드</y-th>
      </y-tr>
      <y-tr>
        <y-td style="text-align:center;">정기점검 프로그램</y-td>
        <y-td style="text-align:center;">V{{data.periodicProgram}}</y-td>
        <y-td style="text-align:center;"><el-button size="small" @click="clickDownloadPeriodic">다운로드</el-button></y-td>
      </y-tr>
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
  
  <InquiryInsert :close="clickClose" ref="apt" />
</template>

<script setup lang="ts">

import { ref, reactive, onMounted, computed, watch, watchEffect } from "vue"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import router from '~/router'
import { Apt, Repair, Program, Periodic } from "~/models"
import { util }  from "~/global"


const store = useStore()
const route = useRoute()

const repairInsert = ref({});
const detailInsert = ref({});
const periodicInsert = ref({});

const data = reactive({
  id: 0,
  apt: {
    id: 0,
    tel: ''
  },
  repair: {
    reportdate: ''
  },
  periodic: {
    category: 1
  },
  visible: false,
  visibleRepair: false,
  visibleDetail: false,
  visiblePeriodic: false,
  visibleDownload: false,
  menu: '',
  periodicProgram: '',
  patrolProgram: ''
})

const apt = ref({})


watch(() => route.params.id, async () => {
  data.id = util.getInt(route.params.id)
})

watch(() => route.params.apt, async () => {
  const aptid = util.getInt(route.params.apt)

  if (aptid == 0) {
    data.apt = { id: 0, tel: '' }
    return
  }

  const res = await Apt.get(aptid)
  let apt = res.item
  data.apt = apt
})

watch(() => route.path, async () => {
  var s = route.path.split('/')

  let menu = ''
  if (s[1] == 'management') {
    if (s.length == 3) {
      menu = s[1]
    } else {
      menu = `${s[1]}/${s[2]}`
    }
  } else {
    menu = s[2]
  }

  if (data.menu != menu) {
    data.menu = menu
  }

  if (menu == 'repair') {
    if (data.id == 0) {
      data.repair = { reportdate: '' }
      return
    } else {
      const res = await Repair.get(data.id)
      data.repair = res.item
    }
  } else if (menu == 'periodic') {
    if (data.id == 0) {
      data.periodic = { category: 1 }
      return
    } else {
      const res = await Periodic.get(data.id)
      data.periodic = res.item
    }
  }
})


const handleCloseRepair = (done: () => void) => {
  repairInsert.value.reset()
  done()
}

const handleCloseDetail = (done: () => void) => {
  detailInsert.value.reset()
  done()
}

const handleClosePeriodic = (done: () => void) => {
  periodicInsert.value.reset()
  done()
}

onMounted(async () => {    
})

function clickLogout() {
  store.commit('setLogout')
  router.push('/')
}

function clickHome() {
  if (store.getters['getLevel'] == 'admin') {
    router.push('/management/sales')
  } else {
    router.push('/')
  }
}

function clickBase(url) {
  router.push(url)
}

function clickAptView() {
  apt.value?.readData(data.apt.id)
}

function clickDetail() {
}

function clickPatrol() {
}

function clickAptApt() {
  router.push(`/${data.apt.id}/apt/apt`)
}

function clickAptRepair() {
  data.visibleRepair = true
}

function clickAptDetail() {
  data.visibleDetail = true
}

function clickAptPeriodic() {
  data.visiblePeriodic = true
}

function clickAptPatrol() {
  router.push(`/${data.apt.id}/patrol/patrol`)
}

function closeRepair() {
  data.visibleRepair = false
}

function closeDetail() {
  data.visibleDetail = false
}

function closePeriodic() {
  data.visiblePeriodic = false
}

async function clickDownload() {
  let res = await Program.find({orderby: 'p_id desc'})

  let items = res.items

  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    if (item.type == 1) {
      data.periodicProgram = item.version
    }

    if (item.type == 3) {
      data.patrolProgram = item.version
    }
  }

  data.visibleDownload = true
}

function clickDownloadPeriodic() {
  let version = data.periodicProgram
  const url = `/webdata/apk/periodic-V${version}.apk`
  const filename = `ANB-정기점검프로그램-V${version}.apk`

  util.download(store, url, filename)  
}

function clickDownloadPatrol() {
  let version = data.patrolProgram
  const url = `/webdata/apk/patrol-V${version}.apk`
  const filename = `ANB-순찰프로그램-V${version}.apk`

  util.download(store, url, filename)  
}

function clickClose() {
  
}
</script>

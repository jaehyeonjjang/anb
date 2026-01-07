<template>
  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="width:150px;">점검 기간</y-th>
      <y-td>{{util.viewDate(data.periodic.startdate)}} ~ {{util.viewDate(data.periodic.enddate)}} </y-td>
    </y-tr>
    <y-tr>
      <y-th>관리주체</y-th>
      <y-td>{{data.periodic.manager}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>대표자</y-th>
      <y-td>{{data.periodic.agent}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>소유주</y-th>
      <y-td>{{data.periodic.owner}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>공동수급</y-th>
      <y-td>{{data.periodic.supply}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>계약방법</y-th>
      <y-td>{{data.periodic.contract}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>시설물구분</y-th>
      <y-td>{{data.facilitytypes[data.aptperiodic.facilitytype]}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>종류</y-th>
      <y-td>{{getFacilitycategorys(data.aptperiodic.facilitycategory)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>종별</y-th>
      <y-td>{{data.facilitydivisions[data.aptperiodic.facilitydivision]}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>점검금액(천원)</y-th>
      <y-td>{{data.periodic.price}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>안전등급</y-th>
      <y-td>{{data.periodic.safetygrade}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>시설물 위치</y-th>
      <y-td>{{data.aptItem.position}}</y-td>
    </y-tr>
    <y-tr v-if="data.dongs.length == 0">
      <y-th>시설물 규모</y-th>
      <y-td>건축연면적 : {{data.aptItem.area}} (지하 {{data.aptItem.undergroundfloor}}층 / 지상 {{data.aptItem.groundfloor}}층)</y-td>
    </y-tr>

  </y-table>

  <el-dialog
    v-model="data.visible"
    title="일반 현황 수정"
    :before-close="handleClose"
  >


  <y-table>
    <y-tr>
      <y-th style="width:150px;">점검 기간</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.startdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.enddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>관리주체</y-th>
      <y-td><el-input v-model="data.item.manager" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>대표자</y-th>
      <y-td><el-input v-model="data.item.agent" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>소유주</y-th>
      <y-td><el-input v-model="data.item.owner" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>공동수급</y-th>
      <y-td><el-input v-model="data.item.supply" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>계약방법</y-th>
      <y-td>
        <el-radio-group v-model="data.item.contract">
              <el-radio-button size="small" label="수의계약">수의계약</el-radio-button>
              <el-radio-button size="small" label="입찰계약">입찰계약</el-radio-button>              
        </el-radio-group>
        
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>시설물구분</y-th>
      <y-td>
        <el-select v-model.number="data.item.facilitytype" style="width:100%;">
          <el-option v-for="(item, index) in data.facilitytypes" :key="index" :label="item" :value="index" />
        </el-select>
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>종류</y-th>
      <y-td>
        <el-select v-model.number="data.item.facilitycategory" style="width:100%;">
          <el-option v-for="item in data.facilitycategorys" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>


      </y-td>
    </y-tr>

    <y-tr>
      <y-th>종별</y-th>
      <y-td>
        <el-radio-group v-model.number="data.item.facilitydivision">
              <el-radio-button size="small" label="1">1종</el-radio-button>
              <el-radio-button size="small" label="2">2종</el-radio-button>
              <el-radio-button size="small" label="3">3종</el-radio-button>
        </el-radio-group>
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>점검금액(천원)</y-th>
      <y-td><el-input v-model="data.item.price" placeholder="" /></y-td>
    </y-tr>

    <y-tr>
      <y-th>안전등급</y-th>
      <y-td><el-input v-model="data.item.safetygrade" placeholder="" /></y-td>
    </y-tr>

    <y-tr>
      <y-th>시설물 위치</y-th>
      <y-td><el-input v-model="data.item.position" placeholder="" /></y-td>
    </y-tr>

    <y-tr v-if="data.dongs.length == 0">
      <y-th>시설물 규모</y-th>
      <y-td>건축연면적 : <el-input v-model="data.item.area" placeholder="" style="width:100px;" /> (지하 <el-input v-model.number="data.item.undergroundfloor" placeholder="" style="width:50px;" />층 / 지상 <el-input v-model.number="data.item.groundfloor" placeholder=""  style="width:50px;" />층)</y-td>
    </y-tr>

  </y-table>

  <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">저장</el-button>
    </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Aptperiodic, Periodic, Facilitycategory, Aptdong } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()


const aptItem = {
  id: 0,
  position: '',
  area: 0,
  groundfloor: 0,
  undergroundfloor: 0
}

const item = {
  id: 0,
  startdate: null,
  enddate: null,
  manager: '',
  agent: '',
  owner: '',
  supply: '',
  contract: '',
  facilitydivision: 0,
  facilitytype: 0,
  facilitycategory: 0,
  price: 0,
  safetygrade: 0,
  position: '',
  area: 0,
  groundfloor: 0,
  undergroundfloor: 0
}

const periodic = {
  id: 0,
  name: '',
  startdate: null,
  enddate: null,
  supply: '',
  contract: '',
  price: 0,
  safetygrade: 0,
  status: 0,
  apt: 0,
  date: ''
}

const aptperiodic = {
  id: 0,
  current1: '',
  current2: '',
  current3: '',
  current4: '',
  current5: '',
  current6: '',
  current7: '',
  current8: '',
  current9: '',
  current10: '',
  current11: '',
  current12: '',
  current13: '',
  current14: '',
  current15: '',
  current16: '',
  current17: '',
  current18: '',
  current19: '',
  current20: '',
  current21: '',
  current22: '',
  current23: '',
  deligate: '',
  facilitydivision: 0,
  facilitytype: 0,
  facilitycategory: 0,
  date: ''
}

const data = reactive({
  id: 0,
  aptItem: util.clone(aptItem),
  item: util.clone(item),
  periodic: util.clone(periodic),
  aptperiodic: util.clone(aptperiodic),
  visible: false,
  facilitycategorys: [],
  facilitydivisions: [' ', '1종', '2종', '3종'],
  facilitytypes: [' ', '건축물'],
  dongs: []
})

function getFacilitycategorys(value) {
  if (util.isNull(value) || value == '') {
    return ''
  }

  for (let i = 0; i < data.facilitycategorys.length; i++) {
    let item = data.facilitycategorys[i]

    if (item.id == value) {
      return item.name
    }
  }

  return ''
}

async function initData() {
  let res = await Facilitycategory.find({orderby: 'fc_order'})

  data.facilitycategorys = [{id: 0, name: ' '}, ...res.items]

  res = await Aptdong.find({apt: data.apt})
  if (res.items == null) {
    res.items = []
  }
  
  data.dongs = res.items  
}

async function getItems() {
  let res = await Aptperiodic.get(data.apt)
  const aptperiodic = res.item
  if (aptperiodic == null) {
    let newItem = util.clone(item)
    newItem.id = data.apt
    await Aptperiodic.insert(newItem)
    data.aptperiodic = newItem
  } else {
    data.aptperiodic = aptperiodic
  }

  res = await Periodic.get(data.id)
  data.periodic = res.item

  res = await Apt.get(data.apt)
  data.aptItem = res.item
}

async function clickUpdate(pos) {
  let newItem = util.clone(item)

  newItem.startdate = data.periodic.startdate
  newItem.enddate = data.periodic.enddate

  newItem.manager = data.periodic.manager
  newItem.agent = data.periodic.agent
  newItem.owner = data.periodic.owner
  newItem.supply = data.periodic.supply
  if (newItem.supply == '') {
    newItem.supply = '독자수행 100%'
  }
  newItem.contract = data.periodic.contract
  if (newItem.contract == '') {
    newItem.contract = '수의계약'
  }

  newItem.facilitydivision = data.aptperiodic.facilitydivision
  newItem.facilitytype = data.aptperiodic.facilitytype
  newItem.facilitycategory = data.aptperiodic.facilitycategory
  newItem.price = data.periodic.price
  newItem.safetygrade = data.periodic.safetygrade


  newItem.position = data.aptItem.position
  if (newItem.position == '') {
    newItem.position = data.aptItem.address
  }

  newItem.area = data.aptItem.area
  newItem.groundfloor = data.aptItem.groundfloor
  newItem.undergroundfloor = data.aptItem.undergroundfloor

  data.item = newItem
  data.visible = true
}

async function clickSubmit(type) {
  let res = await Aptperiodic.get(data.apt)
  let aptperiodic = res.item

  aptperiodic.facilitydivision = util.getInt(data.item.facilitydivision)
  aptperiodic.facilitytype = util.getInt(data.item.facilitytype)
  aptperiodic.facilitycategory = util.getInt(data.item.facilitycategory)

  await Aptperiodic.update(aptperiodic)

  res = await Periodic.get(data.id)
  let periodic = res.item

  periodic.startdate = util.convertDBDate(data.item.startdate)
  periodic.enddate = util.convertDBDate(data.item.enddate)
  periodic.supply = data.item.supply
  periodic.contract = data.item.contract
  periodic.price = data.item.price
  periodic.safetygrade = data.item.safetygrade

  periodic.manager = data.item.manager
  periodic.agent = data.item.agent
  periodic.owner = data.item.owner

  await Periodic.update(periodic)

  res = await Apt.get(data.apt)
  let apt = res.item

  apt.position = data.item.position
  apt.area = data.item.area
  apt.groundfloor = util.getInt(data.item.groundfloor)
  apt.undergroundfloor = util.getInt(data.item.undergroundfloor)

  await Apt.update(apt)

  util.info('수정되었습니다')

  await getItems()

  data.visible = false
}

const handleClose = (done: () => void) => {
  /*
     util.confirm('팝업창을 닫으시겠습니까', function() {
     done()
     })
   */

  done()
}

onMounted(async () => {
  util.loading(true)
  
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  await initData()
  await getItems()

  util.loading(false)
})

</script>

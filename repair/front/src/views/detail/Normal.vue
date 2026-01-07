<template>
  <Title title="일반 현황" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th>점검 기간</y-th>
      <y-td>{{util.viewDate(data.detail.startdate)}} ~ {{util.viewDate(data.detail.enddate)}} </y-td>
    </y-tr>
    <y-tr>
      <y-th>관리주체</y-th>
      <y-td>{{data.aptdetail.current3}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>대표자</y-th>
      <y-td>{{data.aptdetail.deligate}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>공동수급</y-th>
      <y-td>{{data.detail.supply}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>계약방법</y-th>
      <y-td>{{data.detail.contract}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>시설물구분</y-th>
      <y-td>{{data.facilitytypes[data.aptdetail.facilitytype]}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>종류</y-th>
      <y-td>{{getFacilitycategorys(data.aptdetail.facilitycategory)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>종별</y-th>
      <y-td>{{data.facilitydivisions[data.aptdetail.facilitydivision]}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>점검금액(천원)</y-th>
      <y-td>{{data.detail.price}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>안전등급</y-th>
      <y-td>{{data.detail.safetygrade}}</y-td>
    </y-tr>

  </y-table>

  <el-dialog
    v-model="data.visible"
    title="일반 현황 수정"
    :before-close="handleClose"
  >


  <y-table>
    <y-tr>
      <y-th>점검 기간</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.startdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.enddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>관리주체</y-th>
      <y-td><el-input v-model="data.item.current3" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>대표자</y-th>
      <y-td><el-input v-model="data.item.deligate" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>공동수급</y-th>
      <y-td><el-input v-model="data.item.supply" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>계약방법</y-th>
      <y-td><el-input v-model="data.item.contract" placeholder="" /></y-td>
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
import { Apt, Aptdetail, Detail, Facilitycategory } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()


const item = {
  id: 0,
  startdate: null,
  enddate: null,
  current3: '',
  deligate: '',
  supply: '',
  contract: '',
  facilitydivision: 0,
  facilitytype: 0,
  facilitycategory: 0,
  price: 0,
  safetygrade: 0
}

const detail = {
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

const aptdetail = {
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
  item: util.clone(item),
  detail: util.clone(detail),
  aptdetail: util.clone(aptdetail),
  visible: false,
  facilitycategorys: [],
  facilitydivisions: [' ', '1종', '2종', '3종'],
  facilitytypes: [' ', '건축물']
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

}

async function getItems() {
  let res = await Aptdetail.get(data.apt)
  const aptdetail = res.item
  if (aptdetail == null) {
    let newItem = util.clone(item)
    newItem.id = data.apt
    await Aptdetail.insert(newItem)
    data.aptdetail = newItem
  } else {
    data.aptdetail = aptdetail
  }

  res = await Detail.get(data.id)
  data.detail = res.item

  res = await Apt.get(data.apt)
  data.aptItem = res.item
}

async function clickUpdate(pos) {
  let newItem = util.clone(item)

  newItem.startdate = data.detail.startdate
  newItem.enddate = data.detail.enddate

  newItem.current3 = data.aptdetail.current3
  newItem.deligate = data.aptdetail.deligate
  newItem.supply = data.detail.supply
  newItem.contract = data.detail.contract

  newItem.facilitydivision = data.aptdetail.facilitydivision
  newItem.facilitytype = data.aptdetail.facilitytype
  newItem.facilitycategory = data.aptdetail.facilitycategory
  newItem.price = data.detail.price
  newItem.safetygrade = data.detail.safetygrade

  data.item = newItem
  data.visible = true
}

async function clickSubmit(type) {
  let res = await Aptdetail.get(data.apt)
  let aptdetail = res.item

  aptdetail.current3 = data.item.current3
  aptdetail.deligate = data.item.deligate
  aptdetail.facilitydivision = util.getInt(data.item.facilitydivision)
  aptdetail.facilitytype = util.getInt(data.item.facilitytype)
  aptdetail.facilitycategory = util.getInt(data.item.facilitycategory)

  console.log(aptdetail)

  await Aptdetail.update(aptdetail)

  res = await Detail.get(data.id)
  let detail = res.item

  detail.startdate = util.convertDBDate(data.item.startdate)
  detail.enddate = util.convertDBDate(data.item.enddate)
  detail.supply = data.item.supply
  detail.contract = data.item.contract
  detail.price = util.getInt(data.item.price)
  detail.safetygrade = data.item.safetygrade

  await Detail.update(detail)

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
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  await initData()
  await getItems()
})

</script>

<template>
  <Title title="시설물 현황" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th>시설물번호</y-th>
      <y-td>{{data.detail.current1}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>관리번호</y-th>
      <y-td>{{data.detail.current2}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>관리주체</y-th>
      <y-td>{{data.detail.current3}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>관리책임자</y-th>
      <y-td>{{data.detail.current4}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>관리책임자 연락처</y-th>
      <y-td>{{data.detail.current5}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>층수</y-th>
      <y-td>지하 : {{data.detail.current6}}, 지상 : {{data.detail.current7}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>기초형식</y-th>
      <y-td>{{data.current8s[data.detail.current8]}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>지하깊이</y-th>
      <y-td>{{data.detail.current9}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>기둥표준간격</y-th>
      <y-td>{{data.detail.current11}} × {{data.detail.current12}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>지하수위</y-th>
      <y-td>{{data.detail.current10}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>내진설계여부</y-th>
      <y-td>
        <span v-if="data.detail.current13 == 1">O</span>
        <span v-else>X</span>

        <span v-if="data.detail.current14 == 1">, 구조계산서 미보유</span>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>기준층 슬래브 두께</y-th>
      <y-td>{{data.detail.current15}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>콘크리트 설계강도</y-th>
      <y-td>{{data.detail.current16}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>철근종류</y-th>
      <y-td>{{data.detail.current17}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>철골종류</y-th>
      <y-td>{{data.detail.current18}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>방수공법</y-th>
      <y-td>지붕층 : {{data.detail.current19}}, 지하층 : {{data.detail.current20}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>사용승인도면 보관여부</y-th>
      <y-td>
        건축( <span v-if="data.detail.current22 & 1">√</span> ),
        구조( <span v-if="data.detail.current22 & 2">√</span> ),
        토목( <span v-if="data.detail.current22 & 4">√</span> ),
        조경( <span v-if="data.detail.current22 & 8">√</span> ),
        전기( <span v-if="data.detail.current22 & 16">√</span> ),
        기계( <span v-if="data.detail.current22 & 32">√</span> )
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>사용승인서류 보관여부</y-th>
      <y-td>
        구조계산서( <span v-if="data.detail.current23 & 1">√</span> ),
        지질조사 보고서( <span v-if="data.detail.current23 & 2">√</span> ),
        시방서( <span v-if="data.detail.current23 & 4">√</span> ),
        품질관리계획서( <span v-if="data.detail.current23 & 8">√</span> ),
        내역서( <span v-if="data.detail.current23 & 16">√</span> )
      </y-td>
    </y-tr>
  </y-table>

  <el-dialog
    v-model="data.visible"
    title="시설물 현황 수정"
    :before-close="handleClose"
  >


  <y-table>
    <y-tr>
      <y-th>시설물번호</y-th>
      <y-td><el-input v-model.model="data.item.current1" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>관리번호</y-th>
      <y-td><el-input v-model.model="data.item.current2" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>관리주체</y-th>
      <y-td><el-input v-model.model="data.item.current3" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>관리책임자</y-th>
      <y-td><el-input v-model.model="data.item.current4" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>관리책임자 연락처</y-th>
      <y-td><el-input v-model.model="data.item.current5" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>층수</y-th>
      <y-td>지하 : <el-input v-model.model="data.item.current6" style="width:100px;" />, 지상 : <el-input v-model.model="data.item.current7" style="width:100px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>기초형식</y-th>
      <y-td>
        <el-select v-model.number="data.item.current8" style="width:100%;">
          <el-option v-for="(item, index) in data.current8s" :key="index" :label="item" :value="index" />
        </el-select>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>지하깊이</y-th>
      <y-td><el-input v-model.model="data.item.current9" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>기둥표준간격</y-th>
      <y-td><el-input v-model.model="data.item.current11" style="width:100px;" /> × <el-input v-model.model="data.item.current12" style="width:100px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>지하수위</y-th>
      <y-td><el-input v-model.model="data.item.current10" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>내진설계여부</y-th>
      <y-td>
        <el-radio-group v-model.number="data.item.current13" style="float:left;display:block;">
          <el-radio-button size="small" label="1">O</el-radio-button>
          <el-radio-button size="small" label="2">X</el-radio-button>
        </el-radio-group>

        <el-checkbox style="display:block;float:left;margin-left:10px;margin-top:3px;" size="small" v-model="data.item.check14">구조계산서 미보유</el-checkbox>
        <div style="clear:both;"></div>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>기준층 슬래브 두께</y-th>
      <y-td><el-input v-model.model="data.item.current15" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>콘크리트 설계강도</y-th>
      <y-td><el-input v-model.model="data.item.current16" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>철근종류</y-th>
      <y-td><el-input v-model.model="data.item.current17" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>철골종류</y-th>
      <y-td><el-input v-model.model="data.item.current18" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>방수공법</y-th>
      <y-td>지붕층 : <el-input v-model="data.item.current19" style="width:150px;" />, 지하층 : <el-input v-model="data.item.current20" style="width:150px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>사용승인도면 보관여부</y-th>
      <y-td>
        <el-checkbox size="small" v-model="data.check1">건축</el-checkbox>
        <el-checkbox size="small" v-model="data.check2">구조</el-checkbox>
        <el-checkbox size="small" v-model="data.check4">토목</el-checkbox>
        <el-checkbox size="small" v-model="data.check8">조경</el-checkbox>
        <el-checkbox size="small" v-model="data.check16">전기</el-checkbox>
        <el-checkbox size="small" v-model="data.check32">기계</el-checkbox>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>사용승인서류 보관여부</y-th>
      <y-td>
        <el-checkbox size="small" v-model="data.chk1">구조계산서</el-checkbox>
        <el-checkbox size="small" v-model="data.chk2">지질조사 보고서</el-checkbox>
        <el-checkbox size="small" v-model="data.chk4">시방서</el-checkbox>
        <el-checkbox size="small" v-model="data.chk8">품질관리계획서</el-checkbox>
        <el-checkbox size="small" v-model="data.chk16">내역서</el-checkbox>
      </y-td>
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
import { Apt, Aptdetail } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
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
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  detail: util.clone(item),
  visible: false,
  total: 0,
  check1: false,
  check2: false,
  check4: false,
  check8: false,
  check16: false,
  check32: false,
  chk1: false,
  chk2: false,
  chk4: false,
  chk8: false,
  chk16: false,
  current8s: [' ', '온통기초', '독립기초', '줄기초', '복합기초']
})

function calculateDetail() {
  return util.getInt(data.detail.current3) + util.getInt(data.detail.current4) + util.getInt(data.detail.current5)
}

function calculateTotal() {
  data.total = util.getInt(data.item.current3) + util.getInt(data.item.current4) + util.getInt(data.item.current5)
}

watch(() => data.item.current3, () => {
  calculateTotal()
})

watch(() => data.item.current4, () => {
  calculateTotal()
})

watch(() => data.item.current5, () => {
  calculateTotal()
})


async function initData() {
}

async function getItems() {
  let res = await Aptdetail.get(data.apt)
  const aptdetail = res.item
  if (aptdetail == null) {
    let newItem = util.clone(item)
    newItem.id = data.apt
    await Aptdetail.insert(newItem)
    data.detail = newItem
  } else {
    data.detail = aptdetail
  }

  res = await Apt.get(data.apt)
  const apt = res.item
  data.aptItem = apt
}

async function clickUpdate(pos) {
  let res = await Aptdetail.get(data.apt)
  const item = res.item

  if (item.current14 == 1) {
    item.check14 = true
  } else {
    item.check14 = false
  }

  if (item.current13 == 0) {
    item.current13 = 2
  }

  data.check1 = item.current22 & 1 ? true : false
  data.check2 = item.current22 & 2 ? true : false
  data.check4 = item.current22 & 4 ? true : false
  data.check8 = item.current22 & 8 ? true : false
  data.check16 = item.current22 & 16 ? true : false
  data.check32 = item.current22 & 32 ? true : false

  data.chk1 = item.current23 & 1 ? true : false
  data.chk2 = item.current23 & 2 ? true : false
  data.chk4 = item.current23 & 4 ? true : false
  data.chk8 = item.current23 & 8 ? true : false
  data.chk16 = item.current23 & 16 ? true : false

  data.item = util.clone(item)

  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  if (item.check14 == true) {
    item.current14 = 1
  } else {
    item.current14 = 2
  }

  item.current6 = util.getInt(item.current6)
  item.current7 = util.getInt(item.current7)

  item.current11 = util.getFloat(item.current11)
  item.current12 = util.getFloat(item.current12)
  item.current13 = util.getFloat(item.current13)
  item.current14 = util.getFloat(item.current14)

  var checked = 0

  if (data.check1) checked += 1
  if (data.check2) checked += 2
  if (data.check4) checked += 4
  if (data.check8) checked += 8
  if (data.check16) checked += 16
  if (data.check32) checked += 32

  item.current22 = checked

  checked = 0

  if (data.chk1) checked += 1
  if (data.chk2) checked += 2
  if (data.chk4) checked += 4
  if (data.chk8) checked += 8
  if (data.chk16) checked += 16

  item.current23 = checked

  await Aptdetail.update(item)

  util.info('수정되었습니다')

  data.detail = item
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

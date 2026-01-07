<template>
  <Title title="주소 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <el-input v-model="search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    <TotalDiv :total="data.total" />
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :span-method="spanMethod"  :cell-class-name="cellClassName" :height="height(200)" v-infinite="getItems" @row-click="clickUpdate">
    <el-table-column prop="id" label="ID" width="50" align="center" />
    <el-table-column prop="name" label="아파트명">
      <template #default="scope">
          <div style="color:#000;" >{{scope.row.name}}</div>
      </template>
    </el-table-column>

    <el-table-column prop="tel" label="전화번호" width="100">
      <el-table-column prop="fax" label="팩스번호" width="100">
        <template #default="scope">
          <span v-if="scope.$index % 2 == 0">{{scope.row.tel}}</span>
          <span v-if="scope.$index % 2 != 0">{{scope.row.fax}}</span>
          &nbsp;
        </template>
      </el-table-column>
    </el-table-column>
    <el-table-column prop="email" label="공용메일주소">
      <el-table-column prop="personalemail" label="담당자 연락처">
        <template #default="scope">
          <span class="value" v-if="scope.$index % 2 == 0">{{scope.row.email}}</span>
          <span class="value" v-if="scope.$index % 2 != 0">{{scope.row.personalname}} {{scope.row.personalhp}} {{scope.row.personalemail}}</span>
          &nbsp;
        </template>
      </el-table-column>
      </el-table-column>
    <el-table-column prop="zip" label="우편번호" align="center" width="55" />
    <el-table-column prop="address" label="도로명주소">
      <el-table-column prop="address2" label="지번주소">
        <template #default="scope">
          <span class="value" v-if="scope.$index % 2 == 0">{{scope.row.address}}</span>
          <span class="value" v-if="scope.$index % 2 != 0">{{scope.row.address2}}</span>
          &nbsp;
        </template>
      </el-table-column>
      </el-table-column>
    <el-table-column label="장기수선" width="80" align="center" >
      <template #default="scope">        
        <el-button size="small" v-if="scope.row.contracttype & 1" type="success" @click="clickRepair(scope.row)">장기수선</el-button>
      </template>
    </el-table-column>
    <el-table-column label="안전점검" width="80" align="center" >
      <template #default="scope">
        <el-button size="small" v-if="(scope.row.contracttype & 2) || (scope.row.contracttype & 4) || (scope.row.contracttype & 512)" type="primary" @click="clickPeriodic(scope.row)">안전점검</el-button>
      </template>
    </el-table-column>
    <el-table-column label="순찰" width="80" align="center" >
      <template #default="scope">
        <el-button style="width:64px;" size="small" v-if="scope.row.contracttype & 256" type="warning" @click="clickPatrol(scope.row)">순찰</el-button>
      </template>
    </el-table-column>
  </el-table>  
  <div style="margin-top:10px;display:flex;justify-content: space-between;">
    <div>
      <el-button size="small" type="success" @click="clickInsert">등록</el-button>
      <el-button size="small" type="warning" @click="clickInsertAddress">주소 검색 등록</el-button>
    </div>
    <div>
      <el-button size="small" type="danger" @click="clickDownloadRepair">장기수선 목록 다운로드</el-button>
      <el-button size="small" type="danger" @click="clickDownload">전체 목록 다운로드</el-button>
    </div>
  </div>

  <InquiryInsert :close="clickClose" ref="apt" />

  <el-dialog v-model="data.visibleRepair" width="800px">
    <RepairInsert :id="data.id" />
  </el-dialog>

  <el-dialog v-model="data.visibleDetail" width="800px">
    <DetailInsert :id="data.id" />
  </el-dialog>

  <el-dialog v-model="data.visiblePeriodic" width="800px">
    <PeriodicInsert :id="data.id" />
  </el-dialog>
</template>

<script setup lang="ts">

import { ref, reactive, onMounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Category, Repair } from "~/models"
import { useStore } from 'vuex'

const { width, height } = size()

const store = useStore()

const apt = ref({})

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems(true)
}

const item = {
  id: 0,
  name: '',
  contracttype: 0,
  contractprice: '',
  completeyear: '',
  flatcount: '',
  type: '',
  floor: '',
  familycount: '',
  tel: '',
  fax: '',
  email: '',
  personalemail: '',
  zip: '',
  address: '',
  address2: '',
  testdate: '',
  nexttestdate: '',
  repair: '',
  safety: '',
  fault: '',
  contractdate: '',
  contractduration: '',
  invoice: '',
  depositdate: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  items: [],
  total: 0,
  page: 1,
  pagesize: 30,
  index: 0,
  item: util.clone(item),
  visibleRepair: false,
  visibleDetail: false,
  visiblePeriodic: false,
  check1: false,
  check2: false,
  check3: false,
  check4: false,
  check5: false,
  check6: false,
  check7: false,
  check8: false,
  check9: false
})

async function initData() {
}

async function getItems(reset) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }
  
  let res = await Apt.search({page: data.page, pagesize: data.pagesize, search: search.text})

  if (res.items == null) {
    res.items = []
  }

  let items = []

  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    item.index = ((data.page - 1) * data.pagesize) + i
    items.push(item)
    items.push(item)
  }
  
  data.total = res.total
  data.items = data.items.concat(items)

  data.page++
}

function clickInsert() {
  data.item = util.clone(item)
  
  data.item.name = ''
  data.item.zip = ''
  data.item.address = ''
  data.item.address2 = ''
  
  apt.value?.insert(data.item)
}

function clickUpdate(item, pos, pos2) {
  if (pos.no > 5) {
    return
  }  
  
  data.index = item.index * 2
  data.apt = item.id
  
  apt.value?.readData(item.id)
}

async function clickSubmit() {
  const item = data.item

  let contracttype = 0;
  
  if (data.check1 == true) {
    contracttype += 1;
  }

  if (data.check2 == true) {
    contracttype += 2;
  }

  if (data.check3 == true) {
    contracttype += 4;
  }

  if (data.check4 == true) {
    contracttype += 8;
  }

  if (data.check5 == true) {
    contracttype += 16;
  }

  if (data.check6 == true) {
    contracttype += 32;
  }

  if (data.check7 == true) {
    contracttype += 64;
  }

  if (data.check8 == true) {
    contracttype += 128;
  }

  if (data.check9 == true) {
    contracttype += 256;
  }

  item.contracttype = contracttype;
  
  if (item.name === '') {
    util.error('아파트명을 입력하세요')
    return    
  }

  if (item.contractype === 0) {
    util.error('계약구분을 선택하세요')
    return
  }
  
  let res;

  if (item.id === 0) {
    item.position = item.address
    res = await Apt.insert(item)
  } else {
    res = await Apt.update(item)
  }

  if (res.code === 'ok') {
    util.info('등록되었습니다')

    if (item.id == 0) {
      item.id = res.id
      let items = data.items
      data.items = [item, item, ...items]
    } else {
      data.items[data.index] = item
      data.items[data.index + 1] = item
    }
    
    data.visible = false
  } else {
    util.error('오류가 발생했습니다')
  }
}

onMounted(async () => {
  util.loading(true)

  await initData()
  await getItems()

  util.loading(false)
})

function checkContracttype(item, index) {
  
  
  let check = false
  let d = 1;
  
  for (let i = 1; i <= 8; i++) {
    if (index == i && item & d) {
      check = true
    }

    d = d * 2
  }
  
  return check
}

const spanMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {

  if (columnIndex == 0 || columnIndex == 1 || columnIndex == 4  || columnIndex >= 6) {
    if (rowIndex % 2 == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }
  
  return {rowspan: 1, colspan: 1}
}

function cellClassName({row, columnIndex}) {
  return 'value'    
}

function clickDownload() {
  const url = '/api/download/address'
  const filename = `주소록.xlsx`

  util.download(store, url, filename)
}

function clickDownloadRepair() {
  const url = '/api/download/addressrepair'
  const filename = `장기수선 주소록.xlsx`

  util.download(store, url, filename)
}

function clickZip() {
  new window.daum.Postcode({
    oncomplete: (item) => {
      data.item.name = item.buildingName
      data.item.zip = item.zonecode
      data.item.address = item.roadAddress
      data.item.address2 = item.jibunAddress      
    }
  }).open()  
}

function clickInsertAddress() {
  new window.daum.Postcode({
    oncomplete: (v) => {
      data.item = util.clone(item)
      
      data.item.name = v.buildingName
      data.item.zip = v.zonecode
      data.item.address = v.roadAddress
      data.item.address2 = v.jibunAddress

      apt.value?.insert(data.item)
    }
  }).open()
}

function clickSetting(pos, item) {
  router.push(`/${item.id}/apt/apt`)
}

async function clickRepair(item) {
  data.id = item.id
  data.visibleRepair = true
}

async function clickDetail(item) {
  data.id = item.id
  data.visibleDetail = true
}

async function clickPeriodic(item) {
  data.id = item.id
  data.visiblePeriodic = true  
}

function clickPatrol(item) {
  router.push(`/${item.id}/patrol/patrol`)
}

async function clickClose(item) {
  if (item.id == 0) {
    await getItems(true)
  } else {
    data.items[data.index] = item
    data.items[data.index + 1] = item
  }
}

</script>
<style>
.title {
  background-color: #fafafa;
}

.value {
  background-color: #FFF;
  width: 100%;
  overflow:hidden;
  text-overflow:ellipsis;
  white-space:nowrap;
}

.el-checkbox__label {
  font-size: 12px;
}
</style>  

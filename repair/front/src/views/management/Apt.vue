<template>
  <Title title="작업 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <el-input v-model="search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>

    <TotalDiv :total="data.total" />
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :cell-class-name="cellClassName" :height="height(170)" v-infinite="getItems">
    <el-table-column prop="id" label="ID" width="50" align="center" />
    <el-table-column prop="name" label="아파트명">
      <template #default="scope">
          <div style="color:#000;" @click="clickSetting(scope.$index, scope.row)">{{scope.row.name}}</div>
      </template>
    </el-table-column>

    <el-table-column prop="tel" label="전화번호" width="100" />          
    <el-table-column prop="zip" label="우편번호" align="center" width="55" />
    <el-table-column prop="address" label="도로명주소" />
    <el-table-column prop="address2" label="지번주소" />    
    <el-table-column label="장기수선" width="80" align="center" >
      <template #default="scope">        
        <el-button size="small" v-if="scope.row.contracttype & 1" type="success" @click="clickRepair(scope.row)">장기수선</el-button>
      </template>
    </el-table-column>
    <el-table-column label="정기점검" width="80" align="center" >
      <template #default="scope">
        <el-button size="small" v-if="(scope.row.contracttype & 2) || (scope.row.contracttype & 4)" type="primary" @click="clickPeriodic(scope.row)">안전점검</el-button>
      </template>
    </el-table-column>
    <el-table-column label="순찰" width="80" align="center" >
      <template #default="scope">
        <el-button style="width:64px;" size="small" v-if="scope.row.contracttype & 256" type="warning" @click="clickPatrol(scope.row)">순찰</el-button>
      </template>
    </el-table-column>
  </el-table>
  <!-- 
  <div style="margin-top:10px;display:flex;justify-content: space-between;">
    <div>
      <el-button size="small" type="success" @click="clickInsert">등록</el-button>
      <el-button size="small" type="warning" @click="clickInsertAddress">주소 검색 등록</el-button>
    </div>
    <el-button size="small" type="danger" @click="clickDownload">전체 목록 다운로드</el-button>
  </div>
  -->
  
  <el-dialog
    v-model="data.visible"
    title="주소 등록/수정"
    width="1100px"
    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="100px">

      <div>
        <div style="float:left;text-align:right;">
          <el-form-item label="ID" v-show="data.item.id != 0">
            {{ data.item.id }}
          </el-form-item>
        </div>
        <div style="float:right;text-align:right;">
          <el-button size="small" @click="data.visible = false">취소</el-button>
          <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
        </div>
        <div style="clear:both;"></div>
      </div>
      
      
      <el-form-item label="아파트명" style="margin:0px 0px;">
        <el-input v-model="data.item.name" />
      </el-form-item>

      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 33%;">

          <el-form-item label="준공년도" style="margin:0px 0px;">
            <el-input v-model="data.item.completeyear" />
          </el-form-item>

          <el-form-item label="건축물형태" style="margin:0px 0px;">
            <el-input v-model="data.item.type" />
          </el-form-item>
          

          <el-form-item label="전체동수" style="margin:0px 0px;">
            <el-input v-model="data.item.flatcount" />
          </el-form-item>

          <el-form-item label="세대수" style="margin:0px 0px;">
            <el-input v-model="data.item.familycount" />
          </el-form-item>

          <el-form-item label="층수" style="margin:0px 0px;">
            <el-input v-model="data.item.floor" />
          </el-form-item>


          

          <el-form-item label="전화번호" style="margin:0px 0px;">
            <el-input v-model="data.item.tel" />
          </el-form-item>

          <el-form-item label="팩스번호" style="margin:0px 0px;">
            <el-input v-model="data.item.fax" />
          </el-form-item>
          

      
          <el-form-item label="공용메일주소" style="margin:0px 0px;">
            <el-input v-model="data.item.email" />
          </el-form-item>
          
          
    
          <el-form-item label="개인메일주소" style="margin:3px 0px 0px 0px;">
            <el-input v-model="data.item.personalemail" :rows="2" type="textarea" style="font-size:12px;" />
          </el-form-item>
          
          
         
        </div>
        <div style="flex:1 1 67%;">

          
          
          <el-form-item label="장기수선계획" style="margin:5px 0px 0px 0px;">
            <el-input v-model="data.item.repair" :rows="5" type="textarea" style="font-size:12px;" />
          </el-form-item>
          <el-form-item label="안전점검" style="margin:5px 0px;">
            <el-input v-model="data.item.safety" :rows="5" type="textarea" style="font-size:12px;" />
          </el-form-item>
          <el-form-item label="기타" style="margin:0px 0px;">
            <el-input v-model="data.item.fault" :rows="4" type="textarea" style="font-size:12px;" />
          </el-form-item>      
          
          
          
          
        </div>
      </div>

      
      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 16%;">

          <el-form-item label="우편번호" style="margin:0px 0px;">
            <el-input v-model="data.item.zip" @click="clickZip" />
          </el-form-item>
        </div>

        
        <div style="flex:1 1 42%;">


          <el-form-item label="도로명주소" style="margin:0px 0px;">
            <el-input v-model="data.item.address" />
          </el-form-item>
          
        </div>

        <div style="flex:1 1 42%;">

          <el-form-item label="지번주소" style="margin:0px 0px;">
            <el-input v-model="data.item.address2" />
          </el-form-item>
          
        </div>
      </div>
      
      

      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 33%;">
          <el-form-item label="정밀점검일자" style="margin:0px 0px;">
            <el-input v-model="data.item.testdate" />
          </el-form-item>

          <!--
          <el-form-item label="정밀안전점검 차기일자">
            <el-input v-model="data.item.nexttestdate" />
          </el-form-item>
          -->
          
        </div>
        <div style="flex:1 1 33%;">
          <el-form-item label="FMS 아이디" style="margin:0px 0px;">
            <el-input v-model="data.item.fmsloginid" />
          </el-form-item>
          
        </div>
        <div style="flex:1 1 33%;">
          <el-form-item label="FMS 비번" style="margin:0px 0px;">
            <el-input v-model="data.item.fmspasswd" />
          </el-form-item>
          
        </div>
        
        


        
      </div>

      
      <el-form-item label="계약구분" style="margin:0px 0px;font-size:12px;">        

        <el-checkbox size="small" label="장기수선계획" v-model="data.check1" style="font-size:12px;" />
        <el-checkbox size="small" label="정밀" v-model="data.check2" style="font-size:12px;" />
        <el-checkbox size="small" label="정기" v-model="data.check3" style="font-size:12px;" />
        <el-checkbox size="small" label="하자보수" v-model="data.check4" style="font-size:12px;" />
        <el-checkbox size="small" label="하자조사" v-model="data.check5" style="font-size:12px;" />
        <el-checkbox size="small" label="정밀안전진단" v-model="data.check6" style="font-size:12px;" />
        <el-checkbox size="small" label="감리" v-model="data.check7" style="font-size:12px;" />
        <el-checkbox size="small" label="기술자문" v-model="data.check8" style="font-size:12px;" />
        <el-checkbox size="small" label="순찰" v-model="data.check9" style="font-size:12px;" />
      </el-form-item>

      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 18%;">
          
          <el-form-item label="계약날짜" style="margin:0px 0px;">
            <el-input v-model="data.item.contractdate" />
          </el-form-item>

        </div>
        <div style="flex:1 1 28%;">
          
          <el-form-item label="계약기간" style="margin:0px 0px;">
            <el-input v-model="data.item.contractduration" />
          </el-form-item>

        </div>

        <div style="flex:1 1 18%;">
          
          <el-form-item label="계약금액" style="margin:0px 0px;">
            <el-input v-model="data.item.contractprice" />
          </el-form-item>

        </div>
        <div style="flex:1 1 18%;">          
          
          <el-form-item label="계산서 발행" style="margin:0px 0px;">
            <el-input v-model="data.item.invoice" />            
          </el-form-item>

        </div>
        <div style="flex:1 1 18%;">

          <el-form-item label="입금날짜" style="margin:0px 0px;">
            <el-input v-model="data.item.depositdate" />
          </el-form-item>

        </div>
      </div>
          
        
          
    </el-form>

    <template #footer>
      <el-button size="small" style="float:left;" type="danger" @click="clickDelete(data.item)">삭제</el-button>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>

  <el-dialog v-model="data.visibleRepair" width="800px" :before-close="handleClose">
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

import { reactive, onMounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Category, Repair } from "~/models"
import { useStore } from 'vuex'

const { width, height } = size()

const store = useStore()

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
  id: 0,
  items: [],
  total: 0,
  page: 1,
  pagesize: 30,
  index: 0,
  item: util.clone(item),
  visible: false,
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
  
  let res = await Apt.search({page: data.page, pagesize: data.pagesize, search: search.text, contracttype: 263})

  if (res.items == null) {
    res.items = []
  }

  let items = []

  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]
    items.push(item)
  }
  
  data.total = res.total
  data.items = data.items.concat(items)

  data.page++
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true
}

function clickUpdate(pos, item) {  
  data.index = pos
  
  data.item = util.clone(item)

  if (data.item.contracttype & 1) {
    data.check1 = true;
  } else {
    data.check1 = false;
  }

  if (data.item.contracttype & 2) {
    data.check2 = true;
  } else {
    data.check2 = false;    
  }

  if (data.item.contracttype & 4) {
    data.check3 = true;
  } else {
    data.check3 = false;    
  }

  if (data.item.contracttype & 8) {
    data.check4 = true;
  } else {
    data.check4 = false;    
  }

  if (data.item.contracttype & 16) {
    data.check5 = true;
  } else {
    data.check5 = false;    
  }

  if (data.item.contracttype & 32) {
    data.check6 = true;
  } else {
    data.check6 = false;    
  }

  if (data.item.contracttype & 64) {
    data.check7 = true;
  } else {
    data.check7 = false;    
  }

  if (data.item.contracttype & 128) {
    data.check8 = true;
  } else {
    data.check8 = false;    
  }

  if (data.item.contracttype & 256) {
    data.check9 = true;
  } else {
    data.check9 = false;
  }
  
  data.visible = true
}

function clickDelete(item) {
  util.confirm('삭제하시겠습니까', async function() {
    util.confirm('한번 삭제한 주소는 복구가 불가능합니다. 삭제하시겠습니까', async function() {
      let res = await Apt.remove(item)
      if (res.code === 'ok') {
        util.info('삭제되었습니다')
        let items = data.items
        items.splice(data.index, 2)
        data.items = items

        data.visible = false
      }
    })
  })
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

      data.visible = true
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

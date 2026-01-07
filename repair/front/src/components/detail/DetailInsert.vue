<template>
  <Title :title="data.title" subtitle="정밀점검" />


  <el-tabs v-model="data.menu">
    <el-tab-pane label="작업 목록" name="list">

      <el-table :data="data.items" border :height="height(566)" @row-click="clickDetail" ref="listRef">
        <el-table-column prop="index" label="NO" align="center" width="60" />
        <el-table-column prop="name" label="작업명" />        
        <!--<el-table-column prop="startdate" label="시작일" align="center" />
        <el-table-column prop="enddate" label="종료일" align="center" />-->
        <el-table-column prop="reportdate" label="리포트작업일" align="center" />
        <el-table-column label="상태" align="center" width="60">
          <template #default="scope">
            <el-tag :type="Detail.getStatusType(scope.row.status)">{{Detail.getStatus(scope.row.status)}}</el-tag>                        
          </template>
        </el-table-column>
        <el-table-column prop="date" label="등록일" align="center" />        
        <el-table-column label="" align="center" width="100">
          <template #default="scope">
            <el-button size="small" type="success" style="margin-right:0px;" @click="clickDetail(scope.row)">정밀점검</el-button>
          </template>
        </el-table-column>
        
      </el-table>
      
    </el-tab-pane>
    <el-tab-pane label="관리" name="management">


      
      <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
        <div style="flex:1;text-align:right;gap:5;">
          <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;float:left;">삭제</el-button>        
          <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">작업 생성</el-button>
        </div>
      </div>  


      <el-table :data="data.items" border :width="data.width" :height="height(600)" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
        <el-table-column type="selection" width="40" align="center" />
        <el-table-column prop="index" label="NO" align="center" width="60" />    
        <el-table-column prop="name" label="작업명" />        
        <!--<el-table-column prop="startdate" label="시작일" align="center" />
        <el-table-column prop="enddate" label="종료일" align="center" />-->
        <el-table-column prop="reportdate" label="리포트작업일" align="center" />
        <el-table-column label="상태" align="center" width="60">
          <template #default="scope">
            <el-tag :type="Detail.getStatusType(scope.row.status)">{{Detail.getStatus(scope.row.status)}}</el-tag>                        
          </template>
        </el-table-column>
        <el-table-column prop="date" label="등록일" align="center" />        
        <el-table-column label="" align="center" width="100">
          <template #default="scope">
            <el-button size="small" type="warning" style="margin-right:0px;" @click="clickCopy(scope.row)">작업복제</el-button>
          </template>
        </el-table-column>
      </el-table>  

    </el-tab-pane>
  </el-tabs>

  <el-dialog v-model="data.visible" width="800px">
    <el-form label-width="80px">
      <el-form-item label="ID" v-if="data.item.id > 0">
        {{data.item.id}}
      </el-form-item>
      
      <el-form-item label="작업명">
        <el-input v-model="data.item.name" placeholder="" />
      </el-form-item>
      
      <el-form-item label="작업일" v-show="false">
        <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.startdate" placeholder="" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.enddate" placeholder="" /> 
      </el-form-item>

      <el-form-item label="리포트작업일">
        <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.reportdate" placeholder="" />
      </el-form-item>

      <el-form-item label="상태">
        <el-radio-group v-model.number="data.item.status">
          <el-radio-button size="small" label="1">준비</el-radio-button>
          <el-radio-button size="small" label="2">착수</el-radio-button>
          <el-radio-button size="small" label="3">완료</el-radio-button>
          <el-radio-button size="small" label="4">중단</el-radio-button>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="참여기술진" v-show="false">
        <div style="text-align:right;margin-bottom:5px;width:100%;">
          <el-button size="small" type="primary" style="margin-right:0px;" @click="clickTechnicianInsert">참여기술진 추가</el-button>
        </div>
        <el-table :data="data.detailtechnicians" border height="300">
          <el-table-column prop="type" label="참여구분" align="center" width="70">
            <template #default="scope">
              {{Detailtechnician.getType(scope.row.type)}}
            </template>
          </el-table-column>
          <el-table-column prop="extra.technician.original" label="성명" width="70" />    
          <el-table-column prop="part" label="참여분야" />        
          <el-table-column label="기술등급" align="center" width="90">
            <template #default="scope">
              {{Technician.getGrade(scope.row.extra.technician.grade)}}
            </template>
          </el-table-column>
          <el-table-column prop="remark" label="비고" />                
          <el-table-column label="" align="center" width="150">
            <template #default="scope">
              <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickTechnicianUp(scope.row, scope.$index)"><el-icon><ArrowUp /></el-icon></el-button>
              <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickTechnicianDown(scope.row, scope.$index)"><el-icon><ArrowDown /></el-icon></el-button>
              <el-button size="small" type="info" style="width:30px;margin-right:-7px;" @click="clickTechnicianUpdate(scope.row, scope.$index)">수정</el-button>
              <el-button size="small" type="danger" style="width:30px;margin-right:0px;" @click="clickTechnicianDelete(scope.row, scope.$index)">삭제</el-button>
            </template>
          </el-table-column>
        </el-table>
        
      </el-form-item>
            
    </el-form>

    
    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>


  <el-dialog v-model="data.visibleInsert" width="500px">
    <el-form label-width="70px">
      <el-form-item label="참여 구분">
        <el-radio-group v-model.number="data.detailtechnician.type">
          <el-radio-button size="small" label="1">책임기술자</el-radio-button>
          <el-radio-button size="small" label="2">참여기술자</el-radio-button>          
        </el-radio-group>
      </el-form-item>

      <el-form-item label="기술자">
        <el-select v-model.number="data.detailtechnician.technician" style="width:100%;">
          <el-option v-for="item in data.alltechnicians" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      
      <el-form-item label="참여분야">
        <el-input v-model="data.detailtechnician.part" placeholder="" />
      </el-form-item>

      <el-form-item label="비고">
        <el-input v-model="data.detailtechnician.remark" placeholder="" />
      </el-form-item>
                        
    </el-form>

    
    <template #footer>
      <el-button size="small" @click="data.visibleInsert = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickInsertSubmit">등록</el-button>
    </template>
  </el-dialog>  

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted, computed, watch, watchEffect } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Detail, Technician, Detailtechnician } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const { width, height } = size()

const props = defineProps({
  id: Number,
  close: Function
})

defineExpose({
  reset
})

const store = useStore()
const route = useRoute()

const search = reactive({
  text: ''
})

const item = {
  id: 0,
  name: '',  
  startdate: null,
  enddate: null,
  reportdate: null,
  status: 1  
}

const detailtechnician = {
  id: 0,
  type: 0,
  part: '',
  remark: '',
  order: 0,
  user: 0,
  detail: 0,
  extra: {
    user: {
      name: ''
    }
  }
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  visibleInsert: false,
  search: '',
  title: '',
  menu: 'list',
  level: 0,
  alltechnicians: [],
  detailtechnicians: [],
  detailtechnician: util.clone(detailtechnician),
  detailtechnicianIndex: -1
})

watch(() => props.id, () => {
  data.menu = 'list'
  data.apt = props.id
})

watchEffect(() => {
  data.apt = props.id
  getItems()
})

async function initData() {
  let res = await Technician.find({orderby: 'te_name'})
  if (res.items == undefined) {
    res.items = []
  }

  let items = [{id: 0, name: ' '}]

  res.items.forEach(item => {
    items.push({id: item.id, name: `${item.name} - ${Technician.getGrade(item.grade)}`, original: item.name, grade: item.grade});
  })
  data.alltechnicians = items
}

async function getItems() {
  if (data.apt == 0) {
    return
  }

  let res = await Apt.get(data.apt)
  data.title = res.item.name

  res = await Detail.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'd_id'
  })

  if (res.items == null) {
    res.items = []
  }

  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1
  }

  data.total = res.total
  data.items = res.items
}

function makeItems(items) {
  return items
}

function clickInsert() {
  data.item = util.clone(item)
  data.visible = true
}

async function clickUpdate(item, index) {
  if (index.no == 0 || index.no == 6) {
    return
  }

  if (data.menu == 'management') {
    data.item = util.clone(item)

    let res = await Detailtechnician.find({detail: item.id, orderby: 'dt_order'})
    if (res.items == undefined) {
      res.items = []
    }

    data.detailtechnicians = res.items
    data.visible = true
    return
  }  
}

function clickDelete() {
  let item = data.batchs[0]  
  util.confirm('삭제하시겠습니까', async function() {
    util.confirm('삭제한 데이터는 복구가 불가능합니다. 삭제하시겠습니까', async function() {
      let res = await Detail.remove(item)
      if (res.code === 'ok') {
        util.info('삭제되었습니다')
        getItems()
        
      }
    })
  })
}

const handleClose = (done: () => void) => {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      done()
    })
  } else {
    done()
  }

  data.menu = 'list'
}

onMounted(async () => {
  if (store.getters["getUser"] != null) {
    data.level = store.getters["getUser"].level
  }
  
  util.loading(true)
  
  await initData()
  await getItems()

  util.loading(false)
})

const listRef = ref<InstanceType<typeof ElTable>>()
const listSelection = ref([])
const toggleListSelection = (rows) => {
  if (rows) {
    rows.forEach((row) => {
      listRef.value!.toggleRowSelection(row, undefined)
    })
  } else {
    listRef.value!.clearSelection()
  }
}
const changeList = (val) => {
  listSelection.value = val
}

const listTechnicianRef = ref<InstanceType<typeof ElTable>>()
const listTechnicianSelection = ref([])
const toggleListTechnicianSelection = (rows) => {
  if (rows) {
    rows.forEach((row) => {
      listTechnicianRef.value!.toggleRowSelection(row, undefined)
    })
  } else {
    listTechnicianRef.value!.clearSelection()
  }
}
const changeTechnicianList = (val) => {
  listTechnicianSelection.value = val
}

function clickDeleteMulti() {
  if (listSelection.value.length == 0) {
    util.error('선택된 항목이 없습니다')
    return
  }
  
  util.confirm('삭제하시겠습니까', async function() {
    util.confirm('삭제한 데이터는 복구가 불가능합니다. 삭제하시겠습니까', async function() {
      let items = []
      for (let i = 0; i < listSelection.value.length; i++) {
        let value = listSelection.value[i]

        let item = {
          id: value.id
        }

        items.push(item)      
      }

      await Detail.removebatch(items)
      util.info('삭제되었습니다')
      getItems()
    })
  })
}


function clickRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function clickAdd(count) {
  let items = []
  for (let i = 0; i < count; i++) {
    items.push(util.clone(item))
  }

  data.batchs = data.batchs.concat(items)
}

async function clickSubmit() {
  let item = util.clone(data.item)
  if (util.isNull(item.name) || item.name == '') {
    util.error('작업명을 입력하세요')
    return
  }

  item.startdate = util.convertDBDate(item.startdate)
  item.enddate = util.convertDBDate(item.enddate)
  item.reportdate = util.convertDBDate(item.reportdate)

  item.prestartdate = util.convertDBDate()
  item.preenddate = util.convertDBDate()
  item.researchstartdate = util.convertDBDate()
  item.researchenddate = util.convertDBDate()
  item.analyzestartdate = util.convertDBDate()
  item.analyzeenddate = util.convertDBDate()
  item.ratingstartdate = util.convertDBDate()
  item.ratingenddate = util.convertDBDate()
  item.writestartdate = util.convertDBDate()
  item.writeenddate = util.convertDBDate()
  item.printstartdate = util.convertDBDate()
  item.printenddate = util.convertDBDate() 
  
  item.apt = data.apt

  if (item.id > 0) {
    await Detail.update(item)
    //await Detailtechnician.deleteByDetail({detail: item.id})
  } else {  
    let res = await Detail.insert(item)
    item.id = res.id
  }

  /*
  for (let i = 0; i < data.detailtechnicians.length; i++) {
    let detailtechnicians = util.clone(data.detailtechnicians[i])
    detailtechnicians.id = 0
    detailtechnicians.detail = item.id
    detailtechnicians.order = i + 1
    detailtechnicians.date = ''
    await Detailtechnician.insert(detailtechnicians)
  nn}
  */

  util.info('등록되었습니다')
  
  await getItems()
  data.visible = false  
}

async function clickDetail(item) {
  let res = await Detail.get(item.id)
  let detail = res.item
  
  if (!util.isNull(props.close)) {
    props.close()
  }

  router.push(`/${detail.apt}/detail/${detail.id}/detail`)
}

function clickChangeStatus(status) {
  if (listSelection.value.length == 0) {
    util.error('선택된 항목이 없습니다')
    return
  }

  let title = ''

  if (status == 1) {
    title = '마감 해제'
  } else {
    title = '마감 처리'
  }
  util.confirm(title + ' 하시겠습니까', async function() {
    let items = []
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      await Detail.updateStatusById(status, value.id)
    }

    util.info(title + ' 되었습니다')
    getItems()
  })  
}

function reset() {
  data.menu = 'list'
}

function clickTechnicianInsert() {
  data.detailtechnician = util.clone(detailtechnician)
  data.detailtechnicianIndex = -1
  data.visibleInsert = true
}

function clickTechnicianUpdate(row, index) {
  data.detailtechnician = util.clone(row)
  data.detailtechnicianIndex = index
  data.visibleInsert = true  
}

function clickTechnicianDelete(row, index) {
  let items = util.clone(data.detailtechnicians)
  items.splice(index, 1)
  data.detailtechnicians = items
}

function clickTechnicianUp(row, index) {
  if (index == 0) {
    return
  }

  let items = util.clone(data.detailtechnicians)
  let temp = items[index]
  items[index] = items[index - 1]
  items[index - 1] = temp

  data.detailtechnicians = items
}

function clickTechnicianDown(row, index) {
  if (index >= data.detailtechnicians.length - 1) {
    return
  }

  let items = util.clone(data.detailtechnicians)
  let temp = items[index]
  items[index] = items[index + 1]
  items[index + 1] = temp

  data.detailtechnicians = items
}

function clickInsertSubmit() {
  let item = util.clone(data.detailtechnician)

  if (item.type == 0) {
    util.error('참여 구분을 선택하세요')
    return
  }

  if (item.technician == 0) {
    util.error('기술자를 선택하세요')
    return
  }

  for (let i = 0; i < data.alltechnicians.length; i++) {
    if (data.alltechnicians[i].id == item.technician) {
      item.extra = {
        technician: util.clone(data.alltechnicians[i])
      }
    }
  }

  if (data.detailtechnicianIndex == -1) {
    data.detailtechnicians = [...data.detailtechnicians, item]
  } else {
    let items = util.clone(data.detailtechnicians)
    items[data.detailtechnicianIndex] = item
    data.detailtechnicians = items
  }

  data.visibleInsert = false
}

function clickCopy(item) {
  util.confirm('정밀점검을 생성하시겠습니까<BR/>생성시 약 약간의 시간이 소요됩니다.<BR/>브라우져 창을 닫지 말고 기다려주세요', async function() {
    util.loading(true)
    
    await Detail.duplication(item.id)
    await getItems()

    util.loading(false)
  })
  
}

</script>

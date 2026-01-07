<template>
  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" >등록</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :height="height(324)" @row-click="clickUpdate"  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />    
    <el-table-column prop="type" label="점검종류" align="center" width="100" />
    <el-table-column prop="company" label="시행자" align="center" width="150" />
    <el-table-column prop="name" label="점검자" align="center" width="100" />
    <el-table-column label="점검기간" align="center" width="250">
      <template #default="scope">
        {{util.viewDate(scope.row.repairstartdate)}} ~ {{util.viewDate(scope.row.repairenddate)}}
      </template>
    </el-table-column>
    <el-table-column prop="content" label="점검 결과 및 주요 보수보강 요약" align="center" />    
    <el-table-column prop="grade" label="안전등급" align="center" width="80" />
    <!--
    <el-table-column label="" align="center" width="105">
      <template #default="scope">
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickUp(scope.row, scope.$index)"><el-icon><ArrowUp /></el-icon></el-button>
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickDown(scope.row, scope.$index)"><el-icon><ArrowDown /></el-icon></el-button>        
      </template>
    </el-table-column>
    -->
  </el-table>  

  <div class="resulttext" style="height:80px;overflow:auto;">
    <div style="margin-top:4px;">{{data.periodic.past}}</div>
    <el-button size="small" type="success" @click="clickUpdatePast">수정</el-button>
  </div>

  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1050px"
  >

    <y-table>
      <y-tr>
        <y-th>점검종류</y-th>
        <y-td><el-input v-model="data.item.type" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>시행자</y-th>
        <y-td><el-input v-model="data.item.company" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>점검자</y-th>
        <y-td><el-input v-model="data.item.name" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>점검기간</y-th>
        <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.repairstartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.repairenddate" /> </y-td>
      </y-tr>
      <y-tr>
        <y-th>점검결과</y-th>
        <y-td><el-input v-model="data.item.content" :rows="10" type="textarea" style="font-size:12px;" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>안전등급</y-th>
        <y-td><el-input v-model="data.item.grade" /></y-td>
      </y-tr>      
    </y-table>

      <template #footer>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

  <el-dialog
    v-model="data.visiblePast"
  >

    <y-table>
      <y-tr>
        <y-td><el-input v-model="data.past" :rows="10" type="textarea" style="font-size:12px;" /></y-td>
      </y-tr>
    </y-table>

      <template #footer>
        <el-button size="small" @click="data.visiblePast = false">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmitPast">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Periodic, Periodicpast } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodicpast

const item = {
  id: 0,
  type: '',
  company: '',
  name: '',
  repairstartdate: null,
  repairenddate: null,
  content: '',
  grade: '',
  ordre: 0,
  periodic: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  visiblePast: false,
  periodic: { past: '' },
  past: ''
})

async function initData() {
  let res = await Periodic.get(data.id)

  if (res.item.past == '') {
    res.item.past = '가. 금번 정기안전점검에서 상기와 같은 결과를 검토·분석하여 기 발생된 결함사항에 대해 구조적 문제의 발생 및 결함의 진행, 보수·보강 여부 등을 중점으로 검토하였다.'
  }

  data.periodic = res.item
}

async function getItems() {
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    periodic: data.id,
    orderby: 'pp_repairstartdate desc,pp_id desc'    
  })


  data.total = res.total
  data.items = res.items
}

function clickInsert() {  
  data.item = util.clone(item)
  
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0 || index.no == 6) {
    return
  }

  if (item.repairstartdate == '1000-01-01') {
    item.repairstartdate = null
  }

  if (item.repairenddate == '1000-01-01') {
    item.repairenddate = null
  }
  
  data.item = util.clone(item)
  data.mode = 'normal'
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await model.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      await getItems()
    }
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
}

onMounted(async () => {
  data.id = parseInt(route.params.id)
  data.apt = parseInt(route.params.apt)  
  
  util.loading(true)
  
  await initData()
  await getItems()

  data.visible = false
  util.loading(false)
})

function clickCancel() {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      data.visible = false
    })
  } else {
    data.visible = false
  }
}

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

function clickDeleteMulti() {
  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)
    
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id
      }

      await model.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  util.loading(true)

  let item = util.clone(data.item)  
  
  item.periodic = data.id
  item.repairstartdate = util.convertDBDate(data.item.repairstartdate)
  item.repairenddate = util.convertDBDate(data.item.repairenddate)

  if (item.id == 0) {
    let max = 0
    for (let i = 0; i < data.items.length; i++) {
      if (data.items[i].order > max) {
        max = data.items[i].order
      }
    }

    item.order = max + 1
  }

  if (item.id > 0) {
    await model.update(item)
  } else { 
    await model.insert(item)
  }

  util.info('등록되었습니다')

  await getItems()
  data.visible = false  
  util.loading(false)  
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

async function clickUp(row, index) {
  if (index == 0) {
    return
  }

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index - 1].order
  items[index - 1].order = temp

  await model.update(items[index])
  await model.update(items[index - 1])

  await getItems()
}

async function clickDown(row, index) {
  if (index >= data.items.length - 1) {
    return
  }

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index + 1].order
  items[index + 1].order = temp

  await model.update(items[index])
  await model.update(items[index + 1])

  await getItems()
}

function clickUpdatePast() {
  data.past = data.periodic.past

  data.visiblePast = true
}

async function clickSubmitPast() {
  util.loading(true)

  let res = await Periodic.get(data.id)
  let item = res.item
  item.past = data.past

  await Periodic.update(item)

  data.periodic = item

  util.loading(false)
  data.visiblePast = false
}
</script>

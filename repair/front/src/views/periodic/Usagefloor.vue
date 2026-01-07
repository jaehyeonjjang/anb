<template>
  <Title title="용도현황 (층구분)" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄처리</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :height="height(170)" @row-click="clickUpdate"  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />    
    <el-table-column prop="floor" label="층구분" align="center" />
    <el-table-column prop="purpose" label="주용도" align="center" />
    <el-table-column prop="area" label="면적" align="center" />        
    <el-table-column prop="remark" label="비고" align="center" />
    <el-table-column label="" align="center" width="105">
      <template #default="scope">
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickUp(scope.row, scope.$index)"><el-icon><ArrowUp /></el-icon></el-button>
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickDown(scope.row, scope.$index)"><el-icon><ArrowDown /></el-icon></el-button>        
      </template>
    </el-table-column>    
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="800px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="층구분" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].floor" />
          </template>
        </el-table-column>
        <el-table-column label="주용도" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].purpose" />
          </template>
        </el-table-column>
        <el-table-column label="바닥면적" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].area" />
          </template>
        </el-table-column>
        <el-table-column label="비고" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].remark" />
          </template>
        </el-table-column>

      </el-table>


    </el-form>

      <template #footer>
        <el-button size="small" type="danger" v-if="data.mode != 'batch' && (data.batchs.length > 0 && data.batchs[0].id > 0)" style="float:left;" @click="clickDelete">삭제</el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(1)"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(10)"><el-icon><Plus /></el-icon> &nbsp;10</el-button>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Aptusagefloor } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  floor: '',
  purpose: '',
  area: '',
  remark: '',
  order: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false
})

async function initData() {  
}

async function getItems() {
  let res = await Aptusagefloor.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'af_order,af_id'    
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }
  }

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }

  data.items = res.items
}

function clickInsert() {  
  data.item = util.clone(item)

  let items = [data.item]
  
  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0 || index.no == 5) {
    return
  }

  let items = [util.clone(item)]

  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Aptusagefloor.remove(item)
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
  data.apt = parseInt(route.params.apt)  
  
  util.loading(true)
  
  await initData()
  await getItems()

  data.visible = false
  util.loading(false)
})

function clickBatch() {
  let items = util.clone(data.items)

  if (items == null) {
    items = []
  }

  if (items.length == 0) {
    for (let i = 0; i < 5; i++) {
      items.push(util.clone(data.item))
    }
  }

  data.mode = 'batch'
  data.batchs = items
  data.visible = true  
}

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

      await Aptusagefloor.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  util.loading(true)

  if (data.mode == 'batch') {
    for (let i = 0; i < data.items.length; i++) {
      let item = data.items[i]
      let flag = false;
      for (let j = 0; j < data.batchs.length; j++) {
        if (data.items[i].id == data.batchs[j].id) {
          flag = true
          break
        }
      }

      if (flag == false) {      
        await Aptusagefloor.remove(item)
      }
    }
  } else {
    let max = 0
    console.log(data.items)
    for (let i = 0; i < data.items.length; i++) {
      let item = data.items[i]

      if (item.order > max) {
        max = item.order
      }
    }

    console.log(max)

    max++

    data.batchs[0].order = max
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    if (item.floor == '' && item.purpose == '' && item.area == '' && item.remark == '') {
      continue
    }
    
    item.apt = data.apt
    item.floor = String(item.floor)
    item.purpose = String(item.purpose)    
    item.area = String(item.area)
    item.remark = String(item.remark)
    if (data.mode == 'batch') {
      item.order = i + 1
    }

    if (item.id > 0) {
        await Aptusagefloor.update(item)
    } else { 
      await Aptusagefloor.insert(item)
    }
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

  await Aptusagefloor.update(items[index])
  await Aptusagefloor.update(items[index - 1])

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

  await Aptusagefloor.update(items[index])
  await Aptusagefloor.update(items[index + 1])

  await getItems()
}

</script>

<template>
  <Title title="항목검토" />


  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <el-input v-model="data.search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />
    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>

    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">등록</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="id" label="NO" align="center" width="60" />
    <el-table-column prop="name" label="검토사유" >
      <template #default="scope">
        <span v-html="util.nl2br(scope.row.content)" />
      </template>
    </el-table-column>

    <el-table-column prop="name" label="조정내용" >
      <template #default="scope">
        <span v-html="util.nl2br(scope.row.adjust)" />
      </template>
    </el-table-column>

    <el-table-column prop="order" label="순번" align="center" width="60" />

  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1000px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="검토사유" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content" />
          </template>
        </el-table-column>

        <el-table-column label="조정내용" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].adjust" />
          </template>
        </el-table-column>

        <el-table-column label="순번" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].order" />
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
import { util }  from "~/global"
import { Reviewbasic } from "~/models"
import { useStore } from 'vuex'
import { ElTable } from 'element-plus'

const store = useStore()

const search = reactive({
  text: ''
})

function clickSearch() {
    getItems()
}

function paging(page) {
  data.page = page
  getItems()
}

const item = {
  id: 0,
  content: '',
  order: 1000
}

const data = reactive({
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,    
  search: {
    text: ''
  }
})

async function initData() {  
}

async function getItems() {
  let res = await Reviewbasic.find({
    page: data.page,
    pagesize: data.pagesize,
    content: data.search.text,
    orderby: 'rv_order,rv_id'
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

function makeItems(items) {
  return items
}
function clickInsert() {  
  data.item = util.clone(item)

  let items = makeItems([data.item])

  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0) {
    return
  }

  let items = makeItems([util.clone(item)])

  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Reviewbasic.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      getItems()
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

function setWindowSize() {
  data.width = (window.innerWidth - 500) + 'px'
  data.height = (window.innerHeight - 170) + 'px'
}

onMounted(() => {
  util.loading(true)
  initData()
  getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  data.visible = false
  util.loading(false)
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
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

  items = makeItems(items)
  
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

      await Reviewbasic.remove(item)
    }

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    
    if (item.content == '') {
      util.error('검토사유를 입력하세요')
      return    
    }

    if (item.adjust == '') {
      util.error('조정내용을 입력하세요')
      return    
    }
  }
  
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
        await Reviewbasic.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.order = util.getInt(item.order)

    if (item.id > 0) {
      await Reviewbasic.update(item)
    } else { 
      await Reviewbasic.insert(item)
    }
  }

  util.info('등록되었습니다')
  getItems()
  data.visible = false

  setTimeout(function() {
    listRef.value!.setScrollTop(data.items.length  * 100 + 1000)
  }, 500)
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

</script>

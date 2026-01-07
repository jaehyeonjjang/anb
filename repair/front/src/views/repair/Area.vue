<template>
  <Title title="공동주택" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄처리</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" :summary-method="getSummaries" show-summary  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" sortable />
    <el-table-column prop="name" label="평형" align="center" sortable />
    <el-table-column prop="familycount" label="세대수" align="right" sortable />
    <el-table-column prop="size" label="면적" align="right" sortable />    
    <el-table-column label="총면적" align="right">
      <template #default="scope">
        {{util.area(scope.row.familycount*scope.row.size)}}
      </template>
    </el-table-column>

    <el-table-column prop="order" label="순번" align="center" />
    <el-table-column prop="remark" label="비고" align="center" />
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="800px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="평형" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].name" />
          </template>
        </el-table-column>
        <el-table-column label="세대수" align="center" width="100">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].familycount" @keyup="onKeyup(scope.$index)" />            
          </template>
        </el-table-column>
        <el-table-column label="면적" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].size" @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>
        <el-table-column label="총면적" align="right" width="100">
          <template #default="scope">
            {{data.batchs[scope.$index].totalsize}}
          </template>
        </el-table-column>
        <el-table-column label="순번" align="center" width="100">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].order" />
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
import { util }  from "~/global"
import { Area } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const store = useStore()
const route = useRoute()

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
  name: '',
  familycount: null,
  size: null,
  order: 0,
  remark: ''  
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
  search: ''
})

async function initData() {  
}

async function getItems() {
  let res = await Area.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'ar_order,ar_id'    
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

  console.log(res.items)
  data.items = res.items
}

function makeItems(items) {
  for (let i = 0; i < items.length; i++) {
    let item = items[i]
    items[i].totalsize = util.area(util.getFloat(item.familycount) * util.getFloat(item.size))
  }

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
    let res = await Area.remove(item)
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

function setWindowSize() {
  data.width = (window.innerWidth - 500) + 'px'
  data.height = (window.innerHeight - 170) + 'px'
}

onMounted(async () => {
  data.apt = parseInt(route.params.id)  
  
  util.loading(true)
  
  await initData()
  await getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  data.visible = false
  util.loading(false)
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
})

function getDong(id) {
  for (let i = 0; i < data.dongs.length; i++) {
    let item = data.dongs[i]

    if (item.id == id) {
      return item.name
    }
  }

  return ''
}

function getDongElevator(id, elevator) {
  if (elevator == null || elevator == 0) {
    return getDong(id)
  }

  for (let i = 0; i < data.dongs.length; i++) {
    let item = data.dongs[i]

    if (item.id == id) {
      return item.name + ` ${elevator}호기`
    }
  }

  return ''
}

function getCategory(id) {
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.id == id) {
      return item.name
    }
  }

  return ''
}

function getStandard(id) {
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.id == id) {
      return item.name
    }
  }

  return ''
}

function changeTopcategory(id) {
  let subcategorys = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      subcategorys.push(item)
    }
  }

  data.subcategorys = subcategorys
  
  data.categorys = []
  data.standards = []
  data.methods = []

  data.item.subcategory = null
  data.item.category = null
  data.item.standard = null
  data.item.method = null
}

function changeSubcategory(id) {
  let categorys = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      categorys.push(item)
    }
  }

  data.categorys = categorys

  data.standards = []
  data.methods = []
  
  data.item.category = null
  data.item.standard = null
  data.item.method = null
}

function changeCategory(id) {
  let methods = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      methods.push(item)
    }
  }

  data.methods = methods

  data.item.method = null

  let standards = []
  
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.category == id) {
      standards.push(item)
    }
  }

  data.standards = standards
  data.item.standard = null
}

function changeSearchTopcategory(id) {
  let subcategorys = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      subcategorys.push(item)
    }
  }

  data.search.subcategorys = subcategorys
  
  data.search.categorys = []
  data.search.standards = []
  data.search.methods = []

  data.search.subcategory = null
  data.search.category = null
  data.search.standard = null
  data.search.method = null
}

function changeSearchSubcategory(id) {
  let categorys = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      categorys.push(item)
    }
  }

  data.search.categorys = categorys
  
  data.search.standards = []
  data.search.methods = []

  data.search.category = null
  data.search.standard = null
  data.search.method = null
}

function changeSearchCategory(id) {
  let methods = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      methods.push(item)
    }
  }

  data.search.methods = methods

  data.search.method = null
  
  let standards = []
  
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.category == id) {
      standards.push(item)
    }
  }

  data.search.standards = standards

  data.search.standard = null  
}

const getSummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index === 2) {
      sums[index] = '총 주택공급면적'
    } else if (index == 5) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += util.getFloat(item.familycount) * util.getFloat(item.size)
        })
      }
      
      sums[index] = util.area(total)    
    }
  })

  return sums
}

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

      await Area.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems()

    util.loading(false)
  })
}

function onKeyup(index) {
  let item = data.batchs[index]

  let totalsize = util.area(util.getFloat(item.familycount) * util.getFloat(item.size))
  data.batchs[index].totalsize = totalsize

  data.batchs[index].name = getPyong(item.size)
  console.log('평 변경')
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
        await Area.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.familycount = util.getInt(item.familycount)
    item.size = util.getFloat(item.size)    
    item.order = util.getInt(item.order)
    item.name = String(item.name)

    console.log(item)
    if (item.id > 0) {
      await Area.update(item)
    } else { 
      await Area.insert(item)
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

function getPyong(value) {
  let ret = parseInt(util.getInt(util.getFloat(value) / 3.3))
  if (ret == 0) {
    return ''
  } else {
    return ret
  }
}
</script>

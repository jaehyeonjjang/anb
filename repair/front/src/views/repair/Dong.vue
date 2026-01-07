<template>
  <Title title="시설물" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄처리</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" sortable />
    <el-table-column prop="name" label="시설물" />      
    <el-table-column prop="ground" label="지상층" align="center" />
    <el-table-column prop="underground" label="지하층" align="center" />
    <el-table-column prop="parking" label="주차장" align="center" />
    <el-table-column prop="elevator" label="승강기" align="center" />
    <el-table-column prop="familycount" label="세대수" align="center" />
    <el-table-column label="공용/전용" align="center" width="80">
      <template #default="scope">
        {{data.basics[scope.row.basic].name}}        
      </template>
    </el-table-column>

    <el-table-column prop="order" label="순번" align="center" />
    <el-table-column prop="remark" label="비고" />    
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
        <el-table-column label="시설물" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].name" />
          </template>
        </el-table-column>
        <el-table-column label="지상층" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].ground" />
          </template>
        </el-table-column>
        <el-table-column label="지하층" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].underground" />
          </template>
        </el-table-column>
        <el-table-column label="주차장" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].parking" />
          </template>
        </el-table-column>
        <el-table-column label="승강기" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].elevator" />
          </template>
        </el-table-column>
        <el-table-column label="세대수" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].familycount" />
          </template>
        </el-table-column>
        <el-table-column label="공용/전용" align="center" width="80">
          <template #default="scope">
            <el-select v-model.number="data.batchs[scope.$index].basic" placeholder="공용/전용" style="width:70px;">           
              <el-option
                v-for="item in data.basics"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
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
import { Dong, Breakdown } from "~/models"
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
  ground: 0,
  underground: 0,
  parking: 0,
  elevator: 0,
  familycount: 0,
  remark: '',
  basic: 1
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
  search: '',
  basics: [{id: 0, name: ''}, {id: 1, name: '전용'}, {id: 2, name: '공용'}]
})

async function initData() {  
}

async function getItems() {
  let res = await Dong.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'd_order,d_id'    
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
    let count = await Breakdown.countByAptDong(data.apt, item.id)
    if (count > 0) {
      util.error('세부내역에 등록된 시설물은 삭제할수 없습니다')
      return
    }
    
    let res = await Dong.remove(item)
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

    let flag = false
    
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id
      }

      let count = await Breakdown.countByAptDong(data.apt, item.id)
      if (count > 0) {
        flag = true        
        continue
      }
      
      await Dong.remove(item)
    }

    if (flag == true) {
      util.error('세부내역에 등록된 시설물은 삭제할수 없습니다')
    } else {
      util.info('삭제되었습니다')
    }
    getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    
    if (item.name == '') {
      util.error('시설물명을 입력하세요')
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
        await Dong.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.familycount = util.getInt(item.familycount)
    item.size = util.getFloat(item.size)    
    item.order = util.getInt(item.order)

    item.basic = util.getInt(item.basic)
    if (item.basic == 0) {
      item.basic = 1
    }

    if (item.id > 0) {
      await Dong.update(item)
    } else { 
      await Dong.insert(item)
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

</script>

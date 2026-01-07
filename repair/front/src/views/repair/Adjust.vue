<template>
  <Title title="단가 조정" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
        <div style="flex:1;text-align:right;gap:5;">
          <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
          <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">등록</el-button>
        </div>
  </div>
  
  <div style="display:flex;">
    <div style="flex:2;">

      <el-tabs v-model="data.menu" @tab-click="clickTab">      
        <el-tab-pane label="총론" name="totalplan">
          <div :style="{'height': data.height, 'overflow': 'auto'}">
            <Totalplan ref="totalplan" />
          </div>
        </el-tab-pane>
        <el-tab-pane label="공사종별 수선계획금액 집계표" name="summary">
          <div :style="{'height': data.height, 'overflow': 'auto'}">
            <Summary ref="summary" />
          </div>
        </el-tab-pane>
      </el-tabs>

      
    </div>
    <div style="width:10px;"></div>
    <div style="flex:2;">
        

      
      <el-table :data="data.items" border :width="data.width" :height="height(170)" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
        <el-table-column type="selection" width="40" align="center" />        
        <el-table-column label="대분류" width="80" :show-overflow-tooltip="true">
          <template #default="scope">
            {{scope.row.topcategory.name}}
          </template>
        </el-table-column>
        <el-table-column label="중분류" width="80" :show-overflow-tooltip="true">
          <template #default="scope">
            {{scope.row.subcategory.name}}
          </template>
        </el-table-column>
        <el-table-column label="공사종별">
          <template #default="scope">
            {{getCategory(scope.row.category).name}}
          </template>
        </el-table-column>
        <el-table-column label="규격">
          <template #default="scope">
            {{getStandard(scope.row.standard).name}}
          </template>
        </el-table-column>        
        <el-table-column label="비율" align="right" width="50">
          <template #default="scope">
            {{util.getFloat(scope.row.rate)}}
          </template>
        </el-table-column>
        <el-table-column prop="order" label="순번" align="center" width="40" />
      </el-table>  

    </div>
  </div>
  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="700px"
  >

          
      <el-table :data="data.batchs" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;width:100%;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="공사종별" align="center" width="300">
          <template #default="scope">
            <el-tree-select style="width:290px;" v-model="data.batchs[scope.$index].category" :data="data.categorys" check-strictly :default-expand-all="false" :render-after-expand="false" placeholder="공사종별" @node-click="changeCategory" />
          </template>
        </el-table-column>
        <el-table-column label="규격" align="center" width="180">
          <template #default="scope">
            <el-select v-model.number="data.item.standard" class="m-2" placeholder="규격">
              <el-option
                v-for="item in data.standards"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
            

          </template>
        </el-table-column>
        <el-table-column label="비율" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].rate" />
          </template>
        </el-table-column>
        <el-table-column label="순번" align="center" width="100">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].order" />
          </template>
        </el-table-column>                

      </el-table>




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
import { Adjust, Standard, Category, Report } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const { width, height } = size()

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
  type: null,
  category: null,
  standard: null,
  rate: 0,
  order: 0
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
  allcategorys: [],
  categorys: [],
  standards: [],
  allstandards: [],
  width: 0,
  height: 0,
  menu: 'totalplan'
})

const summary = ref({});
const totalplan = ref({});

function clickTab(item) {
  let name = item.props.name

  if (name == 'summary') {
    summary.value.readData()
  } else if (name == 'totalplan') {
    totalplan.value.readData()    
  }
}

async function initData() {
  if (data.apt == 0) {
    return
  }
  
  let {allcategorys, categorys} = await util.getCategoryTree(data.apt, '공사종별')
  data.allcategorys = allcategorys 
  data.categorys = categorys

  let res = await Standard.findByApt(data.apt)
  data.allstandards = res.items
}

async function getItems() {
  let res = await Adjust.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'aj_order,aj_id'    
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      let item = res.items[i]

      let category = getCategory(item.category)
      let subcategory = getCategory(category.parent)
      let topcategory = getCategory(subcategory.parent)

      res.items[i].topcategory = topcategory
      res.items[i].subcategory = subcategory
      res.items[i].index = i + 1
    }
  }

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }
  data.items = res.items

  if (data.menu == 'summary') {
    summary.value?.readData()
  } else {
    totalplan.value?.readData()
  }
}

function makeItems(items) {
  for (let i = 0; i < items.length; i++) {
    let item = items[i]    
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
    let res = await Adjust.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')

      if (data.menu == 'summary') {
        summary.value.readData()
      } else {
        totalplan.value.readData()
      }
      
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

function getCategory(id) {
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.id == id) {
      return item
    }
  }

  return {id: 0, name: ''}
}

function getStandard(id) {
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.id == id) {
      return item
    }
  }

  return {id: 0, name: ''}
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

    let items = []
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id,
        apt: data.apt
      }

      items.push(item)
    }

    await Adjust.removebatch(items)
    util.info('삭제되었습니다')

    if (data.menu == 'summary') {
      summary.value.readData()
    } else {
      totalplan.value.readData()
    }
    
    getItems()

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
        await Adjust.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.category = util.getInt(item.category)
    item.standard = util.getInt(item.standard)
    item.rate = util.getFloat(item.rate)
    item.order = util.getInt(item.order)

    if (item.id > 0) {
      await Adjust.update(item)
    } else { 
      await Adjust.insert(item)
    }
  }

  util.info('등록되었습니다')

  if (data.menu == 'summary') {
    summary.value.readData()
  } else {
    totalplan.value.readData()
  }
  
  getItems()
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

function changeCategory(item) {
  let index = 0
  let id = item.value

  let standards = [{id: 0, name: '규격'}]  
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.category == id) {
      standards.push(item)
    }
  }

  data.standards = standards
  data.batchs[index].standard = null
}

</script>

<template>
  <Title title="공사종별" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <el-tree-select style="width:290px;" v-model="data.search.category" :data="data.categorys" check-strictly :default-expand-all="false" :render-after-expand="false"  @node-click="clickSearch" placeholder="분류" />

    <el-input v-model="data.search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />
    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">등록</el-button>      
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" />
    <el-table-column label="대분류">
      <template #default="scope">
        <div v-if="scope.row.level == 1">{{scope.row.name}}</div>
      </template>
    </el-table-column>
    <el-table-column label="중분류">
      <template #default="scope">
        <div v-if="scope.row.level == 2">{{scope.row.name}}</div>
      </template>
    </el-table-column>
    <el-table-column label="공사종별">
      <template #default="scope">
        <div v-if="scope.row.level == 3">{{scope.row.name}}</div>
      </template>
    </el-table-column>
    <el-table-column label="수선방법">
      <template #default="scope">
        <div v-if="scope.row.level == 4">{{scope.row.name}}</div>
      </template>
    </el-table-column>
    <el-table-column label="구분" align="center" width="80">
      <template #default="scope">
        {{getLevel(scope.row.level)}}
      </template>
    </el-table-column>
    <el-table-column label="수선주기" align="center" width="80">
      <template #default="scope">
        <div v-if="scope.row.level == 4">{{scope.row.cycle}}</div>
      </template>
    </el-table-column>
    <el-table-column label="수선율" align="center" width="80">
      <template #default="scope">
        <div v-if="scope.row.level == 4">{{scope.row.percent}}</div>
      </template>
    </el-table-column>
    <el-table-column label="승강기" width="60" align="center">
      <template #default="scope">
        <div v-if="scope.row.elevator == 1">사용</div>
      </template>
    </el-table-column>
    <el-table-column prop="order" label="순번" width="100" align="center" />
    <el-table-column prop="remark" label="비고" />
        
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1100px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>

        
        
        <el-table-column prop="name" label="분류" width="300">
          <template #default="scope">
            <el-tree-select style="width:290px;" v-model="data.batchs[scope.$index].parent" :data="data.categorys" check-strictly :default-expand-all="false" :render-after-expand="false" placeholder="분류" />                    
          </template>
        </el-table-column>
        <el-table-column prop="content" label="분류명">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].name" />
          </template>
        </el-table-column>
        <el-table-column label="수선주기" align="center" width="70">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].cycle" />
          </template>
        </el-table-column>
        <el-table-column label="수선율" align="center" width="70">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].percent" />
          </template>
        </el-table-column>
        <el-table-column label="승강기" width="100" align="center">
          <template #default="scope">
            <el-select v-model.number="data.batchs[scope.$index].elevator" class="m-2" placeholder="승강기">          
              <el-option
                v-for="item in data.elevators"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>        
        <el-table-column label="순번" align="center" width="70">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].order" />
          </template>
        </el-table-column>
        <el-table-column label="비고">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].remark" />
          </template>
        </el-table-column>
      </el-table>


    </el-form>

      <template #footer>
        <el-button size="small" type="danger" v-if="data.mode != 'batch' && (data.batchs.length > 0 && data.batchs[0].id > 0)" style="float:left;" @click="clickDelete">삭제</el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Category } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const props = defineProps({
  height: Number
})

const store = useStore()
const route = useRoute()

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
  level: null,
  parent: null,
  cycle: '',
  percent: '',
  remark: '',
  unit: '',
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
  search: {
    text: '',
    category: null
  },
  allcategorys: [],
  categorys: [],  
  elevators: [{id: 1, name: '사용'}, {id: 2, name: '사용안함'}],
  height: 700
})

async function initData() {
  let {allcategorys, categorys} = await util.getCategoryTree(data.apt, '분류')
  data.allcategorys = allcategorys 
  data.categorys = categorys  
}

async function getItems() {
  let res = await Category.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    name: data.search.text,
    orderby: 'c_order,c_id'    
  })

  if (res.items == null) {
    res.items = []
  }

  let items = res.items 

  /*
     if (data.search.text != '') {
     items = items.filter(item => item.name.indexOf(data.search.text) >= 0)        
     }
   */

  if (data.search.category != null && data.search.category != 0) {
    items = items.filter(item => item.parent == data.search.category || item.id == data.search.category)  
  }

  if (items.length > 0 && items[0].id == 0) {
    items.splice(0, 1)
  }
  
  for (let i = 0; i < items.length; i++) {
    items[i].index = i + 1
  }  

  data.total = res.total  
  data.items = items
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

  let items = util.clone(data.items)

  data.batchs = items

  let row = item.index - 1
  
  data.batchs = [data.batchs[row]]
  data.mode = 'normal'  
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Category.remove(item)
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
  data.height = (window.innerHeight - 170 - props.height) + 'px'  
}

onMounted(async () => {
  data.apt = parseInt(route.params.id)  
  
  if (util.getInt(data.apt) == 0) {
    data.apt = -1
  }  
  
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

  return {
    id: 0,
    name: ''
  }
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

  data.batchs = items
  
  data.mode = 'batch' 
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

      await Category.remove(item)
    }

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    
    if (item.name == '') {
      util.error('분류명을 입력하세요')
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
        await Category.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.cycle = util.getInt(item.cycle)
    item.percent = util.getInt(item.percent)
    item.parent = util.getInt(item.parent)
    item.elevator = util.getInt(item.elevator)
    if (item.elevator == 0) {
      item.elevator = 2
    }
    item.order = util.getInt(item.order)
    item.category = util.getInt(item.category)

    if (item.parent == 0) {
      item.level = 1
    } else {
      let parent = getCategory(item.parent)
      item.level = parent.level + 1
    } 

    if (item.id > 0) {
      await Category.update(item)
    } else { 
      await Category.insert(item)
    }
  }

  util.info('등록되었습니다')
  
  getItems()
  data.visible = false  
  util.loading(false)  
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function clickAdd() {  
  data.batchs.push(util.clone(item))
}

const levels = ['', '대분류', '중분류', '공사종별', '수선방법']
function getLevel(value) {
  return levels[value]
}

</script>

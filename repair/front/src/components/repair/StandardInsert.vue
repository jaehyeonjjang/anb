
<template>
  <Title title="규격" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <el-tree-select style="width:290px;" v-model="data.search.category" :data="data.categorys" check-strictly :default-expand-all="false" :render-after-expand="false" placeholder="공사종별" />

    <el-input v-model="data.search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />
    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;" v-if="data.apt > 0 || isAdmin()">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">등록</el-button>      
    </div>
  </div>  
    
  <el-table :data="data.items" border :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList" v-infinite="getItems">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" />


    <el-table-column label="대분류">
      <template #default="scope">
        {{getTopcategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column label="중분류">
      <template #default="scope">
        {{getSubcategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column label="공사종별">
      <template #default="scope">
        {{getCategory(scope.row.category)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="규격명" />
    
    <el-table-column label="재료비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.direct)}}
      </template>
    </el-table-column>
    
    <el-table-column label="노무비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.labor)}}
      </template>
    </el-table-column>
    
    <el-table-column label="경비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.cost)}}
      </template>
    </el-table-column>
    
    <el-table-column label="단가" align="right" width="90">
      <template #default="scope">
        {{util.money(util.calculatePrice(scope.row.direct, scope.row.labor, scope.row.cost))}}
      </template>
    </el-table-column>

    <el-table-column prop="unit" label="규격" align="center" width="60" />
    <el-table-column prop="order" label="순번" align="center" width="60" />
        
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

        
        
        <el-table-column prop="name" label="공사종별" width="300">
          <template #default="scope">
            <el-tree-select style="width:290px;" v-model="data.batchs[scope.$index].category" :data="data.categorys" :default-expand-all="false" :render-after-expand="false" placeholder="공사종별" />                    
          </template>
        </el-table-column>

        <el-table-column label="규격명">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].name" />
          </template>
        </el-table-column>
        
        <el-table-column label="재료비" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].direct"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>
        <el-table-column label="노무비" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].labor"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>
        <el-table-column label="경비" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].cost"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>
        
        <el-table-column label="단가" align="right" width="100">
          <template #default="scope">
            {{util.money(util.calculatePrice(scope.row.direct, scope.row.labor, scope.row.cost))}}
          </template>
        </el-table-column>

        <el-table-column label="규격" align="center" width="80">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].unit" />
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
      <el-button size="small" style="float:left;" @click="data.visibleCategory = true">공사종별</el-button>
      <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd"><el-icon><Plus /></el-icon></el-button>

      <div style="float:left;margin-left:20px;" v-if="data.mode != 'batch'">
        <el-checkbox v-model="data.batch" label="기존 작업에 일괄등록" size="small" style="margin:0px 0px;" v-show="data.apt == -1" />
      </div>
      
      <el-button size="small" @click="clickCancel">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleCategory"
    width="1100px"
    :before-close="handleCloseCategory"
  >
    <CategoryInsert :height="300" />
  </el-dialog>  
  
</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted, computed, watch } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Category, Standard, Standardlist, Repair } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const props = defineProps({
  height: Number,
  category: Number
})

const store = useStore()
const route = useRoute()

async function clickSearch() {  
  await getItems(true)
}

const item = {
  id: 0,
  name: '',
  category: null,
  direct: 0,
  labor: 0,
  cost: 0,
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
  pagesize: 100,
  item: util.clone(item),
  visible: false,
  visibleCategory: false,
  search: {
    text: '',
    category: 0
  },
  categorys: [],
  allcategorys: [],
  elevators: [{id: 1, name: '사용'}, {id: 2, name: '사용안함'}],
  height: 700,
  batch: false,
  batchs: [],
  selectMode: null
})

/*
   watch(() => data.search.category, (first, second) => {
   console.log(data.search.category)
   console.log(
   "Watch props.selected function called with args:",
   first,
   second
   );
   });
 */

async function initData() {
  let {allcategorys, categorys} = await util.getCategoryTree(data.apt, '공사종별')
  data.allcategorys = allcategorys 
  data.categorys = categorys    
}

async function readItems() {  
  await getItems()
}

defineExpose({
  readItems,
  setCategory,
  setSelectMode,
  getSelectMode
})

async function getItems(reset: boolean) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }

  util.loading(true)
  
  let topcategory = 0
  let subcategory = 0
  let category = 0
  
  let searchCategory = getCategoryInfo(data.search.category)
  if (searchCategory.level == 1) {
    topcategory = searchCategory.id
  } else if (searchCategory.level == 2) {
    subcategory = searchCategory.id
  } else if (searchCategory.level == 3) {
    category = searchCategory.id
  }

  let res = await Standardlist.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    topcategory: topcategory,
    subcategory: subcategory,
    category: category,
    name: data.search.text,
    orderby: 's_categoryorder,s_category,s_order,s_id'    
  })

  if (res.items == null) {
    res.items = []
  } else {
    data.page = data.page + 1
  }

  let items = data.items.concat(res.items) 

  /*
     if (data.search.text != '') {
     items = items.filter(item => item.name.indexOf(data.search.text) >= 0)
     }

     if (util.getInt(data.search.category) != 0) {
     let category = util.findCategory(data.search.category, data.categorys)
     let categorys = util.getCategoryChildren(category)

     items = items.filter(item => categorys.find(c => c.value == item.category))
     }
   */

  if (items.length > 0 && items[0].id == 0) {
    items.splice(0, 1)
  }
  
  for (let i = 0; i < items.length; i++) {
    items[i].index = i + 1
  }  

  data.total = res.total  
  data.items = items

  util.loading(false)
}

function makeItems(items) {
  items.map(item => item.price = util.money(util.calculatePrice(item.direct, item.labor, item.cost)))
  return items
}

function clickInsert() {  
  data.item = util.clone(item)

  let category = getCategoryInfo(data.search.category)
  if (category.level == 3) {
    data.item.category = data.search.category
  }
  
  let items = makeItems([data.item])

  /*
     if (data.apt == -1) {
     data.batch = true
     } else {
     data.batch = false
     }
   */
  data.batch = false
  
  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickUpdate(item, index) {
  if (!isAdmin() && data.apt <= 0) {
    return
  }
  
  if (index.no == 0) {
    return
  }

  if (data.selectMode != null) {
    data.selectMode(util.clone(item))
    return
  }

  let items = util.clone(data.items)

  data.batchs = items

  let row = item.index - 1
  
  data.batchs = [data.batchs[row]]

  /*
     if (data.apt == -1) {
     data.batch = true
     } else {
     data.batch = false
     }
   */
  data.batch = false
  
  data.mode = 'normal'  
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Standard.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      await getItems(true)
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

async function setCategory(category) {
  data.search.category = category
  await clickSearch()  
}

onMounted(async () => {
  data.apt = parseInt(route.params.id)
  
  if (util.getInt(data.apt) == 0) {
    data.apt = -1
  }  
  
  await initData()
  await getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  data.visible = false
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

      await Standard.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems(true)

    util.loading(false)
  })
}

async function clickSubmit() {
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    
    if (util.getInt(item.category) == 0) {
      util.error('공사종별을 선택하세요')
      return    
    }
    
    if (item.name == '') {
      util.error('규격명을 입력하세요')
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
        await Standard.remove(item)
      }
    }
  }

  let flag = false

  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.direct = util.getInt(item.direct)
    item.labor = util.getInt(item.labor)
    item.cost = util.getInt(item.cost)
    item.category = util.getInt(item.category)    

    if (item.id > 0) {
      if (data.batch == true) {
        await Standard.updateall(item)

        flag = true
      } else {
        await Standard.update(item)
      }

      data.items[item.index - 1] = item
    } else {
      flag = true

      if (data.batch == true) {
        await Standard.all(item)
      } else { 
        await Standard.insert(item)
      }
    }
  }

  util.loading(false)
  if (data.batch == true) {
    util.info('등록되었습니다')
  }
  
  if (flag == true) {
    await getItems(true)
  }
  
  data.visible = false  
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function clickAdd() {  
  data.batchs.push(util.clone(item))
}

function getCategoryInfo(id) {
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.id == id) {
      return item;
    }
  }

  return {
    id: 0,
    name: ''
  }
}

function getTopcategory(id) {
  let category = getCategoryInfo(id)

  let subcategory = getCategoryInfo(category.parent)
  let topcategory = getCategoryInfo(subcategory.parent)


  return topcategory.name
}

function getSubcategory(id) {
  let category = getCategoryInfo(id)
  let subcategory = getCategoryInfo(category.parent)

  return subcategory.name  
}

function getCategory(id) {
  let category = getCategoryInfo(id)

  return category.name  
}

function onKeyup(index) {
  let item = data.batchs[index]
  let direct = util.getInt(item.direct)
  let labor = util.getInt(item.labor)
  let cost = util.getInt(item.cost)

  let price = util.money(util.calculatePrice(direct, labor, cost)) 
  data.batchs[index].price = price
}

function changeSearchCategory(item) {
  data.search.category = item.value

  clickSearch()
}

const handleCloseCategory = async (done: () => void) => {
  data.categorys = await util.getCategoryTree(data.apt, '공사종별')
  
  done()
}

function setSelectMode(value) {
  data.selectMode = value
}

function getSelectMode() {
  if (data.selectMode == null) {
    return false
  } else {
    return true
  }
}

function isAdmin() {
  if (store.getters['getUser'].loginid == 'yuhki') {
    return true
  }

  if (store.getters['getUser'].loginid == 'A001') {
    return true
  }

  return false
}
</script>

<style>
.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>

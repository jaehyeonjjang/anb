<template>
  <Title title="사용현황" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <el-select v-model.number="data.search.year" placeholder="년도" style="width:100px;">           
      <el-option
        v-for="item in data.years"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>

    <el-tree-select style="width:290px;" v-model="data.search.category" :data="data.categorys" check-strictly :default-expand-all="false" :render-after-expand="false" placeholder="공사종별" />

    <el-input v-model="data.search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />
    
    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄등록</el-button>
      <!--<el-button size="small" type="warning" @click="clickBatch">일괄처리</el-button>-->
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height"  :summary-method="getSummaries" show-summary ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" sortable />
    <el-table-column prop="name" label="년도" align="center" width="100" sortable>
      <template #default="scope">
        {{scope.row.year}}년 {{util.pad(scope.row.month, 2)}}월
      </template>
    </el-table-column>
    <el-table-column prop="name" label="대분류" sortable>
      <template #default="scope">
        {{getCategory(scope.row.topcategory).name}}
      </template>
    </el-table-column>
    <el-table-column prop="name" label="중분류" sortable>
      <template #default="scope">
        {{getCategory(scope.row.subcategory).name}}
      </template>
    </el-table-column>
    <el-table-column prop="name" label="공사종별" sortable>
      <template #default="scope">
        {{getCategory(scope.row.category).name}}
      </template>
    </el-table-column>    
    <el-table-column prop="content" label="보수내역" />
    <el-table-column label="사용금액" align="right" width="100">
      <template #default="scope">
        <div style="color:#000;">{{util.money(scope.row.price)}}</div>
      </template>
    </el-table-column>
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    :fullscreen="data.fullscreen"
    width="900px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border :height="data.popupHeight" :key="data.width+''+data.popupHeight" style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="년월" width="130" align="center">
          <template #default="scope">
            <el-date-picker style="margin: 0px 0px;height: 24px;width:120px;" v-model="data.batchs[scope.$index].yearmonth" type="month" placeholder="년월" />
          </template>
        </el-table-column>
        <el-table-column label="공사종별" width="300">
          <template #default="scope">
            <el-tree-select style="width:290px;" v-model="data.batchs[scope.$index].category" :data="data.categorys" :default-expand-all="false" :render-after-expand="false" placeholder="공사종별" />
          </template>
        </el-table-column>

        <el-table-column prop="content" label="보수내역">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content" class="inputText" />
          </template>
        </el-table-column>
        <el-table-column label="사용금액" width="100" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].pricestring"
                      :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                      :parser="(value) => value.replace(/\$\s?|(,*)/g, '')"
                      class="inputNumber"
                      @keypress.enter.native="clickNextInput(scope.$index)"
            />
          </template>
        </el-table-column>
      </el-table>


    </el-form>

      <template #footer>
        <el-button size="small" type="danger" v-if="data.mode != 'batch' && (data.batchs.length > 0 && data.batchs[0].id > 0)" style="float:left;" @click="clickDelete">삭제</el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(1)"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(10)"><el-icon><Plus /></el-icon> &nbsp;10</el-button>
        <!--<el-button size="small" style="float:left;" @click="data.visibleCategory = true">공사종별</el-button>-->        
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

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Category, History } from "~/models"
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
  year: '',
  month: '',
  topcategory: null,
  subcategory: null,
  category: null,
  content: '',
  price: 0,
  yearmonth: '',
  pricestring: ''
}

const data = reactive({
  apt: 0,
  repair: null,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  visibleCategory: false,
  search: {
    text: '',
    category: null,
    year: null
  },
  allcategorys: [],
  categorys: [],
  fullscreen: false,
  width: 0,
  height: 0,
  popupHeight: 0,
  oldHeight: 0,
  years: [],
  categoryMap: {}
})

async function initData() {
  if (data.apt == 0) {
    return
  }
  
  data.repair = await util.getRepair(data.apt)

  let {allcategorys, categorys} = await util.getCategoryTree(data.apt, '공사종별')
  data.allcategorys = allcategorys 
  data.categorys = categorys

  let categoryMap = {}
  allcategorys.forEach((item) => {
    categoryMap[item.id] = item
  })
  data.categoryMap = categoryMap
  
  let topcategorys = []
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.level == 1) {
      topcategorys.push(item)
    }
  }

  data.topcategorys = topcategorys  
}

async function getItems() {
  let topcategory = 0
  let subcategory = 0
  let category = 0
  
  let searchCategory = getCategory(data.search.category)
  if (searchCategory.level == 1) {
    topcategory = searchCategory.id
  } else if (searchCategory.level == 2) {
    subcategory = searchCategory.id
  } else if (searchCategory.level == 3) {
    category = searchCategory.id
  } 
  
  let res = await History.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    //orderby: 'h_year,h_month,c_order,h_id',
    orderby: 'h_id',    
    content: data.search.text,
    topcategory: topcategory,
    subcategory: subcategory,
    category: category,
    year: util.getInt(data.search.year)
  })

  let years = [{id: 0, name: '년도'}]
  let year = util.getInt(data.search.year)
  let flag = false 
  
  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
      years.push({
        id: res.items[i].year,
        name: res.items[i].year
      })

      if (year != 0) {
        if (year == res.items[i].year) {
          flag = true
        }
      }
    }
  }

  if (flag == false) {
    data.search.year = null
  }

  data.years = years

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }
  data.items = res.items
}

function makeItems(items) {
  for (let i = 0; i < items.length; i++) {
    let item = items[i];

    if (util.getInt(item.year) != 0 && util.getInt(item.month) != 0) { 
      items[i].yearmonth = util.pad(item.year, 4) + '-' + util.pad(item.month, 2)
    }

    if (items[i].price == 0) {
      items[i].pricestring = ''
    } else {
      items[i].pricestring = util.money(items[i].price)
    }
  }

  return items
}

function clickInsert() {  
  data.item = util.clone(item)

  let items = makeItems([data.item])

  data.mode = 'insert'
  data.batchs = items    
  data.popupHeight = '58px'
  data.fullscreen = false
  data.visible = true
}

function clickUpdate(item, index) {
  if (index.no == 0) {
    return
  }
  
  let row = item.index - 1
  item = util.clone(data.items[row])
  item.yearmonth = util.pad(item.year, 4) + '-' + util.pad(item.month, 2)
  if (item.price == 0) {
    item.pricestring = ''
  } else {
    item.pricestring = util.money(item.price)
  }

  
  data.batchs = [item]
  data.popupHeight = '58px'
  data.mode = 'update'
  data.fullscreen = false
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await History.remove(item)
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
  data.popupHeight = (window.innerHeight - 170) + 'px'
  data.oldHeight = (window.innerHeight - 170) + 'px'
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
  let item = data.categoryMap[id]

  if (item == null || item == undefined) {
    return {
      id: 0,
      name: ''
    } 
  }

  return item  
}

const getSummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index === 2) {
      sums[index] = '사용 계'
    } else if (index == 7) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += item.price
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

function clickBatch() {
  let items = util.clone(data.items)

  if (items == null) {
    items = []
  }

  items = []
  
  if (items.length == 0) {
    for (let i = 0; i < 5; i++) {
      items.push(util.clone(item))
    }
  }
  
  items = makeItems(items)

  data.batchs = items
  
  data.mode = 'batch'
  data.popupHeight = data.oldHeight
  data.fullscreen = true
  data.visible = true  
}

function clickCancel() {
  util.confirm('팝업창을 닫으시겠습니까', function() {
    data.visible = false
  })
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
        id: value.id
      }

      items.push(item)
    }

    await History.removebatch(items)

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  let d = new Date()
  let currentYear = d.getFullYear()
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]


    if (typeof item.yearmonth == 'string') {
      let strs = item.yearmonth.split('-')
      if (strs.length != 2 && data.mode != 'batch') {
        util.error('년월을 입력하세요')
        return
      }

      item.year = util.getInt(strs[0])
      item.month = util.getInt(strs[1])
    } else {
      if (item.yearmonth == null) {
        if (strs.length != 2 && data.mode != 'batch') {
          util.error('년월을 입력하세요')
          return
        }

        item.year = 0
        item.month = 0
      } else {
        item.year = item.yearmonth.getFullYear()
        item.month = item.yearmonth.getMonth() + 1
        }
    }
    
    if (item.id > 0 || data.mode != 'batch') {
      if (util.isNull(item.yearmonth)) {
        util.error('년월을 입력하세요')
        return
      }
      

      let year = util.getInt(item.year)
      let month = util.getInt(item.month)
      if (year == 0) {
        util.error('년도를 입력하세요')
        return    
      }

      if (data.repair.completionyear > 0) { 
        if (year < data.repair.completionyear) {
          util.error('년도가 사용검사 년도 이전이 될 수 없습니다')
          return    
        }

        if (year > currentYear) {
          util.error('년도값이 정확하지 않습니다')
          return    
        }
      }

      if (month == 0) {
        util.error('월을 입력하세요')
        return    
      }

      if (month > 12) {
        util.error('월값이 정확하지 않습니다')
        return    
      }
      
      if (item.content == '') {
        util.error('보수내역을 입력하세요')
        return    
      }
    }
  }

  util.loading(true)

  /*
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
        await History.remove(item)
      }
    }
  }
  */

  let count = 0
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.year = util.getInt(item.year)
    item.month = util.getInt(item.month)
    item.price = util.getInt(item.pricestring.replace(/,/g, ''))
    item.category = util.getInt(item.category)

    if (item.id == 0) {
      if (item.year == 0 || item.month == 0) {
        continue
      }
    }
    
    let category = getCategory(item.category)
    let subcategory = getCategory(category.parent)
    let topcategory = getCategory(subcategory.parent)

    item.subcategory = subcategory.id
    item.topcategory = topcategory.id


    if (item.id > 0) {
      await History.update(item)
    } else { 
      await History.insert(item)
      count++
    }
  }

  util.info('등록되었습니다')
  
  getItems()
  

  data.visible = false  
  util.loading(false)

  if (data.mode != 'update' && count > 0) {
    setTimeout(function() {
      listRef.value!.setScrollTop(data.items.length  * 100 + 1000)
    }, 500)
  }
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

const handleCloseCategory = async (done: () => void) => {
  data.categorys = await util.getCategoryTree(data.apt, '공사종별')
  
  done()
}

function clickNextInput(index) {
  if (index == data.batchs.length - 1) {
    clickSubmit()
  }
}
</script>
<style>
.inputNumber .el-input__inner {
  text-align: right;
}

.inputText .el-input__inner {
  text-align: left;
}

</style>

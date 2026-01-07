<template>
  <Title title="항목검토" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
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

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" sortable />
    <el-table-column label="구분" align="center" width="50">
      <template #default="scope">
        <el-tag v-if="scope.row.id == 0" type="success">자동</el-tag>
        <el-tag v-else type="warning">직접</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="name" label="대분류" sortable width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.topcategory).name}}
      </template>
    </el-table-column>
    <el-table-column prop="name" label="중분류" sortable width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.subcategory).name}}
      </template>
    </el-table-column>
    <el-table-column prop="name" label="공사종별" sortable width="90" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.category).name}}
      </template>
    </el-table-column>
    <el-table-column prop="name" label="규격" sortable width="150" :show-overflow-tooltip="true">
      <template #default="scope">
        {{scope.row.extra.standard.name}}
      </template>
    </el-table-column>
    <el-table-column label="수선방법" width="60" align="center">
      <template #default="scope">
        {{getCategory(scope.row.method).name}}
      </template>
    </el-table-column>
    <el-table-column prop="content" label="검토사유" />
    <el-table-column prop="adjust" label="조정내용" />
    <el-table-column prop="cycle" label="주기" align="center" width="50"  />
    <el-table-column prop="percent" label="수선율" align="center" width="50" />
    <el-table-column prop="count" label="수량" align="right" width="50" />
    <el-table-column label="수선금액" align="right" width="100">
      <template #default="scope">
        <div style="color:#000;">{{util.money(scope.row.price)}}</div>
      </template>
    </el-table-column>
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    :fullscreen="data.fullscreen"
    width="1200px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border :height="data.popupHeight" :key="data.width+''+data.popupHeight" style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>        
        <el-table-column label="공사종별" width="300">
          <template #default="scope">
            <el-tree-select style="width:290px;" v-model="data.batchs[scope.$index].category" :data="data.batchs[scope.$index].categorys" :default-expand-all="false" placeholder="공사종별" @node-click="changeCategory" />
          </template>
        </el-table-column>

        <el-table-column label="규격">
          <template #default="scope">
            <el-select v-model.number="data.batchs[scope.$index].standard" placeholder="규격">
              <el-option
                v-for="item in data.batchs[scope.$index].standards"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>

        <el-table-column label="수선방법" width="100" align="center">
          <template #default="scope">
            <el-select v-model.number="data.batchs[scope.$index].method" placeholder="수선방법">
              <el-option
                v-for="item in data.batchs[scope.$index].methods"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>
        
        <el-table-column label="검토사유">
          <template #default="scope">
            <el-select v-model="data.batchs[scope.$index].content" placeholder="" style="width:100%;"
                       filterable
                       allow-create
                       default-first-option
                       @change="changeContent(scope.$index, scope.row)"
                       :reserve-keyword="false">

              <el-option
                v-for="item in data.basics"
                :key="item.id"
                :label="item.content"
                :value="item.content"
              />
            </el-select>            
          </template>
        </el-table-column>
        <el-table-column label="조정내용">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].adjust" class="inputText" />
          </template>
        </el-table-column>
        <el-table-column label="주기" width="60" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].cycle" class="inputText" />
          </template>
        </el-table-column>
        <el-table-column label="수선율" width="50" align="center">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].percent" class="inputNumber" />
          </template>
        </el-table-column>
        <el-table-column label="수량" width="80" align="center">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].count" class="inputNumber" />
          </template>
        </el-table-column>
        <el-table-column label="수선금액" width="100" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].pricestring"
                      :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                      :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
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
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Repair, Standardlist, Category, Review, Reviewbasic, Breakdownhistory } from "~/models"
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
  topcategory: null,
  subcategory: null,
  category: null,
  standard: null,
  method: null,
  content: '',
  adjust: '',
  cycle: '',
  percent: null,
  count: null,
  price: null
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
  search: {
    text: '',
    category: null
  },
  allcategorys: [],
  categorys: [],
  fullscreen: false,
  width: 0,
  height: 0,
  popupHeight: 0,
  oldHeight: 0,
  categoryMap: {},
  basics: [],
  allstandards: []
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

  let res = await Reviewbasic.find({
    orderby: 'rv_order,rv_id'
  })

  if (res.items == null) {
    res.items = []
  }

  data.basics = [{id: 0, content: '', adjust: ''}, ...res.items]

  res = await Standardlist.find({
    apt: data.apt,
    orderby: 's_categoryorder,s_category,s_order,s_id'
  })

  if (res.items == null) {
    res.items = []
  }
  data.allstandards = res.items
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

  
  let res = await Breakdownhistory.auto(data.apt)

  let autos = []
  if (res.items != null) {
    for (let i = 0; i < res.items.length; i++) {
      let item = res.items[i]
      if (category > 0) {
        if (category != item.category) {
          continue
        }
      }

      if (subcategory > 0) {
        if (subcategory != item.subcategory) {
          continue
        }
      }

      if (topcategory > 0) {
        if (topcategory != item.topcategory) {
          continue
        }
      }
      
      autos.push(item)
    }
  }
  
  res = await Review.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 're_id',    
    content: data.search.text,
    topcategory: topcategory,
    subcategory: subcategory,
    category: category
  })

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }

  res.items = autos.concat(res.items)

  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1
  }
  
  data.items = res.items
}

function makeItems(items) {
  for (let i = 0; i < items.length; i++) {
    let item = items[i];    
  }

  return items
}

function clickInsert() {
  let index = 0
  data.item = util.clone(item)
  data.item.index = index
  
  let topcategorys = util.clone(data.categorys)
  for (let i = 1; i < topcategorys.length; i++) {
    let topcategory = topcategorys[i]

    topcategorys[i].index = index 
    for (let j = 0; j < topcategory.children.length; j++) {
      let subcategory = topcategorys[i].children[j]

      topcategorys[i].children[j].index = index
      
      for (let k = 0; k < subcategory.children.length; k++) {
        topcategorys[i].children[j].children[k].index = index
      }
    }
  }

  data.item.categorys = topcategorys

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

  if (item.id == 0) {
    return
  }
  
  let row = item.index - 1
  item = util.clone(data.items[row])
  if (item.price == 0) {
    item.pricestring = ''
  } else {
    item.pricestring = util.money(item.price)
  }

  if (item.standard == 0) {
    item.standard = null
  }

  let topcategorys = util.clone(data.categorys)
  for (let i = 1; i < topcategorys.length; i++) {
    let topcategory = topcategorys[i]

    topcategorys[i].index = index 
    for (let j = 0; j < topcategory.children.length; j++) {
      let subcategory = topcategorys[i].children[j]

      topcategorys[i].children[j].index = index
      
      for (let k = 0; k < subcategory.children.length; k++) {
        topcategorys[i].children[j].children[k].index = index
      }
    }
  }

  item.categorys = topcategorys

  let methods = [{id: 0, name: '수선방법'}]
  let standards = [{id: 0, name: '규격'}]
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let nitem = data.allcategorys[i]

    if (nitem.parent == item.category) {
      methods.push(nitem)
    }
  }

  item.methods = methods

  let flag = false
  
  for (let i = 0; i < data.allstandards.length; i++) {
    let nitem = data.allstandards[i]

    if (nitem.category == item.category) {
      standards.push(nitem)
    }
  }

  item.standards = standards
  
  data.batchs = [item]
  data.popupHeight = '58px'
  data.mode = 'update'
  data.fullscreen = false
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Review.remove(item)
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

function clickBatch() {
  let items = util.clone(data.items)

  if (items == null) {
    items = []
  }

  items = []
  
  if (items.length == 0) {
    for (let l = 0; l < 5; l++) {
      let n = util.clone(item)
      let index = l
      n.index = index
      
      let topcategorys = util.clone(data.categorys)
      for (let i = 1; i < topcategorys.length; i++) {
        let topcategory = topcategorys[i]

        topcategorys[i].index = index 
        for (let j = 0; j < topcategory.children.length; j++) {
          let subcategory = topcategorys[i].children[j]

          topcategorys[i].children[j].index = index
          
          for (let k = 0; k < subcategory.children.length; k++) {
            topcategorys[i].children[j].children[k].index = index
          }
        }
      }

      n.categorys = topcategorys
      
      items.push(n)
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

      if (value.id == 0) {
        continue
      }
      
      let item = {
        id: value.id
      }

      items.push(item)
    }

    await Review.removebatch(items)

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

async function clickSubmit() {  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    if (data.batchs.length == 1 && util.getInt(item.category) == 0) {
      util.error('공사종별을 선택하세요')
      return
    } if (util.getInt(item.category) == 0) {
      continue    
    }

    if (util.getInt(item.method) == 0) {
      util.error('수선방법을 선택하세요')
      return    
    }

    if (util.isNull(item.content)) {
      util.error('검토사유를 입력하세요')
      return    
    }

    if (util.isNull(item.adjust)) {
      util.error('조정내용을 입력하세요')
      return    
    }
    
    if (util.isNull(item.cycle)) {
      util.error('주기를 입력하세요')
      return    
    }

    if (util.isNull(item.percent)) {
      util.error('수선율을 입력하세요')
      return    
    }
    
    if (util.isNull(item.count)) {
      util.error('수량을 입력하세요')
      return    
    }

    if (util.isNull(item.pricestring)) {
      util.error('수선금액을 입력하세요')
      return    
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
     await Review.remove(item)
     }
     }
     }
   */

  let count = 0
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    if (util.getInt(item.category) == 0) {
      continue
    }
    
    item.apt = data.apt
    item.category = util.getInt(item.category)
    
    let category = getCategory(item.category)
    let subcategory = getCategory(category.parent)
    let topcategory = getCategory(subcategory.parent)

    item.subcategory = subcategory.id
    item.topcategory = topcategory.id

    item.count = util.getInt(item.count)
    item.price = util.getInt(item.pricestring)

    if (item.id > 0) {
      await Review.update(item)
    } else { 
      await Review.insert(item)
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

  for (let l = 0; l < data.batchs.length; l++) {
    index = l
    data.batchs[l].index = index    

    let topcategorys = data.batchs[l].categorys
    for (let i = 1; i < topcategorys.length; i++) {
      let topcategory = topcategorys[i]

      data.batchs[l].categorys[i].index = index 
      for (let j = 0; j < topcategory.children.length; j++) {
        let subcategory = topcategorys[i].children[j]

        data.batchs[l].categorys[i].children[j].index = index
        
        for (let k = 0; k < subcategory.children.length; k++) {
          data.batchs[l].categorys[i].children[j].children[k].index = index
        }
      }
    }    
  }
}

function clickAdd(count) {
  let items = []
  let rows = data.batchs.length
  for (let l = 0; l < count; l++) {
    let n = util.clone(item)
    let index = rows + l
    n.index = index
    
    let topcategorys = util.clone(data.categorys)
    for (let i = 1; i < topcategorys.length; i++) {
      let topcategory = topcategorys[i]

      topcategorys[i].index = index 
      for (let j = 0; j < topcategory.children.length; j++) {
        let subcategory = topcategorys[i].children[j]

        topcategorys[i].children[j].index = index
        
        for (let k = 0; k < subcategory.children.length; k++) {
          topcategorys[i].children[j].children[k].index = index
        }
      }
    }

    n.categorys = topcategorys
    
    items.push(n)
  }

  data.batchs = data.batchs.concat(items)
}

function clickNextInput(index) {
  if (index == data.batchs.length - 1) {
    clickSubmit()
  }
}

function changeContent(index, row) {
  let str = ''
  data.basics.forEach((item) => {
    if (row.content == item.content) {      
      str = item.adjust
    }
  })

  data.batchs[index].adjust = str
}

function changeCategory(item, node) {
  let category = getCategory(item.value)  
  if (category.level != 3) {
    return
  }
  
  let id = category.id
  let index = 0

  if (data.batchs.length > 1) {
    index = node.data.index
  }  

  let methods = [{id: 0, name: '수선방법'}]
  let standards = [{id: 0, name: '규격'}]
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      methods.push(item)
    }
  }

  console.log('---------------')
                console.log(index)
  console.log(data.batchs[index])
  console.log(methods)
  data.batchs[index].methods = methods

  let flag = false
  
  if (methods.length == 2) {
    flag = true
    data.batchs[index].method = methods[1].id
  } else { 
    data.batchs[index].method = null
  }

  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.category == id) {
      standards.push(item)
    }
  }

  data.batchs[index].standards = standards    
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

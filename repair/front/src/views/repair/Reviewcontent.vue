<template>
  <Title title="검토사항" />


  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="success" @click="clickBatch" style="margin-right:0px;">등록</el-button>
    </div>
  </div>  


  <div style="display:flex;space-between;gap:5px;">
    <div style="flex:1;">
      <el-table :data="[1]" border>
        <el-table-column label="수립일자" align="center">
          <template #default="scope">
            {{getDate(data.yearmonth)}}&nbsp;
          </template>
        </el-table-column>              
        
      </el-table>
    </div>

    <div style="flex:1;display:none;">
      <el-table :data="data.itemdates" border>
        <el-table-column label="조정이력" align="center">
          <template #default="scope">
            <div v-if="scope.$index == 0">{{data.repair.completionyear}}년 {{data.repair.completionmonth}}월</div>
            <div v-if="scope.$index > 0">
              <div v-if="data.itemdates[scope.$index].year == 0 || data.itemdates[scope.$index].month == 0">&nbsp;</div> 
              <div v-if="data.itemdates[scope.$index].year != 0 && data.itemdates[scope.$index].month != 0">{{data.itemdates[scope.$index].year}}년 {{data.itemdates[scope.$index].month}}월</div>
            </div>
          </template>
        </el-table-column>              
        
      </el-table>
    </div>

    <div style="flex:4;">
      <el-table :data="data.items" border>
        <el-table-column prop="name" label="내용" >
          <template #default="scope">
            <span v-html="util.nl2br(scope.row.content)" />&nbsp;
          </template>
        </el-table-column>

      </el-table>  
    </div>
  </div>
  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="800px"
  >

    
    

    <div style="display:flex;space-between;gap:5px;">
      <div style="flex:1;">
        <el-table :data="[1]" border style="margin-top:15px;">
          <el-table-column label="수립일자" align="center">
            <template #default="scope">
              <el-date-picker style="margin: 0px 0px;height: 24px;width:120px;" v-model="data.batchyearmonth" type="month" placeholder="년월" />
            </template>
          </el-table-column>              
          
        </el-table>
      </div>

      <div style="flex:1;display:none;">
        <el-table :data="data.batchdates" border style="margin-top:15px;">
          <el-table-column label="조정이력" align="center">
            <template #default="scope">
              <div v-if="scope.$index == 0">{{data.repair.completionyear}}-{{data.repair.completionmonth}}</div>
              <el-date-picker v-if="scope.$index > 0" style="margin: 0px 0px;height: 24px;width:120px;" v-model="data.batchdates[scope.$index].yearmonth" type="month" placeholder="년월" />
            </template>
          </el-table-column>              
          
        </el-table>
      </div>

      <div style="flex:4;">
        <el-table :data="data.batchs" border style="margin-top:15px;">
          <el-table-column label="검토내용" align="center">
            <template #default="scope">
              <el-select v-model="data.batchs[scope.$index].content" placeholder="" style="width:100%;"
                         filterable
                         allow-create
                         default-first-option
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
          
        </el-table>

      </div>

    </div>  

      <template #footer>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Repair, Reviewcontent, Reviewcontentbasic, Reviewdate } from "~/models"
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
  content: '',
  order: 0
}

const itemdate = {
  id: 0,
  year: null,
  month: null,
  order: 0
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  itemdates: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  itemdate: util.clone(itemdate),
  visible: false,    
  search: {
    text: ''
  },
  basics: [],
  batchs: [],
  batchdates: [],
  batchyearmonth: null,
  dates: [],
  yearmonth: null,
  repair: {
    completionyear: '',
    completionmonth: ''
  }
})

async function initData() {
  let res = await Reviewcontentbasic.find({
    orderby: 'rb_order,rb_id'
  })

  if (res.items == null) {
    res.items = []
  }

  let res3 = await Repair.get(data.apt)
  let repair = res3.item
  
  if (repair.type == 2) {
    let res2 = await Repair.find({
      apt: repair.apt,
      orderby: 'r_reportdate,r_id'
    })

    let years = res2.items.filter(item => item.reportdate != '' && item.reportdate < repair.reportdate).map(item => util.getYear(item.reportdate) + '년').join(', ')
    res.items.push({id: 0, content: years + '에 진행(보수)한 부분에 대하여 반영하고,'})

  }

  data.basics = [{id: 0, content: ''}, ...res.items]  
}

async function getItems() {
  let res = await Reviewcontent.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    content: data.search.text,
    orderby: 'rc_id'
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


  res = await Reviewdate.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'rd_id'
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      let item = res.items[i]
      res.items[i].index = i + 1
      if (item.year == 0 && item.month == 0) {
        res.items[i].yearmonth = null
      } else {
        res.items[i].yearmonth = `${item.year}-${item.month}`
      }
    }
  }

  if (res.items == null) {
    res.items = []
  }
  data.itemdates = res.items

  res = await Repair.get(data.apt)
  data.repair = res.item

  data.yearmonth = data.repair.reviewcontent1
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
    let res = await Reviewcontent.remove(item)
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

function clickBatch() {
  let items = util.clone(data.items)

  if (items == null) {
    items = []
  }

  if (items.length < 12) {
    let length = 12 - items.length
    for (let i = 0; i < length; i++) {
      items.push(util.clone(data.item))
    }
  }

  items = makeItems(items)

  data.batchs = items

  items = util.clone(data.itemdates)

  if (items == null) {
    items = []
  }

  if (items.length < 7) {
    let length = 7 - items.length
    for (let i = 0; i < length; i++) {
      items.push(util.clone(data.itemdate))
    }
  }

  data.batchdates = items  

  data.batchyearmonth = data.yearmonth
  
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

async function clickSubmit() {
  /*
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    
    if (item.content == '') {
      util.error('내용을 입력하세요')
      return    
    }
  }
  */
  
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
        await Reviewcontent.remove(item)
      }
    }

    for (let i = 0; i < data.itemdates.length; i++) {
      let item = data.itemdates[i]
      let flag = false;
      for (let j = 0; j < data.batchdates.length; j++) {
        if (data.itemdates[i].id == data.batchdates[j].id) {
          flag = true
          break
        }
      }

      if (flag == false) {      
        await Reviewdate.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.order = util.getInt(item.order)

    if (item.id > 0) {
      await Reviewcontent.update(item)
    } else { 
      await Reviewcontent.insert(item)
    }
  }

  for (let i = 0; i < data.batchdates.length; i++) {
    let item = data.batchdates[i]

    item.apt = data.apt
    item.order = util.getInt(item.order)

    if (typeof item.yearmonth == 'string') {
      let strs = item.yearmonth.split('-')

      if (strs.length == 2) {
        item.year = util.getInt(strs[0])
        item.month = util.getInt(strs[1])
      } else {
        item.year = 0
        item.month = 0
      }
    } else if (item.yearmonth == null) {
        item.year = 0
        item.month = 0
    } else {
      item.year = item.yearmonth.getFullYear()
      item.month = item.yearmonth.getMonth() + 1
    }
    
    if (item.id > 0) {
      await Reviewdate.update(item)
    } else { 
      await Reviewdate.insert(item)
    }
  }

  let res = await Repair.get(data.apt)
  let repair = res.item


  let yearmonth = data.batchyearmonth
  let year = 0
  let month = 0
  if (typeof yearmonth == 'string') {
    let strs = yearmonth.split('-')

    if (strs.length == 2) {
      year = util.getInt(strs[0])
      month = util.getInt(strs[1])
    } else {
      year = 0
      month = 0
    }
  } else if (yearmonth == null) {
    year = 0
    month = 0
  } else {
    year = yearmonth.getFullYear()
    month = yearmonth.getMonth() + 1
  }
  
  repair.reviewcontent1 = `${year}-${month}`
  await Repair.update(repair)

  util.info('등록되었습니다')
  getItems()
  data.visible = false
  
  util.loading(false)  
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function clickAdd(count) {
  if (data.batchs.length >= 12) {
    return
  }
  
  let items = []
  for (let i = 0; i < count; i++) {
    items.push(util.clone(item))
  }

  data.batchs = data.batchs.concat(items)
}

function getDate(str) {
  if (str == null) {
    return ''
  }
  
  if (str == '') {
    return ''
  }

  let strs = str.split('-')

  if (strs.length != 2) {
    return ''
  }

  return `${strs[0]}년 ${strs[1]}월`
}

</script>

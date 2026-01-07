<template>
  <Title title="충당금적립" />


  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="success" @click="clickBatch" style="margin-right:0px;">등록</el-button>
    </div>
  </div>  


  <el-table :data="data.items" border :summary-method="getSummaries" show-summary>
    <el-table-column prop="year" label="년도" align="center" width="100" />
    
    <el-table-column prop="name" label="전기이월" align="right">
      <template #default="scope">
        {{util.money(scope.row.forward)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="적립액" align="right">
      <template #default="scope">
        {{util.money(scope.row.saving)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="이자"  align="right">
      <template #default="scope">
        {{util.money(scope.row.interest)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="이익잉여금" align="right">
      <template #default="scope">
        {{util.money(scope.row.surplus)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="기타" align="right">
      <template #default="scope">
        {{util.money(scope.row.etc)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="사용액" align="right">
      <template #default="scope">
        {{util.money(scope.row.use)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="계" align="right">
      <template #default="scope">
        {{scope.row.total}}
      </template>
    </el-table-column>


    <el-table-column prop="remark" label="비고" />
      
  </el-table>  


  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1000px"
  >

    
    
    <el-table :data="data.batchs" border style="margin-top:15px;">
      <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
      </el-table-column>
      <el-table-column prop="year" label="년도" align="center" width="100">
        <template #default="scope">
          <el-date-picker v-if="scope.$index == 0" style="margin: 0px 0px;height: 24px;width:90px;" v-model="data.batchs[scope.$index].year" type="year" placeholder="년도" @change="changeYear" />
          <div v-if="scope.$index != 0">{{data.batchs[scope.$index].year}}</div>
        </template>
      </el-table-column>
      
        <el-table-column prop="name" label="전기이월" align="right">
          <template #default="scope">
            <el-input v-if="scope.$index == 0" v-model="data.batchs[scope.$index].forward" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')" :parser="(value) => value.replace(/\$\s?|(,*)/g, '')" class="inputNumber" @keyup="onKeyup(scope.$index)" />
            <div v-if="scope.$index != 0">{{util.money(data.batchs[scope.$index].forward)}}</div>
          </template>
        </el-table-column>

        <el-table-column prop="name" label="적립액" align="right">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].saving" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')" :parser="(value) => value.replace(/\$\s?|(,*)/g, '')" class="inputNumber"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>

        <el-table-column prop="name" label="이자"  align="right">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].interest" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')" :parser="(value) => value.replace(/\$\s?|(,*)/g, '')" class="inputNumber"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>

        <el-table-column prop="name" label="이익잉여금" align="right">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].surplus" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')" :parser="(value) => value.replace(/\$\s?|(,*)/g, '')" class="inputNumber"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>

        <el-table-column prop="name" label="기타" align="right">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].etc" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')" :parser="(value) => value.replace(/\$\s?|(,*)/g, '')" class="inputNumber"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>


        <el-table-column prop="name" label="사용액" align="right">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].use" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')" :parser="(value) => value.replace(/\$\s?|(,*)/g, '')" class="inputNumber"  @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>

        <el-table-column prop="name" label="계" align="right">
          <template #default="scope">
            {{data.batchs[scope.$index].total}}
          </template>
        </el-table-column>

        <el-table-column prop="remark" label="비고" width="150">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].remark" />
          </template>
        </el-table-column>

        
      </el-table>

      
      <template #footer>
        <el-button size="small" type="danger" v-if="data.mode != 'batch' && (data.batchs.length > 0 && data.batchs[0].id > 0)" style="float:left;" @click="clickDelete">삭제</el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(1)"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" type="warning" style="margin-right:130px;" @click="clickUsage">사용액 가져오기</el-button>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Saving, History } from "~/models"
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
  year: 0,
  forward: 0,
  interest: 0,
  surplus: 0,
  etc: 0,
  saving: 0,
  use: 0,
  remark: ''
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
  batchs: []    
})

async function initData() {    
}

async function getItems() {
  let res = await Saving.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'sa_year,sa_id'
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      let item = res.items[i]
      res.items[i].index = i + 1

      let total = util.getInt(item.forward) + util.getInt(item.interest) + util.getInt(item.surplus) + util.getInt(item.etc) + util.getInt(item.saving) - util.getInt(item.use)
      res.items[i].total = util.money(total)
    }
  }

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }
  data.items = res.items  
}

function makeItems(items) {
  for (let i = 0; i < items.length; i++) {
    let item = items[i]
    let total = util.getInt(item.forward) + util.getInt(item.interest) + util.getInt(item.surplus) + util.getInt(item.etc) + util.getInt(item.saving) - util.getInt(item.use)
    items[i].total = util.money(total)
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
    let res = await Saving.remove(item)
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

  if (items.length == 0) {
    let length = 15 - items.length
    for (let i = 0; i < length; i++) {
      let n = util.clone(data.item)
      n.index = i + 1
      items.push(n)
    }
  } else {
    let d = new Date()
    d.setFullYear(items[0].year)
    items[0].year = d
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
        await Saving.remove(item)
      }
    }
  }

  for (let i = 0; i < data.batchs.length; i++) {
    let item = util.clone(data.batchs[i])

    item.apt = data.apt

    if (i == 0) {
      let year = data.batchs[0].year
      if (typeof year == 'string' || typeof year == 'number') {
        item.year = util.getInt(year)
      } else if (year == null || year == undefined || year == 0) {
        util.error('년월을 입력하세요')
        return
      } else {
        item.year = year.getFullYear()
      }
    } else {
      item.year = util.getInt(item.year)
    }

    item.forward = util.getInt(item.forward)
    item.interest = util.getInt(item.interest)
    item.surplus = util.getInt(item.surplus)
    item.etc = util.getInt(item.etc)
    item.saving = util.getInt(item.saving)
    item.use = util.getInt(item.use)

    if (item.id > 0) {
      await Saving.update(item)
    } else { 
      await Saving.insert(item)
    }
  }

  util.info('등록되었습니다')
  getItems()
  data.visible = false
  
  util.loading(false)  
}

function calculate() {
  let item = data.batchs[0].year

  let year
  if (typeof item === 'string') {
    year = util.getInt(item)
  } else if (typeof item === 'number') {
    year = item
  } else if (item == null || item == undefined) {
    year = 0
  } else {
    year = item.getFullYear()
  }

  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    if (i == 1) {
      data.batchs[i].forward = util.getInt(data.batchs[i - 1].total)
      data.batchs[i].year = year + 1
    } else if (i > 1) {
      data.batchs[i].forward = util.getInt(data.batchs[i - 1].total)
      data.batchs[i].year = data.batchs[i - 1].year + 1
    }
    
    let total = util.getInt(item.forward) + util.getInt(item.interest) + util.getInt(item.surplus) + util.getInt(item.etc) + util.getInt(item.saving) - util.getInt(item.use)
    data.batchs[i].total = util.money(total)    
  }  
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)

  calculate()
}

function clickAdd(count) {
  let items = []
  for (let i = 0; i < count; i++) {
    items.push(util.clone(item))
  }

  data.batchs = data.batchs.concat(items)

  calculate()
}

function getSum(item) {
}

function onKeyup(index) {
  calculate()
}

function getYear(index) {
  let year = util.getInt(data.batchs[0].year)

  if (year > 0) {
    return year + index + 1
  }

  return ''
}

function changeYear(item) {
  let year
  if (typeof item === 'string') {
    year = util.getInt(item)
  } else if (item == null || item == undefined) {
    year = 0
  } else {
    year = item.getFullYear()
  }

  if (year == 0) {
    for (let i = 1; i < data.batchs.length; i++) {
      data.batchs[i].year = ''  
    }
  } else { 
    for (let i = 1; i < data.batchs.length; i++) {
      data.batchs[i].year = year + i  
    }
  }
}

const getSummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index === 0) {
      sums[index] = '합계'
    } else if (index >= 2 && index <= 7) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          if (index == 2) {
            total += item.saving
          } else if (index == 3) {
            total += item.interest
          } else if (index == 4) {
            total += item.surplus
          } else if (index == 5) {
            total += item.etc
          } else if (index == 6) {
            total += item.use
          } else if (index == 7) {
            total += util.getInt(item.total)
          }
        })
      }
      
      sums[index] = util.money(total)
    }
  })

  return sums
}

function clickUsage() {
  util.confirm('사용액을 자동으로 가져오시겠습니까', async function() {
    util.loading(true) 

    const res = await History.find({
      apt: data.apt
    })

    let years = {}
    let items = res.items
    for (let i = 0; i < items.length; i++) {
      const item = items[i]

      if (years[item.year] == undefined) {
        years[item.year] = 0
      }

      years[item.year] += item.price
    }

    items = util.clone(data.batchs)

    for (let i = 0; i < items.length; i++) {
      let item = items[i]

      if (years[item.year] == undefined) {
        items[i].use = 0
        continue
      }

      items[i].use = years[item.year]
    }

    data.batchs = items
    util.loading(false)
  })
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

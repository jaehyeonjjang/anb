<template>
  <div v-for="item in data.items">
    <PageTitle :title="`${getCategory(item.topcategory).name} > ${getCategory(item.subcategory).name}`" />

    <el-table :data="item.datas" border :span-method="spanMethod" :cell-class-name="cellClassName">
      <el-table-column prop="index" label="순번" align="center" width="40">
        <template #default="scope">
          {{scope.row.index}}
        </template>
      </el-table-column>
      <el-table-column label="소분류">
        <template #default="scope">
          {{getCategory(scope.row.category).name}}
        </template>
      </el-table-column>
      <el-table-column label="규격">
        <template #default="scope">
          <div v-if="scope.row.type == 1">{{scope.row.extra.standard.name}}</div>
          <div v-if="scope.row.type == 2">{{scope.row.date}}</div>
          <div v-if="scope.row.type == 3" style="text-align:center;">{{scope.row.title}}</div>
        </template>
      </el-table-column>
      <el-table-column label="수선방법" width="70" align="center">
        <template #default="scope">
          <div v-if="scope.row.type == 1">{{scope.row.extra.category.name}}</div>
          <div v-if="scope.row.type == 2">{{scope.row.content}}</div>
        </template>
      </el-table-column>
      <el-table-column prop="extra.category.cycle" label="수선주기" align="center" width="60" />
      <el-table-column prop="extra.category.percent" label="수선율" align="center" width="50" />
      <el-table-column prop="extra.standard.unit" label="단위" align="center" width="50" />
      <el-table-column prop="count" label="수량" align="right" width="50" />
      <el-table-column label="재료비" align="right" width="80">
        <template #default="scope">
          {{util.money(scope.row.extra.standard.direct)}}
        </template>
      </el-table-column>
      <el-table-column label="노무비" align="right" width="70">
        <template #default="scope">
          {{util.money(scope.row.extra.standard.labor)}}
        </template>
      </el-table-column>
      <el-table-column label="경비" align="right" width="70">
        <template #default="scope">
          {{util.money(scope.row.extra.standard.cost)}}
        </template>
      </el-table-column>      
      <el-table-column label="단가" align="right" width="80">
        <template #default="scope">
          {{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}
        </template>
      </el-table-column>

      <el-table-column label="수선금액" align="right" width="90">
        <template #default="scope">
          <div v-if="scope.row.type != 3">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</div>
          <div v-if="scope.row.type == 3">{{util.money(scope.row.price)}}</div>
        </template>
      </el-table-column>

      <el-table-column label="연평균 적립금액" align="right" width="90">
        <template #default="scope">
          <div v-if="scope.row.type != 3">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent) / util.getFloat(scope.row.extra.category.cycle))}}</div>
          <div v-if="scope.row.type == 3">{{util.money(scope.row.savingprice)}}</div>
        </template>
      </el-table-column>
      
      <el-table-column prop="lastdate" label="최종수선" align="center" width="55" />
      <el-table-column prop="duedate" label="수선예정" align="center" width="55" />
      

    </el-table>  


  </div>


  
  
</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Repair, Category, Dong, Standard, History, Breakdown, Totalreport } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import type { TabsPaneContext } from 'element-plus'

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
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,    
  search: '',
  dongs: [],
  categorys: [],
  standards: [],
  width: 0,
  height: 0,
  historys: [],
  categoryMap: {},
  parcelrate: 0
})

async function initData() {
  let res = await Dong.findByApt(data.apt)
  data.dongs = res.items

  res = await Standard.findByApt(data.apt)
  data.standards = res.items
  
  res = await Category.findByApt(data.apt)
  data.categorys = res.items

  let categoryMap = {}
  res.items.forEach((item) => {
    categoryMap[item.id] = item
  })

  data.categoryMap = categoryMap
}

function getHistory(category) {
  let items = data.historys.filter(item => item.category == category)
  return items
}  

function getHistoryItem(history) {
  if (history == null) {
    history = {
      year: '',
      month: '',
      content: ''
    }
  }
  
  let item = {
    category: 0,
    extra: {
      standard: {
        name: '',
        direct: 0,
        labor: 0,
        cost: 0
      },
      category: {
        name: '',
        cycle: 0,
        percent: 0,
        unit: ''
      }
    },
    count: 0,
    lastdate: '',
    duedate: '',
    type: 2,
    year: history.year,
    month: history.month,
    content: history.content,
    date: `${history.year}년 ${history.month}월`,
  span: 0
  }

  return item
}


function getTotalItem(title, price, savingprice) {
  let item = {
    category: 0,
    extra: {
      standard: {
        name: '',
        direct: 0,
        labor: 0,
        cost: 0
      },
      category: {
        name: '',
        cycle: 0,
        percent: 0,
        unit: ''
      }
    },
    count: 0,
    lastdate: '',
    duedate: '',
    type: 3,
    year: 0,
    month: 0,
    content: '',
    date: '',
    span: 0,
    title: title,
    price: price,
    savingprice: savingprice
  }

  return item
}

async function getItems() {
  if (data.apt == 0) {
    return
  }

  let res = await Repair.get(data.apt)
  //data.parcelrate = res.item.parcelrate
  
  res = await History.find({
    page: 0,
    pagesize: 0,
    apt: data.apt,
    orderby: 'h_year,h_month,c_order,h_id'
  })
  
  if (res.items == null) {
    res.items = []
  }    

  data.historys = res.items

  res = await Totalreport.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'b_topcategory,b_subcategory,b_category,b_standard,b_method'
  })
  
  data.total = res.total
  if (res.items == null) {
    res.items = []
  }

  let items = []
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    let flag = false
    for (let j = 0; j < items.length; j++) {
      let subcategory = items[j]

      if (subcategory.subcategory == item.subcategory) {
        flag = true
        break
      }
    }

    if (flag == false) {
      items.push(item)
    }
  }

  
  
  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    items[i].datas = []
    
    let index = 1;
    let last = -1
    let lastFlag = true

    let partPrice = 0
    let partTotalprice = 0
    let allPrice = 0
    let allTotalprice = 0
    
    for (let j = 0; j < res.items.length; j++) {
      let d = res.items[j]

      if (item.subcategory != d.subcategory) {
        continue
      }

      if (last != -1) {
        if (d.category != last) {
          let historys = getHistory(last)

          if (historys.length > 0) {
            let header = getHistoryItem(null)
            header.data = i
            header.span = historys.length + 1
            header.index = '보수이력'
            header.date = '날짜'
            header.content = '상세 내역'

            items[i].datas.push(header)

            for (let j = 0; j < historys.length; j++) {
              let history = historys[j]
              history.data = i
              history.index = ''
              items[i].datas.push(getHistoryItem(history))
            }
          }
        }
      }

      d.type = 1
      d.data = i
      d.index = index
      index++

      items[i].datas.push(d)
      
      lastFlag = false      
      last = d.category

      let price = parseInt(util.calculateRepair(d.extra.standard.direct, d.extra.standard.labor, d.extra.standard.cost, d.rate, data.parcelrate, d.count, d.extra.category.percent))
      let totalprice = parseInt(util.calculateRepair(d.extra.standard.direct, d.extra.standard.labor, d.extra.standard.cost, d.rate, data.parcelrate, d.count, d.extra.category.percent) / util.getFloat(d.extra.category.cycle))
      
      if (d.extra.category.percent == 100) {
        allPrice += price
        allTotalprice += totalprice
      } else {
        partPrice += price
        partTotalprice += totalprice
      }
    }

    if (lastFlag == false) {
      let historys = getHistory(last)

      if (historys.length > 0) {
        let header = getHistoryItem(null)
        header.data = i
        header.span = historys.length + 1
        header.index = '보수이력'
        header.date = '날짜'
        header.content = '상세 내역'

        items[i].datas.push(header)

        for (let j = 0; j < historys.length; j++) {
          let history = historys[j]
          history.data = i
          items[i].datas.push(getHistoryItem(history))
        }
      }
    }

    items[i].datas.push(getTotalItem('(부분수리)합계', partPrice, partTotalprice))
    items[i].datas.push(getTotalItem('(전면수리)합계', allPrice, allTotalprice))
    items[i].datas.push(getTotalItem('합계', partPrice+allPrice, partTotalprice + allTotalprice))
  }

  data.items = items  
}

function setWindowSize() {
  data.width = (window.innerWidth - 500) + 'px'
  data.height = (window.innerHeight - 170) + 'px'
}

onMounted(() => {
  data.apt = parseInt(route.params.id)  
  /*
     
     
     initData()
     getItems()
   */
  setWindowSize()

  window.addEventListener('resize', setWindowSize)
})

async function readData() {
  util.loading(true)
  
  await initData()
  await getItems()
  
  util.loading(false)
}

defineExpose({
  readData
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

function getStandard(id) {
  for (let i = 0; i < data.standards.length; i++) {
    let item = data.standards[i]

    if (item.id == id) {
      return item
    }
  }

  return {
    id: 0,
    name: ''
  }
}

const spanMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  
  if (row.type == 2) {
    if (columnIndex === 0) {
      return {rowspan: row.span, colspan: 2}
    } else if (columnIndex === 1) {
      return {rowspan: 0, colspan: 0}
    } else if (columnIndex === 2) {
      return {rowspan: 1, colspan: 1}      
    } else if (columnIndex === 3) {
      return {rowspan: 1, colspan: 12}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  } else if (row.type == 3) {
    if (columnIndex === 0) {
      return {rowspan: 1, colspan: 1}
    } else if (columnIndex === 1) {
      return {rowspan: 1, colspan: 1}
    } else if (columnIndex === 2) {
      return {rowspan: 1, colspan: 10}      
    } else if (columnIndex === 12) {
      return {rowspan: 1, colspan: 1}
    } else if (columnIndex === 13) {
      return {rowspan: 1, colspan: 1}
    } else if (columnIndex === 14) {
      return {rowspan: 1, colspan: 1}
    } else if (columnIndex === 15) {
      return {rowspan: 1, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }
  return {rowspan: 1, colspan: 1}
}


function cellClassName({row, columnIndex}) {
  if (row.type == 2) {
    return 'title'
  } else {
    return 'value'    
  }
}

</script>
<style>
.title {
  background-color: #fafafa;
}

.value {
  background-color: #FFF;  
}
</style>  

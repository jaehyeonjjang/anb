<template>

  <PageTitle title="장기수선충담금 계획금액" />
  
  <el-table :data="data.totals" border :cell-class-name="cellClassNameTotal" :span-method="spanMethodTotal">
    <el-table-column prop="index" label="구분" align="center" width="50" />
    <el-table-column prop="title" label="내용" />
    <el-table-column prop="content1" label="" align="center" />
    <el-table-column prop="content2" label="계산식" align="center" width="50" />
    <el-table-column prop="content3" label="" align="center" />
    <el-table-column prop="price" label="금액" align="center" />
  </el-table>



  <PageTitle title="장기수선충당금 계획금액 (* 50년 가상 금액)" />
  
  <el-table :data="data.totalplans" border :cell-class-name="cellClassNameTotal" :span-method="spanMethodTotalPlan">
    <el-table-column prop="index" label="구분" align="center" width="50" />
    <el-table-column prop="title" label="내용" />
    <el-table-column prop="content1" label="" align="center" />
    <el-table-column prop="content2" label="계산식" align="center" width="50" />
    <el-table-column prop="content3" label="" align="center" />
    <el-table-column prop="price" label="금액" align="center" />
  </el-table>

  <PageTitle title="장기수선충당금 계획금액과 장기수선충당금 적립금액의 비교" />
  
  <el-table :data="data.totalcompares" border >
    <el-table-column prop="index" label="구분" align="center" width="50" />
    <el-table-column prop="title" label="" />
    <el-table-column prop="price" label="장기수선계획금액" align="right" />
    <el-table-column prop="rate" label="비율" align="center" width="60">
      <template #default="scope">
        {{util.fixed(100.0 * util.getFloat(scope.row.price2) / util.getFloat(scope.row.price), 2)}}
      </template>
    </el-table-column>
    <el-table-column prop="price2" label="적립예정금액" align="right" />
    <el-table-column prop="title2" label="" />
  </el-table>


  <PageTitle title="수립 예정단가로 본 적립 요율 적용(관리규약 요율 및 단가)" />


  <el-table :data="data.items" border>          
    <el-table-column label="시설물의 내구 연한" align="center" width="140">
      <template #default="scope">
        <div style="display:flex;font-size:12px;">
          <div style="margin-left:7px;width:50px;">{{scope.row.startyear}}.{{util.pad(scope.row.startmonth, 2)}}</div> <div>~</div> <div style="margin-left:5px;"> {{scope.row.endyear}}.{{util.pad(scope.row.endmonth, 2)}}</div>
        </div>              
      </template>
    </el-table-column>

    <el-table-column label="구간별 적립금액" align="right">
      <template #default="scope">
        <span style="font-size:12px;">{{scope.row.totalprice}}</span>
      </template>
    </el-table-column>

    <el-table-column label="적용 적립요율" align="center">
      <template #default="scope">
        {{scope.row.rate}}
      </template>
    </el-table-column>

    <el-table-column label="m2당 단가" align="center">
      <template #default="scope">
        {{scope.row.price}}              
      </template>
    </el-table-column>
    
    <el-table-column label="누적요율" align="center">
      <template #default="scope">
        {{scope.row.totalrate}}
      </template>
    </el-table-column>

    <el-table-column label="비고">
      <template #default="scope">
        {{scope.row.remark}}
      </template>
    </el-table-column>
    
  </el-table>



  <PageTitle title="수립 예정단가로 본 적립 요율 적용(향후 개정시 적용해야할 요율 및 단가)" />

  <el-table :data="data.itemplans" border>          
    <el-table-column label="시설물의 내구 연한" align="center" width="140">
      <template #default="scope">
        <div style="display:flex;font-size:12px;">
          <div style="margin-left:7px;width:50px;">{{scope.row.startyear}}.{{util.pad(scope.row.startmonth, 2)}}</div> <div>~</div> <div style="margin-left:5px;"> {{scope.row.endyear}}.{{util.pad(scope.row.endmonth, 2)}}</div>
        </div>              
      </template>
    </el-table-column>

    <el-table-column label="구간별 적립금액" align="right">
      <template #default="scope">
        <span style="font-size:12px;">{{scope.row.totalprice}}</span>
      </template>
    </el-table-column>

    <el-table-column label="적용 적립요율" align="center">
      <template #default="scope">
        {{scope.row.rate}}
      </template>
    </el-table-column>

    <el-table-column label="m2당 단가" align="center">
      <template #default="scope">
        {{scope.row.price}}              
      </template>
    </el-table-column>
    
    <el-table-column label="누적요율" align="center">
      <template #default="scope">
        {{scope.row.totalrate}}
      </template>
    </el-table-column>

    <el-table-column label="비고">
      <template #default="scope">
        {{scope.row.remark}}
      </template>
    </el-table-column>
    
  </el-table>

</template>

<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Report, Repair, Outline, Outlineplan } from "~/models"
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
  totalprice: 0,
  rate: 0,
  price: 0,
  totalrate: 0,
  remark: ''
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  itemplans: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  itemplan: util.clone(item),
  visible: false,    
  search: {
    text: ''
  },
  basics: [],
  batchs: [],
  batchplans: [],
  repair: {
    completionyear: '',
    completionmonth: ''
  },
  batchrepair: {
    savingprice: '',
    price1: '',
    price2: '',
    price3: '',
    price4: '',
    price2: ''
  },
  report: {
    price: 0,
    totalprice: 0,
    totalsize: 0
  },
  totals: [],
  totalplans: [],
  totalcompares: []
})

async function initData() {  
}

async function getItems() {
  if (data.apt == 0) {
    return
  }
  
  let res = await Repair.get(data.apt)
  data.repair = res.item
  data.parcelrate = res.item.parcelrate
  if (data.parcelrate == 0) {
    data.parcelrate = 100
  }

  res = await Report.total(data.apt)

  data.report = res
  
  res = await Outline.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'o_id'
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

  let alltotalprice = 0
  let sum = 0
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    sum += util.getFloat(item.rate)

    let per = parseFloat(util.getFloat(item.rate))
    let totalprice = util.getInt((util.getFloat(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)

    res.items[i].totalprice = util.money(totalprice)
    res.items[i].totalrate = util.getFloat(sum)

    let duration = getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth)    
    res.items[i].duration = duration

    //let price = parseFloat(totalprice) / (util.getFloat(data.report.totalsize) * util.getFloat(duration))
    //res.items[i].price = util.moneyfloat(res.items[i].price)

    if (i < res.items.length - 1) {
      alltotalprice += totalprice
    }
  }

  if (res.items.length > 0 && res.items[res.items.length - 1].totalrate === 100.0) {    
    res.items[res.items.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - alltotalprice)
  }
  
  data.items = res.items

  res = await Outlineplan.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'op_id'
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }
  }
  
  if (res.items == null) {
    res.items = []
  }

  alltotalprice = 0
  sum = 0
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    sum += util.getFloat(item.rate)

    let per = parseFloat(util.getFloat(item.rate))
    let totalprice = util.getInt((util.getFloat(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)

    res.items[i].totalprice = util.money(totalprice)
    res.items[i].totalrate = util.getFloat(sum)

    let duration = getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth)    
    res.items[i].duration = duration

    //let price = parseFloat(totalprice) / (util.getFloat(data.report.totalsize) * util.getFloat(duration))
    //res.items[i].price = util.moneyfloat(res.items[i].price)

    if (i < res.items.length - 1) {
      alltotalprice += totalprice
    }
  }

  if (res.items.length > 0 && res.items[res.items.length - 1].totalrate === 100.0) {    
    res.items[res.items.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - alltotalprice)
  }
  
  data.itemplans = res.items

  let plan = await Report.plan(data.apt)

  let years = parseFloat(util.getPlanyears(data.repair.planyears) * 12)
  
  let items = [
    {
      index: 1,
      title: '총 계획금액',
      content1: plan.startyear,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(plan.total),
      remark: '사용검사년도는 제외'
    },
    {
      index: '',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.total) / years * 12)),
      content2: '*',
      content3: util.getPlanyears(data.repair.planyears),
      price: 0,
      remark: ''
    },
    {
      index: '2',
      title: '기준년도까지 적립금액',
      content1: plan.startyear,
      content2: '~',
      content3: plan.reportyear,
      price: util.money(Math.round(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1) * util.getFloat(plan.total) / years * 12),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.total) / years * 12)),
      content2: '*',
      content3: util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1,
      price: 0,
      remark: ''
    },
    {
      index: '3',
      title: '기준년도 이후 적립금액',
      content1: util.getInt(plan.reportyear) + 1,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(Math.round(util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1)) * util.getFloat(plan.total) / years * 12),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.total) / years * 12)),
      content2: '*',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1),
      price: 0,
      remark: ''
    },
    {
      index: '4',
      title: '년 평균 계획금액',
      content1: '월 평균 계획금액',
      content2: '*',
      content3: '12개월',
      price: util.money(Math.round(util.getFloat(plan.total) / years * 12)),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.total) / years)),
      content2: '*',
      content3: 12,
      price: 0,
      remark: ''
    },
    {
      index: '5',
      title: '월 평균 계획금액',
      content1: '총 계획금액',
      content2: '/',
      content3: '적립개월수',
      price: util.money(Math.round(util.getFloat(plan.total) / years)),
      remark: '사용검사년도는 제외'
    },
    {
      index: '2',
      title: '',
      content1: util.money(plan.total),
      content2: '/',
      content3: util.getInt(years),
      price: 0,
      remark: ''
    },
    {
      index: '6',
      title: 'm2 평균 계획단가',
      content1: '월 평균 계획금액',
      content2: '/',
      content3: '총 주택공급 면적',
      price: Math.round(util.getFloat(plan.total) / years / plan.area),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.total) / years)),
      content2: '/',
      content3: util.fixed(util.getFloat(plan.area), 3),
      price: 0,
      remark: ''
    },
  ]

  data.totals = items
  
  let itemplans = [
    {
      index: 1,
      title: '총 계획금액',
      content1: plan.startyear,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12 * util.getPlanyears(data.repair.planyears))),
      remark: '사용검사년도는 제외'
    },
    {
      index: '',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12)),
      content2: '*',
      content3: util.getPlanyears(data.repair.planyears),
      price: 0,
      remark: ''
    },
    {
      index: '2',
      title: '기준년도까지 적립금액',
      content1: plan.startyear,
      content2: '~',
      content3: plan.reportyear,
      price: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12 * util.getFloat(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1))),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12)),
      content2: '*',
      content3: util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1,
      price: 0,
      remark: ''
    },
    {
      index: '3',
      title: '기준년도 이후 적립금액',
      content1: util.getInt(plan.reportyear) + 1,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12 * util.getFloat(util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1)))),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12)),
      content2: '*',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1),
      price: 0,
      remark: ''
    },
    {
      index: '4',
      title: '년 평균 적립금액',
      content1: '월 평균 계획금액',
      content2: '*',
      content3: '12개월',
      price: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12)),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice))),
      content2: '*',
      content3: 12,
      price: 0,
      remark: ''
    },
    {
      index: '5',
      title: '월 평균 적립금액',
      content1: 'm2 평균 적립금액',
      content2: '/',
      content3: '총 부과면적',
      price: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice))),
      remark: '사용검사년도는 제외'
    },
    {
      index: '2',
      title: '',
      content1: util.money(data.repair.savingprice),
      content2: '/',
      content3: util.fixed(util.getFloat(plan.area), 3),
      price: 0,
      remark: ''
    },
    {
      index: '6',
      title: 'm2 평균 적립금액',
      content1: '기준일 현재 부과금액',
      content2: '/',
      content3: '총 주택공급 면적',
      price: util.money(data.repair.savingprice),
      remark: ''
    }    
  ]

  data.totalplans = itemplans

  let compares = [
    {
      index: 1,
      title: '총 계획금액',
      title2: '총 적립금액',
      price: util.money(plan.total * data.parcelrate / 100),
      price2: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12 * util.getPlanyears(data.repair.planyears) * data.parcelrate / 100))
    },
    {
      index: 2,
      title: '기준일까지 계획금액',
      title2: '기준일까지 적립금액',
      price: util.money(Math.round((util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1) * util.getFloat(plan.total) / years * 12 * data.parcelrate / 100)),
      price2: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12 * util.getFloat(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1) * data.parcelrate / 100))
    },
    {
      index: 3,
      title: '년 평균 계획금액',
      title2: '년 평균 적립금액',
      price: util.money(Math.round(util.getFloat(plan.total) / years * 12 * data.parcelrate / 100)),
      price2: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * 12 * data.parcelrate / 100))
    },
    {
      index: 4,
      title: '월 평균 계획금액',
      title2: '월 평균 적립금액',
      price: util.money(Math.round(util.getFloat(plan.total) / years * data.parcelrate / 100)),
      price2: util.money(Math.round(util.getFloat(plan.area) * util.getFloat(data.repair.savingprice) * data.parcelrate / 100))
    },
    {
      index: 5,
      title: 'm2 평균 계획금액',
      title2: 'm2 평균 적립금액',
      price: Math.round(util.getFloat(plan.total) / years / plan.area * data.parcelrate / 100),
      price2: util.money(data.repair.savingprice * data.parcelrate / 100)
    }
  ]
  
  data.totalcompares = compares
}

function getDurationMonth(endyear, endmonth, startyear, startmonth) {
  endyear = util.getInt(endyear)
  startyear = util.getInt(startyear)
  endmonth = util.getInt(endmonth)
  startmonth = util.getInt(startmonth)

  if (startyear > endyear) {
    return ''
  }
  
  if (startyear == endyear) {
    if (startmonth > endmonth) {
      return ''
    }
    
    return endmonth - startmonth + 1
  }

  let total = 12 - startmonth + 1
  total += (endyear - startyear - 1) * 12
  total += endmonth

  return total
}

function makeItems(items) {
  let alltotalprice = 0
  let sum = 0
  for (let i = 0; i < items.length; i++) {
    let item = items[i]
    let per = parseFloat(util.getFloat(item.rate))
    let totalprice = util.getInt((util.getFloat(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    let duration = util.getFloat(getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth))
    let price = parseFloat(totalprice) / (util.getFloat(data.report.totalsize) * duration)  
    items[i].totalprice = util.money(totalprice)    
    
    sum += util.getFloat(item.rate)

    items[i].price = util.moneyfloat(item.price)
    items[i].totalrate = util.getFloat(sum)

    if (i < items.length - 1) {
      alltotalprice += totalprice
    }
  }

  if (items.length > 0 && items[items.length - 1].totalrate === 100.0) {    
    items[items.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - alltotalprice)
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
    let res = await Outline.remove(item)
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

onMounted(() => {
  data.apt = parseInt(route.params.id)
  
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

function clickCancel() {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      data.visible = false
    })
  } else {
    data.visible = false
  }
}

const spanMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  
  if (columnIndex === 0) {
    if (rowIndex == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }
  return {rowspan: 1, colspan: 1}
}

function cellClassName({row, columnIndex}) {
  return 'title'  
}

const getSummaries = (param: SummaryMethodProps) => {
  const columns = param.columns
  const items = param.data
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index == 1) {
      sums[index] = util.money(util.getFloat(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears)))
    } else if (index == 2) {
      let total = 0
      if (items != null) {        
        items.forEach((item) => {
          total += util.getInt(item.totalprice)
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

const getSummariesPlan = (param: SummaryMethodProps) => {
  const columns = param.columns
  const items = param.data
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index == 1) {
      sums[index] = util.money(util.getFloat(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears)))
    } else if (index == 2) {
      let total = 0
      if (items != null) {        
        items.forEach((item) => {
          total += util.getInt(item.totalprice)
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

function cellClassNameTotal({row, columnIndex}) {
  if (row.index % 2 < 2) {
    return 'value'
  } else {
    return 'title'    
  }
}

const spanMethodTotal = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  if (columnIndex == 0 || columnIndex == 1 || columnIndex == 5 || columnIndex == 6) {
    if (rowIndex % 2 == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }

  return {rowspan: 1, colspan: 1}
}

const spanMethodTotalPlan = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  if (rowIndex == 10) {
    if (columnIndex == 2) {
      return {rowspan: 1, colspan: 3}
    } else if (columnIndex == 3 || columnIndex == 4) {
      return {rowspan: 0, colspan: 0}
    } else {
      return {rowspan: 1, colspan: 1}
    }
  }
  
  if (columnIndex == 0 || columnIndex == 1 || columnIndex == 5 || columnIndex == 6) {
    if (rowIndex % 2 == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }

  return {rowspan: 1, colspan: 1}
}

</script>
<style>
.inputNumber .el-input__inner {
  text-align: right;
}

.inputText .el-input__inner {
  text-align: left;
}

.title {
  background-color: #fafafa;
}

.value {
  background-color: #FFF;  
}
</style>

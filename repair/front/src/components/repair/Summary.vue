<template>
  <el-table :data="data.totals" border :summary-method="getSummaries" show-summary>
    <el-table-column prop="title" label="공사종별" />
    <el-table-column prop="price" label="항목별 일회 수선금액" align="right" :formatter="money" />
    <el-table-column prop="saveprice" label="년간 추정 적립 금액" align="right" :formatter="money" />
    <el-table-column prop="totalsaveprice" label="항목별 총계획 금액" align="right" :formatter="money" />
  </el-table>

  <div v-for="(item, index) in data.items">
    <el-table :data="item.items" border style="margin-top:10px;" :show-header="index == 0">
      <el-table-column prop="title" label="공사종별">
        <template #default="scope">
          <span v-if="scope.$index == 0" style="color:#af0202;">{{scope.row.title}}</span>
          <span v-else>{{scope.row.title}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="method" label="수선방법" align="center" width="60" />
      <el-table-column prop="cycle" label="수선주기" align="center" width="70">
        <template #default="scope">
          <span v-if="scope.$index == 0" style="color:#af0202;">{{scope.row.cycle}}</span>
          <span v-else>{{scope.row.cycle}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="percent" label="수선율" align="center" width="70">
        <template #default="scope">
          <span v-if="scope.$index == 0" style="color:#af0202;">{{scope.row.percent}}</span>
          <span v-else>{{scope.row.percent}}</span>
        </template>
      </el-table-column>
        <el-table-column prop="price" label="항목별 일회 수선금액" align="right">
          <template #default="scope">
          <span v-if="scope.$index == 0" style="color:#af0202;">{{util.money(scope.row.price)}}</span>
          <span v-else>{{util.money(scope.row.price)}}</span>
          </template>
        </el-table-column>
      <el-table-column prop="saveprice" label="년간 추정 적립 금액" align="right">
        <template #default="scope">
          <span v-if="scope.$index == 0" style="color:#af0202;">{{util.money(scope.row.saveprice)}}</span>
          <span v-else>{{util.money(scope.row.saveprice)}}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>

</template>
<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Repair, Report } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import type { TabsPaneContext } from 'element-plus'

const store = useStore()
const route = useRoute()

const data = reactive({
  apt: 0,
  items: [],
  totals: [],
  visible: false
})

async function initData() {
}

async function getItems() {
  if (data.apt == 0) {
    return
  }

  let res = await Repair.get(data.apt)
  //data.parcelrate = res.item.parcelrate
  
  res = await Report.summary(data.apt)
  console.log(res)
  data.totals = res.totals
  data.items = res.items  
}

onMounted(() => {
  data.apt = parseInt(route.params.id)  
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

const getSummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index == 0) {
      sums[index] = '총 계획금액'
    } else if (index === 1) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += item.price
        })
      }
      
      sums[index] = util.money(total)    
    } else if (index === 2) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += item.saveprice
        })
      }
      
      sums[index] = util.money(total)    
    } else if (index === 3) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += item.totalsaveprice
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

function money(row, column, value, index) {
  return util.money(value)
}

</script>

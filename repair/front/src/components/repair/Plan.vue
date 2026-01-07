<template>
  <div v-for="item in data.items">
    <PageTitle :title="item.title" />

    <el-table :data="item.datas" border>
      <el-table-column prop="index" label="순번" align="center" width="40">
        <template #default="scope">
          {{scope.row.index}}
        </template>
      </el-table-column>
      <el-table-column label="대분류">
        <template #default="scope">
          {{getCategory(scope.row.topcategory).name}}
        </template>
      </el-table-column>
      <el-table-column label="중분류">
        <template #default="scope">
          {{getCategory(scope.row.subcategory).name}}
        </template>
      </el-table-column>
      <el-table-column label="소분류">
        <template #default="scope">
          {{getCategory(scope.row.category).name}}
        </template>
      </el-table-column>
      <el-table-column label="규격">
        <template #default="scope">
          {{scope.row.extra.standard.name}}          
        </template>
      </el-table-column>
      <el-table-column label="수선방법" width="70" align="center">
        <template #default="scope">
          {{scope.row.extra.category.name}}
        </template>
      </el-table-column>
    </el-table>  
  </div>
  
  
</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Category, Dong, Standard, History, Breakdown, Totalyearreport, Repair } from "~/models"
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
  categoryMap: {},
  parcelrate: 0
})

async function initData() {
  if (data.apt == 0) {
    return
  }
  
  let res = await Standard.findByApt(data.apt)
  data.standards = res.items
  
  res = await Category.findByApt(data.apt)
  data.categorys = res.items

  
  let categoryMap = {}
  res.items.forEach((item) => {
    categoryMap[item.id] = item
  })

  data.categoryMap = categoryMap  
}

async function getItems() {
  if (data.apt == 0) {
    return
  }
  
  let res = await Repair.get(data.apt)
  //data.parcelrate = res.item.parcelrate
  
  res = await Totalyearreport.find({    
    apt: data.apt,
    orderby: 'b_duedate'
  })
    
  if (res.items == null) {
    res.items = []
  }    

  let reports = res.items

  let d = new Date()
  let year = d.getFullYear()

  res = await Repair.get(data.apt)
  let repair = res.item

  console.log(repair)
  if (repair.reportdate != '') {
    console.log(repair.reportdate)
    let temp = repair.reportdate.split('-')

    year = util.getInt(temp[0])

    console.log(year)
  }


  let items = []
  let yearsCount = 0
  
  for (let i = 0; i < 4; i++) {    
    let datas = []
    let index = 1

    let flag = false
    for (let j = 0; j < reports.length; j++) {
      let report = reports[j]

      if (year + i == report.duedate) {
        report.index = index
        datas.push(report)

        flag = true
        index++                    
      }
    }

    if (flag == false) {
      continue
    }
    
    items.push({
      year: year + i,
      title: `수립대상시설 ${year + i}년`,
               datas: datas
    })
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
     util.loading(true)
     
     initData()
     getItems()
   */
  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  //util.loading(false)
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

</script>

<template>
  <Title title="순찰 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <!--
    <el-input v-model="data.search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    -->

    <TotalDiv :total="data.total" />          
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :height="height(160)" v-infinite="getItems" @row-click="clickUpdate" >
    <el-table-column prop="id" label="ID" width="50" align="center" />
    <el-table-column prop="extra.apt.name" label="아파트명" />
    <el-table-column prop="extra.user.name" label="순찰자" width="80" align="center" />
    <el-table-column prop="location" label="장소" />
    <el-table-column prop="content" label="점검결과" />
    <el-table-column prop="process" label="처리결과" />
    <el-table-column prop="opinion" label="점검자 의견" />
    <el-table-column prop="startdate" label="시작 시간" align="center" width="140" />
    <el-table-column prop="enddate" label="종료 시간" align="center" width="140" />
  </el-table>  
  
</template>

<script setup lang="ts">

import { reactive, onMounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Patrol, Apt } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

const { width, height } = size()

const store = useStore()
const route = useRoute()

function clickSearch() {
  getItems(true)
}

const item = {
  id: 0,
  apt: 0,
  location: '',
  content: '',
  process: '',
  opinion: '',
  status: 0,
  startdate: '',
  enddate: ''  
}

const data = reactive({
  items: [],
  total: 0,
  page: 1,
  pagesize: 30,
  item: util.clone(item),
  search: {
    status: 0,
    text: ''
  }
})

async function initData() {
  data.statuss = [{id: 0, name: ' '}, {id: 1, name: '진행'}, {id: 2, name: '완료'}]
}

async function getItems(reset: boolean) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }
  
  let res = await Patrol.find({page: data.page, pagesize: data.pagesize, status: data.search.status})

  console.log(res)


  if (res.items == null) {
    res.items = []
  }
  
  data.total = res.total
  data.items = data.items.concat(res.items)

  data.page++
}

onMounted(async () => {
  util.loading(true)
  
  await initData()
  await getItems()

  util.loading(false)
})

async function clickView(item) {
  const apt = item.id
  const id = item.repairid
  router.push(`/${apt}/repair/${id}/breakdown`)
}

function clickUpdate(item, index) {
  const apt = item.id
  router.push(`/${apt}/patrol/patrol`)
}
</script>

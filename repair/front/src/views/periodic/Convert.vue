<template>
  <Title title="기존 자료 변환" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="danger" @click="clickSubmit">데이터 변환</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="width:80px;">작업명</y-th>
      <y-td>
        <div style="display:flex;justify-content:space-between;">
          <div style="margin-top:3px;">{{data.select.search}}</div> 
          <el-button size="small" type="primary" @click="clickAptSearch">작업 찾기</el-button>
        </div>
      </y-td>        
    </y-tr>    
  </y-table>

  <y-table style="margin-top:10px;">
    <y-tr>
      <y-th>
        기존 도면
      </y-th>
      <y-th>
        변환 도면
      </y-th>
    </y-tr>
    <y-tr v-for="(item, index) in data.images">
      <y-td>
        <div :style="{'margin-left': `${item.level * 30}px`}">{{item.name}}</div>
      </y-td>
      <y-td>
        <el-select v-model.number="data.images[index].target" style="width:100%;">
          <el-option v-for="item in data.blueprints" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>        
      </y-td>
    </y-tr>
  </y-table>
  
  <el-dialog
    v-model="data.visible"
    width="800"
  >

    <div style="display:flex;gap: 10px;margin-bottom:10px;">
      <el-input v-model="data.search" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />
      <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    </div>  
    <el-table :data="data.items" border style="width: 100%;"  :height="'300px'" v-infinite="getItems" @row-click="clickUpdate" >
      <el-table-column prop="id" label="ID" width="50" align="center" />      
      <el-table-column prop="search" label="작업명" />
      <el-table-column label="작업시작일" align="center" width="160">
        <template #default="scope">
          {{scope.row.workstartdate.replace('T', ' ').replace('Z', '')}}
        </template>
      </el-table-column>
      <el-table-column label="작업종료일" align="center" width="160">
        <template #default="scope">
          {{scope.row.workenddate.replace('T', ' ').replace('Z', '')}}
        </template>
      </el-table-column>
    </el-table>    

    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>      
    </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Oldapt, Oldimage, Blueprint } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  date: ''
}

const data = reactive({
  id: 0,
  page: 1,
  pagesize: 50,
  item: util.clone(item),
  original: util.clone(item),
  visible: false,
  search: '',
  items: [],
  select: {
    id: 0,
    search: ''
  },
  images: [],
  blueprints: []
})

async function initData() {
  
}

async function getItems(reset) {
  if (reset) {
    data.page = 1
    data.items = []
  }

  util.loading(true)
  
  let res = await Oldapt.find({
    page: data.page,
    pagesize: data.pagesize,
    search: data.search,
    orderby: 'a_name'    
  })

  if (res.items == null) {
    res.items = []
  } else {
    data.page = data.page + 1
  }

  let items = data.items.concat(res.items) 

  data.total = res.total  
  data.items = items

  util.loading(false)  
}

async function clickUpdate(item, index) {
  util.loading(true)
  
  data.select = item

  let res = await Oldimage.find({
    apt: data.select.id,
    orderby: 'i_order'
  })

  let title = ''
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]
    
    res.items[i].target = null

    if (item.level == 0) {
      title = item.name
      res.items[i].fullname = item.name
    } else {
      res.items[i].fullname = title + ' - ' + item.name
    }
  }

  let images = res.items  

  res = await Blueprint.find({apt: data.apt, category: 1, orderby: 'bp_parentorder,bp_order desc,bp_id'})
  
  title = ''
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    if (item.level == 1) {
      title = item.name
    } else {
      res.items[i].name = title + ' - ' + item.name
    }
  }

  res.items = res.items.filter(item => item.upload != 0)

  for (let i = 0; i < images.length; i++) {
    let item = images[i]
    
    for (let j = 0; j < res.items.length; j++) {
      if (item.fullname == res.items[j].name) {
        images[i].target = res.items[j].id
        break
      }
    }
  }

  data.images = images
  data.blueprints = res.items
  
  data.visible = false

  util.loading(false)
}

onMounted(async () => {
  util.loading(true)
  
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  await initData()
  await getItems()

  util.loading(false)
})

function clickAptSearch() {
  data.visible = true
}

async function clickSearch() {
  await getItems(true)
}

async function clickSubmit() {
  util.confirm('데이터 변환을 하면 기존 데이터가 모두 삭제됩니다. 실행하시겠습니까', async function() {
    util.loading(true)
    
    let items = util.clone(data.images)

    for (let i = 0; i < items.length; i++) {
      items[i].company = items[i].target
      items[i].master = data.id
    }
    
    await Oldapt.convert(items)

    util.info('변환되었습니다')

    util.loading(false)
  })  
}

</script>

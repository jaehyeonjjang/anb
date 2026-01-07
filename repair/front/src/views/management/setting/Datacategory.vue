<template>
  <Title title="입력항목 관리" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
        
    <el-select v-model.number="data.search.category" placeholder="분류" style="width:150px;">           
      <el-option
        v-for="item in data.categorys"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert">등록</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :height="height(170)" @row-click="clickUpdate"  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />    
    <el-table-column label="분류" align="center" width="100">
      <template #default="scope">
        {{getCategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column prop="name" label="명칭" align="left" />
    <el-table-column prop="remark" label="비고" align="left" />
    <el-table-column label="구분" align="center" width="100">
      <template #default="scope">
        <span v-if="scope.row.type==1">기본</span>
        <span v-if="scope.row.type==2">기타</span>
      </template>
    </el-table-column>    
    <el-table-column prop="order" label="순번" align="center" width="100" />
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    width="800px"
  >

      <y-table>
        <y-tr>
          <y-th>분류</y-th>
          <y-td>
            <el-select v-model.number="data.item.category" placeholder="분류" style="width:150px;">           
              <el-option
                v-for="item in data.categorys"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>명칭</y-th>
          <y-td>
            <el-input v-model="data.item.name" />
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>비고</y-th>
          <y-td>
            <el-input v-model="data.item.remark" />
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>구분</y-th>
          <y-td>
            <el-radio-group v-model.number="data.item.type">
              <el-radio-button size="small" label="1">기본</el-radio-button>
              <el-radio-button size="small" label="2">기타</el-radio-button>
            </el-radio-group>
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>순번</y-th>
          <y-td>
            <el-input v-model.number="data.item.order" />
          </y-td>
        </y-tr>        
      </y-table>

      <template #footer>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Datacategory } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Datacategory

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const item = {
  id: 0,
  type: 2,
  filename: '',
  name: '',
  use: 1,
  order: 0,
  periodic: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  search: {
    category: 0
  },
  categorys: [
    {id: 0, name: ' '},
    {id: 1, name: '부위'},
    {id: 2, name: '부재 - 빨강'},
    {id: 10, name: '부재 - 파랑'},
    {id: 3, name: '결함종류 - 빨강'},
    {id: 11, name: '결함종류 - 파랑'},
    {id: 4, name: '폭'},
    {id: 5, name: '길이'},
    {id: 6, name: '개소'},
    {id: 7, name: '진행사항'},
    {id: 8, name: '비고'},
    {id: 20, name: '기울기 비고'},
    {id: 21, name: '보'}
  ]
})

function getCategory(pos) {
  for (let i = 0; i < data.categorys.length; i++) {
    let item = data.categorys[i]

    if (item.id == pos) {
      return item.name
    }
  }

  return ''
}

async function clickSearch() {
  await getItems(true)
}

async function initData() {  
}

async function getItems() {
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    category: data.search.category,
    orderby: 'dc_category,dc_order,dc_id'
  })

  if (res.items == null) {
    res.items = []
  }

  let items = []
  
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    item.index = i + 1
    items.push(item)
  }

  data.total = res.total
  data.items = items
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0) {
    return
  }

  data.item = util.clone(item)
  data.visible = true  
}

onMounted(async () => {
  data.apt = parseInt(route.params.apt)
  
  util.loading(true)
  
  await initData()
  await getItems()

  data.visible = false
  util.loading(false)
})

function clickCancel() {
  data.visible = false
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
    
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id
      }

      await model.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  util.loading(true)

  let item = util.clone(data.item)
  
  item.type = util.getInt(item.type)
  item.category = util.getInt(item.category)
  item.order = util.getInt(item.order)

  if (item.id > 0) {
    await model.update(item)
  } else {
    await model.insert(item)
  }

  util.info('등록되었습니다')
  
  await getItems()

  data.visible = false  
  util.loading(false)  
}

</script>

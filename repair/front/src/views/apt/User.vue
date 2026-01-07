<template>
  <Title title="사용자 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <el-input v-model="search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>

    <TotalDiv :total="data.total" />
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :height="height(200)" v-infinite="getItems">
    <el-table-column prop="id" label="ID" width="80" align="center" />
    <el-table-column prop="loginid" label="아이디" />
    <el-table-column prop="name" label="이름" />
    <el-table-column prop="email" label="이메일" />    
    <el-table-column prop="level" width="100" label="권한" align="center">
      <template #default="scope">
        <el-tag :type="User.getLevelType(scope.row.level)">
          {{User.getLevel(scope.row.level)}}
        </el-tag>
      </template>
    </el-table-column>    
    <el-table-column label="" width="200" align="center" >
      <template #default="scope">
        <el-button size="small" @click="clickUpdate(scope.$index, scope.row)">수정</el-button>
        <el-button size="small" type="danger" @click="clickDelete(scope.$index, scope.row)">삭제</el-button>
      </template>
    </el-table-column>
  </el-table>  
  <div style="margin-top:10px;display:flex;justify-content: space-between;">
    <el-button size="small" type="success" @click="clickInsert">등록</el-button>
  </div>

  <el-dialog
    v-model="data.visible"
    title="사용자 등록/수정"
    width="600px"
    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="80px">
      <el-form-item label="ID" v-show="data.item.id != 0">
        {{ data.item.id }}
      </el-form-item>
      <el-form-item label="아이디">
        <el-input v-model="data.item.loginid" />
      </el-form-item>
      <el-form-item label="비밀번호">
        <el-input v-model="data.item.passwd" type="password" show-password />
      </el-form-item>
      <el-form-item label="이름">
        <el-input v-model="data.item.name" />
      </el-form-item>
      <el-form-item label="이메일">
        <el-input v-model="data.item.email" />
      </el-form-item>
      
      <el-form-item label="권한">
        <el-select v-model.number="data.item.level" class="m-2" placeholder="권한">
          <el-option
            v-for="(item, index) in User.aptlevels"
            :key="index"
            :label="item"
            :value="index"
          />
        </el-select>

      </el-form-item>      
    </el-form>

    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>  
</template>

<script setup lang="ts">

import { reactive, onMounted } from "vue"
import router from '~/router'
import { useRoute } from 'vue-router'
import { util, size }  from "~/global"
import { User } from "~/models"

const route = useRoute()

const { width, height } = size()

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems(true)
}

const item = {
  id: 0,
  loginid: '',
  passwd: '',
  name: '',
  email: '',
  level: 1,
  apt: 0,
  memo: ''
}

const data = reactive({
  apt: 0,
  items: [],
  total: 0,
  page: 1,
  pagesize: 100,
  item: util.clone(item),
  visible: false
})

async function initData() {  
}

async function getItems(reset) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }  

  let res = await User.find({page: data.page, pagesize: data.pagesize, apt: data.apt, name: search.text})

  if (res.items == null) {
    res.items = []
  }
  
  data.total = res.total
  data.items = data.items.concat(res.items)
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true
}

function clickUpdate(pos, item) {
  data.item = util.clone(item)
  data.visible = true
}

function clickDelete(pos, item) {
  util.confirm('삭제하시겠습니까', async function() {
    let res = await User.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      getItems(true)
    }
  })
}

async function clickSubmit() {
  const item = data.item
  if (item.loginid === '') {
    util.error('아이디를 입력하세요')
    return    
  }

  if (item.passwd === '') {
    util.error('비밀번호를 입력하세요')
    return
  }

  if (item.passwd === '') {
    util.error('이름을 입력하세요')
    return
  }

  if (item.level === 0) {
    util.error('권한을 선택하세요')
    return
  }
  
  let res;


  let count = await User.countByLoginid(item.loginid)
  if (count > 0) {
    util.error('이미 등록된 아이디입니다. 다른 아이디를 입력하세요')
    return
  }
  
  item.apt = data.apt

  if (item.id === 0) {
    res = await User.insert(item)
  } else {
    res = await User.update(item)
  }

  if (res.code === 'ok') {
    util.info('등록되었습니다')
    getItems(true)
    data.visible = false
  } else {
    util.error('오류가 발생했습니다')
  }
}

const handleClose = (done: () => void) => {
  util.confirm('팝업창을 닫으시겠습니까', function() {
    done()
  })  
}

onMounted(() => {  
  data.apt = parseInt(route.params.apt)  
  
  initData()
  getItems()
})

</script>

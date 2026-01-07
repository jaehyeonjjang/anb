<template>
  <Title title="견적 기본값 관리" />

  <y-table>
    <y-tr>
      <y-th>기술사</y-th>
      <y-td>{{util.money(data.item.person1)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>특급기술자</y-th>
      <y-td>{{util.money(data.item.person2)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>고급기술자</y-th>
      <y-td>{{util.money(data.item.person3)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>중급기술자</y-th>
      <y-td>{{util.money(data.item.person4)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>초급기술자</y-th>
      <y-td>{{util.money(data.item.person5)}}</y-td>
    </y-tr>
  </y-table>
  
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
        <el-input v-model="data.item.passwd" />
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
            v-for="item in data.levels"
            :key="item.id"
            :label="item.name"
            :value="item.id"
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
import { util, size }  from "~/global"
import { Standardwage } from "~/models"

const { width, height } = size()

const model = Standardwage

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems(true)
}

const item = {    
}

const data = reactive({
  items: [],
  total: 0,  
  item: util.clone(item),
  visible: false  
})

async function initData() {  
}

async function getItems(reset) {
  let res = await model.get(1)

  data.item = res.item  
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true
}

function clickUpdate(pos, item) {
  data.item = util.clone(item)
  data.visible = true
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

onMounted(async () => {
  util.loading(true)
  
  await initData()
  await getItems()

  util.loading(false)
})

</script>

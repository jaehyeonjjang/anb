<template>
  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="width:150px;">작업명</y-th>
      <y-td>{{data.periodic.name}}</y-td>
    </y-tr>
    <y-tr>
      <y-th style="width:150px;">아파트명</y-th>
      <y-td>
        <span v-if="data.periodic.aptname != ''">{{data.periodic.aptname}}</span>
        <span v-else>{{data.aptItem.name}}</span>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th style="width:150px;">과업범위</y-th>
      <y-td>{{data.periodic.taskrange}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>리포트 작업일</y-th>
      <y-td>{{util.viewDate(data.periodic.reportdate)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>상태</y-th>
      <y-td>{{Periodic.getStatus(data.periodic.status)}}</y-td>
    </y-tr>
  </y-table>

  <el-dialog
    v-model="data.visible"
    title="기본 정보 수정"
    :before-close="handleClose"
  >


  <y-table>
    <y-tr>
      <y-th style="width:150px;">작업명</y-th>
      <y-td><el-input v-model.model="data.item.name" /></y-td>
    </y-tr>
    <y-tr>
      <y-th style="width:150px;">아파트명</y-th>
      <y-td><el-input v-model.model="data.item.aptname" />
        <div style="margin-top:10px;">*입력하지 않으면 원래의 아파트명이 입력됩니다</div>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th style="width:150px;">과업범위</y-th>
      <y-td><el-input v-model.model="data.item.taskrange" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>리포트 작업일</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.reportdate" placeholder="" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>상태</y-th>
      <y-td>
        <el-radio-group v-model.number="data.item.status">
          <el-radio-button size="small" label="1">준비</el-radio-button>
          <el-radio-button size="small" label="2">착수</el-radio-button>
          <el-radio-button size="small" label="3">완료</el-radio-button>
          <el-radio-button size="small" label="4">중단</el-radio-button>
        </el-radio-group>
      </y-td>
    </y-tr>
  </y-table>

  <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">저장</el-button>
    </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Periodic, Apt } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  name: '',
  startdate: null,
  enddate: null,
  supply: '',
  contract: '',
  price: 0,
  safetygrade: 0,
  aptname: '',
  taskrange: '',
  status: 0,
  apt: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  item: util.clone(item),
  periodic: util.clone(item),
  visible: false,
  aptItem: {
    name: ''
  }
})

watch(() => route.params.id, async () => {
  data.apt = util.getInt(route.params.apt)
  data.id = util.getInt(route.params.id)
  
  await initData()
  await getItems()
})

async function initData() {
  let res = await Apt.get(data.apt)
  data.aptItem = res.item
}

async function getItems() {
  let res = await Periodic.get(data.id)
  data.periodic = res.item
}

async function clickUpdate(pos) {
  let res = await Periodic.get(data.id)
  const item = res.item

  data.item = util.clone(item)
  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.status = util.getInt(item.status)

  item.reportdate = util.convertDBDate(item.reportdate)
  
  await Periodic.update(item)

  util.info('수정되었습니다')

  data.periodic = item
  data.visible = false
}

const handleClose = (done: () => void) => {
  /*
     util.confirm('팝업창을 닫으시겠습니까', function() {
     done()
     })
   */

  done()
}

onMounted(async () => {
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  await initData()
  await getItems()
})

</script>

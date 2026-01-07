<template>
  <Title title="발생원인 분석" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="width:80px;">구분</y-th>
      <y-th>층별 주요 증상</y-th>      
    </y-tr>
    <y-tr>
      <y-td>지붕층</y-td>
      <y-td>{{data.original.cause1}}</y-td>            
    </y-tr>
    <y-tr>
      <y-td>지상층</y-td>
      <y-td>{{data.original.cause2}}</y-td>            
    </y-tr>
    <y-tr>
      <y-td>지하층</y-td>
      <y-td>{{data.original.cause3}}</y-td>    
    </y-tr>
    <y-tr>
      <y-td>외벽</y-td>
      <y-td>{{data.original.cause4}}</y-td>            
    </y-tr>
    <y-tr>
      <y-td>부대시설</y-td>
      <y-td>{{data.original.cause5}}</y-td>            
    </y-tr>
  </y-table>

  <div class="resulttext">
    <div style="margin-top:4px;">{{data.original.cause6}}</div>    
  </div>

  
  <el-dialog
    v-model="data.visible"
  >


    <y-table>
    <y-tr>
      <y-th style="width:80px;">구분</y-th>
      <y-th>층별 주요 증상</y-th>      
    </y-tr>
    <y-tr>
      <y-td>지붕층</y-td>
      <y-td><el-input v-model="data.item.cause1" :rows=5 type="textarea" style="font-size:12px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-td>지상층</y-td>
      <y-td><el-input v-model="data.item.cause2" :rows=5 type="textarea" style="font-size:12px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-td>지하층</y-td>
      <y-td><el-input v-model="data.item.cause3" :rows=5 type="textarea" style="font-size:12px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-td>외벽</y-td>
      <y-td><el-input v-model="data.item.cause4" :rows=5 type="textarea" style="font-size:12px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-td>부대시설</y-td>
      <y-td><el-input v-model="data.item.cause5" :rows=5 type="textarea" style="font-size:12px;" /></y-td>
    </y-tr>
    </y-table>
    
    <div style="margin-top:5px;">
      <el-input v-model="data.item.cause6" :rows=2 type="textarea" style="font-size:12px;" />
  </div>


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
import { Periodicopinion } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodicopinion

const item = {
  id: 0,
  result1: '',
  result2: '',
  result3: '',
  result4: '',
  result5: '',
  result6: '',
  status1: '',
  status2: '',
  status3: '',
  status4: '',
  status5: '',
  status6: '',
  position1: '',
  position2: '',
  position3: '',
  position4: '',
  position5: '',
  position6: '',
  periodic: 0,
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  original: util.clone(item),
  visible: false
})

async function initData() {
}

async function getItems() {
  let res = await model.getByPeriodic(data.id)
  data.original = res.item
}

async function clickUpdate(pos) {
  let res = await model.getByPeriodic(data.id)
  const item = res.item

  data.item = util.clone(item)
  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)
  
  await model.update(item)

  util.info('수정되었습니다')

  data.original = item
  data.visible = false
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

</script>

<template>
  <Title title="종합의견" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="width:100px;">안전등급</y-th>
      <y-td>
        <span v-if="data.original.grade == 1">양호</span>
        <span v-if="data.original.grade == 2">보통</span>
        <span v-if="data.original.grade == 3">불량</span>
      </y-td>            
    </y-tr>
    <y-tr>
      <y-th rowspan="3" style="width:100px;">내용</y-th>
      <y-td>{{data.original.content1}}</y-td>            
    </y-tr>
    <y-tr>
      <y-td>{{data.original.content2}}</y-td>            
    </y-tr>
    <y-tr>
      <y-td>{{data.original.content3}}</y-td>    
    </y-tr>        
  </y-table>

  <div class="resulttext">* 내용을 입력하지 않으면 자동으로 생성되어 입력됩니다.</div>


  
  <el-dialog
    v-model="data.visible"
  >

    
  <y-table>
    <y-tr>
      <y-td>
        <el-radio-group v-model.number="data.item.grade">
          <el-radio-button size="small" label="1">양호</el-radio-button>
          <el-radio-button size="small" label="2">보통</el-radio-button>
          <el-radio-button size="small" label="3">불량</el-radio-button>
        </el-radio-group>
      </y-td>
    </y-tr>
    <y-tr>
      <y-td><el-input v-model="data.item.content1" :rows=5 type="textarea" style="font-size:12px;" /></y-td>            
    </y-tr>
    <y-tr>
      <y-td><el-input v-model="data.item.content2" :rows=5 type="textarea" style="font-size:12px;" /></y-td>            
    </y-tr>
    <y-tr>
      <y-td><el-input v-model="data.item.content3" :rows=5 type="textarea" style="font-size:12px;" /></y-td>            
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

  item.grade = util.getInt(item.grade)
  
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

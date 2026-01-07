<template>
  <Title title="건축물 구조상태" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th>최고층고</y-th>
      <y-td>부위 : {{data.detail.struct1}}, 층고 : {{data.detail.struct2}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>기둥최장간격</y-th>
      <y-td>{{data.detail.struct3}} × {{data.detail.struct4}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>지정형식</y-th>
      <y-td>{{data.struct5s[data.detail.struct5]}}</y-td>
    </y-tr>    
    <y-tr>
      <y-th>PILE․PIER의 근입심도</y-th>
      <y-td>{{data.detail.struct6}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>PILE의 지지방법</y-th>
      <y-td>{{data.detail.struct7}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>강재종류</y-th>
      <y-td>{{data.detail.struct8}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>외벽 주요 마감자재</y-th>
      <y-td>{{data.detail.struct9}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>실내바닥 마감자재</y-th>
      <y-td>{{data.detail.struct10}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>실내벽체 마감자재</y-th>
      <y-td>{{data.detail.struct11}}</y-td>
    </y-tr>
  </y-table>

  <el-dialog
    v-model="data.visible"
    title="건축물 구조상태 수정"
    :before-close="handleClose"
  >


  <y-table>    
    <y-tr>
      <y-th>최고층고</y-th>
      <y-td>부위 : <el-input v-model.model="data.item.struct1" style="width:100px;" />, 층고 : <el-input v-model.model="data.item.struct2" style="width:100px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>기둥최장간격</y-th>
      <y-td><el-input v-model.model="data.item.struct3" style="width:100px;" /> × <el-input v-model.model="data.item.struct4" style="width:100px;" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>지정형식</y-th>
      <y-td>
        <el-select v-model.number="data.item.struct5" style="width:100%;">
          <el-option v-for="(item, index) in data.struct5s" :key="index" :label="item" :value="index" />
        </el-select>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>PILE․PIER의 근입심도</y-th>
      <y-td><el-input v-model="data.item.struct6" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>PILE의 지지방법</y-th>
      <y-td><el-input v-model="data.item.struct7" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>강재종류</y-th>
      <y-td><el-input v-model="data.item.struct8" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>외벽 주요 마감자재</y-th>
      <y-td><el-input v-model="data.item.struct9" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>실내바닥 마감자재</y-th>
      <y-td><el-input v-model="data.item.struct10" /></y-td>
    </y-tr>
    <y-tr>
      <y-th>실내벽체 마감자재</y-th>
      <y-td><el-input v-model="data.item.struct11" /></y-td>
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
import { Aptdetail } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  struct1: '',
  struct2: '',
  struct3: '',
  struct4: '',
  struct5: 0,
  struct6: '',
  struct7: '',
  struct8: '',
  struct9: '',
  struct10: '',
  struct11: '',  
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  detail: util.clone(item),
  visible: false,  
  struct5s: [' ', 'PHC말뚝', '현장말뚝', '모래잡석', '피어']
})

async function initData() {
}

async function getItems() {
  let res = await Aptdetail.get(data.apt)
  data.detail = res.item  
}

async function clickUpdate(pos) {
  let res = await Aptdetail.get(data.apt)
  const item = res.item

  data.item = util.clone(item)

  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.struct5 = util.getInt(item.struct5)  

  await Aptdetail.update(item)

  util.info('수정되었습니다')

  data.detail = item
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

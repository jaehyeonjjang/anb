<template>
  <Title title="점검 일반사항" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <PageTitle title="설계도서류" />
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th>시공관계 사진철 보관유무</y-th>
      <y-td>{{viewRadio(data.detail.blueprint1)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>도서보관함 설치 유무</y-th>
      <y-td>
        <span v-if="data.detail.blueprint2 & 1">▣ 양호&nbsp;&nbsp;&nbsp;</span>
        <span v-if="data.detail.blueprint2 & 2">▣ 보통&nbsp;&nbsp;&nbsp;</span>
        <span v-if="data.detail.blueprint2 & 4">▣ 일반 케비넷 사용&nbsp;&nbsp;&nbsp;</span>
        <span v-if="data.detail.blueprint2 & 8">▣ 없음</span>
      </y-td>        
    </y-tr>
    <y-tr>
      <y-th>재하시험 보고서</y-th>
      <y-td>{{viewRadio(data.detail.blueprint3)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>인·허가 서류</y-th>
      <y-td>{{viewRadio(data.detail.blueprint4)}}</y-td>
    </y-tr>
  </y-table>

  <PageTitle title="건출물 관리대장 활용" />
  
  <y-table>
    <y-tr>
      <y-th>작성유무 및 보관 실태</y-th>
      <y-td>{{viewRadio(data.detail.blueprint5)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>내용 갱신 유무</y-th>
      <y-td>{{viewRadio(data.detail.blueprint6)}}</y-td>
    </y-tr>
  </y-table>

  <PageTitle title="건출물 유지관리 계획 수립·시행" />
  
  <y-table>
    <y-tr>
      <y-th>유지관리 계획서 작성 유무</y-th>
      <y-td>{{viewRadio(data.detail.blueprint7)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>유지관리 계획서 보고 유무</y-th>
      <y-td>{{viewRadio(data.detail.blueprint8)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>정기점검 실시 유무</y-th>
      <y-td>{{viewRadio(data.detail.blueprint9)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>정기점검 실시간격</y-th>
      <y-td>{{data.detail.blueprint10}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>정기점검자 자격</y-th>
      <y-td>
        <span v-if="data.detail.blueprint11 & 1">▣ 관리주체직원&nbsp;&nbsp;&nbsp;</span>
        <span v-if="data.detail.blueprint11 & 2">▣ 외부점검 전문기관의뢰&nbsp;&nbsp;&nbsp;</span>
        <span v-if="data.detail.blueprint11 & 4">▣ 유자격자&nbsp;&nbsp;&nbsp;</span>
        <span v-if="data.detail.blueprint11 & 8">▣ 무자격자</span>
      </y-td>
    </y-tr>
  </y-table>

  
  <el-dialog
    v-model="data.visible"
    title="점검 일반사항 수정"
    :before-close="handleClose"
  >

    <PageTitle title="설계도서류" />


    <y-table>
      <y-tr>
        <y-th>시공관계 사진철 보관유무</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint1" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>          
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>도서보관함 설치 유무</y-th>
        <y-td>
          <el-checkbox size="small" v-model="data.check1">양호</el-checkbox>
          <el-checkbox size="small" v-model="data.check2">보통</el-checkbox>
          <el-checkbox size="small" v-model="data.check4">일반 캐비넷 사용</el-checkbox>
          <el-checkbox size="small" v-model="data.check8">없음</el-checkbox>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>재하시험 보고서</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint3" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>인·허가 서류</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint4" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>        
      </y-tr>
    </y-table>

    <PageTitle title="건출물 관리대장 활용" />
    
    <y-table>
      <y-tr>
        <y-th>작성유무 및 보관 실태</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint5" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>        
      </y-tr>
      <y-tr>
        <y-th>내용 갱신 유무</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint6" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>
      </y-tr>
    </y-table>

    <PageTitle title="건출물 유지관리 계획 수립·시행" />
    
    <y-table>
      <y-tr>
        <y-th>유지관리 계획서 작성 유무</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint7" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>유지관리 계획서 보고 유무</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint8" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>정기점검 실시 유무</y-th>
        <y-td>
          <el-radio-group v-model="data.item.blueprint9" style="margin-left:5px;">
            <el-radio label="1" size="small">유</el-radio>
            <el-radio label="2" size="small">무</el-radio>
          </el-radio-group>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>정기점검 실시간격</y-th>
        <y-td>
          <el-input v-model="data.item.blueprint10" placeholder="" />
        </y-td>
      </y-tr>      
      <y-tr>
        <y-th>정기점검자 자격</y-th>
        <y-td>
          <el-checkbox size="small" v-model="data.chk1">관리주체직원</el-checkbox>
          <el-checkbox size="small" v-model="data.chk2">외부점검 전문기관의뢰</el-checkbox>
          <el-checkbox size="small" v-model="data.chk4">유자격자</el-checkbox>
          <el-checkbox size="small" v-model="data.chk8">무자격자</el-checkbox>
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
import { Detail } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  blueprint1: 0,
  blueprint2: 0,
  blueprint3: 0,
  blueprint4: 0,
  blueprint5: 0,
  blueprint6: 0,
  blueprint7: 0,
  blueprint8: 0,
  blueprint9: 0,
  blueprint10: 0,  
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  detail: util.clone(item),
  visible: false,
  total: 0,
  check1: false,
  check2: false,
  check4: false,
  check8: false,  
  chk1: false,
  chk2: false,
  chk4: false,
  chk8: false
})

async function initData() {
}

async function getItems() {
  let res = await Detail.get(data.id)
  data.detail = res.item  
}

async function clickUpdate(pos) {
  let res = await Detail.get(data.id)
  const item = res.item


  item.blueprint1 = String(item.blueprint1)
  item.blueprint3 = String(item.blueprint3)
  item.blueprint4 = String(item.blueprint4)
  item.blueprint5 = String(item.blueprint5)
  item.blueprint6 = String(item.blueprint6)
  item.blueprint7 = String(item.blueprint7)
  item.blueprint8 = String(item.blueprint8)
  item.blueprint9 = String(item.blueprint9)
  item.blueprint10 = String(item.blueprint10)

  data.check1 = item.blueprint2 & 1 ? true : false
  data.check2 = item.blueprint2 & 2 ? true : false
  data.check4 = item.blueprint2 & 4 ? true : false
  data.check8 = item.blueprint2 & 8 ? true : false  

  data.chk1 = item.blueprint11 & 1 ? true : false
  data.chk2 = item.blueprint11 & 2 ? true : false
  data.chk4 = item.blueprint11 & 4 ? true : false
  data.chk8 = item.blueprint11 & 8 ? true : false

  data.item = util.clone(item)

  console.log(data.item)
  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.blueprint1 = util.getInt(item.blueprint1)
  item.blueprint3 = util.getInt(item.blueprint3)
  item.blueprint4 = util.getInt(item.blueprint4)
  item.blueprint5 = util.getInt(item.blueprint5)
  item.blueprint6 = util.getInt(item.blueprint6)
  item.blueprint7 = util.getInt(item.blueprint7)
  item.blueprint8 = util.getInt(item.blueprint8)
  item.blueprint9 = util.getInt(item.blueprint9)  

  var checked = 0

  if (data.check1) checked += 1
  if (data.check2) checked += 2
  if (data.check4) checked += 4
  if (data.check8) checked += 8

  item.blueprint2 = checked

  checked = 0

  if (data.chk1) checked += 1
  if (data.chk2) checked += 2
  if (data.chk4) checked += 4
  if (data.chk8) checked += 8

  item.blueprint11 = checked

  console.log(item)
  await Detail.update(item)

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

function viewRadio(value) {
  value = util.getInt(value)

  if (value == 1) {
    return '유'
  } else if (value == 2) {
    return '무'
  }

  return ''
}
</script>

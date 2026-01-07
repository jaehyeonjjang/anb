<template>
  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="text-align:center;" colspan="2">보존대상목록</y-th>
      <y-th style="width:80px;text-align:center;">유무</y-th>
      <y-th style="text-align:center;">관리주체 보유현황</y-th>
      <y-th style="text-align:center;">비고</y-th>      
    </y-tr>
    <y-tr>
      <y-td rowspan="2" style="width:100px;">설계도서</y-td>
      <y-td style="width:100px;">공통</y-td>
      <y-td style="text-align:center;" v-if="data.original.status1 == 1">▣ 유  □ 무</y-td>
      <y-td style="text-align:center;" v-else>□ 유  ▣ 무</y-td>    
      <y-td><div v-html="util.nl2br(data.original.content1)"></div></y-td>
      <y-td><div v-html="util.nl2br(data.original.remark1)"></div></y-td>
    </y-tr>
    <y-tr>      
      <y-td>설계도면</y-td>
      <y-td style="text-align:center;" v-if="data.original.status2 == 1">▣ 유  □ 무</y-td>
      <y-td style="text-align:center;" v-else>□ 유  ▣ 무</y-td>    
      <y-td><div v-html="util.nl2br(data.original.content2)"></div></y-td>
      <y-td><div v-html="util.nl2br(data.original.remark2)"></div></y-td>      
    </y-tr>
    <y-tr>      
      <y-td colspan="2">시설물 관리대장</y-td>
      <y-td style="text-align:center;" v-if="data.original.status3 == 1">▣ 유  □ 무</y-td>
      <y-td style="text-align:center;" v-else>□ 유  ▣ 무</y-td>    
      <y-td><div v-html="util.nl2br(data.original.content3)"></div></y-td>
      <y-td><div v-html="util.nl2br(data.original.remark3)"></div></y-td>
    </y-tr>
    <y-tr>      
      <y-td colspan="2">시공관련 자료</y-td>
      <y-td style="text-align:center;" v-if="data.original.status4 == 1">▣ 유  □ 무</y-td>
      <y-td style="text-align:center;" v-else>□ 유  ▣ 무</y-td>    
      <y-td><div v-html="util.nl2br(data.original.content4)"></div></y-td>
      <y-td><div v-html="util.nl2br(data.original.remark4)"></div></y-td>
    </y-tr>
    <y-tr>      
      <y-td colspan="2">안전점검 및 정밀안전진단 자료</y-td>
      <y-td style="text-align:center;" v-if="data.original.status5 == 1">▣ 유  □ 무</y-td>
      <y-td style="text-align:center;" v-else>□ 유  ▣ 무</y-td>    
      <y-td><div v-html="util.nl2br(data.original.content5)"></div></y-td>
      <y-td><div v-html="util.nl2br(data.original.remark5)"></div></y-td>
    </y-tr>
    <y-tr>      
      <y-td colspan="2">보수.보강 자료</y-td>
      <y-td style="text-align:center;" v-if="data.original.status6 == 1">▣ 유  □ 무</y-td>
      <y-td style="text-align:center;" v-else>□ 유  ▣ 무</y-td>    
      <y-td><div v-html="util.nl2br(data.original.content6)"></div></y-td>
      <y-td><div v-html="util.nl2br(data.original.remark6)"></div></y-td>      
    </y-tr>
  </y-table>


  <el-dialog
    v-model="data.visible"
    title="자료 보유 현황 수정"
    width="1000px"
  >

    <y-table>
      <y-tr>
        <y-th style="text-align:center;" colspan="2">보존대상목록</y-th>
        <y-th style="width:80px;text-align:center;">유무</y-th>
        <y-th style="text-align:center;">관리주체 보유현황</y-th>
        <y-th style="text-align:center;width:200px;">비고</y-th>      
      </y-tr>
      <y-tr>
        <y-td rowspan="2" style="width:100px;">설계도서</y-td>
        <y-td style="width:100px;">공통</y-td>
        <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.item.status1">
              <el-radio-button size="small" label="1">유</el-radio-button>
              <el-radio-button size="small" label="2">무</el-radio-button>              
          </el-radio-group>
        </y-td>              
        <y-td><el-input :rows="5" type="textarea" style="font-size:12px;" v-model="data.item.content1"/></y-td>
        <y-td><el-input :rows="5" type="textarea" style="font-size:12px;" v-model="data.item.remark1"/></y-td>        
      </y-tr>
      <y-tr>      
        <y-td>설계도면</y-td>
            <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.item.status2">
              <el-radio-button size="small" label="1">유</el-radio-button>
              <el-radio-button size="small" label="2">무</el-radio-button>              
          </el-radio-group>
        </y-td>              
        <y-td><el-input :rows="3" type="textarea" style="font-size:12px;" v-model="data.item.content2"/></y-td>
        <y-td><el-input :rows="3" type="textarea" style="font-size:12px;" v-model="data.item.remark2"/></y-td>         
      </y-tr>
      <y-tr>      
        <y-td colspan="2">시설물 관리대장</y-td>
        <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.item.status3">
            <el-radio-button size="small" label="1">유</el-radio-button>
            <el-radio-button size="small" label="2">무</el-radio-button>              
          </el-radio-group>
        </y-td>              
        <y-td><el-input :rows="3" type="textarea" style="font-size:12px;" v-model="data.item.content3"/></y-td>
        <y-td><el-input :rows="3" type="textarea" style="font-size:12px;" v-model="data.item.remark3"/></y-td> 
      </y-tr>
      <y-tr>      
        <y-td colspan="2">시공관련 자료</y-td>
        <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.item.status4">
            <el-radio-button size="small" label="1">유</el-radio-button>
            <el-radio-button size="small" label="2">무</el-radio-button>              
          </el-radio-group>
        </y-td>              
        <y-td><el-input :rows="3" type="textarea" style="font-size:12px;" v-model="data.item.content4"/></y-td>
        <y-td><el-input :rows="3" type="textarea" style="font-size:12px;" v-model="data.item.remark4"/></y-td> 
      </y-tr>
      <y-tr>      
        <y-td colspan="2">안전점검 및 정밀안전진단 자료</y-td>
        <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.item.status5">
            <el-radio-button size="small" label="1">유</el-radio-button>
            <el-radio-button size="small" label="2">무</el-radio-button>              
          </el-radio-group>
        </y-td>              
        <y-td><el-input :rows="2" type="textarea" style="font-size:12px;" v-model="data.item.content5"/></y-td>
        <y-td><el-input :rows="2" type="textarea" style="font-size:12px;" v-model="data.item.remark5"/></y-td> 
      </y-tr>
      <y-tr>      
        <y-td colspan="2">보수.보강 자료</y-td>
        <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.item.status6">
            <el-radio-button size="small" label="1">유</el-radio-button>
            <el-radio-button size="small" label="2">무</el-radio-button>              
          </el-radio-group>
        </y-td>              
        <y-td><el-input :rows="2" type="textarea" style="font-size:12px;" v-model="data.item.content6"/></y-td>
        <y-td><el-input :rows="2" type="textarea" style="font-size:12px;" v-model="data.item.remark6"/></y-td> 
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
import { Periodickeep } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodickeep

const item = {
  id: 0,
  status1: 1,
  status2: 1,
  status3: 1,
  status4: 2,
  status5: 1,
  status6: 2,
  content1: '',
  content2: '',
  content3: '',
  content4: '',
  content5: '',
  content6: '',
  remark1: '',
  remark2: '',
  remark3: '',
  remark4: '',
  remark5: '',
  remark6: '',
  periodic: 0,
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  original: util.clone(item),
  visible: false,
  total: 0,
})

async function initData() {
}

async function getItems() {
  let res = await model.getByPeriodic(data.id)  
  data.original = res.item  
}

async function clickUpdate(pos) {
  let res = await model.getByPeriodic(data.id)  
  
  data.item = util.clone(res.item)

  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.status1 = util.getInt(item.status1)
  item.status2 = util.getInt(item.status2)
  item.status3 = util.getInt(item.status3)
  item.status4 = util.getInt(item.status4)
  item.status5 = util.getInt(item.status5)
  item.status6 = util.getInt(item.status6)

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

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  await initData()
  await getItems()

  util.loading(false)
})

</script>

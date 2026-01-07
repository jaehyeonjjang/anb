<template>
  <div style="display:flex;justify-content:space-between;">
    <PageTitle title="건축물의 개요" />
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th style="width:100px;">대지면적</y-th>
      <y-td>{{data.periodic.outline1}}m²</y-td>
      <y-th style="width:100px;">건축면적</y-th>
      <y-td>{{data.periodic.outline2}}m²</y-td>    
      <y-th style="width:100px;">연면적</y-th>
      <y-td>{{data.periodic.outline6}}m²</y-td>
    </y-tr>
    <y-tr>
      <y-th>구조형식</y-th>
      <y-td colspan="5">
        <span v-if="data.periodic.outline7==7">{{data.periodic.outline7content}}</span>
        <span v-else>{{data.outline7s[data.periodic.outline7]}}</span>
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>최고높이</y-th>
      <y-td colspan="5">{{data.periodic.outline8}}m</y-td>
    </y-tr>
    <y-tr>
      <y-th>주용도</y-th>
      <y-td colspan="5">
        <span v-if="data.periodic.outline9==5">{{data.periodic.outline9content}}</span>
        <span v-else>{{data.outline9s[data.periodic.outline9]}}</span>
      </y-td>
    </y-tr>
  </y-table>

  <PageTitle title="건축물 이력사항" />

  <y-table>
    <y-tr>
      <y-th style="width:100px;">설계자</y-th>
      <y-td>{{data.periodic.record1}}</y-td>    
      <y-th style="width:100px;">감리자</y-th>
      <y-td>{{data.periodic.record2}}</y-td>    
      <y-th style="width:100px;">시공자</y-th>
      <y-td>{{data.periodic.record3}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>공사기간</y-th>
      <y-td colspan="5">
        {{util.viewDate(data.periodic.record4)}} ~ {{util.viewDate(data.periodic.record5)}}
      </y-td>
    </y-tr>
    <y-tr>
      <y-th colspan="5">사용승인일</y-th>
      <y-td>{{util.viewDate(data.aptItem.useapproval)}}</y-td>
    </y-tr>
  </y-table>

  <PageTitle title="설계도서 보존실태" />

  <y-table>
    <y-tr>
      <y-th style="width:100px;">사용승인도면</y-th>
      <y-td v-if="data.periodic.blueprint1 == 1">▣ 유  □ 무
        <span v-if="data.periodic.blueprint1save == 1">(FMS 등록)</span>
        <span v-else>(FMS 미등록)</span>
      </y-td>
      <y-td v-else>□ 유  ▣ 무
      </y-td>
    </y-tr>
    <y-tr>
      <y-th>시방서</y-th>
      <y-td v-if="data.periodic.blueprint2 == 1">▣ 유  □ 무</y-td>
      <y-td v-else>□ 유  ▣ 무</y-td>
    </y-tr>
    <y-tr>
      <y-th>구조계산서</y-th>
      <y-td v-if="data.periodic.blueprint3 == 1">▣ 유  □ 무</y-td>
      <y-td v-else>□ 유  ▣ 무</y-td>
    </y-tr>
    <y-tr>
      <y-th>공사관계철</y-th>
      <y-td v-if="data.periodic.blueprint4 == 1">▣ 유  □ 무</y-td>
      <y-td v-else>□ 유  ▣ 무</y-td>      
    </y-tr>
    <y-tr>
      <y-th>관리대장</y-th>
      <y-td v-if="data.periodic.blueprint5 == 1">▣ 유  □ 무</y-td>
      <y-td v-else>□ 유  ▣ 무</y-td>
    </y-tr>
    <y-tr>
      <y-th>유지관리계획서</y-th>
      <y-td v-if="data.periodic.blueprint6 == 1">▣ 유  □ 무</y-td>
      <y-td v-else>□ 유  ▣ 무</y-td>
    </y-tr>
  </y-table>
  

  <el-dialog
    v-model="data.visible"
    title="개요 수정"
    width="1000px"
    :before-close="handleClose"
  >


    <PageTitle title="건축물의 개요" />


    <y-table>
      <y-tr>
        <y-th style="width:100px;">대지면적</y-th>
        <y-td><el-input v-model="data.item.outline1" style="width:100px;" /> m²</y-td>      
        <y-th style="width:100px;">건축면적</y-th>
        <y-td><el-input v-model="data.item.outline2" style="width:100px;" /> m²</y-td>      
        <y-th style="width:100px;">연면적</y-th>
        <y-td><el-input v-model="data.item.outline6" style="width:100px;" /> m²</y-td>
      </y-tr>
      <y-tr>
        <y-th>구조형식</y-th>
        <y-td colspan="5">
          <el-select v-model.number="data.item.outline7" style="width:100%;">
            <el-option v-for="(item, index) in data.outline7s" :key="index" :label="item" :value="index" />
          </el-select>

          <div style="margin-top:5px;" v-if="data.item.outline7 == 7">
            <el-input v-model="data.item.outline7content"  />
          </div>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>최고높이</y-th>
        <y-td colspan="5"><el-input v-model="data.item.outline8" style="width:100px;" /> m</y-td>
      </y-tr>
      <y-tr>
        <y-th>주용도</y-th>
        <y-td colspan="5">
          <el-select v-model.number="data.item.outline9" style="width:100%;">
            <el-option v-for="(item, index) in data.outline9s" :key="index" :label="item" :value="index" />
          </el-select>
          <div style="margin-top:5px;" v-if="data.item.outline9 == 5">
            <el-input v-model="data.item.outline9content"  />
          </div>
        </y-td>
      </y-tr>
    </y-table>

    <PageTitle title="설계도서 보존실태" />

    <y-table>
      <y-tr>
        <y-th style="width:100px;">설계자</y-th>
        <y-td><el-input v-model="data.item.record1" /></y-td>
        <y-th style="width:100px;">감리자</y-th>
        <y-td><el-input v-model="data.item.record2" /></y-td>      
        <y-th style="width:100px;">시공자</y-th>
        <y-td><el-input v-model="data.item.record3" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>공사기간</y-th>
        <y-td colspan="5">
          <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.record4" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.record5" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>사용승인일</y-th>
        <y-td colspan="5">
          <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.useapproval" />
        </y-td>
      </y-tr>
    </y-table>


    <PageTitle title="건축물 이력사항" />

    <y-table>
      <y-tr>
        <y-th style="width:100px;" rowspan="2">사용승인서류<br>보관여부</y-th>
        <y-td>
          <el-checkbox size="small" v-model="data.item.blueprint1">사용승인도면</el-checkbox>
          <el-checkbox size="small" v-model="data.item.blueprint2">시방서</el-checkbox>
          <el-checkbox size="small" v-model="data.item.blueprint3">구조계산서</el-checkbox>
          <el-checkbox size="small" v-model="data.item.blueprint4">공사관계철</el-checkbox>
          <el-checkbox size="small" v-model="data.item.blueprint5">관리대장</el-checkbox>
          <el-checkbox size="small" v-model="data.item.blueprint6">유지관리계획서</el-checkbox>
        </y-td>
      </y-tr>
      <y-tr>
        <y-td>
          <el-checkbox size="small" v-model="data.item.blueprint1save">사용승인도면 FMS 등록</el-checkbox>
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
import { Apt, Aptperiodic, Periodic } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  outline1: '',
  outline2: '',
  outline3: '',
  outline4: '',
  outline5: '',
  outline6: '',
  outline7: '',
  outline7content: '',
  outline8: '',
  outline9: '',
  outline9content: '',
  blueprint1: '',
  blueprint1save: '',
  blueprint2: '',
  blueprint3: '',
  blueprint4: '',
  blueprint5: '',
  blueprint6: '',
  useapproval: '',
  date: ''
}

const aptItem = {
  id: 0,
  useapproval: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  periodic: util.clone(item),
  aptItem: util.clone(aptItem),
  visible: false,
  total: 0,
  outline7s: [' ', '철근콘크리트 구조', '철골.철근콘크리트 구조', '철근콘크리트라멘조+벽식 구조', '철골.철근콘크리트구조 및 복합구조', '프리케스트콘크리트 구조', '철근콘트리트 PC구조', '직접 입력'],
  outline9s: [' ', '공동주택(아파트)', '제 1종 근린생활시설', '제 2종 근린생활시설', '업무시설', '직접 입력']
})

function calculatePeriodic() {
  return util.getInt(data.periodic.outline3) + util.getInt(data.periodic.outline4) + util.getInt(data.periodic.outline5)
}

function calculateTotal() {
  data.total = util.getInt(data.item.outline3) + util.getInt(data.item.outline4) + util.getInt(data.item.outline5)
}

watch(() => data.item.outline3, () => {
  calculateTotal()
})

watch(() => data.item.outline4, () => {
  calculateTotal()
})

watch(() => data.item.outline5, () => {
  calculateTotal()
})


async function initData() {
}

async function getItems() {
  let res = await Aptperiodic.get(data.apt)
  console.log(res)
  const aptperiodic = res.item
  if (aptperiodic == null) {
    let newItem = util.clone(item)
    newItem.id = data.apt
    await Aptperiodic.insert(newItem)
    data.periodic = newItem
  } else {
    data.periodic = aptperiodic
  }

  res = await Periodic.get(data.id)
  let periodic = res.item
  data.periodic.blueprint1 = periodic.blueprint1
  data.periodic.blueprint1save = periodic.blueprint1save
  data.periodic.blueprint2 = periodic.blueprint2
  data.periodic.blueprint3 = periodic.blueprint3
  data.periodic.blueprint4 = periodic.blueprint4
  data.periodic.blueprint5 = periodic.blueprint5
  data.periodic.blueprint6 = periodic.blueprint6

  res = await Apt.get(data.apt)
  const apt = res.item
  data.aptItem = apt
}

async function clickUpdate(pos) {
  let res = await Aptperiodic.get(data.apt)
  const item = res.item

  res = await Periodic.get(data.id)
  let periodic = res.item
  item.blueprint1 = periodic.blueprint1 == 1 ? true : false
  item.blueprint1save = periodic.blueprint1save == 1 ? true : false
  item.blueprint2 = periodic.blueprint2 == 1 ? true : false
  item.blueprint3 = periodic.blueprint3 == 1 ? true : false
  item.blueprint4 = periodic.blueprint4 == 1 ? true : false
  item.blueprint5 = periodic.blueprint5 == 1 ? true : false
  item.blueprint6 = periodic.blueprint6 == 1 ? true : false

  console.log(item)

  item.useapproval = data.aptItem.useapproval
  
  data.item = util.clone(item)

  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.outline3 = util.getInt(item.outline3)
  item.outline4 = util.getInt(item.outline4)
  item.outline5 = util.getInt(item.outline5)
  item.outline7 = util.getInt(item.outline7)
  item.outline9 = util.getInt(item.outline9)

  item.outline1 = util.getFloat(item.outline1)
  item.outline2 = util.getFloat(item.outline2)
  item.outline6 = util.getFloat(item.outline6)
  item.outline8 = util.getFloat(item.outline8)

  item.record4 = util.convertDBDate(item.record4)
  item.record5 = util.convertDBDate(item.record5)

  await Aptperiodic.update(item)

  let res = await Periodic.get(data.id)
  let periodicItem = res.item

  periodicItem.blueprint1 = item.blueprint1 ? 1 : 0
  periodicItem.blueprint1save = item.blueprint1save ? 1 : 0
  periodicItem.blueprint2 = item.blueprint2 ? 1 : 0
  periodicItem.blueprint3 = item.blueprint3 ? 1 : 0
  periodicItem.blueprint4 = item.blueprint4 ? 1 : 0
  periodicItem.blueprint5 = item.blueprint5 ? 1 : 0
  periodicItem.blueprint6 = item.blueprint6 ? 1 : 0

  await Periodic.update(periodicItem)  

  res = await Apt.get(data.apt)
  let aptItem = res.item
  aptItem.useapproval = util.convertDBDate(item.useapproval)

  await Apt.update(aptItem)

  util.info('수정되었습니다')

  data.aptItem = aptItem
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

<template>
  <Title title="개요" />


  <div style="display:flex;justify-content:space-between;">
    <PageTitle title="건축물의 개요" />
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th>대지면적</y-th>
      <y-td>{{data.detail.outline1}}m²</y-td>
    </y-tr>
    <y-tr>
      <y-th>건축면적</y-th>
      <y-td>{{data.detail.outline2}}m²</y-td>
    </y-tr>
    <y-tr>
      <y-th>동수</y-th>
      <y-td>1종 - {{data.detail.outline3}}개동, 2종 - {{data.detail.outline4}}개동, 기타 - {{data.detail.outline5}}개동 (총계 - {{calculateDetail()}}개동)</y-td>
    </y-tr>
    <y-tr>
      <y-th>연면적</y-th>
      <y-td>{{data.detail.outline6}}m²</y-td>
    </y-tr>
    <y-tr>
      <y-th>구조형식</y-th>
      <y-td>{{data.outline7s[data.detail.outline7]}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>최고높이</y-th>
      <y-td>{{data.detail.outline8}}m</y-td>
    </y-tr>
    <y-tr>
      <y-th>주용도</y-th>
      <y-td>{{data.outline9s[data.detail.outline9]}}</y-td>
    </y-tr>
  </y-table>

  <PageTitle title="건축물 이력사항" />

  <y-table>
    <y-tr>
      <y-th>설계자</y-th>
      <y-td>{{data.detail.record1}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>감리자</y-th>
      <y-td>{{data.detail.record2}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>시공자</y-th>
      <y-td>{{data.detail.record3}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>공사기간</y-th>
      <y-td>
        {{util.viewDate(data.detail.record4)}} ~ {{util.viewDate(data.detail.record5)}}
      </y-td>
    </y-tr>
  </y-table>


  <el-dialog
    v-model="data.visible"
    title="개요 수정"

    :before-close="handleClose"
  >


    <PageTitle title="건축물의 개요" />


    <y-table>
      <y-tr>
        <y-th>대지면적</y-th>
        <y-td><el-input v-model="data.item.outline1" style="width:100px;" /> m²</y-td>
      </y-tr>
      <y-tr>
        <y-th>건축면적</y-th>
        <y-td><el-input v-model="data.item.outline2" style="width:100px;" /> m²</y-td>
      </y-tr>
      <y-tr>
        <y-th>동수</y-th>
        <y-td>1종 - <el-input v-model="data.item.outline3" style="width:50px;" /> 개동, 2종 - <el-input v-model="data.item.outline4" style="width:50px;" /> 개동, 기타 - <el-input v-model="data.item.outline5" style="width:50px;" /> 개동 (총계 - {{data.total}}개동)</y-td>
      </y-tr>
      <y-tr>
        <y-th>연면적</y-th>
        <y-td><el-input v-model="data.item.outline6" style="width:100px;" /> m²</y-td>
      </y-tr>
      <y-tr>
        <y-th>구조형식</y-th>
        <y-td>
          <el-select v-model.number="data.item.outline7" style="width:100%;">
            <el-option v-for="(item, index) in data.outline7s" :key="index" :label="item" :value="index" />
          </el-select>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>최고높이</y-th>
        <y-td><el-input v-model="data.item.outline8" style="width:100px;" /> m</y-td>
      </y-tr>
      <y-tr>
        <y-th>주용도</y-th>
        <y-td>
          <el-select v-model.number="data.item.outline9" style="width:100%;">
            <el-option v-for="(item, index) in data.outline9s" :key="index" :label="item" :value="index" />
          </el-select>
        </y-td>
      </y-tr>
    </y-table>

    <PageTitle title="건축물 이력사항" />

    <y-table>
      <y-tr>
        <y-th>설계자</y-th>
        <y-td><el-input v-model="data.item.record1" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>감리자</y-th>
        <y-td><el-input v-model="data.item.record2" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>시공자</y-th>
        <y-td><el-input v-model="data.item.record3" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>공사기간</y-th>
        <y-td>
          <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.record4" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.record5" />
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
import { Apt, Aptdetail } from "~/models"
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
  outline8: '',
  outline9: '',
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  detail: util.clone(item),
  visible: false,
  total: 0,
  outline7s: [' ', '철근콘트리트 구조', '철골.철근콘크리트 구조'],
  outline9s: [' ', '공동주택(아파트)']
})

function calculateDetail() {
  return util.getInt(data.detail.outline3) + util.getInt(data.detail.outline4) + util.getInt(data.detail.outline5)
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
  let res = await Aptdetail.get(data.apt)
  const aptdetail = res.item
  if (aptdetail == null) {
    let newItem = util.clone(item)
    newItem.id = data.apt
    await Aptdetail.insert(newItem)
    data.detail = newItem
  } else {
    data.detail = aptdetail
  }

  res = await Apt.get(data.apt)
  const apt = res.item
  data.aptItem = apt

}

async function clickUpdate(pos) {
  let res = await Aptdetail.get(data.apt)
  const item = res.item

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

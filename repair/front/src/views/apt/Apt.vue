<template>
  <Title title="기본 정보" />

  <div style="text-align:left;margin-bottom:10px;">
    <el-button size="small" type="success" @click="clickUpdate()">수정</el-button>
  </div>
  
  <y-table>
    <y-tr>
      <y-th>ID</y-th>
      <y-td>{{ data.item.id }}</y-td>
    </y-tr>
    
    <y-tr>
      <y-th>아파트명</y-th>
      <y-td>{{data.item.name}}</y-td>
    </y-tr>


    <y-tr>
      <y-th>준공년도</y-th><y-td>
        {{ data.item.completeyear }}
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>건축물형태</y-th><y-td>
        {{ data.item.type }}
      </y-td>
    </y-tr>
    
    <y-tr>
      <y-th>전체동수</y-th><y-td>
        {{ data.item.flatcount }}
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>세대수</y-th><y-td>
        {{ data.item.familycount }}
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>층수</y-th><y-td>
        {{ data.item.floor }}
      </y-td>
    </y-tr>

    <y-tr>
      <y-th>전화번호</y-th><y-td>
        {{ data.item.tel }}
      </y-td></y-tr>

    <y-tr>
      <y-th>팩스번호</y-th><y-td>
        {{ data.item.fax }}
      </y-td></y-tr>
    
    <y-tr>
      <y-th>공용메일주소</y-th><y-td>
        {{ data.item.email }}
      </y-td></y-tr>
    
    <y-tr v-if="data.level > 2">
      <y-th>개인메일주소</y-th><y-td>
        {{ data.item.personalemail }}
      </y-td></y-tr>

    <y-tr v-if="data.level > 2">
      <y-th>장기수선계획</y-th><y-td>
        <span v-html="util.nl2br(data.item.repair)" />
      </y-td></y-tr>
    <y-tr v-if="data.level > 2">
      <y-th>안전점검</y-th><y-td>
        <span v-html="util.nl2br(data.item.safety)" />      
      </y-td></y-tr>
    <y-tr v-if="data.level > 2">
      <y-th>기타</y-th><y-td>
        <span v-html="util.nl2br(data.item.fault)" />        
      </y-td></y-tr>      


    <y-tr>
      <y-th>우편번호</y-th><y-td>
        {{ data.item.zip }}
      </y-td></y-tr>
    


    <y-tr>
      <y-th>도로명주소</y-th><y-td>
        {{ data.item.address }}
      </y-td></y-tr>
    

    <y-tr>
      <y-th>지번주소</y-th><y-td>
        {{ data.item.address2 }}
      </y-td></y-tr>
    

    <y-tr>
      <y-th>정밀점검일자</y-th><y-td>
        {{ data.item.testdate }}
      </y-td></y-tr>


    <y-tr v-if="data.level > 2">
      <y-th>FMS 아이디</y-th><y-td>
        {{ data.item.fmsloginid }}
      </y-td></y-tr>
    
    

    <y-tr v-if="data.level > 2">
      <y-th>FMS 비번</y-th><y-td>
        {{ data.item.fmspasswd }}
      </y-td></y-tr>
    
    
    
    


    
    

    
    <y-tr v-if="data.level > 2">
      <y-th>계약구분</y-th><y-td>        
      </y-td></y-tr>
    
    <y-tr v-if="data.level > 2">
      <y-th>계약날짜</y-th><y-td>
        {{ data.item.contractdate }}
      </y-td></y-tr>

    
    <y-tr v-if="data.level > 2">
      <y-th>계약기간</y-th><y-td>
        {{ data.item.contractduration }}
      </y-td></y-tr>
    <y-tr v-if="data.level > 2">
      <y-th>계약금액</y-th><y-td>
        {{ data.item.contractprice }}
      </y-td></y-tr>
    <y-tr v-if="data.level > 2">
      <y-th>계산서 발행</y-th><y-td>
        {{ data.item.invoice }}            
      </y-td></y-tr>
    <y-tr v-if="data.level > 2">
      <y-th>입금날짜</y-th><y-td>
        {{ data.item.depositdate }}
      </y-td></y-tr>

  </y-table>

    
  <AptInsert :id="data.apt" :close="clickClose" ref="apt" />

</template>

<script setup lang="ts">

import { reactive, onMounted, ref } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Category, Repair } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const data = reactive({
  apt: 0,
  item: {
    id: 0
  },
  level: 0
})

const apt = ref({})

async function initData() {
}

async function getItems() {
  let res = await Apt.get(data.apt)

  data.item = res.item
}

function clickUpdate() {
  apt.value?.readData(data.apt)
}

onMounted(async () => {
  data.apt = parseInt(route.params.apt)

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  util.loading(true)
  
  await initData()
  await getItems()  
  
  util.loading(false)  
})

function clickClose() {
  getItems()
}

</script>
<style>
td {
  text-align: left;
}
</style>

<template>
  <Title title="예외적 집행" />


  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">등록</el-button>
    </div>
  </div>  


  <div :style="{'height': data.height, 'overflow': 'auto', 'padding-right': '10px'}">
  <div class="bigtitle">
    03. 장기수선충당금 사용시 예외적 사항의 인정
  </div>
  
  <div class="content"><span v-html="util.nl2br(data.item.content1)" /></div>

  <div class="maintitle">
    2.1 긴급공사 및 소액지출(예측불가한 사고)
  </div>

  <div class="subtitle">
    (1) 내용
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content2)" /></div>

  <div class="subtitle">
    (2) 수선대상
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content3)" /></div>

  <div class="subtitle">
    (3) 근거마련
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content4)" /></div>

  <div class="subtitle">
    (4) 장기수선계획 조정시기
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content3)" /></div>

  <div class="subtitle">
    (5) 사용절차
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content6)" /></div>

  <div class="subtitle">
    (6) 긴급공사 대상 시설물
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content7)" /></div>

  <div class="subtitle">
    (7) 소액지출 사용요건
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content8)" /></div>

  <div class="maintitle">
    2.2 장기수선충당금 사용 방법 및 절차
  </div>

  <div class="content"><span v-html="nl2number(data.item.content9)" /></div>

  <div class="maintitle">
    2.3 장기수선충당금 사용 금액의 범위
  </div>

  <div class="content"><span v-html="nl2dot(data.item.content10)" /></div>

  <div class="bigtitle">
    04. 향후 장기수선충담금에 대한 고찰
  </div>

  <div class="content"><span v-html="util.nl2br(data.item.content11)" /></div>

  </div>
  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1100px"
  >


    <div style="height:500px;overflow:auto;padding: 10px 10px;">
      <div class="bigtitle">
    장기수선충당금 사용시 예외적 집행
  </div>

  <el-input v-model="data.repair.content1" :rows="2" type="textarea" style="font-size:12px;" />


  <div class="maintitle">
    2.1 긴급공사 및 소액지출(예측불가한 사고)
  </div>

  <div class="subtitle">
    (1) 내용
  </div>

  <el-input v-model="data.repair.content2" :rows="4" type="textarea" style="font-size:12px;" />

  <div class="subtitle">
    (2) 수선대상
  </div>

  <el-input v-model="data.repair.content3" :rows="2" type="textarea" style="font-size:12px;" />

  <div class="subtitle">
    (3) 근거마련
  </div>

  <el-input v-model="data.repair.content4" :rows="2" type="textarea" style="font-size:12px;" />

  <div class="subtitle">
    (4) 장기수선계획 조정시기
  </div>

  <el-input v-model="data.repair.content5" :rows="2" type="textarea" style="font-size:12px;" />
  

  <div class="subtitle">
    (5) 사용절차
  </div>

  <el-input v-model="data.repair.content6" :rows="2" type="textarea" style="font-size:12px;" />
  
  <div class="subtitle">
    (6) 긴급공사 대상 시설물
  </div>

  <el-input v-model="data.repair.content7" :rows="20" type="textarea" style="font-size:12px;" />

  <div class="subtitle">
    (7) 소액지출 사용요건
  </div>

  <el-input v-model="data.repair.content8" :rows="6" type="textarea" style="font-size:12px;" />

  <div class="maintitle">
    2.2 장기수선충당금 사용 방법 및 절차
  </div>

  <el-input v-model="data.repair.content9" :rows="8" type="textarea" style="font-size:12px;" />

  <div class="maintitle">
    2.3 장기수선충당금 사용 금액의 범위
  </div>

  <el-input v-model="data.repair.content10" :rows="2" type="textarea" style="font-size:12px;" />  

  <div class="bigtitle">
    03. 향후 장기수선충담금에 대한 고찰
  </div>

  <el-input v-model="data.repair.content11" :rows="3" type="textarea" style="font-size:12px;" />

  </div>

  <template #footer>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util }  from "~/global"
import { Repair } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const store = useStore()
const route = useRoute()

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems()
}

function paging(page) {
  data.page = page
  getItems()
}

const item = {
  id: 0,
  content1: '',
  content2: '',
  content3: '',
  content4: '',
  content5: '',
  content6: '',
  content7: '',
  content8: '',
  content9: '',
  content10: '',
  content11: ''
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  itemplan: util.clone(item),
  visible: false,    
  search: {
    text: ''
  },
  batchs: []
})

async function initData() {  
}

async function getItems() {
  let res = await Repair.get(data.apt)
  data.item = res.item  
}

function clickInsert() {  
  data.repair = util.clone(data.item)

  data.mode = 'normal'
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0) {
    return
  }

  let items = makeItems([util.clone(item)])

  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Outline.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      getItems()
    }
  })
}

const handleClose = (done: () => void) => {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      done()
    })
  } else {
    done()
  }
}

function setWindowSize() {
  data.width = (window.innerWidth - 500) + 'px'
  data.height = (window.innerHeight - 170) + 'px'
}

onMounted(async () => {
  data.apt = parseInt(route.params.id)
  
  util.loading(true)
  
  await initData()
  await getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  data.visible = false
  util.loading(false)
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
})

function clickCancel() {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      data.visible = false
    })
  } else {
    data.visible = false
  }
}

async function clickSubmit() {
  util.loading(true)

  let item = data.repair
  let res = await Repair.get(data.apt)
  res.item.content1 = data.repair.content1
  res.item.content2 = data.repair.content2
  res.item.content3 = data.repair.content3
  res.item.content4 = data.repair.content4
  res.item.content5 = data.repair.content5
  res.item.content6 = data.repair.content6
  res.item.content7 = data.repair.content7
  res.item.content8 = data.repair.content8
  res.item.content9 = data.repair.content9
  res.item.content10 = data.repair.content10
  res.item.content11 = data.repair.content11
  await Repair.update(res.item)

  util.info('등록되었습니다')
  getItems()
  data.visible = false
  
  util.loading(false)  
}

function nl2dot(value) {
  return value.split('\n').map(item => `● ${item}`).join('<BR/>') 
}

function nl2number(value) {
  return value.split('\n').map((item, index) => `(${index+1}) ${item}`).join('<BR/>')
}  
</script>
<style>
.bigtitle {
  font-size:16px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;
}

.maintitle {
  font-size:14px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;background: #FFF;
}

.subtitle {
  font-size:12px;font-weight:bold;text-align:left;margin-top:5px;margin-bottom:5px;
}

.content {
  text-align: left;
}
</style>

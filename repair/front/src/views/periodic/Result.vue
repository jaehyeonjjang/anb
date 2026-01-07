<template>
  <Title title="변경사항" />

  <PageTitle title="용도변경" />    
  
  <div style="display:flex;justify-content:space-between;margin-top:-27px;margin-bottom:5px;margin-left:175px;">
    <div style="margin-top:8px;text-align:left;">
      <span v-if="data.periodic.result1 == 1">▣ 유 &nbsp; □ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result1 == 2">□ 유 &nbsp; ▣ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result1 == 3">□ 유 &nbsp; □ 무 &nbsp; ▣ 불명</span>
    </div>
    <el-button size="small" type="success" @click="clickUpdate(1)">수정</el-button>
  </div>

  <el-table :data="data.items1" border>    
    <el-table-column prop="content1" label="부위(층수)" align="center" />
    <el-table-column prop="" label="변경전" align="center">
      <el-table-column prop="content2" label="용도" align="center" />
      <el-table-column prop="content3" label="면적" align="center" />
    </el-table-column>
    <el-table-column prop="" label="변경후" align="center">
      <el-table-column prop="content4" label="용도" align="center" />
      <el-table-column prop="content5" label="면적" align="center" />
    </el-table-column>
    <el-table-column prop="content6" label="설계자" align="center" />
    <el-table-column prop="content7" label="날짜" align="center" />        
  </el-table>

  <div class="resulttext">{{data.periodic.resulttext1}}</div>

  <PageTitle title="구조 부재의 변경사항" />

  <div style="display:flex;justify-content:space-between;margin-top:-27px;margin-bottom:5px;margin-left:175px;">
    <div style="margin-top:8px;text-align:left;">
      <span v-if="data.periodic.result2 == 1">▣ 유 &nbsp; □ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result2 == 2">□ 유 &nbsp; ▣ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result2 == 3">□ 유 &nbsp; □ 무 &nbsp; ▣ 불명</span>
    </div>
    <el-button size="small" type="success" @click="clickUpdate(2)">수정</el-button>
  </div>    

  <el-table :data="data.items2" border>    
    <el-table-column prop="content1" label="부재" align="center" />
    <el-table-column prop="content2" label="위치" align="center" />
    <el-table-column prop="content3" label="내용" align="center" />
    <el-table-column prop="content4" label="담당자" align="center" />
    <el-table-column prop="content7" label="날짜" align="center" />        
  </el-table>

  <div class="resulttext">{{data.periodic.resulttext2}}</div>  
  
  <PageTitle title="주변조건의 변경사항" />

  <div style="display:flex;justify-content:space-between;margin-top:-27px;margin-bottom:5px;margin-left:175px;">
    <div style="margin-top:8px;text-align:left;">
      <span v-if="data.periodic.result3 == 1">▣ 유 &nbsp; □ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result3 == 2">□ 유 &nbsp; ▣ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result3 == 3">□ 유 &nbsp; □ 무 &nbsp; ▣ 불명</span>
    </div>
    <el-button size="small" type="success" @click="clickUpdate(3)">수정</el-button>
  </div>  

  <el-table :data="data.items3" border>    
    <el-table-column prop="content1" label="구분" align="center" />
    <el-table-column prop="content2" label="위치" align="center" />
    <el-table-column label="변경사항" align="center">
      <el-table-column prop="content3" label="변경전" align="center" />
      <el-table-column prop="content4" label="변경후" align="center" />
    </el-table-column>      
  </el-table>

  <div class="resulttext">{{data.periodic.resulttext3}}</div>

  <PageTitle title="증·개축" />
  
  <div style="display:flex;justify-content:space-between;margin-top:-27px;margin-bottom:5px;margin-left:175px;">
    <div style="margin-top:8px;text-align:left;">
      <span v-if="data.periodic.result4 == 1">▣ 유 &nbsp; □ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result4 == 2">□ 유 &nbsp; ▣ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result4 == 3">□ 유 &nbsp; □ 무 &nbsp; ▣ 불명</span>
    </div>
    <el-button size="small" type="success" @click="clickUpdate(4)">수정</el-button>
  </div>

  <el-table :data="data.items4" border>    
    <el-table-column prop="content1" label="부위(층수)" align="center" />
    <el-table-column label="증·개축 전" align="center">
      <el-table-column prop="content2" label="용도" align="center" />
      <el-table-column prop="content3" label="면적" align="center" />
    </el-table-column>
    <el-table-column label="증·개축 후" align="center">
      <el-table-column prop="content4" label="용도" align="center" />
      <el-table-column prop="content5" label="면적" align="center" />
    </el-table-column>
    <el-table-column prop="content6" label="설계자" align="center" />
    <el-table-column prop="content7" label="날짜" align="center" />        
  </el-table>

  <div class="resulttext">{{data.periodic.resulttext4}}</div>

  <PageTitle title="보수·보강 이력사항" />

  <div style="display:flex;justify-content:space-between;margin-top:-27px;margin-bottom:5px;margin-left:175px;">
    <div style="margin-top:8px;text-align:left;">
      <span v-if="data.periodic.result5 == 1">▣ 유 &nbsp; □ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result5 == 2">□ 유 &nbsp; ▣ 무 &nbsp; □ 불명</span>
      <span v-if="data.periodic.result5 == 3">□ 유 &nbsp; □ 무 &nbsp; ▣ 불명</span>
    </div>
    <el-button size="small" type="success" @click="clickUpdate(5)">수정</el-button>
  </div>

  <el-table :data="data.items5" border>    
    <el-table-column prop="content1" label="기간" align="center" />
    <el-table-column prop="content2" label="공사종류" align="center" />
    <el-table-column prop="content3" label="설계자" align="center" />
    <el-table-column prop="content4" label="시공자" align="center" />
    <el-table-column prop="content5" label="감리자" align="center" />
    <el-table-column prop="content6" label="비고" align="center" />        
  </el-table>

  <div class="resulttext">{{data.periodic.resulttext5}}</div>
  

  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1050px"
  >

    <div style="text-align:left;">
      <el-radio-group v-model="data.result">
        <el-radio-button size="small" label="1">유</el-radio-button>
        <el-radio-button size="small" label="2">무</el-radio-button>
        <el-radio-button size="small" label="3">해당없음</el-radio-button>
      </el-radio-group>
    </div>
    
    <el-form label-width="100px">      
      <el-table :data="data.batchs" border style="margin-top:15px;" v-if="data.pos == 1">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>

        
        <el-table-column prop="content1" label="부위(층수)" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content1" />
          </template>
        </el-table-column>
        <el-table-column prop="" label="변경전" align="center">
          <el-table-column prop="content2" label="용도" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content2" />
            </template>
          </el-table-column>
          <el-table-column prop="content3" label="면적" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content3" />
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column prop="" label="변경후" align="center">
          <el-table-column prop="content4" label="용도" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content4" />
            </template>
          </el-table-column>
          <el-table-column prop="content5" label="면적" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content5" />
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column prop="content6" label="설계자" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content6" />
          </template>
        </el-table-column>
        <el-table-column prop="content7" label="날짜" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content7" />
          </template>
        </el-table-column>
        
      </el-table>




      <el-table :data="data.batchs" border style="margin-top:15px;" v-if="data.pos == 2">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        
        <el-table-column prop="content1" label="부재" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content1" />
          </template>
        </el-table-column>
        <el-table-column prop="content2" label="위치" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content2" />
          </template>
        </el-table-column>
        <el-table-column prop="content3" label="내용" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content3" />
          </template>
        </el-table-column>
        <el-table-column prop="content4" label="담당자" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content4" />
          </template>
        </el-table-column>
        <el-table-column prop="content7" label="날짜" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content7" />
          </template>
        </el-table-column>
      </el-table>

      <el-table :data="data.batchs" border style="margin-top:15px;" v-if="data.pos == 3">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        
        <el-table-column prop="content1" label="구분" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content1" />
          </template>
        </el-table-column>
        <el-table-column prop="content2" label="위치" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content2" />
          </template>
        </el-table-column>
        <el-table-column label="변경사항" align="center">
          <el-table-column prop="content3" label="변경전" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content3" />
            </template>
          </el-table-column>
          <el-table-column prop="content4" label="변경후" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content4" />
            </template>
          </el-table-column>
        </el-table-column>      
      </el-table>

      <el-table :data="data.batchs" border style="margin-top:15px;" v-if="data.pos == 4">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        
        <el-table-column prop="content1" label="부위(층수)" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content1" />
          </template>
        </el-table-column>
        <el-table-column label="증·개축 전" align="center">
          <el-table-column prop="content2" label="용도" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content2" />
            </template>
          </el-table-column>
          <el-table-column prop="content3" label="면적" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content3" />
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column label="증·개축 후" align="center">
          <el-table-column prop="content4" label="용도" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content4" />
            </template>
          </el-table-column>
          <el-table-column prop="content5" label="면적" align="center">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].content5" />
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column prop="content6" label="설계자" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content6" />
          </template>
        </el-table-column>
        <el-table-column prop="content7" label="날짜" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content7" />
          </template>
        </el-table-column>
      </el-table>

      <el-table :data="data.batchs" border style="margin-top:15px;" v-if="data.pos == 5">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        
        <el-table-column prop="content1" label="기간" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content1" />
          </template>
        </el-table-column>
        <el-table-column prop="content2" label="공사종류" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content2" />
          </template>
        </el-table-column>
        <el-table-column prop="content3" label="설계자" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content3" />
          </template>
        </el-table-column>
        <el-table-column prop="content4" label="시공자" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content4" />
          </template>
        </el-table-column>
        <el-table-column prop="content5" label="감리자" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content5" />
          </template>
        </el-table-column>
        <el-table-column prop="content6" label="비고" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].content6" />
          </template>
        </el-table-column>
      </el-table>


      
    </el-form>

    <div style="margin-top:10px;">
    <el-input v-model="data.resulttext" />
    </div>

      <template #footer>
        <el-button size="small" type="danger" v-if="data.mode != 'batch' && (data.batchs.length > 0 && data.batchs[0].id > 0)" style="float:left;" @click="clickDelete">삭제</el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(1)"><el-icon><Plus /></el-icon></el-button>        
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Periodic, Periodicchange } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  content1: '',
  content2: '',
  content3: '',
  content4: '',
  content5: '',
  content6: '',
  content7: '',
  type: 0,
  order: 0,
  periodic: 0,
  date: ''
}

const periodic = {
  result1: 0,
  result2: 0,
  result3: 0,
  result4: 0,
  result5: 0,
  resulttext1: '',
  resulttext2: '',
  resulttext3: '',
  resulttext4: '',
  resulttext5: ''
}

const data = reactive({
  id: 0,
  pos: 1,
  item: util.clone(item),
  visible: false,
  total: 0,
  batchs: [],
  result: 0,
  resulttext: '',
  items1: [],
  items2: [],
  items3: [],
  items4: [],
  items5: [],
  periodic: util.clone(periodic)
})



async function initData() {  
}

async function getItems() {
  let res = await Periodic.get(data.id)
  data.periodic = res.item

  for (let i = 1; i <= 5; i++) {
    let res = await Periodicchange.find({periodic: data.id, type: i, orderby: 'pc_order,pc_id'})
    data[`items${i}`] = res.items
  }
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

function clickAdd(count) {
  let items = []
  for (let i = 0; i < count; i++) {
    items.push(util.clone(item))
  }

  data.batchs = data.batchs.concat(items)
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await model.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      await getItems()
    }
  })
}

async function clickUpdate(pos) {
  let res = await Periodic.get(data.id)
  const item = res.item

  data.result = item[`result${pos}`]
  data.resulttext = item[`resulttext${pos}`]  
  
  data.item = util.clone(item)

  res = await Periodicchange.find({periodic: data.id, type: pos, orderby: 'pc_order,pc_id'})
  data.batchs = res.items

  data.pos = pos
  data.mode = 'batch'
  data.visible = true
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function clickCancel() {
  data.visible = false
}

async function clickSubmit() {
  util.loading(true)
  
  let pos = data.pos
  let items = util.clone(data.batchs)

  let res = await Periodic.get(data.id)
  let periodicItem = res.item
  
  periodicItem[`result${pos}`] = util.getInt(data.result)
  periodicItem[`resulttext${pos}`] = data.resulttext

  await Periodic.update(periodicItem)

  res = await Periodicchange.find({periodic: data.id, type: pos, orderby: 'pc_order,pc_id'})
  for (let i = 0; i < res.items.length; i++) {
    let item = {
      id: res.items[i].id
    }

    await Periodicchange.remove(item)
  }

  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    item.type = pos
    item.periodic = data.id
    item.order = i + 1

    item.content1 = String(item.content1)
    item.content2 = String(item.content2)
    item.content3 = String(item.content3)
    item.content4 = String(item.content4)
    item.content5 = String(item.content5)
    item.content6 = String(item.content6)
    item.content7 = String(item.content7)
    
    await Periodicchange.insert(item)
  }

  await getItems()

  data.visible = false
  util.loading(false)  
  util.info('저장되었습니다')  
}

</script>

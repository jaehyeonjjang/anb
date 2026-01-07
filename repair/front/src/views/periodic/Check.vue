<template>
  <Title title="점검표" />

  <div style="display:flex;">
    <div style="border-right:1px solid #ccc;padding-right:10px;width:150px;">
      <el-table :data="data.dongs" border @row-click="clickView">
        <el-table-column prop="dong" label="시설물명" align="left" />
      </el-table>
    </div>
    <div style="flex:1;margin-left:10px;" v-if="data.dongs.length > 0">
      <div style="text-align:left;font-size:16px;font-weight:bold;margin-bottom:10px;margin-left:3px;display:flex;justify-content: space-between;">
        <div>{{data.dong.dong}}</div>
        <div>
          <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
        </div>
      </div>

      <div style="overflow:auto;" :style="{height: height(170)}">
        <y-table>
          <!--
          <y-tr>
            <y-th style="width:150px;">용도변경</y-th>
            <y-td>{{data.original.content1}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">구조부재 변경</y-th>
            <y-td>{{data.original.content2}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">증·개축</y-th>
            <y-td>{{data.original.content3}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">주변조건 변경</y-th>
            <y-td>{{data.original.content4}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">균열발생상황</y-th>
            <y-td>자동 입력</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">누수·백태현황</y-th>
            <y-td>자동 입력</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">철근노출 및 부식 상태</y-th>
            <y-td>자동 입력</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">강재구조 노후</y-th>
            <y-td>자동 입력</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">마감재</y-th>
            <y-td>자동 입력</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">지반(포장)</y-th>
            <y-td>{{data.original.content10}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">옹벽(축대)</y-th>
            <y-td>{{data.original.content11}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">담장</y-th>
            <y-td>{{data.original.content12}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">공중이 이용하는 부위</y-th>
            <y-td>
              <y-table>
                <y-tr>
                  <y-td>
                    <span v-if="data.original.use1 == 1">▣유 □무</span>
                    <span v-else>□유 ▣무</span>
                  </y-td>
                  <y-td>추락방지시설</y-td>
                  <y-td>전체적으로 양호함</y-td>
                  <y-td>
                    <span v-if="data.original.need1 == 1">▣보수필요 □불필요</span>
                    <span v-else>□보수필요 ▣불필요</span>                    
                  </y-td>
                </y-tr>
                <y-tr>
                  <y-td>
                    <span v-if="data.original.use2 == 1">▣유 □무</span>
                    <span v-else>□유 ▣무</span>                    
                  </y-td>
                  <y-td>도로포장</y-td>
                  <y-td>전체적으로 양호함</y-td>
                  <y-td>
                    <span v-if="data.original.need2 == 1">▣보수필요 □불필요</span>
                    <span v-else>□보수필요 ▣불필요</span>                    
                  </y-td>
                </y-tr>
                <y-tr>
                  <y-td>
                    <span v-if="data.original.use3 == 1">▣유 □무</span>
                    <span v-else>□유 ▣무</span>                    
                  </y-td>
                  <y-td>도로부 신축이음부</y-td>
                  <y-td>해당사항 없음</y-td>
                  <y-td>
                    <span v-if="data.original.need3 == 1">▣보수필요 □불필요</span>
                    <span v-else>□보수필요 ▣불필요</span>                    
                  </y-td>
                </y-tr>
                <y-tr>
                  <y-td>
                    <span v-if="data.original.use4 == 1">▣유 □무</span>
                    <span v-else>□유 ▣무</span>
                  </y-td>
                  <y-td>환기구 등의 덮게</y-td>
                  <y-td>전체적으로 양호함</y-td>
                  <y-td>
                    <span v-if="data.original.need4 == 1">▣보수필요 □불필요</span>
                    <span v-else>□보수필요 ▣불필요</span>                    
                  </y-td>
                </y-tr>
              </y-table>
            </y-td>
          </y-tr>
          -->
          <y-tr>
            <y-th style="width:150px;">천장 및 채광창 상태</y-th>
            <y-td>{{data.original.content13}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">기타</y-th>
            <y-td>{{data.original.content14}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">특기사항</y-th>
            <y-td>{{data.original.content15}}</y-td>
          </y-tr>
          <y-tr>
            <y-th style="width:150px;">점검자 의견</y-th>
            <y-td>{{data.original.content16}}</y-td>
          </y-tr>
        </y-table>
      </div>
    </div>
  </div>

  <el-dialog
    v-model="data.visible"
    title="점검표 수정"
    width="1000"    
  >

    <div style="overflow:auto;">
      <y-table>
        <!--
        <y-tr>
          <y-th style="width:150px;">용도변경</y-th>
          <y-td><el-input v-model="data.item.content1" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">구조부재 변경</y-th>
          <y-td><el-input v-model="data.item.content2" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">증·개축</y-th>
          <y-td><el-input v-model="data.item.content3" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">주변조건 변경</y-th>
          <y-td><el-input v-model="data.item.content4" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">균열발생상황</y-th>
          <y-td>자동 입력</y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">누수·백태현황</y-th>
          <y-td>자동 입력</y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">철근노출 및 부식 상태</y-th>
          <y-td>자동 입력</y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">강재구조 노후</y-th>
          <y-td>자동 입력</y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">마감재</y-th>
          <y-td>자동 입력</y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">지반(포장)</y-th>
          <y-td><el-input v-model="data.item.content10" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">옹벽(축대)</y-th>
          <y-td><el-input v-model="data.item.content11" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">담장</y-th>
          <y-td><el-input v-model="data.item.content12" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">공중이 이용하는 부위</y-th>
          <y-td>
            <y-table>
              <y-tr>
                <y-td>
                  <el-radio-group v-model.number="data.item.use1">
                    <el-radio-button size="small" label="1">유</el-radio-button>
                    <el-radio-button size="small" label="2">무</el-radio-button>              
                  </el-radio-group>
                </y-td>
                <y-td>추락방지시설</y-td>
                <y-td>전체적으로 양호함</y-td>
                <y-td>
                  <el-radio-group v-model.number="data.item.need1">
                    <el-radio-button size="small" label="1">보수필요</el-radio-button>
                    <el-radio-button size="small" label="2">불필요</el-radio-button>              
                  </el-radio-group>
                </y-td>
              </y-tr>
              <y-tr>
                <y-td>
                  <el-radio-group v-model.number="data.item.use2">
                    <el-radio-button size="small" label="1">유</el-radio-button>
                    <el-radio-button size="small" label="2">무</el-radio-button>              
                  </el-radio-group>
                </y-td>
                <y-td>도로포장</y-td>
                <y-td>전체적으로 양호함</y-td>
                <y-td>
                  <el-radio-group v-model.number="data.item.need2">
                    <el-radio-button size="small" label="1">보수필요</el-radio-button>
                    <el-radio-button size="small" label="2">불필요</el-radio-button>              
                  </el-radio-group>
                </y-td>
              </y-tr>
              <y-tr>
                <y-td>
                  <el-radio-group v-model.number="data.item.use3">
                    <el-radio-button size="small" label="1">유</el-radio-button>
                    <el-radio-button size="small" label="2">무</el-radio-button>              
                  </el-radio-group>
                </y-td>
                <y-td>도로부 신축이음부</y-td>
                <y-td>해당사항 없음</y-td>
                <y-td>
                  <el-radio-group v-model.number="data.item.need3">
                    <el-radio-button size="small" label="1">보수필요</el-radio-button>
                    <el-radio-button size="small" label="2">불필요</el-radio-button>              
                  </el-radio-group>
                </y-td>
              </y-tr>
              <y-tr>
                <y-td>
                  <el-radio-group v-model.number="data.item.use4">
                    <el-radio-button size="small" label="1">유</el-radio-button>
                    <el-radio-button size="small" label="2">무</el-radio-button>              
                  </el-radio-group>
                </y-td>
                <y-td>환기구 등의 덮게</y-td>
                <y-td>전체적으로 양호함</y-td>
                <y-td>
                  <el-radio-group v-model.number="data.item.need4">
                    <el-radio-button size="small" label="1">보수필요</el-radio-button>
                    <el-radio-button size="small" label="2">불필요</el-radio-button>              
                  </el-radio-group>
                </y-td>
              </y-tr>
            </y-table>
          </y-td>
        </y-tr>
        -->
        <y-tr>
          <y-th style="width:150px;">천장 및 채광창 상태</y-th>
          <y-td><el-input v-model="data.item.content13" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">기타</y-th>
          <y-td><el-input v-model="data.item.content14" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">특기사항</y-th>
          <y-td><el-input v-model="data.item.content15" :rows="2" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
        <y-tr>
          <y-th style="width:150px;">점검자 의견</y-th>
          <y-td><el-input v-model="data.item.content16" :rows="3" type="textarea" style="font-size:12px;"/></y-td>
        </y-tr>
      </y-table>
    </div>

    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">저장</el-button>
    </template>
  </el-dialog>
</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Periodiccheck, Aptdong } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodiccheck

const item = {
  id: 0,
  content1: '⦁용도 변경사항 없음(이전 점검일 이후)',
  content2: '⦁구조부재 변경사항 없음',
  content3: '⦁증.개축 없음',
  content4: '⦁주변조건 변경사항 없음',
  content5: '',
  content6: '',
  content7: '',
  content8: '',
  content9: '',
  content10: '⦁전반적으로 양호함',
  content11: '⦁해당사항 없음',
  content12: '⦁해당사항 없음',
  content13: '⦁해당사항 없음',
  content14: '⦁전반적으로 양호함',
  content15: '⦁특기사항 없음',
  content16: '⦁발생된 결함 사항에 대하여 보수·보강 우선순위 선정 후 유지관리계획에 따라 보수하여 관리하는 것이 바람직함',
  use1: 1,
  use2: 1,
  use3: 1,
  use4: 1,
  need1: 2,
  need2: 2,
  need3: 2,
  need4: 2,
  aptdong: 0,
  periodic: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  original: util.clone(item),  
  item: util.clone(item),
  visible: false,
  dongs: [],
  dong: {title: ''}
})

async function initData() {
  let res = await Aptdong.find({apt: data.apt, orderby: 'au_order,au_id'})

  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    if (item.private == 3) {
      data.dongs = [item]
      await clickView(data.dongs[0])
      return
    }
  }

  let items = []
  
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]
    
    if (item.private != 1) {
      continue
    }

    items.push(item)
  }
  
  data.dongs = items

  if (data.dongs.length > 0) {
    await clickView(data.dongs[0])
  }
}

onMounted(async () => {
  data.id = parseInt(route.params.id)
  data.apt = parseInt(route.params.apt)  
  
  util.loading(true)
  
  await initData()

  data.visible = false
  util.loading(false)
})

async function clickView(dong) {
  data.dong = dong

  let res = await model.find({peroidic:data.id, aptdong:dong.id})

  if (res.items.length == 0) {
    data.item = util.clone(item)
    data.original = util.clone(item)
  } else {
    data.item = res.items[0]
    data.original = util.clone(res.items[0])
  }
}

function clickUpdate() {
  data.item = util.clone(data.original)
  data.visible = true
}

function clickSubmit() {
  util.confirm('저장하시겠습니까', async function() {
    util.loading(true)

    let item = util.clone(data.item)

    item.periodic = data.id
    item.aptdong = data.dong.id

    if (item.id > 0) {
      await model.update(item)
    } else {
      item.id = await model.insert(item)
    }

    util.loading(false)
    util.info('저장되었습니다')

    data.original = util.clone(item)
    data.visible = false
  })
}
</script>

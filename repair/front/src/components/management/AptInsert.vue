<template>
  <el-dialog
    v-model="data.visible"
    title="주소 등록/수정"
    width="600px"
    :before-close="handleClose"
  >
    <y-table style="margin-top:-20px;">
            <y-tr>
            <y-th>ID</y-th>
            <y-td>
              <span v-if="data.item.id > 0">{{data.item.id}}</span>
            </y-td>
            </y-tr>
            <y-tr>
            <y-th>아파트명</y-th>
            <y-td>
              <el-input v-model="data.item.name" />
            </y-td>
          </y-tr>
          <y-tr>
            
            <y-th>준공년도</y-th>
            <y-td><el-input v-model="data.item.completeyear" /></y-td>
          </y-tr>
          <y-tr>

            <y-th>건축물형태</y-th>
            <y-td><el-input v-model="data.item.type" /></y-td>
          </y-tr>
          <y-tr>

            <y-th>전체동수</y-th>
            <y-td><el-input v-model="data.item.flatcount" /></y-td>

          </y-tr>
          <y-tr>

            <y-th>세대수</y-th>
            <y-td>
              <span style="width:60px;">{{data.item.familycount}} 세대</span> (상가&nbsp;<el-input class="date" style="width:35px;" v-model.number="data.item.familycount1" @keyup="onKeyup" />,
              오피&nbsp;<el-input class="date" style="width:35px;" v-model.number="data.item.familycount2" @keyup="onKeyup" />,
              아파트&nbsp;<el-input class="date" style="width:35px;" v-model.number="data.item.familycount3" @keyup="onKeyup" />
              )
            </y-td>
            
          </y-tr>
          <y-tr>
            <y-th>층수</y-th>
            <y-td>
              <el-input v-model="data.item.floor" />
            </y-td>
          </y-tr>
          <y-tr>

            <y-th>전화번호</y-th>
            <y-td>
              <el-input v-model="data.item.tel" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>팩스번호</y-th>
            <y-td>
              
              <el-input v-model="data.item.fax" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>공용메일주소</y-th>
            <y-td>
              <el-input v-model="data.item.email" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>담당자명</y-th>
            <y-td>
              
              <el-input v-model="data.item.personalname" />
            </y-td>
            
          </y-tr>
          <y-tr>
            <y-th>담당자 연락처</y-th>
            <y-td>
              
              <el-input v-model="data.item.personalhp" />
            </y-td>
            
          </y-tr>
          <y-tr>
            <y-th>담당자 메일</y-th>
            <y-td>
              
              <el-input v-model="data.item.personalemail" :rows="2" type="textarea" style="font-size:12px;" />
            </y-td>
            
          </y-tr>
          <y-tr>
            <y-th>우편번호</y-th>
            <y-td>
              <el-input v-model="data.item.zip" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>도로명주소</y-th>
            <y-td>
              <el-input v-model="data.item.address" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>지번주소</y-th>
            <y-td>
              
              <el-input v-model="data.item.address2" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>정밀점검일자</y-th>
            <y-td>
              <el-input v-model="data.item.testdate" />
            </y-td>
          </y-tr>
          <y-tr>

            <y-th>FMS 아이디</y-th>
            <y-td>
              <el-input v-model="data.item.fmsloginid" />
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>FMS 비번</y-th>
            <y-td>
              
              <el-input v-model="data.item.fmspasswd" />
            </y-td>
          </y-tr>
            </y-table>

    <template #footer>
      <el-button size="small" @click="clickClose">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, computed, watch, watchEffect } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Category, Repair, Facilitycategory } from "~/models"
import { useStore } from 'vuex'

const props = defineProps({
  id: Number,
  visible: Boolean,
  close: Function
})

const { width, height } = size()

const store = useStore()

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems(true)
}

const item = {
  id: 0,
  name: '',
  contracttype: 0,
  contractprice: '',
  completeyear: '',
  flatcount: '',
  type: '',
  floor: '',
  familycount: '',
  familycount1: 0,
  familycount2: 0,
  familycount3: 0,
  tel: '',
  fax: '',
  email: '',
  personalemail: '',
  zip: '',
  address: '',
  address2: '',
  testdate: '',
  nexttestdate: '',
  repair: '',
  safety: '',
  fault: '',
  contractdate: '',
  contractduration: '',
  invoice: '',
  depositdate: ''
}

const data = reactive({
  id: 0,  
  item: util.clone(item),
  check1: false,
  check2: false,
  check3: false,
  check4: false,
  check5: false,
  check6: false,
  check7: false,
  check8: false,
  check9: false,
  visible: false
})

async function initData() {
  let res = await Facilitycategory.find({orderby: 'fc_order'})

  data.facilitycategorys = [{id: 0, name: ' '}, ...res.items]
}

async function getItems(reset) {
  if (data.id != 0) {
    let res = await Apt.get(data.id)
    data.item = res.item
  }

  if (data.item.contracttype & 1) {
    data.check1 = true;
  } else {
    data.check1 = false;
  }

  if (data.item.contracttype & 2) {
    data.check2 = true;
  } else {
    data.check2 = false;    
  }

  if (data.item.contracttype & 4) {
    data.check3 = true;
  } else {
    data.check3 = false;    
  }

  if (data.item.contracttype & 8) {
    data.check4 = true;
  } else {
    data.check4 = false;    
  }

  if (data.item.contracttype & 16) {
    data.check5 = true;
  } else {
    data.check5 = false;    
  }

  if (data.item.contracttype & 32) {
    data.check6 = true;
  } else {
    data.check6 = false;    
  }

  if (data.item.contracttype & 64) {
    data.check7 = true;
  } else {
    data.check7 = false;    
  }

  if (data.item.contracttype & 128) {
    data.check8 = true;
  } else {
    data.check8 = false;    
  }

  if (data.item.contracttype & 256) {
    data.check9 = true;
  } else {
    data.check9 = false;    
  }
}

async function clickSubmit() {

  const item = data.item

  let contracttype = 0;
  
  if (data.check1 == true) {
    contracttype += 1;
  }

  if (data.check2 == true) {
    contracttype += 2;
  }

  if (data.check3 == true) {
    contracttype += 4;
  }

  if (data.check4 == true) {
    contracttype += 8;
  }

  if (data.check5 == true) {
    contracttype += 16;
  }

  if (data.check6 == true) {
    contracttype += 32;
  }

  if (data.check7 == true) {
    contracttype += 64;
  }

  if (data.check8 == true) {
    contracttype += 128;
  }

  if (data.check9 == true) {
    contracttype += 256;
  }

  item.contracttype = contracttype;
  
  if (item.name === '') {
    util.error('아파트명을 입력하세요')
    return    
  }

  if (item.contractype === 0) {
    util.error('계약구분을 선택하세요')
    return
  }
  
  let res
  let mode

  if (item.id === 0) {
    item.position = item.address
    res = await Apt.insert(item)
    mode = 'insert'
  } else {
    res = await Apt.update(item)
    mode = 'update'
  }

  if (res.code === 'ok') {
    util.info('등록되었습니다')
    
    data.visible = false

    if (props.close != null && props.close != undefined) {
      props.close(item)
    }
  } else {
    util.error('오류가 발생했습니다')
  }
}

function clickDelete() {
  util.confirm('삭제하시겠습니까', async function() {
    util.confirm('한번 삭제한 주소는 복구가 불가능합니다. 삭제하시겠습니까', async function() {
      let res = await Apt.remove(data.item)
      if (res.code === 'ok') {
        util.info('삭제되었습니다')
        done()
      }
    })
  })
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
  await initData()
  //getItems()
})

function clickClose() {
  data.visible = false
}

async function readData(id) {
  if (!util.isNull(id)) {
    data.id = id
  } else {
    data.id = 0
    data.item = util.clone(item)
  }

  if (data.id == 0 && !util.isNull(data)) {
    data.item = data
  }
  
  await getItems()
  data.visible = true
}

function insert(item) {
  data.id = 0
  if (!util.isNull(item)) {
    data.item = item
  }
  
  data.visible = true
}

function onKeyup() {
  let count = util.getInt(data.item.familycount1) + util.getInt(data.item.familycount2) + util.getInt(data.item.familycount3)

  data.item.familycount = String(count)
}

defineExpose({
  readData,
  insert
})

</script>
<style>
.date .el-input__wrapper {
  padding:0px;  
}

.date .el-input__inner {
  text-align:center;

}
</style>  

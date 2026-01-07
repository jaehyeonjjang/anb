<template>
  <Title title="비교견적업체 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <el-input v-model="search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>

    <TotalDiv :total="data.total" />
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :height="height(200)" v-infinite="getItems">
    <el-table-column prop="id" label="ID" width="80" align="center" />
    <el-table-column prop="name" label="회사명" />
    <el-table-column prop="ceo" label="대표" />
    <el-table-column label="구분" align="center">
      <template #default="scope">
          {{viewType(scope.row)}}
      </template>
    </el-table-column>
    <el-table-column width="80" label="기본" align="center">
      <template #default="scope">
          {{Comparecompany.getDefault(scope.row.default)}}
      </template>
    </el-table-column>
    <el-table-column prop="order" width="50" label="순번" align="center" />
    <!-- <el-table-column label="이미지" align="center" width="60">
         <template #default="scope">
         <el-image v-if="scope.row.image2 != ''" :src="util.getImagePath(scope.row.image2)" fit="cover" style="width:50px;height:50px;" :preview-src-list="[util.getImagePath(scope.row.image2)]"/>
         </template>
         </el-table-column>
         <el-table-column label="도장" align="center" width="60">
         <template #default="scope">
         <el-image v-if="scope.row.image != ''" :src="util.getImagePath(scope.row.image)" fit="cover" style="width:50px;height:50px;" :preview-src-list="[util.getImagePath(scope.row.image)]"/>
         </template>
         </el-table-column> -->
    <el-table-column label="" width="130" align="center" >
      <template #default="scope">
        <el-button size="small" @click="clickUpdate(scope.$index, scope.row)">수정</el-button>
        <!-- <el-button size="small" type="danger" @click="clickDelete(scope.$index, scope.row)">삭제</el-button> -->
      </template>
    </el-table-column>
  </el-table>  
  <div style="margin-top:10px;display:flex;justify-content: space-between;">
    <el-button size="small" type="success" @click="clickInsert">등록</el-button>
  </div>

  <el-dialog
    v-model="data.visible"
    title="비교견적업체 등록/수정"
    width="600px"
    :before-close="handleClose"
  >

    <y-table>
      <y-tr>
        <y-th style="width:70px;">구분</y-th>
        <y-td colspan="3">
          <el-checkbox size="small" label="장기수선계획" v-model="data.check1" style="font-size:12px;" />
          <el-checkbox size="small" label="정밀" v-model="data.check2" style="font-size:12px;" />
          <el-checkbox size="small" label="정기" v-model="data.check3" style="font-size:12px;" />
          <el-checkbox size="small" label="하자보수" v-model="data.check4" style="font-size:12px;" />
          <el-checkbox size="small" label="하자조사" v-model="data.check5" style="font-size:12px;" />
          <el-checkbox size="small" label="정밀안전진단" v-model="data.check6" style="font-size:12px;" />
          <el-checkbox size="small" label="감리" v-model="data.check7" style="font-size:12px;" />
          <el-checkbox size="small" label="기술자문" v-model="data.check8" style="font-size:12px;" />
          <el-checkbox size="small" label="순찰" v-model="data.check9" style="font-size:12px;" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>회사명</y-th>
        <y-td>
          <el-input v-model="data.item.name" />
        </y-td>
        <y-th style="width:70px;">대표자명</y-th>
        <y-td>
          <el-input v-model="data.item.ceo" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>주소</y-th>
        <y-td colspan="3">
          <div>
            <el-input v-model="data.item.address" />
          </div>
          <div style="margin-top:5px;">
            <el-input v-model="data.item.addressetc" />
          </div>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>전화번호</y-th>
        <y-td>
          <el-input v-model="data.item.tel" />
        </y-td>
        <y-th>FAX</y-th>
        <y-td>
          <el-input v-model="data.item.fax" />
        </y-td>
      </y-tr>
      <!-- <y-tr>
           <y-th>이미지</y-th>
           <y-td>
           <el-upload
           style="float:left;"
           class="upload-demo"
           ref="upload"
           :action="data.upload"
           :headers="headers"
           :limit="1"
           :on-exceed="handleExceed"
           :on-success="handelSuccess"
           :show-file-list="false"
           :auto-upload="true"
           >
           <el-button size="small" type="danger" @click="submitUpload(2)">이미지 업로드</el-button>

           </el-upload>
           </y-td>
           <y-th>도장</y-th>
           <y-td>
           <el-upload
           style="float:left;"
           class="upload-demo"
           ref="upload"
           :action="data.upload"
           :headers="headers"
           :limit="1"
           :on-exceed="handleExceed"
           :on-success="handelSuccess"
           :show-file-list="false"
           :auto-upload="true"
           >
           <el-button size="small" type="danger" @click="submitUpload(1)">이미지 업로드</el-button>

           </el-upload>
           </y-td>
           </y-tr> -->
      <y-tr>
        <y-th>기본</y-th>
        <y-td>
          <el-radio-group v-model="data.item.default">
            <el-radio :label="1" size="small">기본</el-radio>
            <el-radio :label="2" size="small">기본 아님</el-radio>
          </el-radio-group>        
        </y-td>
        <y-th>순번</y-th>
        <y-td>
          <el-input v-model="data.item.order" style="width:50px;" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>금액조정</y-th>
        <y-td colspan="3">
          <el-input v-model="data.item.adjust" style="width:100px;" /> 원
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>제경비</y-th>
        <y-td>
          <el-input v-model.number="data.item.financialprice" style="width:50px;" /> %
        </y-td>
        <y-th>기술료</y-th>
        <y-td>
          <el-input v-model.number="data.item.techprice" style="width:50px;" /> %
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>직접경비</y-th>
        <y-td>
          <el-input v-model.number="data.item.directprice" style="width:100px;" /> 원
        </y-td>
        <y-th>여비</y-th>
        <y-td>
          <el-input v-model.number="data.item.travelprice" style="width:100px;" /> 원
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>인쇄비</y-th>
        <y-td>
          <el-input v-model.number="data.item.printprice" style="width:100px;" /> 원
        </y-td>
        <y-th>차량운행비</y-th>
        <y-td>
          <el-input v-model.number="data.item.gasprice" style="width:100px;" /> 원
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>추가경비</y-th>
        <y-td>
          <el-input v-model.number="data.item.extraprice" style="width:100px;" /> 원
        </y-td>
        <y-th>위험수당</y-th>
        <y-td>
          <el-input v-model.number="data.item.dangerprice" style="width:50px;" /> %
        </y-td>
      </y-tr>
      <y-tr>
        <y-th></y-th>
        <y-td>
        </y-td>
        <y-th>기계기구손료</y-th>
        <y-td>
          <el-input v-model.number="data.item.machineprice" style="width:50px;" /> %
        </y-td>
      </y-tr>
    </y-table>


    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>  
</template>

<script setup lang="ts">

import { reactive, onMounted, ref } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Comparecompany } from "~/models"
import { useStore } from 'vuex'

const { width, height } = size()

const store = useStore()

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems(true)
}

const item = {
  id: 0,
  name: '',  
  address: '',
  addressetc: '',
  tel: '',
  fax: '',
  ceo: '',
  format: '',
  default: 2,
  image: '',
  image2: '',
  adjust: 0,
  financialprice: 0,
  techprice: 0,
  directprice: 0,
  printprice: 0,
  extraprice: 0,
  travelprice: 0,
  gasprice: 0,
  dangerprice: 0,
  machineprice: 0,
  remark: '',
  type: 0,
  order: 0,
  date: ''
}

const data = reactive({
  items: [],
  total: 0,
  page: 1,
  pagesize: 100,
  item: util.clone(item),
  visible: false,
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`,
  uploadtype: 1,
  check1: false,
  check2: false,
  check3: false,
  check4: false,
  check5: false,
  check6: false,
  check7: false,
  check8: false,
  check9: false,
})

async function initData() {  
}

async function getItems(reset) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }

  let res = await Comparecompany.find({page: data.page, pagesize: data.pagesize, name: search.text})

  if (res.items == undefined) {
    res.items = []
  }

  data.total = res.total
  data.items = data.items.concat(res.items)
}

function clickInsert() {  
  data.check1 = false
  data.check2 = false
  data.check3 = false
  data.check4 = false
  data.check5 = false
  data.check6 = false
  data.check7 = false
  data.check8 = false
  data.check9 = false
  
  data.item = util.clone(item)
  data.visible = true
}

function viewType(item) {
  let titles = []
  if (item.type & 1) {
    titles.push('장기수선계획') 
  }
  if (item.type & 2) {
    titles.push('정밀') 
  }
  if (item.type & 4) {
    titles.push('정기') 
  }
  if (item.type & 8) {
    titles.push('하자보수') 
  }
  if (item.type & 16) {
    titles.push('하자조사') 
  }
  if (item.type & 32) {
    titles.push('정밀안전진단') 
  }
  if (item.type & 64) {
    titles.push('감리') 
  }
  if (item.type & 128) {
    titles.push('기술자문') 
  }
  if (item.type & 256) {
    titles.push('순찰') 
  }

  return titles.join(', ')
}

function clickUpdate(pos, item) {
  data.check1 = false
  data.check2 = false
  data.check3 = false
  data.check4 = false
  data.check5 = false
  data.check6 = false
  data.check7 = false
  data.check8 = false
  data.check9 = false
  if (item.type & 1) {
    data.check1 = true
  }
  if (item.type & 2) {
    data.check2 = true
  }
  if (item.type & 4) {
    data.check3 = true
  }
  if (item.type & 8) {
    data.check4 = true
  }
  if (item.type & 16) {
    data.check5 = true
  }
  if (item.type & 32) {
    data.check6 = true
  }
  if (item.type & 64) {
    data.check7 = true
  }
  if (item.type & 128) {
    data.check8 = true
  }
  if (item.type & 256) {
    data.check9 = true
  }
  data.item = util.clone(item)
  data.visible = true
}

function clickDelete(pos, item) {
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Comparecompany.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      getItems(true)
    }
  })
}

async function clickSubmit() {
  const item = data.item
  if (item.name === '') {
    util.error('회사명을 입력하세요')
    return    
  }

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

  item.type = contracttype

  
  item.adjust = util.getInt(item.adjust)
  item.default = util.getInt(item.default)
  item.order = util.getInt(item.order)

  let res;

  if (item.id === 0) {
    res = await Comparecompany.insert(item)
  } else {
    res = await Comparecompany.update(item)
  }

  if (res.code === 'ok') {
    util.info('등록되었습니다')
    getItems(true)
    data.visible = false
  } else {
    util.error('오류가 발생했습니다')
  }
}

const handleClose = (done: () => void) => {
  util.confirm('팝업창을 닫으시겠습니까', function() {
    done()
  })  
}

onMounted(async () => {
  util.loading(true)
  
  await initData()
  await getItems()

  util.loading(false)
})

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
}

async function handelSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  if (data.uploadtype == 1) {
    data.item.image = response.filename
  } else {
    data.item.image2 = response.filename
  } 
}

const submitUpload = (uploadtype) => {
  data.uploadtype = uploadtype
  upload.value.clearFiles()
  upload.value!.submit()
}

</script>

<template>
  <Title title="기술자 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <el-input v-model="search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>

    <TotalDiv :total="data.total" />
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :height="height(200)" v-infinite="getItems">
    <el-table-column prop="id" label="ID" width="80" align="center" />
    <el-table-column prop="name" label="성명" />
    <el-table-column width="200" label="기술등급" align="center">
      <template #default="scope">
        <el-tag :type="Technician.getGradeType(scope.row.grade)">
          {{Technician.getGrade(scope.row.grade)}}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column label="도장" align="center" width="60">
      <template #default="scope">
        <el-image v-if="scope.row.stamp != ''" :src="util.getImagePath(scope.row.stamp)" fit="cover" style="width:50px;height:50px;" :preview-src-list="[util.getImagePath(scope.row.stamp)]"/>
      </template>
    </el-table-column>
    <el-table-column label="" width="200" align="center" >
      <template #default="scope">
        <el-button size="small" @click="clickUpdate(scope.$index, scope.row)">수정</el-button>
        <el-button size="small" type="danger" @click="clickDelete(scope.$index, scope.row)">삭제</el-button>
      </template>
    </el-table-column>
  </el-table>  
  <div style="margin-top:10px;display:flex;justify-content: space-between;">
    <el-button size="small" type="success" @click="clickInsert">등록</el-button>
  </div>

  <el-dialog
    v-model="data.visible"
    title="기술자 등록/수정"
    width="600px"
    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="80px">
      <el-form-item label="ID" v-show="data.item.id != 0">
        {{ data.item.id }}
      </el-form-item>
      <el-form-item label="성명">
        <el-input v-model="data.item.name" />
      </el-form-item>
      <el-form-item label="기술등급">
        <el-select v-model.number="data.item.grade" class="m-2" placeholder="기술등급">
          <el-option
            v-for="(item, index) in Technician.grades"
            :key="index"
            :label="item"
            :value="index"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="도장">

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
          <el-button size="small" type="danger" @click="submitUpload">이미지 업로드</el-button>

        </el-upload>
      </el-form-item>
      </el-form>

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
import { Technician } from "~/models"
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
  grade: 0,
  stamp: '',
  date: ''
}

const data = reactive({
  items: [],
  total: 0,
  page: 1,
  pagesize: 100,
  item: util.clone(item),
  visible: false,
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`
})

async function initData() {  
}

async function getItems(reset) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }

  let res = await Technician.find({page: data.page, pagesize: data.pagesize, name: search.text})

  if (res.items == undefined) {
    res.items = []
  }

  data.total = res.total
  data.items = data.items.concat(res.items)
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true
}

function clickUpdate(pos, item) {
  data.item = util.clone(item)
  data.visible = true
}

function clickDelete(pos, item) {
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Technician.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      getItems(true)
    }
  })
}

async function clickSubmit() {
  const item = data.item
  if (item.name === '') {
    util.error('성명을 입력하세요')
    return    
  }

  if (item.grade === 0) {
    util.error('기술등급을 선택하세요')
    return
  }
  
  let res;

  if (item.id === 0) {
    res = await Technician.insert(item)
  } else {
    res = await Technician.update(item)
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
  data.item.stamp = response.filename
}

const submitUpload = () => {
  upload.value.clearFiles()
  upload.value!.submit()
}

</script>

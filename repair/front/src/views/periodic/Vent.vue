<template>
  <Title title="환기구 덮개" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <el-table :data="data.items" border>
    <el-table-column prop="name" label="점검내용"  />
    <el-table-column prop="result" label="점검결과" align="center" width="100">
      <template #default="scope">
        <span v-if="scope.row.type == 1"><span v-if="scope.row.result == 1">양호</span><span v-else>보통</span></span>
        <span v-if="scope.row.type == 2"><span v-if="scope.row.result == 1">없음</span><span v-else>있음</span></span>
      </template>
    </el-table-column>
    <el-table-column prop="status" label="상태"  />
    <el-table-column prop="position" label="해당 위치" />
    <el-table-column label="이미지" align="center" width="80">
      <template #default="scope">
        <div style="height:30px;">
          <div v-if="scope.row.filename != ''">
            <el-image v-for="(item, index) in scope.row.filename.split(',')" style="width: 20px; height: 20px; top:4px;left:0px;position:relative;margin-right:5px;"
                      :src="util.getImagePath(item)"                    
                      fit="cover"
                      @click="clickPreviews(scope.row.filename, index)"
            />
          </div>          
        </div>
      </template>
    </el-table-column>        
  </el-table>  

  <div class="resulttext">
    <div style="margin-top:4px;">{{data.original.content2}}</div>    
  </div>

  
  <el-dialog
    v-model="data.visible"
    width="1000px"
  >


    <y-table>
      <y-tr>
        <y-th>점검내용</y-th>
        <y-th style="text-align:center;width:100px;">점검결과</y-th>
        <y-th style="width:200px;">상태</y-th>
        <y-th style="width:200px;">해당 위치</y-th>
        <y-th style="text-align:center;width:120px;">이미지</y-th>        
      </y-tr>
      <y-tr v-for="(item, index) in data.batchs" :key="item.id">
        <y-td>{{data.batchs[index].name}}</y-td>
        <y-td style="text-align:center;">
          <el-radio-group v-model.number="data.batchs[index].result">
            <el-radio-button size="small" label="1"><span v-if="data.batchs[index].type == 1">양호</span><span v-else>없음</span></el-radio-button>
            <el-radio-button size="small" label="2"><span v-if="data.batchs[index].type == 1">보통</span><span v-else>있음</span></el-radio-button>              
          </el-radio-group>
        </y-td>
        <y-td>
          <el-input v-model="data.batchs[index].status" />
        </y-td>
        <y-td>
          <el-input v-model="data.batchs[index].position" />
        </y-td>
        <y-td style="text-align:center;">
          <div style="height:30px;">
            <el-upload
              accept="image/jpeg,image/png"
              ref="upload"
              style="display:block;float:left;margin:5px 10px 0px 5px;"
              :data="{path:'periodic'}"
              :action="data.upload"
              :headers="headers"
              :limit="1"
              :on-exceed="(file, files) => handleExceed(file, files, index)"
              :on-success="(res, file, files) => handleSuccess(res, file, files, index)"
              :show-file-list="false"
              :auto-upload="true"
              :before-upload="beforeImageUpload"
            >

              
              <el-icon style="font-size:20px;"><Plus /></el-icon>
            </el-upload>

            
            <div v-if="data.batchs[index].filename != ''" style="float:left;">
              <el-image v-for="(item, pos) in data.batchs[index].filename.split(',')" style="width: 20px; height: 20px; top:4px;left:0px;position:relative;margin-right:5px;"
                        :src="util.getImagePath(item)"                    
                        fit="cover"
                        @click="clickImageDelete(data.batchs[index].filename, index, pos)"
              />
            </div>          
          </div>

        </y-td>

      </y-tr>
    </y-table>



  <div style="margin-top:5px;">
    <el-input v-model="data.item.content2" :rows=2 type="textarea" />
  </div>


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
import { Periodicother, Periodicotheretc } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'
import type { UploadInstance } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodicother

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const item = {
  id: 0,
  name: '',
  type: 0,
  result: 0,
  status: '',
  position: '',
  filename: '',
  offlinefilename: '',
  category: 0,
  order: 0,
  periodic: 0,
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  original: {content2: ''},
  visible: false,
  items: [],
  batchs: [],
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`
})

async function initData() {
}

async function getItems() {
  let res = await model.find({
    periodic: data.id,
    category: 2,
    orderby: 'po_order,po_id'    
  })

  if (res.items == null) {
    res.items = []
  }
  
  data.items = res.items

  data.uploads = new Array(res.items.length)
  
  res = await Periodicotheretc.getByPeriodic(data.id)
  data.original = res.item
}

async function clickUpdate(pos) {
  let res = await Periodicotheretc.getByPeriodic(data.id)
  const item = res.item
  
  data.item = util.clone(item)

  data.batchs = util.clone(data.items)
  data.visible = true
}

async function clickSubmit(type) {
  util.loading(true)

  for (let i = 0; i < data.batchs.length; i++) {
    let item = util.clone(data.batchs[i])

    item.result = util.getInt(item.result)
    await model.update(item)
  }
  
  let res = await Periodicotheretc.getByPeriodic(data.id)
  let item = res.item

  item.content2 = data.item.content2
  
  await Periodicotheretc.update(item)

  data.items = util.clone(data.batchs)

  data.original = item
  data.visible = false
  
  util.info('수정되었습니다')
  util.loading(false)
}

onMounted(async () => {
  util.loading(true)
  
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  await initData()
  await getItems()

  util.loading(false)
})

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles, index) => {  
}

function handleSuccess(response: any, file: UploadFile, files: UploadFiles, index: number) {
  const filename = response.filename
  
  let items = util.clone(data.batchs)
  items[index].filename = filename
  
  data.batchs = items

  for (let i = 0; i < upload.value!.length; i++) {
    upload.value![i].clearFiles(['success'])
  }  
}

function beforeImageUpload(file) {
  const isImage = (file.type === 'image/jpeg' || file.type === 'image/png')

  if (!isImage) {
    util.error('이미지 파일만 업로드 가능합니다 (jpg, png)')
  }

  return isImage
}

function clickPreviews(str, index) {
  const imgs = str.split(',').map(item => util.getImagePath(item)) 
  v3ImgPreviewFn({images:imgs, index: index})  
}

function clickImageDelete(item, index, pos) {
  util.confirm('삭제하시겠습니까', function() {
    let items = util.clone(data.batchs)
    let temp = items[index].filename.split(',')
    temp.splice(pos, 1)

    items[index].filename = temp.join(',')
    data.batchs = items
    util.info('삭제되었습니다')
  })
}
</script>

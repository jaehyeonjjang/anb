<template>
  <Container :height="80">  
    <Title title="도면" />
    
    <div style="margin-bottom:10px;">
      <el-button style="display:block;float:left;" size="small" type="primary" @click="clickSave">저장</el-button>
      <el-button style="display:block;float:left;margin-left:10px;" size="small" @click="clickCancel">취소</el-button>

      <el-button style="display:block;float:right;" size="small" type="warning" @click="clickCopyFloor">층 복사</el-button>
      <el-button style="display:block;float:right;" size="small" type="warning" @click="clickCopyDong">동 복사</el-button>
      <div style="clear:both;"></div>
    </div>

    <div style="overflow:auto;border:1px solid #ccc;padding:10px 10px;" :style="{height: height(190)}">
      <div v-for="(item, index) in data.items" :key="item.id" class="block" :style="{marginLeft: (item.level - 1) * 40 + 'px'}">
        <div style="flex:1;height:19px;padding-top:5px;" @click="clickPreview(item)">{{item.name}}</div>
        <el-upload
          v-if="item.upload == 1"
          accept="image/jpeg,image/png"
          ref="upload"
          style="display:block;"
          :data="{path:'blueprint'}"
          :action="data.upload"
          :headers="headers"
          :limit="1"
          :on-exceed="(file, files) => handleExceed(file, files, index)"
          :on-success="(res, file, files) => handleSuccess(res, file, files, index)"
          :show-file-list="false"
          :auto-upload="true"
          :before-upload="beforeImageUpload"
        >


          <img v-if="item.filename" :src="util.getImagePath(item.filename)" style="height:20px;margin-right:2px;" class="plan">
          <el-icon v-else class="btn"><Picture /></el-icon>
        </el-upload>

        <div class="clear" />
      </div>
    </div>

    <el-dialog title="동 복사" v-model="data.visibleCopyDong">
      <div style="margin:0px auto 10px auto;width:560px;">
            <el-select v-model.number="data.dong" style="width:100%;">
              <el-option v-for="item in data.alldongs" :key="item.key" :label="item.label" :value="item.key" />
            </el-select>
      </div>

      <el-transfer v-model="data.dongs" :data="data.alldongs" />


      <template #footer>
        <el-button size="small" @click="data.visibleCopyDong = false">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmitCopyDong">등록</el-button>
      </template>
    </el-dialog>

    <el-dialog title="층 복사" v-model="data.visibleCopyFloor">
      <div style="margin:0px auto 10px auto;width:560px;">
            <el-select v-model.number="data.dong" style="width:100%;" @change="changeDong"> 
              <el-option v-for="item in data.alldongs" :key="item.key" :label="item.label" :value="item.key" />
            </el-select>
      </div>


      <div style="margin:0px auto 10px auto;width:560px;">
        <el-select v-model.number="data.floor" style="width:100%;">
          <el-option v-for="item in data.floors" :key="item.key" :label="item.label" :value="item.key" />
        </el-select>
      </div>

      <el-transfer v-model="data.targetfloors" :data="data.floors" />


      <template #footer>
        <el-button size="small" @click="data.visibleCopyFloor = false">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmitCopyFloor">등록</el-button>
      </template>
    </el-dialog>    

    <el-dialog v-model="data.visiblePreview">
    <el-image      
      :src="data.previewUrl"      
      fit="cover"
    />
    
  </el-dialog>

  </Container>


</template>

<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted, computed, watch, getCurrentInstance } from "vue"
import { util, size }  from "~/global"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import type { UploadInstance } from 'element-plus'
import { Blueprint } from "~/models"
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size();
const store = useStore()
const route = useRoute()

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const data = reactive({
  apt: 0,
  masters: [],
  summarytypes: [],
  allsubmasters: [],
  submasters: [],
  allusers: [],
  users: [],
  prefix: '',
  postfix: '',
  start: '',
  end: '',
  imageId: -1,
  loading: true,
  tempRoute: {},
  aptgroups: [],
  imageUrl: '',
  items: [],
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`,
  floortypes: [{id: 0, name: ' ' }, {id: 1, name: '지하'}, {id: 2, name: '지상'}, {id: 3, name: '옥상'}],
  visiblePreview: false,
  visibleCopyDong: false,
  visibleCopyFloor: false,
  previewUrl: '',
  dong: null,
  dongs: [],
  alldongs: [],
  floor: null,
  floors: []
})

function initData() {
  let items = []
  items.push(getImage())
  data.items = items
}

async function getItems() {
  let res = await Blueprint.find({apt: data.apt, category: 1, orderby: 'bp_parentorder,bp_order desc,bp_id'})
  if (res.items == null) {
    res.items = []
  }

  data.items = res.items

  data.uploads = new Array(res.items.length)

  let dongs = []

  res.items.filter(c => c.level == 1).map((item) => {

    dongs.push({
      key: item.id,
      label: item.name,
      disabled: false
    })
  })

  data.alldongs = dongs

}

onMounted(async () => {
  data.apt = parseInt(route.params.apt)
  
  util.loading(true)
  
  await initData()
  await getItems()  
  
  util.loading(false)
})

function clickSave() {
  util.confirm('저장하시겠습니까', async function() {
    util.loading(true)    

    for (let i = 0; i < data.items.length; i++) {
      const item = data.items[i]

      await Blueprint.update(item)
    }

    util.loading(false)    
    util.info('저장되었습니다')
  })
}

function clickCancel() {
  util.confirm('취소하시겠습니까', async function() {
    await getItems()    
  })
}

function getImage() {
  data.imageId--
  return {
    id: data.imageId,
    name: '',
    level: 1,
    parent: 0,    
    floortype: '0',
    filename: '',
    upload: 0,
    parentorder: 0,
    order: 0,
    aptdong: 0,
    apt: 0,
    date: ''
  }
}

function getImageEnd(pos) {
  const item = data.items[pos]
  let end = pos
  for (let i = pos + 1; i < data.items.length; i++) {
    if (data.items[i].level <= item.level) {
      break
    }

    end = i
  }

  return end
}

function clickImageInsert(pos) {
  const image = getImage()

  if (pos > 0) {
    image.level = data.items[pos - 1].level
  }
  let items = data.items
  items.splice(pos, 0, image)

  data.items = items
}

function clickImageDelete(pos) {
  util.confirm('삭제하시겠습니까', async function() {
    const end = getImageEnd(pos)
    data.items.splice(pos, end - pos + 1)
  })
}

function clickImageDown(pos) {  
  if (pos >= data.items.length - 1) {
    return
  }

  const temp = data.items[pos]
  data.items[pos] = data.items[pos + 1]
  data.items[pos + 1] = temp
}

function clickImageUp(pos) {
  if (pos === 0) {
    return
  }

  const temp = data.items[pos]
  data.items[pos] = data.items[pos - 1]
  data.items[pos - 1] = temp
}

function clickImageLeft(pos) {
  const item = data.items[pos]

  if (item.level <= 1) {
    return
  }

  const end = getImageEnd(pos)

  for (let i = pos; i <= end; i++) {
    data.items[i].level--
  }
}

function clickImageRight(pos) {
  const item = data.items[pos]

  const max = 4

  if (pos === 0) {
    return
  }

  if (item.level >= max) {
    return
  }

  if (pos > 0) {
    if (item.level - data.items[pos - 1].level >= 1) {
      return
    }
  }

  let end = pos
  for (let i = pos + 1; i < data.items.length; i++) {
    if (data.items[i].level <= item.level) {
      break
    }

    if (data.items[i].level >= max) {
      return
    }

    end = i
  }

  for (let i = pos; i <= end; i++) {
    data.items[i].level++
  }
}

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles, index) => {  
}

function handleSuccess(response: any, file: UploadFile, files: UploadFiles, index: number) {
  const filename = response.filename
  
  let items = util.clone(data.items)
  items[index].filename = filename
  
  data.items = items

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

async function clickCopy() {
  const id = this.$route.params && this.$route.params.id
  util.confirm('삭제하시겠습니까', async function() {
    await request({
      method: 'POST',
      url: '/api/apt/' + id + '/copy'
    })

    util.info('복사되었습니다')

    getList()
  })
}

function clickPreview(item) {
  if (item.filename == '') {
    return
  }

  v3ImgPreviewFn(util.getImagePath(item.filename))   
}

function clickCopyFloor() {
  data.dong = null
  data.floor = null
  data.floors = []
  data.targetfloors = []
  data.visibleCopyFloor = true
}

function clickCopyDong() {
  data.dong = null
  data.dongs = []
  data.visibleCopyDong = true
}

function clickSubmitCopyDong() {
  if (data.dong == null || data.dong == 0) {
    util.error('동을 선택하세요')
    return
  }

  if (data.dongs.length == 0) {
    util.error('복사할 동을 선택하세요')
    return
  }

  let sources = data.items.filter(item => item.parent == data.dong)


  for (let i = 0; i < sources.length; i++) {
    let source = sources[i]

    for (let j = 0; j < data.dongs.length; j++) {
      let dong = data.dongs[j]

      if (dong == data.dong) {
        continue
      }

      for (let k = 0; k < data.items.length; k++) {
        let item = data.items[k]

        if (item.parent != dong) {
          continue
        }

        if (source.name == item.name) {
          data.items[k].filename = source.filename
          break
        }
      }
    }
  }

  util.info('복사되었습니다')
  data.visibleCopyDong = false
}

function clickSubmitCopyFloor() {
  if (data.dong == null || data.dong == 0) {
    util.error('동을 선택하세요')
    return
  }

  if (data.floor == null || data.floor == 0) {
    util.error('층을 선택하세요')
    return
  }

  if (data.floors.length == 0) {
    util.error('복사할 층을 선택하세요')
    return
  }

  let source
  
  for (let k = 0; k < data.items.length; k++) {
    let item = data.items[k]

    if (item.id == data.floor) {
      source = item
      break
    }
  }

  let items = util.clone(data.items)
  
  for (let j = 0; j < data.targetfloors.length; j++) {
    let target = data.targetfloors[j]

    for (let k = 0; k < data.items.length; k++) {
      let item = data.items[k]

      if (item.id != target) {
        continue
      }

      items[k].filename = source.filename
      break
    }
  }
  
  data.items = items

  util.info('복사되었습니다')
  data.visibleCopyFloor = false  
}

function changeDong() {
  let items = []
  for (let i = 0; i < data.items.length; i++) {
    let item = data.items[i]
    
    if (item.parent == data.dong) {
      items.push({
        key: item.id,
        label: item.name
      })
    }
  }

  data.floors = items
}
</script>

<style lang="scss" scoped>
.block {
  border: 1px solid #aaa;
  border-radius: 3px;
  padding: 5px 10px;
  margin-bottom: 5px;
  text-align: left;
  font-weight: bold;  
  display:flex;
  flex-direction:row;
  justify-content: space-between;
  
  .input {
    display:block;
    float:left;
    width:400px;
    margin-right: 10px;
    padding: 2px 5px;
  }

  .btn {
    cursor: hand;
    cursor: pointer;
    display:block;
    float:right;
    font-size:20px;
    margin: 2px 5px 0px 5px;
  }

  .clear {
    clear:both;
  }
}

.block:hover {
  background:#eeeeee;
}

.plan-uploader {
  float:right;
  margin-top: 4px;
  margin-right: 20px;
  height: 34px;

  .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;    
  }

  .el-upload:hover {
    border-color: #409EFF;
  }
}

.plan-uploader-icon {
  font-size: 24px;
  color: #8c939d;
  width: 32px;
  height: 32px;
  line-height: 32px;
  text-align: center;
}
.plan {
  width: 24px;
  height: 24px;
  display: block;
}


</style>

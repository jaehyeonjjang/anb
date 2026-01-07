<template>
  <Title title="수량" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert">등록</el-button>
      <el-button size="small" type="primary" @click="clickInsertMulti">일괄 등록</el-button>
    </div>
  </div>


  <el-table :data="data.items" border :height="height(170)" ref="listRef" @row-click="clickUpdate" v-infinite="getItems" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" width="40" align="center" />
    <el-table-column prop="extra.apt.name" label="아파트명" align="left" v-if="data.apt == 0" />
    <el-table-column prop="name" label="명칭" align="left" />
    <!-- <el-table-column label="이미지" align="center">
         <template #default="scope">
         <el-image :src="util.getImagePath(scope.row.filename)" fit="cover" style="position:relative;top:5px;width:75px;height:50px;" @click="v3ImgPreviewFn(util.getImagePath(scope.row.filename))" />
         </template>
         </el-table-column> -->
    <!--<el-table-column prop="order" label="순번" align="center" width="50" />-->
    <el-table-column prop="date" label="등록일" align="center" width="150" />
    <el-table-column label="" align="center" width="80">
      <template #default="scope">
        <el-button size="small" type="primary" @click="clickDraw(scope.row)">입력</el-button>
      </template>
    </el-table-column>
  </el-table>


  <el-dialog
    v-model="data.visible"
    width="800px"
  >

      <y-table>
        <y-tr v-if="data.apt == 0">
          <y-th>아파트명</y-th>
          <y-td>
            <div style="display:flex;justify-content: space-between;">
              <div style="margin-top:3px;">{{data.selectapt.name}}</div>
              <el-button size="small" type="primary" @click="clickApt()">아파트 검색</el-button>
            </div>
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>명칭</y-th>
          <y-td>
            <el-input v-model="data.item.name" />
          </y-td>
        </y-tr>
        <!--
        <y-tr>
          <y-th>순번</y-th>
          <y-td>
            <el-input v-model.number="data.item.order" />
          </y-td>
        </y-tr>
        -->
        <y-tr>
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
              :on-success="handleSuccess"
              :show-file-list="true"
              :auto-upload="true"
              :data="{path:'area'}"
              :accept="'image/*'"
            >
              <el-button size="small" type="danger" @click="submitUpload">이미지 업로드</el-button>

            </el-upload>


          </y-td>
        </y-tr>
      </y-table>

      <template #footer>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">저장</el-button>
      </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleMulti"
    width="900px"
  >

      <y-table>
        <y-tr v-if="data.apt == 0">
          <y-th style="width:60px;">아파트명</y-th>
          <y-td>
            <div style="display:flex;justify-content: space-between;">
              <div style="margin-top:3px;">{{data.selectapt.name}}</div>
              <el-button size="small" type="primary" @click="clickApt()">아파트 검색</el-button>
            </div>
          </y-td>
        </y-tr>
      </y-table>
      <y-table style="margin-top:10px;">
        <y-tr v-for="(item, index) in data.names">
          <y-th style="width:60px;">명칭</y-th>
          <y-td style="width:300px;">
            <el-input v-model="data.names[index]" />
          </y-td>
          <y-th style="width:60px;">이미지</y-th>
          <y-td>

            <el-upload
              style="float:left;"
              class="upload-demo"
              ref="uploads"
              :action="data.upload"
              :headers="headers"
              :limit="1"
              :on-exceed="handleExceedMulti"
              :on-success="handleSuccessMulti"
              :show-file-list="true"
              :auto-upload="true"
              :data="{path:'area', param: data.uploadIndex}"
              :accept="'image/*'"
            >
              <el-button size="small" type="danger" @click="submitUploads(index)">이미지 업로드</el-button>

            </el-upload>


          </y-td>
        </y-tr>
      </y-table>

      <template #footer>
        <el-button size="small" @click="clickCancelMulti">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmitMulti">저장</el-button>
      </template>
  </el-dialog>


  <div style="position:absolute;left:0px;top:0px;background:#FFF;z-index:999;" :style="{width: width(0), height: height(0)}" v-show="data.visibleDraw">
    <div style="display:flex;flex-direction: row;justify-content: space-between;padding: 10px 10px;">
      <span style="font-size:16px;font-weight:bold;">{{getAptName()}}</span>
      <el-radio-group v-model.number="data.type" @change="changeType">
        <el-radio-button size="small" label="1">기준</el-radio-button>
        <el-radio-button size="small" label="2">면적</el-radio-button>
        <el-radio-button size="small" label="3">연속길이</el-radio-button>
        <el-radio-button size="small" label="4">면적 + 연속길이</el-radio-button>

        <div style="width:50px;"></div>
        <el-button size="small" @click="clickZoomUp"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" @click="clickZoomDown"><el-icon><Minus /></el-icon></el-button>

        <div style="width:50px;"></div>
        <el-button size="small" @click="clickCut"><el-icon><Scissor /></el-icon></el-button>        

        <div style="width:50px;"></div>
        <el-button size="small" @click="clickUndo" :disabled="data.undos.length == 0"><el-icon><ArrowLeft /></el-icon></el-button>
        <el-button size="small" @click="clickRedo" :disabled="data.redos.length == 0"><el-icon><ArrowRight /></el-icon></el-button>

        <div style="width:50px;"></div>
        <el-checkbox size="small" label="SHIFT" v-model="data.usershift" @change="changeUsershift" />
      </el-radio-group>

      <div>
      <el-button size="small" @click="clickCancelDraw">닫기</el-button>
      <el-button size="small" @click="clickSubmitDraw" type="primary" >저장</el-button>
      </div>
    </div>

    <div id="canvasFrame" style="margin-left:10px;border:1px solid #999;text-align:left;overflow:auto;" :style="{width: width(20), height: height(70)}">
      <canvas id="canvas" @mousedown.left="clickMouseLeft" @mousedown.right="clickMouseRight" @mousemove="moveMouse" @contextmenu.prevent></canvas>
    </div>

    <div style="position:absolute;background:white;overflow:auto;" :style="{left: getPosition(), top: '50px', height: height(90)}">
      <y-table>
        <y-tr>
          <y-th style="width:30px;text-align:center;"><span @click="clickPosition">NO</span></y-th>
          <y-th style="width:120px;">길이</y-th>
          <y-th>삭제</y-th>
        </y-tr>
        <y-tr v-for="(item, index) in lengths">
          <y-td style="text-align:center;">
            <span v-if="index==0">기준</span>
            <span v-else>{{item[0].index+1}}</span>
          </y-td>
          <y-td style="text-align:right;" v-if="index==0">
            <el-input v-model="data.item.standard" style="text-align:right;" @keypress="inputKey" />
          </y-td>
          <y-td style="text-align:right;" v-if="index!=0">
            <span v-html="getValue(item)" />
          </y-td>
          <y-td style="text-align:center;">
            <el-button size="small" type="danger" @click="clickDeletePoint(item[0].index)"><el-icon><Delete /></el-icon></el-button>
          </y-td>
        </y-tr>
        <y-tr v-if="lengths.length > 1">
          <y-td style="text-align:center;">합계</y-td>
          <y-td style="text-align:right;">
            {{getTotal()}}
          </y-td>
          <y-td style="text-align:center;">

          </y-td>
        </y-tr>
      </y-table>


      <y-table style="margin-top:10px;">
        <y-tr>
          <y-th style="width:30px;text-align:center;"><span @click="clickPosition">NO</span></y-th>
          <y-th style="width:120px;">면적</y-th>
          <y-th>삭제</y-th>
        </y-tr>
        <y-tr v-for="(item, index) in areas">
          <y-td style="text-align:center;">{{item[0].index+1}}</y-td>
          <y-td style="text-align:right;">
            <span v-html="getArea(item)" />
          </y-td>
          <y-td style="text-align:center;">
            <el-button size="small" type="danger" @click="clickDeletePoint(item[0].index)"><el-icon><Delete /></el-icon></el-button>
          </y-td>
        </y-tr>
        <y-tr v-if="areas.length > 0">
          <y-td style="text-align:center;">합계</y-td>
          <y-td style="text-align:right;">
            {{getTotalArea()}}
          </y-td>
          <y-td style="text-align:center;">

          </y-td>
        </y-tr>
      </y-table>
    </div>
  </div>

  <el-dialog
    v-model="data.visibleApt"
    width="800px"
  >

    <div style="display:flex;gap: 10px;margin-bottom:10px;">
      <el-input v-model="data.aptname" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearchApt" />

      <el-button size="small" class="filter-item" type="primary" @click="clickSearchApt">검색</el-button>
    </div>


    <el-table :data="data.apts" border :height="'300px'" v-infinite="getApts" @row-click="selectApt">

      <el-table-column prop="name" label="아파트명">
        <template #default="scope">
          <div style="color:#000;" >{{scope.row.name}}</div>
        </template>
      </el-table-column>

      <el-table-column prop="tel" label="전화번호" width="100">
        <template #default="scope">
          <span>{{scope.row.tel}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="address" label="도로명주소">
        <template #default="scope">
          <span class="value">{{scope.row.address}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="address2" label="지번주소">
        <template #default="scope">
          <span class="value">{{scope.row.address2}}</span>
        </template>
      </el-table-column>


    </el-table>

  </el-dialog>

</template>


<script setup lang="ts">
const props = defineProps({
  apt: Number,
})

import { ref, reactive, onMounted, onUnmounted, computed } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Repairarea, Apt } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Repairarea

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const canvas = ref(null);
const ctx = ref(null);

const item = {
  id: 0,
  name: '',
  filename: '',
  order: 0,
  apt: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 50,
  aptpage: 1,
  aptpagesize: 20,
  item: util.clone(item),
  visible: false,
  visibleMulti: false,
  visibleDraw: false,
  visibleApt: false,
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`,
  filename: '',
  start: true,
  points: [],
  type: 1,
  zoom: 1.0,
  undos: [],
  redos: [],  
  position: 2,
  aptname: '',
  selectapt: {
    id: 0,
    name: ''
  },
  shift: true,
  usershift: true,
  saved: true,
  names: ['', '', '', '', '', '', '', '', '', ''],
  filenames: ['', '', '', '', '', '', '', '', '', ''],
  uploadIndex: 0,
})

async function initData() {
}

function changeUsershift() {
  data.shift = data.usershift
}

async function getItems(reset) {
  if (reset == true) {
    data.items = []
    data.page = 1
  }
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    periodic: data.id,
    apt: data.apt,
    orderby: 'a_name, ra_name'
  })

  // if (res.items != null) {
  //   for (let i = 0; i < res.items.length; i++) {
  //     res.items[i].index = i + 1
  //   }
  // }

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }

  let items = util.clone(data.items).concat(res.items)
  for (let i = 0; i < items.length; i++) {
    items[i].index = i + 1
  }

  data.items = items
  data.page++
}

function clickInsert() {
  data.item = util.clone(item)
  data.visible = true
}

function clickInsertMulti() {
  data.item = util.clone(item)

  data.names = ['', '', '', '', '', '', '', '', '', '']
  data.filenames = ['', '', '', '', '', '', '', '', '', '']
  data.visibleMulti = true
}

onMounted(async () => {
  canvas.value = document.getElementById("canvas")
  ctx.value = canvas.value.getContext("2d")

  data.apt = util.getInt(route.params.apt)

  util.loading(true)

  await initData()
  await getItems(true)

  data.visible = false
  util.loading(false)
})

function clickCancel() {
  data.visible = false
}

function removeEvent() {
  window.removeEventListener('keyup', keyUp)
  window.removeEventListener('keydown', keyDown)
}

async function clickCancelDraw() {
  if (data.saved == false) {
    util.confirm('작업을 취소하시겠습니까', async function() {
      await getItems(true)
      data.visibleDraw = false
      removeEvent()
    })
  } else {
    await getItems(true)
    data.visibleDraw = false
    removeEvent()
  }
}

const listRef = ref<InstanceType<typeof ElTable>>()
const listSelection = ref([])
const toggleListSelection = (rows) => {
  if (rows) {
    rows.forEach((row) => {
      listRef.value!.toggleRowSelection(row, undefined)
    })
  } else {
    listRef.value!.clearSelection()
  }
}
const changeList = (val) => {
  listSelection.value = val
}

function clickDeleteMulti() {
  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)

    let items = []
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id
      }

      await model.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems(true)

    util.loading(false)
  })
}

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
}

async function handleSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  const filename = response.filename
  data.filename = filename
}

const submitUpload = () => {
  upload.value.clearFiles()
  upload.value!.submit()
}

const uploads = ref([])

const handleExceedMulti: UploadProps['onExceed'] = (files, uploadFiles) => {
}

async function handleSuccessMulti(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  const filename = response.filename
  data.filenames[util.getInt(response.param)] = filename
}

const submitUploads = (index) => {
  data.uploadIndex = index
  uploads.value[index].clearFiles()
  uploads.value![index].submit()
}

async function clickSubmit() {
  if (data.apt > 0) {
    data.selectapt.id = data.apt
  }

  if (data.selectapt.id == 0) {
    util.error('아파트를 선택하세요')
    return
  }

  util.loading(true)

  let item = util.clone(data.item)
  item.apt = data.selectapt.id
  item.order = util.getInt(item.order)
  item.filename = data.filename
  await model.insert(item)

  await getItems(true)
  util.info('등록되었습니다')
  util.loading(false)

  data.visible = false
}

function repaint() {
  ctx.value.beginPath()
  ctx.value.clearRect(0, 0, canvas.value.width, canvas.value.height)
  ctx.value.lineWidth = 1
  ctx.value.fillStyle = 'rgb(255, 255, 255)'
  ctx.value.strokeStyle = 'rgb(255, 0, 0)'
  ctx.value.textAlign = "center"
  ctx.value.font = '20px'

  let color = 'rgb(255, 0, 0)'
  ctx.value.drawImage(data.img, 0, 0, data.imgWidth * data.zoom, data.imgHeight * data.zoom)
  ctx.value.beginPath()

  for (let i = 0; i < data.points.length; i++) {
    let item = data.points[i]

    let typeid = item[0].type
    
    if (typeid == 1) {
      color = 'rgb(255, 0, 255)'
    } else if (typeid == 2) {
      color = 'rgb(0, 0, 255)'
    } else if (typeid == 3) {
      color = 'rgb(0, 200, 0)'
    } else if (typeid == 4) {
      color = 'rgb(255, 0, 0)'
    }

    ctx.value.strokeStyle = color
    
    ctx.value.beginPath()

    for (let j = 1; j < item.length; j++) {
      let point = item[j]
      let x = point.x * data.zoom
      let y = point.y * data.zoom

      if (data.start == false) {
        if (item.length > 1 && j == item.length - 1 && i == data.points.length - 1 && data.shift == true) {
          if (Math.abs(item[j - 1].y - y) > Math.abs(item[j - 1].x - x)) {
            x = item[j - 1].x * data.zoom
          } else {
            y = item[j - 1].y * data.zoom
          }
        }
      }

      if (j == 1) {
        if (data.start == false && i == data.points.length - 1) {
          ctx.value.arc(x, y, 5, 0, 2 * Math.PI)
        }

        ctx.value.moveTo(x, y)
      } else {
        ctx.value.lineTo(x, y)
      }
    }

    ctx.value.stroke()

    if (data.start == false && i == data.points.length - 1) {
    } else {
      ctx.value.fillStyle = 'rgb(255, 255, 255)'
      ctx.value.beginPath()
      let cx = (item[1].x - 10) * data.zoom
      let cy = (item[1].y - 10) * data.zoom
      ctx.value.arc(cx, cy, 8, 0, 2 * Math.PI)
      ctx.value.fill()
      ctx.value.stroke()
      ctx.value.fillStyle = color
      ctx.value.fillText((i+1), cx, cy+3)
    }
  }
}

function clickDraw(item) {
  data.points = []

  data.item = util.clone(item)

  if (item.content != '') {
    data.points = JSON.parse(item.content)
  }

  canvas.value.width = parseInt(width(20).replace('px', ''))
  canvas.value.height = parseInt(height(70).replace('px', ''))

  if (data.item.standard == 0) {
    data.item.standard = 10  
  }

  let img = new Image()
  img.src = util.getImagePath(item.filename)
  img.onload = function() {
    data.imgWidth = img.width
    data.imgHeight = img.height

    canvas.value.width = img.width
    canvas.value.height = img.height

    data.img = img
    repaint()
    
    if (data.points.filter(item => item[0].type == 1).length > 0) {
      data.type = 4
    } else {
      data.type = 1
    }
    
    window.addEventListener('keyup', keyUp)
    window.addEventListener('keydown', keyDown)

    data.visibleDraw = true
  }
}

function clickUpdate(row, index) {
  if (index.no == 0 || index.no == 5) {
    return
  }

  clickDraw(row)
}

let _lastDate = new Date()

function clickMouseLeft(e) {  
  _lastDate = new Date()
  
  let margin = document.getElementById("canvas").getBoundingClientRect()
  let x = (e.clientX - margin.x + 1) / data.zoom
  let y = (e.clientY - margin.y + 1) / data.zoom
  let point = {x:x, y:y}

  let points = util.clone(data.points)

  if (data.start) {
    points.push([{type: data.type}, point, point])

    data.start = false
  } else {
    let start = points[points.length - 1][1]

    if (data.type == 2 || data.type == 4) {
      if (Math.sqrt(Math.pow(Math.abs(start.x - x), 2) * Math.pow(Math.abs(start.y - y), 2)) < 5) {
        point = {x:start.x, y:start.y}
        data.start = true
        afterObject()

        points[points.length - 1].push(point)
        data.points = points
        data.saved = false
        
        afterStandard()
        repaint()
        return
      }
    }

    if (data.shift == true) {
      let item = points[points.length - 1]
      let pos = item.length
      if (Math.abs(item[pos - 2].y - item[pos - 1].y) > Math.abs(item[pos - 2].x - item[pos - 1].x)) {
        points[points.length - 1][pos - 1].x = item[pos - 2].x
      } else {
        points[points.length - 1][pos - 1].y = item[pos - 2].y
      }

      if (data.type != 1) {
        points[points.length - 1].push(point)
      }
    } else {
      points[points.length - 1].push(point)
    }

    if (data.type == 1) {
      data.start = true
      afterObject()
    }
  }

  data.points = points
  data.saved = false
  
  afterStandard()
  repaint()
}

function afterStandard() {
  if (data.start == false) {
    return
  }

  let points = util.clone(data.points)

  let count = 0
  for (let i = 0; i < points.length; i++) {
    if (points[i][0].type == 1) {
      count++
    }
  }

  if (count > 1) {
    let items = []

    for (let i = 0; i < points.length - 1; i++) {
      if (points[i][0].type != 1) {
        items.push(points[i])
      }
    }

    data.points = [points[points.length - 1], ...items]
  }
}

function clickMouseRight() {
  if (data.start == false && (data.type == 2 || data.type == 4)) {
    let points = util.clone(data.points)
    
    if (points.length > 0) {
      if (points[points.length - 1].length >= 3) {
        points[points.length - 1][points[points.length - 1].length - 1] = util.clone(points[points.length - 1][1])
      }
    }

    data.points = points        
  }

  if (data.start == false && data.type == 3) {    
    // if (data.shift == true) {
    //   let points = util.clone(data.points)
    //   let item = points[points.length - 1]
    // 
    //   let j = item.length - 1
    //   
    //   let point = item[j]
    //   let x = point.x * data.zoom
    //   let y = point.y * data.zoom
    //   
    //   
    //   if (Math.abs(item[j - 1].y - y) > Math.abs(item[j - 1].x - x)) {
    //     x = item[j - 1].x * data.zoom
    //     points[points.length - 1][points[points.length - 1].length - 1].x = x
    //   } else {
    //     y = item[j - 1].y * data.zoom
    //     points[points.length - 1][points[points.length - 1].length - 1].y = y
    //   }      
    // 
    //   data.points = points
    // }    
  }

  
  data.saved = false
  data.start = true
  afterObject()

  afterStandard()
  repaint()
}

function moveMouse(e) {
  if (data.start == true) {
    return
  }

  if (data.points.length == 0) {
    return
  }
  
  let margin = document.getElementById("canvas").getBoundingClientRect()
  let x = (e.clientX - margin.x + 1) / data.zoom
  let y = (e.clientY - margin.y + 1) / data.zoom
  let point = {x:x, y:y}

  let points = util.clone(data.points)

  
  points[points.length - 1][points[points.length - 1].length - 1] = point

  data.points = points
  repaint()
}

function changeType() {
  if (data.start != true) {
    data.start = true
    let points = util.clone(data.points)
    points.pop()
    data.points = points
    repaint()
  }
}

function clickCut() {
  const now = new Date()
  const diff = now - _lastDate

  // if (data.start == true && diff / 1000 > 10) {
  //   return
  // }

  let points = util.clone(data.points)

  if (points.length == 0) {
    return
  }

  if (points[points.length - 1].length <= 3) {
    return
  }

  points[points.length - 1].splice(points[points.length - 1].length - 2, 1)
  
  data.points = points

  data.start = false

  repaint()
  afterStandard()

  data.saved = false
}

function clickZoomUp() {
  if (data.zoom >= 10.0) {
    return
  }

  if (data.zoom < 1.0) {
    data.zoom *= 2
  } else {
    data.zoom += 0.5
  }

  canvas.value.width = data.imgWidth * data.zoom
  canvas.value.height = data.imgHeight * data.zoom

  repaint()
}

function clickZoomDown() {
  if (data.zoom < 0.05) {
    return
  }
  
  if (data.zoom <= 1.0) {
    data.zoom /= 2
  } else {
    data.zoom -= 0.5
  }

  canvas.value.width = data.imgWidth * data.zoom
  canvas.value.height = data.imgHeight * data.zoom

  repaint()
}

function keyDown(e) {
  if (e.key == 'Shift') {
    if (data.usershift == true) {
      data.shift = false
    } else {
      data.shift = true
    }
    repaint()
  }
}

function keyUp(e) {
  let frame = document.getElementById('canvasFrame')
  let top = frame.scrollTop
  let left = frame.scrollLeft

  let width = frame.getBoundingClientRect().width
  let height = frame.getBoundingClientRect().height

  let inc = 100 * data.zoom

  if (data.usershift == true) {
    data.shift = true
  } else {
    data.shift = false
  }
  
  if(e.key === 'Escape') {
    clickCut()
    /*
       if (data.type == 2 || data.type == 4) {
       let points = util.clone(data.points)
       let start = points[points.length - 1][1]

       let point = {x:start.x, y:start.y}
       points[points.length - 1].push(point)

       data.points = points
       }

       data.start = true
       afterObject()

       afterStandard()
       repaint()
     */    
  } else if (e.key == 'z') {
    clickZoomUp()
  } else if (e.key == 'x') {
    clickZoomDown()        
  } else if (e.key == 'a') {
    left -= inc
    if (left <= 0) {
      left = 0
    }

    frame.scroll({
      top: top,
      left: left,
      behavior: 'smooth',
    })
  } else if (e.key == 'd') {
    left += inc
    let max = (data.imgWidth * data.zoom) - width
    if (left >= max) {
      left = max
    }

    frame.scroll({
      top: top,
      left: left,
      behavior: 'smooth',
    })
  } else if (e.key == 'w') {
    top -= inc
    if (top <= 0) {
      top = 0
    }

    frame.scroll({
      top: top,
      left: left,
      behavior: 'smooth',
    })
  } else if (e.key == 's') {
    top += inc
    let max = (data.imgHeight * data.zoom) - height
    if (top >= max) {
      top = max
    }

    frame.scroll({
      top: top,
      left: left,
      behavior: 'smooth',
    })
  } else if (e.key == '1') {
    data.type = 1
    data.start = true
    repaint()
    afterObject()
  } else if (e.key == '2') {
    data.type = 2
    data.start = true
    repaint()
    afterObject()
  } else if (e.key == '3') {
    data.type = 3
    data.start = true
    repaint()
    afterObject()
  } else if (e.key == '4') {
    data.type = 4
    data.start = true
    repaint()
    afterObject()
  } else {
    repaint()
  }
}

function clickUndo() {
  if (data.undos.length == 0) {
    return
  }

  let undos = util.clone(data.undos)
  let points = undos.pop()

  data.redos.push(points)

  if (undos.length > 0) {
    data.points = undos[undos.length - 1]
  } else {
    data.points = []
  }
  data.undos = undos
  repaint()

  data.start = true
  data.saved = false
}

function clickRedo() {
  if (data.redos.length == 0) {
    return
  }

  let redos = util.clone(data.redos)
  let points = redos.pop()

  data.points = points
  data.undos.push(util.clone(points))
  data.redos = redos
  repaint()

  data.start = true
  data.saved = false
}

function afterObject() {
  let points = util.clone(data.points)
  data.undos.push(points)
  data.redos = []
}

function getLength(point1, point2) {
  return Math.sqrt(Math.pow(point2.x - point1.x, 2) + Math.pow(point2.y - point1.y, 2))
}

function getValue(item) {
  if (data.points.length == 0) {
    return 0
  }

  let point = data.points[0]

  let standard = getLength(point[1], point[2])
  let length = getLength(item[1], item[2])

  if (standard == 0) {
    return 0
  }

  let type = item[0].type

  if (type == 1) {
    return util.getFloat(data.item.standard * length / standard).toFixed(4)
  } else if (type == 3 || type == 4) {
    let total = 0.0
    let lengths = []
    for (let i = 1; i < item.length - 1; i++) {
      let length = getLength(item[i], item[i+1])
      if (length == 0) {
        continue
      }

      let value = util.getFloat(data.item.standard * length / standard).toFixed(4)
      total += util.getFloat(value)
    }

    return total.toFixed(4)
  }
}

function getPosition() {
  if (data.position == 2) {
    return width(253)
  } else {
    return '16px'
  }
}

function clickPosition() {
  if (data.position == 2) {
    data.position = 1
  } else {
    data.position = 2
  }
}

function clickDeletePoint(index) {
  let points = util.clone(data.points)
  data.undos.push(points)

  points = util.clone(data.points)
  points.splice(index, 1)

  data.points = points
  repaint()

  data.saved = false  
}

function getTotal() {
  let standard = 0
  if (data.points.length > 0) {
    let point = data.points[0]
    standard = getLength(point[1], point[2])
  }

  let total = 0.0

  for (let j = 1; j < data.points.length; j++) {
    let item = data.points[j]

    let type = item[0].type

    if (type == 1) {
      let length = getLength(item[1], item[2])
      total += parseFloat(util.getFloat(data.item.standard * length / standard).toFixed(4))
    } else if (type == 3 || type == 4) {
      let lengths = []
      for (let i = 1; i < item.length - 1; i++) {
        let length = getLength(item[i], item[i+1])
        if (length == 0) {
          continue
        }

        let value = util.getFloat(data.item.standard * length / standard).toFixed(4)
        total += parseFloat(value)
      }
    }
  }

  return total.toFixed(4)
}

function getArea(item) {
  let point = data.points[0]
  let standard = getLength(point[1], point[2])

  var sum = 0.0;

  for (var i = 1; i < item.length - 1; i++) {
    sum += ((item[i].x * item[i + 1].y) / 2.0 - (item[i + 1].x * item[i].y) / 2.0);
  }

  return util.getFloat(data.item.standard * data.item.standard * Math.abs(sum) / standard / standard).toFixed(4)
}

function getTotalArea() {
  let total = 0.0
  for (let i = 0; i < data.points.length; i++) {
    let item = data.points[i]

    if (item[0].type != 2 && item[0].type != 4) {
      continue
    }

    total += util.getFloat(getArea(item))
  }

  return total.toFixed(4)
}

const lengths = computed(() => {
  let points = util.clone(data.points)
  return points.map(function(item, index) {
    item[0].index = index

    return item
  }).filter(item => item[0].type != 2)
})

const areas = computed(() => {
  let points = util.clone(data.points)
  return points.map(function(item, index) {
    item[0].index = index

    return item
  }).filter(item => item[0].type == 2 || item[0].type == 4)
})

function clickSubmitDraw() {
  util.confirm('저장하시겠습니까', async function() {
    util.loading(true)

    let item = util.clone(data.item)
    item.standard = util.getFloat(item.standard)
    item.content = JSON.stringify(data.points)
    await model.update(item)

    util.info('저장되었습니다')
    //await getItems()

    util.loading(false)
    data.saved = true
  })
}

function clickApt() {
  data.aptpage = 1
  data.apts = []

  data.selectapt = {
    id: 0,
    name: ''
}

    data.visibleApt = true
}

async function clickSearchApt() {
  await getApts(true)
}

async function getApts(reset) {
  if (reset == true) {
    data.aptpage = 1
    data.apts = []
  }

  let res = await Apt.search({page: data.aptpage, pagesize: data.aptpagesize, search: data.aptname})

  if (res.items == null) {
    res.items = []
  }

  data.apts = data.apts.concat(res.items)

  data.aptpage++  
}

function selectApt(item) {
  data.selectapt = item
  data.visibleApt = false
}

function inputKey() {
  data.saved = false
}

async function clickSubmitMulti() {
  if (data.apt > 0) {
    data.selectapt.id = data.apt
  }

  if (data.selectapt.id == 0) {
    util.error('아파트를 선택하세요')
    return
  }

  let count = 0
  for (let i = 0; i < uploads.value.length; i++) {
    if (data.names[i] != '' && data.filenames[i] != '') {
      count++
    }
  }

  if (count == 0) {
    util.error('데이터를 입력하세요')
    return
  }

  util.loading(true)

  for (let i = 0; i < uploads.value.length; i++) {
    if (data.names[i] == '') {
      continue
    }

    if (data.filenames[i] == '') {
      continue
    }

    let item = util.clone(data.item)
    item.apt = data.selectapt.id
    item.name = data.names[i]
    item.filename = data.filenames[i]
    await model.insert(item)
  }

  await getItems(true)
  util.info('등록되었습니다')
  util.loading(false)

  data.visibleMulti = false
}

function clickCancelMulti() {
  data.visibleMulti = false
}

function getAptName() {
  let name = ''
  if (!util.isNull(data.item.extra)) {
    if (!util.isNull(data.item.extra.apt)) {
      if (!util.isNull(data.item.extra.apt)) {
        name = data.item.extra.apt.name + ' - '
      }
    }
  }

  name += data.item.name

  return name
}
</script>
<style>
.el-message {
  z-index: 99999 !important;
}

</style>

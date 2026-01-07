<template>
  <Title title="장기수선 관리" />
  
  <div style="display:flex;gap: 10px;margin-bottom:10px;">
    <el-select v-model.number="data.search.status" placeholder="상태" style="width:80px;">           
      <el-option
        v-for="item in data.statuss"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>      
    
    <el-input v-model="data.search.text" placeholder="검색할 내용을 입력해 주세요" style="width:300px;" @keypress.enter.native="clickSearch" />

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>


    <TotalDiv :total="data.total" />          
  </div>  
  
  <el-table :data="data.items" border style="width: 100%;" :height="height(170)" v-infinite="getItems" @row-click="clickUpdate" >
    <el-table-column prop="id" label="ID" width="50" align="center" />
    <el-table-column prop="name" label="아파트명" />      
    <el-table-column prop="name" label="구분" align="center" width="80">
      <template #default="scope">
        <el-tag :type="Repair.getTypeType(scope.row.repairtype)">{{Repair.getType(scope.row.repairtype)}}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="reportdate" label="리포트작성일" align="center" width="80" />
    <el-table-column label="설명">
      <template #default="scope">
        <span v-if="scope.row.repairtype == 3">{{scope.row.info1}}</span>
      </template>
    </el-table-column>
    <el-table-column prop="address" label="도로명주소" />
    <el-table-column prop="repairdate" label="등록일" align="center" width="130" />
    <el-table-column label="상태" align="center" width="60">
          <template #default="scope">
            <span v-if="scope.row.status == 2 || scope.row.repairtype == 3">마감</span>
            <span v-if="scope.row.repairtype != 3 && scope.row.status == 1">진행</span>                        
          </template>
    </el-table-column>
    <el-table-column label="" width="80" align="center" >
      <template #default="scope">        
        <el-button v-if="scope.row.repairtype != 3 && scope.row.status == 1" size="small" type="success" @click="clickRepair(scope.row)">장기수선</el-button>
      </template>
    </el-table-column>
  </el-table>  
  <div style="margin-top:10px;display:flex;justify-content: space-between;">
    <!--<el-button size="small" type="success" @click="clickInsert">등록</el-button>-->    
    <div style="flex:0;"></div>
  </div>

  <el-dialog
    v-model="data.visible"
    title="주소 등록/수정"
    width="1100px"
    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="100px">

      <div>
        <div style="float:left;text-align:right;">
          <el-form-item label="ID" v-show="data.item.id != 0">
            {{ data.item.id }}
          </el-form-item>
        </div>
        <div style="float:right;text-align:right;">
          <el-button size="small" @click="data.visible = false">취소</el-button>
          <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
        </div>
        <div style="clear:both;"></div>
      </div>
      
      
      <el-form-item label="아파트명" style="margin:5px 0px;">
        <el-input v-model="data.item.name" />
      </el-form-item>

          

      
      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 33%;">


          <el-form-item label="준공년도" style="margin:5px 0px;">
            <el-input v-model="data.item.completeyear" />
          </el-form-item>

          <el-form-item label="건축물형태" style="margin:5px 0px;">
            <el-input v-model="data.item.type" />
          </el-form-item>
          

          <el-form-item label="전체동수" style="margin:5px 0px;">
            <el-input v-model="data.item.flatcount" />
          </el-form-item>

          <el-form-item label="세대수" style="margin:5px 0px;">
            <el-input v-model="data.item.familycount" />
          </el-form-item>

          <el-form-item label="층수" style="margin:5px 0px;">
            <el-input v-model="data.item.floor" />
          </el-form-item>


          

          <el-form-item label="전화번호" style="margin:5px 0px;">
            <el-input v-model="data.item.tel" />
          </el-form-item>

          <el-form-item label="팩스번호" style="margin:5px 0px;">
            <el-input v-model="data.item.fax" />
          </el-form-item>
          

      
          <el-form-item label="공용메일주소" style="margin:5px 0px;">
            <el-input v-model="data.item.email" />
          </el-form-item>
          
          
    
          <el-form-item label="개인메일주소" style="margin:5px 0px;">
            <el-input v-model="data.item.personalemail" :rows="2" type="textarea" />
          </el-form-item>
          
          
         
        </div>
        <div style="flex:1 1 67%;">

          
          
          <el-form-item label="장기수선계획" style="margin:5px 0px;">
            <el-input v-model="data.item.repair" :rows="5" type="textarea" />
          </el-form-item>
          <el-form-item label="안전점검" style="margin:5px 0px;">
            <el-input v-model="data.item.safety" :rows="5" type="textarea" />
          </el-form-item>
          <el-form-item label="기타" style="margin:5px 0px;">
            <el-input v-model="data.item.fault" :rows="5" type="textarea" />
          </el-form-item>      
          
          
          
          
        </div>
      </div>

      
      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 20%;">

          <el-form-item label="우편번호" style="margin:5px 0px;">
            <el-input v-model="data.item.zip" />
          </el-form-item>
        </div>

        
        <div style="flex:1 1 80%;">


          <el-form-item label="도로명주소" style="margin:5px 0px;">
            <el-input v-model="data.item.address" />
          </el-form-item>
          
        </div>
      </div>
      
      
      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 20%;">

        </div>

        
        <div style="flex:1 1 80%;">
          
          <el-form-item label="지번주소" style="margin:5px 0px;">
            <el-input v-model="data.item.address2" />
          </el-form-item>
          
          
          
        </div>
      </div>


      

      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 33%;">
          <el-form-item label="정밀점검일자" style="margin:5px 0px;">
            <el-input v-model="data.item.testdate" />
          </el-form-item>

          <!--
          <el-form-item label="정밀안전점검 차기일자">
            <el-input v-model="data.item.nexttestdate" />
          </el-form-item>
          -->
          
        </div>
        <div style="flex:1 1 33%;">
          <el-form-item label="FMS 아이디" style="margin:5px 0px;">
            <el-input v-model="data.item.fmsloginid" />
          </el-form-item>
          
        </div>
        <div style="flex:1 1 33%;">
          <el-form-item label="FMS 비번" style="margin:5px 0px;">
            <el-input v-model="data.item.fmspasswd" />
          </el-form-item>
          
        </div>
        
        


        
      </div>

      
      <el-form-item label="계약구분" style="margin:5px 0px;">        

        <el-checkbox label="장기수선계획" v-model="data.check1" />
        <el-checkbox label="정밀" v-model="data.check2" />
        <el-checkbox label="정기" v-model="data.check3" />
        <el-checkbox label="하자보수" v-model="data.check4" />
        <el-checkbox label="하자조사" v-model="data.check5" />
        <el-checkbox label="정밀안전진단" v-model="data.check6" />
        <el-checkbox label="감리" v-model="data.check7" />
        <el-checkbox label="기술자문" v-model="data.check8" />        
      </el-form-item>

      <div style="display:flex;flex-direction:row;">
        <div style="flex:1 1 25%;">
          
          <el-form-item label="계약날짜" style="margin:5px 0px;">
            <el-input v-model="data.item.contractdate" />
          </el-form-item>

        </div>
        <div style="flex:1 1 25%;">
          
          <el-form-item label="계약기간" style="margin:5px 0px;">
            <el-input v-model="data.item.contractduration" />
          </el-form-item>

        </div>
        <div style="flex:1 1 25%;">          
          
          <el-form-item label="계산서 발행" style="margin:5px 0px;">
            <el-input v-model="data.item.invoice" />            
          </el-form-item>

        </div>
        <div style="flex:1 1 25%;">

          <el-form-item label="입금날짜" style="margin:5px 0px;">
            <el-input v-model="data.item.depositdate" />
          </el-form-item>

        </div>
      </div>
          
        
          
    </el-form>

    <template #footer>
      <el-button size="small" style="float:left;" type="danger" @click="clickDelete(data.item)">삭제</el-button>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>


  <el-dialog
    v-model="data.visibleRepair"
    width="800px"
    :before-close="handleClose"
  >
    <RepairInsert :id="data.id" />
  </el-dialog>
  
</template>

<script setup lang="ts">

import { reactive, onMounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Repair, Repairlist, Apt } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

const { width, height } = size()

const store = useStore()
const route = useRoute()

function clickSearch() {
  getItems(true)
}

const item = {
  id: 0,
  apt: 0,
  name: '',
  contracttype: 0,
  completeyear: '',
  flatcount: '',
  type: '',
  floor: '',
  familycount: '',
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
  items: [],
  total: 0,
  page: 1,
  pagesize: 50,
  item: util.clone(item),
  visible: false,
  visibleRepair: false,
  search: {
    status: 1,
    text: ''
  }
})

async function initData() {
  data.statuss = [{id: 0, name: ' '}, {id: 1, name: '진행'}, {id: 2, name: '마감'}]
}

async function getItems(reset: boolean) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }
  
  let res = await Repairlist.search({page: data.page, pagesize: data.pagesize, status: data.search.status, search: data.search.text})  
  
  if (res.items == null) {
    res.items = []
  }
  
  data.total = res.total
  data.items = data.items.concat(res.items)

  data.page++
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
  
  await initData()
  await getItems()

  util.loading(false)
})

async function clickRepair(item) {
  const apt = item.id
  const id = item.repairid
  router.push(`/${apt}/repair/${id}/breakdown`)
}

function clickUpdate(item, index) {
  if (index.no == 8) {
    return
  }

  data.id = item.id
  data.visibleRepair = true
}
</script>

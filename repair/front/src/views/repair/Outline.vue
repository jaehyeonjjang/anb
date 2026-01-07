<template>
  <Title title="총론" />


  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="success" @click="clickBatch" style="margin-right:0px;">등록</el-button>
    </div>
  </div>  

  <div style="display:flex;">
    <div style="flex:1;">

      <div :style="{'height': data.height, 'overflow': 'auto', 'padding-right': '10px'}">
        
        <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
          장기수선충당금 계획금액
        </div>
      
      <el-table :data="data.totals" border :cell-class-name="cellClassNameTotal" :span-method="spanMethodTotal">
        <el-table-column prop="index" label="구분" align="center" width="50" />
        <el-table-column prop="title" label="내용" />
        <el-table-column prop="content1" label="" align="center" />
        <el-table-column prop="content2" label="계산식" align="center" width="50" />
        <el-table-column prop="content3" label="" align="center" />
        <el-table-column prop="price" label="금액" align="center" />
      </el-table>


      <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
        장기수선충당금 계획금액 (* 50년 가상 금액)
      </div>
      
      <el-table :data="data.totalplans" border :cell-class-name="cellClassNameTotal" :span-method="spanMethodTotalPlan">
        <el-table-column prop="index" label="구분" align="center" width="50" />
        <el-table-column prop="title" label="내용" />
        <el-table-column prop="content1" label="" align="center" />
        <el-table-column prop="content2" label="계산식" align="center" width="50" />
        <el-table-column prop="content3" label="" align="center" />
        <el-table-column prop="price" label="금액" align="center" />
      </el-table>

      <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
        장기수선충당금 계획금액과 장기수선충당금 적립금액의 비교
      </div>
      
      <el-table :data="data.totalcompares" border >
        <el-table-column prop="index" label="구분" align="center" width="50" />
        <el-table-column prop="title" label="" />
        <el-table-column prop="price" label="장기수선계획금액" align="right" />
        <el-table-column prop="rate" label="비율" align="center" width="60">
          <template #default="scope">
            {{util.fixed(100.0 * util.getFloat(scope.row.price2) / util.getFloat(scope.row.price), 2)}}
          </template>
        </el-table-column>
        <el-table-column prop="price2" label="적립예정금액" align="right" />
        <el-table-column prop="title2" label="" />
      </el-table>

      
      </div>
      
    </div>
    <div style="flex:1;">

      <div :style="{'height': data.height, 'overflow': 'auto'}">
        
        <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
        수립 예정단가로 본 적립 요율 적용(관리규약 요율 및 단가)
      </div>

      
      <el-table :data="data.items" border>          
        <el-table-column label="시설물의 내구 연한" align="center" width="180">
          <template #default="scope">
            <div style="display:flex;font-size:12px;">
              <div style="text-align:left;margin-left:7px;width:50px;">{{scope.row.startyear}}.{{util.pad(scope.row.startmonth, 2)}}</div> <div>~</div> <div style="text-align:left;margin-left:5px;width:50px;"> {{scope.row.endyear}}.{{util.pad(scope.row.endmonth, 2)}}</div>
              <div style="margin-left:5px;width:40px;">({{scope.row.duration}})</div>
            </div>              
          </template>
        </el-table-column>

        <el-table-column label="구간별 적립금액" align="right">
          <template #default="scope">
            <span style="font-size:12px;">{{scope.row.totalprice}}</span>
          </template>
        </el-table-column>

        <el-table-column label="적용 적립요율" align="center">
          <template #default="scope">
            {{scope.row.rate}}
          </template>
        </el-table-column>

        <el-table-column label="m2당 단가" align="center">
          <template #default="scope">
            {{scope.row.price}}              
          </template>
        </el-table-column>
        
        <el-table-column label="누적요율" align="center">
          <template #default="scope">
            {{scope.row.totalrate}}
          </template>
        </el-table-column>

        <el-table-column label="비고">
          <template #default="scope">
            {{scope.row.remark}}
          </template>
        </el-table-column>
        
      </el-table>



      <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
        수립 예정단가로 본 적립 요율 적용(향후 개정시 적용해야할 요율 및 단가)
      </div>

      <el-table :data="data.itemplans" border>          
        <el-table-column label="시설물의 내구 연한" align="center" width="180">
          <template #default="scope">
            <div style="display:flex;font-size:12px;">
              <div style="text-align:left;margin-left:7px;width:50px;">{{scope.row.startyear}}.{{util.pad(scope.row.startmonth, 2)}}</div> <div>~</div> <div style="text-align:left;margin-left:5px;width:50px;"> {{scope.row.endyear}}.{{util.pad(scope.row.endmonth, 2)}}</div>
              <div style="margin-left:5px;width:40px;">({{scope.row.duration}})</div>              
            </div>              
          </template>
        </el-table-column>

        <el-table-column label="구간별 적립금액" align="right">
          <template #default="scope">
            <span style="font-size:12px;">{{scope.row.totalprice}}</span>
          </template>
        </el-table-column>

        <el-table-column label="적용 적립요율" align="center">
          <template #default="scope">
            {{scope.row.rate}}
          </template>
        </el-table-column>

        <el-table-column label="m2당 단가" align="center">
          <template #default="scope">
            {{scope.row.price}}              
          </template>
        </el-table-column>
        
        <el-table-column label="누적요율" align="center">
          <template #default="scope">
            {{scope.row.totalrate}}
          </template>
        </el-table-column>

        <el-table-column label="비고">
          <template #default="scope">
            {{scope.row.remark}}
          </template>
        </el-table-column>
        
      </el-table>

      </div>
    </div>
  </div>
  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1100px"
  >

     <el-table :data="[1,2]" border :span-method="spanMethod" :cell-class-name="cellClassName" style="width:120px;">
       <el-table-column label="m2 평균 적립금액" align="center" width="120">
         <template #default="scope">
           <el-input v-if="scope.$index == 0" v-model="data.batchrepair.savingprice" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                     :parser="(value) => value.replace(/\$\s?|(,*)/g, '')"
                     class="inputNumber"
           />
         </template>
       </el-table-column>
       <!--
       <el-table-column label="기타 금액">
            <template #default="scope">
              <div v-if="scope.$index == 0">승강기 부품교체 공사시 종합 유지보수계약 후 장기수선계획에 따른 부품교체 공사비용은 장기수선충당 금<el-input v-model="data.batchrepair.price1" style="width:100px;" />으로 집행하고,유지보수비용은 관리비<el-input v-model="data.batchrepair.price2" style="width:100px;" />로 집행</div>

              <div v-if="scope.$index == 1">긴급의 단일공사 <el-input v-model="data.batchrepair.price3" style="width:100px;" />만원 이내이고 년간공사 <el-input v-model="data.batchrepair.price4" style="width:100px;" />만원 이내로 하고 소액은 <el-input v-model="data.batchrepair.price5" style="width:100px;" />만원 수의계약가능하나 그 이상은 입찰로 진행한다.</div>
            </template>
          </el-table-column>              
          -->
    </el-table>


    <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
      수립 예정단가로 본 적립 요율 적용(관리규약 요율 및 단가)
    </div>

    
    <el-table :data="data.batchs" border :summary-method="getSummaries" show-summary>
          <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon v-if="scope.$index > 0" @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
          </el-table-column>
          <el-table-column label="시설물의 내구 연한" align="center" width="240">
            <template #default="scope">
                <div style="display:flex;font-size:12px;">
                  <div style="width:80px;">
                    <el-input class="date" v-if="scope.$index == 0" v-model="data.batchs[scope.$index].startyear" style="margin-left:0px;width:40px;" @keyup="onKeyup(scope.$index, 1)" />
                    <el-input class="date" v-if="scope.$index == 0" v-model="data.batchs[scope.$index].startmonth" style="margin-left:3px;width:30px;" @keyup="onKeyup(scope.$index, 1)" />
                    
                    <span  v-if="scope.$index != 0">{{data.batchs[scope.$index].startyear}}.{{util.pad(data.batchs[scope.$index].startmonth, 2)}}</span>
                  </div>

                  <div style="text-align:center;width:20px;">~</div>
                  <div>
                    <el-input class="date" v-model="data.batchs[scope.$index].endyear" style="width:40px;" @keyup="onKeyup(scope.$index, 1)" />
                    <el-input class="date" v-model="data.batchs[scope.$index].endmonth" style="margin-left:3px;width:30px;" @keyup="onKeyup(scope.$index, 1)" />
                  </div>

                  <div style="width: 50px;">{{data.batchs[scope.$index].duration}}</div>
                </div>              
            </template>
          </el-table-column>

          <el-table-column label="구간별 적립금액" align="right" width="120">
            <template #default="scope">
              <span style="font-size:12px;">{{data.batchs[scope.$index].totalprice}}</span>
            </template>
          </el-table-column>

          <el-table-column label="적용 적립요율" align="center" width="170">
            <template #default="scope">
              <el-input v-if="scope.$index != data.batchs.length - 1" v-model="data.batchs[scope.$index].rate"  @input="onKeyup(scope.$index, 2)" />

              <div v-if="scope.$index == data.batchs.length - 1">
                <el-input v-model="data.batchs[scope.$index].rate"  @input="onKeyup(scope.$index, 2)" style="width:110px;display:block;float:left;" />
                <div style="display:block;float:left;width:45px;margin-left:5px;margin-top:3px;">
                  <el-button size="small" type="warning" @click="clickHundred" style="width:45px;height:24px;">100%</el-button>
                </div>
              </div>
              
            </template>
          </el-table-column>

          <el-table-column label="m2당 단가" align="center" width="120">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].price" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                        :parser="(value) => value.replace(/\$\s?|(,*)/g, '')"
                        class="inputNumber"
                        @input="onKeyup(scope.$index, 3)"
              />
            </template>
          </el-table-column>
          
          <el-table-column label="누적요율" align="center" width="140">
            <template #default="scope">
              {{data.batchs[scope.$index].totalrate}}
            </template>
          </el-table-column>

          <el-table-column label="비고">
            <template #default="scope">
              <el-input v-model="data.batchs[scope.$index].remark" />
            </template>
          </el-table-column>
          
        </el-table>

        <div style="margin-top:5px;text-align:left;">
          <el-button size="small" v-if="data.mode == 'batch'" @click="clickAdd"><el-icon><Plus /></el-icon></el-button>
        </div>





        <div style="font-size:12px;font-weight:bold;text-align:left;margin-top:15px;margin-bottom:5px;">
          수립 예정단가로 본 적립 요율 적용(향후 개정시 적용해야할 요율 및 단가)
        </div>

        

        <el-table :data="data.batchplans" border  :summary-method="getSummariesPlan" show-summary>
          <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon v-if="scope.$index > 0" @click="clickRegistPlanDelete(scope.$index)"><Delete /></el-icon>
          </template>
          </el-table-column>
          <el-table-column label="시설물의 내구 연한" align="center" width="240">
            <template #default="scope">
              <div style="display:flex;font-size:12px;">
                <div style="width:80px;">
                  <el-input class="date" v-if="scope.$index == 0" v-model="data.batchplans[scope.$index].startyear" style="margin-left:0px;width:40px;" @keyup="onKeyupPlan(scope.$index, 1)" />
                  <el-input class="date" v-if="scope.$index == 0" v-model="data.batchplans[scope.$index].startmonth" style="margin-left:3px;width:30px;" @keyup="onKeyupPlan(scope.$index, 1)" />
                  <span  v-if="scope.$index != 0">{{data.batchplans[scope.$index].startyear}}.{{util.pad(data.batchplans[scope.$index].startmonth, 2)}}</span></div> <div style="text-align:center;width:20px;">~</div>
                  <div>
                    <el-input class="date" v-model="data.batchplans[scope.$index].endyear" style="width:40px;" @keyup="onKeyupPlan(scope.$index, 1)" />
                    <el-input class="date" v-model="data.batchplans[scope.$index].endmonth" style="margin-left:3px;width:30px;" @keyup="onKeyupPlan(scope.$index, 1)" />
                  </div>

                <div style="width: 50px;">{{data.batchplans[scope.$index].duration}}</div>
              </div>

                
            </template>
          </el-table-column>

          <el-table-column label="구간별 적립금액" align="right" width="120">
            <template #default="scope">
              <span style="font-size:12px;">{{data.batchplans[scope.$index].totalprice}}</span>
            </template>
          </el-table-column>

          <el-table-column label="적용 적립요율" align="center" width="170">
            <template #default="scope">
              <el-input  v-if="scope.$index != data.batchplans.length - 1" v-model="data.batchplans[scope.$index].rate"  @input="onKeyupPlan(scope.$index, 2)" />

              <div v-if="scope.$index == data.batchplans.length - 1">
                <el-input v-model="data.batchplans[scope.$index].rate"  @input="onKeyupPlan(scope.$index, 2)" style="width:110px;display:block;float:left;" />
                <div style="display:block;float:left;width:45px;margin-left:5px;margin-top:3px;">
                  <el-button size="small" type="warning" @click="clickHundredPlan" style="width:45px;height:24px;">100%</el-button>
                </div>
              </div>
            </template>
          </el-table-column>

          <el-table-column label="m2당 단가" align="center" width="120">
            <template #default="scope">
              <el-input v-model="data.batchplans[scope.$index].price" :formatter="(value) => `${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                        :parser="(value) => value.replace(/\$\s?|(,*)/g, '')"
                        class="inputNumber"
                        @input="onKeyupPlan(scope.$index, 3)"
              />
            </template>
          </el-table-column>
          
          <el-table-column label="누적요율" align="center" width="140">
            <template #default="scope">
              {{data.batchplans[scope.$index].totalrate}}
            </template>
          </el-table-column>

          <el-table-column label="비고">
            <template #default="scope">
              <el-input v-model="data.batchplans[scope.$index].remark" />
            </template>
          </el-table-column>
          
        </el-table>

        <div style="margin-top:5px;text-align:left;">
          <el-button size="small" v-if="data.mode == 'batch'" @click="clickAddPlan"><el-icon><Plus /></el-icon></el-button>
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
import { Report, Repair, Outline, Outlineplan } from "~/models"
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
  totalprice: 0,
  rate: 0,
  price: 0,
  totalrate: 0,
  remark: ''
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  itemplans: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  itemplan: util.clone(item),
  visible: false,    
  search: {
    text: ''
  },
  basics: [],
  batchs: [],
  batchplans: [],
  repair: {
    completionyear: '',
    completionmonth: ''
  },
  batchrepair: {
    savingprice: '',
    price1: '',
    price2: '',
    price3: '',
    price4: '',
    price2: ''
  },
  report: {
    price: 0,
    totalprice: 0,
    totalsize: 0
  },
  totals: [],
  totalplans: [],
  totalcompares: []
})

async function initData() {  
}

async function getItems() {
  let res = await Repair.get(data.apt)
  data.repair = res.item  

  res = await Report.total(data.apt)

  data.report = res
  
  res = await Outline.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'o_id'
  })
  
  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }
  }
  
  data.total = res.total
  if (res.items == null) {
    res.items = []
  }

  let alltotalprice = 0
  let sum = 0
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    sum += util.getFloatFixed(item.rate)

    let per = parseFloat(util.getFloatFixed(item.rate))
    let totalprice = 0
    if (per > 0) {      
      totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    }

    res.items[i].totalprice = util.money(totalprice)
    res.items[i].totalrate = sum

    let duration = getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth)    
    res.items[i].duration = duration
    
    if (i < res.items.length - 1) {
      alltotalprice += totalprice
    }
  }

  if (res.items.length > 0 && res.items[res.items.length - 1].totalrate === 100.0) {    
    res.items[res.items.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - alltotalprice)
  }
  
  data.items = res.items

  res = await Outlineplan.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'op_id'
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }
  }
  
  if (res.items == null) {
    res.items = []
  }

  alltotalprice = 0
  sum = 0
  
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    sum += util.getFloatFixed(item.rate)

    let per = parseFloat(util.getFloatFixed(item.rate))
    let totalprice = 0

    if (per > 0) {
      totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    }

    res.items[i].totalprice = util.money(totalprice)
    res.items[i].totalrate = sum

    let duration = getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth)    
    res.items[i].duration = duration

    if (i < res.items.length - 1) {
      alltotalprice += totalprice
    }
  }

  if (res.items.length > 0 && res.items[res.items.length - 1].totalrate === 100.0) {    
    res.items[res.items.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - alltotalprice)
  }


  data.itemplans = res.items

  let plan = await Report.plan(data.apt)

  let years = parseFloat(util.getPlanyears(data.repair.planyears) * 12)  
  
  let items = [
    {
      index: 1,
      title: '총 계획금액',
      content1: plan.startyear,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(plan.total),
      remark: '사용검사년도는 제외'
    },
    {
      index: '',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.total) / years * 12)),
      content2: '*',
      content3: util.getPlanyears(data.repair.planyears),
      price: 0,
      remark: ''
    },
    {
      index: '2',
      title: '기준년도까지 적립금액',
      content1: plan.startyear,
      content2: '~',
      content3: plan.reportyear,
      price: util.money(Math.round(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1) * util.getFloatFixed(plan.total) / years * 12),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.total) / years * 12)),
      content2: '*',
      content3: util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1,
      price: 0,
      remark: ''
    },
    {
      index: '3',
      title: '기준년도 이후 적립금액',
      content1: util.getInt(plan.reportyear) + 1,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(Math.round(util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1)) * util.getFloatFixed(plan.total) / years * 12),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.total) / years * 12)),
      content2: '*',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1),
      price: 0,
      remark: ''
    },
    {
      index: '4',
      title: '년 평균 계획금액',
      content1: '월 평균 계획금액',
      content2: '*',
      content3: '12개월',
      price: util.money(Math.round(util.getFloatFixed(plan.total) / years * 12)),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.total) / years)),
      content2: '*',
      content3: 12,
      price: 0,
      remark: ''
    },
    {
      index: '5',
      title: '월 평균 계획금액',
      content1: '총 계획금액',
      content2: '/',
      content3: '적립개월수',
      price: util.money(Math.round(util.getFloatFixed(plan.total) / years)),
      remark: '사용검사년도는 제외'
    },
    {
      index: '2',
      title: '',
      content1: util.money(plan.total),
      content2: '/',
      content3: util.getInt(years),
      price: 0,
      remark: ''
    },
    {
      index: '6',
      title: 'm2 평균 계획단가',
      content1: '월 평균 계획금액',
      content2: '/',
      content3: '총 주택공급 면적',
      price: Math.round(util.getFloatFixed(plan.total) / years / plan.area),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.total) / years)),
      content2: '/',
      content3: util.fixed(util.getFloatFixed(plan.area), 3),
      price: 0,
      remark: ''
    },
  ]

  data.totals = items

  let itemplans = [
    {
      index: 1,
      title: '총 계획금액',
      content1: plan.startyear,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12 * util.getPlanyears(data.repair.planyears))),
      remark: '사용검사년도는 제외'
    },
    {
      index: '',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12)),
      content2: '*',
      content3: util.getPlanyears(data.repair.planyears),
      price: 0,
      remark: ''
    },
    {
      index: '2',
      title: '기준년도까지 적립금액',
      content1: plan.startyear,
      content2: '~',
      content3: plan.reportyear,
      price: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12 * util.getFloatFixed(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1))),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12)),
      content2: '*',
      content3: util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1,
      price: 0,
      remark: ''
    },
    {
      index: '3',
      title: '기준년도 이후 적립금액',
      content1: util.getInt(plan.reportyear) + 1,
      content2: '~',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - 1,
      price: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12 * util.getFloatFixed(util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1)))),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12)),
      content2: '*',
      content3: util.getInt(plan.startyear) + util.getPlanyears(data.repair.planyears) - (util.getInt(plan.reportyear) + 1),
      price: 0,
      remark: ''
    },
    {
      index: '4',
      title: '년 평균 적립금액',
      content1: '월 평균 계획금액',
      content2: '*',
      content3: '12개월',
      price: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12)),
      remark: ''
    },
    {
      index: '2',
      title: '',
      content1: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice))),
      content2: '*',
      content3: 12,
      price: 0,
      remark: ''
    },
    {
      index: '5',
      title: '월 평균 적립금액',
      content1: 'm2 평균 적립금액',
      content2: '/',
      content3: '총 부과면적',
      price: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice))),
      remark: '사용검사년도는 제외'
    },
    {
      index: '2',
      title: '',
      content1: util.money(data.repair.savingprice),
      content2: '/',
      content3: util.fixed(util.getFloatFixed(plan.area), 3),
      price: 0,
      remark: ''
    },
    {
      index: '6',
      title: 'm2 평균 적립금액',
      content1: '기준일 현재 부과금액',
      content2: '/',
      content3: '총 주택공급 면적',
      price: util.money(data.repair.savingprice),
      remark: ''
    }    
  ]

  data.totalplans = itemplans

  let compares = [
    {
      index: 1,
      title: '총 계획금액',
      title2: '총 적립금액',
      price: util.money(plan.total),
      price2: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12 * util.getPlanyears(data.repair.planyears)))
    },
    {
      index: 2,
      title: '기준일까지 계획금액',
      title2: '기준일까지 적립금액',
      price: util.money(Math.round(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1) * util.getFloatFixed(plan.total) / years * 12),
      price2: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12 * util.getFloatFixed(util.getInt(plan.reportyear) - util.getInt(plan.startyear) + 1)))
    },
    {
      index: 3,
      title: '년 평균 계획금액',
      title2: '년 평균 적립금액',
      price: util.money(Math.round(util.getFloatFixed(plan.total) / years * 12)),
      price2: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice) * 12))
    },
    {
      index: 4,
      title: '월 평균 계획금액',
      title2: '월 평균 적립금액',
      price: util.money(Math.round(util.getFloatFixed(plan.total) / years)),
      price2: util.money(Math.round(util.getFloatFixed(plan.area) * util.getFloatFixed(data.repair.savingprice)))
    },
    {
      index: 5,
      title: 'm2 평균 계획금액',
      title2: 'm2 평균 적립금액',
      price: Math.round(util.getFloatFixed(plan.total) / years / plan.area),
      price2: util.money(data.repair.savingprice)
    }
  ]
  
  data.totalcompares = compares
}

function makeItems(items) {
  let alltotalprice = 0
  let sum = 0
  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    let per = parseFloat(util.getFloatFixed(item.rate))
    let totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    let duration = util.getFloatFixed(getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth))      
    items[i].totalprice = util.money(totalprice)    
    
    sum += util.getFloatFixed(item.rate)

    //items[i].price = item.price
    console.log('sum:', sum)
    items[i].totalrate = sum

    if (i < items.length - 1) {
      alltotalprice += totalprice
    }

    items[i].duration = getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth)    
  }

  if (items.length > 0 && items[items.length - 1].totalrate === 100.0) {    
    items[items.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - alltotalprice)
  }
  
  return items
}

function clickInsert() {  
  data.item = util.clone(item)

  let items = makeItems([data.item])

  data.mode = 'normal'
  data.batchs = items
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

async function clickBatch() {
  let res = await Repair.get(data.apt)
  if (util.getInt(res.item.completionyear) == 0) {
    //util.error('개요 메뉴에서 사업계획승인 일자를 먼저 등록하세요')
    util.messagebox('', '개요 메뉴에서 사업계획승인 일자를 먼저 등록하세요')
    return
  }

  data.repair = res.item
  
  let items = util.clone(data.items)

  if (items == null) {
    items = []
  }

  
  if (items.length == 0) {
    let length = 5 - items.length
    for (let i = 0; i < length; i++) {
      items.push(util.clone(item))
    }

    let d = new Date()
    let year = d.getFullYear()
    items[0].startyear = data.repair.completionyear + 1
    items[0].endyear = year + 1

    for (let i = 1; i <= items.length - 1; i++) {
      items[i].startyear = year + 1 + (i-1)*10 + 1
      items[i].endyear = year + 1 + i*10
      items[i].duration = getDurationMonth(items[i].endyear, 12, items[i].startyear, 1)
    }
  }    

  items = makeItems(items)
  data.batchs = items

  items = util.clone(data.itemplans)

  if (items == null) {
    items = []
  }

  if (items.length == 0) {
    let length = 5 - items.length
    for (let i = 0; i < length; i++) {
      items.push(util.clone(data.item))
    }

    let d = new Date()
    let year = d.getFullYear()
    items[0].startyear = data.repair.completionyear + 1
    items[0].endyear = year + 1

    for (let i = 1; i <= items.length - 1; i++) {
      items[i].startyear = year + 1 + (i-1)*10 + 1
      items[i].endyear = year + 1 + i*10
      items[i].duration = getDurationMonth(items[i].endyear, 12, items[i].startyear, 1)
    }
  }

  items = makeItems(items)

  data.batchrepair = util.clone(data.repair)

  data.batchplans = items
  data.mode = 'batch'
  data.visible = true  
}

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

  if (data.mode == 'batch') {
    for (let i = 0; i < data.items.length; i++) {
      let item = data.items[i]
      let flag = false;
      for (let j = 0; j < data.batchs.length; j++) {
        if (data.items[i].id == data.batchs[j].id) {
          flag = true
          break
        }
      }

      if (flag == false) {      
        await Outline.remove(item)
      }
    }

    for (let i = 0; i < data.itemplans.length; i++) {
      let item = data.itemplans[i]
      let flag = false;
      for (let j = 0; j < data.batchplans.length; j++) {
        if (data.itemplans[i].id == data.batchplans[j].id) {
          flag = true
          break
        }
      }

      if (flag == false) {      
        await Outlineplan.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    item.startyear = util.getInt(item.startyear)
    item.endyear = util.getInt(item.endyear)
    item.startmonth = util.getInt(item.startmonth)
    item.endmonth = util.getInt(item.endmonth)
    item.rate = util.getFloat(item.rate)
    item.price = util.getFloat(item.price)

    if (item.id > 0) {
      await Outline.update(item)
    } else { 
      await Outline.insert(item)
    }
  }

  for (let i = 0; i < data.batchplans.length; i++) {
    let item = data.batchplans[i]

    item.apt = data.apt
    item.startyear = util.getInt(item.startyear)
    item.endyear = util.getInt(item.endyear)
    item.startmonth = util.getInt(item.startmonth)
    item.endmonth = util.getInt(item.endmonth)
    item.rate = util.getFloat(item.rate)
    item.price = util.getFloat(item.price)

    if (item.id > 0) {
      await Outlineplan.update(item)
    } else { 
      await Outlineplan.insert(item)
    }
  }

  let res = await Repair.get(data.apt)
  let repair = res.item

  repair.savingprice = util.getFloat(data.batchrepair.savingprice)
  repair.price1 = data.batchrepair.price1
  repair.price2 = data.batchrepair.price2
  repair.price3 = data.batchrepair.price3
  repair.price4 = data.batchrepair.price4
  repair.price5 = data.batchrepair.price5

  await Repair.update(repair)

  util.info('등록되었습니다')
  getItems()
  data.visible = false
  
  util.loading(false)  
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)

  for (let i = 0; i < data.batchs.length - 1; i++) {
    let endyear = util.getInt(data.batchs[i].endyear)
    let endmonth = util.getInt(data.batchs[i].endmonth)

    if (endmonth == 12) {
      data.batchs[i + 1].startyear = endyear + 1
      data.batchs[i + 1].startmonth = 1
    } else {
      data.batchs[i + 1].startyear = endyear
      data.batchs[i + 1].startmonth = endmonth + 1
    }
  }
}

function clickRegistPlanDelete(index) {
  data.batchplans.splice(index, 1)

  for (let i = 0; i < data.batchplans.length - 1; i++) {
    let endyear = util.getInt(data.batchplans[i].endyear)
    let endmonth = util.getInt(data.batchplans[i].endmonth)

    if (endmonth == 12) {
      data.batchplans[i + 1].startyear = endyear + 1
      data.batchplans[i + 1].startmonth = 1
    } else {
      data.batchplans[i + 1].startyear = endyear
      data.batchplans[i + 1].startmonth = endmonth + 1
    }
  }
}

function clickAdd() {
  let items = []

  let endyear = util.getInt(data.batchs[data.batchs.length - 1].endyear)
  let endmonth = util.getInt(data.batchs[data.batchs.length - 1].endmonth)

  if (endmonth == 12) {
    item.startyear = endyear + 1
    item.startmonth = 1
  } else {
    item.startyear = endyear
    item.startmonth = endmonth + 1
  }
  
  items.push(util.clone(item))

  data.batchs = data.batchs.concat(items)

  let sum = 0
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    sum += util.getFloatFixed(item.rate)

    data.batchs[i].totalrate = sum
  }
}

function clickAddPlan() {
  let items = []

  let endyear = util.getInt(data.batchplans[data.batchplans.length - 1].endyear)
  let endmonth = util.getInt(data.batchplans[data.batchplans.length - 1].endmonth)

  if (endmonth == 12) {
    item.startyear = endyear + 1
    item.startmonth = 1
  } else {
    item.startyear = endyear
    item.startmonth = endmonth + 1
  }
  items.push(util.clone(item))

  data.batchplans = data.batchplans.concat(items)

  let sum = 0
  for (let i = 0; i < data.batchplans.length; i++) {
    let item = data.batchplans[i]

    sum += util.getFloatFixed(item.rate)

    data.batchplans[i].totalrate = sum
  }
}

const spanMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  
  if (columnIndex === 0) {
    if (rowIndex == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }
  return {rowspan: 1, colspan: 1}
}

function cellClassName({row, columnIndex}) {
  return 'title'  
}

function onKeyup(index, col) {
  if (col == 1) {
    if (index >= data.batchs.length - 1) {
      return
    }    
    
    let endyear = util.getInt(data.batchs[index].endyear)
    let endmonth = util.getInt(data.batchs[index].endmonth)

    if (endmonth == 12) {
      data.batchs[index + 1].startyear = endyear + 1
      data.batchs[index + 1].startmonth = 1
    } else {
      data.batchs[index + 1].startyear = endyear
      data.batchs[index + 1].startmonth = endmonth + 1
    }

    data.batchs[index].duration = getDurationMonth(data.batchs[index].endyear, data.batchs[index].endmonth, data.batchs[index].startyear, data.batchs[index].startmonth)
  } else if (col == 2) {
    let item = data.batchs[index]
    let per = parseFloat(util.getFloatFixed(item.rate))
    let totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    let duration = util.getFloatFixed(getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth))
    let price = util.getFloatFixed(parseFloat(totalprice) / (util.getFloat(data.report.totalsize) * duration))

    data.batchs[index].totalprice = util.money(totalprice)    
    data.batchs[index].price = util.moneyfloat(price)

    data.batchs[index].duration = util.getInt(duration)
  } else if (col == 3) {
    let item = data.batchs[index]

    let duration = util.getFloatFixed(getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth))
    let per = (util.getFloatFixed(data.report.totalsize) * duration) * util.getFloatFixed(item.price) / (util.getFloatFixed(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears))) * 100

    let totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    data.batchs[index].rate = util.getFloatFixed(per)
    data.batchs[index].totalprice = util.money(totalprice)
  }

  let totalprice = 0
  let sum = 0.0
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    sum += util.getFloatFixed(item.rate)    
    data.batchs[i].totalrate = util.getFloatFixed(sum)

    if (i < data.batchs.length - 1) {
      totalprice += util.getInt(item.totalprice)
    }
  }

  if (data.batchs.length > 0 && data.batchs[data.batchs.length - 1].totalrate === 100.0) {    
    data.batchs[data.batchs.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - totalprice)
  }
}

function onKeyupPlan(index, col) {
  console.log('repair ------------------')
  console.log(data.repair)
  console.log('report ------------------')
  console.log(data.report)
    
  if (col == 1) {
    if (index >= data.batchplans.length - 1) {
      return
    }
    
    let endyear = util.getInt(data.batchplans[index].endyear)
    let endmonth = util.getInt(data.batchplans[index].endmonth)

    if (endmonth == 12) {
      data.batchplans[index + 1].startyear = endyear + 1
      data.batchplans[index + 1].startmonth = 1
    } else {
      data.batchplans[index + 1].startyear = endyear
      data.batchplans[index + 1].startmonth = endmonth + 1
    }

    data.batchplans[index].duration = getDurationMonth(data.batchplans[index].endyear, data.batchplans[index].endmonth, data.batchplans[index].startyear, data.batchplans[index].startmonth)
  } else if (col == 2) {
    let item = data.batchplans[index]
    console.log('item', item)
    let per = parseFloat(util.getFloatFixed(item.rate))
    console.log('per', per)
    let totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    console.log('totalprice', totalprice)
    let duration = util.getFloatFixed(getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth))
    console.log('duration', duration)
    let price = util.getFloatFixed(parseFloat(totalprice) / (util.getFloat(data.report.totalsize) * duration))
    console.log('duration', price)
    data.batchplans[index].totalprice = util.money(totalprice)    
    data.batchplans[index].price = util.moneyfloat(price)

    data.batchplans[index].duration = util.getInt(duration)
  } else if (col == 3) {
    let item = data.batchplans[index]
    console.log('item', item)
    let duration = util.getFloatFixed(getDurationMonth(item.endyear, item.endmonth, item.startyear, item.startmonth))
    console.log('duration', duration)

    let calc1 = (util.getFloatFixed(data.report.totalsize) * duration) * util.getFloatFixed(item.price)
    console.log('calc1', calc1)

    let calc2 = (util.getFloatFixed(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears)))
    console.log('calc2', calc2)

    let calc3 = (util.getFloatFixed(data.report.totalsize) * duration) * util.getFloatFixed(item.price) / (util.getFloatFixed(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears))) 
    console.log('calc3', calc3)
    
    let per = (util.getFloatFixed(data.report.totalsize) * duration) * util.getFloatFixed(item.price) / (util.getFloatFixed(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears))) * 100
    console.log('per', per)

    let totalprice = util.getInt((util.getFloatFixed(data.report.saveprice)) * parseFloat(util.getPlanyears(data.repair.planyears)) / 100 * per)
    console.log('totalprice', totalprice)
    data.batchplans[index].rate = util.getFloatFixed(per)
    data.batchplans[index].totalprice = util.money(totalprice)    
  }

  let totalprice = 0
  let sum = 0.0
  for (let i = 0; i < data.batchplans.length; i++) {
    let item = data.batchplans[i]
    sum += util.getFloatFixed(item.rate)
    data.batchplans[i].totalrate = util.getFloatFixed(sum)    

    if (i < data.batchplans.length - 1) {
      totalprice += util.getInt(item.totalprice)
    }
  }

  if (data.batchplans.length > 0 && data.batchplans[data.batchplans.length - 1].totalrate === 100.0) {    
    data.batchplans[data.batchplans.length - 1].totalprice = util.money(util.getInt(data.report.saveprice) * util.getPlanyears(data.repair.planyears) - totalprice)
  }
}

const getSummaries = (param: SummaryMethodProps) => {
  const columns = param.columns
  const items = param.data
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index == 1) {
      sums[index] = util.money(util.getFloatFixed(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears)))
    } else if (index == 2) {
      let total = 0
      if (items != null) {        
        items.forEach((item) => {
          total += util.getInt(item.totalprice)
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

const getSummariesPlan = (param: SummaryMethodProps) => {
  const columns = param.columns
  const items = param.data
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index == 1) {
      sums[index] = util.money(util.getFloatFixed(data.report.saveprice) * parseFloat(util.getPlanyears(data.repair.planyears)))
    } else if (index == 2) {
      let total = 0
      if (items != null) {        
        items.forEach((item) => {
          total += util.getInt(item.totalprice)
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

function clickHundred() {
  let total = 0.0
  let index = data.batchs.length - 1

  data.batchs.map((item, i) => {
    if (i == index) {
      return
    }

    total += util.getFloat(item.rate)
  })

  if (parseFloat(total) > 100.0) {
    return
    
  }

  data.batchs[index].rate = util.getFloatFixed(100.0 - total)

  onKeyup(index, 2)
}

function clickHundredPlan() {
  let total = 0.0
  let index = data.batchplans.length - 1

  data.batchplans.map((item, i) => {
    if (i == index) {
      return
    }

    total += util.getFloat(item.rate)
  })

  if (total > 100.0) {
    return
  }

  console.log('go')

  data.batchplans[index].rate = util.getFloatFixed(100.0 - total)

  onKeyupPlan(index, 2)
}

function cellClassNameTotal({row, columnIndex}) {
  if (row.index % 2 < 2) {
    return 'value'
  } else {
    return 'title'    
  }
}

const spanMethodTotal = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  if (columnIndex == 0 || columnIndex == 1 || columnIndex == 5 || columnIndex == 6) {
    if (rowIndex % 2 == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }

  return {rowspan: 1, colspan: 1}
}

const spanMethodTotalPlan = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  if (rowIndex == 10) {
    if (columnIndex == 2) {
      return {rowspan: 1, colspan: 3}
    } else if (columnIndex == 3 || columnIndex == 4) {
      return {rowspan: 0, colspan: 0}
    } else {
      return {rowspan: 1, colspan: 1}
    }
  }
  
  if (columnIndex == 0 || columnIndex == 1 || columnIndex == 5 || columnIndex == 6) {
    if (rowIndex % 2 == 0) {
      return {rowspan: 2, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }

  return {rowspan: 1, colspan: 1}
}

function getDuration(end, start) {
  end = util.getInt(end)
  start = util.getInt(start)

  if (end - start + 1 > 0) {
    return end - start + 1
  }

  return ''
}

function getDurationMonth(endyear, endmonth, startyear, startmonth) {
  endyear = util.getInt(endyear)
  startyear = util.getInt(startyear)
  endmonth = util.getInt(endmonth)
  startmonth = util.getInt(startmonth)

  if (startyear > endyear) {
    return ''
  }
  
  if (startyear == endyear) {
    if (startmonth > endmonth) {
      return ''
    }
    
    return endmonth - startmonth + 1
  }

  let total = 12 - startmonth + 1
  total += (endyear - startyear - 1) * 12
  total += endmonth

  return total
}

</script>
<style>
.inputNumber .el-input__inner {
  text-align: right;
}

.inputText .el-input__inner {
  text-align: left;
}

.title {
  background-color: #fafafa;
}

.value {
  background-color: #FFF;  
}

.date .el-input__wrapper {
  padding:0px;  
}

.date .el-input__inner {
  text-align:center;

}
</style>

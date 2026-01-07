<template>
  <Title title="세부내역" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
    
    
    <el-select v-model.number="data.search.dong" placeholder="시설물" style="width:80px;">           
      <el-option
        v-for="item in data.dongs"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>

    <el-tree-select style="width:200px;" v-model="data.search.category" :data="data.categorys" check-strictly :default-expand-all="false" :render-after-expand="false" placeholder="공사종별" />    

    <el-select v-model.number="data.search.standard" style="width:150px;" placeholder="규격">
      <el-option
        v-for="item in data.search.standards"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>

    <el-select v-model.number="data.search.method" placeholder="수선방법" style="width:90px;">           
      <el-option
        v-for="item in data.search.methods"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      />
    </el-select>
    

    <el-button size="small" class="filter-item" type="primary" @click="clickSearch">검색</el-button>
    <div style="flex:1;text-align:right;gap:5;">



      <div style="display:inline-flex;margin-right:30px;">
        <el-upload
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
          <el-button size="small" type="danger" @click="submitUpload">기존자료 비교</el-button>

        </el-upload>

      </div>

      <!--<el-button size="small" type="danger" @click="clickDeduplication" style="margin-right:-5px;">중복 삭제</el-button>-->
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="info" @click="clickChangeYear" style="margin-right:-5px;">년도 수정</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">개별 등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄 등록</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" :summary-method="getSummaries" show-summary  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="30" align="center" />
    <!--<el-table-column prop="index" label="NO" align="center" width="40" />-->      
    <el-table-column prop="name" label="시설물" align="center" width="70" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getDongElevator(scope.row.dong, scope.row.elevator)}}
      </template>
    </el-table-column>

    <el-table-column label="대분류" width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.topcategory).name}}
      </template>
    </el-table-column>
    <el-table-column label="중분류" width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.subcategory).name}}
      </template>
    </el-table-column>
    <el-table-column label="공사종별" width="90" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.category).name}}
      </template>
    </el-table-column>
    <el-table-column prop="extra.standard.name" label="규격" width="150" :show-overflow-tooltip="true" />      
    
    <el-table-column label="수선방법" align="center" width="60">
      <template #default="scope">
        {{getCategory(scope.row.method).name}}
      </template>
    </el-table-column>

    <el-table-column prop="extra.category.cycle" label="주기" align="center" width="40" />      
    <el-table-column prop="extra.category.percent" label="수선율" align="center" width="50" />
    <el-table-column prop="extra.standard.unit" label="단위" align="center" width="40" />
    <el-table-column prop="count" label="수량" align="right" width="60" />
    
    <el-table-column label="단가" align="right">
      <template #default="scope">
        <span v-if="scope.row.rate == 0">{{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}</span>
        <span v-else style="color:#af2020;">{{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}</span>
      </template>
    </el-table-column>

    <el-table-column label="수선금액" align="right">
      <template #default="scope">
        <span v-if="scope.row.rate == 0">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</span>
        <span v-else style="color:#af2020;">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</span>
      </template>
    </el-table-column>
    
    
    <el-table-column prop="lastdate" label="최종수선" align="center" />
    <el-table-column prop="duedate" label="수선예정" align="center" />

    <el-table-column prop="remark" label="기타" align="left" />
  </el-table>  

  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1100px"
  >
    <el-form :model="data.item" label-width="100px">

      <div style="text-align:left;">
        <el-tree-select style="width:290px;" v-model="data.item.category" :data="data.categorys" :default-expand-all="false" :render-after-expand="false" @node-click="changeCategory" placeholder="공사종별" />        
      </div>

      
      <el-table :data="[data.item]" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">        
        <el-table-column label="시설물" align="center" width="90">
          <template #default="scope">
            <el-select v-model.number="data.item.dong" class="m-2" placeholder="시설물" @change="changeDongForEvelator(scope.row)">           
              <el-option
                v-for="item in data.dongs"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column label="승강기" align="center" width="90">
          <template #default="scope">
            <el-select v-model.number="data.item.elevator" class="m-2" placeholder="승강기">           
              <el-option
                v-for="item in data.elevators"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column prop="standard" label="규격">
          <template #default="scope">
            <el-select v-model.number="data.item.standard" class="m-2" placeholder="규격" @change="changeStandard(scope.row)">
              <el-option
                v-for="item in data.standards"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
            
          </template>
        </el-table-column>
        <el-table-column prop="method" label="수선방법" align="center" width="100">
          <template #default="scope">
            <el-select v-model.number="data.item.method" class="m-2" placeholder="수선방법" @change="changeMethod(scope.row)">          
              <el-option
                v-for="item in data.methods"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>

          </template>
        </el-table-column>
        <el-table-column prop="cycle" label="주기" align="center" width="50" />
        <el-table-column prop="percent" label="수선율" align="center" width="60" />
        <el-table-column prop="unit" label="단위"  align="center" width="50" />
        <el-table-column label="수량" align="center" width="70">
          <template #default="scope">
            <el-input v-model.number="data.item.count" @keyup="onKeyupSingle(false)" />
          </template>
        </el-table-column>
        <el-table-column label="단가" align="right" width="80">
          <template #default="scope">
            <span v-if="scope.row.rate == 0">{{scope.row.price}}</span>
            <span v-else style="color:#af2020;">{{scope.row.price}}</span>
          </template>
        </el-table-column>        
        <el-table-column label="수선금액" width="100" align="right">
          <template #default="scope">            
            <span v-if="scope.row.rate == 0">{{scope.row.totalprice}}</span>
            <span v-else style="color:#af2020;">{{scope.row.totalprice}}</span>
          </template>
        </el-table-column>
        
        <el-table-column label="최종수선" align="center" width="65">
          <template #default="scope">
            <el-input v-model.number="data.item.lastdate" @keyup="onKeyupLastdateSingle" />
          </template>
        </el-table-column>
        <el-table-column label="수선예정" align="center" width="65">
          <template #default="scope">
            <el-input v-model.number="data.item.duedate" />
          </template>
        </el-table-column>
        <el-table-column label="기타">
          <template #default="scope">
            <el-input v-model="data.item.remark" />
          </template>
        </el-table-column>

      </el-table>


      <el-table :data="data.historys" border height="300px" :summary-method="getHistorySummaries" show-summary style="margin-top:15px;">
        <el-table-column prop="index" label="NO" align="center" width="60" sortable />
        <el-table-column prop="name" label="년도" align="center" width="100" sortable>
          <template #default="scope">
            {{scope.row.year}}년 {{util.pad(scope.row.month, 2)}}월
          </template>
        </el-table-column>
        <el-table-column prop="name" label="대분류" width="80" :show-overflow-tooltip="true">
          <template #default="scope">
            {{getCategory(scope.row.topcategory).name}}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="중분류"  width="80" :show-overflow-tooltip="true">
          <template #default="scope">
            {{getCategory(scope.row.subcategory).name}}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="공사종별"  width="150" :show-overflow-tooltip="true">
          <template #default="scope">
            {{getCategory(scope.row.category).name}}
          </template>
        </el-table-column>    
        <el-table-column prop="content" label="보수내역" />
        <el-table-column label="사용금액" align="right" width="100">
          <template #default="scope">
            {{util.money(scope.row.price)}}
          </template>
        </el-table-column>
      </el-table>


    </el-form>

    <template #footer>
      <el-button size="small" v-if="data.item.id > 0" style="float:left;" type="danger" @click="clickDelete(data.item)">삭제</el-button>
      <!--<el-button size="small" style="float:left;" @click="data.visibleCategory = true">공사종별</el-button>-->
      <el-button size="small" style="float:left;" @click="clickShowStandard">규격</el-button>
      <el-button size="small" style="float:left;" @click="clickShowHistory('normal')">사용현황</el-button>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleStandard"
    width="1100px"
    :before-close="handleCloseStandard"    
  >
    <StandardInsert :height="300" :category="data.item.category" ref="standardPopupRef" />
  </el-dialog>
  
  <el-dialog
    v-model="data.visibleCategory"
    width="1100px"
    :before-close="handleCloseCategory"
  >
    <CategoryInsert :height="300" />
  </el-dialog>  

  <el-dialog
    v-model="data.visibleHistory"
    width="1100px"
    title="사용현황"
  >

    <el-table :data="data.historys" border height="400px" :summary-method="getHistorySummaries" show-summary>
      <el-table-column prop="index" label="NO" align="center" width="60" sortable />
      <el-table-column prop="name" label="년도" align="center" width="100" sortable>
        <template #default="scope">
          {{scope.row.year}}년 {{util.pad(scope.row.month, 2)}}월
        </template>
      </el-table-column>
      <el-table-column prop="name" label="대분류" width="80" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getCategory(scope.row.topcategory).name}}
        </template>
      </el-table-column>
      <el-table-column prop="name" label="중분류" width="80" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getCategory(scope.row.subcategory).name}}
        </template>
      </el-table-column>
      <el-table-column prop="name" label="공사종별" width="150" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getCategory(scope.row.category).name}}
        </template>
      </el-table-column>    
      <el-table-column prop="content" label="보수내역" />
      <el-table-column label="사용금액" align="right" width="100">
        <template #default="scope">
          {{util.money(scope.row.price)}}
        </template>
      </el-table-column>
    </el-table>  

    <template #footer>    
      <el-button size="small" @click="data.visibleHistory = false">닫기</el-button>
    </template>
    
  </el-dialog>

  
  <el-dialog
    v-model="data.visibleBatch"
    :before-close="handleCloseBatch"
    :fullscreen="true"
  >

    <el-tabs size="small" v-model="data.menu" style="margin-top:-20px;">
      <el-tab-pane label="규격 선택" name="standard">
        

        

        <div style="display:flex;flex-direction:row;margin:-10px 0px -10px 0px;">
          <div style="width:150px;text-align:left;">
            

            <el-table :data="data.batchdongs" border height="268" @row-click="clickDong" style="margin-top:10px;" ref="dongRef" @selection-change="changeDong">
              <el-table-column type="selection" width="40" align="center" />
              <el-table-column prop="name" label="시설물" />          
            </el-table>
            <el-button style="margin-top:7px;width:100%;" size="small" type="info" @click="clickAddDong"><el-icon style="margin-right:5px;"><Plus /></el-icon> 시설물 추가</el-button>
            
          </div>
          <div style="flex:2;padding:10px 10px;">
            <el-table :data="data.batchtopcategorys" border height="300" highlight-current-row @row-click="clickTopcategory">
              <el-table-column prop="label" label="대분류" />
            </el-table>
          </div>

          <div style="flex:2;padding:10px 10px;">
            <el-table :data="data.batchsubcategorys" border height="300" highlight-current-row @row-click="clickSubcategory">
              <el-table-column prop="label" label="중분류" />
            </el-table>
          </div>

          <div style="flex:2;padding:10px 10px;">
            <el-table :data="data.batchcategorys" border height="300" highlight-current-row @row-click="clickCategory">
              <el-table-column prop="label" label="공사종별" />
            </el-table>
          </div>

          <div style="flex:3;padding:10px 10px;">
            <el-table :data="data.batchstandards" border height="300" @row-click="clickStandard" ref="standardRef" @selection-change="changeStandardData">
              <el-table-column type="selection" width="40" align="center" />
              <el-table-column prop="label" label="규격" />
            </el-table>
          </div>
          
          <div style="flex:2;">

            
            <el-form :model="data.item" label-width="80px">
              
              <el-form-item label="수선방법" style="margin-bottom:-12px;">        
                <el-select v-model.number="data.item.method" class="m-2" placeholder="수선방법" @change="changeMethodData(data.item)">          
                  <el-option
                    v-for="item in data.methods"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>

              <el-form-item label="수선주기" style="margin-bottom:-12px;font-size:12px;">
                {{data.item.cycle}}
              </el-form-item>

              <el-form-item label="수선율" style="margin-bottom:-12px;font-size:12px;">
                {{data.item.percent}}
              </el-form-item>

              <el-form-item label="단위" style="margin-bottom:-12px;font-size:12px;">
                {{data.item.unit}}
              </el-form-item>
              
              <el-form-item label="수량" style="margin-bottom:-12px;" @keyup="onKeyupSingle(true)">
                <el-input v-model.number="data.item.count" style="width:70px;" />
              </el-form-item>

              <el-form-item label="단가" style="margin-bottom:-12px;font-size:12px;">
                {{data.item.price}}
              </el-form-item>

              <el-form-item label="수선금액" style="margin-bottom:-12px;font-size:12px;">
                {{data.item.totalprice}}            
              </el-form-item>
              
              <el-form-item label="최종수선년도" style="margin-bottom:-12px;">
                <el-input v-model.number="data.item.lastdate" style="width:70px;" @keyup="onKeyupLastdateSingle" />
              </el-form-item>

              <el-form-item label="수선예정년도" style="margin-bottom:0px;">
                <el-input v-model.number="data.item.duedate" style="width:70px;" />
              </el-form-item>

            </el-form>

            <el-button style="margin-top:22px;width:100%;" size="small" type="warning" @click="clickBatchAdd">추가</el-button>
          </div>
        </div>
        
        
      </el-tab-pane>
      <el-tab-pane label="사용현황" name="history">
        <el-table :data="data.historys" border height="300px" :summary-method="getHistorySummaries" show-summary>
          <el-table-column prop="index" label="NO" align="center" width="60" sortable />
          <el-table-column prop="name" label="년도" align="center" width="100" sortable>
            <template #default="scope">
              {{scope.row.year}}년 {{util.pad(scope.row.month, 2)}}월
            </template>
          </el-table-column>
          <el-table-column prop="name" label="대분류" sortable>
            <template #default="scope">
              {{getCategory(scope.row.topcategory).name}}
            </template>
          </el-table-column>
          <el-table-column prop="name" label="중분류" sortable>
            <template #default="scope">
              {{getCategory(scope.row.subcategory).name}}
            </template>
          </el-table-column>
          <el-table-column prop="name" label="공사종별" sortable>
            <template #default="scope">
              {{getCategory(scope.row.category).name}}
            </template>
          </el-table-column>    
          <el-table-column prop="content" label="보수내역" />
          <el-table-column label="사용금액" align="right" width="100">
            <template #default="scope">
              {{util.money(scope.row.price)}}
            </template>
          </el-table-column>
        </el-table>

      </el-tab-pane>    
    </el-tabs>
    


    
    <el-form label-width="100px" style="margin-bottom:-27px;">

      <el-table :data="data.batchs" border :height="data.batchheight" :key="data.width+''+data.height" style="margin-top:15px;" @row-click="changeBatchInsert">
        <el-table-column label="" align="center" width="35">
          <template #default="scope">
            <el-icon @click="clickBatchRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="시설물" align="center" width="110">
          <template #default="scope">
            {{getDongElevator(scope.row.dong, scope.row.elevator)}}
          </template>
        </el-table-column>
        <el-table-column label="공사종별">
          <template #default="scope">
            {{getCategory(scope.row.category).name}}
          </template>
        </el-table-column>
        <el-table-column prop="standard" label="규격">
          <template #default="scope">
            {{getStandard(scope.row.standard)}}
          </template>
        </el-table-column>
        <el-table-column prop="method" label="수선방법" align="center" width="60">
          <template #default="scope">
            {{getCategory(scope.row.method).name}}
          </template>
        </el-table-column>
        <el-table-column prop="cycle" label="주기" align="center" width="40" />
        <el-table-column prop="percent" label="수선율" align="center" width="45" />
        <el-table-column prop="unit" label="단위"  align="center" width="40" />
        <el-table-column label="수량" align="center" width="70">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].count" @keyup="onKeyup(scope.$index)" />
          </template>
        </el-table-column>
        <el-table-column label="단가" align="right" width="80">
          <template #default="scope">
            <span v-if="scope.row.rate == 0">{{util.money(scope.row.price)}}</span>
            <span v-else style="color:#af2020;">{{util.money(scope.row.price)}}</span>
          </template>
        </el-table-column>
        <el-table-column label="수선금액" width="100" align="right">
          <template #default="scope">
            <span v-if="scope.row.rate == 0">{{scope.row.totalprice}}</span>
            <span v-else style="color:#af2020;">{{scope.row.totalprice}}</span>
          </template>
        </el-table-column>
        <el-table-column label="최종수선" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].lastdate" @keyup="onKeyupLastdate(scope.$index)" />
          </template>
        </el-table-column>
        <el-table-column label="수선예정" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].duedate" />
          </template>
        </el-table-column>
        <el-table-column label="기타" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].remark" />
          </template>
        </el-table-column>

      </el-table>


    </el-form>

    
    <template #footer>
      <!--<el-button size="small" style="float:left;" @click="data.visibleCategory = true">공사종별</el-button>-->
      <el-button size="small" style="float:left;" @click="clickShowStandard">규격</el-button>
      <el-button size="small" style="float:left;" @click="clickShowHistory('batch')">사용현황</el-button>
      <el-button size="small" @click="data.visibleBatch = false">취소</el-button>      
      <el-button size="small" type="primary" @click="clickBatchSubmit">등록</el-button>
    </template>
  </el-dialog>


  <el-dialog
    v-model="data.visibleYear"
    :before-close="handleClose"
    width="1100px"
    title="수선예정년도 수정"
  >

    <el-table :data="data.batchs" border :height="height(450)" :key="data.width+''+data.height" style="margin-top:0px;" :span-method="spanMethod">
      <el-table-column label="시설물" align="center" width="70" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getDongElevator(scope.row.dong, scope.row.elevator)}}
        </template>
      </el-table-column>
      <el-table-column label="대분류" width="80" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getCategory(scope.row.topcategory).name}}
        </template>
      </el-table-column>
      <el-table-column label="중분류" width="80" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getCategory(scope.row.subcategory).name}}
        </template>
      </el-table-column>
      <el-table-column label="공사종별" width="90" :show-overflow-tooltip="true">
        <template #default="scope">
          {{getCategory(scope.row.category).name}}
        </template>
      </el-table-column>
      <el-table-column label="규격">
        <template #default="scope">
          {{getStandard(scope.row.standard)}}

        </template>
      </el-table-column>

      <el-table-column label="수선방법" align="center" width="60">
        <template #default="scope">
          {{getCategory(scope.row.method).name}}
        </template>
      </el-table-column>

      <el-table-column prop="extra.category.cycle" label="주기" align="center" width="40" />
      <el-table-column prop="extra.category.percent" label="수선율" align="center" width="50" />
      <el-table-column prop="extra.standard.unit" label="단위"  align="center" width="40" />
      <el-table-column prop="count" label="수량" align="center" width="50" />
      <el-table-column label="단가" align="right" width="80">
        <template #default="scope">
          <span v-if="scope.row.rate == 0">{{scope.row.price}}</span>
          <span v-else style="color:#af2020;">{{scope.row.price}}</span>
        </template>
      </el-table-column>
      <el-table-column label="수선금액" width="100" align="right">
        <template #default="scope">
          <span v-if="scope.row.rate == 0">{{scope.row.totalprice}}</span>
          <span v-else style="color:#af2020;">{{scope.row.totalprice}}</span>
        </template>
      </el-table-column>

      <el-table-column prop="lastdate" label="최종수선" align="center" width="70">
        <template #default="scope">
          <span v-if="scope.$index < data.batchs.length - 1">{{scope.row.lastdate}}</span>
          <span v-else>
            <el-input v-model.number="data.lastdate" />
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="duedate" label="수선예정" align="center" width="70">
        <template #default="scope">
          <span v-if="scope.$index < data.batchs.length - 1">{{scope.row.duedate}}</span>
          <span v-else>
            <el-input v-model.number="data.duedate" />
          </span>
        </template>
      </el-table-column>

    </el-table>



    <template #footer>
      <el-button size="small" @click="data.visibleYear = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitYear">등록</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleAddDong"
    width="800px"
  >

    <el-form label-width="100px">      
      <el-table :data="[data.dong]" border style="margin-top:15px;">        
        <el-table-column label="시설물" align="center">
          <template #default="scope">
            <el-input v-model="data.dong.name" />
          </template>
        </el-table-column>
        <el-table-column label="지상층" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.ground" />
          </template>
        </el-table-column>
        <el-table-column label="지하층" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.underground" />
          </template>
        </el-table-column>
        <el-table-column label="주차장" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.parking" />
          </template>
        </el-table-column>
        <el-table-column label="승강기" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.elevator" />
          </template>
        </el-table-column>
        <el-table-column label="세대수" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.familycount" />
          </template>
        </el-table-column>
        <el-table-column label="공용/전용" align="center" width="80">
          <template #default="scope">
            <el-select v-model.number="data.dong.basic" placeholder="공용/전용" style="width:70px;">           
              <el-option
                v-for="item in data.basics"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column label="순번" align="center" width="100">
          <template #default="scope">
            <el-input v-model.number="data.dong.order" />
          </template>
        </el-table-column>        
        <el-table-column label="비고" align="center">
          <template #default="scope">
            <el-input v-model="data.dong.remark" />
          </template>
        </el-table-column>

      </el-table>


    </el-form>

    <template #footer>        
      <el-button size="small" @click="clickAddDongCancel">취소</el-button>
      <el-button size="small" type="primary" @click="clickAddDongSubmit">등록</el-button>
    </template>
  </el-dialog>

  
  <el-dialog
    v-model="data.visibleDong"
    :before-close="handleClose"
    width="800px"
  >

    <el-form label-width="100px">      
      <el-table :data="[1]" border style="margin-top:15px;">        
        <el-table-column label="시설물" align="center">
          <template #default="scope">
            <el-input v-model="data.dong.name" />
          </template>
        </el-table-column>
        <el-table-column label="지상층" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.ground" />
          </template>
        </el-table-column>
        <el-table-column label="지하층" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.underground" />
          </template>
        </el-table-column>
        <el-table-column label="주차장" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.parking" />
          </template>
        </el-table-column>
        <el-table-column label="승강기" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.elevator" />
          </template>
        </el-table-column>
        <el-table-column label="세대수" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.dong.familycount" />
          </template>
        </el-table-column>
        <el-table-column label="공용/전용" align="center" width="80">
          <template #default="scope">
            <el-select v-model.number="data.dong.basic" placeholder="공용/전용" style="width:70px;">           
              <el-option
                v-for="item in data.basics"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column label="순번" align="center" width="100">
          <template #default="scope">
            <el-input v-model.number="data.dong.order" />
          </template>
        </el-table-column>        
        <el-table-column label="비고" align="center">
          <template #default="scope">
            <el-input v-model="data.dong.remark" />
          </template>
        </el-table-column>

      </el-table>


    </el-form>

    <template #footer>
      <el-button size="small" @click="data.visibleDong = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitDong">등록</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visiblePrice"
    :before-close="handleClosePrice"
    width="1000px"
  >

    <el-table :data="[data.price]" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">
      <el-table-column prop="name" label="대분류" width="100">
        <template #default="scope">
          {{getCategory(data.price.topcategory).name}}
        </template>
      </el-table-column>

      <el-table-column prop="name" label="중분류별" width="100">
        <template #default="scope">
          {{getCategory(data.price.subcategory).name}}
        </template>
      </el-table-column>

      <el-table-column prop="name" label="공사종별" width="100">
        <template #default="scope">
          {{getCategory(data.price.category).name}}
        </template>
      </el-table-column>

      <el-table-column label="규격명">
        <template #default="scope">
          {{data.price.extra.standard.name}}
        </template>
      </el-table-column>

      <el-table-column label="재료비" align="center" width="100">
        <template #default="scope">
          <el-input v-model="data.standard.direct"  @keyup="onKeyupStandard" />
        </template>
      </el-table-column>
      <el-table-column label="노무비" align="center" width="100">
        <template #default="scope">
          <el-input v-model="data.standard.labor"  @keyup="onKeyupStandard" />
        </template>
      </el-table-column>
      <el-table-column label="경비" align="center" width="100">
        <template #default="scope">
          <el-input v-model="data.standard.cost"  @keyup="onKeyupStandard" />
        </template>
      </el-table-column>

      <el-table-column label="단가" align="right" width="100">
        <template #default="scope">
          <span v-if="data.price.rate == 0">{{util.money(util.calculatePriceRate(data.standard.direct, data.standard.labor, data.standard.cost, data.price.rate, data.parcelrate))}}</span>
          <span v-else style="color:#af2020;">{{util.money(util.calculatePriceRate(data.standard.direct, data.standard.labor, data.standard.cost, data.price.rate, data.parcelrate))}}</span>            
        </template>
      </el-table-column>

    </el-table>


    
    <template #footer>
      <el-button size="small" @click="data.visiblePrice = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitPrice">등록</el-button>
    </template>
  </el-dialog>


  <el-dialog
    v-model="data.visiblePeriod"
    :before-close="handleClosePeriod"
    width="1000px"
  >

    <el-table :data="[data.period]" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">
      <el-table-column prop="name" label="대분류">
        <template #default="scope">
          {{getCategory(data.period.topcategory).name}}
        </template>
      </el-table-column>

      <el-table-column prop="name" label="중분류별">
        <template #default="scope">
          {{getCategory(data.period.subcategory).name}}
        </template>
      </el-table-column>

      <el-table-column prop="name" label="공사종별">
        <template #default="scope">
          {{getCategory(data.period.category).name}}
        </template>
      </el-table-column>        

      <el-table-column label="규격명">
        <template #default="scope">
          {{data.period.extra.standard.name}}
        </template>
      </el-table-column>

      <el-table-column label="수선방법" align="center" width="60">
        <template #default="scope">
          {{getCategory(data.period.method).name}}
        </template>
      </el-table-column>

      <el-table-column label="주기" align="center" width="80">
        <template #default="scope">
          <el-input v-model="data.method.cycle" />
        </template>
      </el-table-column>        

      <!--
           <el-table-column label="단가" align="right" width="100">
           <template #default="scope">
           <span v-if="data.period.rate == 0">{{util.money(util.calculatePriceRate(data.period.extra.standard.direct, data.period.extra.standard.labor, data.period.extra.standard.cost, data.period.rate, data.parcelrate))}}</span>
           <span v-else style="color:#af2020;">{{util.money(util.calculatePriceRate(data.period.extra.standard.direct, data.period.extra.standard.labor, data.period.extra.standard.cost, data.period.rate, data.parcelrate))}}</span>            
           </template>
           </el-table-column>

           <el-table-column label="수선금액" align="right">
           <template #default="scope">
           <span v-if="scope.row.rate == 0">{{util.money(util.calculateRepair(data.period.extra.standard.direct, data.period.extra.standard.labor, data.period.extra.standard.cost, data.period.rate, data.parcelrate, data.period.count, data.period.extra.category.percent))}}</span>
           <span v-else style="color:#af2020;">{{util.money(util.calculateRepair(data.period.extra.standard.direct, data.period.extra.standard.labor, data.period.standard.cost, data.period.rate, data.parcelrate, data.period.count, data.period.extra.category.percent))}}</span>
           </template>
           </el-table-column>
      -->
    </el-table>


    
    <template #footer>
      <el-button size="small" @click="data.visiblePeriod = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitPeriod">등록</el-button>
    </template>
  </el-dialog>


  <el-dialog
    v-model="diff.visible"
    :fullscreen="true"    
  >

    <div style="font-weight:bold;text-align:left;margin-bottom:5px;">삭제된 규격</div>
    
    <el-table :data="diff.newStandards" border :height="100" ref="newStandardRef" @selection-change="changeNewStandard">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" />


    <el-table-column label="대분류">
      <template #default="scope">
        {{getTopcategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column label="중분류">
      <template #default="scope">
        {{getSubcategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column label="공사종별">
      <template #default="scope">
        {{getCategory2(scope.row.category)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="규격명" />
    
    <el-table-column label="재료비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.direct)}}
      </template>
    </el-table-column>
    
    <el-table-column label="노무비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.labor)}}
      </template>
    </el-table-column>
    
    <el-table-column label="경비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.cost)}}
      </template>
    </el-table-column>
    
    <el-table-column label="단가" align="right" width="90">
      <template #default="scope">
        {{util.money(util.calculatePrice(scope.row.direct, scope.row.labor, scope.row.cost))}}
      </template>
    </el-table-column>

    <el-table-column prop="unit" label="규격" align="center" width="60" />
    <el-table-column prop="order" label="순번" align="center" width="60" />
    
    </el-table>    


    <div style="font-weight:bold;text-align:left;margin-bottom:5px;margin-top:10px;">변경된 규격</div>    
    
    <el-table :data="diff.changeStandards" border :height="100" ref="changeStandardRef" @selection-change="changeChangeStandard">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" />


    <el-table-column label="대분류">
      <template #default="scope">
        {{getTopcategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column label="중분류">
      <template #default="scope">
        {{getSubcategory(scope.row.category)}}
      </template>
    </el-table-column>
    <el-table-column label="공사종별">
      <template #default="scope">
        {{getCategory2(scope.row.category)}}
      </template>
    </el-table-column>

    <el-table-column prop="name" label="규격명" />
    
    <el-table-column label="재료비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.direct)}}
      </template>
    </el-table-column>
    
    <el-table-column label="노무비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.labor)}}
      </template>
    </el-table-column>
    
    <el-table-column label="경비" align="right" width="90">
      <template #default="scope">
        {{util.money(scope.row.cost)}}
      </template>
    </el-table-column>
    
    <el-table-column label="단가" align="right" width="90">
      <template #default="scope">
        {{util.money(util.calculatePrice(scope.row.direct, scope.row.labor, scope.row.cost))}}
      </template>
    </el-table-column>

    <el-table-column prop="unit" label="규격" align="center" width="60" />
    <el-table-column prop="order" label="순번" align="center" width="60" />
    
    </el-table>    

     <div style="font-weight:bold;text-align:left;margin-bottom:5px;margin-top:10px;">추가 세부내역</div>
    
     <el-table :data="diff.remainBreakdowns" border :height="200" ref="remainBreakdownRef" @selection-change="changeRemainBreakdown">
    <!--<el-table-column prop="index" label="NO" align="center" width="40" />-->      
    <el-table-column prop="name" label="시설물" align="center" width="70" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getDongElevator(scope.row.dong, scope.row.elevator)}}
      </template>
    </el-table-column>

    <el-table-column label="대분류" width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.topcategory).name}}
      </template>
    </el-table-column>
    <el-table-column label="중분류" width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.subcategory).name}}
      </template>
    </el-table-column>
    <el-table-column label="공사종별" width="90" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.category).name}}
      </template>
    </el-table-column>
    <el-table-column prop="extra.standard.name" label="규격" width="150" :show-overflow-tooltip="true" />      
    
    <el-table-column label="수선방법" align="center" width="60">
      <template #default="scope">
        {{getCategory(scope.row.method).name}}
      </template>
    </el-table-column>

    <el-table-column prop="extra.category.cycle" label="주기" align="center" width="40" />      
    <el-table-column prop="extra.category.percent" label="수선율" align="center" width="50" />
    <el-table-column prop="extra.standard.unit" label="단위" align="center" width="40" />
    <el-table-column prop="count" label="수량" align="right" width="60" />
    
    <el-table-column label="단가" align="right">
      <template #default="scope">
        <span v-if="scope.row.rate == 0">{{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}</span>
        <span v-else style="color:#af2020;">{{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}</span>
      </template>
    </el-table-column>

    <el-table-column label="수선금액" align="right">
      <template #default="scope">
        <span v-if="scope.row.rate == 0">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</span>
        <span v-else style="color:#af2020;">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</span>
      </template>
    </el-table-column>
    
    
    <el-table-column prop="lastdate" label="최종수선" align="center" />
    <el-table-column prop="duedate" label="수선예정" align="center" />

    <el-table-column prop="remark" label="기타" align="left" />
     </el-table>  


     
    <div style="font-weight:bold;text-align:left;margin-bottom:5px;margin-top:10px;">삭제된 세부내역</div>
    
    <el-table :data="diff.newBreakdowns" border :height="200" ref="newBreakdownRef" @selection-change="changeNewBreakdown">
    <el-table-column type="selection" width="30" align="center" />
    <!--<el-table-column prop="index" label="NO" align="center" width="40" />-->      
    <el-table-column prop="name" label="시설물" align="center" width="70" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getDongElevator(scope.row.dong, scope.row.elevator)}}
      </template>
    </el-table-column>

    <el-table-column label="대분류" width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.topcategory).name}}
      </template>
    </el-table-column>
    <el-table-column label="중분류" width="80" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.subcategory).name}}
      </template>
    </el-table-column>
    <el-table-column label="공사종별" width="90" :show-overflow-tooltip="true">
      <template #default="scope">
        {{getCategory(scope.row.category).name}}
      </template>
    </el-table-column>
    <el-table-column prop="extra.standard.name" label="규격" width="150" :show-overflow-tooltip="true" />      
    
    <el-table-column label="수선방법" align="center" width="60">
      <template #default="scope">
        {{getCategory(scope.row.method).name}}
      </template>
    </el-table-column>

    <el-table-column prop="extra.category.cycle" label="주기" align="center" width="40" />      
    <el-table-column prop="extra.category.percent" label="수선율" align="center" width="50" />
    <el-table-column prop="extra.standard.unit" label="단위" align="center" width="40" />
    <el-table-column prop="count" label="수량" align="right" width="60" />
    
    <el-table-column label="단가" align="right">
      <template #default="scope">
        <span v-if="scope.row.rate == 0">{{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}</span>
        <span v-else style="color:#af2020;">{{util.money(util.calculatePriceRate(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate))}}</span>
      </template>
    </el-table-column>

    <el-table-column label="수선금액" align="right">
      <template #default="scope">
        <span v-if="scope.row.rate == 0">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</span>
        <span v-else style="color:#af2020;">{{util.money(util.calculateRepair(scope.row.extra.standard.direct, scope.row.extra.standard.labor, scope.row.extra.standard.cost, scope.row.rate, data.parcelrate, scope.row.count, scope.row.extra.category.percent))}}</span>
      </template>
    </el-table-column>
    
    
    <el-table-column prop="lastdate" label="최종수선" align="center" />
    <el-table-column prop="duedate" label="수선예정" align="center" />

    <el-table-column prop="remark" label="기타" align="left" />
     </el-table>  



     
    <template #footer>
      <el-button size="small" @click="diff.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitDiff">등록</el-button>
    </template>
    
  </el-dialog>


  

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Repair, Dong, Category, Breakdown, Standard, History, Adjust, Upload } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'
import type { UploadInstance } from 'element-plus'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

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
  dong: null,
  topcategory: null,
  subcategory: null,
  category: null,
  method: null,
  count: 0,
  standard: null,  
  lastdate: '',
  duedate: '',
  remark: '',
  elevator: null
}

const dong = {
  id: 0,
  name: '',
  ground: 0,
  underground: 0,
  parking: 0,
  elevator: 0,
  familycount: 0,
  remark: '',
  basic: 1
}

const diff = reactive({
  visible: false,
  newStandards: [],
  changeStandards: [],
  newBreakdowns: [],
  remainBreakdowns: []
})

const data = reactive({
  apt: 0,
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  visibleBatch: false,
  visibleStandard: true,
  visibleCategory: false,
  visibleHistory: false,
  visibleYear: false,
  visibleDong: false,
  visiblePrice: false,
  visiblePeriod: false,
  dongs: [],
  batchdongs: [],
  standards: [],
  methods: [],
  allcategorys: [],
  allstandards: [],
  search: {
    dong: null,    
    category: null,
    standard: null,
    method: null,
    standards: [],
    methods: []    
  },
  elevators: [],
  batchs: [{category: 0}],
  repair: null,
  topcategorys: [],
  categorys: [],
  height: 0,
  batchheight: 0,
  save: null,
  historys: [],
  batchcategorys: [],
  batchsubcategorys: [],
  batchcategorys: [],
  batchstandards: [],
  menu: 'standard',
  categoryMap: {},
  adjusts: [],
  duedate: '',
  lastdate: '',
  parcelrate: 0,
  dong: util.clone(dong),
  visibleAddDong: false,
  basics: [{id: 0, name: ''}, {id: 1, name: '전용'}, {id: 2, name: '공용'}],
  index: -1,
  price: null,  
  standard: null,
  period: null,
  method: null,
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`,
  filename: ''
})

async function initData() {
  data.repair = await util.getRepair(data.apt)

  let {allcategorys, categorys} = await util.getCategoryTree(data.apt, '공사종별')
  data.allcategorys = allcategorys 
  data.categorys = categorys

  let categoryMap = {}
  allcategorys.forEach((item) => {
    categoryMap[item.id] = item
  })
  data.categoryMap = categoryMap
  
  let topcategorys = []
  categorys.forEach((item) => {
    if (item.id == 0) {
      return
    }

    topcategorys.push(item)
  })

  data.topcategorys = topcategorys
  data.batchtopcategorys = topcategorys.splice(0, 1)
  
  let res = await Dong.findByApt(data.apt)
  if (res.items == null) {
    res.items = []
  }
  
  data.dongs = [{id: 0, name: '시설물'}, ...res.items, {id: -1, name: '시설물 등록'}]
  data.batchdongs = res.items

  res = await Standard.findByApt(data.apt)
  data.allstandards = res.items  
}

async function getItems() {
  if (data.apt == 0) {
    return
  }

  util.loading(true)
  
  let topcategory = 0
  let subcategory = 0
  let category = 0
  
  let searchCategory = getCategory(data.search.category)
  if (searchCategory.level == 1) {
    topcategory = searchCategory.id
  } else if (searchCategory.level == 2) {
    subcategory = searchCategory.id
  } else if (searchCategory.level == 3) {
    category = searchCategory.id
  }

  let res = await Repair.get(data.apt)
  // data.parcelrate = res.item.parcelrate

  if (data.parcelrate == 0) {
    data.parcelrate = 100.0
  }

  res = await Breakdown.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    //orderby: 'd_order,d_id, b_elevator,c_order,c_id,s_order,s_id,b_lastdate,b_id',
    orderby: 'b_id',
    dong: data.search.dong,
    topcategory: topcategory,
    subcategory: subcategory,
    category: category,
    standard: data.search.standard,
    method: data.search.method
  })

  if (res.items != null) {    
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }
  }

  data.total = res.total
  data.items = res.items

  res = await Adjust.find({
    apt: data.apt,
    orderby: 'aj_order,aj_id'    
  })

  if (res.items == null) {
    res.items = []
  }
  data.adjusts = res.items

  util.loading(false)
}

function getRate(item) {
  let rate = 0

  let category = getCategory(item.category)
  let subcategory = getCategory(category.parent)
  let topcategory = getCategory(subcategory.parent)

  data.adjusts.forEach(v => {
    if (v.category != 0) {
      let c = getCategory(v.category)

      if (c.level == 1) {
        if (topcategory.id != c.id) {
          return
        }
      } else if (c.level == 2) {
        if (subcategory.id != c.id) {
          return
        }
      } else if (c.level == 3) {
        if (item.category != c.id) {
          return
        }
      }      
    }

    if (v.standard != 0) {
      if (item.standard != v.standard) {
        return
      }
    }

    rate = v.rate
  })

  return rate
}  

function clickInsert() {  
  data.item = util.clone(item)

  if (data.repair.completionyear > 0) {
    data.item.lastdate = data.repair.completionyear
  }
  
  data.item.cycle = ''
  data.item.percent = ''
  data.item.unit = ''
  
  data.methods = []
  data.standards = []
  data.historys = []
  
  data.visible = true
}

async function clickUpdate(item, index) {
  if (index.no == 0) {
    return
  }

  if (index.no == 7) {
    let res = await Category.get(item.method)

    data.method = res.item
    data.visiblePeriod = true
    data.period = util.clone(item)
    return
  }
  
  if (index.no == 11) {
    let res = await Standard.get(item.standard)
    console.log(res)

    data.standard = res.item
    data.visiblePrice = true
    data.price = util.clone(item)
    return
  }

  data.item = util.clone(item)
  
  changeCategory({value: item.category}, false)
  
  data.item.category = item.category
  data.item.method = item.method
  data.item.standard = item.standard

  changeMethod(data.item, false)
  data.visible = true
}

function clickDelete(item) {
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Breakdown.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      getItems()
    }
  })
}

async function clickSubmit() {
  const item = data.item
  item.apt = data.apt

  if (util.getInt(item.category) == 0) {
    util.error('공사종별을 선택하세요')
    return    
  }

  if (util.getInt(item.standard) == 0) {
    util.error('규격을 선택하세요')
    return    
  }

  if (util.getInt(item.method) == 0) {
    util.error('수선방법을 선택하세요')
    return    
  }

  if (item.lastdate == '') {
    util.error('최종수선년도를 입력하세요')
    return    
  }

  if (item.duedate == '') {
    util.error('수선예정년도를 입력하세요')
    return    
  }

  if (util.getInt(item.duedate) >= 2100) {
    util.error('수선예정년도를 정확하게 입력하세요')
    return    
  }

  

  let category = getCategory(item.category)
  let subcategory = getCategory(category.parent)
  let topcategory = getCategory(subcategory.parent)

  item.subcategory = subcategory.id
  item.topcategory = topcategory.id

  item.dong = util.getInt(item.dong)
  item.count = util.getInt(item.count)
  item.lastdate = util.getInt(item.lastdate)
  item.duedate = util.getInt(item.duedate)
  item.elevator = util.getInt(item.elevator)
  
  let res;

  if (item.id === 0) {
    res = await Breakdown.insert(item)
  } else {
    res = await Breakdown.update(item)
  }

  if (res.code === 'ok') {
    util.info('등록되었습니다')
    getItems()
    data.visible = false

    if (item.id == 0) {
      setTimeout(function() {
        listRef.value!.setScrollTop(data.items.length  * 100 + 1000)
      }, 500)
    }
  } else {
    util.error('오류가 발생했습니다')
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

const handleCloseBatch = (done: () => void) => {
  util.confirm('팝업창을 닫으시겠습니까', function() {
    done()
  })
}

function setWindowSize() {
  data.height = (window.innerHeight - 170) + 'px'
  data.batchheight = (window.innerHeight - 170 - 290) + 'px'
}

watch(() => route.params.id, async () => {
  data.visibleBatch = false
  data.visibleStandard = false

  data.id = util.getInt(route.params.id)

  await initData()
  await getItems()
})

onMounted(async () => {
  data.visibleBatch = false
  data.visibleStandard = false
  
  data.apt = parseInt(route.params.id)

  await initData()
  await getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)    
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
})

function getDong(id) {
  for (let i = 0; i < data.dongs.length; i++) {
    let item = data.dongs[i]

    if (item.id == id) {
      return item.name
    }
  }

  return ''
}

function getDongElevator(id, elevator) {
  if (elevator == null || elevator == 0) {
    return getDong(id)
  }

  for (let i = 0; i < data.dongs.length; i++) {
    let item = data.dongs[i]

    if (item.id == id) {
      return item.name + ` ${elevator}호기`
    }
  }

  return ''
}

function getCategory(id) {
  let item = data.categoryMap[id]

  if (item == null || item == undefined) {
    return {
      id: 0,
      name: ''
    } 
  }

  return item  
}

function getStandard(id) {
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.id == id) {
      return item.name
    }
  }

  return ''
}

function getStandardInfo(id) {
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.id == id) {
      return item
    }
  }

  return {}
}

function changeCategory(item, updateFlag) {
  data.item.category = item.value
  let id = data.item.category

  let methods = [{id: 0, name: '수선방법'}]
  let standards = [{id: 0, name: '규격'}]
  
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == id) {
      methods.push(item)
    }
  }

  data.methods = methods

  let flag = false
  
  if (methods.length == 2) {
    flag = true
    data.item.method = methods[1].id
  } else { 
    data.item.method = null
  }

  data.item.cycle = ''
  data.item.percent = ''
  data.item.unit = ''
  
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.category == id) {
      standards.push(item)
    }
  }

  data.standards = standards
  if (standards.length == 2) {
    flag = true
    data.item.standard = standards[1].id
  } else {
    data.item.standard = null
  }
  
  changeMethod(data.item, updateFlag)

  readHistory(data.item.category)
}

function changeSearchCategory(item) {
  data.search.category = item.value
  let id = data.search.category
  
  let category = getCategory(id)
  
  let methods = [{id: 0, name: '수선방법'}]
  let standards = [{id: 0, name: '규격'}]
  
  if (category.level == 3) {
    for (let i = 0; i < data.allcategorys.length; i++) {
      let item = data.allcategorys[i]

      if (item.parent == id) {
        methods.push(item)
      }
    }
    
    for (let i = 0; i < data.allstandards.length; i++) {
      let item = data.allstandards[i]

      if (item.category == id) {
        standards.push(item)
      }
    }
  }
  
  data.search.methods = methods
  data.search.method = null
  data.search.standards = standards
  data.search.standard = null
  
  clickSearch()
}

const getSummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index === 1) {
      sums[index] = '합계'
    } else if (index == 12) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += util.calculateRepair(item.extra.standard.direct, item.extra.standard.labor, item.extra.standard.cost, item.rate, data.parcelrate, item.count, item.extra.category.percent)
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

function clickBatch() {
  data.menu = 'standard'
  data.item = util.clone(item)
  data.item.count = null

  if (data.repair.completionyear > 0) {
    data.item.lastdate = data.repair.completionyear
  }

  data.methods = []
  data.standards = []

  /*
     if (data.save != null) {
     data.item = data.save.item
     data.methods = data.save.methods
     data.standards = data.save.standards
     }
   */
  
  data.visibleBatch = true

  data.batchs = []

  if (dongRef.value != null) {    
    dongRef.value!.clearSelection()
    //dongRef.value!.toggleAllSelection(true)

    /*
       for (let i = 0; i < data.batchdongs.length; i++) {
       let dong = data.batchdongs[i]

       if (dong.basic == 1) {
       dongRef.value!.toggleRowSelection(dong, undefined)
       }
       }
     */
  }

  data.batchtopcategorys = util.clone(data.topcategorys)
  data.batchsubcategorys = []
  data.batchcategorys = []
  data.batchstandards = []
}

const dongRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref([])
const toggleSelection = (rows) => {
  if (rows) {
    rows.forEach((row) => {      
      dongRef.value!.toggleRowSelection(row, undefined)
    })
  } else {
    dongRef.value!.clearSelection()
  }
}
const changeDong = (val) => {
  multipleSelection.value = val
}

function clickDong(item) {
  dongRef.value!.toggleRowSelection(item, undefined)
}

function clickBatchAdd() {
  const item = data.item
  item.apt = data.apt

  if (multipleSelection.value.length == 0) {
    util.error('시설물을 선택하세요')
    return
  }

  if (util.getInt(item.category) == 0) {
    util.error('공사종별을 선택하세요')
    return
  }

  if (multipleSelectionStandard.value.length == 0) {
    util.error('규격을 선택하세요')
    return
  }

  item.count = util.getInt(item.count)

  /*
     let standard = null

     if (multipleSelectionStandard.value.length == 1) {
     for (let i = 0; i < data.allstandards.length; i++) {
     let standardItem = data.allstandards[i]

     if (standardItem.id == multipleSelectionStandard.value[0].value) {
     standard = standardItem
     break
     }
     }
     }
   */

  let elevator = false
  let ca = getCategory(item.category)  
  if (ca.elevator == 1) {
    elevator = true
  }

  let items = []

  multipleSelectionStandard.value.forEach((standardInfo) => {
    let standard = null

    for (let i = 0; i < data.allstandards.length; i++) {
      let standardItem = data.allstandards[i]

      if (standardItem.id == standardInfo.value) {
        standard = standardItem
        break
      }
    }

    if (standard == null) {
      return
    }

    multipleSelection.value.forEach((value) => {
      let start = 0
      let end = 1

      if (elevator == true) {
        start = 1
        end = value.elevator + 1
      }

      for (let i = start; i < end; i++) {
        let methodEnd = data.methods.length

        if (methodEnd == 4) {
          methodEnd = 3
        }        

        for (let j = 1; j < methodEnd; j++) {
          if (item.method != null) {
            if (item.method > 0 && data.methods[j].id != item.method) {
              continue
            }
          }

          let category = null

          for (let k = 0; k < data.allcategorys.length; k++) {
            if (data.methods[j].id == data.allcategorys[k].id) {
              category = data.allcategorys[k]
              break
            }
          }

          let n = {
            once: 0,
            apt: data.apt,
            id: 0,
            dong: value.id,          
            category: item.category,
            method: data.methods[j].id,
            count: item.count == 0 ? '' : item.count,
            standard: standard.id,
            lastdate: item.lastdate,
            duedate: item.duedate,
            remark: item.remark,
            elevator: i,
            cycle: category.cycle,
            percent: category.percent,
            unit: standard.unit,
            price: 0,
            totalprice: 0
          }
          
          let rate = getRate(n)
          n.rate = rate

          if (rate == 0) {
            rate = 100
          }
          
          let price = util.calculatePriceRate(standard.direct, standard.labor, standard.cost, rate, data.parcelrate)
          let totalprice = util.money(util.calculateRepair(standard.direct, standard.labor, standard.cost, rate, data.parcelrate, item.count, category.percent))

          let once = 0

          if (item.method == -1) {
            once = j
          }

          n.price = price
          n.totalprice = totalprice
          n.once = once

          items.push(n)
        }
      }
    })
  })

  data.batchs = data.batchs.concat(items)
  for (let i = 0; i < data.batchs.length; i++) {
    data.batchs[i].index = i
  }

  data.save = util.clone({
    item: data.item,
    standards: data.standards,
    methods: data.methods
  })

  readHistory(item.category)
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
  if (listSelection.value.length == 0) {
    util.error('선택된 항목이 없습니다')
    return
  }

  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)

    let items = []
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id
      }

      items.push(item)
    }

    await Breakdown.removebatch(items)

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

function getTitle() {
  if (data.batchs.length == 0) {
    return ''
  }

  let item = data.batchs[0] 
  let category = getCategory(item.category)
  let subcategory = getCategory(category.parent)
  let topcategory = getCategory(subcategory.parent)
  
  return `${topcategory.name} > ${subcategory.name} > ${category.name}`
}

function onKeyup(index) {
  let item = data.batchs[index]
  let count = util.getInt(item.count)

  let rate = item.rate
  if (rate == 0) {
    rate = 100
  }  

  let totalprice = util.money(item.price * item.percent * count / 100 * rate / 100.0 * data.parcelrate / 100.0)
  data.batchs[index].totalprice = totalprice

  if (item.once == 0) {
    return
  } else if (item.once == 1) {
    data.batchs[index+1].count = item.count

    item = data.batchs[index+1]
    let rate = item.rate
    if (rate == 0) {
      rate = 100
    }
    totalprice = util.money(item.price * item.percent * count / 100 * rate / 100.0 * data.parcelrate / 100.0)
    data.batchs[index+1].totalprice = totalprice
  } else {
    data.batchs[index-1].count = item.count

    item = data.batchs[index-1]
    let rate = item.rate
    if (rate == 0) {
      rate = 100
    }
    totalprice = util.money(item.price * item.percent * count / 100 * rate / 100.0 * data.parcelrate / 100.0)
    data.batchs[index-1].totalprice = totalprice
  }
}

async function clickBatchSubmit() {
  util.loading(true)

  let items = util.clone(data.batchs)
  for (let i = 0; i < items.length; i++) {
    let item = items[i]
    let category = getCategory(item.category)
    let subcategory = getCategory(category.parent)
    let topcategory = getCategory(subcategory.parent)

    items[i].subcategory = subcategory.id
    items[i].topcategory = topcategory.id
    
    items[i].count = util.getInt(item.count)
    items[i].lastdate = util.getInt(item.lastdate)
    items[i].duedate = util.getInt(item.duedate)

    if (items[i].duedate >= 2100) {      
      util.error('수선예정년도를 정확하게 입력하세요')
      util.loading(false)
      return
    }
  }

  await Breakdown.insertbatch(items)

  util.info('등록되었습니다')
  getItems()
  data.visibleBatch = false
  
  util.loading(false)
  setTimeout(function() {
    listRef.value!.setScrollTop(data.items.length  * 100 + 1000)
  }, 500)
}

function clickBatchRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function changeMethod(item, updateFlag) {
  let method = getCategory(item.method)

  let standard = null
  
  for (let i = 0; i < data.allstandards.length; i++) {
    let standardItem = data.allstandards[i]

    if (standardItem.id == item.standard) {
      standard = standardItem
    }
  }

  let rate = getRate(item)
  data.item.rate = rate
  if (rate == 0) {
    rate = 100
  }

  if (standard == null) {
    data.item.unit = ''
  } else {
    data.item.unit = standard.unit
  }
  
  if (method == null || method.id == 0) {
    data.item.cycle = ''
    data.item.percent = ''

    data.item.price = ''
    data.item.totalprice = ''
  } else {
    data.item.cycle = method.cycle
    data.item.percent = method.percent

    let cycle = util.getInt(data.item.cycle)
    if (cycle > 0) {
      let lastdate = util.getInt(data.item.lastdate)
      if (lastdate > 0) {
        lastdate += data.item.cycle
      }

      let d = new Date()
      let year = d.getFullYear();
      while (lastdate < year) {
        lastdate += data.item.cycle
      }

      if (updateFlag != false) {
        data.item.duedate = lastdate
      }
    }

    for (let i = 0; i < data.allstandards.length; i++) {
      let standard = data.allstandards[i]

      if (standard.id == item.standard) {
        data.item.price = util.money(util.calculatePriceRate(standard.direct, standard.labor, standard.cost, rate, data.parcelrate))
        
        if (util.getInt(data.item.count) > 0) {
          data.item.totalprice = util.money(util.calculateRepair(standard.direct, standard.labor, standard.cost, rate, data.parcelrate, data.item.count, method.percent))
        } else {
          data.item.totalprice = ''
        }
        
        return        
      }
    }

    data.item.price = ''
    data.item.totalprice = ''    
  }
}

function changeStandard(item) {
  changeMethod(item)
}

function onKeyupSingle(multi) {
  console.log('onKeyupSingle')
  let item = data.item

  let rate = getRate(item)
  if (rate == 0) {
    rate = 100
  }

  let standard = null

  if (multi == true) {
    if (multipleSelectionStandard.value.length == 1) {
      for (let i = 0; i < data.allstandards.length; i++) {
        let standardItem = data.allstandards[i]

        if (standardItem.id == multipleSelectionStandard.value[0].value) {
          standard = standardItem
          break
        }
      }
    }
  } else {
    if (util.isNull(item.standard) || item.standard == 0) {
      data.item.totalprice = 0
      return
    }

    for (let i = 0; i < data.allstandards.length; i++) {
      let standardItem = data.allstandards[i]

      if (standardItem.id == item.standard) {
        standard = standardItem
        break
      }
    }
  }

  if (standard == null) {
    data.item.totalprice = 0
    return

  }

  let percent = util.getFloat(item.percent)
  data.item.totalprice = util.money(util.calculateRepair(standard.direct, standard.labor, standard.cost, rate, data.parcelrate, item.count, percent))
}

function changeDongForEvelator() {  
  if (util.getInt(data.item.dong) == -1) {    
    const item = {
      id: 0,
      name: '',
      ground: 0,
      underground: 0,
      parking: 0,
      elevator: 0,
      familycount: 0,
      remark: '',
      basic: 1
    }

    data.dong = item
    
    data.visibleDong = true

    return
  }
  
  for (let i = 0; i < data.dongs.length; i++) {
    let item = data.dongs[i]

    if (item.id == data.item.dong) {      
      let items = [{id:0, name:'승강기'}]
      if (item.elevator >= 0) {
        for (let j = 1; j <= item.elevator; j++) {
          items.push({
            id: j,
            name: `${j}호기`
          })
        }
      }

      data.elevators = items
      
      return
    }
  }  
}

const handleCloseStandard = async (done: () => void) => {
  if (standardPopupRef.value?.getSelectMode() != true) {
    let res = await Standard.findByApt(data.apt)
    data.allstandards = res.items

    let standard = data.item.standard
    let method = data.item.method
    changeCategory({value: data.item.category})
    clickCategory({value: data.item.category})

    data.item.standard = standard
    data.item.method = method
  }

  standardPopupRef.value?.setSelectMode(null)
  done()
}

const handleCloseCategory = async (done: () => void) => {
  data.categorys = await util.getCategoryTree(data.apt, '공사종별')
  
  done()
}

async function readHistory(category) {
  let res = await History.find({
    apt: data.apt,
    orderby: 'h_year,h_month,c_order,h_id',
    category: category
  })

  if (res.items == null) {
    res.items = []
  }

  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1
  }

  data.historys = res.items
}

function clickShowHistory(mode) {
  util.loading(true)

  readHistory(data.item.category)

  util.loading(false)
  data.visibleHistory = true
}

const getHistorySummaries = (param: SummaryMethodProps) => {
  const { columns, data } = param
  const sums: string[] = []
  columns.forEach((column, index) => {
    if (index === 1) {
      sums[index] = '사용 계'
    } else if (index == 6) {
      let total = 0
      if (data != null) {
        data.forEach((item) => {
          total += item.price
        })
      }
      
      sums[index] = util.money(total)    
    }
  })

  return sums
}

function clickTopcategory(row) {
  data.batchsubcategorys = row.children
  data.batchcategorys = []
  data.batchstandards = []
}

function clickSubcategory(row) {
  data.batchcategorys = row.children
  data.batchstandards = []
}

function clickCategory(row) {
  data.item.category = row.value
  
  let standards = []
  
  for (let i = 0; i < data.allstandards.length; i++) {
    let item = data.allstandards[i]

    if (item.category == row.value) {
      standards.push({
        label: item.name,
        value: item.id
      })
    }
  }

  data.batchstandards = standards

  let category = getCategory(row.value)
  
  let methods = [{id: 0, name: '수선방법'}]

  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.parent == category.id) {
      methods.push(item)
    }
  }

  if (methods.length == 3) {
    methods.push({
      id: -1,
      name: '전면 + 부분'
    })
  }

  data.methods = methods

  
  /*
     if (standards.length == 2) {
     flag = true
     data.item.standard = standards[1].id
     } else {
     data.item.standard = null
     }
     
     changeMethod(data.item)
   */
}

const standardRef = ref<InstanceType<typeof ElTable>>()
const multipleSelectionStandard = ref([])

const changeStandardData = (val) => {
  multipleSelectionStandard.value = val
}

function clickStandard(item) {
  standardRef.value!.toggleRowSelection(item, undefined)

  changeMethodData(item)
}

function changeMethodData(item) {
  let method = getCategory(data.item.method)

  let count = 0
  multipleSelectionStandard.value.forEach((value) => {
  })

  let standard = null
  
  if (multipleSelectionStandard.value.length == 1) {
    for (let i = 0; i < data.allstandards.length; i++) {
      let standardItem = data.allstandards[i]

      if (standardItem.id == multipleSelectionStandard.value[0].value) {
        standard = standardItem
        break
      }
    }
  }

  if (standard == null) {
    data.item.unit = ''
  } else {
    data.item.unit = standard.unit
  }
  
  if (method == null || method.id == 0) {
    data.item.cycle = ''
    data.item.percent = ''

    data.item.price = ''
    data.item.totalprice = ''
  } else {
    data.item.cycle = method.cycle
    data.item.percent = method.percent

    let cycle = util.getInt(data.item.cycle)
    if (cycle > 0) {
      let lastdate = util.getInt(data.item.lastdate)
      if (lastdate > 0) {
        lastdate += data.item.cycle
      }

      let d = new Date()
      let year = d.getFullYear();
      while (lastdate < year) {
        lastdate += data.item.cycle
      }

      //if (util.getInt(data.item.duedate) == 0) {
      data.item.duedate = lastdate
      //}
    }

    if (standard != null) {
      data.item.price = util.money(util.calculatePriceRate(standard.direct, standard.labor, standard.cost, data.item.rate, data.parcelrate))
      
      if (util.getInt(data.item.count) > 0) {
        data.item.totalprice = util.money(util.calculateRepair(standard.direct, standard.labor, standard.cost, data.item.rate, data.parcelrate, data.item.count, method.percent))
      } else {
        data.item.totalprice = ''
      }
    } else {
      data.item.price = ''
      data.item.totalprice = ''
    }
  }
}


const standardPopupRef = ref()

function clickShowStandard() {
  //standardPopupRef.value?.readItems(data.item.category)
  standardPopupRef.value?.setCategory(data.item.category)
  data.visibleStandard = true
}

async function selectPopupStandard(standard) {
  let res = await Standard.findByApt(data.apt)
  data.allstandards = res.items

  const index = data.index
  let item = util.clone(data.batchs[index])

  let rate = item.rate
  if (item.rate == 0) {
    rate = 100
  }

  item.standard = standard.id
  let price = util.money(util.calculatePriceRate(standard.direct, standard.labor, standard.cost, rate, data.parcelrate))
  item.price = price

  data.batchs[index] = item

  onKeyup(index)

  data.visibleStandard = false
}

function changeBatchInsert(row, index) {
  if (index.no == 3) {
    data.index = row.index
    standardPopupRef.value?.setSelectMode(selectPopupStandard)
    clickShowStandard()
  }
  
  readHistory(row.category)
}

function onKeyupLastdateSingle() {
  let cycle = util.getInt(data.item.cycle)
  if (cycle > 0) {
    let lastdate = util.getInt(data.item.lastdate)
    if (lastdate > 0) {
      lastdate += data.item.cycle
    }

    let d = new Date()
    let year = d.getFullYear();
    while (lastdate < year) {
      lastdate += data.item.cycle
    }

    data.item.duedate = lastdate
  }
}

function onKeyupLastdate(index) {
  let item = data.batchs[index]

  let cycle = util.getInt(item.cycle)
  if (cycle > 0) {
    let lastdate = util.getInt(item.lastdate)
    if (lastdate > 0) {
      lastdate += item.cycle
    }

    let d = new Date()
    let year = d.getFullYear();
    while (lastdate < year) {
      lastdate += item.cycle
    }

    data.batchs[index].duedate = lastdate
  }
}

function clickDeduplication() {
  util.confirm('중복된 데이터를 삭제하시겠습니까', async function() {
    util.loading(true)

    const item = {
      id: data.apt
    }
    
    //await Breakdown.deduplication(item)

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

function clickChangeYear() {
  if (listSelection.value.length == 0) {
    util.error('선택된 항목이 없습니다')
    return
  }

  let items = []
  for (let i = 0; i < listSelection.value.length; i++) {
    let item = listSelection.value[i]

    let standard = item.extra.standard
    let category = item.extra.category
    let rate = item.rate

    if (item.rate == 0) {
      rate = 100
    }

    let price = util.money(util.calculatePriceRate(standard.direct, standard.labor, standard.cost, rate, data.parcelrate))
    let totalprice = util.money(util.calculateRepair(standard.direct, standard.labor, standard.cost, rate, data.parcelrate, item.count, category.percent))

    item.price = price
    item.totalprice = totalprice

    item.index = i
    items.push(item)
  }

  let n = util.clone(item)
  n.count = ''
  items.push(n)

  data.lastdate = ''
  data.duedate = ''

  data.batchs = items

  data.visibleYear = true
}

async function clickSubmitYear() {
  let duedate = util.getInt(data.duedate)
  let lastdate = util.getInt(data.lastdate)

  if (duedate == 0 && lastdate == 0) {
    util.error('수선예정년도 혹은 최종수선년도를 입력하세요')
    return
  }

  let d = new Date()
  let year = d.getFullYear()
  
  if (duedate > 0) {
    if (duedate > year + util.getPlanyears(data.repair.planyears)) {
      util.error('수선예정년도를 정확하게 입력하세요')
      return
    }
  }

  if (lastdate > 0) {    
    if (lastdate > year + util.getPlanyears(data.repair.planyears)) {
      util.error('최종수선년도를 정확하게 입력하세요')
      return
    }
  }

  util.loading(true)

  let ids = []
  for (let i = 0; i < data.batchs.length - 1; i++) {
    let item = data.batchs[i]
    ids.push(item.id)

    for (let j = 0; j < data.items.length; j++) {
      if (data.items[j].id == item.id) {
        if (duedate > 0) {
          data.items[j].duedate = duedate
        }

        if (lastdate > 0) {
          data.items[j].lastdate = lastdate
        }
      }
    }
  }

  let id = ids.join(',')

  if (lastdate > 0) {
    await Breakdown.updateLastdate(lastdate, id)
  }

  if (duedate > 0) {
    await Breakdown.updateDuedate(data.apt, duedate, id)
  }

  listRef.value!.clearSelection()
  
  util.info('수정되었습니다')
  util.loading(false)
  data.visibleYear = false
}

const spanMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  if (rowIndex == data.batchs.length - 1) {
    if (columnIndex == 0) {
      return {rowspan: 1, colspan: 12}
    } else if (columnIndex == 12) {
      return {rowspan: 1, colspan: 1}      
    } else if (columnIndex == 13) {
      return {rowspan: 1, colspan: 1}
    } else {
      return {rowspan: 0, colspan: 0}
    }
  }

  return {rowspan: 1, colspan: 1}
}

function clickAddDong() {
  data.dong = util.clone(dong)
  data.visibleAddDong = true
}

function clickAddDongCancel() {
  data.visibleAddDong = false
}

async function clickAddDongSubmit() {
  let item = data.dong
  if (item.name == '') {
    util.error('시설물명을 입력하세요')
    return    
  }

  
  util.loading(true)

  item.apt = data.apt
  item.familycount = util.getInt(item.familycount)
  item.size = util.getFloat(item.size)    
  item.order = util.getInt(item.order)

  item.basic = util.getInt(item.basic)
  if (item.basic == 0) {
    item.basic = 1
  }

  await Dong.insert(item)

  util.info('등록되었습니다')

  let res = await Dong.findByApt(data.apt)
  if (res.items == null) {
    res.items = []
  }

  data.dongs = [{id: 0, name: '시설물'}, ...res.items, {id: -1, name: '시설물 등록'}]
  data.batchdongs = res.items
  
  data.visibleAddDong = false  
  util.loading(false)  
}

async function clickSubmitDong() {
  let item = data.dong

  if (item.name == '') {
    util.error('시설물명을 입력하세요')
    return    
  }  
  
  util.loading(true)

  
  item.apt = data.apt
  item.familycount = util.getInt(item.familycount)
  item.size = util.getFloat(item.size)    
  item.order = util.getInt(item.order)

  item.basic = util.getInt(item.basic)
  if (item.basic == 0) {
    item.basic = 1
  }

  let res = await Dong.insert(item)
  let id = res.id
  
  util.info('등록되었습니다')

  res = await Dong.findByApt(data.apt)
  if (res.items == null) {
    res.items = []
  }

  data.dongs = [{id: 0, name: '시설물'}, ...res.items, {id: -1, name: '시설물 등록'}]
  data.batchdongs = res.items

  data.item.dong = id

  data.visibleDong = false
  util.loading(false)
}

const handleClosePrice = (done: () => void) => {
  done()
}

async function clickSubmitPrice() {
  util.loading(true)
  
  let item = util.clone(data.standard)

  item.direct = util.getInt(item.direct)
  item.labor = util.getInt(item.labor)
  item.cost = util.getInt(item.cost)
  
  await Standard.update(item)

  for (let i = 0; i < data.items.length; i++) {
    if (data.items[i].standard != item.id) {
      continue
    }

    data.items[i].extra.standard.direct = item.direct
    data.items[i].extra.standard.labor = item.labor
    data.items[i].extra.standard.cost = item.cost
  }
  
  util.info('변경되었습니다')
  util.loading(false)

  data.visiblePrice = false
}

function onKeyupStandard() {
  let item = data.standard
  let direct = util.getInt(item.direct)
  let labor = util.getInt(item.labor)
  let cost = util.getInt(item.cost)

  let price = util.money(util.calculatePrice(direct, labor, cost))
  data.standard.price = price
}

const handleClosePeriod = (done: () => void) => {
  done()
}

async function clickSubmitPeriod() {  
  util.loading(true)
  
  let item = util.clone(data.method)
  item.cycle = util.getInt(item.cycle)
  
  await Category.update(item)

  for (let i = 0; i < data.items.length; i++) {
    if (data.items[i].method != item.id) {
      continue
    }

    data.items[i].extra.category.cycle = item.cycle    
  }
  
  util.info('변경되었습니다')
  util.loading(false)

  data.visiblePeriod = false
}

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
}

async function handelSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  util.loading(true)

  let params = {
    id: data.apt,
    filename: response.filename
  }

  data.filename = response.filename

  console.log('params', params)
  let res = await Upload.diff(data.apt, params)
  
  util.loading(false)

  console.log(res)


  for (let i = 0; i < res.newBreakdown.length; i++) {
    res.newBreakdown[i].extra = {}
    res.newBreakdown[i].extra.standard = getStandardInfo(res.newBreakdown[i].standard)
    res.newBreakdown[i].extra.category = getCategory(res.newBreakdown[i].method)
  }
  
  diff.newStandards = res.newStandard
  diff.changeStandards = res.changeStandard
  diff.newBreakdowns = res.newBreakdown
  diff.remainBreakdowns = res.remainBreakdown
  
  diff.visible = true
}

const submitUpload = () => {
  upload.value.clearFiles()
  upload.value!.submit()
}

function getCategoryInfo(id) {
  for (let i = 0; i < data.allcategorys.length; i++) {
    let item = data.allcategorys[i]

    if (item.id == id) {
      return item;
    }
  }

  return {
    id: 0,
    name: ''
  }
}

function getTopcategory(id) {
  let category = getCategoryInfo(id)

  let subcategory = getCategoryInfo(category.parent)
  let topcategory = getCategoryInfo(subcategory.parent)


  return topcategory.name
}

function getSubcategory(id) {
  let category = getCategoryInfo(id)
  let subcategory = getCategoryInfo(category.parent)

  return subcategory.name  
}

function getCategory2(id) {
  let category = getCategoryInfo(id)

  return category.name  
}

const newStandardRef = ref<InstanceType<typeof ElTable>>()
const newStandardSelection = ref([])
const changeNewStandard = (val) => {
  newStandardSelection.value = val
}

const changeStandardRef = ref<InstanceType<typeof ElTable>>()
const changeStandardSelection = ref([])
const changeChangeStandard = (val) => {
  changeStandardSelection.value = val
}

const newBreakdownRef = ref<InstanceType<typeof ElTable>>()
const newBreakdownSelection = ref([])
const changeNewBreakdown = (val) => {
  newBreakdownSelection.value = val
}

const remainBreakdownRef = ref<InstanceType<typeof ElTable>>()
const remainBreakdownSelection = ref([])
const changeRemainBreakdown = (val) => {
  remainBreakdownSelection.value = val
}

async function clickSubmitDiff() {
  util.loading(true)

  let params = {
    id: data.apt,
    newStandard: util.clone(newStandardSelection.value),
    changeStandard: util.clone(changeStandardSelection.value),
    newBreakdown: util.clone(newBreakdownSelection.value),
    filename: util.clone(data.filename)
  }
        
  let res = await Upload.diffupdate(params)
  
  util.loading(false)
  
  diff.visible = false
  await getItems(true)
  util.loading(false)
}
</script>
<style>
.el-table__body tr.current-row> td{
  background-color: #409eff! important;
  color: #fff;
}

.overflow {
  overflow:hidden;
  text-overflow:ellipsis;
  white-space:nowrap;
}
</style>

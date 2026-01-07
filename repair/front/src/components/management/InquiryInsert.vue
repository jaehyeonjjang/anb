<template>
  <link href="https://cdn.jsdelivr.net/npm/remixicon@2.2.0/fonts/remixicon.css" rel="stylesheet">
  <el-dialog
    v-model="data.visibleApt"
    title="주소 등록"
    width="600px"    
  >


        <y-table style="margin-top:-20px;">          
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
              <el-input v-model="data.item.floor" style="width:40px;" /> 층
              (지하 <el-input v-model.number="data.item.undergroundfloor" style="width:40px;" /> 층, 지상 <el-input v-model.number="data.item.groundfloor" style="width:40px;" /> 층)
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>연면적</y-th>
            <y-td>
              <el-input v-model="data.item.area" />              
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
            <y-th>담당자 메일주소</y-th>
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
      <el-button v-if="data.id > 0" style="float:left;" size="small" type="danger" @click="clickDelete">삭제</el-button>
      
      <el-button size="small" @click="data.visibleApt = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitInsert">등록</el-button>
    </template>
</el-dialog>


  <el-dialog
    v-model="data.visible"
    title="상담"
    :fullscreen="true"
    :before-close="handleClose"
  >

    <div style="margin-top:-20px;display:flex;gap:10px;" :style="{height: height(150)}">
      <div style="width:400px;">
        <y-table>
          <y-tr>
            <y-th>ID</y-th>
            <y-td>
              <span v-if="data.id != 0">{{ data.item.id }}</span>
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
              <el-input v-model="data.item.floor" style="width:40px;" /> 층
              (지하 <el-input v-model.number="data.item.undergroundfloor" style="width:40px;" /> 층, 지상 <el-input v-model.number="data.item.groundfloor" style="width:40px;" /> 층)
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>연면적</y-th>
            <y-td>
              <el-input v-model="data.item.area" />              
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

        <div style="text-align:right;margin-top:10px;">
          <el-button size="small" type="primary" @click="clickSubmit">저장</el-button>
        </div>
      </div>
      <div style="flex:1;">

        <y-table>
          <y-tr>
            <y-th>구분</y-th>
            <y-td>
              <el-checkbox size="small" label="장기수선계획" v-model="data.check1" style="font-size:12px;" />
              <el-checkbox size="small" label="정밀" v-model="data.check2" style="font-size:12px;" />
              <el-checkbox size="small" label="정기" v-model="data.check3" style="font-size:12px;" />
              <el-checkbox size="small" label="하자보수" v-model="data.check4" style="font-size:12px;" />
              <el-checkbox size="small" label="하자조사" v-model="data.check5" style="font-size:12px;" />
              <el-checkbox size="small" label="구조안전진단" v-model="data.check6" style="font-size:12px;" />
              <el-checkbox size="small" label="감리" v-model="data.check7" style="font-size:12px;" />
              <el-checkbox size="small" label="기술자문" v-model="data.check8" style="font-size:12px;" />
              <el-checkbox size="small" label="순찰" v-model="data.check9" style="font-size:12px;" />
              <el-checkbox size="small" label="점검프로그램" v-model="data.check10" style="font-size:12px;" />
              
              
            </y-td>
          </y-tr>
          <y-tr>
            <y-th>상담내용</y-th>
            <y-td>              
              <TiptapEditor v-model="data.inquiry.content" />
            </y-td>
          </y-tr>
          <y-tr>
            
            <y-th>상담결과</y-th>
            <y-td>
              <el-radio-group v-model.number="data.inquiry.status">
                <el-radio-button size="small" label="1">완료</el-radio-button>
                <el-radio-button size="small" label="2">진행</el-radio-button>
                <el-radio-button size="small" label="3">예약</el-radio-button>                
              </el-radio-group>
            </y-td>              
          </y-tr>
          
         </y-table>

         <div style="margin-top:10px;">
           <el-button style="display:block;float:left;" size="small" type="success" @click="clickEstimateInsert">견적 등록</el-button>
           <el-button style="display:block;float:left;" size="small" type="success" @click="clickContractInsert(null)">계약 등록</el-button>
           <el-button style="display:block;float:right;" size="small" type="success" @click="clickSubmitInquiry">상담내용 등록</el-button>
           <div style="clear:both;"></div>
         </div>


         <el-tabs v-model="data.menu" @tab-click="clickTab">
           <el-tab-pane label="상담 내역" name="inquiry">

             
             <el-table :data="data.items" border :height="height(540)"  @row-click="clickInquiry" v-infinite="getItems">
               <el-table-column prop="index" label="NO" align="center" width="60" />
               <el-table-column prop="floor" label="구분" align="center" width="90">
                 <template #default="scope">
                   <div v-for="item in getType(scope.row.type)">{{item}}</div>
                 </template>
               </el-table-column>
               <el-table-column label="내용" align="center">
                 <template #default="scope">
                   <div style="text-align:left;line-height:80%;" v-html="scope.row.content.replace('  ', '&nbsp;&nbsp;')"></div>
                 </template>
               </el-table-column>
               <el-table-column label="상태" align="center" width="50">
                 <template #default="scope">
                   <span v-if="scope.row.status == 1">완료</span>
                   <span v-if="scope.row.status == 2">진행</span>
                   <span v-if="scope.row.status == 3">예약</span>
                 </template>
               </el-table-column>
               <el-table-column prop="extra.user.name" label="상담자" align="center" width="60" />
               <el-table-column prop="date" label="등록일" align="center" width="150" />
             </el-table>  

             
           </el-tab-pane>

           <el-tab-pane label="견적 내역" name="estimate">
             <el-table :data="data.estimates" border :height="height(540)" @row-click="clickEstimate">
               <el-table-column prop="index" label="NO" align="center" width="30" />
               <el-table-column prop="type" label="구분" align="center" width="150">
                 <template #default="scope">
                   <div v-if="scope.row.type == 1">장기수선 (<span v-if="scope.row.subtype == 1">조정</span><span v-if="scope.row.subtype == 2">재수립</span>)</div>
                   <div v-if="scope.row.type == 2">정밀점검</div>
                   <div v-if="scope.row.type == 3">정기점검 (<span v-if="scope.row.subtype == 1">상반기</span><span v-if="scope.row.subtype == 2">하반기</span><span v-if="scope.row.subtype == 3">연간</span><span v-if="scope.row.subtype == 4">정기 5회</span>)</div>
                   <div v-if="scope.row.type == 4">하자보수</div>
                   <div v-if="scope.row.type == 5">하자조사</div>
                   <div v-if="scope.row.type == 6">구조안전진단</div>
                   <div v-if="scope.row.type == 7">감리</div>
                   <div v-if="scope.row.type == 8">기술자문</div>
                   <div v-if="scope.row.type == 9">순찰</div>
                   <div v-if="scope.row.type == 10">점검프로그램 (<span v-if="scope.row.subtype == 1">상반기</span><span v-if="scope.row.subtype == 2">하반기</span><span v-if="scope.row.subtype == 3">연간</span><span v-if="scope.row.subtype == 4">연간-1회무상</span>)</div>
                 </template>
               </el-table-column>               
               <el-table-column label="견적일" align="center" width="110">
                 <template #default="scope">
                   {{util.viewDate(scope.row.writedate)}}
                 </template>
               </el-table-column>               
               <el-table-column label="금액" align="right" width="90">
                 <template #default="scope">
                   <span v-if="scope.row.price > 0">{{util.money(scope.row.price)}} 원</span>
                 </template>
               </el-table-column>
               <el-table-column label="총액" align="right" width="90">
                 <template #default="scope">
                   <span v-if="scope.row.price > 0">
                     <span v-if="scope.row.type == 3 && scope.row.subtype == 3">
                     {{util.money(scope.row.price * 2)}} 원
                     </span>
                     <span v-else>
                     {{util.money(scope.row.price)}} 원
                     </span>

                   </span>
                 </template>
               </el-table-column>
               <el-table-column prop="remark" label="비고" align="left" />
               <el-table-column prop="extra.user.name" label="상담자" align="center" width="60" />
               <el-table-column prop="date" label="등록일" align="center" width="130" />
               <el-table-column align="left" width="260">
                 <template #default="scope">
                   <el-button size="small" type="success" @click="clickCopyEstimate(scope.row)">복사</el-button>
                   <el-button size="small" type="danger" @click="clickInsertContract(scope.row)" style="margin-left:5px;">계약</el-button>
                   <el-button size="small" type="primary" @click="clickDownloadEstimate(scope.row, 0, '', 1)" style="margin-left:5px;">다운</el-button>                   
                   <el-button v-for="compareestimate in scope.row.extra.compareestimate" size="small" type="warning" @click="clickDownloadEstimate(scope.row, compareestimate.comparecompany, compareestimate.extra.comparecompany.name, 1)" style="margin-left:5px;">{{compareestimate.extra.comparecompany.name.substring(0, 2)}}</el-button>                   
                 </template>
               </el-table-column>
             </el-table>
           </el-tab-pane>
           
           <el-tab-pane label="계약 내역" name="contract">
             <el-table :data="data.contracts" border :height="height(540)" @row-click="clickContract">
               <el-table-column prop="index" label="NO" align="center" width="40" />
               <el-table-column prop="type" label="구분" align="center" width="100">
                 <template #default="scope">
                   <div v-for="item in getType(scope.row.type)">{{item}}</div>                   
                 </template>
               </el-table-column>               
               <el-table-column label="계약일" align="center" width="110">
                 <template #default="scope">
                   {{util.viewDate(scope.row.contractdate)}}
                 </template>
               </el-table-column>
               <el-table-column label="계약기간" align="center" width="220">
                 <template #default="scope">
                   {{util.viewDate(scope.row.contractstartdate)}} ~ {{util.viewDate(scope.row.contractenddate)}}
                 </template>
               </el-table-column>
               <el-table-column label="금액" align="right" width="70">
                 <template #default="scope">
                   <span v-if="scope.row.price > 0">{{util.money(scope.row.price)}} 만원</span>
                 </template>
               </el-table-column>
               <el-table-column label="VAT" align="center" width="45">
                 <template #default="scope">
                   <span v-if="scope.row.vat == 1">포함</span>
                   <span v-if="scope.row.vat == 2">별도</span>
                 </template>
               </el-table-column>
               <el-table-column label="세금계산서발행" align="center" width="110">
                 <template #default="scope">
                   {{util.viewDate(scope.row.invoice)}}
                 </template>
               </el-table-column>
               <el-table-column label="입금일" align="center" width="110">
                 <template #default="scope">
                   {{util.viewDate(scope.row.depositdate)}}
                 </template>
               </el-table-column>
               <el-table-column prop="remark" label="비고" align="left" />
               <el-table-column prop="extra.user.name" label="상담자" align="center" width="60" />
               <el-table-column prop="date" label="등록일" align="center" width="130" />
               <el-table-column align="center" width="120">
                 <template #default="scope">
                   <el-button v-if="scope.row.estimate > 0" size="small" type="primary" @click="clickDownloadEstimateByContract(scope.row, 0, '', 1)" style="margin-left:5px;">다운</el-button>                   
                   <el-button style="margin-left:5px;" v-if="scope.row.estimate > 0" size="small" type="success" @click="clickEstimateByContract(scope.row.estimate)">견적</el-button>
                 </template>
               </el-table-column>
             </el-table>
           </el-tab-pane>
           
           <el-tab-pane label="장기수선계획" name="repair">

             <el-table :data="data.repairs" border :height="height(540)" @row-click="clickRepair">
               <el-table-column prop="index" label="NO" align="center" width="60" />
               <el-table-column prop="type" label="구분" align="center" width="100">
                 <template #default="scope">
                   <el-tag :type="Repair.getTypeType(scope.row.type)">{{Repair.getType(scope.row.type)}}</el-tag>
                 </template>
               </el-table-column>
               <el-table-column prop="reportdate" label="리포트 작성일" align="center" />
               <el-table-column label="설명">
                 <template #default="scope">
                   {{scope.row.remark}}
                 </template>
               </el-table-column>
               <el-table-column prop="date" label="등록일" align="center" />
               <el-table-column label="상태" align="center" width="60">
                 <template #default="scope">
                   <span v-if="scope.row.status == 2 || scope.row.type == 3">마감</span>
                   <span v-if="scope.row.type != 3 && scope.row.status == 1 && scope.$index < data.items.length - 1">진행</span>
                   <span v-if="scope.row.type != 3 && scope.row.status == 1 && scope.$index == data.items.length - 1" style="color:#af2020;">현재진행</span>            
                 </template>
               </el-table-column>               
               
             </el-table>
           </el-tab-pane>
           <el-tab-pane label="정기점검" name="periodic">

             
             <el-table :data="data.periodics" border :height="height(540)" @row-click="clickPeriodic">
               <el-table-column prop="index" label="NO" align="center" width="60" />
               <el-table-column prop="name" label="작업명" />                       
               <el-table-column prop="reportdate" label="리포트작업일" align="center" />
               <el-table-column label="상태" align="center" width="60">
                 <template #default="scope">
                   <el-tag :type="Periodic.getStatusType(scope.row.status)">{{Periodic.getStatus(scope.row.status)}}</el-tag>
                 </template>
               </el-table-column>
               <el-table-column prop="date" label="등록일" align="center" />                       
               
             </el-table>
           </el-tab-pane>           
         </el-tabs>

      </div>
    </div>

    <template #footer>
      <el-button v-if="data.id > 0" style="float:left;" size="small" type="danger" @click="clickDelete">삭제</el-button>
      
      <el-button size="small" @click="clickClose">닫기</el-button>
    </template>
  </el-dialog>

<el-dialog
    v-model="data.visibleEstimate"
    :title="data.estimate.id > 0 ? '견적 정보 수정' : '견적 정보 등록'"
    width="1220px"
    top="20px"
  >

  <div style="display:flex;justify-content:space-between;gap:10px;align-items:flex-start;">
    <y-table style="width:400px;">
      <y-tr>
        <y-td>
          <el-select v-model.number="data.estimate.type" style="width:110px;" size="small" placeholder="" @change="changeEstimateType">
            <el-option
              v-for="item in data.types"
              :key="item.id"
              :label="item.title"
              :value="item.id"
            />
          </el-select>

          <span style="margin-left:10px;" v-if="data.estimate.type == 7 || data.estimate.type == 8">공사명: </span><el-input v-model="data.estimate.name" v-if="data.estimate.type == 7 || data.estimate.type == 8" style="width: 200px;" />
          <el-radio-group v-model.number="data.estimate.subtype"  v-if="data.estimate.type == 1" style="margin-left:10px;">
            <el-radio-button size="small" label="1" @click="clickSubtype(1)">조정</el-radio-button>
            <el-radio-button size="small" label="2" @click="clickSubtype(2)">재수립</el-radio-button>
          </el-radio-group>

          <el-radio-group v-model.number="data.estimate.subtype"  v-if="data.estimate.type == 3" style="margin-left:10px;">
            <el-radio-button size="small" label="1" @click="clickSubtype(1)">상반기</el-radio-button>
            <el-radio-button size="small" label="2" @click="clickSubtype(2)">하반기</el-radio-button>
            <el-radio-button size="small" label="3" @click="clickSubtype(3)">연간</el-radio-button>
            <!-- <el-radio-button size="small" :value="4" @click="clickSubtype(4)">정기 5회</el-radio-button> -->
          </el-radio-group>

          <el-radio-group v-model.number="data.estimate.subtype"  v-if="data.estimate.type == 10" style="margin-left:10px;">
            <el-radio-button size="small" label="1" @click="clickSubtype(1)">상반기</el-radio-button>
            <el-radio-button size="small" label="2" @click="clickSubtype(2)">하반기</el-radio-button>
            <el-radio-button size="small" label="3" @click="clickSubtype(3)">연간</el-radio-button>
            <el-radio-button size="small" label="4" @click="clickSubtype(4)">연간-1회무상</el-radio-button>
          </el-radio-group>

          <el-select v-model="data.estimate.start" style="margin-left:20px;width:150px;" size="small" placeholder=""  v-if="(data.estimate.type == 3 && data.estimate.subtype == 4) || data.estimate.type == 2">
            <el-option
              v-for="item in data.options"
              :key="item.value"
              :label="item.id"
              :value="item.value"
            />
          </el-select>          
        </y-td>        
      </y-tr>
      <y-tr>
        <y-td style="display:flex;justify-content:space-between;border:none;">
          <div>
            <b>견적일 : </b>
            <el-date-picker style="margin: 0px 0px;height: 24px;width:120px;" v-model="data.estimate.writedate" placeholder="" />
          </div>

          <div style="margin-right:5px;">
            <el-checkbox size="small" label="이벤트" v-model="data.estimate.event" style="font-size:12px;" v-show="data.estimate.type == 1" />
            <el-checkbox size="small" label="택배배송" v-model="data.estimate.parcel" style="font-size:12px;" v-show="data.estimate.type == 1" />
          </div>
        </y-td>
      </y-tr>
      <y-tr>
        <y-td>
          <div :style="{height: height(480), overflow: 'auto'}">
          <y-table v-show="data.estimate.type != 10">
            <y-tr>
              <y-th rowspan="2" style="width:140px;">직접 인건비</y-th>              
              <y-th rowspan="2" style="width:80px;text-align:right;">단가</y-th>
              <y-th colspan="2" style="width:90px;text-align:center;">수량</y-th>              
              <y-th rowspan="2" style="width:70px;text-align:right;padding-right:20px;">계</y-th>
            </y-tr>

            <y-tr>              
              <y-th style="text-align:right;width:45px;">외업</y-th>              
              <y-th style="text-align:right;width:45px;">내업</y-th>
            </y-tr>

            <!--
            <y-tr>
              <y-th>기술사</y-th>
              <y-td style="width:120px;text-align:right;">{{util.money(data.standardwage.person1)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person1" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person6" style="width:70px;" @input="changePrice" v-if="data.estimate.type != 1" />
              </y-td>              
              <y-td style="width:120px;text-align:right;padding-right:20px;">{{util.money(calc.price1)}}</y-td>
            </y-tr>
            -->
            <y-tr>
              <y-th>특급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person2)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person7" style="width:35px;" @input="changePrice" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person2" style="width:35px;" @input="changePrice" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.price2)}}</y-td>              
            </y-tr>
            <y-tr>
              <y-th>고급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person3)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person8" style="width:35px;" @input="changePrice" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person3" style="width:35px;" @input="changePrice" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.price3)}}</y-td>              
            </y-tr>
            <y-tr>              
              <y-th>중급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person4)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person9" style="width:35px;" @input="changePrice" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person4" style="width:35px;" @input="changePrice" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.price4)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>초급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person5)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person10" style="width:35px;" @input="changePrice" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.estimate.person5" style="width:35px;" @input="changePrice" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.price5)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>계</y-th>
              <y-td></y-td>
              <y-td></y-td>
              <y-td></y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.price)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>제경비</y-th>
              <y-td colspan="3">
                직접인건비 ✕ <el-input v-model="data.estimate.financialprice" class="inputNumber" style="width:45px;" @input="changePrice" /> %
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.financialprice)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>기술료</y-th>              
              <y-td colspan="3">
                (직접인건비 + 제경비) ✕ <el-input v-model="data.estimate.techprice" class="inputNumber" style="width:35px;" @input="changePrice" /> %
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.techprice)}}</y-td>
            </y-tr>
            
            <y-tr v-if="data.estimate.type == 1">
              <y-th>직접경비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.directprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
            </y-tr>

            <y-tr v-if="data.estimate.type == 1">
              <y-th>인쇄비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.printprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
            </y-tr>

            
            <y-tr v-if="data.estimate.type == 1">
              <y-th>추가경비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.extraprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
            </y-tr>

            <y-tr v-if="data.estimate.type != 1">
              <y-th>직접경비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.directprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;여비 및<br>&nbsp;&nbsp;&nbsp;현장체재비</y-th>
                            
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.travelprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.days"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:40px;" @input="changeDays" />
                일
                ✕ 
                {{data.estimate.travel}}                
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.travelprice)}}</y-td>
            </y-tr>            
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;차량운행비</y-th>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.carprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <!-- <el-input v-model="data.estimate.car"
                     :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                     :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                     class="inputNumber" style="width:70px;" @input="changePrice" /> -->
                {{data.estimate.days}}                
                일
                ✕ 
                {{data.estimate.car}}                
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.carprice)}}</y-td>
            </y-tr>
            <!--
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;시간당 손료 계산</y-th>
                            
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.lossprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:100px;" @input="changePrice" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.loss"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.lossprice)}}</y-td>
            </y-tr>            
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;주연료</y-th>
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.gasprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:100px;" @input="changePrice" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.gas"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.gasprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;잡품</y-th>              
              <y-td style="text-align:right;padding-right:10px;">                
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.etc"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.etcprice)}}</y-td>
            </y-tr>
            -->
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;위험수당</y-th>              
              <y-td style="text-align:right;padding-right:10px;">                
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.danger"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.dangerprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;기계기구손료</y-th>              
              <y-td style="text-align:right;padding-right:10px;">                
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.machine"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.machineprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;인쇄비</y-th>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.printprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.estimate.print"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePrice" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(calc.printprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type == 6">
              <y-th>구조안정성 검토</y-th>
              <y-td></y-td>              
              <y-td colspan="3" style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.stability"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:100px;" @input="changePrice" />
              </y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type == 6">
              <y-th>내진성능평가</y-th>
              <y-td></y-td>              
              <y-td colspan="3" style="text-align:right;padding-right:10px;">
                <el-input v-model="data.estimate.earthquake"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:100px;" @input="changePrice" />
              </y-td>
            </y-tr>
          </y-table>
          </div>
          <y-table>
            <y-tr v-if="data.estimate.type != 10">
              <y-th style="width:70px;">합계금액</y-th>
              <y-td style=""></y-td>              
              <y-td style="width:90px;text-align:right;padding-right:15px;">{{util.money(calc.totalprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 10">
              <y-th>절삭</y-th>
              <y-td></y-td>              
              <y-td style="text-align:right;padding-right:5px;">
                <el-input v-model="data.estimate.saleprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:90px;" @input="changeSaleprice" />
              </y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type == 10">
              <y-th style="width:70px;">합계금액</y-th>
              <y-td style=""></y-td>              
              <y-td style="width:90px;text-align:right;padding-right:15px;">{{util.money(calc.programprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type == 10">
              <y-th>
                <span v-if="data.estimate.subtype != 4">절삭금액</span>
                <span v-if="data.estimate.subtype == 4">1회 무상제공</span>
              </y-th>
              <y-td></y-td>              
              <y-td style="text-align:right;padding-right:5px;">
                <el-input v-model="data.estimate.saleprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:90px;" @input="changeTotalprice" />
              </y-td>
            </y-tr>
            <y-tr>
              <y-th>견적금액</y-th>
              <y-td></y-td>              
              <y-td style="text-align:right;padding-right:5px;">
                <el-input v-model="data.estimate.price"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:90px;" @input="changeTotalprice" />
              </y-td>
            </y-tr>
          </y-table>
        </y-td>
      </y-tr>            
      <y-tr>
        <y-td><el-input v-model="data.estimate.remark" :rows="3" type="textarea" /></y-td>
      </y-tr>
    </y-table>

    <y-table style="width:400px;border:" v-for="(estimate, index) in data.compareestimates">
      <y-tr>
        <y-td>
          비교견적업체 : <el-select v-model.number="data.compareestimates[index].comparecompany" style="width:130px;" size="small" placeholder="">
            <el-option
              v-for="item in data.comparecompanys"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </y-td>        
      </y-tr>
      <y-tr v-if="estimate.comparecompany > 0">
        <y-td style="display:flex;justify-content:space-between;border:none;">
          <div>
            <b>견적일 : </b>
            <el-date-picker style="margin: 0px 0px;height: 24px;width:120px;" v-model="data.compareestimates[index].writedate" placeholder="" />
          </div>

          <div>
            <b>자동견적 : </b>
            <el-input v-model="data.compareestimates[index].adjust" style="width:70px;" @keyup="autoCompare" />
          </div>
        </y-td>
      </y-tr>
      <y-tr v-if="estimate.comparecompany > 0">
        <y-td>
          <div :style="{height: height(480), overflow: 'auto'}">
          <y-table>
            <y-tr>
              <y-th rowspan="2" style="width:150px;">직접 인건비</y-th>              
              <y-th rowspan="2" style="width:80px;text-align:right;">단가</y-th>
              <y-th colspan="2" style="width:90px;text-align:center;">수량</y-th>              
              <y-th rowspan="2" style="width:70px;text-align:right;padding-right:20px;">계</y-th>
            </y-tr>

            <y-tr>              
              <y-th style="text-align:right;width:45px;">외업</y-th>              
              <y-th style="text-align:right;width:45px;">내업</y-th>
            </y-tr>

            <y-tr>
              <y-th>특급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person2)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person7" style="width:35px;" @input="changePriceCompare(index)" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person2" style="width:35px;" @input="changePriceCompare(index)" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].price2)}}</y-td>              
            </y-tr>
            <y-tr>
              <y-th>고급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person3)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person8" style="width:35px;" @input="changePriceCompare(index)" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person3" style="width:35px;" @input="changePriceCompare(index)" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].price3)}}</y-td>              
            </y-tr>
            <y-tr>              
              <y-th>중급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person4)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person9" style="width:35px;" @input="changePriceCompare(index)" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person4" style="width:35px;" @input="changePriceCompare(index)" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].price4)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>초급기술자</y-th>
              <y-td style="text-align:right;">{{util.money(data.standardwage.person5)}}</y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person10" style="width:35px;" @input="changePriceCompare(index)" v-if="data.estimate.type != 1" />
              </y-td>
              <y-td style="text-align:right;">
                <el-input v-model.number="data.compareestimates[index].person5" style="width:35px;" @input="changePriceCompare(index)" />
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].price5)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>계</y-th>
              <y-td></y-td>
              <y-td></y-td>
              <y-td></y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].price)}}</y-td>
            </y-tr>
            <y-tr v-if="estimate.comparecompany != 3">
              <y-th>제경비</y-th>
              <y-td colspan="3">
                직접인건비 ✕ <el-input v-model="data.compareestimates[index].financialprice" class="inputNumber" style="width:45px;" @input="changePriceCompare(index)" /> %
              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].financialprice)}}</y-td>
            </y-tr>
            <y-tr v-if="estimate.comparecompany != 3">
              <y-th>기술료</y-th>              
              <y-td colspan="3">
                (직접인건비 + 제경비) ✕ <el-input v-model="data.compareestimates[index].techprice" class="inputNumber" style="width:35px;" @input="changePriceCompare(index)" /> %
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].techprice)}}</y-td>
            </y-tr>
            
            <y-tr v-if="estimate.comparecompany == 3">
              <y-th>제경비,기술료</y-th>              
              <y-td colspan="3">
                직접인건비 ✕ <el-input v-model="data.compareestimates[index].financialprice" class="inputNumber" style="width:60px;" @input="changePriceCompare(index)" /> %
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].financialprice)}}</y-td>
            </y-tr>

            <y-tr v-if="data.estimate.type == 1 && estimate.comparecompany != 3">
              <y-th>직접경비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.compareestimates[index].directprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
            </y-tr>

            <y-tr v-if="data.estimate.type == 1">
              <y-th>인쇄비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.compareestimates[index].printprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
            </y-tr>

            
            <y-tr v-if="data.estimate.type == 1 && estimate.comparecompany != 3">
              <y-th>추가경비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.compareestimates[index].extraprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
            </y-tr>

            <y-tr v-if="data.estimate.type != 1">
              <y-th>직접경비</y-th>
              <y-td colspan="3"></y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].directprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;여비 및<br>&nbsp;&nbsp;&nbsp;현장체재비</y-th>
                            
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.compareestimates[index].travelprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                {{data.estimate.days}}                
                일
                ✕ 
                {{data.compareestimates[index].travel}}                
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].travelprice)}}</y-td>
            </y-tr>            
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;차량운행비</y-th>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.compareestimates[index].carprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <!-- <el-input v-model="data.compareestimates[index].car"
                     :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                     :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                     class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" /> -->
                {{data.estimate.days}}                
                일
                ✕ 
                {{data.compareestimates[index].car}}                

              </y-td>              
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].carprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;위험수당</y-th>              
              <y-td style="text-align:right;padding-right:10px;">                
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.compareestimates[index].danger"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].dangerprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;기계기구손료</y-th>              
              <y-td style="text-align:right;padding-right:10px;">                
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.compareestimates[index].machine"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].machineprice)}}</y-td>
            </y-tr>
            <y-tr v-if="data.estimate.type != 1">
              <y-th>&nbsp;&nbsp;&nbsp;인쇄비</y-th>              
              <y-td style="text-align:right;padding-right:10px;">
                <el-input v-model="data.compareestimates[index].printprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
              <y-td colspan="2" style="text-align:right;">
                <el-input v-model="data.compareestimates[index].print"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:70px;" @input="changePriceCompare(index)" />
              </y-td>
              <y-td style="text-align:right;padding-right:10px;">{{util.money(comparecalc[index].printprice)}}</y-td>
            </y-tr>
            
          </y-table>
          </div>
          <y-table>
            <y-tr>
              <y-th style="width:70px;">합계금액</y-th>
              <y-td style=""></y-td>              
              <y-td style="width:90px;text-align:right;padding-right:15px;">{{util.money(comparecalc[index].totalprice)}}</y-td>
            </y-tr>
            <y-tr>
              <y-th>절삭</y-th>
              <y-td></y-td>              
              <y-td style="text-align:right;padding-right:5px;">
                <el-input v-model="data.compareestimates[index].saleprice"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:90px;" @input="changeSalepriceCompare(index)" />
              </y-td>
            </y-tr>
            <y-tr>
              <y-th>견적금액</y-th>
              <y-td></y-td>              
              <y-td style="text-align:right;padding-right:5px;">
                <el-input v-model="data.compareestimates[index].price"
                          :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                          :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                          class="inputNumber" style="width:90px;"  @input="changeTotalpriceCompare(index)" />
              </y-td>
            </y-tr>
          </y-table>
        </y-td>
      </y-tr>            
      <y-tr v-if="estimate.comparecompany > 0">
        <y-td><el-input v-model="data.compareestimates[index].remark" :rows="3" type="textarea" /></y-td>
      </y-tr>
    </y-table>

  </div>

    <template #footer>
      <el-button v-if="data.estimate.id > 0" style="float:left;" size="small" type="danger" @click="clickDeleteEstimate">삭제</el-button>
      
      <el-button size="small" @click="data.visibleEstimate = false">닫기</el-button>
      <el-button size="small" type="primary" @click="clickSubmitEstimate">저장</el-button>
    </template>
</el-dialog>



  <el-dialog
    v-model="data.visibleContract"
    :title="data.contract.id > 0 ? '계약 정보 수정' : '계약 정보 등록'"
    width="900"
  >

    <y-table>
      <y-tr>
        <y-th style="width:100px;">구분</y-th>
        <y-td>
          <el-checkbox size="small" label="장기수선계획" v-model="data.c1" style="font-size:12px;" />
          <el-checkbox size="small" label="정밀" v-model="data.c2" style="font-size:12px;" />
          <el-checkbox size="small" label="정기" v-model="data.c3" style="font-size:12px;" />
          <el-checkbox size="small" label="하자보수" v-model="data.c4" style="font-size:12px;" />
          <el-checkbox size="small" label="하자조사" v-model="data.c5" style="font-size:12px;" />
          <el-checkbox size="small" label="구조안전진단" v-model="data.c6" style="font-size:12px;" />
          <el-checkbox size="small" label="감리" v-model="data.c7" style="font-size:12px;" />
          <el-checkbox size="small" label="기술자문" v-model="data.c8" style="font-size:12px;" />
          <el-checkbox size="small" label="순찰" v-model="data.c9" style="font-size:12px;" />
          <el-checkbox size="small" label="점검프로그램" v-model="data.c10" style="font-size:12px;" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>계약일</y-th>
        <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.contract.contractdate" placeholder="" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>계약기간</y-th>
        <y-td>
          <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.contract.contractstartdate" placeholder="" /> ~
          <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.contract.contractenddate" placeholder="" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>금액</y-th>
        <y-td>
          <el-input v-model="data.contract.price"
                    :formatter="(value) => `${value}`.replace(/\D/g,'').replace(/\B(?=(\d{3})+(?!\d))/g, ',')"
                    :parser="(value) => value.replace(/\D/g,'').replace(/\$\s?|(,*)/g, '')"
                    class="inputNumber" style="width:150px;" /> 만원
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>
          VAT
        </y-th>
        <y-td>
          <el-radio-group v-model.number="data.contract.vat">
                <el-radio-button size="small" label="1">포함</el-radio-button>
                <el-radio-button size="small" label="2">별도</el-radio-button>                
          </el-radio-group>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>세금계산서 발행</y-th>
        <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.contract.invoice" placeholder="" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>입금일</y-th>
        <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.contract.depositdate" placeholder="" /></y-td>
      </y-tr>
      <y-tr>
        <y-th>비고</y-th>
        <y-td><el-input v-model="data.contract.remark" :rows="5" type="textarea" style="font-size:14px;" /></y-td>
      </y-tr>
    </y-table>

    <template #footer>
      <el-button v-if="data.contract.id > 0" style="float:left;" size="small" type="danger" @click="clickDeleteContract">삭제</el-button>
      
      <el-button size="small" @click="data.visibleContract = false">닫기</el-button>
      <el-button size="small" type="primary" @click="clickSubmitContract">저장</el-button>
    </template>
  </el-dialog>


  <el-dialog
    v-model="data.visibleInquiry"
    title="상담내용 수정"
    width="900"
  >

    <y-table>
      <y-tr>
        <y-th>구분</y-th>
        <y-td>
          <el-checkbox size="small" label="장기수선계획" v-model="data.c1" style="font-size:12px;" />
          <el-checkbox size="small" label="정밀" v-model="data.c2" style="font-size:12px;" />
          <el-checkbox size="small" label="정기" v-model="data.c3" style="font-size:12px;" />
          <el-checkbox size="small" label="하자보수" v-model="data.c4" style="font-size:12px;" />
          <el-checkbox size="small" label="하자조사" v-model="data.c5" style="font-size:12px;" />
          <el-checkbox size="small" label="구조안전진단" v-model="data.c6" style="font-size:12px;" />
          <el-checkbox size="small" label="감리" v-model="data.c7" style="font-size:12px;" />
          <el-checkbox size="small" label="기술자문" v-model="data.c8" style="font-size:12px;" />
          <el-checkbox size="small" label="순찰" v-model="data.c9" style="font-size:12px;" />
          <el-checkbox size="small" label="점검프로그램" v-model="data.c10" style="font-size:12px;" />
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>상담내용</y-th>
        <y-td>
          <TiptapEditor v-model="data.editcontent" />          
        </y-td>
      </y-tr>            
      <y-tr>
        <y-th>
          상태
        </y-th>
        <y-td>
          <el-radio-group v-model.number="data.editinquiry.status">
                <el-radio-button size="small" label="1">완료</el-radio-button>
                <el-radio-button size="small" label="2">진행</el-radio-button>
                <el-radio-button size="small" label="3">예약</el-radio-button>                
          </el-radio-group>          
        </y-td>
      </y-tr>      
    </y-table>

    <template #footer>
      <el-button style="float:left;" size="small" type="danger" @click="clickDeleteInquiry">삭제</el-button>
      
      <el-button size="small" @click="data.visibleInquiry = false">닫기</el-button>
      <el-button size="small" type="primary" @click="clickUpdateInquiry">저장</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">

import { reactive, onMounted, computed, watch, watchEffect, ref } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Category, Repair, Facilitycategory, Standardwage, Contract, Estimate, Inquiry, Periodic, Comparecompany, Compareestimate } from "~/models"
import { useStore } from 'vuex'
import axios from 'axios'

const content = ref('')
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
  depositdate: '',
  area: '',
  undergroundfloor: 0,
  groundfloor: 0
}

const estimate = {
  id: 0,
  type: 1,
  subtype: 1,
  originalprice: 0,
  saleprice: 0,
  price: 0,
  discount: 0,
  person1: 0,
  person2: 0,
  person3: 0,
  person4: 0,
  person5: 0,
  person6: 0,
  person7: 0,
  person8: 0,
  person9: 0,
  person10: 0,
  writedate: null,
  financialprice: 90,
  techprice: 20,
  directprice: util.money(100000),
  printprice: util.money(100000),
  extraprice: 0,
  carprice: util.money(20000),
  stability: util.money(2000000),
  earthquake: 0,
  car: 0,
  user: 0,
  apt: 0,
  start: '',
  event: false,
  parcel: false,
  days: 1,
  date: ''
}

const contract = {
  id: 0,
  type: 0,
  contractdate: null,
  contractstartdate: null,
  contractenddate: null,
  price: '',
  vat: 2,
  invoice: null,
  depositdate: null,
  remark: '',  
  apt: 0,
  date: ''
}

const inquiry = {
  id: 0,
  type: 0,
  content: '',
  status: 2,
  user: 0,
  apt: 0,
  date: ''
}

const data = reactive({
  id: 0,  
  item: util.clone(item),
  estimate: util.clone(estimate),
  contract: util.clone(contract),
  inquiry: util.clone(inquiry),
  check1: false,
  check2: false,
  check3: false,
  check4: false,
  check5: false,
  check6: false,
  check7: false,
  check8: false,
  check9: false,
  check10: false,
  c1: false,
  c2: false,
  c3: false,
  c4: false,
  c5: false,
  c6: false,
  c7: false,
  c8: false,
  c9: false,  
  c10: false,  
  visible: false,
  visibleApt: false,
  visibleInquiry: false,
  visibleEstimate: false,
  visibleContract: false,
  page: 1,
  pagesize: 20,
  items: [],
  menu: 'inquiry',
  content: '',
  standardwage: null,
  types: [
    {id: 1, title:'장기수선계획'},
    {id: 2, title:'정밀점검'},
    {id: 3, title:'정기점검'},
    {id: 4, title:'하자보수'},
    {id: 5, title:'하자조사'},
    {id: 6, title:'구조안전진단'},
    {id: 7, title:'감리'},
    {id: 8, title:'기술자문'},
    {id: 10, title:'점검프로그램'},
  ],
  allcomparecompany: [],
  comparecompany: [],
  compareestimate: [],
  originalcompareestimate: [],
})

const calc = reactive({
  price1: 0,
  price2: 0,
  price3: 0,
  price4: 0,
  price5: 0,
  price: 0,
  financialprice: 0,
  techprice: 0,
  totalprice: 0,
  travelprice: 0,
  lossprice: 0,
  gasprice: 0,
  etcprice: 0,
  dangerprice: 0,
  machineprice: 0,
  printprice: 0,
  carprice: 0,
  programprice: 0
})

const comparecalc = reactive([
    {
      price1: 0,
      price2: 0,
      price3: 0,
      price4: 0,
      price5: 0,
      price: 0,
      financialprice: 0,
      techprice: 0,
      totalprice: 0,
      travelprice: 0,
      lossprice: 0,
      gasprice: 0,
      etcprice: 0,
      dangerprice: 0,
      machineprice: 0,
      printprice: 0,
      carprice: 0
    },
    {
      price1: 0,
      price2: 0,
      price3: 0,
      price4: 0,
      price5: 0,
      price: 0,
      financialprice: 0,
      techprice: 0,
      totalprice: 0,
      travelprice: 0,
      lossprice: 0,
      gasprice: 0,
      etcprice: 0,
      dangerprice: 0,
      machineprice: 0,
      printprice: 0,
      carprice: 0
    },
])

async function initData() {
  let res = await Facilitycategory.find({orderby: 'fc_order'})

  data.facilitycategorys = [{id: 0, name: ' '}, ...res.items]

  res = await Standardwage.get(1)
  data.standardwage = res.item

  res = await Comparecompany.find({orderby: 'cc_order,cc_id'})
  data.allcomparecompanys = [{id: 0, name: ' '}, ...res.items]
}

async function getItems(reset) {
  if (data.id != 0) {
    let res = await Apt.get(data.id)
    data.item = res.item
  }

  await readInquiry(reset)
  //await readContract(reset)  
}

async function readInquiry(reset) {
  if (reset == true) {
    data.page = 1
    data.items = []
  }

  let res = await Inquiry.find({apt:data.id, page: data.page, pagesize: data.pagesize, orderby: 'in_id desc'})

  if (res.items == null) {
    res.items = []
  }

  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1    
  }

  data.total = res.total
  data.items = data.items.concat(res.items)  

  data.page++
}

async function readEstimate(reset) {
  if (reset == true) {
    data.page = 1
    data.estimates = []
  }

  let res = await Estimate.find({apt:data.id, orderby: 'e_id desc'})

  if (res.items == null) {
    res.items = []
  }

  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1
  }

  data.total = res.total
  data.estimates = res.items
}

async function readContract(reset) {
  if (reset == true) {
    data.page = 1
    data.contracts = []
  }

  let res = await Contract.find({apt:data.id, orderby: 'co_id desc'})

  if (res.items == null) {
    res.items = []
  }

  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1
  }

  data.total = res.total
  data.contracts = res.items
}

async function clickSubmitInquiry() {
  let item = util.clone(data.inquiry)
  
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

  if (data.check10 == true) {
    contracttype += 512;
  }

  item.type = contracttype
  item.apt = data.id
  item.content = data.inquiry.content
  item.status = util.getInt(data.inquiry.status)
  item.user = store.getters['getUser'].id

  util.loading(true)
  
  await Inquiry.insert(item)
  await readInquiry(true)

  data.inquiry = util.clone(inquiry)

  util.loading(false)
  util.info('등록되었습니다')
}

async function clickSubmitInsert() {
  const item = data.item

  if (item.name === '') {
    util.error('아파트명을 입력하세요')
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

    data.id = res.id
    data.item.id = res.id
    
    data.visibleApt = false
    data.visible = true    
  } else {
    util.error('오류가 발생했습니다')
  }
}

async function clickSubmit() {
  const item = data.item

  if (item.name === '') {
    util.error('아파트명을 입력하세요')
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
    util.info('저장되었습니다')

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

  if (!util.isNull(props.close)) {
    props.close(data.item)    
  }
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
  
  await getItems(true)
  data.menu = 'inquiry'
  data.inquiry = util.clone(inquiry)
  data.visible = true
}

function insert(item) {
  data.id = 0
  if (!util.isNull(item)) {
    data.item = item
  }
  
  data.visibleApt = true
}

function onKeyup() {
  let count = util.getInt(data.item.familycount1) + util.getInt(data.item.familycount2) + util.getInt(data.item.familycount3)

  data.item.familycount = String(count)
}

function getType(type) {
  if (type == 0) {
    return ['기타']
  }

  let strs = []
  
  if (type & 1) strs.push('장기수선계획')
  if (type & 2) strs.push('정밀')
  if (type & 4) strs.push('정기')
  if (type & 8) strs.push('하자보수')
  if (type & 16) strs.push('하자조사')
  if (type & 32) strs.push('정밀안전진단')
  if (type & 64) strs.push('감리')
  if (type & 128) strs.push('기술자문')
  if (type & 256) strs.push('순찰')
  if (type & 512) strs.push('점검프로그램')

  return strs 
}

defineExpose({
  readData,
  insert
})

function clickUpdate() {
}

async function clickRepair(item) {
  return
  
  let res = await Repair.get(item.id)
  let repair = res.item

  if (repair.type == 3) {
    return
  }

  if (repair.status == 2) {
    return
  }

  if (!util.isNull(props.close)) {
    props.close()
  }

  router.push(`/${repair.apt}/repair/${repair.id}/breakdown`)
}

async function clickPeriodic(item) {
  return
  
  let res = await Periodic.get(item.id)
  let periodic = res.item
  
  if (!util.isNull(props.close)) {
    props.close()
  }

  router.push(`/${periodic.apt}/periodic/${periodic.id}/periodic`)
}

function clickContractInsert(estimate, index) {
  data.menu = 'contract'

  data.c1 = false
  data.c2 = false
  data.c3 = false
  data.c4 = false
  data.c5 = false
  data.c6 = false
  data.c7 = false
  data.c8 = false
  data.c9 = false  
  data.c10 = false  

  let item = util.clone(contract)
  if (estimate != null) {
    item.estimate = estimate.id
    let price = util.getInt(estimate.price) / 10000
    if (estimate.type == 3 && estimate.subtype == 3) {
      price = price * 2
    }

    item.price = price
    data[`c${estimate.type}`] = true
  }
  data.contract = item
  data.visibleContract = true
}

function clickContract(item, index) {
  if (index.no == 11) {
    return
  }

  data.contract = util.clone(item)

  data.c1 = (item.type & 1) > 0 ? true : false
  data.c2 = (item.type & 2) > 0 ? true : false
  data.c3 = (item.type & 4) > 0 ? true : false
  data.c4 = (item.type & 8) > 0 ? true : false
  data.c5 = (item.type & 16) > 0 ? true : false
  data.c6 = (item.type & 32) > 0 ? true : false
  data.c7 = (item.type & 64) > 0 ? true : false
  data.c8 = (item.type & 128) > 0 ? true : false
  data.c9 = (item.type & 256) > 0 ? true : false
  data.c10 = (item.type & 512) > 0 ? true : false

  data.visibleContract = true
}

function clickInquiry(item, index) {
  data.editinquiry = util.clone(item)

  data.c1 = (item.type & 1) > 0 ? true : false
  data.c2 = (item.type & 2) > 0 ? true : false
  data.c3 = (item.type & 4) > 0 ? true : false
  data.c4 = (item.type & 8) > 0 ? true : false
  data.c5 = (item.type & 16) > 0 ? true : false
  data.c6 = (item.type & 32) > 0 ? true : false
  data.c7 = (item.type & 64) > 0 ? true : false
  data.c8 = (item.type & 128) > 0 ? true : false
  data.c9 = (item.type & 256) > 0 ? true : false
  data.c10 = (item.type & 512) > 0 ? true : false

  data.editcontent = item.content.replace(' ', '&nbsp;')

  data.visibleInquiry = true
}

async function clickSubmitContract() {
  let item = util.clone(data.contract)

  let contracttype = 0;
  
  if (data.c1 == true) {
    contracttype += 1;
  }

  if (data.c2 == true) {
    contracttype += 2;
  }

  if (data.c3 == true) {
    contracttype += 4;
  }

  if (data.c4 == true) {
    contracttype += 8;
  }

  if (data.c5 == true) {
    contracttype += 16;
  }

  if (data.c6 == true) {
    contracttype += 32;
  }

  if (data.c7 == true) {
    contracttype += 64;
  }

  if (data.c8 == true) {
    contracttype += 128;
  }

  if (data.c9 == true) {
    contracttype += 256;
  }

  if (data.c10 == true) {
    contracttype += 512;
  }

  item.type = contracttype
  item.contractdate = util.convertDate(item.contractdate)
  item.contractstartdate = util.convertDate(item.contractstartdate)
  item.contractenddate = util.convertDate(item.contractenddate)
  item.price = util.getInt(item.price)
  item.vat = util.getInt(item.vat)
  item.invoice = util.convertDate(item.invoice)
  item.depositdate = util.convertDate(item.depositdate)

  item.apt = data.id
  item.user = store.getters['getUser'].id

  util.loading(true)
  
  if (item.id > 0) {
    await Contract.update(item)
    util.info('저장되었습니다')
  } else {
    await Contract.insert(item)
    util.info('등록되었습니다')
  }


  let res = await Apt.get(data.id)
  let aptItem = res.item
  data.item = aptItem 
  
  await readContract(true)
  if (props.close != null && props.close != undefined) {
    props.close(item)
  }

  util.loading(false)
  data.visibleContract = false
}

async function clickUpdateInquiry() {
  let item = util.clone(data.editinquiry)

  let contracttype = 0;
  
  if (data.c1 == true) {
    contracttype += 1;
  }

  if (data.c2 == true) {
    contracttype += 2;
  }

  if (data.c3 == true) {
    contracttype += 4;
  }

  if (data.c4 == true) {
    contracttype += 8;
  }

  if (data.c5 == true) {
    contracttype += 16;
  }

  if (data.c6 == true) {
    contracttype += 32;
  }

  if (data.c7 == true) {
    contracttype += 64;
  }

  if (data.c8 == true) {
    contracttype += 128;
  }

  if (data.c9 == true) {
    contracttype += 256;
  }

  if (data.c10 == true) {
    contracttype += 512;
  }

  item.type = contracttype  
  item.content = data.editcontent
  item.status = util.getInt(data.editinquiry.status)

  util.loading(true)

  await Inquiry.update(item)
  util.info('저장되었습니다')  
  
  await readInquiry(true)
  util.loading(false)
  data.visibleInquiry = false  
}

function clickDeleteContract() {
  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)
    await Contract.remove(data.contract)
    
    await readContract(true)

    if (props.close != null && props.close != undefined) {
      props.close(item)
    }
    
    util.loading(false)
    util.info('삭제되었습니다')
    data.visibleContract = false
  })
}

function clickDeleteInquiry() {
  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)
    await Inquiry.remove({id: data.editinquiry.id})
    
    await readInquiry(true)
    util.loading(false)
    util.info('삭제되었습니다')
    data.visibleInquiry = false
  })  
}

async function clickTab(tab) {
  const name = tab.props.name

  if (name == 'contract') {
    await readContract(true)
  } else if (name == 'repair') {
    let res = await Repair.find({apt: data.id, orderby: 'r_reportdate,r_id'})
    if (res.items == null) {
      res.items = []
    }

    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }

    data.repairs = res.items
  } else if (name == 'periodic') {
    let res = await Periodic.find({apt: data.id, category: 1, orderby: 'd_id'})
    if (res.items == null) {
      res.items = []
    }

    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }

    data.periodics = res.items
  } else if (name == 'estimate') {
    await readEstimate(true)    
  }
}

async function makeComparecompany() {
  const type = util.getInt(data.estimate.type)
  const companytype = Math.pow(2, (type - 1))

  let comparecompanys = []
  for (let i = 0; i < data.allcomparecompanys.length; i++) {
    const item = data.allcomparecompanys[i]

    if (item.type & companytype) {
      comparecompanys.push(item)
    }
  }  

  data.comparecompanys = [{id: 0, name: ' '}, ...comparecompanys]

  let originalcompareestimates = []
  comparecompanys = []
  if (data.estimate.id > 0) {
    const res = await Compareestimate.find({estimate: data.estimate.id, orderby: 'e_id'})
    comparecompanys = res.items
    originalcompareestimates = util.clone(res.items)

    for (let i = 0; i < comparecompanys.length; i++) {
      comparecompanys[i].id = comparecompanys[i].comparecompany
    }

  } else {
    for (let i = 0; i < data.allcomparecompanys.length; i++) {
      const item = data.allcomparecompanys[i]

      if (item.type & companytype) {
        if (item.default == 1) {
          comparecompanys.push(item)
        }
      }
    }  
  }

  let estimates = []
  let count = 2
  // let count = comparecompanys.length
  // if (count > 2) {
  //   count = 2
  // }
  for (let i = 0; i < count; i++) {
    let  item = {
      id: 0,
      type: 1,
      subtype: 1,
      originalprice: 0,
      saleprice: 0,
      price: 0,
      discount: 0,
      person1: 0,
      person2: 0,
      person3: 0,
      person4: 0,
      person5: 0,
      person6: 0,
      person7: 0,
      person8: 0,
      person9: 0,
      person10: 0,
      writedate: null,
      financialprice: 90,
      techprice: 20,
      directprice: util.money(100000),
      printprice: util.money(100000),
      extraprice: 0,
      carprice: 0,
      car: 0,
      user: 0,
      apt: 0,
      start: '',
      date: '',
      writedate: util.getCurrentDate(),
      comparecompany: 0
    }

    if (i < comparecompanys.length) {
      const target = comparecompanys[i]
      const keys = Object.keys(target)
      for (let i = 0; i < keys.length; i++) {
        const key = keys[i]
        item[key] = target[key]
      }

      item['danger'] = target.danger
      item['machine'] = target.machine
      item['carprice'] = target.carprice
      item['print'] = 1
      item.comparecompany = util.getInt(target.id)
    }

    estimates.push(item)
  }

  data.compareestimates = estimates
  data.originalcompareestimates = util.clone(estimates)

  for (let i = 0; i < data.compareestimates.length; i++) {
    changePriceCompare(i)
  }

  for (let i = 0; i < originalcompareestimates.length; i++) {
    let item = originalcompareestimates[i]
    let originalprice = util.getInt(item.originalprice)
    let saleprice = util.getInt(item.saleprice)
    let price = util.getInt(item.price)
    comparecalc[i].totalprice = util.money(originalprice)
    data.compareestimates[i].saleprice = util.money(saleprice)
    data.compareestimates[i].price = util.money(price)
  }
}

function changeEstimateType() {
  if (data.estimate.type == 10) {
    data.estimate.saleprice = 0
    data.estimate.price = 500000
    data.estimate.subtype = 1
  } else {
    changePrice()
  }

  changeTotalprice()

  makeComparecompany()
}

function clickEstimateInsert() {
  data.menu = 'estimate'

  data.estimate = util.clone(estimate)
  
  data.estimate.financialprice = util.money(data.standardwage.financialprice1)
  data.estimate.techprice = util.money(data.standardwage.techprice1)
  data.estimate.printprice = util.money(data.standardwage.printprice1)
  data.estimate.directprice = util.money(data.standardwage.directprice)

  data.estimate.travelprice = util.money(data.standardwage.travelprice)
  data.estimate.lossprice = util.money(data.standardwage.lossprice)
  data.estimate.gasprice = util.money(data.standardwage.gasprice)


  data.estimate.travel = util.getInt(data.standardwage.travel)
  data.estimate.loss = util.getInt(data.standardwage.loss)
  data.estimate.gas = util.getInt(data.standardwage.gas)
  data.estimate.etc = util.getInt(data.standardwage.etc)
  data.estimate.danger = util.getInt(data.standardwage.danger)
  data.estimate.machine = util.getInt(data.standardwage.machine)
  data.estimate.print = util.getInt(data.standardwage.print)
  
  data.estimate.writedate = util.getCurrentDate()

  if (!util.isNull(data.estimate.writedate)) {
    let temp = data.estimate.writedate.split('-')
    let year = util.getInt(temp[0])

    let items = []
    if (util.getInt(temp[1]) > 6) {
      items.push({id: `${year-1}년 상반기`, value: `${year-1}-1`})
      items.push({id: `${year-1}년 하반기`, value: `${year-1}-2`})
      items.push({id: `${year-1}년 연간`, value: `${year-1}-0`})
      items.push({id: `${year}년 상반기`, value: `${year}-1`})
      items.push({id: `${year}년 하반기`, value: `${year}-2`})        
      items.push({id: `${year}년 연간`, value: `${year}-0`})
      items.push({id: `${year+1}년 상반기`, value: `${year+1}-1`})
      items.push({id: `${year+1}년 하반기`, value: `${year+1}-2`})
      items.push({id: `${year+1}년 연간`, value: `${year+1}-0`})
    } else {
      items.push({id: `${year-1}년 상반기`, value: `${year-1}-1`})
      items.push({id: `${year-1}년 하반기`, value: `${year-1}-2`})
      items.push({id: `${year-1}년 연간`, value: `${year-1}-0`})
      items.push({id: `${year}년 상반기`, value: `${year}-1`})
      items.push({id: `${year}년 하반기`, value: `${year}-2`})        
      items.push({id: `${year}년 연간`, value: `${year}-0`})
      items.push({id: `${year+1}년 상반기`, value: `${year+1}-1`})
      items.push({id: `${year+1}년 하반기`, value: `${year+1}-2`})
      items.push({id: `${year+1}년 연간`, value: `${year+1}-0`})
    }

    data.options = items
  }
  
  changePrice()

  makeComparecompany()
  
  data.visibleEstimate = true
}

function clickEstimate(item, index) {
  if (index.no == 8) {
    return
  }

  data.estimate = util.clone(item)
  
  let saleprice = util.money(data.estimate.saleprice)
  let price = util.money(data.estimate.price)

  if (!util.isNull(data.estimate.writedate)) {
    let temp = data.estimate.writedate.split('-')
    let year = util.getInt(temp[0])

    let items = []
    if (util.getInt(temp[1]) > 6) {
      items.push({id: `${year-1}년 상반기`, value: `${year-1}-1`})
      items.push({id: `${year-1}년 하반기`, value: `${year-1}-2`})
      items.push({id: `${year-1}년 연간`, value: `${year-1}-0`})
      items.push({id: `${year}년 상반기`, value: `${year}-1`})
      items.push({id: `${year}년 하반기`, value: `${year}-2`})        
      items.push({id: `${year}년 연간`, value: `${year}-0`})
      items.push({id: `${year+1}년 상반기`, value: `${year+1}-1`})
      items.push({id: `${year+1}년 하반기`, value: `${year+1}-2`})
      items.push({id: `${year+1}년 연간`, value: `${year+1}-0`})
    } else {
      items.push({id: `${year-1}년 상반기`, value: `${year-1}-1`})
      items.push({id: `${year-1}년 하반기`, value: `${year-1}-2`})
      items.push({id: `${year-1}년 연간`, value: `${year-1}-0`})
      items.push({id: `${year}년 상반기`, value: `${year}-1`})
      items.push({id: `${year}년 하반기`, value: `${year}-2`})        
      items.push({id: `${year}년 연간`, value: `${year}-0`})
      items.push({id: `${year+1}년 상반기`, value: `${year+1}-1`})
      items.push({id: `${year+1}년 하반기`, value: `${year+1}-2`})
      items.push({id: `${year+1}년 연간`, value: `${year+1}-0`})
    }

    data.options = items
  }  

  data.estimate.directprice = util.money(data.estimate.directprice)
  data.estimate.printprice = util.money(data.estimate.printprice)
  data.estimate.extraprice = util.money(data.estimate.extraprice)

  data.estimate.travelprice = util.money(data.estimate.travelprice)
  data.estimate.lossprice = util.money(data.estimate.lossprice)  
  data.estimate.gasprice = util.money(data.estimate.gasprice)
  data.estimate.etcprice = util.money(data.estimate.etcprice)

  if (item.event == 1) {
    data.estimate.event = true
  } else {
    data.estimate.event = false
  }

  if (item.parcel == 1) {
    data.estimate.parcel = true
  } else {
    data.estimate.parcel = false
  }

  changePrice()
  changeTotalprice()

  data.estimate.saleprice = saleprice
  data.estimate.price = price

  makeComparecompany()
  
  data.visibleEstimate = true
}

async function clickSubmitEstimate() {
  let item = util.clone(data.estimate)

  if (item.type == 3 && item.subtype == 4) {
    if (data.estimate.start == '') {
      util.error('시작 분기를 선택하세요')
      return
    }
  }

  if (item.type == 2) {
    if (data.estimate.start == '') {
      util.error('점검 분기를 선택하세요')
      return
    }
  }

  let compareestimates = util.clone(data.compareestimates)
  
  if (item.event) {
    item.event = 1
  } else {
    item.event = 2
  }

  if (item.parcel) {
    item.parcel = 1
  } else {
    item.parcel = 2
  }

  item.days = util.getInt(item.days)
  
  item.person1 = util.getInt(item.person1)
  item.person2 = util.getInt(item.person2)
  item.person3 = util.getInt(item.person3)
  item.person4 = util.getInt(item.person4)
  item.person5 = util.getInt(item.person5)
  item.person6 = util.getInt(item.person6)
  item.person7 = util.getInt(item.person7)
  item.person8 = util.getInt(item.person8)
  item.person9 = util.getInt(item.person9)
  item.person10 = util.getInt(item.person10)

  item.personprice1 = util.getInt(data.standardwage.person1)
  item.personprice2 = util.getInt(data.standardwage.person2)
  item.personprice3 = util.getInt(data.standardwage.person3)
  item.personprice4 = util.getInt(data.standardwage.person4)
  item.personprice5 = util.getInt(data.standardwage.person5)
  item.personprice6 = util.getInt(data.standardwage.person6)
  item.personprice7 = util.getInt(data.standardwage.person7)
  item.personprice8 = util.getInt(data.standardwage.person8)
  item.personprice9 = util.getInt(data.standardwage.person9)
  item.personprice10 = util.getInt(data.standardwage.person10)

  item.financialprice = util.getInt(item.financialprice)
  item.techprice = util.getInt(item.techprice)

  item.stability = util.getInt(item.stability)
  item.earthquake = util.getInt(item.earthquake)

  const type = item.type

  if (item.type == 1) {
    item.directprice = util.getInt(item.directprice)
    item.printprice = util.getInt(item.printprice)
    item.extraprice = util.getInt(item.extraprice)
  } else {
    item.travelprice = util.getInt(item.travelprice)
    item.travel = util.getInt(item.travel)
    item.lossprice = util.getInt(item.lossprice)
    item.loss = util.getInt(item.loss)
    item.gasprice = util.getInt(item.gasprice)
    item.gas = util.getInt(item.gas)
    item.carprice = util.getInt(item.carprice)
    item.car = util.getInt(item.car)
    
    item.etc = util.getInt(item.etc)
    item.danger = util.getInt(item.danger)    
    item.machine = util.getInt(item.machine)
    item.printprice = util.getInt(item.printprice)
    item.print = util.getInt(item.print)    
  }

  item.saleprice = util.getInt(item.saleprice)
  item.price = util.getInt(item.price)
  item.originalprice = item.saleprice + item.price

  item.writedate = util.convertDBDate(item.writedate)
  item.apt = data.id
  item.user = store.getters['getUser'].id

  const id = item.id


  for (let i = 0; i < compareestimates.length; i++) {
    let item = compareestimates[i]

    item.adjust = util.getInt(item.adjust)

    item.person1 = util.getInt(item.person1)
    item.person2 = util.getInt(item.person2)
    item.person3 = util.getInt(item.person3)
    item.person4 = util.getInt(item.person4)
    item.person5 = util.getInt(item.person5)
    item.person6 = util.getInt(item.person6)
    item.person7 = util.getInt(item.person7)
    item.person8 = util.getInt(item.person8)
    item.person9 = util.getInt(item.person9)
    item.person10 = util.getInt(item.person10)

    item.personprice1 = util.getInt(data.standardwage.person1)
    item.personprice2 = util.getInt(data.standardwage.person2)
    item.personprice3 = util.getInt(data.standardwage.person3)
    item.personprice4 = util.getInt(data.standardwage.person4)
    item.personprice5 = util.getInt(data.standardwage.person5)
    item.personprice6 = util.getInt(data.standardwage.person6)
    item.personprice7 = util.getInt(data.standardwage.person7)
    item.personprice8 = util.getInt(data.standardwage.person8)
    item.personprice9 = util.getInt(data.standardwage.person9)
    item.personprice10 = util.getInt(data.standardwage.person10)

    item.financialprice = util.getInt(item.financialprice)
    item.techprice = util.getInt(item.techprice)

    if (type == 1) {
      item.directprice = util.getInt(item.directprice)
      item.printprice = util.getInt(item.printprice)
      item.extraprice = util.getInt(item.extraprice)
    } else {
      item.travelprice = util.getInt(item.travelprice)
      item.travel = util.getInt(item.travel)
      item.lossprice = util.getInt(item.lossprice)
      item.loss = util.getInt(item.loss)
      item.gasprice = util.getInt(item.gasprice)
      item.gas = util.getInt(item.gas)
      item.carprice = util.getInt(item.carprice)
      item.car = util.getInt(item.car)
      
      item.etc = util.getInt(item.etc)
      item.danger = util.getInt(item.danger)    
      item.machine = util.getInt(item.machine)
      item.printprice = util.getInt(item.printprice)
      item.print = util.getInt(item.print)    
    }

    item.saleprice = util.getInt(item.saleprice)
    item.price = util.getInt(item.price)
    item.originalprice = item.saleprice + item.price

    item.writedate = util.convertDBDate(item.writedate)
    item.apt = data.id
    item.user = store.getters['getUser'].id
    item.id = id

    compareestimates[i] = item
  }

  const params = {
    estimate: item,
    compares: compareestimates
  }

  util.loading(true)
  
  if (item.id > 0) {
    await Estimate.update(params)
    util.info('저장되었습니다')
  } else {
    await Estimate.insert(params)
    util.info('등록되었습니다')
  }


  let res = await Apt.get(data.id)
  let aptItem = res.item
  data.item = aptItem 
  
  await readEstimate(true)
  if (props.close != null && props.close != undefined) {
    props.close(item)
  }

  util.loading(false)
  data.visibleEstimate = false
}

function clickDeleteEstimate() {
  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)
    await Estimate.remove(data.estimate)
    
    await readEstimate(true)

    if (props.close != null && props.close != undefined) {
      props.close(item)
    }
    
    util.loading(false)
    util.info('삭제되었습니다')
    data.visibleEstimate = false
  })
}

function autoCompare() {
  if (data.compareestimates == undefined) {
    return
  }

  for (let i = 0; i < data.compareestimates.length; i++) {
    let item = data.compareestimates[i]

    if (item.comparecompany == 0) {
      continue
    }

    const adjust = util.getInt(item.adjust)
    if (adjust == 0) {
      continue
    }

    item.person1 = data.estimate.person1
    item.person2 = data.estimate.person2
    item.person3 = data.estimate.person3
    item.person4 = data.estimate.person4
    item.person5 = data.estimate.person5
    item.person6 = data.estimate.person6
    item.person7 = data.estimate.person7
    item.person8 = data.estimate.person8
    item.person9 = data.estimate.person9
    item.person10 = data.estimate.person10

    item.days = data.estimate.days

    let persons = [[],[]]
    for (let j = 1; j <= 5; j++) {
      persons[0].push(item[`person${j}`])
    }
    for (let j = 6; j <= 10; j++) {
      persons[1].push(item[`person${j}`])
    }

    changePriceCompare(i)
    item = util.clone(data.compareestimates[i])

    let comparetotalprice = util.getInt(item.price)
    let totalprice = util.getInt(data.estimate.price) 

    if (data.estimate.type != 1) {
      let totalpersons = 0
      for (let l = 6; l <=10; l++) {
        totalpersons += item[`person${l}`]
      }
      item.car = totalpersons
      data.compareestimates[i] = item
    }
    
    if (adjust > 0) {
      if (comparetotalprice > totalprice + adjust) {
        const saleprice = comparetotalprice - (totalprice + adjust) + util.getInt(item.saleprice)
        item.saleprice = util.money(saleprice)

        data.compareestimates[i] = item
        changeSalepriceCompare(i)
      } else if (comparetotalprice < totalprice + adjust) {
        let flag = false
        for (let j = 10; j >= 1; j--) {
          if (item[`person${j}`] > 0) {
            for (let k = 0; k < 100; k++) {
              item[`person${j}`]++
              if (data.estimate.type != 1) {
                let totalpersons = 0
                for (let l = 6; l <=10; l++) {
                  totalpersons += item[`person${l}`]
                }
                item.car = totalpersons
              }
              data.compareestimates[i] = item
              changePriceCompare(i)
              item = util.clone(data.compareestimates[i])
              
              comparetotalprice = util.getInt(item.price)
              totalprice = util.getInt(data.estimate.price) 
              if (comparetotalprice > totalprice + adjust) {
                const saleprice = comparetotalprice - (totalprice + adjust) + util.getInt(item.saleprice)
                item.saleprice = util.money(saleprice)

                data.compareestimates[i] = item
                changeSalepriceCompare(i)
                item = util.clone(data.compareestimates[i])
                flag = true
                break
              }
            }
          }
          
          if (flag == true) {
            break
          }
        }
      }
    } else {
      let flag = false
      if (comparetotalprice > totalprice + adjust) {
        for (let j = 10; j >= 1; j--) {
          if (item[`person${j}`] > 1) {
            item[`person${j}`]--
            if (data.estimate.type != 1) {
              let totalpersons = 0
              for (let l = 6; l <= 10; l++) {
                totalpersons += item[`person${l}`]
              }
              item.car = totalpersons
            }
            data.compareestimates[i] = item
            changePriceCompare(i)
            item = util.clone(data.compareestimates[i])
            
            comparetotalprice = util.getInt(item.price)
            totalprice = util.getInt(data.estimate.price) 
            if (comparetotalprice <= totalprice + adjust) {
              item[`person${j}`]++
              if (data.estimate.type != 1) {
                let totalpersons = 0
                for (let l = 6; l <= 10; l++) {
                  totalpersons += item[`person${l}`]
                }
                item.car = totalpersons
              }
              data.compareestimates[i] = item
              changePriceCompare(i)
              item = util.clone(data.compareestimates[i])

              comparetotalprice = util.getInt(item.price)
              const saleprice = comparetotalprice - (totalprice + adjust) + util.getInt(item.saleprice)
              item.saleprice = util.money(saleprice)

              data.compareestimates[i] = item
              changeSalepriceCompare(i)
              item = util.clone(data.compareestimates[i])
              flag = true
              break
            }
          }
          
          if (flag == true) {
            break
          }
        }
      } else {
        for (let j = 10; j >= 1; j--) {
          if (item[`person${j}`] > 0) {
            item[`person${j}`]++
            if (data.estimate.type != 1) {
              let totalpersons = 0
              for (let l = 6; l <= 10; l++) {
                totalpersons += item[`person${l}`]
              }
              item.car = totalpersons
            }
            data.compareestimates[i] = item
            changePriceCompare(i)
            item = util.clone(data.compareestimates[i])
            
            comparetotalprice = util.getInt(item.price)
            totalprice = util.getInt(data.estimate.price) 
            if (comparetotalprice >= totalprice + adjust) {
              const saleprice = comparetotalprice - (totalprice + adjust) + util.getInt(item.saleprice)
              item.saleprice = util.money(saleprice)

              data.compareestimates[i] = item
              changeSalepriceCompare(i)
              item = util.clone(data.compareestimates[i])
              flag = true
              break
            }
          }
          
          if (flag == true) {
            break
          }
        }
      }

      if (flag == false) {
        let totalpersons = 0
        for (let l = 1; l <= 10; l++) {
          totalpersons += item[`person${l}`]
        }

        if (totalpersons > 1) {
          for (let j = 10; j >= 1; j--) {
            if (item[`person${j}`] > 0) {
              item[`person${j}`]--
              if (data.estimate.type != 1) {
                let totalpersons = 0
                for (let l = 6; l <= 10; l++) {
                  totalpersons += item[`person${l}`]
                }
                item.car = totalpersons
              }
              data.compareestimates[i] = item
              changePriceCompare(i)
              item = util.clone(data.compareestimates[i])
              
              comparetotalprice = util.getInt(item.price)
              totalprice = util.getInt(data.estimate.price) 
              if (comparetotalprice <= totalprice + adjust) {
                item[`person${j}`]++
                if (data.estimate.type != 1) {
                  let totalpersons = 0
                  for (let l = 6; l <= 10; l++) {
                    totalpersons += item[`person${l}`]
                  }
                  item.car = totalpersons
                }

                data.compareestimates[i] = item
                changePriceCompare(i)
                item = util.clone(data.compareestimates[i])

                comparetotalprice = util.getInt(item.price)
                const saleprice = comparetotalprice - (totalprice + adjust) + util.getInt(item.saleprice)
                item.saleprice = util.money(saleprice)

                data.compareestimates[i] = item
                changeSalepriceCompare(i)
                item = util.clone(data.compareestimates[i])
                flag = true
                break
              }
            }
            
            if (flag == true) {
              break
            }
          }
        }
      }

      if (flag == false) {
        comparetotalprice = util.getInt(item.price)
        totalprice = util.getInt(data.estimate.price) 
        const saleprice = comparetotalprice - (totalprice + adjust) + util.getInt(item.saleprice)
        item.saleprice = util.money(saleprice)

        data.compareestimates[i] = item
        changeSalepriceCompare(i)
      }
    }
  }
}

function changePrice() {
  let days = 1
  if (data.estimate.type != 1) {    
    days = data.estimate.days
  }

  if (days == 0) {
    days = 1
  }

  calc.price1 = util.getInt(data.estimate.person1) * data.standardwage.person1
  calc.price2 = util.getInt(data.estimate.person2) * data.standardwage.person2
  calc.price3 = util.getInt(data.estimate.person3) * data.standardwage.person3
  calc.price4 = util.getInt(data.estimate.person4) * data.standardwage.person4
  calc.price5 = util.getInt(data.estimate.person5) * data.standardwage.person5
  
  if (data.estimate.type != 1) {
    calc.price1 += util.getInt(data.estimate.person6) * data.standardwage.person1 * days
    calc.price2 += util.getInt(data.estimate.person7) * data.standardwage.person2 * days
    calc.price3 += util.getInt(data.estimate.person8) * data.standardwage.person3 * days
    calc.price4 += util.getInt(data.estimate.person9) * data.standardwage.person4 * days
    calc.price5 += util.getInt(data.estimate.person10) * data.standardwage.person5 * days
  }

  data.estimate.travel = (util.getInt(data.estimate.person6) + util.getInt(data.estimate.person7) + util.getInt(data.estimate.person8) + util.getInt(data.estimate.person9) + util.getInt(data.estimate.person10))
  data.estimate.car = (util.getInt(data.estimate.person6) + util.getInt(data.estimate.person7) + util.getInt(data.estimate.person8) + util.getInt(data.estimate.person9) + util.getInt(data.estimate.person10))
  calc.price = calc.price1 + calc.price2 + calc.price3 + calc.price4 + calc.price5
  calc.financialprice = util.getInt(Math.round(util.getFloat(calc.price) * util.getFloat(data.estimate.financialprice) / 100.0))  
  calc.techprice = util.getInt(Math.round((util.getFloat(calc.price) + util.getFloat(calc.financialprice)) * util.getFloat(data.estimate.techprice) / 100.0))

  calc.carprice = util.getInt(data.estimate.carprice) * util.getInt(data.estimate.car) * days
  calc.travelprice = util.getInt(data.estimate.travelprice) * util.getInt(data.estimate.travel) * days
  calc.lossprice = util.getInt(data.estimate.lossprice) * util.getInt(data.estimate.loss)
  calc.gasprice = util.getInt(data.estimate.gasprice) * util.getInt(data.estimate.gas) * 10
  calc.etcprice = util.getInt(Math.round(util.getFloat(calc.gasprice) * util.getFloat(data.estimate.etc) / 100)) 

  let outperson = util.getFloat(util.getInt(data.estimate.person6) * data.standardwage.person1 * days
                              + util.getInt(data.estimate.person7) * data.standardwage.person2 * days
                              + util.getInt(data.estimate.person8) * data.standardwage.person3 * days
                              + util.getInt(data.estimate.person9) * data.standardwage.person4 * days
                              + util.getInt(data.estimate.person10) * data.standardwage.person5 * days)

  calc.dangerprice = util.getInt(Math.round(outperson * util.getFloat(data.estimate.danger) / 100))  
  calc.machineprice = util.getInt(Math.round(calc.price * util.getFloat(data.estimate.machine) / 100))
  calc.printprice = util.getInt(data.estimate.printprice) * util.getInt(data.estimate.print)

  //calc.carprice = calc.lossprice + calc.gasprice + calc.etcprice  
  calc.directprice = calc.travelprice + calc.carprice + calc.dangerprice + calc.machineprice + calc.printprice

  if (data.estimate.type == 1) {    
    calc.totalprice = calc.price + calc.financialprice + calc.techprice + util.getInt(data.estimate.directprice) + util.getInt(data.estimate.printprice) + util.getInt(data.estimate.extraprice)
  } else if (data.estimate.type == 6) {    
    calc.totalprice = calc.price + calc.financialprice + calc.techprice + calc.directprice + util.getInt(data.estimate.stability) + util.getInt(data.estimate.earthquake)
  } else {
    calc.totalprice = calc.price + calc.financialprice + calc.techprice + calc.directprice
  }
  
  if (calc.totalprice > 100000) {    
    let saleprice = calc.totalprice - (parseInt(calc.totalprice / 100000) * 100000)
    data.estimate.saleprice = util.money(saleprice)
    data.estimate.price = util.money(calc.totalprice - saleprice)
  }  

  autoCompare()
}

function changeSaleprice() {  
  data.estimate.price = util.money(calc.totalprice - util.getInt(data.estimate.saleprice))
}

function changePriceCompare(index) {
  let days = 1
  if (data.estimate.type != 1) {    
    days = data.estimate.days
  }

  if (days == 0) {
    days = 1
  }

  let calc = util.clone(comparecalc[index])
  let estimate = util.clone(data.compareestimates[index])
  calc.price1 = util.getInt(estimate.person1) * data.standardwage.person1
  calc.price2 = util.getInt(estimate.person2) * data.standardwage.person2
  calc.price3 = util.getInt(estimate.person3) * data.standardwage.person3
  calc.price4 = util.getInt(estimate.person4) * data.standardwage.person4
  calc.price5 = util.getInt(estimate.person5) * data.standardwage.person5
  
  if (data.estimate.type != 1) {
    calc.price1 += util.getInt(estimate.person6) * data.standardwage.person1 * days
    calc.price2 += util.getInt(estimate.person7) * data.standardwage.person2 * days
    calc.price3 += util.getInt(estimate.person8) * data.standardwage.person3 * days
    calc.price4 += util.getInt(estimate.person9) * data.standardwage.person4 * days
    calc.price5 += util.getInt(estimate.person10) * data.standardwage.person5 * days
  }

  estimate.car = util.getInt(estimate.person6) + util.getInt(estimate.person7) + util.getInt(estimate.person8) + util.getInt(estimate.person9) + util.getInt(estimate.person10) 
  estimate.travel = util.getInt(estimate.person6) + util.getInt(estimate.person7) + util.getInt(estimate.person8) + util.getInt(estimate.person9) + util.getInt(estimate.person10) 
  calc.price = calc.price1 + calc.price2 + calc.price3 + calc.price4 + calc.price5
  calc.financialprice = util.getInt(Math.round(util.getFloat(calc.price) * util.getFloat(estimate.financialprice) / 100.0))  
  calc.techprice = util.getInt(Math.round((util.getFloat(calc.price) + util.getFloat(calc.financialprice)) * util.getFloat(estimate.techprice) / 100.0))

  calc.carprice = util.getInt(estimate.carprice) * util.getInt(estimate.car) * days
  calc.travelprice = util.getInt(estimate.travelprice) * util.getInt(estimate.travel) * days
  calc.lossprice = util.getInt(estimate.lossprice) * util.getInt(estimate.loss)
  calc.gasprice = util.getInt(estimate.gasprice) * util.getInt(estimate.gas) * 10
  calc.etcprice = util.getInt(Math.round(util.getFloat(calc.gasprice) * util.getFloat(estimate.etc) / 100)) 

  let outperson = util.getFloat(util.getInt(estimate.person6) * data.standardwage.person1 * days
                              + util.getInt(estimate.person7) * data.standardwage.person2 * days
                              + util.getInt(estimate.person8) * data.standardwage.person3 * days
                              + util.getInt(estimate.person9) * data.standardwage.person4 * days
                              + util.getInt(estimate.person10) * data.standardwage.person5 * days) 

  if (estimate.comparecompany == 2) {
    calc.dangerprice = util.getInt(Math.round(calc.price * util.getFloat(estimate.danger) / 100))
  } else {
    calc.dangerprice = util.getInt(Math.round(outperson * util.getFloat(estimate.danger) / 100))
  }
  calc.machineprice = util.getInt(Math.round(calc.price * util.getFloat(estimate.machine) / 100))
  calc.printprice = util.getInt(estimate.printprice) * util.getInt(estimate.print)

  //calc.carprice = calc.lossprice + calc.gasprice + calc.etcprice  
  calc.directprice = calc.travelprice + calc.carprice + calc.dangerprice + calc.machineprice + calc.printprice

  if (estimate.comparecompany == 3) {
    estimate.directprice = 0
    estimate.extraprice = 0
    calc.techprice = 0
  }

  if (data.estimate.type == 1) {    
    calc.totalprice = calc.price + calc.financialprice + calc.techprice + util.getInt(estimate.directprice) + util.getInt(estimate.printprice) + util.getInt(estimate.extraprice)
  } else if (data.estimate.type == 6) {    
    calc.totalprice = calc.price + calc.financialprice + calc.techprice + calc.directprice + util.getInt(data.estimate.stability) + util.getInt(data.estimate.earthquake)
  } else {
    calc.totalprice = calc.price + calc.financialprice + calc.techprice + calc.directprice
  }

  console.log(calc.totalprice)
  
  if (calc.totalprice > 100000) {    
    let saleprice = calc.totalprice - (parseInt(calc.totalprice / 100000) * 100000)
    estimate.saleprice = util.money(saleprice)
    estimate.price = util.money(calc.totalprice - saleprice)
  }  

  comparecalc[index] = calc
  let compareestimates = util.clone(data.compareestimates)
  compareestimates[index] = estimate
  data.compareestimates = compareestimates

  // if (util.getInt(estimate.adjust) > 0) {
  // }
}

function changeSalepriceCompare(index) {  
  let estimates = util.clone(data.compareestimates)
  estimates[index].price = util.money(util.getInt(comparecalc[index].totalprice) - util.getInt(estimates[index].saleprice))
  data.compareestimates = estimates
}

function clickType(value) {
  if (value == 1) {
    if (data.estimate.subtype > 2) {
      data.estimate.subtype = 1
    }
    
    if (data.estimate.subtype == 1) {
      data.estimate.financialprice = data.standardwage.financialprice1
      data.estimate.techprice = util.money(data.standardwage.techprice1)
    } else {
      data.estimate.financialprice = data.standardwage.financialprice2
      data.estimate.techprice = util.money(data.standardwage.techprice2)
    }
    
    data.estimate.printprice = util.money(data.standardwage.printprice1)
  } else if (value == 2) {
    data.estimate.techprice = util.money(data.standardwage.techprice3)
    data.estimate.financialprice = data.standardwage.financialprice3
    data.estimate.printprice = util.money(data.standardwage.printprice2)
  } else {
    data.estimate.techprice = util.money(data.standardwage.techprice4)
    data.estimate.financialprice = data.standardwage.financialprice4
    data.estimate.printprice = util.money(data.standardwage.printprice2)
  }

  data.estimate.type = value
  changePrice() 
}

function clickSubtype(value) {
  if (data.estimate.type == 1) {
    if (value == 1) {
      data.estimate.financialprice = data.standardwage.financialprice1
    } else {
      data.estimate.financialprice = data.standardwage.financialprice2
    }
  }

  // if (data.estimate.type == 3) {
  //   if (value == 3) {
  //     const compareestimates = util.clone(data.compareestimates)
  //     if (compareestimates.length > 1) {
  //       compareestimates[1].comparecompany = 0
  //       data.compareestimates = compareestimates
  //     }
  //   } else {
  //     const compareestimates = util.clone(data.compareestimates)
  //     if (compareestimates.length > 1) {
  //       if (data.originalcompareestimates.length > 1) {
  //         if (compareestimates[1].comparecompany == 0) {
  //           compareestimates[1].comparecompany = data.originalcompareestimates[1].comparecompany
  //           data.compareestimates = compareestimates
  //         }
  //       }
  //     }
  //   }
  // }
  // 

  if (data.estimate.type == 10) {
    if (value == 1) {
      data.estimate.saleprice = 0
      data.estimate.price = 500000
    } else if (value == 2) {
      data.estimate.saleprice = 0
      data.estimate.price = 500000
    } else if (value == 3) {
      data.estimate.saleprice = 0
      data.estimate.price = 1000000
    } else {
      data.estimate.saleprice = 600000
      data.estimate.price = 600000
    }

    changeTotalprice()
  }

  changePrice()
}

function clickCopyEstimate(item) {
  const id = item.id
  item.id = 0
  data.estimate = util.clone(item)

  data.estimate.directprice = util.money(data.estimate.directprice)
  data.estimate.printprice = util.money(data.estimate.printprice)
  data.estimate.extraprice = util.money(data.estimate.extraprice)

  data.estimate.travelprice = util.money(data.estimate.travelprice)
  data.estimate.lossprice = util.money(data.estimate.lossprice)  
  data.estimate.gasprice = util.money(data.estimate.gasprice)
  data.estimate.etcprice = util.money(data.estimate.etcprice)
  data.estimate.carprice = util.money(data.estimate.carprice)
  
  let saleprice = util.money(data.estimate.saleprice)
  let price = util.money(data.estimate.price)
  
  changePrice()
  data.estimate.saleprice = saleprice
  data.estimate.price = price

  data.estimate.id = id
  makeComparecompany()
  data.estimate.id = 0
  
  data.visibleEstimate = true  
}


function download(url: string, filename: string) {
  axios.get(import.meta.env.VITE_REPORT_URL + url, {
    responseType: 'blob',
    headers: {
      Authorization: 'Bearer ' + store.state.token
    }
  }).then(response => {
    const url = window.URL.createObjectURL(new Blob([response.data]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', filename);
    document.body.appendChild(link)
    link.click()
    
    util.info('다운로드가 완료되었습니다')
    util.loading(false)
  }).catch(exception => {
    util.error('다운로드가 실패하였습니다')
    util.loading(false)
  });        
}

async function clickDownloadEstimate(item, type, name, contract) {
  let res = await Apt.get(item.apt)
  let aptItem = res.item
  
  let url = '/api/download/estimate/' + item.id + '?type=' + type
  let filename

  let subtype = item.subtype

  let ext = '.xlsx'
  let typename = '견적서'

  let subtitle = item.name

  if (contract == 2) {
    if (item.type == 1 || item.type == 7 || item.type == 10) {
      url = '/api/download/contract/' + item.id + '?type=' + type
      ext = '.hml'

      if (item.estimate > 0) {
        const res = await Estimate.get(item.estimate)
        const estimate = res.item
        
        subtype = 2
        if (!util.isNull(estimate)) {
          subtype = estimate.subtype
        }

        subtitle = estimate.name
      }
    }

    typename = '계약서'
  }

  const titles = ['', '장기수선계획', '정밀안전점검', '정기안전점검', '하자보수', '하자조사', '구조안전진단', '감리', '기술자문', '순찰', '점검프로그램 사용']
  if (item.type == 1) {
    if (subtype == 1) {
      filename = `(주)에이앤비 - ${aptItem.name} - 장기수선계획 조정 ${typename}`
    } else {
      filename = `(주)에이앤비 - ${aptItem.name} - 장기수선계획 수립(조정 포함) ${typename}`      
    }
  } else if (item.type == 7 || item.type == 8) {
    if (!util.isNull(subtitle)) {
      filename = `(주)에이앤비 - ${aptItem.name} - ${subtitle} ${titles[item.type]} ${typename}`
    } else {
      filename = `(주)에이앤비 - ${aptItem.name} - ${titles[item.type]} ${typename}`
    }
  } else {
    filename = `(주)에이앤비 - ${aptItem.name} - ${titles[item.type]} ${typename}`
  }

  if (type > 0) {
    filename += ' - 비교 견적 ' + name
  }

  if (item.type == 3 && item.subtype == 3 && type == 2) {
    ext = '.zip'
  }

  filename += ext

  download(url, filename)  
}

function clickInsertContract(item) {
  clickContractInsert(item)
}

async function clickEstimateByContract(id) {
  const res = await Estimate.get(id)
  clickEstimate(res.item, 0)
}

async function clickDownloadEstimateByContract(target) {
  let item = util.clone(target)
  const res = await Estimate.get(item.estimate)
  if (!util.isNull(res.item)) {
    item.type = res.item.type

    if (item.type == 1) {
      clickDownloadEstimate(item, 0, '', 2)
    } else if (item.type == 7) {
      clickDownloadEstimate(item, 0, '', 2)
    } else if (item.type == 10) {
      clickDownloadEstimate(item, 0, '', 2)
    } else {
      clickDownloadEstimate(res.item, 0, '', 1)
    }
  } else {
    if (item.type & 1) {
      clickDownloadEstimate(item, 0, '', 2)
    }
  }
}

function changeTotalprice() {
  calc.programprice = util.getInt(data.estimate.price) + util.getInt(data.estimate.saleprice)
}

function changeTotalpriceCompare(index) {
}

function changeDays() {
  changePrice()
  for (let i = 0; i < data.compareestimates.length; i++) {
    changePriceCompare(i)
  }
}

</script>
<style>
.date .el-input__wrapper {
  padding:0px;  
}

.date .el-input__inner {
  text-align:center;

}

p {
  line-height: 80%;
}

.inputNumber .el-input__inner {
  text-align: right;
}

</style>  

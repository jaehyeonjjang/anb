<template>
  <Title title="리포트" />

  <el-tabs v-model="data.menu" @tab-click="clickTab">
    <el-tab-pane label="통합총괄자료" name="total">
      <div :style="{'height': data.height, 'overflow': 'auto'}">
        <Total ref="total" />
      </div>
    </el-tab-pane>
    <el-tab-pane label="총론" name="totalplan">
      <div :style="{'height': data.height, 'overflow': 'auto'}">
        <Totalplan ref="totalplan" />
      </div>
    </el-tab-pane>
    <el-tab-pane label="총론2" name="plan">
      <div :style="{'height': data.height, 'overflow': 'auto'}">
        <Plan ref="plan" />
      </div>
    </el-tab-pane>
    <el-tab-pane label="연도별공사예정현황" name="year">
      <div :style="{'height': data.height, 'overflow': 'auto'}">
        <Year ref="year" />
      </div>
    </el-tab-pane>
    <el-tab-pane label="통합세부자료" name="combine">
      <div :style="{'height': data.height, 'overflow': 'auto'}">
        <Combine ref="combine" />
      </div>
    </el-tab-pane>
    <el-tab-pane label="공사종별 수선계획금액 집계표" name="summary">
      <div :style="{'height': data.height, 'overflow': 'auto'}">
        <Summary ref="summary" />
      </div>
    </el-tab-pane>
  </el-tabs>
  
</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import type { TabsPaneContext } from 'element-plus'
import { util }  from "~/global"

const data = reactive({
  menu: 'total',
  height: 0
})

const total = ref({});
const plan = ref({});
const year = ref({});
const combine = ref({});
const summary = ref({});
const totalplan = ref({});

function clickTab(item) {
  let name = item.props.name

  if (name == 'total') {
    total.value.readData()
  } else if (name == 'plan') {
    plan.value.readData()
  } else if (name == 'year') {
    year.value.readData()
  } else if (name == 'combine') {
    combine.value.readData()
  } else if (name == 'summary') {
    summary.value.readData()
  } else if (name == 'totalplan') {
    totalplan.value.readData()    
  }
}

function setWindowSize() {
  data.height = (window.innerHeight - 170 - 20) + 'px'
}

onMounted(() => {
  total.value.readData()

  setWindowSize()
  window.addEventListener('resize', setWindowSize)
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
})

</script>

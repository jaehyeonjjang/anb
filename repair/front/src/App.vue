<template>
  <BaseHeader v-if="store.state.token != null && store.state.token != ''" />

  <div style="display:flex;">
    <RepairMenu v-if="data.menu == 'repair'" />
    <DetailMenu v-if="data.menu == 'detail'" />
    <PeriodicMenu v-if="data.menu == 'periodic'" />
    <AptMenu v-if="data.menu == 'apt'" />
    <ManagementSettingMenu v-if="data.menu == 'management/setting'" />
    <ManagementRepairMenu v-if="data.menu == 'management/repair'" />
    <div style="flex:1;">
      <div style="padding: 10px 10px;">
        <router-view />
      </div>
    </div>
  </div>
  
</template>

<script setup lang="ts">
import { reactive, watchEffect } from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
const store = useStore()
const route = useRoute()

const data = reactive({
  menu: '',
})

watchEffect(() => {
  var s = route.path.split('/')

  if (s[1] == 'management') {
    if (s.length == 3) {
      data.menu = s[1]
    } else {
      data.menu = `${s[1]}/${s[2]}`
    }
  } else {
    data.menu = s[2]    
  }  
})

</script>

<style>
body {
  display: flex;
  flex-direction: column;
  height: 100%;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

</style>

<template>
  <apt-detail :is-edit="true" />
</template>

<script>
import AptDetail from './components/AptDetail'

export default {
  name: 'AptUpdate',
  components: { AptDetail }
}
</script>


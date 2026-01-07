<template>
  <apt-detail :is-edit="false" />
</template>

<script>
import AptDetail from './components/AptDetail'

export default {
  name: 'AptInsert',
  components: { AptDetail }
}
</script>


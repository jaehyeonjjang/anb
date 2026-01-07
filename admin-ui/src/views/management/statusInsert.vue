<template>
  <status-detail :is-edit="false" />
</template>

<script>
import StatusDetail from './components/StatusDetail'

export default {
  name: 'StatusInsert',
  components: { StatusDetail }
}
</script>


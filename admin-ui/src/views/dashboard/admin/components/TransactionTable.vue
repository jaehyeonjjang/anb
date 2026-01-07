<template>
  <el-table :data="items" style="width: 100%;padding-top: 15px;">
    <el-table-column label="Group">
      <template slot-scope="{row}">
        {{ getGroup(row.Group) }}
      </template>
    </el-table-column>
    <el-table-column label="Api">
      <template slot-scope="{row}">
        {{ getApi(row.Api) }}
      </template>
    </el-table-column>
    <el-table-column label="Content">
      <template slot-scope="{row}">
        {{ row.Content }}
      </template>
    </el-table-column>

    <el-table-column width="180px" align="center" label="Reg Date">
      <template slot-scope="{row}">
        <span>{{ row.Date | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
      </template>
    </el-table-column>

  </el-table>
</template>

<script>
import axios from 'axios'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        success: 'success',
        pending: 'danger'
      }
      return statusMap[status]
    },
    orderNoFilter(str) {
      return str.substring(0, 30)
    }
  },
  data() {
    return {
      items: [],
      groups: [],
      apis: []
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    getGroup: function(value) {
      for (let i = 0; i < this.groups.length; i++) {
        if (this.groups[i].Id === value) {
          return this.groups[i].Name
        }
      }

      return ''
    },
    getApi: function(value) {
      for (let i = 0; i < this.apis.length; i++) {
        if (this.apis[i].Id === value) {
          return this.apis[i].Name
        }
      }

      return ''
    },
    fetchData: async function() {
      const params = 'user=' + this.$store.state.user.token + '&page=1&pagesize=10'
      const response = await axios({
        method: 'GET',
        url: 'http://localhost:3001/api/error?' + params
      })

      this.items = response.data.items

      console.log(this.items)

      const groupResponse = await axios({
        method: 'GET',
        url: 'http://localhost:3001/api/group'
      })

      this.groups = groupResponse.data.items

      const apiResponse = await axios({
        method: 'GET',
        url: 'http://localhost:3001/api/api'
      })

      this.apis = apiResponse.data.items
    }
  }
}
</script>

<template>
  <div class="dashboard-editor-container">

    <el-row :gutter="40" class="panel-group">
      <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('newVisitis')">
          <div class="card-panel-icon-wrapper icon-people">
            <svg-icon icon-class="peoples" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              {{ userCount }}
            </div>
            <count-to :start-val="0" :end-val="users" :duration="2600" class="card-panel-num" />
          </div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('messages')">
          <div class="card-panel-icon-wrapper icon-message">
            <svg-icon icon-class="tree-table" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              {{ aptgroupCount }}

            </div>
            <count-to :start-val="0" :end-val="groups" :duration="3000" class="card-panel-num" />
          </div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('purchases')">
          <div class="card-panel-icon-wrapper icon-money">
            <svg-icon icon-class="form" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              {{ aptCount }}
            </div>
            <count-to :start-val="0" :end-val="apis" :duration="3200" class="card-panel-num" />
          </div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('shoppings')">
          <div class="card-panel-icon-wrapper icon-shopping">
            <svg-icon icon-class="excel" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              {{ reportCount }}
            </div>
            <count-to :start-val="0" :end-val="errors" :duration="3600" class="card-panel-num" />
          </div>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-table v-loading="listLoading" :data="items" border fit highlight-current-row style="width: 100%">
        <el-table-column align="center" label="ID" width="80">
          <template slot-scope="{row}">
            <span>{{ row.id }}</span>
          </template>
        </el-table-column>

        <el-table-column v-if="$store.getters.isAdmin" min-width="150px" label="업체">
          <template slot-scope="{row}">
            <span>{{ getCompany(row.company) }}</span>
          </template>
        </el-table-column>

        <el-table-column min-width="150px" label="현장명">
          <template slot-scope="{row}">
            <span>{{ getAptgroup(row.aptgroup) }}</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="작업구분" width="100px">
          <template slot-scope="{row}">
            <span>{{ getType(row.type) }}</span>
          </template>
        </el-table-column>

        <el-table-column min-width="150px" label="작업명">
          <template slot-scope="{row}">
            <router-link :to="'/apt/apt/update/'+row.id" class="link-type">
              <span>{{ row.name }}</span>
            </router-link>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="시작일" width="100px">
          <template slot-scope="{row}">
            <span>{{ row.startdate !== '' ? $moment(row.startdate).format('YYYY-MM-DD') : '' }}</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="종료일" width="100px">
          <template slot-scope="{row}">
            <span>{{ row.enddate !== '' ? $moment(row.enddate).format('YYYY-MM-DD') : '' }}</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="책임 기술자" width="90px">
          <template slot-scope="{row}">
            <span>{{ getUsers(row.master) }}</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="참여 기술자" width="90px">
          <template slot-scope="{row}">
            <span>{{ row.submaster }}</span>
          </template>
        </el-table-column>

        <el-table-column class-name="status-col" label="작업상태" width="80px">
          <template slot-scope="{row}">
            <el-tag :type="row.status | statusFilter">
              {{ getStatus(row.status) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column width="150px" align="center" label="등록일">
          <template slot-scope="{row}">
            <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
          </template>
        </el-table-column>

      </el-table>

    </el-row>
  </div>
</template>

<script>
import request from '@/utils/request'

export default {
  name: 'DashboardAdmin',
  components: {
  },
  filters: {
    statusFilter(value) {
      const values = ['info', 'primary', 'success', 'danger']
      return values[value]
    }
  },
  data() {
    return {
      userCount: 100,
      aptgroupCount: 102,
      aptCount: 103,
      reportCount: 104,
      items: null,
      total: 0,
      downloadLoading: false,
      listLoading: true,
      page: 1,
      pagesize: 10,
      search: '',
      searchstatus: 0,
      searchcompany: 0,
      searchaptgroup: 0,
      searchtype: 'loginid',
      types: [
        { name: '구분', id: 0 },
        { name: '정밀안전점검', id: 1 },
        { name: '정기안전점검', id: 2 },
        { name: '하자조사', id: 3 },
        { name: '장기수선', id: 4 }
      ],
      statuss: [
        { name: '준비', id: 0 },
        { name: '착수', id: 1 },
        { name: '완료', id: 2 },
        { name: '중단', id: 3 }
      ],
      companys: [],
      aptgroups: [],
      users: []
    }
  },
  created: async function() {
    console.log('created')
    const response = await request({
      method: 'GET',
      url: '/api/company'
    })

    this.companys = [{ id: 0, name: '업체' }, ...response.data]

    if (!this.$store.getters.isAdmin) {
      this.searchcompany = this.$store.getters.company
    }

    const responseAptgroup = await request({
      method: 'GET',
      url: '/api/aptgroup?company=' + this.searchcompany
    })

    this.aptgroups = [{ id: 0, name: '현장' }, ...responseAptgroup.data]

    const responseUser = await request({
      method: 'GET',
      url: '/api/user?company=' + this.searchcompany
    })

    this.users = responseUser.data

    this.userCount = this.users.length
    this.aptgroupCount = this.aptgroups.length

    const responseApt = await request({
      method: 'GET',
      url: '/api/apt?company=' + this.searchcompany
    })

    this.aptCount = responseApt.data.length

    const responseReport = await request({
      method: 'GET',
      url: '/api/apt?company=' + this.searchcompany + '&report=4'
    })

    this.reportCount = responseReport.data.length

    this.getList()
  },
  methods: {
    getCompany: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.companys.length; i++) {
        if (value === this.companys[i].id) {
          return this.companys[i].name
        }
      }

      return ''
    },
    getType: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.types.length; i++) {
        if (value === this.types[i].id) {
          return this.types[i].name
        }
      }

      return ''
    },
    getAptgroup: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.aptgroups.length; i++) {
        if (value === this.aptgroups[i].id) {
          return this.aptgroups[i].name
        }
      }

      return ''
    },
    getUsers: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.users.length; i++) {
        if (value === this.users[i].id) {
          return this.users[i].name
        }
      }

      return ''
    },
    getStatus: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.statuss.length; i++) {
        if (value === this.statuss[i].id) {
          return this.statuss[i].name
        }
      }

      return ''
    },
    getList: async function() {
      this.listLoading = true

      const params = `page=${this.page - 1}&size=${this.pagesize}&status=${this.searchstatus}&aptgroup=${this.searchaptgroup}&company=${this.searchcompany}&name=${encodeURIComponent(this.search)}`
      const response = await request({
        method: 'GET',
        url: '/api/apt?' + params
      })

      for (let i = 0; i < response.data.content.length; i++) {
        const item = response.data.content[i]
        const params = `apt=${item.id}`
        const responseSubmaster = await request({
          method: 'GET',
          url: '/api/aptsubmaster?' + params
        })

        response.data.content[i].submaster = responseSubmaster.data.length
      }

      this.items = response.data.content
      this.total = response.data.totalElements
      this.pagesize = response.data.pageable.pageSize

      this.listLoading = false
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}

.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
    border-color: rgba(0, 0, 0, .05);

    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }

      .icon-people {
        background: #40c9c6;
      }

      .icon-message {
        background: #36a3f7;
      }

      .icon-money {
        background: #f4516c;
      }

      .icon-shopping {
        background: #34bfa3
      }
    }

    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 32px;
        margin-top:20px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}

@media (max-width:550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    .svg-icon {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>

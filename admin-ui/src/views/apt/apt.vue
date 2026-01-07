<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="searchstatus" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in statuss"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-select v-if="$store.getters.isAdmin" v-model="searchcompany" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in companys"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-select v-model="searchaptgroup" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in aptgroups"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-input v-model="search" placeholder="검색" style="width:300px;" class="filter-item" @keyup.enter.native="clickSearch" />

      <el-button class="filter-item" type="primary" style="margin-left:15px;" icon="el-icon-search" @click="clickSearch">
        검색
      </el-button>
      <el-button :loading="downloadLoading" style="float:right;" class="filter-item" type="success" icon="el-icon-circle-plus-outline" @click="$router.push('/apt/apt/insert')">
        등록
      </el-button>
    </div>

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
          <span>{{ getUser(row.master) }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="참여 기술자" width="90px">
        <template slot-scope="{row}">
          <span>{{ row.submaster === 0 ? '' : row.submaster + '명' }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="작업상태" width="80px">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ getStatus(row.status) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column width="80px" align="center" label="작성자">
        <template slot-scope="{row}">
          <span>{{ getUser(row.user) }}</span>
        </template>
      </el-table-column>

      <el-table-column width="80px" align="center" label="수정자">
        <template slot-scope="{row}">
          <span>{{ getUser(row.user) }}</span>
        </template>
      </el-table-column>

      <el-table-column width="150px" align="center" label="등록일">
        <template slot-scope="{row}">
          <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="" width="270px">
        <template slot-scope="{row}">

          <el-button type="primary" size="small" icon="el-icon-document-copy" style="margin-right:10px;" @click="clickCopy(row.id)">
            복사
          </el-button>

          <router-link :to="'/apt/apt/update/'+row.id" style="margin-right:10px;">
            <el-button type="primary" size="small" icon="el-icon-edit">
              수정
            </el-button>
          </router-link>

          <el-button type="warning" size="small" icon="el-icon-delete" @click="clickDelete(row.id)">
            삭제
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="page" :limit.sync="pagesize" @pagination="getList" />
  </div>
</template>

<script>
import request from '@/utils/request'
import Pagination from '@/components/Pagination'

export default {
  name: 'AptList',
  components: { Pagination },
  filters: {
    statusFilter(value) {
      const values = ['info', 'primary', 'success', 'danger']
      return values[value]
    }
  },
  data() {
    return {
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
    if (this.$route.params.aptgroup) {
      this.searchaptgroup = this.$route.params.aptgroup
    }

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
    getUser: function(value) {
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
    },
    clickSearch: function() {
      this.getList()
    },
    clickCopy: async function(id) {
      this.$confirm('해당 작업을 복사하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        await request({
          method: 'POST',
          url: '/api/apt/' + id + '/copy'
        })

        this.$message({
          type: 'success',
          message: '복사되었습니다'
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '복사가 취소되었습니다'
        })
      })
    },
    clickDelete: async function(id) {
      this.$confirm('삭제하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        this.$message({
          type: 'success',
          message: '삭제되었습니다'
        })

        await request({
          method: 'DELETE',
          url: '/api/apt/' + id
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제가 취소되었습니다'
        })
      })
    }
  }
}
</script>

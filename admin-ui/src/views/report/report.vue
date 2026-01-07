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

    </div>

    <el-table v-loading="listLoading" :data="items" border fit highlight-current-row style="width: 100%" @header-click="clickHeader">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column v-if="$store.getters.isAdmin" min-width="150px" label="업체" sortable>
        <template slot-scope="{row}">
          <span>{{ getCompany(row.company) }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="150px" label="현장명" sortable>
        <template slot-scope="{row}">
          <span>{{ getAptgroup(row.aptgroup) }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="작업구분" width="100px" sortable>
        <template slot-scope="{row}">
          <span>{{ getType(row.type) }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="150px" label="작업명" sortable>
        <template slot-scope="{row}">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="실시결과 요약표" width="130px">
        <template slot-scope="{row}">
          <el-button type="success" size="small" icon="el-icon-printer" @click="clickSummary(row.id)">
            보기
          </el-button>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="출력" width="100px">
        <template slot-scope="{row}">
          <el-button v-if="row.report === 4" type="primary" size="small" icon="el-icon-printer" @click="clickReport(row.id)">
            출력
          </el-button>
          <span v-if="row.report !== 4">{{ getReport(row.report) }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="" width="100px">
        <template slot-scope="{row}">
          <el-button v-if="row.report === 1 || row.report === 4" type="warning" size="small" icon="el-icon-edit" @click="clickMake(row.id)">
            생성
          </el-button>
        </template>
      </el-table-column>

    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="page" :limit.sync="pagesize" @pagination="getList(false)" />

    <el-dialog
      title="실시결과 요약표"
      :visible.sync="visible"
      width="800px">
      

      <el-tabs v-model="menu" style="margin-top:-30px;">
        <el-tab-pane v-for="(tab, index) in summarys" :label="tab.Name" :name="`menu${index}`">

          <el-table :data="tab.Items" border fit highlight-current-row style="width: 100%" height="500px">
            <el-table-column align="center" label="부위(부재)">
              <template slot-scope="{row}">
                {{row.Name}}
              </template>
            </el-table-column>

            <el-table-column align="center" label="점검 결과">
              <template slot-scope="{row}">
                <div v-for="item in row.Fault">{{item}}</div>
              </template>
            </el-table-column>

            <el-table-column align="center" label="조치 필요사항">
              <template slot-scope="{row}">
                <div v-for="item in row.Method">{{item}}</div>
              </template>
            </el-table-column>
          </el-table>
          
        </el-tab-pane>
      </el-tabs>  
      
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="visible = false">닫기</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import request from '@/utils/request'
import backend from '@/utils/requestBackend'
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
      orderby: 'aptgroupDesc',
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
      users: [],
      visible: false,
      summarys: [],
      menu: 'menu0'
    }
  },
  created: async function() {
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

    this.getList(false)

    const self = this
    setInterval(function() {
      self.getList(true)
    }, 1000 * 5)
  },
  methods: {
    getReport: function(value) {
      const titles = ['', '', '대기중', '생성중', '출력']
      return titles[value]
    },
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
    getList: async function(flag) {
      if (flag === false) {
        this.listLoading = true
      }

      const params = `page=${this.page - 1}&size=${this.pagesize}&orderby=${this.orderby}&status=${this.searchstatus}&aptgroup=${this.searchaptgroup}&company=${this.searchcompany}&name=${encodeURIComponent(this.search)}`
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

      if (flag === false) {
        this.listLoading = false
      }
    },
    clickSearch: function() {
      this.getList()
    },
    clickReport: function(id) {
      location.href = process.env.VUE_APP_BASE_API + '/api/download/report/' + id
    },
    clickMake: function(id) {
      this.$confirm('보고서를 생성하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        const item = {
          apt: id,
          image: 0,
          status: 2
        }

        await request({
          method: 'POST',
          url: '/api/report',
          data: item
        })

        this.$message({
          type: 'success',
          message: '보고서 생성중입니다. 생성이 완료되면 출력 버튼이 나타납니다'
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '보고서 생성이 취소되었습니다'
        })
      })
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
    },
    clickSort(sortProps) {
      this.clickHeader(sortProps.column, event)
    },
    clickHeader(column, event) {
      let orderby = ''
      if (column.label === '현장명') {
        orderby = 'aptgroup'
      } else if (column.label === '작업구분') {
        orderby = 'type'
      } else if (column.label === '작업명') {
        orderby = 'name'
      } else {
        return
      }

      if (column.order === 'descending') {
        orderby += 'Desc'
      } else {
        orderby = ''
      }

      this.orderby = orderby

      this.getList()
    },
    async clickSummary(id) {
      const response = await backend({
        method: 'GET',
        url: '/api/report/summary?id=' + id
      })

      this.summarys = response.items
      this.visible = true
      console.log(response)

    }
  }
}
</script>


<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
         <el-table-column align="left" label="id" prop="ID" width="180">
            <!-- <template #default="scope">{{ formatDate(scope.row.id) }}</template> -->
         </el-table-column>
          <el-table-column align="left" label="币名" prop="tokenName" width="120" />
          <el-table-column align="left" label="合约地址" prop="contractAddress" width="120" />
          <el-table-column align="left" label="rawMsg" prop="msgId" width="120" >
             <template #default="scope"><a :href="'/#/layout/admin/rawmsg?ID='+scope.row.msgId">{{ scope.row.msgId }}</a></template>
          </el-table-column>
          <el-table-column align="left" label="Bond曲线" prop="bondingCurve" width="120" />
          <el-table-column align="left" label="progress" prop="progress" width="120" />
          <el-table-column align="left" label="余额" prop="balance" width="120" />
          <el-table-column align="left" label="创建时间" prop="createdTime" width="120" />
          <el-table-column align="left" label="市场容量" prop="marketCap" width="120" />
          <el-table-column align="left" label="评价数量" prop="commentsCount" width="120" />
          <el-table-column align="left" label="开发钱包" prop="devWallet" width="120" />
          <el-table-column align="left" label="开发钱包" prop="devHolding" width="120" />
          <el-table-column align="left" label="X链接" prop="twitterLink" width="120" />
          <el-table-column align="left" label="T链接" prop="telegramLink" width="120" />
          <el-table-column align="left" label="Web链接" prop="websiteLink" width="120" />
          <el-table-column align="left" label="Price链接" prop="livePriceLink" width="120" />
          <el-table-column align="left" label="PriceWave5m" prop="priceWave5m" width="120" />
          <el-table-column align="left" label="PriceWave1h" prop="priceWave1h" width="120" />
          <el-table-column align="left" label="PriceWave6h" prop="priceWave6h" width="120" />
          <el-table-column align="left" label="TXs" prop="tXs" width="120" />
          <el-table-column align="left" label="vol" prop="vol" width="120" />
          <el-table-column align="left" label="Liq" prop="liq" width="120" />
          <el-table-column align="left" label="Holder" prop="holder" width="120" />
          <el-table-column align="left" label="Open" prop="open" width="120" />
          <el-table-column align="left" label="NoMint" prop="noMint" width="120" />
          <el-table-column align="left" label="BlackList" prop="blackList" width="120" />
          <el-table-column align="left" label="Burnt" prop="burnt" width="120" />
          <el-table-column align="left" label="Top10" prop="top10" width="120" />
          <el-table-column align="left" label="DevAddr" prop="devAddr" width="120" />
          <el-table-column align="left" label="DevInfo" prop="devInfo" width="120" />
          <el-table-column align="left" label="DevBurn" prop="devBurn" width="120" />
          <el-table-column align="left" label="AdvInfo" prop="advInfo" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateCoinsFunc(scope.row)">变更</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id:"  prop="id" >
              <el-date-picker v-model="formData.id" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="币名:"  prop="tokenName" >
              <el-input v-model="formData.tokenName" :clearable="true"  placeholder="请输入币名" />
            </el-form-item>
            <el-form-item label="合约地址:"  prop="contractAddress" >
              <el-input v-model="formData.contractAddress" :clearable="true"  placeholder="请输入合约地址" />
            </el-form-item>
            <el-form-item label="Bond曲线:"  prop="bondingCurve" >
              <el-input v-model="formData.bondingCurve" :clearable="true"  placeholder="请输入Bond曲线" />
            </el-form-item>
            <el-form-item label="progress:"  prop="progress" >
              <el-input v-model="formData.progress" :clearable="true"  placeholder="请输入progress" />
            </el-form-item>
            <el-form-item label="余额:"  prop="balance" >
              <el-input v-model="formData.balance" :clearable="true"  placeholder="请输入余额" />
            </el-form-item>
            <el-form-item label="创建时间:"  prop="createdTime" >
              <el-input v-model="formData.createdTime" :clearable="true"  placeholder="请输入创建时间" />
            </el-form-item>
            <el-form-item label="市场容量:"  prop="marketCap" >
              <el-input v-model="formData.marketCap" :clearable="true"  placeholder="请输入市场容量" />
            </el-form-item>
            <el-form-item label="评价数量:"  prop="commentsCount" >
              <el-input v-model="formData.commentsCount" :clearable="true"  placeholder="请输入评价数量" />
            </el-form-item>
            <el-form-item label="开发钱包:"  prop="devWallet" >
              <el-input v-model="formData.devWallet" :clearable="true"  placeholder="请输入开发钱包" />
            </el-form-item>
            <el-form-item label="开发钱包:"  prop="devHolding" >
              <el-input v-model="formData.devHolding" :clearable="true"  placeholder="请输入开发钱包" />
            </el-form-item>
            <el-form-item label="X链接:"  prop="twitterLink" >
              <el-input v-model="formData.twitterLink" :clearable="true"  placeholder="请输入X链接" />
            </el-form-item>
            <el-form-item label="T链接:"  prop="telegramLink" >
              <el-input v-model="formData.telegramLink" :clearable="true"  placeholder="请输入T链接" />
            </el-form-item>
            <el-form-item label="Web链接:"  prop="websiteLink" >
              <el-input v-model="formData.websiteLink" :clearable="true"  placeholder="请输入Web链接" />
            </el-form-item>
            <el-form-item label="Price链接:"  prop="livePriceLink" >
              <el-input v-model="formData.livePriceLink" :clearable="true"  placeholder="请输入Price链接" />
            </el-form-item>
            <el-form-item label="PriceWave5m:"  prop="priceWave5m" >
              <el-input v-model="formData.priceWave5m" :clearable="true"  placeholder="请输入PriceWave5m" />
            </el-form-item>
            <el-form-item label="PriceWave1h:"  prop="priceWave1h" >
              <el-input v-model="formData.priceWave1h" :clearable="true"  placeholder="请输入PriceWave1h" />
            </el-form-item>
            <el-form-item label="PriceWave6h:"  prop="priceWave6h" >
              <el-input v-model="formData.priceWave6h" :clearable="true"  placeholder="请输入PriceWave6h" />
            </el-form-item>
            <el-form-item label="TXs:"  prop="tXs" >
              <el-input v-model="formData.tXs" :clearable="true"  placeholder="请输入TXs" />
            </el-form-item>
            <el-form-item label="vol:"  prop="vol" >
              <el-input v-model="formData.vol" :clearable="true"  placeholder="请输入vol" />
            </el-form-item>
            <el-form-item label="Liq:"  prop="liq" >
              <el-input v-model="formData.liq" :clearable="true"  placeholder="请输入Liq" />
            </el-form-item>
            <el-form-item label="Holder:"  prop="holder" >
              <el-input v-model="formData.holder" :clearable="true"  placeholder="请输入Holder" />
            </el-form-item>
            <el-form-item label="Open:"  prop="open" >
              <el-input v-model="formData.open" :clearable="true"  placeholder="请输入Open" />
            </el-form-item>
            <el-form-item label="NoMint:"  prop="noMint" >
              <el-input v-model="formData.noMint" :clearable="true"  placeholder="请输入NoMint" />
            </el-form-item>
            <el-form-item label="BlackList:"  prop="blackList" >
              <el-input v-model="formData.blackList" :clearable="true"  placeholder="请输入BlackList" />
            </el-form-item>
            <el-form-item label="Burnt:"  prop="burnt" >
              <el-input v-model="formData.burnt" :clearable="true"  placeholder="请输入Burnt" />
            </el-form-item>
            <el-form-item label="Top10:"  prop="top10" >
              <el-input v-model="formData.top10" :clearable="true"  placeholder="请输入Top10" />
            </el-form-item>
            <el-form-item label="DevAddr:"  prop="devAddr" >
              <el-input v-model="formData.devAddr" :clearable="true"  placeholder="请输入DevAddr" />
            </el-form-item>
            <el-form-item label="DevInfo:"  prop="devInfo" >
              <el-input v-model="formData.devInfo" :clearable="true"  placeholder="请输入DevInfo" />
            </el-form-item>
            <el-form-item label="DevBurn:"  prop="devBurn" >
              <el-input v-model="formData.devBurn" :clearable="true"  placeholder="请输入DevBurn" />
            </el-form-item>
            <el-form-item label="AdvInfo:"  prop="advInfo" >
              <el-input v-model="formData.advInfo" :clearable="true"  placeholder="请输入AdvInfo" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions column="1" border>
                    <el-descriptions-item label="id">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="币名">
                        {{ detailFrom.tokenName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="合约地址">
                        {{ detailFrom.contractAddress }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Bond曲线">
                        {{ detailFrom.bondingCurve }}
                    </el-descriptions-item>
                    <el-descriptions-item label="progress">
                        {{ detailFrom.progress }}
                    </el-descriptions-item>
                    <el-descriptions-item label="余额">
                        {{ detailFrom.balance }}
                    </el-descriptions-item>
                    <el-descriptions-item label="创建时间">
                        {{ detailFrom.createdTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="市场容量">
                        {{ detailFrom.marketCap }}
                    </el-descriptions-item>
                    <el-descriptions-item label="评价数量">
                        {{ detailFrom.commentsCount }}
                    </el-descriptions-item>
                    <el-descriptions-item label="开发钱包">
                        {{ detailFrom.devWallet }}
                    </el-descriptions-item>
                    <el-descriptions-item label="开发钱包">
                        {{ detailFrom.devHolding }}
                    </el-descriptions-item>
                    <el-descriptions-item label="X链接">
                        {{ detailFrom.twitterLink }}
                    </el-descriptions-item>
                    <el-descriptions-item label="T链接">
                        {{ detailFrom.telegramLink }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Web链接">
                        {{ detailFrom.websiteLink }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Price链接">
                        {{ detailFrom.livePriceLink }}
                    </el-descriptions-item>
                    <el-descriptions-item label="PriceWave5m">
                        {{ detailFrom.priceWave5m }}
                    </el-descriptions-item>
                    <el-descriptions-item label="PriceWave1h">
                        {{ detailFrom.priceWave1h }}
                    </el-descriptions-item>
                    <el-descriptions-item label="PriceWave6h">
                        {{ detailFrom.priceWave6h }}
                    </el-descriptions-item>
                    <el-descriptions-item label="TXs">
                        {{ detailFrom.tXs }}
                    </el-descriptions-item>
                    <el-descriptions-item label="vol">
                        {{ detailFrom.vol }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Liq">
                        {{ detailFrom.liq }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Holder">
                        {{ detailFrom.holder }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Open">
                        {{ detailFrom.open }}
                    </el-descriptions-item>
                    <el-descriptions-item label="NoMint">
                        {{ detailFrom.noMint }}
                    </el-descriptions-item>
                    <el-descriptions-item label="BlackList">
                        {{ detailFrom.blackList }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Burnt">
                        {{ detailFrom.burnt }}
                    </el-descriptions-item>
                    <el-descriptions-item label="Top10">
                        {{ detailFrom.top10 }}
                    </el-descriptions-item>
                    <el-descriptions-item label="DevAddr">
                        {{ detailFrom.devAddr }}
                    </el-descriptions-item>
                    <el-descriptions-item label="DevInfo">
                        {{ detailFrom.devInfo }}
                    </el-descriptions-item>
                    <el-descriptions-item label="DevBurn">
                        {{ detailFrom.devBurn }}
                    </el-descriptions-item>
                    <el-descriptions-item label="AdvInfo">
                        {{ detailFrom.advInfo }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createCoins,
  deleteCoins,
  deleteCoinsByIds,
  updateCoins,
  findCoins,
  getCoinsList
} from '@/api/coin/coins'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Coins'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: new Date(),
            tokenName: '',
            contractAddress: '',
            bondingCurve: '',
            progress: '',
            balance: '',
            createdTime: '',
            marketCap: '',
            commentsCount: '',
            devWallet: '',
            devHolding: '',
            twitterLink: '',
            telegramLink: '',
            websiteLink: '',
            livePriceLink: '',
            priceWave5m: '',
            priceWave1h: '',
            priceWave6h: '',
            tXs: '',
            vol: '',
            liq: '',
            holder: '',
            open: '',
            noMint: '',
            blackList: '',
            burnt: '',
            top10: '',
            devAddr: '',
            devInfo: '',
            devBurn: '',
            advInfo: '',
        })



// 验证规则
const rule = reactive({
               id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCoinsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteCoinsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteCoinsByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCoinsFunc = async(row) => {
    const res = await findCoins({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCoinsFunc = async (row) => {
    const res = await deleteCoins({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: new Date(),
        tokenName: '',
        contractAddress: '',
        bondingCurve: '',
        progress: '',
        balance: '',
        createdTime: '',
        marketCap: '',
        commentsCount: '',
        devWallet: '',
        devHolding: '',
        twitterLink: '',
        telegramLink: '',
        websiteLink: '',
        livePriceLink: '',
        priceWave5m: '',
        priceWave1h: '',
        priceWave6h: '',
        tXs: '',
        vol: '',
        liq: '',
        holder: '',
        open: '',
        noMint: '',
        blackList: '',
        burnt: '',
        top10: '',
        devAddr: '',
        devInfo: '',
        devBurn: '',
        advInfo: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCoins(formData.value)
                  break
                case 'update':
                  res = await updateCoins(formData.value)
                  break
                default:
                  res = await createCoins(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCoins({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>

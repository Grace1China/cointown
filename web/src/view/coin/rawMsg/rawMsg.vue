<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-input v-model="searchInfo.ChatId" type="string" placeholder="频道ID"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input v-model="searchInfo.id" type="string" placeholder="消息id"></el-input>
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true"
            v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
        <el-button icon="done" style="margin-left: 10px;" 
          @click="onUpdateRawMsgStatusAll">标为已处理</el-button>
        <el-button icon="delete" style="margin-left: 10px;" 
          @click="onDelete7">清除7天前的消息</el-button>

      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="180">
        </el-table-column>
        <el-table-column align="left" label="新消息" prop="IsNew" width="120" />


        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="时间" prop="currentTime" width="180">
          <template #default="scope">{{ formatDate(scope.row.currentTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="对话Id" prop="chatId" width="120" />
        <el-table-column align="left" label="话题Id" prop="topicId" width="120" />
        <!-- <el-table-column align="left" label="消息体" prop="messageText" width="120" /> -->
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateRawMsgFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="addRawMsgTemplate(scope.row)">添加频道模板</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
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
        <el-form-item label="时间:" prop="currentTime">
          <el-date-picker v-model="formData.currentTime" type="date" style="width:100%" placeholder="选择日期"
            :clearable="true" />
        </el-form-item>
        <el-form-item label="对话Id:" prop="chatId">
          <el-input v-model="formData.chatId" :clearable="true" placeholder="请输入对话Id" />
        </el-form-item>
        <el-form-item label="话题Id:" prop="topicId">
          <el-input v-model="formData.topicId" :clearable="true" placeholder="请输入话题Id" />
        </el-form-item>
        <el-form-item label="消息体:" prop="messageText">
          <el-input type="textarea" rows="5" v-model="formData.messageText" :clearable="true" placeholder="请输入消息体" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="dialogFormVisibleT" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type}}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="channelFormData" label-position="top" ref="elChannelFormRef" :rules="rule" label-width="80px">
        <el-form-item label="时间:" prop="currentTime">
          <el-date-picker v-model="channelFormData.currentTime" type="date" style="width:100%" placeholder=""
            :clearable="true" />
        </el-form-item>
        <el-form-item label="对话Id:" prop="chatId">
          <el-input v-model="channelFormData.chatId" :clearable="true" placeholder="" />
        </el-form-item>
        <el-form-item label="话题Id:" prop="topicId">
          <el-input v-model="channelFormData.topicId" :clearable="true" placeholder="" />
        </el-form-item>
        <el-form-item label="消息体:" prop="lastMessageText">
          <el-input type="textarea" :rows="5" v-model="channelFormData.lastMessageText" :clearable="true"
            placeholder="" />
        </el-form-item>
        <div style="color: forestgreen;">
          {{ "TokenName ,ContractAddress ,BondingCurve ,Progress ,Balance ,Price ,CreatedTime ,MarketCap ,CommentsCount,DevWallet ,DevHolding ,TwitterLink ,TelegramLink ,WebsiteLink ,LivePriceLink ,PriceWave5m ,PriceWave1h,PriceWave6h ,TradeVol1m ,TradeVol5m ,TradeVol1h ,BuyVol1m ,BuyVol5m ,BuyVol1h ,SaleVol1m ,SaleVol5m ,SaleVol1h,TXs ,Vol ,Liq ,Holder ,Open ,NoMint ,BlackList ,Burnt ,Top10 ,DevAddr ,DevInfo ,DevBurn ,AdvInfo ,Circulating,CreatorBal ,CreatorBalPer ,MintCoins ,RecommendAddr ,MsgId ,PriceTest1 ,PriceTest2 ,PriceTest3 ,PriceTestWave1,PriceTestWave2 ,PriceTestWave3 ,PriceTestTime1 ,PriceTestTime2 ,PriceTestTime3 ,SmartBuyer ,ChatId" }}

          <!-- {
            "TokenName":"\\*\\*Name\\*\\*:([a-zA-Z]+)",
            "ContractAddress":"<addr>(.*?)<addr>",
            "MarketCap":"\\*\\*Market Cap\\*\\*:(\\$\\d+)"    
          } -->
        </div>
        <el-form-item label="消息体模板:" prop="template">
          <el-input type="textarea" :rows="5" v-model="channelFormData.template" @input="onTempChange" :clearable="true"
            placeholder="消息体模板" />
        </el-form-item>
        <el-form-item label="匹配结果:" prop="result">
          <el-input type="textarea" :rows="5" v-model="channelFormData.result" :clearable="true" placeholder="自动匹配结果" />
        </el-form-item>
      </el-form>
      
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions column="1" border>
        <el-descriptions-item label="时间">
          {{ detailFrom.currentTime }}
        </el-descriptions-item>
        <el-descriptions-item label="对话Id">
          {{ detailFrom.chatId }}
        </el-descriptions-item>
        <el-descriptions-item label="话题Id">
          {{ detailFrom.topicId }}
        </el-descriptions-item>
        <el-descriptions-item label="消息体">
          {{ detailFrom.messageText }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
    
  </div>
</template>

<script setup>
  import {
    createRawMsg,
    deleteRawMsg,
    deleteRawMsg7D,
    UpdateRawMsgStatusAll,
    deleteRawMsgByIds,
    updateRawMsg,
    findRawMsg,
    getRawMsgList
  } from '@/api/coin/rawMsg'

  import {
    createChannel
  } from '@/api/coin/channel'

  // web/src/api/coin/channel.js

  // 全量引入格式化工具 请按需保留
  import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive ,onBeforeMount} from 'vue'

  defineOptions({
    name: 'RawMsg'
  })

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    currentTime: new Date(),
    chatId: '',
    topicId: '',
    messageText: '',
  })

  // 自动化生成的字典（可能为空）以及字段
  const channelFormData = ref({
    currentTime: new Date(),
    chatId: '',
    topicId: '',
    lastMessageText: '',
    template: ''
  })


// 获取 URL 中的 type 参数
function getHashParameter(name) {
  const hash = window.location.hash.slice(1); // 去掉 '#' 符号
  const urlParams = new URLSearchParams(hash.includes('?') ? hash.split('?')[1] : '');
  return urlParams.get(name);
}


// 在组件挂载前获取 URL 参数，并根据参数设置 showControl 的值
onBeforeMount(async () => {
  const ID = getHashParameter('ID');
  console.log('onBeforeMount',ID)
  if (ID) {
    // 更新行
    const res = await findRawMsg({ ID })
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }
});



  // 验证规则
  const rule = reactive({
  })

  const searchRule = reactive({
    createdAt: [
      {
        validator: (rule, value, callback) => {
          if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
            callback(new Error('请填写结束日期'))
          } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
            callback(new Error('请填写开始日期'))
          } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
            callback(new Error('开始日期应当早于结束日期'))
          } else {
            callback()
          }
        }, trigger: 'change'
      }
    ],
  })

  const elFormRef = ref()
  const elChannelFormRef = ref()

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
    elSearchFormRef.value?.validate(async (valid) => {
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
  const getTableData = async () => {
    const table = await getRawMsgList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  const setOptions = async () => {
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
      deleteRawMsgFunc(row)
    })
  }

   // 删除7天前的消息
  const onDelete7 = () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async  () => {
      const res =  await deleteRawMsg7D()
      console.log('onDelete7',res)
      if(res.code==0){
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

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
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
      const res = await deleteRawMsgByIds({ IDs })
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
  const updateRawMsgFunc = async (row) => {
    const res = await findRawMsg({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }
  const addRawMsgTemplate = async (row) => {
    const res = await findRawMsg({ ID: row.ID })
    type.value = 'addTemplate'
    if (res.code === 0) {
      const { currentTime, chatId, topicId, messageText } = res.data
      channelFormData.value = {
        currentTime, chatId, topicId, lastMessageText: messageText,
        template: `{
        "TokenName":"",
        "ContractAddress":""
      }`

        , result: ''
      }
      dialogFormVisibleT.value = true
    }
  }


  // 删除行
  const deleteRawMsgFunc = async (row) => {
    const res = await deleteRawMsg({ ID: row.ID })
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

  // const deleteRawMsgFunc7D = async () => {
  //   const res = await deleteRawMsg7D()
  //   if (res.code === 0) {
  //     ElMessage({
  //       type: 'success',
  //       message: '删除成功'
  //     })
  //     // if (tableData.value.length === 1 && page.value > 1) {
  //     //   page.value--
  //     // }
  //     getTableData()
  //   }
  // }

  const onUpdateRawMsgStatusAll = async () => {
    const res = await UpdateRawMsgStatusAll()
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      // if (tableData.value.length === 1 && page.value > 1) {
      //   page.value--
      // }
      getTableData()
    }
  }

  

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 弹窗控制标记
  const dialogFormVisibleT = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    if (dialogFormVisible.value) {
      dialogFormVisible.value = false
    }

    if (dialogFormVisibleT.value) {
      dialogFormVisibleT.value = false
    }
    formData.value = {
      currentTime: new Date(),
      chatId: '',
      topicId: '',
      messageText: '',
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    if (dialogFormVisible.value) {
      elFormRef.value?.validate(async (valid) => {
        if (!valid) return
        let res
        switch (type.value) {
          case 'create':
            res = await createRawMsg(formData.value)
            break
          case 'update':
            res = await updateRawMsg(formData.value)
            break
          default:
            res = await createRawMsg(formData.value)
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

    if (dialogFormVisibleT.value) {
      elChannelFormRef.value?.validate(async (valid) => {
        if (!valid) return
        let res
        res = await createChannel(channelFormData.value)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '创建/更改成功'
          })
          // closeDialog()
          // getTableData()
        }
      })
    }
  }

  const onTempChange = (value) => {
    console.log('onTempChange', value)
    value = value.replace(/\n/g, '')
    console.log('onTempChange1', value)
    const msgtxt = channelFormData.value.lastMessageText.replaceAll("`", "<addr>")

    const tpl = JSON.parse(value)
    console.log('tpl.TokenName,tpl.ContractAddress', tpl.TokenName, tpl.ContractAddress)
    channelFormData.value.result = ''
    for (const key of Object.keys(tpl)) {
      let regex = new RegExp(tpl[key])//g; // 定义一个全局搜索模式
      let matches = msgtxt.match(regex); // 使用 match 方法找到所有匹配项
      console.log(matches); // 输出 ["find", "found", "finding"]
      channelFormData.value.result = channelFormData.value.result + '\n' + `${key}:${matches[1]}`
    }
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
    const res = await findRawMsg({ ID: row.ID })
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
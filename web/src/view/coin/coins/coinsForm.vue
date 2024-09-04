<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id:" prop="id">
          <el-date-picker v-model="formData.id" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="币名:" prop="tokenName">
          <el-input v-model="formData.tokenName" :clearable="true"  placeholder="请输入币名" />
       </el-form-item>
        <el-form-item label="合约地址:" prop="contractAddress">
          <el-input v-model="formData.contractAddress" :clearable="true"  placeholder="请输入合约地址" />
       </el-form-item>
        <el-form-item label="Bond曲线:" prop="bondingCurve">
          <el-input v-model="formData.bondingCurve" :clearable="true"  placeholder="请输入Bond曲线" />
       </el-form-item>
        <el-form-item label="progress:" prop="progress">
          <el-input v-model="formData.progress" :clearable="true"  placeholder="请输入progress" />
       </el-form-item>
        <el-form-item label="余额:" prop="balance">
          <el-input v-model="formData.balance" :clearable="true"  placeholder="请输入余额" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdTime">
          <el-input v-model="formData.createdTime" :clearable="true"  placeholder="请输入创建时间" />
       </el-form-item>
        <el-form-item label="市场容量:" prop="marketCap">
          <el-input v-model="formData.marketCap" :clearable="true"  placeholder="请输入市场容量" />
       </el-form-item>
        <el-form-item label="评价数量:" prop="commentsCount">
          <el-input v-model="formData.commentsCount" :clearable="true"  placeholder="请输入评价数量" />
       </el-form-item>
        <el-form-item label="开发钱包:" prop="devWallet">
          <el-input v-model="formData.devWallet" :clearable="true"  placeholder="请输入开发钱包" />
       </el-form-item>
        <el-form-item label="开发钱包:" prop="devHolding">
          <el-input v-model="formData.devHolding" :clearable="true"  placeholder="请输入开发钱包" />
       </el-form-item>
        <el-form-item label="X链接:" prop="twitterLink">
          <el-input v-model="formData.twitterLink" :clearable="true"  placeholder="请输入X链接" />
       </el-form-item>
        <el-form-item label="T链接:" prop="telegramLink">
          <el-input v-model="formData.telegramLink" :clearable="true"  placeholder="请输入T链接" />
       </el-form-item>
        <el-form-item label="Web链接:" prop="websiteLink">
          <el-input v-model="formData.websiteLink" :clearable="true"  placeholder="请输入Web链接" />
       </el-form-item>
        <el-form-item label="Price链接:" prop="livePriceLink">
          <el-input v-model="formData.livePriceLink" :clearable="true"  placeholder="请输入Price链接" />
       </el-form-item>
        <el-form-item label="PriceWave5m:" prop="priceWave5m">
          <el-input v-model="formData.priceWave5m" :clearable="true"  placeholder="请输入PriceWave5m" />
       </el-form-item>
        <el-form-item label="PriceWave1h:" prop="priceWave1h">
          <el-input v-model="formData.priceWave1h" :clearable="true"  placeholder="请输入PriceWave1h" />
       </el-form-item>
        <el-form-item label="PriceWave6h:" prop="priceWave6h">
          <el-input v-model="formData.priceWave6h" :clearable="true"  placeholder="请输入PriceWave6h" />
       </el-form-item>
        <el-form-item label="TXs:" prop="tXs">
          <el-input v-model="formData.tXs" :clearable="true"  placeholder="请输入TXs" />
       </el-form-item>
        <el-form-item label="vol:" prop="vol">
          <el-input v-model="formData.vol" :clearable="true"  placeholder="请输入vol" />
       </el-form-item>
        <el-form-item label="Liq:" prop="liq">
          <el-input v-model="formData.liq" :clearable="true"  placeholder="请输入Liq" />
       </el-form-item>
        <el-form-item label="Holder:" prop="holder">
          <el-input v-model="formData.holder" :clearable="true"  placeholder="请输入Holder" />
       </el-form-item>
        <el-form-item label="Open:" prop="open">
          <el-input v-model="formData.open" :clearable="true"  placeholder="请输入Open" />
       </el-form-item>
        <el-form-item label="NoMint:" prop="noMint">
          <el-input v-model="formData.noMint" :clearable="true"  placeholder="请输入NoMint" />
       </el-form-item>
        <el-form-item label="BlackList:" prop="blackList">
          <el-input v-model="formData.blackList" :clearable="true"  placeholder="请输入BlackList" />
       </el-form-item>
        <el-form-item label="Burnt:" prop="burnt">
          <el-input v-model="formData.burnt" :clearable="true"  placeholder="请输入Burnt" />
       </el-form-item>
        <el-form-item label="Top10:" prop="top10">
          <el-input v-model="formData.top10" :clearable="true"  placeholder="请输入Top10" />
       </el-form-item>
        <el-form-item label="DevAddr:" prop="devAddr">
          <el-input v-model="formData.devAddr" :clearable="true"  placeholder="请输入DevAddr" />
       </el-form-item>
        <el-form-item label="DevInfo:" prop="devInfo">
          <el-input v-model="formData.devInfo" :clearable="true"  placeholder="请输入DevInfo" />
       </el-form-item>
        <el-form-item label="DevBurn:" prop="devBurn">
          <el-input v-model="formData.devBurn" :clearable="true"  placeholder="请输入DevBurn" />
       </el-form-item>
        <el-form-item label="AdvInfo:" prop="advInfo">
          <el-input v-model="formData.advInfo" :clearable="true"  placeholder="请输入AdvInfo" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCoins,
  updateCoins,
  findCoins
} from '@/api/coin/coins'

defineOptions({
    name: 'CoinsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCoins({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>

<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="时间:" prop="currentTime">
          <el-date-picker v-model="formData.currentTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="对话Id:" prop="chatId">
          <el-input v-model="formData.chatId" :clearable="true"  placeholder="请输入对话Id" />
       </el-form-item>
        <el-form-item label="话题Id:" prop="topicId">
          <el-input v-model="formData.topicId" :clearable="true"  placeholder="请输入话题Id" />
       </el-form-item>
        <el-form-item label="消息体:" prop="lastMessageText">
          <el-input v-model="formData.lastMessageText" :clearable="true"  placeholder="请输入消息体" />
       </el-form-item>
        <el-form-item label="消息体模版:" prop="template">
          <el-input v-model="formData.template" :clearable="true"  placeholder="请输入消息体模版" />
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
  createChannel,
  updateChannel,
  findChannel
} from '@/api/coin/channel'

defineOptions({
    name: 'ChannelForm'
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
            currentTime: new Date(),
            chatId: '',
            topicId: '',
            lastMessageText: '',
            template: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findChannel({ ID: route.query.id })
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
               res = await createChannel(formData.value)
               break
             case 'update':
               res = await updateChannel(formData.value)
               break
             default:
               res = await createChannel(formData.value)
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

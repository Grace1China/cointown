import service from '@/utils/request'

// @Tags Channel
// @Summary 创建频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Channel true "创建频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /channel/createChannel [post]
export const createChannel = (data) => {
  return service({
    url: '/channel/createChannel',
    method: 'post',
    data
  })
}

// @Tags Channel
// @Summary 删除频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Channel true "删除频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /channel/deleteChannel [delete]
export const deleteChannel = (params) => {
  return service({
    url: '/channel/deleteChannel',
    method: 'delete',
    params
  })
}

// @Tags Channel
// @Summary 批量删除频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /channel/deleteChannel [delete]
export const deleteChannelByIds = (params) => {
  return service({
    url: '/channel/deleteChannelByIds',
    method: 'delete',
    params
  })
}

// @Tags Channel
// @Summary 更新频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Channel true "更新频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /channel/updateChannel [put]
export const updateChannel = (data) => {
  return service({
    url: '/channel/updateChannel',
    method: 'put',
    data
  })
}

// @Tags Channel
// @Summary 用id查询频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Channel true "用id查询频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /channel/findChannel [get]
export const findChannel = (params) => {
  return service({
    url: '/channel/findChannel',
    method: 'get',
    params
  })
}

// @Tags Channel
// @Summary 分页获取频道列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取频道列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /channel/getChannelList [get]
export const getChannelList = (params) => {
  return service({
    url: '/channel/getChannelList',
    method: 'get',
    params
  })
}

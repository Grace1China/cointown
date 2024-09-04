import service from '@/utils/request'

// @Tags RawMsg
// @Summary 创建原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RawMsg true "创建原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rawmsg/createRawMsg [post]
export const createRawMsg = (data) => {
  return service({
    url: '/rawmsg/createRawMsg',
    method: 'post',
    data
  })
}

// @Tags RawMsg
// @Summary 删除原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RawMsg true "删除原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rawmsg/deleteRawMsg [delete]
export const deleteRawMsg = (params) => {
  return service({
    url: '/rawmsg/deleteRawMsg',
    method: 'delete',
    params
  })
}

// @Tags RawMsg
// @Summary 删除原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RawMsg true "删除原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rawmsg/deleteRawMsg [get]
export const deleteRawMsg7D = () => {
  return service({
    url: '/rawmsg/deleteRawMsg7D',
    method: 'post',
    
  })
}
// @Tags RawMsg
// @Summary 删除原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RawMsg true "删除原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rawmsg/UpdateRawMsgStatusAll [get]
export const UpdateRawMsgStatusAll = () => {
  return service({
    url: '/rawmsg/UpdateRawMsgStatusAll',
    method: 'post',
    
  })
}


// @Tags RawMsg
// @Summary 批量删除原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rawmsg/deleteRawMsg [delete]
export const deleteRawMsgByIds = (params) => {
  return service({
    url: '/rawmsg/deleteRawMsgByIds',
    method: 'delete',
    params
  })
}

// @Tags RawMsg
// @Summary 更新原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RawMsg true "更新原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rawmsg/updateRawMsg [put]
export const updateRawMsg = (data) => {
  return service({
    url: '/rawmsg/updateRawMsg',
    method: 'put',
    data
  })
}

// @Tags RawMsg
// @Summary 用id查询原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RawMsg true "用id查询原始消息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rawmsg/findRawMsg [get]
export const findRawMsg = (params) => {
  return service({
    url: '/rawmsg/findRawMsg',
    method: 'get',
    params
  })
}

// @Tags RawMsg
// @Summary 分页获取原始消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取原始消息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rawmsg/getRawMsgList [get]
export const getRawMsgList = (params) => {
  return service({
    url: '/rawmsg/getRawMsgList',
    method: 'get',
    params
  })
}

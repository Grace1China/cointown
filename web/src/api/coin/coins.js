import service from '@/utils/request'

// @Tags Coins
// @Summary 创建币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coins true "创建币"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /coins/createCoins [post]
export const createCoins = (data) => {
  return service({
    url: '/coins/createCoins',
    method: 'post',
    data
  })
}

// @Tags Coins
// @Summary 删除币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coins true "删除币"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coins/deleteCoins [delete]
export const deleteCoins = (params) => {
  return service({
    url: '/coins/deleteCoins',
    method: 'delete',
    params
  })
}

// @Tags Coins
// @Summary 批量删除币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除币"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coins/deleteCoins [delete]
export const deleteCoinsByIds = (params) => {
  return service({
    url: '/coins/deleteCoinsByIds',
    method: 'delete',
    params
  })
}

// @Tags Coins
// @Summary 更新币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Coins true "更新币"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /coins/updateCoins [put]
export const updateCoins = (data) => {
  return service({
    url: '/coins/updateCoins',
    method: 'put',
    data
  })
}

// @Tags Coins
// @Summary 用id查询币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Coins true "用id查询币"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /coins/findCoins [get]
export const findCoins = (params) => {
  return service({
    url: '/coins/findCoins',
    method: 'get',
    params
  })
}

// @Tags Coins
// @Summary 分页获取币列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取币列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coins/getCoinsList [get]
export const getCoinsList = (params) => {
  return service({
    url: '/coins/getCoinsList',
    method: 'get',
    params
  })
}

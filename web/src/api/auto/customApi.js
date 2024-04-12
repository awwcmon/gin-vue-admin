import service from '@/utils/request'

// @Tags CustomApi
// @Summary 创建CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CustomApi true "创建CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /customApi/createCustomApi [post]
export const createCustomApi = (data) => {
  return service({
    url: '/customApi/createCustomApi',
    method: 'post',
    data
  })
}

// @Tags CustomApi
// @Summary 删除CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CustomApi true "删除CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customApi/deleteCustomApi [delete]
export const deleteCustomApi = (params) => {
  return service({
    url: '/customApi/deleteCustomApi',
    method: 'delete',
    params
  })
}

// @Tags CustomApi
// @Summary 批量删除CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customApi/deleteCustomApi [delete]
export const deleteCustomApiByIds = (params) => {
  return service({
    url: '/customApi/deleteCustomApiByIds',
    method: 'delete',
    params
  })
}

// @Tags CustomApi
// @Summary 更新CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CustomApi true "更新CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customApi/updateCustomApi [put]
export const updateCustomApi = (data) => {
  return service({
    url: '/customApi/updateCustomApi',
    method: 'put',
    data
  })
}

// @Tags CustomApi
// @Summary 用id查询CustomApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CustomApi true "用id查询CustomApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customApi/findCustomApi [get]
export const findCustomApi = (params) => {
  return service({
    url: '/customApi/findCustomApi',
    method: 'get',
    params
  })
}

// @Tags CustomApi
// @Summary 分页获取CustomApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CustomApi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customApi/getCustomApiList [get]
export const getCustomApiList = (params) => {
  return service({
    url: '/customApi/getCustomApiList',
    method: 'get',
    params
  })
}

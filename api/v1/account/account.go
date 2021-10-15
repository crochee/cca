// Date: 2021/10/8

// Package account
package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"cca/pkg/resp"
	"cca/pkg/service/account"
)

// Register godoc
// swagger:operation POST /v1/account 账户 SAccountRegisterRequest
// ---
// summary: 注册账户
// description: 注册账户信息
// Consumes:
// - application/json
// produces:
// - application/json
// responses:
//   '200':
//     type: object
//     "$ref": "#/responses/SAccountRegisterResponseResult"
//   default:
//     type: object
//     "$ref": "#/responses/SResponseCode"
func Register(ctx *gin.Context) {
	var registerRequest account.CreateRequest
	if err := ctx.ShouldBindBodyWith(&registerRequest, binding.JSON); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	if err := account.ValidPassword(registerRequest.Password); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	response, err := account.Create(ctx.Request.Context(), &registerRequest)
	if err != nil {
		resp.Errors(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Retrieves godoc
// swagger:operation GET /v1/account 账户 SAccountRetrievesRequest
// ---
// summary: 查询账户
// description: 根据条件查询账户列表
// produces:
// - application/json
// responses:
//   '200':
//     type: object
//     "$ref": "#/responses/SAccountRetrieveResponses"
//   default:
//     type: object
//     "$ref": "#/responses/SResponseCode"
func Retrieves(ctx *gin.Context) {
	var retrieveRequest account.RetrievesRequest
	if err := ctx.BindQuery(&retrieveRequest); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	response, err := account.Retrieves(ctx.Request.Context(), &retrieveRequest)
	if err != nil {
		resp.Errors(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Update godoc
// swagger:operation PATCH /v1/account/{id} 账户 SAccountUpdateRequest
// ---
// summary: 编辑账户
// description: 编辑指定账户的信息
// Consumes:
// - application/json
// produces:
// - application/json
// responses:
//   '204':
//     type: object
//     "$ref": "#/responses/SNullResponse"
//   default:
//     type: object
//     "$ref": "#/responses/SResponseCode"
func Update(ctx *gin.Context) {
	var user account.User
	if err := ctx.BindUri(&user); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	var modifyRequest account.UpdateRequest
	if err := ctx.ShouldBindBodyWith(&modifyRequest, binding.JSON); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	if err := account.ValidPassword(modifyRequest.Password); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	if err := account.Update(ctx.Request.Context(), &user, &modifyRequest); err != nil {
		resp.Errors(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// Retrieve godoc
// swagger:operation GET /v1/account/{id} 账户 SAccountRetrieveRequest
// ---
// summary: 查询指定账户
// description: 查询指定账户的信息
// produces:
// - application/json
// responses:
//   '200':
//     type: object
//     "$ref": "#/responses/SAccountRetrieveResponse"
//   default:
//     type: object
//     "$ref": "#/responses/SResponseCode"
func Retrieve(ctx *gin.Context) {
	var user account.User
	if err := ctx.BindUri(&user); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	response, err := account.Retrieve(ctx.Request.Context(), &user)
	if err != nil {
		resp.Errors(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Delete godoc
// swagger:operation DELETE /v1/account/{id} 账户 SAccountDeleteRequest
// ---
// summary: 删除指定账户
// description: 删除指定账户信息
// produces:
// - application/json
// responses:
//   '204':
//     type: object
//     "$ref": "#/responses/SNullResponse"
//   default:
//     type: object
//     "$ref": "#/responses/SResponseCode"
func Delete(ctx *gin.Context) {
	var user account.User
	if err := ctx.BindUri(&user); err != nil {
		resp.ErrorParam(ctx, err)
		return
	}
	err := account.Delete(ctx.Request.Context(), &user)
	if err != nil {
		resp.Errors(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

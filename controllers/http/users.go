package http

import (
	"IRIS_WEB/errors"
	"IRIS_WEB/models/dto"
	"IRIS_WEB/services"
	"github.com/kataras/iris"
	"strconv"
)

// RESTFUL /path/1
func ActionCrudUser(ctx iris.Context) {
	var err error
	var user *dto.UserDTO
	var userID int

	switch ctx.Method() {
	case "GET":
		// 根据ID获取用户
		userID, err = strconv.Atoi(ctx.Params().Get("id")) // /path/1
		if err != nil {
			panic(err)
		}
		if user, err = services.FetchUserById(userID); err != nil {
			ctx.JSON(err)
			return
		}

		if user == nil {
			ctx.JSON(errors.NoDataError())
			return
		}

		ctx.JSON(iris.Map{"code": 1000, "data": user})
	case "POST":
		if err = ctx.ReadJSON(&user); err != nil {
			errors.ParamError("invalid json format")
		}

		if err = services.CreateUser(user); err != nil {
			ctx.JSON(err)
			return
		}
		ctx.JSON(iris.Map{"code": 1000, "data": true})
	case "PUT":
		if err = ctx.ReadJSON(&user); err != nil {
			errors.ParamError("invalid json format")
		}

		if err = services.UpdateUser(user); err != nil {
			ctx.JSON(err)
			return
		}
		ctx.JSON(iris.Map{"code": 1000, "data": true})
	case "DELETE":
		userID, err = strconv.Atoi(ctx.Params().Get("id"))
		if err != nil {
			panic(err)
		}
		if err = services.DeleteUser(userID); err != nil {
			ctx.JSON(err)
			return
		}

		ctx.JSON(iris.Map{"code": 1000, "data": true})

	}
}

// PARAMS /path?user_id=1
func ActionGetUser(ctx iris.Context) {
	var err error
	var user *dto.UserDTO
	var params dto.UserParamDTO

	// 绑定参数
	if err = params.Bind(ctx); err != nil {
		ctx.JSON(err)
		return
	}

	// 根据ID获取用户
	if user, err = services.FetchUserById(params.UserId); err != nil {
		ctx.JSON(err)
		return
	}

	if user == nil {
		ctx.JSON(errors.NoDataError())
		return
	}
	ctx.JSON(iris.Map{"code": 1000, "data": user})
}

func ActionGetAllUsers(ctx iris.Context) {
	var err error
	var users []*dto.UserDTO

	// 查询所有用户
	if users, err = services.FetchAllUsers(); err != nil {
		ctx.JSON(err)
		return
	}

	if len(users) == 0 {
		ctx.JSON(errors.NoDataError())
		return
	}

	ctx.JSON(iris.Map{"code": 1000, "data": users})
}

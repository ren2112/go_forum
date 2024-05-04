package middleware

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/vo"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func PostMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" {
			body, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				// 处理错误
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}

			// 打印请求体
			fmt.Println("Request Body:", string(body))

			DB := common.GetDB()
			var requestPost vo.CreatePostRequest
			if err := ctx.ShouldBind(&requestPost); err != nil {
				response.Fail(ctx, nil, "数据验证错误")
				ctx.Abort()
				return
			}
			var category model.Category
			targetId := requestPost.CategoryId
			DB.Model(&category).Where("id=?", targetId).First(&category)
			if category.Name == "" {
				response.Fail(ctx, nil, "不存在该目录")
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}

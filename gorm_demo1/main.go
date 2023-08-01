package main

import (
	"fmt"
	"gorm_demo1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 添加数据
func Add(c *gin.Context) {
	user := models.User{
		Username: "xiaoming",
		Id:       6,
		Age:      99,
	}
	result := models.DB.Create(&user)
	if result.RowsAffected > 1 {
		fmt.Println(user.Id)
	}
	fmt.Println(result.RowsAffected)
	fmt.Println(user.Id)
}

// 查找全部数据
func FindAll(c *gin.Context) {
	user := []models.User{}
	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  user,
	})
}

// 查找指定数据
func FindSome(c *gin.Context) {
	user := []models.User{}
	models.DB.Where("username=?", "wangwu").Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  user,
	})
}

// 修改数据
func Modify(c *gin.Context) {
	user := models.User{Id: 1}
	models.DB.Find(&user)
	user.Username = "xiaozhang"
	user.Age = 11
	models.DB.Save(&user)
}

// 删除数据
func Delete(c *gin.Context) {
	user := models.User{Id: 2}
	models.DB.Delete(&user)
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/add", Modify, func(c *gin.Context) {
		c.String(200, "添加成功")
	})
	r.GET("/findall", FindAll, func(c *gin.Context) {
		c.String(200, "检索到所有数据")
	})
	r.GET("/findsome", FindSome, func(c *gin.Context) {
		c.String(200, "检索到指定数据")
	})
	r.GET("/modify", Modify, func(c *gin.Context) {
		c.String(200, "修改成功")
	})
	r.GET("/delete", Delete, func(c *gin.Context) {
		c.String(200, "删除成功")
	})
	r.Run()
}

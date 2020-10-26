package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Show struct {
	Id           int64  `gorm:"primary_key;AUTO_INCREMENT"`
	SId          int64  `gorm:"type:varchar(50)" form:"sid"`
	CategoryName string `gorm:"type:varchar(50)" form:"cname"`
}

type ShopCar struct {
	Id  int64 `gorm:"primary_key;AUTO_INCREMENT"`
	SId int64 `gorm:"type:varchar(50)" form:"sid"`
}

func main() {

	//链接数据库
	conn := "zhoudi:zhoudi@/fdm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	//创建数据库
	//db.CreateTable(&Show{})
	//showData:=Show{SId: 111,CategoryName: "Test"}
	//db.Create(&showData)
	//find := db.Where("s_id = ?", 125554).Find(&Show{})
	//ret := find.Scan(&Show{})
	//fmt.Println(ret)

	//数据接口
	r := gin.Default()
	r.GET("/api/add", func(c *gin.Context) {
		Cors(c)
		var aShow Show
		if err2 := c.BindQuery(&aShow); err2 != nil {
			panic(err2)
			return
		}
		//创建记录并添加数据
		addData := Show{SId: aShow.SId, CategoryName: aShow.CategoryName}
		//查询数据是否存在
		selectData := Show{}
		db.Where("s_id = ?", aShow.SId).First(&selectData)
		// now use p
		//log.Println(selectData)
		if selectData.SId != 0 {
			//返回状态码
			c.JSON(200, gin.H{
				"message": "No",
				"sid":     aShow.SId,
				"cname":   aShow.CategoryName,
				"code":    0,
			})

		} else {
			db.Create(&addData)
			//返回状态码
			c.JSON(200, gin.H{
				"message": "OK",
				"sid":     aShow.SId,
				"cname":   aShow.CategoryName,
				"code":    1,
			})
		}

	})

	r.POST("/api/getlovelist", func(c *gin.Context) {
		Cors(c)
		var findAllDatas []Show
		db.Find(&findAllDatas)
		fmt.Println(findAllDatas)
		c.JSON(http.StatusOK, gin.H{
			"data": findAllDatas,
		})
	})

	r.GET("/api/removealove", func(c *gin.Context) {
		Cors(c)

		//传过来的ID
		S_Id := c.Query("s_id")
		var toFindAShow Show
		//数据库中查询
		db.Where("s_id = ?", S_Id).Find(&toFindAShow)

		code := -1
		mes := "数据未找到"
		if toFindAShow.SId != 0 {
			//查询找到了存在的数据
			code = 1
			mes = "OK"
			db.Where("s_id = ?", S_Id).Delete(&toFindAShow)
		}
		fmt.Println(toFindAShow)

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"mes":  mes,
		})
	})

	r.GET("/api/addtoshopcar", func(c *gin.Context) {

		//创建 数据库
		//db.CreateTable(&ShopCar{})

		Cors(c)
		var shopcar ShopCar

		if err2 := c.BindQuery(&shopcar); err2 != nil {
			panic(err2)
			return
		}

		addData := ShopCar{SId: shopcar.SId}
		fmt.Println(addData)
		db.Where("s_id = ?", shopcar.SId).Find(&shopcar)
		fmt.Println(shopcar)
		//db.NewRecord(shopcar) // => 主键为空返回`true`

		//db.Create(&shopcar)
		if shopcar.Id != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": shopcar,
				"mes":  "数据已存在",
			})
		} else {
			db.Create(&shopcar)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"data": shopcar,
				"mes":  "插入成功",
			})
		}
	})

	r.POST("/api/getshopcarlist", func(c *gin.Context) {
		Cors(c)
		var findAllDatas []ShopCar
		db.Find(&findAllDatas)
		fmt.Println(findAllDatas)
		c.JSON(http.StatusOK, gin.H{
			"data": findAllDatas,
		})
	})

	r.Run(":9090") // 监听并在 0.0.0.0:9090 上启动服务
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
}

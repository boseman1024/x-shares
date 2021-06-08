package api

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"path"
	"shares/conf"
	"shares/serializer"
	"shares/service"
	"strconv"
)

func ShareIndex(c *gin.Context) {
	var shareListService service.ShareListService
	if err := c.ShouldBindJSON(&shareListService); err == nil {
		res := shareListService.Index()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "分享服务初始化失败",
		})
	}
}
func ShareList(c *gin.Context) {
	var shareListService service.ShareListService
	if err := c.ShouldBindJSON(&shareListService); err == nil {
		res := shareListService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "分享服务初始化失败",
		})
	}
}
func ShareInfo(c *gin.Context) {
	var shareService service.ShareService
	if err := c.ShouldBindJSON(&shareService); err == nil {
		if share, err := shareService.Info(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: share,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "分享服务初始化失败",
		})
	}
}

func ShareAdd(c *gin.Context) {
	var conf conf.Conf
	filePath, confErr := conf.GetConf("filePath")
	if confErr != nil {
		c.JSON(200, serializer.Response{
			Code: 211,
			Msg:  "配置文件加载失败",
		})
		c.Abort()
	}
	fileUrl, confErr := conf.GetConf("fileUrl")
	if confErr != nil {
		c.JSON(200, serializer.Response{
			Code: 211,
			Msg:  "配置文件加载失败",
		})
		c.Abort()
	}

	file, _ := c.FormFile("file")
	userFilePath := filePath + c.PostForm("userId") + "\\"
	randStr := strconv.FormatUint(rand.Uint64(), 10) + "_"
	if _, err := os.Stat(userFilePath); err != nil {
		os.Mkdir(userFilePath, os.ModePerm)
	}
	err := c.SaveUploadedFile(file, userFilePath+randStr+file.Filename)
	if err != nil {
		c.JSON(200, serializer.Response{
			Code: 211,
			Msg:  "文件上传失败",
		})
		c.Abort()
	}

	ext := path.Ext(file.Filename)
	fileStream, fileStreamErr := file.Open()
	if fileStreamErr != nil {
		c.JSON(200, serializer.Response{
			Code: 211,
			Msg:  "文件加载失败",
		})
		c.Abort()
	}
	var img image.Image
	var imgErr error
	if ext == ".jpg" {
		img, imgErr = jpeg.Decode(fileStream)
	}
	if ext == ".png" {
		img, imgErr = png.Decode(fileStream)
	}
	if imgErr != nil {
		c.JSON(200, serializer.Response{
			Code: 211,
			Msg:  "文件转换失败",
		})
		c.Abort()
	}
	thumbnail := imaging.Fill(img, 600, 600, imaging.Center, imaging.Lanczos)
	thumbnailFile, thumbnailFileErr := os.Create(userFilePath + randStr + file.Filename + "_thumbnail.jpg")
	if thumbnailFileErr != nil {
		c.JSON(200, serializer.Response{
			Code: 211,
			Msg:  "文件转换失败",
		})
		c.Abort()
	}
	defer thumbnailFile.Close()
	jpeg.Encode(thumbnailFile, thumbnail, &jpeg.Options{Quality: 100})

	var shareService service.ShareService
	shareService.Img = fileUrl + c.PostForm("userId") + "/" + randStr + file.Filename
	userId, _ := strconv.ParseUint(c.PostForm("userId"), 10, 32)
	tags := c.PostForm("tags")
	categoryId, _ := strconv.ParseUint(c.PostForm("categoryId"), 10, 32)
	if share, err := shareService.Insert(uint(userId), tags, uint(categoryId)); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, serializer.Response{
			Code: 200,
			Data: share,
		})
	}
}
func ShareEdit(c *gin.Context) {
	var shareService service.ShareService
	if err := c.ShouldBindJSON(&shareService); err == nil {
		if share, err := shareService.Update(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: share,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 201,
			Msg:  "分享服务初始化失败",
		})
	}
}
func ShareDel(c *gin.Context) {
	var shareService service.ShareService
	if num, err := shareService.Delete(c.Param("id")); err != nil {
		c.JSON(200, err)
	} else {
		c.JSON(200, serializer.Response{
			Code: 200,
			Data: num,
		})
	}
}

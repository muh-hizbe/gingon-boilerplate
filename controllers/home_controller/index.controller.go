package home_controller

import (
	"gingon-boilerplate/constanta"
	"gingon-boilerplate/helpers"
	"gingon-boilerplate/utils"
	"image"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "bar",
	})
}

func StoreImageHandler(ctx *gin.Context) {
	imageHandler, _ := ctx.FormFile("image")
	if imageHandler == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "image is required",
		})
		return
	}

	extensionFile := filepath.Ext(imageHandler.Filename)
	imagename := utils.RandomFileName(extensionFile, "image")

	isImageValidated := utils.FileValidation(imageHandler, constanta.IMAGE_VALIDTION())
	if !isImageValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "image not valid",
		})
		return
	}

	// >> BEGIN SAVE IMAGE
	file, errFile := imageHandler.Open()
	helpers.PanicIfError(errFile)

	image, _, errDecode := image.Decode(file)
	helpers.PanicIfError(errDecode)

	utils.SaveImage(image, imageHandler, imagename)
	// END SAVE IMAGE <<

	ctx.JSON(200, gin.H{
		"message": "store an image successfully.",
	})
}

func StoreImage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "i don't know the image was stored on server or not!, cause it handled by middleware before.",
	})
}

func StoreFile(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "i don't know the file was stored on server or not!, cause it handled by middleware before.",
	})
}

func StoreFileHandler(ctx *gin.Context) {
	fileHandler, _ := ctx.FormFile("file")
	if fileHandler == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is required",
		})
		return
	}

	extensionFile := filepath.Ext(fileHandler.Filename)
	filename := utils.RandomFileName(extensionFile, "file")

	isFileValidated := utils.FileValidation(fileHandler, constanta.FILE_VALIDTION())
	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file not valid",
		})
		return
	}

	// >> BEGIN SAVE IMAGE
	utils.SaveFile(ctx, fileHandler, filename)
	// END SAVE IMAGE <<

	ctx.JSON(200, gin.H{
		"message": "store an image successfully.",
	})
}

package utils

import (
	"fmt"
	"gingon-boilerplate/constanta"
	"gingon-boilerplate/helpers"
	"image"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

const DefaultPathAssetImage = "./public/images/"
const DefaultPathAssetFile = "./public/files/"

func HandleFileUpload(ctx *gin.Context) {
	imageHandler, errIH := ctx.FormFile("image")
	fileHandler, errFH := ctx.FormFile("file")

	if errIH != nil {
		log.Println("Error image upload ==> ", errIH)
		log.Println("Will skip this upload image proccess ")
	}
	if errFH != nil {
		log.Println("Error file upload ==> ", errFH)
		log.Println("Will skip this upload file proccess ")
	}

	var imagename, filename string

	if imageHandler != nil {
		extensionFile := filepath.Ext(imageHandler.Filename)

		isImageValidated := FileValidation(imageHandler, constanta.IMAGE_VALIDTION())

		if isImageValidated {
			//	GENERATE NEW IMAGE NAME
			imagename = RandomFileName(extensionFile, "image")

			file, errFile := imageHandler.Open()
			helpers.PanicIfError(errFile)

			image, _, errDecode := image.Decode(file)
			helpers.PanicIfError(errDecode)

			statusSave := SaveImage(image, imageHandler, imagename)
			if statusSave {
				ctx.Set("imagename", imagename)
			} else {
				ctx.Set("imagename", "")
			}
		} else {
			log.Println("The provided file format is not allowed. Image format not allowed.")
		}
	} else {
		log.Println("No image stored in asset directory.")
	}

	if fileHandler != nil {
		extensionFile := filepath.Ext(imageHandler.Filename)

		//	GENERATE NEW FILE NAME
		filename = RandomFileName(extensionFile, "file")

		statusSave := SaveFile(ctx, fileHandler, filename)
		if statusSave {
			ctx.Set("filename", filename)
		} else {
			ctx.Set("filename", "")
		}
	} else {
		log.Println("No file stored in asset directory.")
	}

	ctx.Next()
}

func HandleRemoveImage(imagename string) error {
	err := os.Remove(DefaultPathAssetImage + imagename)
	if err != nil {
		log.Println("Failed remove file")
		return err
	}
	return nil
}

func HandleRemoveFile(filename string) error {
	err := os.Remove(DefaultPathAssetFile + filename)
	if err != nil {
		log.Println("Failed remove file")
		return err
	}
	return nil
}

func SaveImage(image image.Image, fileHeader *multipart.FileHeader, imagename string) bool {
	var errSave error
	if fileHeader.Size > 2048000 { // 2Mb
		srcImage := imaging.Resize(image, 800, 0, imaging.Box)
		log.Println("Image do resizing")
		// STORE IMAGE TO THE ASSET DIRECTORY
		errSave = imaging.Save(srcImage, fmt.Sprintf("%s%s", DefaultPathAssetImage, imagename))
	} else {
		// STORE IMAGE TO THE ASSET DIRECTORY Use Imaging Library
		errSave = imaging.Save(image, fmt.Sprintf("%s%s", DefaultPathAssetImage, imagename))
		log.Println("Image skip resizing")
	}

	//// STORE IMAGE TO THE ASSET DIRECTORY
	//errSave := imaging.Save(srcImage, fmt.Sprintf("%s%s",DefaultPathAssetImage, imagename))

	// Use Native Function
	//errSave := ctx.SaveUploadedFile(imageHandler, fmt.Sprintf("%s%s",DefaultPathAssetImage, imagename))

	if errSave != nil {
		helpers.LogIfError("STORE IMAGE", errSave)
		return true
	} else {
		log.Println("Successfully store image to the assets directory")
		return false
	}
}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {
	errSave := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("%s%s", DefaultPathAssetFile, filename))

	if errSave != nil {
		helpers.LogIfError("STORE FILE", errSave)
		return true
	} else {
		log.Println("Successfully store file to the assets directory")
		return false
	}
}

func FileValidation(fileHeader *multipart.FileHeader, imageType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	result := false

	for _, typeImage := range imageType {
		if contentType == typeImage {
			result = true
			break
		}
	}

	return result
}

func RandomFileName(extensionFile string, prefix ...string) string {
	if prefix[0] == "" {
		prefix[0] = "file"
	}
	replacer := strings.NewReplacer(":", "", "/", "", "\\", "", "*", "", "<", "", ">", "", "|", "", " ", "")

	//	GENERATE TIME
	currentTime := time.Now().Format(constanta.TIME_FORMAT)

	//	GENERATE STRING AFTER REPLACED USING REPLACER
	stringReplaced := replacer.Replace(currentTime)

	// CREATE FILENAME
	filename := fmt.Sprintf("%s-%s-%s%s", prefix[0], stringReplaced, helpers.RandomString(4), extensionFile)
	return filename
}
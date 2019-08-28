package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type StatusData struct {
	FileName string `json:"file_name"`
	Status   int    `json:"status"`
}

func BadRequestError(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, err.Error())
	log.Fatalln(err.Error())
}

func UploadPicture(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		BadRequestError(err, context)
	}
	fileName := filepath.Base(header.Filename)
	_, _, err = image.DecodeConfig(buf)
	if err != nil {
		BadRequestError(err, context)
	}
	now := time.Now().Format("20060102150405")
	err = context.SaveUploadedFile(header, "./uploads/"+now+"_"+fileName)
	if err != nil {
		BadRequestError(err, context)
	}
	data := StatusData{Status: http.StatusOK, FileName: now + fileName}
	status, _ := json.Marshal(data)
	context.JSON(http.StatusOK, string(status))
}

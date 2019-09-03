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
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusBadRequest, err.Error())
	log.Fatalln(err.Error())
}

func UploadPicture(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	file, header, err := ctx.Request.FormFile("file")
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		BadRequestError(err, ctx)
	}
	fileName := filepath.Base(header.Filename)
	_, _, err = image.DecodeConfig(buf)
	if err != nil {
		BadRequestError(err, ctx)
	}
	now := time.Now().Format("20060102150405")
	err = ctx.SaveUploadedFile(header, "./uploads/"+now+"_"+fileName)
	if err != nil {
		BadRequestError(err, ctx)
	}
	data := StatusData{Status: http.StatusOK, FileName: "/uploads/" + now + fileName}
	status, _ := json.Marshal(data)
	ctx.JSON(http.StatusOK, string(status))
}

package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	fmt.Printf("Upload File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("Header: %+v\n", file.Header)

	dst, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		return err
	}
	defer dst.Close()

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println(err)
	}
	dst.Write(fileBytes)
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
}

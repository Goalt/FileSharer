package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/labstack/echo/v4"
)

type File struct {
	FileName string `json:"file_name"`
	Data     string `json:"data"`
}

type Token struct {
	Token string `json:"token_id"`
}

type TokenId struct {
	TokenId string `query:"token_id"`
}

func main() {
	e := echo.New()

	e.Static("/imgs", "html/imgs/")
	e.File("/style.css", "html/style.css")
	e.File("/", "html/index.html")
	e.File("/script.js", "html/script.js")
	e.File("/jquery-3.6.0.min.js", "html/jquery-3.6.0.min.js")

	e.POST("/upload", upload)
	e.GET("/download", download)

	_ = e.Start(":8080")
}

func download(c echo.Context) error {
	tokenId := new(TokenId)
	if err := c.Bind(tokenId); err != nil {
		return c.JSON(404, "")
	}

	file := File{
		FileName: "test.txt",
		Data:     "12345",
	}

	return c.JSON(200, file)
}

func upload(c echo.Context) error {
	mediaType, params, err := mime.ParseMediaType(c.Request().Header.Get("Content-Type"))
	if err != nil {
		fmt.Print(err)
		return c.String(500, "")
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(c.Request().Body, params["boundary"])
		for {
			p, err := mr.NextPart()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
				c.String(500, "")
			}

			slurp, err := io.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Part %q: %q\n", p.Header.Get("filename"), slurp)
		}
	}

	// body := c.Request().Body

	// buf := new(strings.Builder)
	// _, _ = io.Copy(buf, body)

	// bufS := buf.String()

	// fmt.Println(bufS)

	// r.ParseMultipartForm(100)
	// mForm := r.MultipartForm

	// for k, _ := range mForm.File {
	//  // k is the key of file part
	//  file, fileHeader, err := r.FormFile(k)
	//  if err != nil {
	//   fmt.Println("inovke FormFile error:", err)
	//   return
	//  }
	//  defer file.Close()
	//  fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
	//   fileHeader.Filename, fileHeader.Size, fileHeader.Header)

	token := Token{
		Token: "rsdlfskdlfmsldkm",
	}

	return c.JSON(200, token)
}

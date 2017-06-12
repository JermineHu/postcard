
package main

import (
	"image"
	"fmt"
	"github.com/fogleman/gg"
	"net/http"
	"log"
	"github.com/disintegration/imaging"
	"encoding/base64"
	"image/png"
	"bytes"
	"os"
)

func LoadImageByUrl(url string) (img image.Image, ext string, err error) {

	ret, err := http.Get(url)
	if err != nil {
		log.Println(url)
		status := map[string]string{}
		status["status"] = "400"
		status["url"] = url
		panic(status)
	}

	return image.Decode(ret.Body)
}

func GeneratorForPostcard(imageUrl, text, logPath, themLogo, fontFamily string) *gg.Context {

	im, _, err := LoadImageByUrl(imageUrl)

	cx := gg.NewContext(976, 1456)
	cx.SetRGB(1, 1, 1)
	cx.Clear()

	im = imaging.Resize(im, 830, 830, imaging.NearestNeighbor)// 重制图片的大小

	cx.DrawImageAnchored(im, 976 / 2, 1456 / 3, 0.5, 0.5)

	im2, err := gg.LoadImage(logPath) // 明信片上的logo
	if err != nil {
		panic(err)
	}

	cx.DrawImageAnchored(im2, 976 / 2, 940, 2.7, 0.2)

	im3, err := gg.LoadImage(themLogo) // 主题文字
	if err != nil {
		panic(err)
	}

	im3 = imaging.Resize(im3, 266, 0, imaging.NearestNeighbor)

	cx.DrawImageAnchored(im3, 976 / 2 + 60, 1380, 0.8, 0.2)

	cx.SetRGB(0, 0, 0)
	if err := cx.LoadFontFace(fontFamily, 24); err != nil {
		panic(err)
	}

	x := float64(976 - 100)
	y := 1000
	yi := 0
	ci := 0

	for _, v := range text {

		if ci % 9 == 0 {
			x = float64(x - 30)
			y = 1000
			yi = 0
		}

		cx.DrawStringWrapped(fmt.Sprintf("%c", v), x, float64(y + (yi * 40)), 1, 1, 16, 24, gg.AlignLeft)

		yi++
		ci++

	}

	return cx

}


//ToSaveImageForHtml  to convert base64 from image.Image
func ToSaveImageForHtml(pathAndName string, ctx *gg.Context) (imagePath string) {

	//  ctx.SavePNG(pathAndName) // Just only need one code if you want to save image

	buf := new(bytes.Buffer)
	err := png.Encode(buf, ctx.Image())
	img_bytes := buf.Bytes()

	if err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(img_bytes)

	f, err := os.Create(pathAndName)

	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("<html><body><img src=\"data:image/png;base64," + base64Str + "\" /></body></html>")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	imagePath=pathAndName

	return imagePath

}

func main() {

	imageUrl := "http://img.tusoapp.com/bb923c29-0d95-449a-a29c-1eaae1ec07ef.jpg"

	text := "123天生丽质难自弃，456天生丽质难自弃！天生丽质难自弃，天生丽质难自弃！天生丽质难自弃，天生丽质难自弃！天生丽质难自弃，天生丽质难自弃！233!"

	logPath := "./logo.png"

	themLogo := "./transformers.png"
	fontFamily := "/Library/Fonts/Arial Unicode.ttf"

	ctx:=GeneratorForPostcard(imageUrl, text, logPath, themLogo, fontFamily)  // to make postcard by url

	saveName:="./index-base-64-test.html"

	pathName:=ToSaveImageForHtml(saveName,ctx) // save image file

	fmt.Println("--->",pathName)


}




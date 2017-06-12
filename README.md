# postcard

#### To make a postcard from remote url or local path

## How to use？


  imageUrl := "http://img.tusoapp.com/bb923c29-0d95-449a-a29c-1eaae1ec07ef.jpg"

	text := "123天生丽质难自弃，456天生丽质难自弃！天生丽质难自弃，天生丽质难自弃！天生丽质难自弃，天生丽质难自弃！天生丽质难自弃，天生丽质难自弃！233!"

	logoPath := "./logo.png"

	themeLogo := "./transformers.png"
	fontFamily := "/Library/Fonts/Arial Unicode.ttf"

	ctx:=GeneratorForPostcard(imageUrl, text, logoPath, themeLogo, fontFamily)  // to make postcard by url

	saveName:="./index-base-64-test.html"

	pathName:=ToSaveImageForHtml(saveName,ctx) // save image file

	fmt.Println("--->",pathName)
  

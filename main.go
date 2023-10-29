package main

import (
	"MaxBitTools/page"
	"MaxBitTools/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	//解决中文乱码
	myApp.Settings().SetTheme(&theme.MyTheme{})
	myWindow := myApp.NewWindow("MaxBit Tools")
	myWindow.Resize(fyne.NewSize(800, 580))

	tabs := container.NewAppTabs(
		container.NewTabItem("About", page.About(myWindow)),
		container.NewTabItem("Base64", page.Base64Page(myWindow)),
		container.NewTabItem("JsonFormat", page.Jsonformat(myWindow)),
		container.NewTabItem("Json2Yaml", page.Json2yaml(myWindow)),
		container.NewTabItem("Jwt解析", page.JwtParsePage(myWindow)),
		container.NewTabItem("时间戳", page.TimePage(myWindow)),
		container.NewTabItem("UrlEnCode", page.UrlEncodePage(myWindow)),
		container.NewTabItem("Unicode", page.UnicodePage(myWindow)),
		container.NewTabItem("公约数", page.GcdPage(myWindow)),
		container.NewTabItem("Wallpaper", page.Wallpaper(myWindow)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
	//page.RequestImg()
}

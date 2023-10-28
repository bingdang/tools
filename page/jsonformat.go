package page

import (
	"bytes"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Jsonformat(w fyne.Window) fyne.CanvasObject {

	input := widget.NewMultiLineEntry()
	show := widget.NewMultiLineEntry()

	input.Wrapping = fyne.TextWrapBreak
	show.Wrapping = fyne.TextWrapBreak

	input.SetPlaceHolder("输入json")
	input.OnChanged = func(s string) {
		var obj bytes.Buffer
		_ = json.Indent(&obj, []byte(s), "", "    ")
		show.SetText(obj.String())
	}

	input.SetText(`{
	"object":{
		"name":"xxx",
		"sex":1,
		"success":true
	},
	"list":[
		1,2,3,4
	],
	"listObject":[
		{"name":"name1","age":12},
		{"name":"name1","age":12},
		{"name":"name1","age":12},
		{"name":"name1","age":12}
	]
}`)
	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(2, input, show),
	)

	return container.NewStack(rows)

}

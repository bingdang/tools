package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func About(w fyne.Window) fyne.CanvasObject {
	text := `
# About
## 这是什么？
这是一个小工具合集，作者为[MaxBit](https://cakepanit.com)站长。

Power By 饼铛(cheng)
`
	about := widget.NewRichTextFromMarkdown(text)
	about.Wrapping = fyne.TextWrapWord
	//About.Scroll = container.ScrollBoth

	for i := range about.Segments {
		if seg, ok := about.Segments[i].(*widget.TextSegment); ok {
			seg.Style.Alignment = fyne.TextAlignCenter
		}
		if seg, ok := about.Segments[i].(*widget.HyperlinkSegment); ok {
			seg.Alignment = fyne.TextAlignCenter
		}
	}
	return container.NewVBox(
		about,
	)
}

package page

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"io"
	"log"
	"net/http"
	"sync"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// 请求参数
type request struct {
	format string
	idx    string
	n      string
	mkt    string
}

// 响应接收
type AutoGenerated struct {
	Images   []Images `json:"images"`
	Tooltips Tooltips `json:"tooltips"`
}
type Images struct {
	Startdate     string        `json:"startdate"`
	Fullstartdate string        `json:"fullstartdate"`
	Enddate       string        `json:"enddate"`
	URL           string        `json:"url"`
	Urlbase       string        `json:"urlbase"`
	Copyright     string        `json:"copyright"`
	Copyrightlink string        `json:"copyrightlink"`
	Title         string        `json:"title"`
	Quiz          string        `json:"quiz"`
	Wp            bool          `json:"wp"`
	Hsh           string        `json:"hsh"`
	Drk           int           `json:"drk"`
	Top           int           `json:"top"`
	Bot           int           `json:"bot"`
	Hs            []interface{} `json:"hs"`
}
type Tooltips struct {
	Loading  string `json:"loading"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	Walle    string `json:"walle"`
	Walls    string `json:"walls"`
}

func NewRequest(Format, Idx, N, Mkt string) *request {
	return &request{
		format: Format,
		idx:    Idx,
		n:      N,
		mkt:    Mkt,
	}
}

func RequestImgURL(Format, Idx, N, Mkt string) (*[]Images, error) {
	newRequest := NewRequest(Format, Idx, N, Mkt)
	resp, err := http.Get("https://cn.bing.com/HPImageArchive.aspx?format=" + newRequest.format + "&idx=" + newRequest.idx + "&n=" + newRequest.n + "&mkt=" + newRequest.mkt)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var alldata []byte
	if resp.StatusCode == 200 {
		alldata, _ = io.ReadAll(resp.Body)
	}
	respstr := &AutoGenerated{}
	json.Unmarshal(alldata, respstr)
	return &respstr.Images, err
}

func GetImgData(iu Images, cimg chan *canvas.Image, wg *sync.WaitGroup) {
	urls, err := storage.ParseURI("https://cn.bing.com" + iu.URL) //解析url字符串
	if err != nil {
		log.Fatal(err)
	}
	cimg <- canvas.NewImageFromURI(urls)
	wg.Done()
}

func introduce() *widget.RichText {
	text := `


## 免责声明
- 数据来自Bing OpenAPI。
- 应该是每日稳定更新。
- 还没学会做保存功能。
- Power By [MaxBit](https://cakepanit.com)。
`
	return widget.NewRichTextFromMarkdown(text)
}

func Wallpaper(w fyne.Window) fyne.CanvasObject {
	cimge := make(chan *canvas.Image, 8)
	imgs, err := RequestImgURL("js", "0", "8", "zh-CN")
	if err != nil {
		label := widget.NewLabel("请检查网络～")
		return container.New(layout.NewGridLayout(1), label)
	} else {
		var wg sync.WaitGroup
		imagelist := &[]*canvas.Image{}
		for _, iu := range *imgs {
			wg.Add(1)
			go GetImgData(iu, cimge, &wg)
		}

		wg.Wait()
		close(cimge)

		for is := range cimge {
			*imagelist = append(*imagelist, is)
		}

		return container.New(layout.NewGridLayout(3), (*imagelist)[0], (*imagelist)[1], (*imagelist)[2], (*imagelist)[3], (*imagelist)[4], (*imagelist)[5], (*imagelist)[6], (*imagelist)[7], introduce())
	}
}

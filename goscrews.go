package goscrews

import (
	"github.com/sav7ng/go-screws.git/aes"
	"github.com/sav7ng/go-screws.git/array"
	"github.com/sav7ng/go-screws.git/bean"
	"github.com/sav7ng/go-screws.git/calendar"
	"github.com/sav7ng/go-screws.git/catch"
	"github.com/sav7ng/go-screws.git/convert"
	"github.com/sav7ng/go-screws.git/emoji"
	"github.com/sav7ng/go-screws.git/http"
	"github.com/sav7ng/go-screws.git/ip"
	"github.com/sav7ng/go-screws.git/mongo"
	"github.com/sav7ng/go-screws.git/qrcode"
	"github.com/sav7ng/go-screws.git/random"
	"github.com/sav7ng/go-screws.git/time"
	"github.com/sav7ng/go-screws.git/validator"
)

var (
	// AES加密工具
	AesTool = &aes.AesTool{}
	// 数组切片工具
	ArrayTool = &array.ArrayTool{}
	// Bean工具
	BeanTool = &bean.BeanTool{}
	// 捕抓工具
	CatchTool = &catch.CatchTool{}
	// 转换器工具
	ConvertTool = &convert.ConvertTool{}
	// emoji表情包工具
	EmojiTool = &emoji.EmojiTool{}
	// http请求工具
	HttpTool = &http.HttpTool{}
	// IP工具
	IPTool = &ip.IPTool{}
	// Mongodb工具
	MongoTool = &mongo.MongoTool{}
	// 二维码工具
	QRCodeTool = &qrcode.QRCodeTool{}
	// 随机数工具
	RandomTool = &random.RandomTool{}
	// 日历工具
	CalendarTool = &calendar.CalendarTool{}
	// 验证器工具
	ValidatorTool = &validator.ValidatorTool{}
	// 日期时间工具
	TimeTool = &time.TimeTool{}
)

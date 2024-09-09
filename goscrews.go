package goscrews

import (
	"github.com/sav7ng/go-screws/aes"
	"github.com/sav7ng/go-screws/array"
	"github.com/sav7ng/go-screws/bean"
	"github.com/sav7ng/go-screws/calendar"
	"github.com/sav7ng/go-screws/catch"
	"github.com/sav7ng/go-screws/convert"
	"github.com/sav7ng/go-screws/emoji"
	"github.com/sav7ng/go-screws/http"
	"github.com/sav7ng/go-screws/ip"
	"github.com/sav7ng/go-screws/mongo"
	"github.com/sav7ng/go-screws/qrcode"
	"github.com/sav7ng/go-screws/random"
	"github.com/sav7ng/go-screws/time"
	"github.com/sav7ng/go-screws/validator"
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

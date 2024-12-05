package feishu

import (
	"context"
	"encoding/json"
	"fmt"
	httpRequest "github.com/Songtingsen/go-utils/request"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"time"
)

const (
	// ResponseOkCode 返回成功状态码
	ResponseOkCode = 0
)

type BotMessage struct {
	// BotSource 机器人资源节点
	BotSource string
}

// At 消息@用户
type At struct {
	// Name 名称
	Name string
	// ID open_id 或 user_id
	ID string
}

// TextContent 文本类型消息内容结构
type TextContent struct {
	// Text 消息内容
	Text string `json:"text"`
}

// RichTextContent 富文本类型消息内容结构
type RichTextContent struct {
	// Title 标题
	Title string `json:"title"`
	// Contents 消息内容
	Contents [][]RichContentItem `json:"content"`
}

// RichContentItem 消息内容
type RichContentItem struct {
	// Tag 标签：text、文本 a、HTML a标签 at、@的人
	Tag string `json:"tag"`
	// Text 消息文本
	Text string `json:"text"`
	// Href a 标签链接地址
	Href string `json:"href"`
	// UserId @的用户ID
	UserId string `json:"user_id"`
}

// Response 返回结构体
type Response struct {
	StatusCode    int64  `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	Code          int64  `json:"code"`
	Data          any    `json:"data"`
	Msg           string `json:"msg"`
}

// NewBotMessage 初始化
func NewBotMessage(source string) *BotMessage {
	return &BotMessage{BotSource: source}
}

// SendTextMsg 发送文本消息
// 参考：https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot#5a997364
// params：
// ctx：上下文
// content：消息内容
// at：消息@的人
// return：
// error：消息发送是否成功，nil表示成功，否则失败
func (l *BotMessage) SendTextMsg(ctx context.Context, content string, at []At) error {
	logc.Infof(ctx, "发送消息参数，content：%s，at：%+v", content, at)

	// 拼接@用户信息
	var atStr string
	if len(at) > 0 {
		for _, item := range at {
			atStr += fmt.Sprintf(`<at user_id="%s">%s</at>`, item.ID, item.Name)
		}
	}

	// 拼接消息体
	text := TextContent{Text: atStr + content}
	textByte, err := json.Marshal(text)
	if err != nil {
		logc.Errorf(ctx, "消息体JSON格式化失败，text：%+v，error：%s", text, err)
		return err
	}

	// 拼接请求参数
	message := map[string]any{
		"msg_type": "text",
		"content":  string(textByte),
	}

	// 设置header
	header := map[string]string{"Content-Type": "application/json"}

	// 请求飞书发送消息
	resByte, err := httpRequest.DoRequest(ctx, l.BotSource, http.MethodPost, message, header, 2*time.Second)
	if err != nil {
		logc.Errorf(ctx, "请求飞书发送消息失败：%s", err)
		return err
	}
	logc.Infof(ctx, "请求飞书发送消息结果：%s", string(resByte))

	// 接收结果
	var res Response
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		logc.Errorf(ctx, "接收结果失败：%s", err)
		return err
	}
	if res.Code != ResponseOkCode {
		logc.Errorf(ctx, "消息发送失败：%s", res.Msg)
		return err
	}

	return nil
}

// SendRichTextMessage 发送富文本消息，富文本支持的标签：文本标签text、超链接标签a、@ 标签at、图片标签img
// 参考：https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot#5a997364
// params：
// ctx：上下文
// content：消息内容
// at：消息@的人
// return：
// error：消息发送是否成功，nil表示成功，否则失败
func (l *BotMessage) SendRichTextMessage(ctx context.Context, content RichTextContent, at []At) error {
	logc.Infof(ctx, "发送富文本消息参数，content：%+v，at：%+v", content, at)

	// 拼接@用户信息
	if len(at) > 0 {
		var RichContentAts []RichContentItem
		for _, item := range at {
			RichContentAts = append(RichContentAts, RichContentItem{
				Tag:    "at",
				UserId: item.ID,
			})
		}
		content.Contents[0] = append(content.Contents[0], RichContentAts...)
	}

	// 格式化消息体
	contentMap := map[string]any{
		"post": map[string]any{
			"zh_cn": content,
		},
	}
	textByte, err := json.Marshal(contentMap)
	if err != nil {
		logc.Errorf(ctx, "消息体JSON格式化失败，text：%+v，error：%s", content, err)
		return err
	}

	// 拼接请求参数
	message := map[string]any{
		"msg_type": "post",
		"content":  string(textByte),
	}

	// 设置header
	header := map[string]string{"Content-Type": "application/json"}

	// 请求飞书发送消息
	resByte, err := httpRequest.DoRequest(ctx, l.BotSource, http.MethodPost, message, header, 2*time.Second)
	if err != nil {
		logc.Errorf(ctx, "请求飞书发送消息失败：%s", err)
		return err
	}
	logc.Infof(ctx, "请求飞书发送消息结果：%s", string(resByte))

	// 接收结果
	var res Response
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		logc.Errorf(ctx, "接收结果失败：%s", err)
		return err
	}
	if res.Code != ResponseOkCode {
		logc.Errorf(ctx, "消息发送失败：%s", res.Msg)
		return err
	}

	return nil
}

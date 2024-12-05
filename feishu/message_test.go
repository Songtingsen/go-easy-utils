package feishu

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

const (
	hookUrl = "飞书hook url"
)

func TestBotMessage_SendMessage(t *testing.T) {
	type fields struct {
		BotSource string
	}
	type args struct {
		ctx     context.Context
		content string
		at      []At
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "发送飞书文本消息",
			fields: fields{
				BotSource: hookUrl,
			},
			args: args{
				ctx: context.Background(),
				content: `消息测试:
日期：2024-01-24
天气：晴天☀️
天空颜色：<span style="color:blue">蓝色</span>`,
				at: []At{
					{Name: "系统消息", ID: "1"},
					{Name: "所有人", ID: "all"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &BotMessage{
				BotSource: tt.fields.BotSource,
			}
			if err := l.SendTextMsg(tt.args.ctx, tt.args.content, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBotMessage_SendRichTextMessage(t *testing.T) {
	type fields struct {
		BotSource string
	}
	type args struct {
		ctx     context.Context
		content RichTextContent
		at      []At
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "飞书富文本消息测试",
			fields: fields{
				BotSource: hookUrl,
			},
			args: args{
				ctx: context.Background(),
				content: RichTextContent{
					Title: "飞书富文本消息测试",
					Contents: [][]RichContentItem{
						{
							{
								Tag:  "text",
								Text: "测试下啦",
							},
							{
								Tag:  "a",
								Text: "带你飞",
								Href: "https://www.baidu.com/",
							},
						},
					},
				},
				at: []At{
					{Name: "系统消息", ID: "1"},
					{Name: "所有人", ID: "all"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &BotMessage{
				BotSource: tt.fields.BotSource,
			}
			if err := l.SendRichTextMessage(tt.args.ctx, tt.args.content, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("SendRichTextMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBotMessage_SendRichTextMessage2(t *testing.T) {
	title := "消息剩余数据通知："
	id := ""
	content := [][]RichContentItem{
		{
			{
				Tag:  "text",
				Text: "模型可使用数量剩余：",
			},
			{
				Tag:  "text",
				Text: strconv.FormatInt(100, 10),
			},
		},
	}
	RichTextContent := RichTextContent{
		Title:    title,
		Contents: content,
	}

	botClient := NewBotMessage(hookUrl)
	var atId []At
	if id == "" {
		atId = []At{}
	} else {
		atId = []At{{ID: id}}
	}
	err := botClient.SendRichTextMessage(context.Background(), RichTextContent, atId)
	if err != nil {
		fmt.Println("错误", err)
		return
	}

	return
}

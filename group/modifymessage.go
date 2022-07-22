package group

import (
	"github.com/DanPlayer/tencent-im/internal/enum"
	"github.com/DanPlayer/tencent-im/internal/types"
)

type ModifyMessage struct {
	id         string           // 发送者
	msgSeq     string           // 消息标识
	body       []*types.MsgBody // 消息体
	customData interface{}      // 自定义数据
}

func (m *ModifyMessage) NewModify(id, key string) *ModifyMessage {
	m.id = id
	m.msgSeq = key
	return m
}

func (m *ModifyMessage) SetID(id string) {
	m.id = id
}

func (m *ModifyMessage) GetID() string {
	return m.id
}

func (m *ModifyMessage) SetMsgSeq(key string) {
	m.msgSeq = key
}

func (m *ModifyMessage) GetMsgSeq() string {
	return m.msgSeq
}

// AddContent 添加消息内容（添加会累加之前的消息内容）
func (m *ModifyMessage) AddContent(msgContent ...interface{}) {
	if m.body == nil {
		m.body = make([]*types.MsgBody, 0)
	}

	if len(msgContent) > 0 {
		var msgType string
		for _, content := range msgContent {
			switch content.(type) {
			case types.MsgTextContent, *types.MsgTextContent:
				msgType = enum.MsgText
			case types.MsgLocationContent, *types.MsgLocationContent:
				msgType = enum.MsgLocation
			case types.MsgFaceContent, *types.MsgFaceContent:
				msgType = enum.MsgFace
			case types.MsgCustomContent, *types.MsgCustomContent:
				msgType = enum.MsgCustom
			case types.MsgSoundContent, *types.MsgSoundContent:
				msgType = enum.MsgSound
			case types.MsgImageContent, *types.MsgImageContent:
				msgType = enum.MsgImage
			case types.MsgFileContent, *types.MsgFileContent:
				msgType = enum.MsgFile
			case types.MsgVideoContent, *types.MsgVideoContent:
				msgType = enum.MsgVideo
			default:
				msgType = ""
			}

			m.body = append(m.body, &types.MsgBody{
				MsgType:    msgType,
				MsgContent: content,
			})
		}
	}
}

// SetContent 设置消息内容（设置会冲掉之前的消息内容）
func (m *ModifyMessage) SetContent(msgContent ...interface{}) {
	if m.body != nil {
		m.body = m.body[0:0]
	}
	m.AddContent(msgContent...)
}

// GetBody 获取消息体
func (m *ModifyMessage) GetBody() []*types.MsgBody {
	return m.body
}

// SetCustomData 设置自定义数据
func (m *ModifyMessage) SetCustomData(data interface{}) {
	m.customData = data
}

// GetCustomData 获取自定义数据
func (m *ModifyMessage) GetCustomData() interface{} {
	return m.customData
}

package coin

import (
	"fmt"
	"time"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
	coinReq "github.com/Grace1China/cointown/server/model/coin/request"
)

type RawMsgService struct{}

// CreateRawMsg 创建原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) CreateRawMsg(rawmsg *coin.RawMsg) (err error) {
	err = global.GVA_DB.Create(rawmsg).Error
	return err
}

// DeleteRawMsg 删除原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) DeleteRawMsg(ID string) (err error) {
	err = global.GVA_DB.Delete(&coin.RawMsg{}, "id = ?", ID).Error
	return err
}

// DeleteRawMsg 删除原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) DeleteRawMsg7D() (err error) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	err = global.GVA_DB.Delete(&[]coin.RawMsg{}, "created_at < ?", sevenDaysAgo).Error
	return err
}

// DeleteRawMsgByIds 批量删除原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) DeleteRawMsgByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]coin.RawMsg{}, "id in ?", IDs).Error
	return err
}

// UpdateRawMsg 更新原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) UpdateRawMsg(rawmsg coin.RawMsg) (err error) {
	err = global.GVA_DB.Model(&coin.RawMsg{}).Where("id = ?", rawmsg.ID).Updates(&rawmsg).Error
	return err
}

// UpdateRawMsg 更新原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) UpdateRawMsgStatusAll() (err error) {
	// err = global.GVA_DB.Model(&coin.RawMsg{}).Where("id = ?", rawmsg.ID).Updates(&rawmsg).Error
	err = global.GVA_DB.Model(&coin.RawMsg{}).Where("is_new != ?", "0").Update("is_new", "0").Error
	return err
}

// GetRawMsg 根据ID获取原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) GetRawMsg(ID string) (rawmsg coin.RawMsg, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&rawmsg).Error
	return
}

// GetRawMsgInfoList 分页获取原始消息记录
// Author [piexlmax](https://github.com/piexlmax)
func (rawmsgService *RawMsgService) GetRawMsgInfoList(info coinReq.RawMsgSearch) (list []coin.RawMsg, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&coin.RawMsg{})
	var rawmsgs []coin.RawMsg
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		fmt.Println("info.StartCreatedAt", info.StartCreatedAt)
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	fmt.Println(info)
	if info.ChatId != "" {
		fmt.Println("info.ChatId", info.ChatId)
		// 添加额外的条件
		db = db.Where("chat_id = ?", info.ChatId)
	}
	if info.Id != "" {
		fmt.Println("info.id", info.Id)
		// 添加额外的条件
		db = db.Where("id = ?", info.Id)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id DESC")
	}

	err = db.Find(&rawmsgs).Error
	return rawmsgs, total, err
}

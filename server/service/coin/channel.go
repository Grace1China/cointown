package coin

import (
	"fmt"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
	coinReq "github.com/Grace1China/cointown/server/model/coin/request"
	"gorm.io/gorm"
)

type ChannelService struct{}

// CreateChannel 创建频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) CreateChannel(channel *coin.Channel) (err error) {
	if channel.HitAt == nil || channel.HitAt.IsZero() {
		channel.HitAt = nil
		err = global.GVA_DB.Create(channel).Error
		return err
	} else {
		return
	}

}

// DeleteChannel 删除频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) DeleteChannel(ID string) (err error) {
	err = global.GVA_DB.Delete(&coin.Channel{}, "id = ?", ID).Error
	return err
}

// DeleteChannelByIds 批量删除频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) DeleteChannelByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]coin.Channel{}, "id in ?", IDs).Error
	return err
}

// UpdateChannel 更新频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) UpdateChannel(channel coin.Channel) (err error) {
	err = global.GVA_DB.Model(&coin.Channel{}).Where("id = ?", channel.ID).Updates(&channel).Error
	return err
}

// UpdateChannel 更新频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) CreateOrUpdateChannel(channel coin.Channel) (err error) {
	err = global.GVA_DB.Where("id = ?", channel.ChatId).First(&channel).Error
	// 检查是否有记录未找到的错误
	if err == gorm.ErrRecordNotFound {
		fmt.Println("没有找到记录", channel)

		err = global.GVA_DB.Create(&channel).Error
		fmt.Println("创建完成", err)
		if err != nil {
			// 如果有其他错误
			fmt.Printf("查询时发生错误: %v\n", err)
		}
		return err
	} else if err != nil {
		// 如果有其他错误
		fmt.Printf("查询时发生错误: %v\n", err)
	} else {
		// 如果查询成功并且找到了记录
		fmt.Printf("查询成功, 记录: %v\n", channel)
		err = global.GVA_DB.Model(&coin.Channel{}).Where("id = ?", channel.ID).Updates(&channel).Error
		return err
	}

	//这里通过对channel.id的判断来看是否是有数
	fmt.Println("CreateOrUpdateChannel", channel, channel.ID)
	// err = global.GVA_DB.Model(&coin.Channel{}).Where("id = ?", channel.ID).Updates(&channel).Error
	return err
}

// GetChannel 根据ID获取频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) GetChannel(ID string) (channel coin.Channel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&channel).Error
	return
}

// GetChannel 根据ID获取频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) GetChannelByChlTopicID(chatId string, topicId string) (list []coin.Channel, total int64, err error) {
	global.GVA_LOG.Info(fmt.Sprintf("GetChannelByChlTopicID:ChatId:%s TopicId:%s", chatId, topicId))
	db := global.GVA_DB.Model(&coin.Channel{})
	var channels []coin.Channel
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("chat_id = ? and topic_id = ?", chatId, topicId)

	err = db.Count(&total).Error

	if err != nil {
		return
	}

	// if limit != 0 {
	// 	db = db.Limit(limit).Offset(offset)
	// }

	err = db.Find(&channels).Error
	return channels, total, err
}

// GetChannelInfoList 分页获取频道记录
// Author [piexlmax](https://github.com/piexlmax)
func (channelService *ChannelService) GetChannelInfoList(info coinReq.ChannelSearch) (list []coin.Channel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&coin.Channel{})
	var channels []coin.Channel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&channels).Error
	return channels, total, err
}

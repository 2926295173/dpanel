// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"github.com/donknap/dpanel/common/accessor"
)

const TableNameSetting = "ims_setting"

// Setting mapped from table <ims_setting>
type Setting struct {
	ID        int32                        `gorm:"column:id;type:INTEGER" json:"id"`
	GroupName string                       `gorm:"column:group_name" json:"groupName"`
	Name      string                       `gorm:"column:name" json:"name"`
	Value     *accessor.SettingValueOption `gorm:"column:value;serializer:json" json:"value"`
}

// TableName Setting's table name
func (*Setting) TableName() string {
	return TableNameSetting
}
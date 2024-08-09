// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTeam = "teams"

// Team mapped from table <teams>
type Team struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	OrganizationID string    `gorm:"column:organization_id;not null" json:"organization_id"`
	Name           string    `gorm:"column:name;not null" json:"name"`
}

// TableName Team's table name
func (*Team) TableName() string {
	return TableNameTeam
}
// Copyright 2021, The Go Authors. All rights reserved.
// Author: crochee
// Date: 2021/3/4

package db

import (
	"time"
)

type Bucket struct {
	ID uint `gorm:"primary_key"`

	Domain string `json:"domain" gorm:"column:domain;type:varchar(50);not null"`
	Bucket string `json:"bucket" gorm:"column:bucket;type:varchar(50);not null"`
	User   string `json:"user" gorm:"column:user;type:varchar(50);not null"`

	BucketFileList []BucketFile `gorm:"ForeignKey:id;AssociationForeignKey:id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Bucket) TableName() string {
	return "bucket_domain"
}
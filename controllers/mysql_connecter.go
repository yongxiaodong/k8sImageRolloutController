package controllers

//package main

import (
	"context"
	"github.com/go-logr/logr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ImageInfo struct {
	DpName    string `gorm:"column:DeploymentName"`
	NameSpace string `gorm:"column:Namespace"`
	Image     string `gorm:"type:text,column:Image"`
	//CreatedAt time.Time `gorm:"autoCreateTime;column:CreateTime;<-:create"`
	//UpdatedAt time.Time `gorm:"autoUpdateTime;column:UpdateTime"`
}

var log2 logr.Logger
var db *gorm.DB
var err error

func init() {
	log2 = log.FromContext(context.Background())
	dsn := "Username:Password@tcp(office.test.com:3306)/k8s-snapshot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log2.Error(err, "connect Mysql error")
	}
}

func MysqlWrite(data ImageInfo) {
	log2.Info("start write data...")
	db.Create(&data)
	if db.Error != nil {
		log2.Error(db.Error, "write db fail.. ")
	}
}

//func MysqlQuery(data ImageInfo) (res []ImageInfo) {
//	log2.Info("start query data...")
//	var datas []ImageInfo
//	if data.NameSpace != "" {
//		db.Where("ImageInfo_name = ? and namespace = ?", data.ImageInfoName, data.NameSpace).Find(&datas)
//
//	} else {
//		db.Where("ImageInfo_name = ?", data.ImageInfoName).Find(&datas)
//	}
//	return datas
//
//}
//
//func MysqlDelete(data ImageInfo) {
//	log2.Info(fmt.Sprintf("delete data.... %s", data.ImageInfoName))
//	db.Where("ImageInfo_name = ?", data.ImageInfoName).Delete(&ImageInfo{})
//}

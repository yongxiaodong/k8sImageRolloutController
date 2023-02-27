package controllers

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
	"time"
)

type PostData struct {
	DpName        string
	SelectedImage DpImage `json:"selectedImage"`
}

type AllDp struct {
	ID     int
	DpName string `gorm:"column:DeploymentName"`
}

type DpImage struct {
	ID        int
	DpName    string    `gorm:"column:DeploymentName"`
	Namespace string    `gorm:"column:NameSpace"`
	Image     string    `gorm:"type:text,column:Image"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:CreateTime;<-:create"`
}
type DpImage2 struct {
	ID           int
	DpName       string    `gorm:"column:DeploymentName"`
	Namespace    string    `gorm:"column:NameSpace"`
	Image        string    `gorm:"type:text,column:Image"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:CreateTime;<-:create"`
	CurrentImage string
}

// get deploymentName
func GetAllDpName() (res []AllDp) {
	res = make([]AllDp, 0)
	db.Model(&ImageInfo{}).Group("DeploymentName").Order("CreateTime desc").Find(&res)
	return res
}

func GetDpImage(dp string) (res []DpImage2) {
	res = make([]DpImage2, 0)
	db.Model(&ImageInfo{}).Where("DeploymentName = ?", dp).Limit(15).Order("CreateTime desc").Find(&res)
	config, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(config)

	dpres, _ := clientset.AppsV1().Deployments(res[0].Namespace).Get(context.TODO(), res[0].DpName, metav1.GetOptions{})
	cs := ""
	for _, v := range dpres.Spec.Template.Spec.Containers {
		cs += v.Image
	}
	res[0].CurrentImage = cs

	return res
}

type Ginres struct {
	Status  int
	Message string
}

func checkMysqlData(dp string, image string) bool {
	res := make([]DpImage, 0)
	db.Model(&ImageInfo{}).Where("DeploymentName = ? and Image = ?", dp, image).Find(&res)
	if len(res) == 1 {
		return true
	}
	return false
}

func SetImage(data PostData) (gres Ginres) {
	// check deploymen name and version
	// CheckMysqlData
	if !checkMysqlData(data.DpName, data.SelectedImage.Image) {
		gres.Status = 402
		gres.Message = "Project Name OR Image Error"
		return gres
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		gres.Status = 401
		return gres
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		gres.Status = 571
		return gres
	}
	dp, err := clientset.AppsV1().Deployments(data.SelectedImage.Namespace).Get(context.TODO(), data.DpName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		gres.Status = 500
		gres.Message = fmt.Sprintf("deployment not fount: %s", err)
		return gres
	} else if err != nil {
		gres.Status = 500
		gres.Message = fmt.Sprintf("unkown fail: %s", err)
		return gres
	}
	Image := data.SelectedImage.Image
	newImage := strings.Split(Image, ",")
	//fmt.Println(dp)
	for i, v := range newImage {
		dp.Spec.Template.Spec.Containers[i].Image = v
		fmt.Printf("update %s: container %v--%v \n", data.DpName, i, v)
	}
	_, err = clientset.AppsV1().Deployments(data.SelectedImage.Namespace).Update(context.TODO(), dp, metav1.UpdateOptions{})
	if err != nil {
		gres.Status = 500
		gres.Message = fmt.Sprintf("update deployment fail: %s", err)
	} else {
		gres.Status = 0
		gres.Message = fmt.Sprintf("project: %s, success", data.DpName)
	}
	return gres

}

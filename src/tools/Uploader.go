package tools

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sports-common/config"
	"sports-common/consts"
	"sports-common/response"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

// FileUploader 文件上传
type FileUploader struct {
	FileTypes  []string //允许上传类型
	UploadPath string   //上传目录
	MaxSize    int64    //最大允许上传, 0表示不限制
	Name       string   //上传文件字段名称
	URLPath    string   //返回显示的URL(图片URL)
}

// Uploader 文件上传
var Uploader = struct {
	New func() *FileUploader
}{
	New: func() *FileUploader {
		return &FileUploader{
			FileTypes:  []string{"jpg", "png", "jpeg", "gif", "webp", "xlsx"},
			UploadPath: consts.UploadPath,
			MaxSize:    2048000, //最大上传
			Name:       "file",  //上传提交表单名称
			URLPath:    consts.UploadURLPath,
		}
	},
}

// SetFileTypes 文件类型
func (ths *FileUploader) SetFileTypes(types []string) *FileUploader {
	ths.FileTypes = types
	return ths
}

// SetUploadPath 设置上传路径
func (ths *FileUploader) SetUploadPath(path string) *FileUploader {
	ths.UploadPath = path
	return ths
}

// SetMaxSize 设置上传允许最大值
func (ths *FileUploader) SetMaxSize(size int64) *FileUploader {
	ths.MaxSize = size
	return ths
}

// GetSaveDir 得到要上传的文件的路径
func (ths *FileUploader) GetSaveDir() (string, string, error) {
	if !IsDir(ths.UploadPath) {
		fmt.Println("路径有误: ", ths.UploadPath)
		return "", "", errors.New("要上传的文件目录不存在")
	}
	year, month, day := Date()
	yearDir := fmt.Sprintf("%s/%d", strings.TrimRight(ths.UploadPath, "/"), year)
	if !IsDir(yearDir) {
		_ = os.Mkdir(yearDir, os.ModePerm)
	}
	monthDir := fmt.Sprintf("%s/%02d", yearDir, month)
	if !IsDir(monthDir) {
		_ = os.Mkdir(monthDir, os.ModePerm)
	}
	dayDir := fmt.Sprintf("%s/%02d", monthDir, day)
	if !IsDir(dayDir) {
		_ = os.Mkdir(dayDir, os.ModePerm)
	}
	urlPath := fmt.Sprintf("%s/%d/%02d/%02d", strings.TrimRight(ths.URLPath, "/"), year, month, day)
	return dayDir, urlPath, nil
}

// Upload 上传文件
func (ths *FileUploader) Upload(c *gin.Context) (string, error) {
	c.Header("Access-Control-Allow-Origin", "*")

	name := ths.Name
	file, err := c.FormFile(name)
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	extArr := strings.Split(fileName, ".")
	extLen := len(extArr)

	if extLen < 2 {
		return "", errors.New("缺少文件扩展名称")
	}
	extName := extArr[extLen-1]
	extAllow := false
	for _, v := range ths.FileTypes {
		if v == extName {
			extAllow = true
			break
		}
	}

	if !extAllow {
		return "", errors.New("不允许上传的文件名称")
	}

	if file.Size > ths.MaxSize {
		return "", errors.New("上传文件过大")
	}

	saveDir, urlPath, err := ths.GetSaveDir()
	if err != nil {
		return "", err
	}

	saveFileName := RandString(16) + "." + extName
	oldFileName := extArr[0]
	isLetterAndNumber, _ := regexp.MatchString(`^[A-Za-z0-9]+$`, oldFileName)
	if isLetterAndNumber { //只有字母和数字 640x960,1242x2208,1440x2560
		saveFileName = RandString(16) + "-" + oldFileName + "." + extName
	}
	saveFilePath := saveDir + "/" + saveFileName
	urlFilePath := urlPath + "/" + saveFileName

	fmt.Println("saveFilePath = ", saveFilePath)
	fmt.Println("urlFilePath = ", urlFilePath)

	err = c.SaveUploadedFile(file, saveFilePath)
	if err != nil {
		return "", err
	}

	UploadToAWS(saveFilePath, urlFilePath) // 上传到aws - s3
	return urlFilePath, err
}

// UploadFile 处理上传文件
func UploadFile(c *gin.Context) {
	path, err := Uploader.New().Upload(c)

	if err != nil {

		response.Err(c, err.Error())
		return
	}

	response.Result(c, struct {
		Path string `json:"path"`
	}{
		Path: path,
	})
}

// UploadToAWS 上传到亚马逊s3
func UploadToAWS(fileName string, uriPath string) {
	accessID := config.Get("aws.access_id", "")
	accessKey := config.Get("aws.access_key", "")
	bucket := config.Get("aws.bucket", "")
	if accessID == "" || accessKey == "" || bucket == "" {
		fmt.Println("未设置 aws - access id/key/bucket = ", accessID, ", ", accessKey, ", ", bucket)
		return
	}

	os.Setenv("AWS_ACCESS_KEY_ID", accessID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", accessKey)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	}))
	uploader := s3manager.NewUploader(sess)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("failed to open file ", fileName, err)
		return
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(uriPath),
		Body:   f,
	})
	if err != nil {
		fmt.Println("上传文件失败: ", err)
		return
	}
	res := aws.StringValue(&result.Location)
	fmt.Printf("文件已上传至 AWS: %s\n", res)
}

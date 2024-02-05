package es

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sports-common/consts"
	"time"

	"github.com/olivere/elastic/v7"
)

// GetClient
func GetClient(platform string) *elastic.Client {
	clientLock.Lock()
	defer clientLock.Unlock()

	if len(clientList) == 0 {
		clientList[platform] = map[string]*elastic.Client{}
	}

	for k, v := range clientList {
		if k != platform {
			continue
		}
		for ck, cv := range v {
			delete(clientList[k], ck)
			return cv
		}
	}

	client, _ := GetClientByPlatform(platform)
	return client
}

// ReturnClient
func ReturnClient(platform string, client *elastic.Client) {
	clientLock.Lock()
	defer clientLock.Unlock()

	if len(clientList) == 0 {
		clientList[platform] = map[string]*elastic.Client{}
	}

	k := fmt.Sprintf("%p", client)
	if len(clientList[platform]) == 0 {
		clientList[platform] = map[string]*elastic.Client{
			k: client,
		}
		return
	}

	clientList[platform][k] = client
}

// 初始化elasticSearch
func GetClientByPlatform(platform string) (client *elastic.Client, err error) {
	conf, exists := ServerUrls[platform]
	if !exists {
		return nil, errors.New("缺少平台" + platform + "的es相关配置信息")
	}

	var cfg = []elastic.ClientOptionFunc{}
	logFile, err := GetLogFile()
	if logFile == nil || err != nil {
		panic("无法建立 elasticsearch 连接:")
	}

	cfg = append(cfg,
		elastic.SetURL(conf...),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)))

	logDebug := false
	if logDebug {
		cfg = append(cfg,
			elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
			elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)))
	}

	client, err = elastic.NewClient(cfg...)
	return client, err
}

// 得到es客户端
func GetEsClient(esURL string, args ...string) (client *elastic.Client, err error) {
	var cfg = []elastic.ClientOptionFunc{}
	esLogDir := consts.LogPath
	if len(args) > 0 {
		esLogDir = args[0]
	}
	logFile, err := GetLogFile(esLogDir)
	if logFile == nil || err != nil {
		panic("无法建立 elasticsearch 连接(打开日志文件失败):" + err.Error())
	}

	cfg = append(cfg,
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)))

	logDebug := false
	if logDebug {
		cfg = append(cfg,
			elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
			elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)))
	}

	client, err = elastic.NewClient(cfg...)
	return client, err
}

// GetLogFile
func GetLogFile(args ...string) (*os.File, error) {
	fileDate := time.Now().Format("200601")
	esLogDir := consts.LogPath
	if len(args) > 0 {
		esLogDir = args[0]
	}
	fileDir := esLogDir + "/es/" + fileDate
	if esLogDir == "./" {
		fileDir = esLogDir + "es/" + fileDir
	}

	var err error
	if !isDir(fileDir) {
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	fileName := "es-client-" + fileDate + ".log"
	fileName = fileDir + "/" + fileName
	var logFile *os.File
	if !isFile(fileName) {
		if logFile, err = os.Create(fileName); err != nil {
			return nil, err
		}

		if err = os.Chmod(fileName, 0766); err != nil {
			return nil, err
		}

		return logFile, nil
	}

	if logFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766); err != nil {
		return nil, err
	}
	return logFile, nil
}

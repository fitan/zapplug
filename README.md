# zapplug
**golang zap日志的插件 完成了切割 循环 最大文件限制 保存时间**
```
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"zapplug"
	"github.com/olivere/elastic"
	"fmt"
)



func main()  {
	pConf := zapplug.CutConf{}
	pConf.FileName = "./logs"
	pConf.Compress = true
	pConf.MaxAge = 3
	pConf.MaxSize = 100
	pConf.MaxBackups = 3

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println(err)
		return
	}


	core := zapcore.NewTee(
		zapplug.ZapCut(pConf,zap.NewDevelopmentEncoderConfig(),zap.DebugLevel),
		zapplug.Es(client, "testlog", "10.106.132.4",zap.DebugLevel, zap.NewDevelopmentEncoderConfig()),
	)
	log := zap.New(core)
	defer log.Sync()
	for {
		log.Info("test",zap.String("test","test"),)
	}
}
```

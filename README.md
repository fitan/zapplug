# zapplug
<<<<<<< HEAD
**golang zap日志的插件 完成了切割 循环 保存时间**
=======
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
	pConf.FileName = "./logs" //文件位置
	pConf.Compress = true //是否压缩
	pConf.MaxAge = 3 //保留旧日志文件的最大天数
	pConf.MaxSize = 100 //文件大小 M
	pConf.MaxBackups = 3 //保留的旧日志文件的最大数量
	
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println(err)
		return
	}


	core := zapcore.NewTee(
		//滚动插件
		zapplug.ZapCut(pConf,zap.NewDevelopmentEncoderConfig(),zap.DebugLevel),
		//es 插件
		//testlog index
		//zap.DebugLevel  leverl
		//zap.NewDevelopmentEncoderConfig()  zap logconf
		zapplug.Es(client, "testlog",zap.DebugLevel, zap.NewDevelopmentEncoderConfig()),
	)
	log := zap.New(core)
	defer log.Sync()
	for {
		log.Info("test",zap.String("test","test"),)
	}
}
```
>>>>>>> a405bf5fc647a64d00a19af3c1e915c047f0489d

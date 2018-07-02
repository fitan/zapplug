package zapplug

import (
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
	"github.com/olivere/elastic"
	"context"
	"unsafe"
)
func Es(client *elastic.Client, index string, l zapcore.Level, z zapcore.EncoderConfig) zapcore.Core {
	elastic.NewClient()
	es := &toes{client:client, index:index}
	es.cIndex()
	zlvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l
	})
	code := zapcore.NewJSONEncoder(z)
	return zapcore.NewCore(code,zapcore.AddSync(es),zlvl)
}

type toes struct {
	client *elastic.Client
	index string
}

func (this *toes) Write(p []byte) (n int, err error) {
	err = this.inset(p)
	if err != nil {
		return 0, err
	}
	return 0,nil
}

func (this *toes) inset(b []byte) error {
	_, err := this.client.Index().Index(this.index).Type("log").BodyJson(*(*string)(unsafe.Pointer(&b))).Do(context.Background())
	return err
}

func (this *toes) cIndex() {
	exists, err := this.client.IndexExists(this.index).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := this.client.CreateIndex(this.index).Do(context.Background())
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			panic("不能创建index")
		}
	}
}

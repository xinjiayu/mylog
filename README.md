这个是为log日志做初始化配置的文件。在工程里只要直接引入这个文件。然后在main方法里调用`mylog.LogInit()`就可以使用支持分级显示的日志了。

在import中引入
```
	"github.com/xinjiayu/mylog"
	"github.com/xinjiayu/log"

```


然后在main方法中进行初始化`mylog.LogInit()`


然后就可以在各处使用log处理。支持如下方式：

```
	aaa := "aaaaaaaaa"
	log.Info("this is test", aaa)
	log.Debug("this is Debug", aaa)
	log.Error("this is Error", aaa)
	log.Fatal("this is Fatal", aaa)

```
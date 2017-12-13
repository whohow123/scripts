package conf

import (
	"gopkg.in/gcfg.v1"
	"runtime"
)

var (
	// global config object
	Conf *Config
)

type Config struct {
	Base struct {
		MaxProc   int      // CPU使用数
		PidPath   string   // 进程pid存储位置
		PprofAddr []string // pprof 监听的端口
		LogPath   string   // 日志配置文件路径
	}
}

// 初始化函数
func init() {
	Conf = &Config{}
	// 初始赋值 保证不为空
	Conf.Base.MaxProc = runtime.NumCPU()
	// 进程pid 文件默认路径
	Conf.Base.PidPath = "/tmp/"
	// 默认pprof 监听路径
	Conf.Base.PprofAddr = []string{}
	// 默认空路径
	Conf.Base.LogPath = ""
}

// 初始化 加载并解析配置文件到 Conf 对象
func Init(confPath string) error {
	if err := gcfg.ReadFileInto(Conf, confPath); err != nil {
		return err
	}
	return nil
}

// 重新加载 重载配置文件到 Conf 对象
func Reload(confPath string) error {
	tmp := &Config{}
	if err := gcfg.ReadFileInto(tmp, confPath); err != nil {
		return err
	}
	Conf = tmp
	return nil
}

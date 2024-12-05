package variable

import (
	"flag"
	"log"
	"nunu-rcc-stack/pkg/config"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var (
	ApplicationUrl   string       //项目ip
	ApplicationPort  int          //项目端口
	ApplicationName  string       //项目名称
	EthClientAddress string       //以太坊连接节点
	ENV              string       // 项目当前的环境
	BasePath         string       //项目根目录
	ClientSwitch     bool         //是否开启后端默认节点配置
	Config           *viper.Viper //配置文件

)

func init() {
	log.Println("**********初始化全局变量***************")
	if curPath, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		log.Println("初始化全局变量路径有误")
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	ENV = env
	configName := ""
	if env == "prod" {
		configName = "prod"
	} else {
		configName = "local"
	}

	var envConf = flag.String("conf", "config/"+configName+".yml", "config path, eg: -conf ./config/"+configName+".yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)
	if conf != nil {
		ApplicationUrl = conf.GetString("http.url")
		ApplicationPort = conf.GetInt("http.port")

		ethUrl := conf.GetString("ethclient.url")
		port := conf.GetInt("ethclient.port")

		if port != 0 {
			EthClientAddress = ethUrl + ":" + strconv.Itoa(port)
		} else {
			EthClientAddress = ethUrl
		}

		switchFlag := conf.GetBool("ethclient.clientSwitch")
		ClientSwitch = switchFlag
		Config = conf
	}
	log.Println("init results --- >   ApplicationUrl: ", ApplicationUrl, "; ApplicationPort: ", ApplicationPort, "; ApplicationName: ", ApplicationName, "; EthClientAddress: ", EthClientAddress, "; Env: ", ENV, "; BasePath: ", BasePath, "; ClientSwitch: ", ClientSwitch)
	log.Println("**********初始化全局变量结束***************")
}

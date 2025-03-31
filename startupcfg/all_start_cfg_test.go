package startupcfg_test

import (
	"fmt"
	"github.com/magic-lib/go-plat-startupcfg/startupcfg"
	"github.com/magic-lib/go-plat-utils/conn"
	"github.com/magic-lib/go-plat-utils/conv"
	"testing"
)

// TemplateURL 模板访问地址
type TemplateURL struct {
	TemplateGetCdBiz            string //获取CD业务信息
	TemplateBindCdBizTag        string //绑定部署标签
	TemplateAppserverSubmitById string //gdp_appserver_go 提交template执行的方法
	GetPodListByCdId            string
	InsertCI                    string
}

// GdpConfig gdp全局配置
type GdpConfig struct {
	HostAndPort              *conn.Connect
	MysqlConnect             *startupcfg.MysqlConfig
	MysqlConnectODP          *startupcfg.MysqlConfig
	RedisConnect             *startupcfg.RedisConfig
	GdpExternalOrigin        string
	ClientSecret             string
	TemplateIdBatchDeleteCd  string //批量删除部署的模板ID
	TemplateIdCopyCdWithCdId string
	TemplateIdCopyCd         string

	DefaultRTXLoginToken string //rtxLoginToken

	DefaultSystemRoleNameMap map[string][]string
}

func TestGetAllApiUrlMap(t *testing.T) {

	keyStr := "tianlin020250214"

	startupcfg.SetDefaultHandler(keyStr)
	enString, _ := startupcfg.Encrypted("aaaaa").Encode()

	fmt.Println(enString)

	one, _ := startupcfg.NewStartupForYamlFile("all_start_cfg_test.yaml")
	mapTemp := one.AllApiUrlMap()
	tempUrl := new(TemplateURL)
	conv.Unmarshal(mapTemp, tempUrl)

	fmt.Println(conv.String(tempUrl))

	cMap, _ := one.AllCustomMap()

	tempCMap := new(GdpConfig)
	conv.Unmarshal(cMap, tempCMap)

	ccTemp, _ := one.AllMysqlMap()
	conv.Unmarshal(ccTemp, tempCMap)

	ccTemp2, _ := one.AllRedisMap()

	conv.Unmarshal(ccTemp2, tempCMap)

	fmt.Println(conv.String(tempCMap))

}

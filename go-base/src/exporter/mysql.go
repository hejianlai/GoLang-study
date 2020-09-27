package main

import (
"net/http"
"os"

"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promhttp"
"github.com/prometheus/common/log"
)

// 指标采集器
type ApiExport struct {
	ApiName         string
	ApiSourceDesc   *prometheus.Desc
	ApiReqTotalDesc *prometheus.Desc
	ApiDeviceDesc   *prometheus.Desc
	ApiBrowserDesc  *prometheus.Desc
	ApiNetWorkDesc  *prometheus.Desc
}

const (
	// 接口请求的来源
	SOURCE_PC      = 1
	SOURCE_ANDROID = 2
	SOURCE_IOS     = 2

	// Device
	DEVICE_PC                = 1
	DEVICE_ANDROID_HUAWEI_10 = 2
	DEVICE_IOS_11            = 3

	// network
	NETWORK_2G = 1
	NETWORK_3G = 2
	NETWORK_4G = 3
	NETWORK_5G = 4

	// Brower
	BROWER_CHROME  = 1
	BROWER_FIREFOX = 2
	BROWER_IE      = 3
	BROWER_SAFARI  = 4
)

// 获取响应指标的数据，测试数据，正式环境需要修改
func (re *ApiExport) getApiSourceInfo() map[string]float64 {
	sourceInfo := map[string]float64{
		"127.0.0.1": float64(SOURCE_PC),
	}

	return sourceInfo
}

func (re *ApiExport) getApiTotalInfo() map[string]float64 {
	sourceInfo := map[string]float64{
		"127.0.0.1": float64(1000),
	}
	return sourceInfo
}

func (re *ApiExport) getApiDeviceInfo() map[string]float64 {
	sourceInfo := map[string]float64{
		"127.0.0.1": float64(DEVICE_PC),
	}
	return sourceInfo
}

func (re *ApiExport) getApiNetWorkInfo() map[string]float64 {
	sourceInfo := map[string]float64{
		"127.0.0.1": float64(NETWORK_5G),
	}
	return sourceInfo
}

func (re *ApiExport) getApiBrowerInfo() map[string]float64 {
	sourceInfo := map[string]float64{
		"127.0.0.1": float64(BROWER_CHROME),
	}
	return sourceInfo
}

/**
用于传递所有可能的指标的定义描述符
可以在程序运行期间添加新的描述，收集新的指标信息
重复的描述符将被忽略。两个不同的Collector不要设置相同的描述符
路径：github.com/prometheus/client_golang@v1.7.1/prometheus/collector.go
实现Collector接口
*/
func (re *ApiExport) Describe(ch chan<- *prometheus.Desc) {
	ch <- re.ApiSourceDesc
	ch <- re.ApiBrowserDesc
	ch <- re.ApiDeviceDesc
	ch <- re.ApiNetWorkDesc
	ch <- re.ApiReqTotalDesc
}

/**
Prometheus的注册器调用Collect执行实际的抓取参数的工作，
并将收集的数据传递到Channel中返回
收集的指标信息来自于Describe中传递，可以并发的执行抓取工作，但是必须要保证线程的安全。
路径：github.com/prometheus/client_golang@v1.7.1/prometheus/collector.go
实现Collector接口
*/
func (re *ApiExport) Collect(ch chan<- prometheus.Metric) {
	sourceInfo := re.getApiSourceInfo()
	for host, source := range sourceInfo {
		ch <- prometheus.MustNewConstMetric(
			re.ApiSourceDesc,
			prometheus.GaugeValue,
			source,
			host,
		)
	}

	browserInfo := re.getApiBrowerInfo()
	for host, source := range browserInfo {
		ch <- prometheus.MustNewConstMetric(
			re.ApiBrowserDesc,
			prometheus.GaugeValue,
			source,
			host,
		)
	}

	deviceInfo := re.getApiDeviceInfo()
	for host, source := range deviceInfo {
		ch <- prometheus.MustNewConstMetric(
			re.ApiDeviceDesc,
			prometheus.GaugeValue,
			source,
			host,
		)
	}

	networkInfo := re.getApiNetWorkInfo()
	for host, source := range networkInfo {
		ch <- prometheus.MustNewConstMetric(
			re.ApiNetWorkDesc,
			prometheus.GaugeValue,
			source,
			host,
		)
	}

	totalInfo := re.getApiTotalInfo()
	for host, source := range totalInfo {
		ch <- prometheus.MustNewConstMetric(
			re.ApiReqTotalDesc,
			prometheus.GaugeValue,
			source,
			host,
		)
	}
}

// 新建一个采集器
func NewApiExport(apiName string) *ApiExport {
	return &ApiExport{
		ApiName: apiName,
		ApiSourceDesc: prometheus.NewDesc(
			"api_source",                           // 指标的名称
			"help~~~~",                             // 帮助信息
			[]string{"host"},                       // label名称数组
			prometheus.Labels{"api_name": apiName}, // labels
		),
		ApiReqTotalDesc: prometheus.NewDesc(
			"api_total",
			"help~~~~",
			[]string{"host"},
			prometheus.Labels{"api_name": apiName},
		),
		ApiDeviceDesc: prometheus.NewDesc(
			"api_device",
			"help~~~~",
			[]string{"host"},
			prometheus.Labels{"api_name": apiName},
		),
		ApiBrowserDesc: prometheus.NewDesc(
			"api_browser",
			"help~~~~",
			[]string{"host"},
			prometheus.Labels{"api_name": apiName},
		),
		ApiNetWorkDesc: prometheus.NewDesc(
			"api_network",
			"help~~~~",
			[]string{"host"},
			prometheus.Labels{"api_name": apiName},
		),
	}

}

func main() {
	order_api := NewApiExport("order_info")
	// 注册一个采集器
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(order_api)

	gatherers := prometheus.Gatherers{reg}

	// 给采集器提供一个http的访问方式
	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})

	// 访问路径也可以是其他的
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	log.Infoln("Start server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error occur when start server %v", err)
		os.Exit(1)
	}
}


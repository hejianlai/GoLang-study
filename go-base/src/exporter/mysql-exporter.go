package main
import (
	"database/sql"
	`github.com/prometheus/client_golang/prometheus`
	`github.com/prometheus/client_golang/prometheus/promhttp`
	`github.com/prometheus/common/log`
	`net/http`
	`os`

	_ "github.com/go-sql-driver/mysql"
)
type ApiExport struct {
	ApiName string
	ApiSourceDesc *prometheus.Desc
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func Getinfo() int {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)
	// 查询数据
	rows, err := db.Query("SELECT count(*) FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		// fmt.Println(count)
		return count
	}
	 db.Close()

	return 0
}
func (re *ApiExport) getApiSourceInfo() map[string]float64 {
	var a = Getinfo()
	sourceInfo := map[string]float64{
		"127.0.0.1": float64(a),
	}
	return sourceInfo
}
func (re *ApiExport) Describe(ch chan<- *prometheus.Desc)  {
	ch <- re.ApiSourceDesc

}
func (re *ApiExport) Collect(ch chan<- prometheus.Metric)  {
	sourceInfo := re.getApiSourceInfo()
	for host,source := range sourceInfo{
		ch <- prometheus.MustNewConstMetric(
			re.ApiSourceDesc,
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
			"userinfo_count",                           // 指标的名称
			"help~~~~",                             // 帮助信息
			[]string{"host"},                       // label名称数组
			prometheus.Labels{"api_name": apiName}, // labels
		),
	}

}

func main() {
	order_api := NewApiExport("table_userinfo_count")
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

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-04 22:47:22
 */
package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "myapp",
		Subsystem: "gin",
		Name:      "http_request_counter",
		Help:      "接口请求统计",
	}, []string{
		"method", "path",
	})

	HttpRquestTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "myapp",
		Subsystem: "gin",
		Name:      "http_request_use_time",
		Help:      "接口请求相应时间统计",
	}, []string{
		"method", "path",
	})
)

func init() {
	prometheus.MustRegister(HttpRequestCounter)
	prometheus.MustRegister(HttpRquestTime)
}

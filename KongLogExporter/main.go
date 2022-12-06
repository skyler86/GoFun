package main

import (
	"KongLogExporter/config/sys"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// _ "go.uber.org/automaxprocs" //由于golang程序要在容器中运行，所以不能被cgroup限制cpu，因此需要再main文件中 import 这个包。
)

//登录验证变量
type UserAuth struct {
	EsUrl    string `json:"url"`
	EsName   string `json:"name"`
	EsPasswd string `json:"passwd"`
}

//模板变量
type KongLogTime struct {
	Klt_startTime string `json:date`
	Klt_endTime   string `json:date`
}

//指标变量
type IndexLabel struct {
	StatusCode   string `json: statuscode`
	DomainName   string `json: domainname`
	ResourcePath string `json: resourcepath`
	ErrorCount   int    `json: errorcount`
}

func GetKongLog(w http.ResponseWriter, r *http.Request) {

	var ua UserAuth
	env_url, env_name, env_password := sys.GetEsEnv()

	ua.EsUrl = env_url
	ua.EsName = env_name
	ua.EsPasswd = env_password
	//获取当前时间
	nowTime := time.Now()
	//当前时间减5分钟
	before5m := nowTime.Add(-5 * time.Minute)
	//当前时间减8小时
	before8h, _ := time.ParseDuration("-8h")
	//获取减8小时零5分钟后的时间
	minusTime8h5m := before5m.Add(before8h)
	minusTime8h := nowTime.Add(before8h)
	// fmt.Println("调整当前时间减8小时零5分钟的时间为:", minusTime8h5m)
	// fmt.Println("调整当前时间减8小时的时间为:", minusTime8h)

	//时间格式化
	startTime := minusTime8h5m.Format("2006-01-02T15:04:05")
	endTime := minusTime8h.Format("2006-01-02T15:04:05")
	// fmt.Println("起始时间：", startTime)
	// fmt.Println("结束时间：", endTime)

	//给模板变量赋值
	k := KongLogTime{
		Klt_startTime: startTime,
		Klt_endTime:   endTime,
	}

	//解析模板
	tpl, err := template.ParseFiles("./config/json/kongQuery.json")
	if err != nil {
		log.Fatal(err)
	}

	//渲染输出
	var ts bytes.Buffer
	err = tpl.Execute(&ts, k)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("User JSON:%v", ts.String())
	// jsonbody := ts.String()
	jsonbody := ts.Bytes()
	// fmt.Printf("User JSON:%v", jsonbody)

	//创建临时存储文件
	jsonFile, error := ioutil.TempFile("./config/json/", "ES_KongReq-*.json")

	if error != nil {
		fmt.Println("创建文件失败")
		// return
	}

	//利用file指针的Write()，将模板文件的内容写入到临时存储文件
	// jsonFile.WriteString(jsonbody)
	if _, err := jsonFile.Write(jsonbody); err != nil {
		jsonFile.Close()
		log.Fatal(err)
	}
	if err := jsonFile.Close(); err != nil {
		log.Fatal(err)
	}

	//通过io读取临时存储文件内容赋值到指定变量
	jsonFileData, _ := ioutil.ReadFile(jsonFile.Name())
	// fmt.Println(string(jsonFileData))

	//销毁临时文件
	defer os.Remove(jsonFile.Name())

	//byte转io.Reader
	body := ioutil.NopCloser(bytes.NewReader(jsonFileData))

	//发送请求
	res, _ := http.NewRequest("GET", ua.EsUrl, body)
	//设置请求头
	res.Header.Set("Content-Type", "application/json")
	//设置请求认证
	res.SetBasicAuth(ua.EsName, ua.EsPasswd)
	//发送请求
	client := http.Client{}
	resp, _ := client.Do(res)
	// fmt.Println("resp: ", resp)
	// fmt.Println("resp.Body: ", resp.Body)

	//读取响应内容
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	// fmt.Println(string(jsonBytes))

	js, err := simplejson.NewJson(jsonBytes)
	if err != nil {
		fmt.Printf("%v", err)
		// return
	}

	var il IndexLabel
	buckets, err := js.Get("aggregations").Get("2").Get("buckets").Array()
	// fmt.Println(buckets)

	//声明注册变量
	var registryLocal = prometheus.NewRegistry()

	//主逻辑开始，遍历buckets数组
	for x, _ := range buckets {
		fieldData := js.Get("aggregations").Get("2").Get("buckets").GetIndex(x)
		il.StatusCode = fieldData.Get("key").MustString()
		// fmt.Printf("状态码=%s\n",il.StatusCode)

		//创建非固定 label 的 gauge 指标对象
		kong_http_code_ := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "kong_http_code_" + il.StatusCode,
			Help: "Get Kong Log Erron Info",
		}, []string{"host", "uri"})

		//将指标对象注册到全局默认注册表中
		registryLocal.MustRegister(kong_http_code_)

		uri_Domain, _ := fieldData.Get("3").Get("buckets").Array()
		for y, _ := range uri_Domain {
			domain := fieldData.Get("3").Get("buckets").GetIndex(y)
			il.DomainName = domain.Get("key").MustString()
			// fmt.Printf("域名=%s\n",il.DomainName)

			uri_Info, _ := domain.Get("4").Get("buckets").Array()
			for z, _ := range uri_Info {
				path := domain.Get("4").Get("buckets").GetIndex(z)
				// fmt.Println(path)
				il.ResourcePath = path.Get("key").MustString()
				il.ErrorCount = path.Get("doc_count").MustInt()
				// fmt.Printf("资源路径=%s\n",il.ResourcePath)
				// fmt.Printf("错误统计=%d\n",il.ErrorCount)
				// fmt.Printf("主机域名: %s 的资源路径: %s 出现 HTTP_Code_%s 的错误数量是: %d\n",il.DomainName,il.ResourcePath,il.StatusCode, il.ErrorCount)

				//格式化字符串
				// message := fmt.Sprintf("主机域名: %s 的资源路径: %s 出现 HTTP_Code_%s 的错误数量是: %d\n", il.DomainName, il.ResourcePath, il.StatusCode, il.ErrorCount)
				// w.Write([]byte(message))

				var count float64 = float64(il.ErrorCount)
				// 针对不同标签值设置不同的指标值
				kong_http_code_.WithLabelValues(il.DomainName, il.ResourcePath).Set(count)
			}
		}
	}
	//主逻辑结束
	h := promhttp.HandlerFor(registryLocal, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

//部署到k8s时需要有健康检查
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {

	//暴露自定义的指标
	http.HandleFunc("/metrics", GetKongLog)
	http.HandleFunc("/health_check", healthCheck)
	//开启监听
	http.ListenAndServe(":9099", nil)

}

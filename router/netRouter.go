package netRouter

import (
	"html/template"
	"log"
	"net/http"
	"lab-server-repair-go/tool/yaml"
)

func netRouterConfigureCenter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fallthrough
	default:
		tpl, gloableErr := template.ParseFiles("html/index.html")
		if gloableErr != nil {
			http.NotFound(w, r)
			return
		}
		tpl.Execute(w, nil)
	}
}

func NetResponseHandler() {
	reader, configureErr := yamlReader.Instance()
	if configureErr != nil {
		log.Fatal("yaml文件配置失败: ", configureErr, "code: ", configureErr.Error())
	}

	// downloadPath := "/" + reader.Configure.Xlsx.DownloadFile + "/"
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	// http.Handle(downloadPath, http.StripPrefix(downloadPath, http.FileServer(http.Dir("file/xlsx"))))

	http.HandleFunc("/", netRouterConfigureCenter)

	err := http.ListenAndServe(":" + reader.Configure.Port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err, "code: ", err.Error())
	}
}

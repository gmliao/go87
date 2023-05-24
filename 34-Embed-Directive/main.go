package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

/*
	34 Embed Directive
	使用golang Embed Directive 去載入json的優點是可以在編譯時將檔案內容嵌入二進位檔案中，
 	不需要在執行時載入外部檔案，因此可以減少應用程式的大小和運行時的開銷。
 	此外，這種方法還可以確保應用程式的安全性，因為它可以防止攻擊者修改或替換外部檔案。
 	缺點是如果json檔案改動了，需要重新編譯原始碼才能更新內容。
	可以應用在系統設定檔案、靜態網頁檔案、靜態圖片檔案等等。
	不過還是可以透過分析檔案的內容來取得檔案的資訊，因此不適合用來儲存機密資訊。
*/

//go:embed config.json
var jsonConfig string

type Config struct {
	DbHost string `json:"dbHost"`
	DbPort int    `json:"dbPort"`
	DbName string `json:"dbName"`
	DbUser string `json:"dbUser"`
	DbPass string `json:"dbPass"`
}

func main() {

	config := Config{}

	json.Unmarshal([]byte(jsonConfig), &config)
	fmt.Printf("%+v\n", config)

}

package test

import (
	"net/http"
	"testing"
)

const pass = "\u2713" //打勾
const fail = "\u2717" //打X

type Url struct {
	url        string
	statuscode int
}

func TestDownload(t *testing.T) {
	urls := []Url{
		Url{
			url:        "http://www.goinggo.net/index.xml",
			statuscode: http.StatusOK,
		},
		Url{
			url:        "http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			statuscode: http.StatusNotFound,
		},
	}

	t.Log("测试下载功能")
	{
		for i, url := range urls {
			resp, err := http.Get(url.url)
			if err != nil {
				t.Error("测试第", i+1, "个url,测试结果:", fail, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == url.statuscode {
				t.Log("测试第", i+1, "个url,测试结果:", pass)
			} else {
				t.Error("测试第", i+1, "个url,测试结果:", fail, "期待结果:", url.statuscode, "实际结果:", resp.StatusCode)
			}
		}
	}
}

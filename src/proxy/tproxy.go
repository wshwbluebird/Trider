package proxy

import(
	"net/http"
	"net/url"
	"fmt"
	"time"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"errors"
)




func NewPoxyer() *Proxyer{
	return &Proxyer{client:http.Client{}}
}

func GetTransportByStr(proxyUrl string) *http.Transport{
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	return &http.Transport{Proxy: proxy}
}

func GetUrlString(host string, port string) string{
	return fmt.Sprintf("http://%s:%s",host,port)
}


func GetUsefulProxy(number int) ([]string, error){
	ans := []string{}
	client := &http.Client{}
	base := "http://www.xicidaili.com/nn/"
	proxyer := NewPoxyer()
	full := false


	for  i := 1 ; i <= number ; i++ {
		path := fmt.Sprintf("%s%d",base,i)
		fmt.Println(path)
		req, _ := http.NewRequest("GET",path,nil)

		req.Header.Set("User-Agent",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/66.0.3359.181 Safari/537.36")

		resp, err :=client.Do(req)
		if err != nil{
			return nil, err
		}
		defer resp.Body.Close()

		node, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil{
			return nil, err
		}
		ip_list := node.Find("table").Find("tbody")

		ip_list.Find("tr").Each(
			func(index int, tr *goquery.Selection) {
			host := tr.Find("td").Eq(1).Text()
			port := tr.Find("td").Eq(2).Text()
			proxyUrl := GetUrlString(host,port)
			//fmt.Printf("Test ip %s\n",proxyUrl)
			if !full && proxyer.isUseful(proxyUrl)   {
				fmt.Printf("find: %s\n",proxyUrl)
				ans = append(ans, proxyUrl)
				if len(ans) == number {
					full = true
				}
			}
		})

		if full {
			return ans, nil
		}

	}

	return nil, errors.New("no enough ip")
}


type Proxyer struct {
	client http.Client
}

func (p *Proxyer) isUseful(proxyUrl string) bool{
	p.client.Timeout = time.Duration(1000 * time.Millisecond)
	p.client.Transport = GetTransportByStr(proxyUrl)
	resp ,err := p.client.Get("http://ip.chinaz.com/getip.aspx")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body , erre := ioutil.ReadAll(resp.Body)
	if erre != nil {
		return  false
	}
	str:= string(body)
	if str[0] == '{'{
		return true
	}
	return false

}






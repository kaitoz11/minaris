package shooting

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

type Gun struct {
	name  string
	color string

	facts map[string]string

	ammos []*Request

	client *http.Client
}

func (g *Gun) AddFacts(facts map[string]string) {
	factMap := make(map[string]string)
	if g.facts == nil {
		g.facts = factMap
	}
	for k, v := range facts {
		g.facts[k] = v
	}
}

func (g *Gun) AddRawFacts(facts []string) {
	factMap := make(map[string]string)
	if g.facts == nil {
		g.facts = factMap
	}
	for _, line := range facts {
		keyValuePair := strings.SplitN(line, "=", 2)
		fmt.Println(keyValuePair)
		g.facts[keyValuePair[0]] = keyValuePair[1]
	}
}

func (g *Gun) PrintFacts() string {
	facts := ""
	for f, v := range g.facts {
		facts += f + "=" + v + "\n"
	}
	return fmt.Sprintf("---\n%s\n", facts)
}

// Send requests
func (g *Gun) Shoot() {
	if g.client == nil {
		return
	}

	facts := make(map[string]string)
	if g.facts != nil {
		facts = g.facts
	}

	for _, req := range g.ammos {
		// subtitute
		for originalInput, modified := range req.Input {
			re := regexp.MustCompile(`{{([A-Z0-9_]+)}}`)
			groups := re.FindStringSubmatch(modified)
			// fmt.Printf("groups: %s\n", groups)
			if len(groups) > 1 {
				modified = strings.Replace(modified, "{{"+groups[1]+"}}", facts[groups[1]], 1)
			}
			// fmt.Printf("mod: %s\n", modified)
			newRaw := strings.Replace(req.Raw, originalInput, modified, 1)
			req.Raw = newRaw
			err := req.LoadRequest()
			if err != nil {
				fmt.Println(err)
				break
			}

			// fmt.Printf("NewRaw: %s\n", newRaw)
		}

		resp, err := req.SendRequest(g.client)
		if err != nil {
			fmt.Println(err)
			break
		}
		defer resp.Body.Close()

		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		result := string(b)

		// fmt.Println("\n===")
		// fmt.Println(result)

		// Extract from response
		fmt.Println("\n====")
		for name, pattern := range req.Output {
			re := regexp.MustCompile(pattern)
			// subt := re.FindString(result)
			subt := re.FindStringSubmatch(result)
			if len(subt) <= 1 {
				fmt.Printf("[-] format err: %s\n", subt)
				break
			}
			facts[name] = subt[1]
			fmt.Printf("%s: %s\n", name, facts[name])
		}
		fmt.Println("====")
		// fmt.Println(facts)
	}
	g.facts = facts
}

func (g *Gun) ShootFlow() {
	if g.client == nil {
		return
	}

}

func (g *Gun) Load(ammos []*Request) {
	g.ammos = ammos
}

func (g *Gun) LoadAmmos(raws map[string]string) {
	ammos := make([]*Request, 0)
	for _, v := range raws {
		// fmt.Printf("key: %s, val: \n%s\n", k, v)
		rawHttpString := v

		req, err := NewFromRaw(rawHttpString)
		if err != nil {
			fmt.Println(err)
		}

		ammos = append(ammos, req)
		g.ammos = ammos
	}
}

// Initialize
func (g *Gun) Init(proxyURL string) {
	g.setEngine(proxyURL)
}

// set proxy
func (g *Gun) setEngine(proxyURL string) {
	transport := http.Transport{}
	client := &http.Client{}
	if proxyURL != "" {
		url_proxy, _ := url.Parse(proxyURL)
		transport.Proxy = http.ProxyURL(url_proxy)                        // set proxy
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl
		client.Transport = &transport
	}

	g.client = client
}

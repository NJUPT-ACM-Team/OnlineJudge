package vjudger

import (
	"OnlineJudge/judger"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	// "os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MAX_WaitTime = 120
)

type HDUJudger struct {
	client   *http.Client
	token    string
	pat      *regexp.Regexp
	username string
	userpass string
}

const HDUToken = "HDU"

var HDURes = map[string]string{
	"Queuing":                                   "wt",
	"Compiling":                                 "wt",
	"Running":                                   "wt",
	"Compilation Error":                         "ce",
	"Accepted":                                  "ac",
	"Runtime Error<br>(STACK_OVERFLOW)":         "re",
	"Runtime Error<br>(ACCESS_VIOLATION)":       "re",
	"Runtime Error<br>(ARRAY_BOUNDS_EXCEEDED)":  "re",
	"Runtime Error<br>(FLOAT_DENORMAL_OPERAND)": "re",
	"Runtime Error<br>(FLOAT_DIVIDE_BY_ZERO)":   "re",
	"Runtime Error<br>(FLOAT_OVERFLOW)":         "re",
	"Runtime Error<br>(FLOAT_UNDERFLOW )":       "re",
	"Runtime Error<br>(INTEGER_OVERFLOW)":       "re",
	"Runtime Error<br>(INTEGER_DIVIDE_BY_ZERO)": "re",
	"Wrong Answer":                              "wa",
	"Time Limit Exceeded":                       "tle",
	"Memory Limit Exceeded":                     "mle",
	"Output Limit Exceeded":                     "ole",
	"Presentation Error":                        "pe",
	"System Error":                              "se"}

var HDULang = map[string]int{
	"c":    1,
	"cpp":  0,
	"java": 5}

var HDUSearchLang = map[string]int{
	"c":    2,
	"cpp":  1,
	"java": 6}

func (h *HDUJudger) Init(_ judger.JudgerInterface) error {
	jar, _ := cookiejar.New(nil)
	h.client = &http.Client{Jar: jar, Timeout: time.Second * 10}
	h.token = HDUToken
	pattern := `(\d+)</td><td>(.*?)</td><td>(?s:.*?)<font color=.*?>(.*?)</font>.*?<td>(\d+)MS</td><td>(\d+)K</td><td><a href="/viewcode.php\?rid=\d+"  target=_blank>(\d+) B</td><td>.*?</td>`

	h.pat = regexp.MustCompile(pattern)
	h.username = "nojvj01"
	h.userpass = "nojvjvj"
	return nil
}

func (h *HDUJudger) Match(token string) bool {
	if strings.ToLower(token) == strings.ToLower(HDUToken) {
		return true
	}
	return false
}
func (h *HDUJudger) Login(_ judger.JudgerInterface) error {

	h.client.Get("http://acm.hdu.edu.cn")

	uv := url.Values{}
	uv.Add("username", h.username)
	uv.Add("userpass", h.userpass)
	uv.Add("login", "Sign In")

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/userloginex.php?action=login", strings.NewReader(uv.Encode()))
	if err != nil {
		return BadInternet
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println("err", err)
		return BadInternet
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "No such user or wrong password.") >= 0 {
		return LoginFailed
	}

	return nil
}

func (h *HDUJudger) Submit(u judger.JudgerInterface) error {

	uv := url.Values{}
	uv.Add("check", "0")
	uv.Add("problemid", u.GetOJPid())
	uv.Add("language", strconv.Itoa(HDULang[u.GetLanguage().GetLang()]))
	uv.Add("usercode", u.GetCode())

	req, err := http.NewRequest("POST", "http://acm.hdu.edu.cn/submit.php?action=submit", strings.NewReader(uv.Encode()))
	if err != nil {
		return BadInternet
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "acm.hdu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.91 Safari/537.36")

	resp, err := h.client.Do(req)
	if err != nil {
		return BadInternet
	}
	defer resp.Body.Close()
	// u.SetSubmitTime(time.Now())

	b, _ := ioutil.ReadAll(resp.Body)
	html := string(b)
	if strings.Index(html, "No such problem") >= 0 {
		return NoSuchProblem
	}
	if strings.Index(html, "One or more following ERROR(s) occurred.") >= 0 {
		return SubmitFailed
	}
	return nil
}

func (h *HDUJudger) GetStatus(u judger.JudgerInterface) error {

	statusUrl := "http://acm.hdu.edu.cn/status.php?first=&" +
		"pid=" + u.GetOJPid() +
		"&user=" + h.username +
		"&lang=" + strconv.Itoa(HDUSearchLang[u.GetLanguage().GetLang()]) + "&status=0"

	endTime := time.Now().Add(MAX_WaitTime * time.Second)

	for true {
		if time.Now().After(endTime) {
			return BadStatus
		}
		resp, err := h.client.Get(statusUrl)
		if err != nil {
			return BadInternet
		}
		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		AllStatus := h.pat.FindAllStringSubmatch(string(b), -1)
		// log.Println(AllStatus)

		layout := "2006-01-02 15:04:05 (MST)" //parse time
		for i := 0; i < len(AllStatus); i++ {
			status := AllStatus[i]
			// log.Println(status)
			t, _ := time.Parse(layout, status[2]+" (CST)")
			t = t.Add(20 * time.Second) //HDU server's time is less 36s.
			// log.Println(t, u.GetSubmitTime())
			// log.Println(status[1:])
			if t.After(u.GetSubmitTime()) {
				rid := status[1] //remote server run id
				sc := HDURes[status[3]]
				if err := u.UpdateResult(status[3], sc); err != nil {
					// LOG
				}

				if sc != "wt" {
					if sc == "ce" {
						CE, err := h.GetCEInfo(rid)
						if err != nil {
							// log.Println(err)
						}
						if err := u.UpdateCEInfo(CE); err != nil {
							// LOG
						}
					}

					Time, _ := strconv.Atoi(status[4])
					Mem, _ := strconv.Atoi(status[5])
					// Length, _ := strconv.Atoi(status[6])
					u.UpdateResource(Time, Mem)
					return nil
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (h *HDUJudger) GetCEInfo(rid string) (string, error) {
	resp, err := h.client.Get("http://acm.hdu.edu.cn/viewerror.php?rid=" + rid)
	if err != nil {
		log.Println(err)
		return "", BadInternet
	}

	b, _ := ioutil.ReadAll(resp.Body)
	pre := "(?s)<pre>(.*?)</pre>"
	re := regexp.MustCompile(pre)
	match := re.FindStringSubmatch(string(b))
	return match[1], nil
}

func (h *HDUJudger) Run(u judger.JudgerInterface) error {
	for _, apply := range []func(judger.JudgerInterface) error{h.Init, h.Login, h.Submit, h.GetStatus} {
		if err := apply(u); err != nil {
			// log.Println(err)
			return err
		}
	}
	return nil
}

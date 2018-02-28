package main

import (
	"github.com/go-vgo/robotgo"
	"fmt"
	"github.com/kr/pretty"
)

const (
	XOFFSET = 30
	YOFFSET = 20
)

func main() {
	fpid, err := robotgo.FindIds("iexplore.exe")
	if err == nil {
		fmt.Printf("%# v", pretty.Formatter(fpid))
	}
	robotgo.ActiveName("iexplore.exe")
	bitmaps := robotgo.OpenBitmap("targetRegion.png")
	bitmapBG := robotgo.OpenBitmap("bg.png")
	bitmap1 := robotgo.OpenBitmap("1.png")
	fxs, fys := robotgo.FindBitmap(bitmaps, bitmapBG)
	fmt.Println("FindBitmap------", fxs, fys)
	x1, y1 := robotgo.FindBitmap(bitmap1, bitmaps)
	fmt.Println("FindBitmap------", fxs+x1, fys+y1)
	bitmap2 := robotgo.OpenBitmap("2.png")
	x2, y2 := robotgo.FindBitmap(bitmap2, bitmaps)
	fmt.Println("FindBitmap------", fxs+x2, fys+y2)
	bitmap3 := robotgo.OpenBitmap("3.png")
	x3, y3 := robotgo.FindBitmap(bitmap3, bitmaps)
	fmt.Println("FindBitmap------", fxs+x3, fys+y3)
	bitmaps2 := robotgo.OpenBitmap("tg2.png")
	bitmapBG2 := robotgo.OpenBitmap("bg2.png")
	bitmap21 := robotgo.OpenBitmap("4.png")
	fxs2, fys2 := robotgo.FindBitmap(bitmaps2, bitmapBG2)
	fmt.Println("FindBitmap------", fxs2, fys2)
	x21, y21 := robotgo.FindBitmap(bitmap21, bitmaps2)
	fmt.Println("FindBitmap------", fxs2+x21, fys2+y21)
	keve := robotgo.AddEvent("space")
	if keve == 0 {

		robotgo.Move(fxs+x1+XOFFSET, fys+y1+YOFFSET)
		robotgo.MouseClick("left", true)
		robotgo.KeyTap("6")
		robotgo.KeyTap("0")
		robotgo.KeyTap("0")

		robotgo.Move(fxs+x2+XOFFSET, fys+y2+YOFFSET)
		robotgo.Click("left")

		robotgo.Move(fxs+x3+XOFFSET, fys+y3+YOFFSET)
		robotgo.Click("left")
		keves := robotgo.AddEvent("enter")
		if keves == 0 {
			robotgo.Move(fxs2+x21+XOFFSET, fys2+y21+YOFFSET)
			robotgo.Click("left")
		}
	}
	/*bitmapcj := robotgo.OpenBitmap("test3.png")
	fmt.Println("...", bitmapcj)

	fxcj, fycj := robotgo.FindBitmap(bitmapcj)

	bitmap := robotgo.OpenBitmap("test.png")
	fmt.Println("...", bitmap)
	robotgo.Move(fxcj + 10, fycj + 10)
	robotgo.Click("left")

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------", fx, fy)
	keve := robotgo.AddEvent("enter")
	if keve == 0 {
		robotgo.Move(fx + 10, fy + 10)
		robotgo.Click("left")
	}*/
}

/*import (
	"fmt"
	"github.com/kr/pretty"
	"git.derbysoft.tm/warrior/derbysoft-common-go.git/util"
	"git.derbysoft.tm/warrior/ccs-client-go.git/sample/sample_use_inject/mapping"
	"reflect"
)

func main() {
	test := &mapping.Account{}
	test2 := false
	test3 := make(map[string]interface{})
	test3["a"] = test2
	fmt.Println(reflect.ValueOf(test3["a"]).Type().)
	testInterface(test)
}

func testInterface(t interface{}) {
	a := util.GetAttributesMap(t)
	fmt.Printf("%# v", pretty.Formatter(a))
}*/

/*
const IHG = "IHG"

const SEPERATOR = "|"

func main() {



}
*/

/*
var Pool redis.Pool

var retryTimes = 0

//TODO Init eRROR FATAL & ADD a switch to turn on/off
func InitPool(conf config.CacheConfig) {
	Pool = redis.Pool{
		MaxActive:   conf.MaxActive,
		MaxIdle:     conf.MaxIdle,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				conf.Url,
				redis.DialConnectTimeout(time.Duration(conf.DialTimeout) * time.Millisecond),
				redis.DialReadTimeout(time.Duration(conf.ReadTimeout) * time.Millisecond),
				redis.DialWriteTimeout(time.Duration(conf.WriteTimeout) * time.Millisecond),
			)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Wait: true,
	}
	log.Info("init cache success")

}

func GetConnection() redis.Conn {
	return Pool.Get()
}

func PutConnection(conn redis.Conn) {
	conn.Close()
}


func main() {
	config := config.CacheConfig{
		Url                 :"54.190.11.182:6379",
		IdleTimeout         :1000,
		DialTimeout         :1000,
		ReadTimeout         :1000,
		WriteTimeout        :1000,
		MaxActive           :50,
		MaxIdle             :50,
		KeyTimeout          :50,
		Processor           :50,
		ProcessorRetryTimes :50,
	}
	InitPool(config)
	con := GetConnection()
	a := `{"CurrencyCode":"CNY","GuestCount":{"AdultCount":1,"ChildAges":"8,9","ChildCount":2,"TPAExtensions":null},"HotelCode":"CP-CANHL","RatePlans":{"RatePlan":[{"CancelPolicy":{"CancelPenalties":null,"Description":"will result in cancellation of the remainder of your reservation.","NonRefundable":true,"TPAExtensions":null},"RatePlanCode":"RATEPLAN01","RatePlanDescription":"RATEPLAN01","RatePlanName":"RATEPLAN01"}],"TPAExtensions":null},"RoomRates":{"RoomRate":[{"CancelPolicy":null,"Fees":{"Fee":[{"Amount":10,"ChargeType":"TaxP","Description":"TaxP","EffectiveDate":null,"ExpireDate":null,"Percent":null,"TPAExtensions":null,"Type":"Exclusive","Unit":"PER_ROOM_PER_NIGHT"}]},"MealPlanCode":"RO","RatePlanCode":"RATEPLAN01","Rates":{"Rate":[{"AmountAfterTax":1200,"AmountBeforeTax":1190,"EffectiveDate":"2018-01-07","ExpireDate":"2018-01-08","ExtraPersonCharge":0,"ExtraPersonChargeAfterTax":0,"TPAExtensions":null},{"AmountAfterTax":1300,"AmountBeforeTax":1290,"EffectiveDate":"2018-01-08","ExpireDate":"2018-01-09","ExtraPersonCharge":0,"ExtraPersonChargeAfterTax":0,"TPAExtensions":null}]},"RoomTypeCode":"ROOMA","TPAExtensions":null}],"TPAExtensions":null},"RoomTypes":{"RoomType":[{"RoomTypeCode":"ROOMA","RoomTypeDescription":"ROOMA","RoomTypeName":"KING FEATURE NONSMOKING"}],"TPAExtensions":null},"TPAExtensions":null}`
	con.Do("SET", "2018-01-07|2018-01-14|CNBOOKING|CP-CANHL|3|0|11", a)
	fmt.Println(a)
	PutConnection(con)
}
*/

/*package main

import (
	"fmt"
	"regexp"
)

*//*const (
	MIN_TOKEN_COUNT = 3
)

type Base struct {
	A string
}

func (h *Base) Test() ([]byte, error) {
	return json.Marshal(h)
}

type BaseEx struct {
	Base
	Aa string
}

func (h *BaseEx) Test() ([]byte, error) {
	return h.Base.Test()
}

type Strs struct {
	A ss `json:","`
}

type ss []string

var specialCases []rune = []rune{45, 47, 46, 61, 39, 65293, 12289, 12290, 65309, 8216, 8217}

type Text []byte*//*

var len int
var gap int

type A map[string]string

func main() {

	var t A
	t = make(map[string]string)

	t["a"] = "b"
	a := make([]byte, 10)
	pat := "\\w*-_G.[jpg|png]"
	src := "http://www.cfmedia.vfmleonardo.com/imageRepo/5/0/85/981/124/wes1513ag-174911-Grand_20Staircase-_G.jpg"
	fmt.Println(regexp.MatchString(pat, src))
	*//*fmt.Println("Please enter your args, lenght of the array: ")
	fmt.Scanln(&len)
	fmt.Println("Please enter your args, gap: ")
	fmt.Scanln(&gap)
	for i := 0 ; i < 1024; i ++ {


	}*//*
	//var ta string = ""
	//var tb *string = &ta
	//tc := []string{"a", "b", "c", "d"}

	fmt.Println(a)

	*//*jsonStr := "{\"userName\":\"public\", \"password\":\"public@derby\"}"
	rs, _, _, _ := common.PostRequestJSON("http://10.200.151.6:8081/v2/user/login", strings.NewReader(string(jsonStr)), "")
	fmt.Print(string(rs))*//*

	*//*tc := make(chan interface{}, 1)
	tc <- new(BaseEx)

	v := <- tc
	switch b := v.(type) {
	case *BaseEx:
		b.A = "a"
		b.Aa = "aa"
		bb, _ := b.Test()
		fmt.Print(string(bb))
	default:
		fmt.Print("Wrong")
	}*//*

	*//*	strs := new(Strs)
		strs.A = make([]string, 0)
		//strs.A = append(strs.A, "abc")
		byts, _ := json.Marshal(strs)
		fmt.Printf("%s\r\n", byts)

		*//*/*t := reflect.TypeOf(*strs)
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("%# v", pretty.Formatter(t.Field(i)))
			s := reflect.StructTag{}("abcd")

			t.Field(i).Tag = s
		}*//*/*
		fmt.Printf("%s\r\n", byts)*//*

}*/

/*func main() {
	ch := make(chan string, 1000)
	returnQ := make(chan string, 1000)
	for i := 0; i < 10; i++ {
		go func(num int) {
			for {
				select {
				case val := <-ch:
					fmt.Println(num)
					//time.Sleep(100*time.Millisecond)
					ch <- val
				default:
					continue
				}
			}
		}(i)
	}

	for i := 0; i < 100; i++ {
		ch <- "hello"
	}
	for {
		select {
		case <-returnQ:
			continue
			//fmt.Println("back")
		}
	}
	//v, ok := <-ch        // 变量v的值为空字符串""，变量ok的值为false
	*//*fmt.Println(SpiltToTokenData("athens, grace", true))
	fmt.Println(SplitToText("athens,g"))
	r, size := utf8.DecodeRune([]byte("￥"))

	fmt.Println(r, " " ,size, unicode.IsSymbol(r))
*//*
	*//*	t := new(BaseEx)
		t.Aa = "aaaaaa"
		rs, _ := t.Test()
		fmt.Println(string(rs))*//*
}*/

/*
func um(x interface{}, b []byte) (error) {
	err := xml.Unmarshal(b, x)
	return err
}

func SplitToText(text string) []string {
	trimTxt := strings.TrimSpace(text)
	textBytes := []byte(trimTxt)
	splitter := new(splitter)
	splitter.indexSet = utils.NewSet()
	texts := splitter.SplitTextToWords(textBytes, false)
	tokens := make([]string, 0)
	for _, row := range texts {
		*/
/*		r, _ := utf8.DecodeRune([]byte(row.Text))
				if isSpecialCase(r) {
					continue
				}*//*

		tokens = append(tokens, string(row.Text))
	}
	return tokens
}

func SpiltToTokenData(text string, isIndex bool) []types.TokenData {
	trimTxt := strings.TrimSpace(text)
	textBytes := []byte(trimTxt)
	splitter := new(splitter)
	splitter.indexSet = utils.NewSet()
	tokens := splitter.SplitTextToWords(textBytes, isIndex)
	return tokens
}

type splitter struct {
	current              int
	inAlphanumeric       bool
	alphanumericStart    int
	minAlphanumericCount int
	indexSet             utils.Set
}

func (s *splitter) SplitTextToWords(text Text, isIndex bool) []types.TokenData {
	output := make([]types.TokenData, 0, len(text)/3)
	for s.current < len(text) {
		r, size := utf8.DecodeRune(text[s.current:])
		if s.isAlpha(r, size) {
			output = s.processAlpha(r, size, text, output, isIndex)
		} else {
			output = s.processNotAlpha(r, size, text, output, isIndex)
		}
		s.current += size
	}

	if s.inAlphanumeric && s.current != 0 && (s.minAlphanumericCount >= MIN_TOKEN_COUNT || isIndex) {
		output = s.tokenAppend(output, types.TokenData{
			Text:      string(toLower(text[s.alphanumericStart:s.current])),
			Locations: []int{s.alphanumericStart},
		})
	}
	return output
}

//Judge the Rune is English word's or not, the min count of the letters is MIN_TOKEN_COUNT.
func (s *splitter) isAlpha(r rune, size int) bool {
	return (size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r))) || isSpecialCase(r)
}

//Process function when the reading Rune is English letter or number.
func (s *splitter) processAlpha(r rune, size int, text Text, output []types.TokenData, isIndex bool) []types.TokenData {
	if !s.inAlphanumeric {
		s.alphanumericStart = s.current
		s.inAlphanumeric = true
	}
	s.minAlphanumericCount += 1
	if !isIndex {
		return output
	}
	if s.current != 0 {
		output = s.tokenAppend(output, types.TokenData{
			Text:      string(toLower(text[s.alphanumericStart:s.current])),
			Locations: []int{s.alphanumericStart},
		})
	}
	return output
}

//Process function when the reading Rune is not an letter, which is something like Chinese character, blank or anything else.
func (s *splitter) processNotAlpha(r rune, size int, text Text, output []types.TokenData, isIndex bool) []types.TokenData {
	if s.inAlphanumeric {
		s.inAlphanumeric = false
		if s.current != 0 {
			output = s.tokenAppend(output, types.TokenData{
				Text:      string(toLower(text[s.alphanumericStart:s.current])),
				Locations: []int{s.alphanumericStart},
			})
		}
	}
	if !isIndex {
		return output
	}
	output = s.tokenAppend(output, types.TokenData{
		Text:      string(text[s.current:s.current+size]),
		Locations: []int{s.current},
	})
	return output
}

func (s *splitter) tokenAppend(output []types.TokenData, data types.TokenData) []types.TokenData {
	if s.indexSet.Exists(data.Text) {
		return output
	}
	s.indexSet.Add(data.Text)
	return append(output, data)
}

func toLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}

func isSpecialCase(r rune) bool {
	for _, row := range specialCases {
		if r == row {
			return true
		}
	}
	return false
}
*/

/*
package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

*/

/*
package main

import (
	"fmt"
	"encoding/json"
	"runtime"
)

func main() {
	t := new(TestH)
	t.A = "test"
	r,_ := json.Marshal(t)
	fmt.Println(string(r))
	runtime.GOMAXPROCS(runtime.NumCPU())

}

type TestH struct {
	A string

}

*/

/*
import (
	"unicode/utf8"
	"unicode"
	"strings"
	"github.com/huichen/wukong/types"
	"fmt"
	_ "github.com/kr/pretty"
	"derbysoft.com/dlog/common/utils"
)

const (
	MIN_TOKEN_COUNT = 3
	BRAND_CODES_REGEX string = "^(intercontinental|kimpton|crowne plaza|hotel indigo|even hotel|holiday inn|holiday inn express" +
		"|holiday inn resort|holiday inn club vacations|staybridge suites|candlewood suites)" +
		"\\w*"
)

type Text []byte

type Test struct {
	A struct {
		  Tt string
	  }
	B string
}

func main() {

	*/
/*src := strings.ToLower("Crowne Plaza Docklands, London, United Kingdom")
	r, _ := regexp.Compile(BRAND_CODES_REGEX)
	src2 := r.ReplaceAllString(src, "")

	if re2, _ := regexp.MatchString(BRAND_CODES_REGEX, src); re2 {
		fmt.Println("a")
	}
	fmt.Println(src2)
	fmt.Println(r.FindString(src), strings.Index(BRAND_CODES_REGEX, r.FindString(src)))*//*

	fmt.Println(SpiltToTokenData("Albany", false))
}

func SplitToText(text string) []string {
	trimTxt := strings.TrimSpace(text)
	textBytes := []byte(trimTxt)
	splitter := new(splitter)
	splitter.indexSet = utils.NewSet()
	texts := splitter.SplitTextToWords(textBytes, false)
	tokens := make([]string, 0)
	for _, row := range texts {
		tokens = append(tokens, string(row.Text))
	}
	return tokens
}

func SpiltToTokenData(text string, isIndex bool) []types.TokenData {
	trimTxt := strings.TrimSpace(text)
	textBytes := []byte(trimTxt)
	splitter := new(splitter)
	splitter.indexSet = utils.NewSet()
	tokens := splitter.SplitTextToWords(textBytes, isIndex)
	*/
/*tokenData := make([]types.TokenData, 0)
	for index, row := range texts {
		data := types.TokenData{
			Text: string(row),
			Locations: []int{index},
		}
		tokenData = append(tokenData, data)
	}*//*

	return tokens
}

type splitter struct {
	current              int
	inAlphanumeric       bool
	alphanumericStart    int
	minAlphanumericCount int
	indexSet             utils.Set
}

func (s *splitter) SplitTextToWords(text Text, isIndex bool) []types.TokenData {
	output := make([]types.TokenData, 0, len(text) / 3)
	for s.current < len(text) {
		r, size := utf8.DecodeRune(text[s.current:])
		if s.isAlpha(r, size) {
			output = s.processAlpha(r, size, text, output, isIndex)
		} else {
			output = s.processNotAlpha(r, size, text, output, isIndex)
		}
		s.current += size
	}

	if s.inAlphanumeric && s.current != 0 {
		output = s.tokenAppend(output, types.TokenData{
			Text: string(toLower(text[s.alphanumericStart:s.current])),
			Locations: []int{s.alphanumericStart},
		})
	}
	return output
}


//Judge the Rune is English word's or not, the min count of the letters is MIN_TOKEN_COUNT.
func (s *splitter) isAlpha(r rune, size int) bool {
	return size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r))
}

//Process function when the reading Rune is English letter or number.
func (s *splitter) processAlpha(r rune, size int, text Text, output []types.TokenData, isIndex bool) []types.TokenData {
	if !s.inAlphanumeric {
		s.alphanumericStart = s.current
		s.inAlphanumeric = true
	}
	if s.minAlphanumericCount < MIN_TOKEN_COUNT {
		s.minAlphanumericCount += 1
		if !isIndex {
			return output
		}
	}
	if s.current != 0 {
		output = s.tokenAppend(output, types.TokenData{
			Text:  string(toLower(text[s.alphanumericStart:s.current])),
			Locations: []int{s.alphanumericStart},
		})
	}
	return output
}
*/

//Process function when the reading Rune is not an letter, which is something like Chinese character, blank or anything else.
/*func (s *splitter) processNotAlpha(r rune, size int, text Text, output []types.TokenData, isIndex bool) []types.TokenData {
	if s.inAlphanumeric {
		s.inAlphanumeric = false
		if s.current != 0 {
			output = s.tokenAppend(output, types.TokenData{
				Text: string(toLower(text[s.alphanumericStart:s.current])),
				Locations: []int{s.alphanumericStart},
			})
		}
	}
	output = s.tokenAppend(output, types.TokenData{
		Text: string(text[s.current:s.current + size]),
		Locations: []int{s.current},
	})
	return output
}*/
/*
func (s *splitter) tokenAppend(output []types.TokenData, data types.TokenData) []types.TokenData {
	if s.indexSet.Exists(data.Text) {
		return output
	}
	s.indexSet.Add(data.Text)
	return append(output, data)
}

func toLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}
/*
import (
	"fmt"
	"reflect"
	"github.com/kr/pretty"
)

type TestS struct {
	Aa string `json:"aa"`
	Bb string `json:"bb"`
}

func main() {
	t := TestS{}
	t.Aa = "a"
	t.Bb = "b"
	fmt.Printf("%# v \r\n", pretty.Formatter(GetAttributesMap(t)))
}

func GetAttributesMap(object interface{}) map[string]interface{} {
	val := reflect.ValueOf(object)//.Elem()
	attrs := make(map[string]interface{})
	fmt.Printf("%# v \r\n", pretty.Formatter(val))
	//reflect.ValueOf()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if !typeField.Anonymous {
			attrs[typeField.Tag.Get("json")] = valueField.Interface()
		}
	}
	return attrs
}
*/

/*package main

import (
	"net/url"
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
	"errors"
	"github.com/kr/pretty"
	"sync"
)

var session_keys = ""

var (
	//globalJarCookies map[string][]*http.Cookie = make(map[string][]*http.Cookie)
	globalJarCookies = &TestJar{perURL : make(map[string][]*http.Cookie)}
)

type TestJar struct {
	m      sync.Mutex
	perURL map[string]map[string]*http.Cookie
}

func (j *TestJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.m.Lock()
	defer j.m.Unlock()
	if j.perURL == nil {
		j.perURL = make(map[string][]*http.Cookie)
	}
	j.perURL[u.Host] = append(j.perURL[u.Host] ,cookies...)
}

func (j *TestJar) Cookies(u *url.URL) []*http.Cookie {
	j.m.Lock()
	defer j.m.Unlock()
	return j.perURL[u.Host]
}

func main() {
	//str := "username=public&password=public%40derby"
	rs, _, _ := GetRequest("http://52.43.244.90:18081/login")
	form := make(url.Values)
	form.Add("username", "public")
	form.Add("password", "public@derby")
	rs, _, _ = PostRequestJSON("http://52.43.244.90:18081/login", nil, form)
	fmt.Println(string(rs))
	rs, _, _ = PostRequestJSON("http://52.43.244.90:18081/api/add?client=facade&add?view_group_list=WARRIOR&view_api_name=aaa2&view_custom_domain=&view_listen_path=a1&view_selected_target_list=&view_target_url=a1&view_target_list=&view_option_authentication_mode=0&view_auth_key_header_name=Authorization&view_rate_limits_allow_requests=1000&view_rate_limits_per_seconds=60&view_quota_resets_every_text=&view_max_quota=-1&view_quota_resets_every=3599", nil, nil)
	fmt.Println(string(rs))
}*/

/*
func main() {
	form := make(url.Values)
	form.Add("username", "public")
	form.Add("password", "public@derby")
	req, _ := http.NewRequest("POST", "http://uatdoorway.derbysoftsec.com:10086/login", strings.NewReader(form.Encode()))
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
*/

/*
func PostRequestJSON(rawUrl string, body io.Reader, form url.Values) (rs []byte, isExpired bool, err error) {
	Url, err := url.Parse(rawUrl)
	if err != nil {
		return rs, false, err
	}
	Url.RawQuery = Url.Query().Encode()
	req, err := http.NewRequest(http.MethodPost, Url.String(), body)
	if err != nil {
		return rs, false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("HOST", "52.43.244.90:18081")
	req.PostForm = form
	//fmt.Println(Url.String())
	client := &http.Client{
		Jar : globalJarCookies,
	}
	//fmt.Printf("%# v", pretty.Formatter(req))
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Printf("%# v", pretty.Formatter(via))
		fmt.Println("Redirected. In function")
		return nil
	}

	resp, err := client.PostForm(Url.String(), form)
	fmt.Printf("%# v", pretty.Formatter(resp.Header))
	fmt.Printf("%# v", pretty.Formatter(client.Jar))
	//resp, err := http.DefaultTransport.RoundTrip(req)
	*/
/*fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	fmt.Println(resp.Cookies())*//*

	if err != nil {
		return rs, false, err
	}
	defer func() {
		if (resp != nil) {
			resp.Body.Close()
		}
	}()
	if resp != nil {
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return rs, false, err
		}

		return resBody, isExpired, nil
	} else {
		return rs, false, errors.New("No response")
	}
}

func GetRequest(rawUrl string) (rs []byte, isExpired bool, err error) {
	Url, err := url.Parse(rawUrl)
	if err != nil {
		return rs, false, err
	}
	Url.RawQuery = Url.Query().Encode()
	req, err := http.NewRequest(http.MethodGet, Url.String(), nil)
	if err != nil {
		return rs, false, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Jar : globalJarCookies,
	}
	fmt.Println(Url.String())
	resp, err := client.Do(req)
	if err != nil {
		return rs, false, err
	}
	defer func() {
		if (resp != nil) {
			resp.Body.Close()
		}
	}()
	if resp != nil {
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return rs, false, err
		}

		return resBody, isExpired, nil
	} else {
		return rs, false, errors.New("No response")
	}
}
*/

/*package main

import (
	"fmt"
	"strings"
	"net/url"
)

var testHotels = []string{"1300","3003","3025","3050","3074","3082","3096","3098","3102","3130","3146","3155","3166","3192","3211","3242","3347","3348","3363","3370","3374","3412","3466","3545","3564","3582","3591","3601","3606","3623","3628","3629","3630","3640","3736","3753","3758","3761","3849","3896","3956","4002","4033","4141","4142","4152","4169","3051","3097","3161","3215","3361","3367","3443","3450","3670","3759","3772","3843","4003","4114","4138","4143","4561","1811","5619","101","1109","1274","1365","1488","149","1539","1553","1757","186","199","277","3024","3081","3382","3522","3524","3528","3565","3635","3714","3802","3877","3929","3965","4110","412","4311","4313","4320","4343","4434","1802","1807","1810","1817","1818","1825","1826","1827","1828","1833","1834","1837","1840","1841","1842","1846","1848","1860","1878","1887","1899","1901","1925","1935","1936","1937","1945","1946","1950","3031","3037","3160","3238","3265","3310","3431","3446","3447","3529","3546","3554","3744","3870","3876","4001","4212","102","1046","105","1056","108","111","1113","1124","1126","1127","1138","1140","1142","1151","1156","1158","1160","1162","1168","118","1195","1203","121","1230","1236","1238","1247","1257","1260","1263","127","129","1296","1298","1326","137","1372","140","1400","142","143","1439","145","1454","146","1460","1468","148","1481","1482","1483","1489","1506","1515","1523","1534","1537","1544","1545","1546","1552","156","1561","1564","1572","1576","1577","1578","1582","1590","164","1709","1713","1716","172","1727","1740","1748","175","1751","1774","1783","180","1832","196","1962","1964","1968","1969","1973","1995","20","205","213","214","224","232","255","268","270","271","292","295","3005","3009","3012","3013","3016","3017","304","3068","307","3078","3086","3089","3093","3094","3107","3110","3116","3117","3119","3120","3134","3138","315","3157","3185","3191","3213","3217","323","3245","3268","327","33","3309","3330","3342","3372","3376","3378","3384","3387","3388","3399","3409","3414","3415","343","3440","3442","3449","3475","3490","3494","3501","3502","3504","3507","3508","3510","3515","3523","3531","3532","3537","3538","3552","3553","3555","3556","3569","3570","3572","3578","3587","3595","3596","3597","3603","3607","3608","361","3611","3613","3615","3617","3620","3622","3626","3627","3638","3639","364","3649","3654","3658","3674","3684","3692","3696","370","3716","3717","3724","3731","3732","3733","3743","3750","3766","3775","3779","3780","3791","3792","3800","3812","3830","3836","384","3856","386","3878","3879","3880","3889","3898","3901","3907","3910","3922","3928","3941","3947","3954","3961","3971","3986","3989","40","4007","4011","4015","4024","4063","4071","408","4095","4097","4113","4116","4117","4139","4146","4158","4176","4177","4186","4191","4193","4234","4241","4242","4256","4267","4278","4341","4349","435","438","445","463","468","471","48","481","482","483","485","497","5604","5607","5611","606","609","636","691","692","693","713","786","822","83","832","865","90","91","931","941","949","951","954","956","989","994","4368","4420","1471","1521","1585","1736","1787","1788","1965","3019","3028","3057","3058","3167","3170","3183","3221","3226","3258","3270","3462","3464","3819","97508","97509","97518","1001","1002","1004","1005","1009","1018","1022","1023","1030","1031","1032","1036","1044","1057","1062","1078","1080","1084","1087","1090","1092","1095","1107","1148","1167","1169","1173","1183","1185","1189","1191","1193","1198","1369","1379","1394","141","1415","1443","1445","1446","1526","1528","1562","1704","1706","1710","1722","1750","1764","1771","202","283","3000","3001","3049","3118","3156","3159","3216","3219","3232","3314","3346","3358","3398","3401","3452","35","3500","3541","3542","3560","3571","3577","3682","3685","3786","3804","3818","3827","3899","3902","4045","4140","4243","4268","5639","5653","70","79","110","1533","1708","1734","193","3080","3121","3198","3199","3200","3482","3817","4130","4195","4315","71"}

type RS struct {
	Message string
	Success string
}

func main() {
	fmt.Println(UrlEncoded("http://10.200.151.6:8081/key/add?client=facade&view_available_policies=WARRIOR+58f70a54e1382352b7f3dbea&view_group_list=WARRIOR"))
	*//*channel := make(chan interface{}, 3)
	go func() {
		for i := 0; i < 1000; i++ {
			channel <- i
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			channel <- i
		}
	}()
	for value := range channel {
		fmt.Println(value)
		//if value == 999 {
		//	break
		//}
	}*//*


}

func UrlEncoded(str string) (string, error) {
	_, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	test, _ := url.ParseQuery(str)
	test.Get("view_available_policies")
	return test.Get("view_available_policies"), nil
}

func getId(uuid string) string{
	return uuid[strings.Index(uuid, "+") + 1 :]
}*/
/*package main

import (
	"fmt"
	"unicode/utf8"
	"unicode"
	"strings"
	"github.com/kr/pretty"
)

type Text []byte

func main() {
	a := SplitToText("New York")
	fmt.Printf("%# v", pretty.Formatter(a))
}

func SplitToText(text string) []string {
	trimTxt := strings.TrimSpace(text)
	textBytes := []byte(trimTxt)
	splitter := new(splitter)
	texts := splitter.SplitTextToWords(textBytes)
	tokens := make([]string, 0)
	for _, row := range texts {
		tokens = append(tokens, string(row))
	}
	return tokens
}

type splitter struct {
	current              int
	inAlphanumeric       bool
	alphanumericStart    int
	minAlphanumericCount int
}

func (s *splitter) SplitTextToWords(text Text) []Text {
	output := make([]Text, 0, len(text) / 3)
	for s.current < len(text) {
		r, size := utf8.DecodeRune(text[s.current:])
		if s.isAlpha(r, size) {
			output = s.processAlpha(r, size, text, output)
		} else {
			output = s.processNotAlpha(r, size, text, output)
		}
		s.current += size
	}

	if s.inAlphanumeric &&  s.current != 0 {
		output = append(output, toLower(text[s.alphanumericStart:s.current]))
	}
	return output
}

//Judge the Rune is English word's or not, the min count of the letters is MIN_TOKEN_COUNT.
func (s *splitter) isAlpha(r rune, size int) bool {
	return size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r))
}

//Process function when the reading Rune is English letter or number.
func (s *splitter) processAlpha(r rune, size int, text Text, output []Text) []Text {
	if s.minAlphanumericCount < 3 {
		if !s.inAlphanumeric {
			s.alphanumericStart = s.current
			s.inAlphanumeric = true
		}
		s.minAlphanumericCount += 1
	} else if s.current != 0 {
		output = append(output, toLower(text[s.alphanumericStart:s.current]))
	}
	return output
}

//Process function when the reading Rune is not an letter, which is something like Chinese character, blank or anything else.
func (s *splitter) processNotAlpha(r rune, size int, text Text, output []Text) []Text {
	if s.inAlphanumeric {
		s.inAlphanumeric = false
		s.minAlphanumericCount = 0
		if s.current != 0 {
			output = append(output, toLower(text[s.alphanumericStart:s.current]))
		}
	}
	output = append(output, text[s.current:s.current + size])
	return output
}

func toLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}

func TestChannel() <-chan int {
	channel := make(chan int)
	return channel
}*/

/*
import (
	"fmt"
	"reflect"
)





func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	v.Type()
	fmt.Println("value:", reflect.ValueOf(x))
}*/

/*
import(
	"fmt"
	"os"
	"path/filepath"
)



func main() {

	workPath, _ := os.Getwd()
	workPath, _ = filepath.Abs(workPath)
	appPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	*/
/*rq := hotelSearch.HotelDescriptiveInfoRequest{}
	rq.UserName = "abc"
	rq.Password = "cde"
	rq.Token ="cdee"
	info := hotelSearch.HotelDescriptiveInfo{
		HotelCode: "test",
	}
	arr := []*hotelSearch.HotelDescriptiveInfo{
		&info,
	}
	infos := hotelSearch.HotelDescriptiveInfos{
		HotelDescriptiveInfo: arr,
		Language: "en-US",
	}
	rq.HotelDescriptiveInfos = &infos
	b, _ := json.Marshal(rq)*//*

	fmt.Println(os.Args[0])
	fmt.Println(appPath)

}
*/

/*
import (
	"github.com/deckarep/golang-set"
	"fmt"
)

func main() {
	sets := mapset.NewSet()
	sets.Add("a")
	fmt.Println(sets.Contains("a"))
}
*/

/*import (
	"fmt"
)
type r interface {}

type u func (a interface{}) interface{}

type Base struct {
	A string
}

func (b *Base)inter() {}

type Child struct {
	Base
	B string
}

type BaseInter interface{
	inter()
}

func main() {
	*//*test := new(Child)
	if _, ok := toInterface(test).(BaseInter); ok {
		fmt.Println(ok)
	}
	fmt.Println("123")*//*
	var a []string
	fmt.Println(a)
}

func toInterface() r {
	return toS("ab")
}

func test(a u) {
	fmt.Println(a("b"))
}

func test2(a r) {
	fmt.Println("c")
}

func toS(a string) string {
	return a
}*/

/*type PerformanceLogFactory struct {
	logs perf.PerformanceLogs
}

func (p *PerformanceLogFactory) AppendMulti(values []string, name string) {
	for index, row := range values {
		logName := p.createMultiLogField(name, index)
		p.logs.Append(logName, row)
	}
}

func (p *PerformanceLogFactory) createMultiLogField(name string, index int) string {
	return fmt.Sprint(name, LOG_SEPARATOR, index)
}

func NewPerf(rq PerformanceLog) perf.PerformanceLogs{
	performanceLogFactory := new(PerformanceLogFactory)
	performanceLogFactory.logs = new(perf.PerformanceLogs)
	fields := util.GetAttributesMap(rq)
	for name, value := range fields {
		if strValue, ok := value.(string); ok {
			performanceLogFactory.logs.Append(name, strValue)
		} else if strsValue, ok := value.([]string); ok {
			performanceLogFactory.AppendMulti(strsValue, name)
		}
	}
	return *performanceLogFactory.logs
}

type PerformanceLog struct {
	Duration         string `json:"process.duration"`
	Process          string `json:"process"`
	Result           string `json:"process.result"`
	Supplier         string `json:"process"`
	Channel          string  `json:"channel"`
	HotelCode        []string `json:"hotel.supplier"`
	ChannelHotelCode []string `json:"hotel.channel"`
	Language         string `json:"language"`
}*/

/*
import (
	"derbysoft.com/common-go/util"
	"fmt"
)

type Test struct {
	Atest string `json:"atest"`
}

func main() {
	test := new(Test)
	test.Atest = "abd"
	maps := util.GetAttributesMap(test)
	for index, value :=range maps {
		fmt.Println(index, value)
	}

}
*/

/*import "fmt"

var chans chan string
func main() {
	chans = make(chan string)
	go showValue()
	chans <- "a"
	go showValue()
	chans <- "b"
	close(chans)
	if _, ok := <- chans; ok {
		fmt.Println("ab")
	}

}

func showValue() {
	if str, ok := <- chans; ok {
		fmt.Println(str)
	}
}*/

/*import (
	"unicode"
	"unicode/utf8"
	"fmt"
)

import (
	"derbysoft.com/common-go/util"
	"fmt"
)

type Result interface {
	SetValue()
}

type arr string

func (a arr) SetArr() {
	a = "a"
}



func main() {
	fmt.Println(util.GenerateUuid())
}*/

/*func test(a ...interface{}) {
	for _, x := range a {
		if v, ok := x.(arr);ok {
			c := append([]string{}, "a")
			copy(v, c)
			fmt.Println(len(v))
			fmt.Println(x)
		}
	}
}*/

/*func test4(t []string) {
	t[0] = "x"
}*/

/*func test2() interface{} {
	arrays := make([]int, 0)
	for i := 0; i < 10; i++ {
		arrays = append(arrays, i)
	}
	return arrays
}

func test3() interface{} {
	d := ""
	e := make([]interface{}, 0)
	e = append(e, d)
	return e
}

type b struct {
	Test string `json:"test"`
}*/
/*
import (
	"fmt"
	"unicode/utf8"
	"unicode"
)

type Text []byte

func main() {
	texts := SpiltToTokens("ASD - Andros Town Airport, Bahamas")
	texts2 := SpiltToTokensOld("ASD - Andros Town Airport, Bahamas")
	for index, row := range texts {
		fmt.Println(string(row))
		fmt.Println(string(texts2[index]))
	}
}

func SpiltToTokens(text string) []Text {
	textBytes := []byte(text)
	splitter := new(splitter)
	return  splitter.SplitTextToWords(textBytes)
}

func SpiltToTokensOld(text string) []Text {
	textBytes := []byte(text)
	return  splitTextToWords(textBytes)
}

func splitTextToWords(text Text) []Text {
	output := make([]Text, 0, len(text) / 3)
	current := 0
	inAlphanumeric := true
	alphanumericStart := 0
	minAlphanumericCount := 0
	for current < len(text) {
		r, size := utf8.DecodeRune(text[current:])
		if size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r)) && minAlphanumericCount < 3 {
			if !inAlphanumeric {
				alphanumericStart = current
				inAlphanumeric = true
			}
			minAlphanumericCount += 1
		} else if size > 2 || !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			if inAlphanumeric {
				inAlphanumeric = false
				minAlphanumericCount = 0
				if current != 0 {
					//adding current word
					output = append(output, toLower(text[alphanumericStart:current]))
				}
			}
			//adding the separator character
			output = append(output, text[current:current + size])
		} else {
			if current != 0 {
				//Adding the current word's letters one after one in sequence.
				output = append(output, toLower(text[alphanumericStart:current]))
			}
		}
		current += size


	}

	if inAlphanumeric {
		if current != 0 {
			output = append(output, toLower(text[alphanumericStart:current]))
		}
	}

	return output
}

type splitter struct {
	current              int
	inAlphanumeric       bool
	alphanumericStart    int
	minAlphanumericCount int
}

func (s *splitter) SplitTextToWords(text Text) []Text {
	output := make([]Text, 0, len(text) / 3)
	for s.current < len(text) {
		r, size := utf8.DecodeRune(text[s.current:])
		if s.isAlpha(r, size) {
			output = s.processAlpha(r, size, text, output)
		} else {
			output = s.processNotAlpha(r, size, text, output)
		}
		s.current += size
	}

	if s.inAlphanumeric && s.current != 0 {
		output = append(output, toLower(text[s.alphanumericStart:s.current]))
	}
	return output
}

//Judge the Rune is English word's or not, the min count of the letters is MIN_TOKEN_COUNT.
func (s *splitter) isAlpha(r rune, size int) bool {
	return size <= 2 && (unicode.IsLetter(r) || unicode.IsNumber(r))
}

//Process function when the reading Rune is English letter or number.
func (s *splitter) processAlpha(r rune, size int, text Text, output []Text) []Text {
	if s.minAlphanumericCount < 3 {
		if !s.inAlphanumeric {
			s.alphanumericStart = s.current
			s.inAlphanumeric = true
		}
		s.minAlphanumericCount += 1
	} else if s.current != 0 {
		output = append(output, toLower(text[s.alphanumericStart:s.current]))
	}
	return output
}

//Process function when the reading Rune is not an letter, which is something like Chinese character, blank or anything else.
func (s *splitter) processNotAlpha(r rune, size int, text Text, output []Text) []Text {
	if s.inAlphanumeric {
		s.inAlphanumeric = false
		s.minAlphanumericCount = 0
		if s.current != 0 {
			output = append(output, toLower(text[s.alphanumericStart:s.current]))
		}
	}
	output = append(output, text[s.current:s.current + size])
	return output
}


func toLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}*/

/*import (
	"sort"
	"fmt"
	"math/rand"
	"strconv"
	"github.com/kr/pretty"
)


type Destination struct {
	Id           uint64 `json:"id"`

	PopRank      string `json:"pop"`

}

type Sorting struct {
	Dests []*Destination
}

func (s Sorting) Less(i, j int) bool {
	popValue64I, err := strconv.ParseFloat(s.Dests[i].PopRank, 5)
	if err != nil {
		return false
	}
	popValue64J, err := strconv.ParseFloat(s.Dests[j].PopRank, 5)
	if err != nil {
		return true
	}
	return popValue64I < popValue64J
}

func (s Sorting) Len() int {

	return len(s.Dests)
}

func (s Sorting) Swap(i, j int) {
	temp := s.Dests[i]
	s.Dests[i] = s.Dests[j]
	s.Dests[j] = temp
}

func main() {
	wbs := make([]*Destination, 0)
	for i := 0; i < 10; i++ {
		wbs = append(wbs, &Destination{
			Id: uint64(i),
			PopRank: strconv.Itoa(rand.Intn(100)),
		})
	}
	s := Sorting{
		Dests: wbs,
	}
	fmt.Printf("%# v", pretty.Formatter(s.Dests))
	sort.Sort(s)
	for _,row := range s.Dests {
		fmt.Println(row)
	}
}*/

/*
import "fmt"

type dest struct {
	Test string
}

func main() {
	maps := make(map[string]dest)
	maps["x"] = dest{
		Test: "a",
	}

	fmt.Println(maps["a"])


}
*/

/*
import (
	"strconv"
	"fmt"
	"strings"
)

func main() {
	f := StringsToInt("zzz")
	fmt.Println(f)
}

func StringsToInt(s string) uint64 {
	strs := ""
	for _, row := range strings.Split(s, "") {
		char := []rune(row)
		for _, c := range char {
			i := int(c)
			str := strconv.Itoa(i)
			strs += str
		}
	}
	if len(strs) > 19 {
		strs = strs[13:19]
	}
	intStrs, err := strconv.Atoi(strs)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return uint64(intStrs)
}
*/

/*import (
	"fmt"
	"net"
	"os"
	"io"
)


type dest struct {
	Dests []*dests `json:"dests"`
}

type dests struct {
	Name string `json:"name"`
	Sub []*subdest `json:"sub"`
	Type int `json:"type"`
}

type subdest struct {
	Name string `json:"name"`
}



func main() {
	host := "0.0.0.0"
	port := "777"
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}

	for {
		s_conn, err := l.Accept()
		if err != nil {
			continue
		}

		d_tcpAddr, _ := net.ResolveTCPAddr("tcp4", "192.168.137.105:80")
		d_conn, err := net.DialTCP("tcp", nil, d_tcpAddr)
		if err != nil {
			fmt.Println(err)
			s_conn.Write([]byte("can't connect 192.168.137.105:80"))
			s_conn.Close()
			continue
		}
		go io.Copy(s_conn, d_conn)
		go io.Copy(d_conn, s_conn)
	}
}*/
/*
	//dictFile, _ := os.Open("D:\\Workspace\\hotelContentService\\src\\derbysoft.com\\hotel\\content\\wukong\\data\\dest_json.txt")
	//reader := bufio.NewReader(dictFile)
	bytesT, _ := ioutil.ReadFile("D:\\Workspace\\hotelContentService\\src\\derbysoft.com\\hotel\\content\\wukong\\data\\location_chs.js")
	//bytesT, _, _ := reader.ReadLine()
	//dests := new(dest)
	str := string(bytesT)
	fmt.Println(str)
	*//*json.Unmarshal(bytesT,dests)
	for _, row := range dests.Dests  {
	  for _, subrow := range row.Sub {
		  fmt.Println(row.Name,subrow.Name)
	  }
	}*//*
}*/

/*client := http.Client{
Timeout: time.Duration(1000) * time.Millisecond,
}
request, _ := http.NewRequest("GET", "http://localhost:8080/json?query=r", nil)
//request.Header.Add("Accept-Encoding", "gzip")

response, _ := client.Do(request)
rs,_ :=ioutil.ReadAll(response.Body)
//json.Unmarshal(rs)
fmt.Printf("%# v", pretty.Formatter(string(rs)))*/

/*

// 载入词典
var segmenter sego.Segmenter
segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")

// 分词
text := []byte("上海")
segments := segmenter.Segment(text)

// 处理分词结果
// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
fmt.Println(sego.SegmentsToString(segments, false))*/

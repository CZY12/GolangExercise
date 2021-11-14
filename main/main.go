package main

import (
	"CzyCoding/service"
	"fmt"
	"net/http"
	"text/template"

	urlshortener "google.golang.org/api/urlshortener/v1"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	a := rand.Int()
	//	fmt.Printf("%d / ",a)
	//}
	//fmt.Println()
	//for i := 0; i < 5; i++ {
	//	r := rand.Intn(8)
	//	fmt.Printf("%d / ",r)
	//}
	//fmt.Println()
	//times := int64(time.Now().Nanosecond())
	//fmt.Printf("%v",times)
	//fmt.Println()
	//rand.Seed(times)
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%2.2f /",100*rand.Float32())
	//
	//}

	voo := new(service.Voodoo)
	voo.Magic()
	voo.Base.Magic()
	voo.MoreMagic()
	merced := new(service.Mercedes)

	merced.SetWheel(4)
	merced.SayHiToMerkel()
	merced.NumberOfWheels()
	merced.Start()
	merced.Stop()

	e := new(service.Employee)
	e.SetSalary(10)
	raise := e.GiveRaise(0.1)
	fmt.Println(raise)

	http.HandleFunc("/", root)
	http.HandleFunc("short", short)
	http.HandleFunc("/long", long)
	http.ListenAndServe("localhost:8080", nil)

	//service.BigAdd()
	//
	//stt1 := "helloworld"
	//fmt.Println(stt1[len(stt1)/2:] + stt1[:len(stt1)/2])
	//
	//fmt.Println(service.ReverseString(stt1))
	//var isss = []string{"1", "2", "2", "3", "7"}
	//fmt.Println(service.SearchNoEq(isss))
	//var i = []string{"1", "2", "3", "7"}
	//i = i[1:2]
	//
	//so := []int{1, 2, 3, 4, 5, 6, 7}
	//
	//so = service.BubbleSort(so)
	//
	//so = service.MapFunc(service.MapFunc1, so)
	//
	//fmt.Println(so)
	//
	//fmt.Println(len(i))
	//fmt.Println(i[0])
	//var str = []string{"1", "2", "3", "7"}
	//var newSlice = []string{"4", "5", "6"}
	//
	//cutStr, end := service.CutStr("dsdsd", 2)
	//
	//fmt.Println(cutStr, "----", end)
	//
	//var ns = service.InsertStringSlice(newSlice, str, 3)
	//
	//for i, num := range ns {
	//	fmt.Println("%d is %n", i, num)
	//}
	//var i1 = 5
	//var intP *int
	//intP = &i1
	//fmt.Printf("An integer: %d, it's lpcation in memory: %p\n", i1, &i1)
	//
	//fmt.Printf("int : %p", intP)

	//var a = 3
	//var b = 1
	//var A, B, C int
	//var A1, B1, C1 int
	////A, B, C = Return1(a, b)
	//fmt.Printf("非命名返回:{0},{1},{2}", A, B, C)
	//A1, B1, C1 = Return2(a, b)
	//fmt.Println(fmt.Printf("非命名返回:{0},{1},{2}", A1, B1, C1))

}

var rootHtmpInpl = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))

func root(w http.ResponseWriter, r *http.Request) {
	rootHtmpInpl.Execute(w, nil)
}

func short(w http.ResponseWriter, r *http.Request) {
	longUtl := r.FormValue("longUrl")
	s, _ := urlshortener.New(http.DefaultClient)
	do, _ := s.Url.Insert(&urlshortener.Url{LongUrl: longUtl}).Do()

	rootHtmpInpl.Execute(w, fmt.Sprintf("Shortened version of %s is : %s",
		longUtl, do.Id))

}

func long(w http.ResponseWriter, r *http.Request) {
	shortUrl := "http://goo.gl/" + r.FormValue("shortUrl")
	s, _ := urlshortener.New(http.DefaultClient)
	do, err := s.Url.Get(shortUrl).Do()
	if err != nil {
		fmt.Println("error: %v", err)
		return
	}
	rootHtmpInpl.Execute(w, fmt.Sprintf("Longer version of %s is : %s",
		shortUrl, do.LongUrl))
}

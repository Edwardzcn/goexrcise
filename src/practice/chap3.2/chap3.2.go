package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html>
<head><style>.error{color:#FF0000;} .fail{color:#F00;} .pass{color:#0F0;}</style></head>
<title>Soundex</title>
<body>
<h3>Soundex</h3>
<p>Compute soundex codes for a list of names.</p>`
	form = `<form action="/" method="POST">
<label for="names">Names (comma or space-separated):</label><br />
<input type="text" name="names" size="30"><br />
<input type="submit" name="compute" value="Compute">
</form>`
	pageBottom   = `</body></html>`
	error        = `<p class="error">%s</p>`
	message      = `<p class="message">%s</p>`
	tableTop     = `<table border="1"><tr><th>Name</th><th>Soundex</th></tr>`
	tableTestTop = `<table border="1"><tr><th>Name</th><th>Soundex</th>
	<th>Expected</th><th>Test</th></tr>`
	tableBottom = `</table>`
)

var testTable map[string]string
var ok bool

func main() {
	http.HandleFunc("/", homePage)

	if testTable, ok = ReadTestFile("soundex-test-data.txt"); ok {
		// fmt.Println(len(testTable))
		http.HandleFunc("/test", testPage)
	} else {
		http.HandleFunc("/test", errorPage)
	}
	if err := http.ListenAndServe(":9012", nil); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

func ReadTestFile(path string) (map[string]string, bool) {
	if data, err := ioutil.ReadFile(path); err != nil {
		return nil, false
	} else {
		lines := strings.Split(string(data), "\n")
		testTable = make(map[string]string, len(lines))
		for _, line := range lines {

			if spaceIndex := strings.Index(line, " "); spaceIndex == -1 {
				continue
			} else {
				mapValue := strings.TrimSpace(line[:spaceIndex])
				mapKey := strings.TrimSpace(line[spaceIndex:])
				testTable[mapKey] = mapValue
			}
		}
		return testTable, true
	}
}

func errorPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, pageTop)
	fmt.Fprintf(writer, error, "Something is error")
	fmt.Fprintf(writer, pageBottom)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, error, err)
	} else {
		if names := processRequest(request); len(names) > 0 {
			soundexes := make([]string, len(names))
			for i, name := range names {
				soundexes[i] = soundex(name)
			}
			fmt.Fprint(writer, formatResults(names, soundexes))
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) (names []string) {
	if slice, found := request.Form["names"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		names = strings.Fields(text)
	}
	return names
}

func formatResults(names, soundexes []string) string {
	text := `<table border="1"><tr><th>Name</th><th>Soundex</th></tr>`
	for i := range names {
		text += "<tr><td>" + html.EscapeString(names[i]) + "</td><td>" +
			html.EscapeString(soundexes[i]) + "</td></tr>"
	}
	return text + "</table>"
}

func testPage(writer http.ResponseWriter, request *http.Request) {
	if len(testTable) == 0 {
		fmt.Fprintf(writer, message, "Empty test list")
		return
	}
	fmt.Fprintf(writer, pageTop)
	fmt.Fprintf(writer, tableTestTop)

	keys := make([]string, 0, len(testTable))
	for key := range testTable {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		expecteValue := testTable[key]
		caculatValue := soundex(key)
		var testValue string
		if expecteValue == caculatValue {
			testValue = `<span class = "pass">PASS</span>`
		} else {
			testValue = `<span class = "fail">FAIL</span>`
		}
		fmt.Fprintf(writer, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", key, caculatValue, expecteValue, testValue)
	}
	fmt.Fprintf(writer, tableBottom)
}

var digitForLetter = []rune{
	0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5,
	5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2}

func soundex(name string) string {
	var nameRune = []rune(strings.ToUpper(name))
	var ansRune []rune
	ansRune = append(ansRune, nameRune[0])
	for index := 1; index < len(nameRune); index++ {
		// 输入不合法字符
		// 需要在这里判断相等，而非判断映射后相等
		if nameRune[index]-'A' < 0 || nameRune[index-1] == nameRune[index] {
			continue
		}
		if digit := digitForLetter[nameRune[index]-'A']; digit == 0 {
			continue
		} else {
			ansRune = append(ansRune, '0'+digit)
		}
	}
	for len(ansRune) < 4 {
		ansRune = append(ansRune, '0')
	}
	// 如果跳过while本身过长，还需裁减
	// fmt.Println(string(nameRune), string(ansRune[:4]))
	return string(ansRune[:4])
}

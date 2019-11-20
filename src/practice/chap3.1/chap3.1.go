// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Song is a struct
type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func main() {

	if length := len(os.Args); length == 1 {
		fmt.Println("Error:No file!")
		fmt.Printf("usage: %s <file.m3u>/<file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if isM3u, isPks := strings.HasSuffix(os.Args[1], ".m3u"), strings.HasSuffix(os.Args[1], ".pls"); !isM3u && !isPks {
		fmt.Println("Error:No .m3u file or .pls file!")
		fmt.Printf("usage: %s <file.m3u>/<file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else if isM3u {
		if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
			log.Fatal(err)
		} else {
			// 将rawBytes作为参数传递给自定义函数readM3uPlaylist，返回Song切片
			songs := readM3uPlaylist(string(rawBytes))
			// 将切片作为参数传递给自定义函数writePlsPlaylist，这个函数没有返回类型
			writePlsPlaylist(songs)
		}
	} else {
		if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
			log.Fatal(err)
		} else {
			songs := readPlsPlaylist(string(rawBytes))
			writeM3uPlaylist(songs)
		}
	}

}

func readM3uPlaylist(data string) (songs []Song) {
	// Go自动初始化为0
	var song Song
	// 按换行符分割
	for _, line := range strings.Split(data, "\n") {
		// 去除开头结尾的空格
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		// 空行或者以#EXTM3U开头（第一行）
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			// 调试
			// fmt.Println(song)
			// 清空
			song = Song{}
		}
	}
	return songs
}

func readPlsPlaylist(data string) (songs []Song) {
	// Go自动初始化为0
	var song Song
	// 按换行符分割
	for _, line := range strings.Split(data, "\n") {
		// 去除开头结尾的空格
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		// 空行或者以#EXTM3U开头（第一行）
		if line == "" || strings.HasPrefix(line, "[playlist]") {
			continue
		}
		if strings.HasPrefix(line, "NumberOf") || strings.HasPrefix(line, "Version") {
			continue
		}
		equalIndex := strings.Index(line, "=")
		if strings.HasPrefix(line, "File") {
			song.Filename = strings.Map(mapPlatformDirSeparator, line[equalIndex+1:])
		} else if strings.HasPrefix(line, "Title") {
			song.Title = line[equalIndex+1:]
		} else {
			if seconds, err := strconv.Atoi(line[equalIndex+1:]); err != nil {
				seconds = -1
				log.Printf("failed to read the duration for '%s': %v\n",
					song.Title, err)
				song.Seconds = seconds
			} else {
				song.Seconds = seconds
			}
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			// 调试
			// fmt.Println(song)
			// 清空
			song = Song{}
		}
	}
	return songs
}

// func parseExtinfLine(line string) (title string, seconds int) {
// 	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
// 		const separator = ","
// 		line = line[i:]
// 		if j := strings.Index(line, separator); j > -1 {
// 			title = line[j+len(separator):]
// 			var err error
// 			if seconds, err = strconv.Atoi(line[:j]); err != nil {
// 				log.Printf("failed to read the duration for '%s': %v\n",
// 					title, err)
// 				seconds = -1
// 			}
// 		}
// 	}
// 	return title, seconds
// }

func parseExtinfLine(line string) (title string, seconds int) {
	i := strings.IndexAny(line, "-0123456789")
	j := strings.LastIndex(line, ",")
	if i != -1 && j != -1 {
		var err error
		if seconds, err = strconv.Atoi(line[i:j]); err != nil {
			seconds = -1
			log.Printf("failed to read the duration for '%s': %v\n",
				title, err)
		}
		title = line[j+1:]
	}
	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func writeM3uPlaylist(songs []Song) {
	fmt.Println("#EXTM3U")
	for i, song := range songs {
		i++
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Printf("Music/%s\n", song.Filename)
	}
}

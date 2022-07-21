package replying

import (
	"bufio"
	"fmt"
	"github.com/Mrs4s/MiraiGo/message"
	"juryo/models"
	"net/http"
	"os"
	"strconv"
)

func download(url string, path int64, fileName string) {
	imgPath := "images/" + strconv.Itoa(int(path))
	_, err := os.Stat(imgPath)
	if err != nil {
		if os.IsExist(err) {
			panic(err)
		}

		if os.IsNotExist(err) {
			err = os.Mkdir(imgPath, 755)
			if err != nil {
				panic(err)
			}
		}
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(imgPath + "/" + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	bytes := make([]byte, 32*1024)
	for {
		len, err := reader.Read(bytes)

		if len < 0 || err != nil {
			return
		}
		// 注意这里byte数组后的[0:len]，不然可能会导致写入多余的数据
		_, _ = writer.Write(bytes[0:len])
		fmt.Printf("%d ", len)
	}

}

func upImage(msg *message.GroupMessage) (int, int) {
	n := 0
	k := 0
	for _, elem := range msg.Elements {
		switch e := elem.(type) {
		case *message.GroupImageElement:
			{
				k++
				models.Upload(e, msg.GroupCode)
				download(e.Url, msg.GroupCode, e.ImageId)
				n++

			}
		default:
			continue
		}

	}

	return n, k
}

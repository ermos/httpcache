package wrapper

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/ermos/httpcache/internal/pkg/cache"
	"github.com/gofiber/fiber/v2"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Handle(c *fiber.Ctx, url string, minutes int) (err error) {
	var result interface{}

	key := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(minutes) + "/" + url))
	expAt := time.Now().Add(time.Minute * time.Duration(minutes)).Unix()

	defer func() {
		if r := recover(); r != nil {
			var body []byte

			body, err = cache.Get(key)
			if err == nil {
				err = json.Unmarshal(body, &result)
				if err == nil {
					err = c.Status(200).JSON(result)
					return
				}
			}

			err = r.(error)
		}
	}()

	body, err := cache.Get(key)
	if err == nil {
		err = json.Unmarshal(body, &result)
		if err != nil {
			panic(err)
		}
	} else {
		body, err = request(url)
		if err != nil {
			panic(err)
		}

		result, err = save(key, body, expAt)
		if err != nil {
			panic(err)
		}
	}

	return c.Status(200).JSON(result)
}

func request(url string) (result []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("an error occured")
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(res.Body)

	return ioutil.ReadAll(res.Body)
}

func save(key string, body []byte, expAt int64) (result interface{}, err error) {
	cache.Set(key, body, expAt)
	err = json.Unmarshal(body, &result)
	return
}

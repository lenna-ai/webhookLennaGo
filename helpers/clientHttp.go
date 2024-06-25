package helpers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// func HitApiJsonWithoutAuth(c *fiber.Ctx, url string, data map[string]interface{}) map[string]any {
// 	var err error
// 	agent := fiber.Post(url)

// 	agent.JSON(data)
// 	statusCode, body, errs := agent.Bytes()
// 	if len(errs) > 0 {
// 		return map[string]any{
// 			"body":       nil,
// 			"statusCode": 0,
// 			"error":      err,
// 		}
// 	}

// 	// pass status code and body received by the proxy

//		var dataApiResult map[string]any
//		err = json.Unmarshal(body, &dataApiResult)
//		if err != nil {
//			panic(err.Error())
//		}
//		return map[string]any{
//			"body":       dataApiResult,
//			"statusCode": statusCode,
//			"error":      nil,
//		}
//	}
func HitApiJsonWithoutAuth(c *fiber.Ctx, url string, data map[string]interface{}, ch chan map[string]any) {
	var err error
	agent := fiber.Post(url)

	agent.JSON(data)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		ch <- map[string]any{
			"body":       nil,
			"statusCode": 0,
			"error":      err,
		}
		return
	}

	// pass status code and body received by the proxy

	var dataApiResult map[string]any
	err = json.Unmarshal(body, &dataApiResult)
	if err != nil {
		panic(err.Error())
	}
	ch <- map[string]any{
		"body":       dataApiResult,
		"statusCode": statusCode,
		"error":      nil,
	}
	return
}

func HitApiBasicAuthJson(c *fiber.Ctx, url string, data map[string]string, ch chan map[string]any) {
	var err error
	agent := fiber.Post(url)

	agent.BasicAuth("lenna", "w4rt36KHARISMA")
	agent.JSON(data)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		ch <- map[string]any{
			"body":       nil,
			"statusCode": 0,
			"error":      err,
		}
		return
	}
	// pass status code and body received by the proxy
	var dataNlpResult map[string]any
	err = json.Unmarshal(body, &dataNlpResult)
	if err != nil {
		panic(err.Error())
	}
	ch <- map[string]any{
		"body":       dataNlpResult,
		"statusCode": statusCode,
		"error":      nil,
	}
	return
}

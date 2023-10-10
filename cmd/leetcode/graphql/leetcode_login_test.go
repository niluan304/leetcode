package graphql

import (
	"context"
	"testing"

	"github.com/machinebox/graphql"
)

func Test_LeetcodeLogin(t *testing.T) {
	client := graphql.NewClient(EndpointZh)

	req := graphql.NewRequest(`mutation signInWithPassword($data: AuthSignInWithPasswordInput!, $nvcData: String!) {
  authSignInWithPasswordByNvc(data: $data, nvcData: $nvcData) {
    output {
      ok
      otpConfig {
        configUrl
        userToken
        __typename
      }
      __typename
    }
    nvcStatus
    __typename
  }
}
`)
	req.Var("data", map[string]string{
		"username": "zwei_elen_2023-10-03",
		"password": "4Wc9FcuvtSshwfWP",
	})
	req.Var("nvcData", "%7B%22a%22%3A%22FFFF000000000179BB25%22%2C%22c%22%3A%221696377017705%3A0.667815277987424%22%2C%22d%22%3A%22nvc_login%22%2C%22j%22%3A%7B%22test%22%3A1%7D%2C%22h%22%3A%7B%22umidToken%22%3A%22T2gA02bDkZRgR2Q9GZ_NuWs947MwFDhCl1eOpbEwNt2DiDKKZVboRbP2KaVUKjyAbqY%3D%22%7D%2C%22b%22%3A%22140%23NOfdfVbozzWACzo22Z3QA6SogFrJddvWCtj0ZA8qb4G6RbhCi%2FSFs77kHkyblJmj2P5YuY9Yf5zdUl7dcuAzruoqlbzxhEXNbR8tzzcFD2kvl3MzzPzbVXl%2FlbcoMIcLI6hqabzi228zPrfxbSziKbhqkQnA2IrkH31fxQdS2R%2BnlfzEKrvZrI7ZbioJzvo0QCx6%2FrS9q6tib2BemPXdKqYK0pi2cuok1izIhrJE2KqmnQaie6GMD6zonLZFZU4%2F2MCUjs8UDRwNPpzHqzV%2BMAfONtQl7LA70s%2F%2FedaCyoMVVSRZ%2F%2BsQcjdsebW365OTw0ZLMuwmAnRA7gy244sMWXfcKcq0f9r2qN%2B%2Bza3PirKsgkgtb24rjtewICzARKracEQcB5PZ7d0CbneiHt0okULNBZkDWdl0RxtRoKsY5Avn%2Ftgqksr7pwvqXSqaTSHiqFs4OK2JfyrUPbTcRbG0kKxlg9BedwELI3XWMZL%2B1im71A3NJuLYaoZT%2F4FW8%2B3%2BT6a%2B93tHhvUo3SF00DHyVMmzFqBAhV8RSVzmhjYdED0xrBgsF88fBO5qPD78XCTaIG3IDeF66KzJtvtnZOVmQy1VU4JUJLIj9K1nNkPdaC7Ypbv4tsJpLhrf2zSAQLkZtdQ2kYaaGOuHKRXDBdwqPuqMbod%2ByOAg35LYAB6tipMg6g9%2FBs6J11kUOELibmo1z1MN9Hg1OeDYm6xef9dNWpsm85ZThNrpgxKLEgi3c0ItQsMXFp3A%2F%2B0pz0Jh0X8LJC6YKZW8HZht6Tj7GAyAkHfmjWZwkepUeaOM3iG81ZtbP9Na4nJvEoDgdxMaPCg%2Bv931ljkaF05qYJ6WLrhyyXzPnqZiWpnpjIssf%2BkW4yIcvxQkrYe4e0tUsLM8iXFw2kW7DQVhZrsRA36YiH%2BtupZxMI1yxS3M3xGvi3oGcTnhB6AqPWbvlowq4%2By5f4sr%2BpwKMpDzq9LqYTqwkTjja0qC8i1AluAtS8igWE2Vz3P7%2Fb5iFp%2BpPNJSXUJbJvVl2%2BKyllhFTJgtaQqyPonx689lYO%2BhoaT9ArNZS4g8HlOOPvmduF%2BRsERQPHYck5c4VzMzePWi6f8pidbK7DOEbQM59uJdaC8DkdbT1Zj39DzGK7Qbk46674%2FlgmazrD9U64%2Bi4IfdfgfhihhgiCQ6yG%2BEnT1CwPWoT9S1SLqYTgtYEZHncXidEcDh1BX0trMus%2FuzT35DPvqWDifOXWAIleXPS1g88V%2BF9A%2BcvfejyNnqEKeO7RVkahwMifi3r83PIK8ZQyWzSOh%2BPDAUdGFdvSzJ1TTDqtP6%2BTDe1zVW874G5ok1%2BaxL6jmei2DDaLxRyLBHo0jAP8Q3WoCwhgDIOTPNsgBxHxMHBSxJTgECh3EGFvIe%2Fy8Nn4fPjkw7NW7AVQGnn820QfBQCzc3Tszgq2RmyqgpMqdWVfqIA2MiVOZDKBifwASzyrz4lKxOOpTsLO3TEQT0dP6TCYLTNaAPEX3watD4oLpTma2jzMR1NFyKMx0FtW3dKjbjhDAa%2BXWEKS5Cas%2BJWSmSGFgONhYLdgIYKjjoicK%2BV%2BH%2FbDHpDtHcSLqQnUKHKESSSnhJSxQinjAynLqRS8JdKf%2F4nSelKa9Vl7rN0j%2F8xwX%2BxuCEMwL34w7WQTYX5O8%2BMe3RZ7vG4eazPW7Lie4U%2FGAALUprNSN0X5zEdnWF0gxdlpIHN6fZArxYqYqoOhgnmK415u39%2FdlFjhOhjs6cVZ6HUMFkkANIR%2F6MRaPp%2BmDQPazgypOuQx5Jg%2B%2BGfFL4hqjn8lIq4uu6tHAcpdqbBqRTthYoCwPa8pipcjcdJb4P3JeM6z%3D%3D%22%2C%22e%22%3A%22b-Zr5yVRdtVlQVfYaS3LUkGgCrqsTEJDGfLy7uXEN6EX3dn3j40nLVqYGCeoaoKiDQyvw1MtY9wAK-CBiWWSETXIBTvvm3mgLiePIaX7rW9OulgfG2WRN2e3JKQmSAqAErhbW72MRf0b2IRTNvflqM2Lu_2nEav3U-f14IFn0mMqWCc5jh5uf0EtlD4AR-Z7hm0rGkzKkCwhhs7BIToSo4RytR_KgDE59WTal9AlggilyQEsnRdXb5MdXNi80Lzo%22%7D")

	cookie := "csrftoken=1B03OL1VuzIG0xLX6n3rxPZ7drNzn29BCeRocYVdw9jBBAx6KZq3d2aZF9xOCsIr; Domain=.leetcode.cn; expires=Sat, 05 Oct 2024 02:22:19 GMT; Max-Age=31449600; Path=/; SameSite=Lax; Secure"
	//"Cookie": _bl_uid=Cnl19nX1btjd2hrv32eUpzUs4t0v; csrftoken=1B03OL1VuzIG0xLX6n3rxPZ7drNzn29BCeRocYVdw9jBBAx6KZq3d2aZF9xOCsIr

	header := map[string]string{
		"Cookie":             cookie,
		"X-Requested-With":   "XMLHttpRequest",
		"Connection":         "keep-alive",
		"DNT":                "1",
		"Host":               "leetcode.cn",
		"Origin":             "https://leetcode.cn",
		"Pragma":             "no-cache",
		"Referer":            "https://leetcode.cn/accounts/login/?next=/list/problem/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.58",
		"X-CSRFToken":        "1B03OL1VuzIG0xLX6n3rxPZ7drNzn29BCeRocYVdw9jBBAx6KZq3d2aZF9xOCsIr",
		"accept":             "*/*",
		"accept-language":    "zh-CN",
		"content-type":       "application/json",
		"random-uuid":        "b035c559-dc13-0017-9a8e-6e5bb9f0ea25",
		"sec-ch-ua":          `"Microsoft Edge";v="117", "Not;A=Brand";v="8", "Chromium";v="117""`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "Windows",
		"uuuserid":           "8e092425283fbe527e3990ae54b391dc",
		"x-definition-name":  "authSignInWithPasswordByNvc",
		"x-operation-name":   "signInWithPassword",
		"x-timezone":         "Asia/Shanghai",
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}

	var resp map[string]interface{}
	err := client.Run(context.Background(), req, &resp)
	if err != nil {
		t.Error(err)
	}

	req2 := graphql.NewRequest("query userHasPassword {\n  userHasPassword\n}\n")
	var res2 map[string]interface{}
	for key, value := range header {
		req2.Header.Add(key, value)
	}

	req2.Header.Set("Cookie", "_bl_uid=Cnl19nX1btjd2hrv32eUpzUs4t0v; "+
		"csrftoken=TFlE64c3ViHJKaLZNGQokLHhA2irnAw7bYiygCXB0FXs60TqBDxDivZnYPNf6xRV; "+
		"LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiNjYwMjUxMyIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiNGYyZGE4MDljYWUyNWZlMTM4NjM3ZmMyMGJkMTRiYzZjMjc2OTA5MDEzZjRmOGM1MjNkZjMwYzdjY2MxZjViYyIsImlkIjo2NjAyNTEzLCJlbWFpbCI6IiIsInVzZXJuYW1lIjoiendlaV9lbGVuXzIwMjMtMTAtMDMiLCJ1c2VyX3NsdWciOiJ6d2VpX2VsZW5fMjAyMy0xMC0wMyIsImF2YXRhciI6Imh0dHBzOi8vYXNzZXRzLmxlZXRjb2RlLmNuL2FsaXl1bi1sYy11cGxvYWQvdXNlcnMvendlaV9lbGVuXzIwMjMtMTAtMDMvYXZhdGFyXzE2OTYzNDM5MTQucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2OTY0MDE3MDAuMjA3NDk4LCJleHBpcmVkX3RpbWVfIjoxNjk4OTUxNjAwLCJ2ZXJzaW9uX2tleV8iOjF9.h3je62J4FjAzvdSCcjhAZYMedePeo8i7rProlKVzOyo")
	err = client.Run(context.Background(), req2, &res2)
	if err != nil {
		t.Error(err)
	}

	t.Log(res2)
}

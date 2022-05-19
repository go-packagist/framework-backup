# support

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/support)](https://goreportcard.com/report/github.com/go-packagist/support)
[![tests](https://github.com/go-packagist/support/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/support/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/support)](https://pkg.go.dev/github.com/go-packagist/support)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

_——The idea came from [Laravel](https://github.com/laravel)_

## Installation

```bash
go get github.com/go-packagist/filesystem
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-packagist/support/coroutine"
	"time"
)

func main() {
	c := coroutine.NewConcurrent(10)

	for i := 1; i <= 100; i++ {
		ii := i
		c.Create(func() {
			time.Sleep(time.Second)

			fmt.Println(ii, time.Now())
		})
	}

	for {
		if c.IsEmpty() {
			break
		}
	}
}
```

Result:

```bazaar
10 2022-05-19 23:24:30.591914 +0800 CST m=+1.005890334
6 2022-05-19 23:24:30.592005 +0800 CST m=+1.005982793
7 2022-05-19 23:24:30.59171 +0800 CST m=+1.005686918
1 2022-05-19 23:24:30.590421 +0800 CST m=+1.004397918
5 2022-05-19 23:24:30.591407 +0800 CST m=+1.005383334
2 2022-05-19 23:24:30.58862 +0800 CST m=+1.002596626
4 2022-05-19 23:24:30.587968 +0800 CST m=+1.001945334
9 2022-05-19 23:24:30.59216 +0800 CST m=+1.006136376
3 2022-05-19 23:24:30.592243 +0800 CST m=+1.006220001
8 2022-05-19 23:24:30.592241 +0800 CST m=+1.006217543
13 2022-05-19 23:24:31.595953 +0800 CST m=+2.009938251
12 2022-05-19 23:24:31.596195 +0800 CST m=+2.010179834
19 2022-05-19 23:24:31.59632 +0800 CST m=+2.010305626
15 2022-05-19 23:24:31.596378 +0800 CST m=+2.010363543
17 2022-05-19 23:24:31.596405 +0800 CST m=+2.010390543
14 2022-05-19 23:24:31.596426 +0800 CST m=+2.010411168
16 2022-05-19 23:24:31.59645 +0800 CST m=+2.010435126
18 2022-05-19 23:24:31.596471 +0800 CST m=+2.010455918
11 2022-05-19 23:24:31.596165 +0800 CST m=+2.010149959
20 2022-05-19 23:24:31.59597 +0800 CST m=+2.009955501
29 2022-05-19 23:24:32.597286 +0800 CST m=+3.011279584
22 2022-05-19 23:24:32.597343 +0800 CST m=+3.011337376
28 2022-05-19 23:24:32.597385 +0800 CST m=+3.011378709
24 2022-05-19 23:24:32.597404 +0800 CST m=+3.011397959
26 2022-05-19 23:24:32.597413 +0800 CST m=+3.011406084
27 2022-05-19 23:24:32.597421 +0800 CST m=+3.011414751
23 2022-05-19 23:24:32.597426 +0800 CST m=+3.011419168
21 2022-05-19 23:24:32.597432 +0800 CST m=+3.011425584
25 2022-05-19 23:24:32.597406 +0800 CST m=+3.011400001
30 2022-05-19 23:24:32.598843 +0800 CST m=+3.012837043
```
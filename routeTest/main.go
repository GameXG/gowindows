package main

import (
	"fmt"
	"net"

	"github.com/GameXG/gowindows"
)

func main() {
	rows, err := gowindows.GetIpForwardTable()
	if err != nil {
		panic(err)
	}
	var defaultRow *gowindows.MibIpForwardRow

	for i, _ := range rows {
		row := &rows[i]
		fmt.Println(row.String())

		if row.GetForwardDest().Equal(net.IPv4(0, 0, 0, 0)) && net.IP(row.GetForwardMask()).Equal(net.IPv4(0, 0, 0, 0)) {
			// 默认网关
			if defaultRow != nil && row.ForwardMetric1 > defaultRow.ForwardMetric1 {
				continue
			}
			defaultRow = row
		}
	}

	if defaultRow == nil {
		fmt.Println("未找到默认网关。")
	}

	fmt.Println("默认网关：", defaultRow.String())

	r1 := *defaultRow
	err = r1.SetForwardDest(net.IPv4(1, 1, 1, 0))
	if err != nil {
		panic(err)
	}
	err = r1.SetForwardMask(net.IPv4Mask(255, 255, 255, 0))
	if err != nil {
		panic(err)
	}
	r1.ForwardMetric1 = 40
	fmt.Printf("添加 %#v\r\n", r1)
	err = gowindows.CreateIpForwardEntry(&r1)
	if err != nil {
		panic(err)
	}

	rows, err = gowindows.GetIpForwardTable()
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		fmt.Println(row.String())
	}

	fmt.Println("等待...")
	v := ""
	fmt.Scanln(&v)

	err = gowindows.DeleteIpForwardEntry(&r1)
	if err != nil {
		panic(err)
	}
}

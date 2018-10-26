package main

import "fmt"

type IPAddr [4]byte //array

// TODO: Add a "String() string" method to IPAddr.

func (i *IPAddr) String() string {
	var result string = ""
	for k, _ := range *i {
		str := fmt.Sprint(i[k])
		if k < len(i)-1 {
			result = result + str + "."
		} else {
			result = result + str
		}

	}
	return result
}

func main() {
	hosts := map[string]*IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
// 就是将String() string 重写即可。

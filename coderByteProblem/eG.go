package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type GoAssignment struct{}

func main() {
	fmt.Println("#########################################################")
	fmt.Printf("Started the execution @ %v \n", time.Now())
	rawdata := GoAssignment{}.loadRawDataList()
	GoAssignment{}.doTheWork(rawdata)
	fmt.Printf("\nFinished the execution @ %v\n", time.Now())
}

// doTheWork handles the required processing
func (ga GoAssignment) doTheWork(rawDataList []string) {
	ipPageTimeMap := make(map[string][]int)

	for _, row := range rawDataList {
		// Split by $#$
		parts := strings.Split(row, "$#$")
		if len(parts) < 3 {
			continue
		}

		// Remove the UUID
		fields := strings.Split(parts[1], "~$~")
		fields = append([]string{parts[1]}, strings.Split(parts[2], "~$~")...)

		// Re-join parts if there are more after 3rd split
		for i := 3; i < len(parts); i++ {
			fields = append(fields, strings.Split(parts[i], "~$~")...)
		}

		data := make(map[string]string)
		for _, field := range fields {
			kv := strings.Split(field, "#$#")
			if len(kv) == 2 {
				data[kv[0]] = kv[1]
			}
		}

		ip := data["ip"]
		pageLoadTimeStr := data["PageLoadTime"]

		if ip != "" && pageLoadTimeStr != "" {
			pageLoadTime, err := strconv.Atoi(pageLoadTimeStr)
			if err == nil {
				ipPageTimeMap[ip] = append(ipPageTimeMap[ip], pageLoadTime)
			}
		}
	}

	// Print Unique IP Count
	fmt.Printf("\nNumber of Unique IPs: %d\n", len(ipPageTimeMap))

	// Sort IPs
	var sortedIPs []string
	for ip := range ipPageTimeMap {
		sortedIPs = append(sortedIPs, ip)
	}
	sort.Strings(sortedIPs)

	// Print average PageLoadTime per IP
	fmt.Println("\nAverage PageLoadTime per IP:")
	for _, ip := range sortedIPs {
		times := ipPageTimeMap[ip]
		sum := 0
		for _, t := range times {
			sum += t
		}
		avg := float64(sum) / float64(len(times))
		fmt.Printf("IP: %s -> Avg PageLoadTime: %.2f ms\n", ip, avg)
	}
}

// loadRawDataList returns the test data
func (ga GoAssignment) loadRawDataList() []string {
	return []string{
		"a121a4cb$#$ip#$#190.25.228.161~$~PageLoadTime#$#1748",
		"bc940d03$#$ip#$#190.25.228.161~$~PageLoadTime#$#1522",
		"17440ba6$#$ip#$#190.25.228.161~$~PageLoadTime#$#2938",
		"7d732744$#$ip#$#190.25.228.161~$~PageLoadTime#$#1350",
		"e57e7965$#$ip#$#190.25.228.161~$~PageLoadTime#$#2611",
		"703f7517$#$ip#$#192.168.1.1~$~PageLoadTime#$#1500",
	}
}

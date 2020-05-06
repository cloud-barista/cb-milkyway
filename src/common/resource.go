package common

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"strconv"
	"github.com/labstack/echo"
	"regexp"
	
)

type benchInfo struct {
	Result string `json:"result"`
	Elapsed string `json:"elapsed"`
}

func RestGetInstall(c echo.Context) error {


	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	// Init fileio
	cmdStr := "wget https://github.com/cloud-barista/cb-milkyway/raw/master/src/script/install.sh -P ~/script/"
	result, err := SysCall(cmdStr)
	if(err != nil){
		mapA := map[string]string{"message": "Error in installation: wget script " + err.Error()}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	// Init DB
	cmdStr = "sudo chmod 755 ~/script/install.sh"
	result2, err := SysCall(cmdStr)
	if(err != nil){
		mapA := map[string]string{"message": "Error in installation: chmod " + err.Error()}
		return c.JSON(http.StatusNotFound, &mapA)
	}
	
	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)

	result += result2

	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetInit(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	// Init fileio
	cmdStr := "sysbench fileio --file-total-size=50M prepare"
	result, err := SysCall(cmdStr)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: Init fileio " + err.Error()}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	// Init DB
	cmdStr = "sysbench /usr/share/sysbench/oltp_read_write.lua --db-driver=mysql --table-size=100000 --mysql-db=sysbench --mysql-user=sysbench --mysql-password=psetri1234ak prepare"
	result2, err := SysCall(cmdStr)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: Init DB " + err.Error()}
		return c.JSON(http.StatusNotFound, &mapA)
	}
	
	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)

	result += result2

	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetClean(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	// Crean fileio
	cmdStr := "sysbench fileio --file-total-size=50M cleanup"
	result, err := SysCall(cmdStr)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: Crean fileio " + err.Error()}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	// Crean DB
	cmdStr = "sysbench /usr/share/sysbench/oltp_read_write.lua --db-driver=mysql --table-size=100000 --mysql-db=sysbench --mysql-user=sysbench --mysql-password=psetri1234ak cleanup"
	result2, err := SysCall(cmdStr)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: Crean DB "  + err.Error()}
		return c.JSON(http.StatusNotFound, &mapA)
	}
	
	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)

	result += result2

	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetCPU(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	cmdStr := "sysbench cpu --cpu-max-prime=10000 run"
	result, err := SysCall(cmdStr)

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	var grepStr = regexp.MustCompile(`execution time \(avg/stddev\):(\s+[+-]?([0-9]*[.])?[0-9]+)/`)
	parseStr := grepStr.FindStringSubmatch(result)	
	if len(parseStr) > 0 {
		parseStr1 := strings.TrimSpace(parseStr[1])
		fmt.Printf("execution time: %s\n", parseStr1)

		result = parseStr1
	}
	
	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetMEMR(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	cmdStr := "sysbench memory --memory-block-size=1K --memory-scope=global --memory-total-size=10G --memory-oper=read run"
	result, err := SysCall(cmdStr)

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	var grepStr = regexp.MustCompile(`execution time \(avg/stddev\):(\s+[+-]?([0-9]*[.])?[0-9]+)/`)
	parseStr := grepStr.FindStringSubmatch(result)	
	if len(parseStr) > 0 {
		parseStr1 := strings.TrimSpace(parseStr[1])
		fmt.Printf("execution time: %s\n", parseStr1)

		result = parseStr1
	}
	
	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetMEMW(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	cmdStr := "sysbench memory --memory-block-size=1K --memory-scope=global --memory-total-size=10G --memory-oper=write run"
	result, err := SysCall(cmdStr)

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	var grepStr = regexp.MustCompile(`execution time \(avg/stddev\):(\s+[+-]?([0-9]*[.])?[0-9]+)/`)
	parseStr := grepStr.FindStringSubmatch(result)	
	if len(parseStr) > 0 {
		parseStr1 := strings.TrimSpace(parseStr[1])
		fmt.Printf("execution time: %s\n", parseStr1)

		result = parseStr1
	}
	
	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetFIOR(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	cmdStr := "sysbench fileio --file-total-size=50M --file-test-mode=rndrd --max-time=30 --max-requests=0 run"
	result, err := SysCall(cmdStr)

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	var grepStr = regexp.MustCompile(`read, MiB/s:(\s+[+-]?([0-9]*[.])?[0-9]+)`)
	parseStr := grepStr.FindStringSubmatch(result)	
	if len(parseStr) > 0 {
		parseStr1 := strings.TrimSpace(parseStr[1])
		fmt.Printf("Throughput read, MiB/s: %s\n", parseStr1)

		result = parseStr1
	}
	
	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetFIOW(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	cmdStr := "sysbench fileio --file-total-size=50M --file-test-mode=rndwr --max-time=30 --max-requests=0 run"
	result, err := SysCall(cmdStr)

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	var grepStr = regexp.MustCompile(`written, MiB/s:(\s+[+-]?([0-9]*[.])?[0-9]+)`)
	parseStr := grepStr.FindStringSubmatch(result)	
	if len(parseStr) > 0 {
		parseStr1 := strings.TrimSpace(parseStr[1])
		fmt.Printf("Throughput write, MiB/s: %s\n", parseStr1)

		result = parseStr1
	}
	
	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func RestGetDBR(c echo.Context) error {

	if(checkInit() != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark: not initialized"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	content := benchInfo{}

	start := time.Now()

	fmt.Println("===============================================")

	cmdStr := "sysbench /usr/share/sysbench/oltp_read_only.lua --db-driver=mysql --table-size=100000 --mysql-db=sysbench --mysql-user=sysbench --mysql-password=psetri1234ak run"
	result, err := SysCall(cmdStr)

	elapsed := time.Since(start)
	elapsedStr := strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64)
	if(err != nil){
		mapA := map[string]string{"message": "Error in excuting the benchmark"}
		return c.JSON(http.StatusNotFound, &mapA)
	}

	var grepStr = regexp.MustCompile(`transactions:(\s+([0-9]*)(\s+)\([+-]?([0-9]*[.])?[0-9]+)`)
	parseStr := grepStr.FindStringSubmatch(result)	
	if len(parseStr) > 0 {
		
		parseStr1 := strings.Split(parseStr[1], "(")
		fmt.Printf("DB Read Transactions/s: %s\n", parseStr1[1])

		result = parseStr1[1]
	}
	
	content.Result = result
	content.Elapsed = elapsedStr 

	PrintJsonPretty(content)
	fmt.Println("===============================================")

	return c.JSON(http.StatusOK, &content)
}

func checkInit() error {
	checkPath, err := SysLookPath("sysbench")
	if(err != nil){
		return err
	}
	fmt.Printf("checkPath: %s\n", checkPath)
	return nil
}



func ApiValidation() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Printf("%v\n", "[API request!]")

			/*
			checkPath, err := SysLookPath("sysbench")
			mapA := map[string]string{"message": "Error in excuting the benchmark: no sysbench"}	
			if(err != nil){
				return echo.NewHTTPError(http.StatusNotFound, &mapA)
			}
			fmt.Printf("checkPath: %s\n", checkPath)
			*/

			return next(c)
		}
	}
}
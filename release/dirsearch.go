package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
 
	"github.com/fatih/color"
)

var wg sync.WaitGroup

func main() {
	clearScreen()
	printLogo()

	domainPtr := flag.String("d", "", "Domain to check (e.g., https://example.com/)")
	domainLongPtr := flag.String("domain", "", "Domain to check (e.g., https://example.com/)")
	wordlistPtr := flag.String("w", "", "Path to the wordlist file")
	wordlistLongPtr := flag.String("wordlist", "", "Path to the wordlist file")
	flag.Parse()

	if *domainPtr == "" && *domainLongPtr == "" || *wordlistPtr == "" && *wordlistLongPtr == "" {
		fmt.Println("Usage:   dirsearch -d|--domain <domain> -w|--wordlist <wordlist>")
		fmt.Println("Example: dirsearch --domain https://example.com/ --wordlist dirs.txt\n")
		os.Exit(1)
	}

	domain := *domainPtr
	if *domainLongPtr != "" {
		domain = *domainLongPtr
	}

	wordlist := *wordlistPtr
	if *wordlistLongPtr != "" {
		wordlist = *wordlistLongPtr
	}

	wordlistFile, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		os.Exit(1)
	}
	defer wordlistFile.Close()

	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		directory := scanner.Text()
		wg.Add(1)
		go checkDirectory(domain, directory)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
		os.Exit(1)
	}
}

func checkDirectory(domain, directory string) {
	defer wg.Done()
	url := constructURL(domain, directory)
	statusCode := getRequestStatusCode(url)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	statusColor := color.New(color.Bold, color.FgHiRed).SprintFunc()
	boldFont := color.New(color.Bold).SprintFunc()
	if statusCode == http.StatusOK {
		fmt.Printf("%s  %s  %s\n", boldFont(timestamp), statusColor("VALID"), boldFont(url))
	} else {
		// fmt.Printf("")
		// fmt.Printf("%s  %s  %s\n", boldFont(timestamp), color.RedString("INVALID"), boldFont(url))
	}
}

func printLogo() {
	logo := ` ____  _____ _____ _____ _____ _____ _____ _____ _____ {1.0.1}
|    \|     | __  |   __|   __|  _  | __  |     |  |  |
|  |  |-   -|    -|__   |   __|     |    -|   --|     |
|____/|_____|__|__|_____|_____|__|__|__|__|_____|__|__|
`
	fmt.Println(color.New(color.Bold).Sprint(logo))
}

func clearScreen() {
	cmd := exec.Command(clearCommand())
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearCommand() string {
	if runtime.GOOS == "windows" {
		return "cls"
	}
	return "clear"
}

func constructURL(domain, directory string) string {
	if strings.HasSuffix(domain, "/") {
		return fmt.Sprintf("%s%s", domain, directory)
	}
	return fmt.Sprintf("%s/%s", domain, directory)
}

func getRequestStatusCode(url string) int {
	response, err := http.Get(url)
	if err != nil {
		return -1
	}
	defer response.Body.Close()
	return response.StatusCode
}

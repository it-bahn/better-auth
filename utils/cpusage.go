package utils

import (
	"bytes"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Process struct {
	pid int
	cpu float64
}

func GetCPUInformation() {
	log.Printf("Getting CPU Information")
	log.Printf("OS: %s", runtime.GOOS)
	//Get MAX CPU CORES
	log.Printf("Max CPU Cores: %v", runtime.NumCPU())
	//Get MAX CPU FREQUENCY
	log.Printf("Runtime GOARCH: %v", runtime.GOARCH)
	//Get MAX CPU THREADS
	log.Printf("Max CPU Threads: %v", runtime.GOMAXPROCS(0))
	GetRamUsage(runtime.GOOS)
}
func GetProcessInfo() {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	for _, p := range processes {
		log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	}
}

func GetRamUsage(osName string) {

	if osName == "linux" {
		cmd := exec.Command("free", "-m")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		for _, line := range strings.Split(out.String(), "\n") {
			log.Printf("Ram Usage: %s", line)
		}
		GetProcessInfo()
	} else if osName == "windows" {
		capacity := exec.Command("wmic", "memorychip", "get", "Capacity")
		cfgClockSpeed := exec.Command("wmic", "memorychip", "get", "ConfiguredClockSpeed")
		interDepth := exec.Command("wmic", "memorychip", "get", "InterleaveDataDepth")
		commands := []*exec.Cmd{capacity, cfgClockSpeed, interDepth}
		var out bytes.Buffer
		for _, cmd := range commands {
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			for _, line := range strings.Split(out.String(), "\n") {
				if line != "" {
					log.Printf("%s", line)
				}
			}
		}
	}

}

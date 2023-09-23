package modules

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"unsafe"
)

// retrieve system information on client and send it to server

func fetchInfo() {
	// get computer name
	computerName, _ := os.Hostname()
	fmt.Print("INFO: Computer Name: \n", computerName)

	// get IP address
	ipaddr, _ := net.InterfaceAddrs()
	for _, address := range ipaddr {
		fmt.Printf("INFO: IP Address: %s\n", address)
	}

	// get MAC address
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		mac := iface.HardwareAddr
		fmt.Printf("INFO: MAC Address: %s\n", mac)
	}

	// get RAM capacity
	var memStatus syscall.MemoryStatusEx
	memStatus.Length = uint32(unsafe.Sizeof(memStatus))
	syscall.GlobalMemoryStatusEx(&memStatus)
	ramCapacity := memStatus.TotalPhys
	fmt.Printf("INFO: RAM Capacity: %d\n", ramCapacity)

	// get Disk capacity
	var freeBytesAvailable, totalBytes, totalFreeBytes int64
	syscall.GetDiskFreeSpaceEx("C:\\", &freeBytesAvailable, &totalBytes, &totalFreeBytes)
	diskSpace := totalBytes / (1024 * 1024 * 1024)
	fmt.Printf("INFO: Disk Capacity: %d\n", diskSpace)

	// get OS version
	var osInfo syscall.RtlOsVersionInfoEx
	info.dwOSVersionInfoSize = uint32(unsafe.Sizeof(info))

	err := syscall.RtlGetVersion((*syscall.RtlOsVersionInfoEx)(unsafe.Pointer(&info)))
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Printf("INFO: OS Version: %d.%d\n", info.dwMajorVersion, info.dwMinorVersion)
	fmt.Println("Build Number: ", info.dwBuildNumber)
	fmt.Println("Platform ID: ", info.dwPlatformId)

	// get CPU info
	var systemInfo syscall.SystemInfo
	syscall.GetSystemInfo(&systemInfo)

	fmt.Println("INFO: CPU Info: ")
	fmt.Println("Processor Architecture: ", systemInfo.ProcessorArchitecture)
	fmt.Println("Number of Processors: ", systemInfo.NumberOfProcessors)
	fmt.Println("Processor Type: ", systemInfo.ProcessorType)

}

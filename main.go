// bunfetch developed by florencio 2022

package main

import (

    "fmt"
    "github.com/shirou/gopsutil/v3/host"

)

func main() {

    // gets system info and saves it to hostname variable, also errors to err variable
    hostname, err := host.Info()

    // checks if theres any errors with the err variable
    if err != nil {
        fmt.Println(err)
    }

    // prints out ascii art w/ info
    fmt.Println(fmt.Sprintf(` (\_/)      os: %+v %+v
 ( . .)     kernel: %+v 
 C(")(")`, hostname.Platform, hostname.KernelArch, hostname.KernelVersion))


}

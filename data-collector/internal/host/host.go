package host

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

func GetInformation() *host.InfoStat {
	host, err := host.Info()
	if err != nil {
		// TODO: Better logging
		fmt.Println(err.Error())
		return nil
	}
	return host
}

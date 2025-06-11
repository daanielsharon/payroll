package httphelper

import (
	"shared/config"
	"shared/constant"
)

type Server struct {
	Name   string
	Port   string
	Target string
}

func RegisteredServers(cfg config.ApplicationConfig) []Server {
	server := cfg.Server
	host := server.Host
	return []Server{
		{
			Name:   constant.ServicePayroll,
			Port:   server.PayrollPort,
			Target: "http://" + host + ":" + server.PayrollPort,
		},
		{
			Name:   constant.ServiceOvertime,
			Port:   server.OvertimePort,
			Target: "http://" + host + ":" + server.OvertimePort,
		},
		{
			Name:   constant.ServiceAttendance,
			Port:   server.AttendancePort,
			Target: "http://" + host + ":" + server.AttendancePort,
		},
		{
			Name:   constant.ServiceReimbursement,
			Port:   server.ReimbursementPort,
			Target: "http://" + host + ":" + server.ReimbursementPort,
		},
		{
			Name:   constant.ServiceAuth,
			Port:   server.AuthPort,
			Target: "http://" + host + ":" + server.AuthPort,
		},
		{
			Name:   constant.ServiceUser,
			Port:   server.UserPort,
			Target: "http://" + host + ":" + server.UserPort,
		},
	}
}

func GetServiceMap(registeredServers []Server) map[string]string {
	serviceMap := make(map[string]string)
	for _, server := range registeredServers {
		serviceMap[server.Name] = server.Target
	}
	return serviceMap
}

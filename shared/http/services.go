package httphelper

import "shared/config"

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
			Name:   "payroll",
			Port:   server.PayrollPort,
			Target: "http://" + host + ":" + server.PayrollPort,
		},
		{
			Name:   "overtime",
			Port:   server.OvertimePort,
			Target: "http://" + host + ":" + server.OvertimePort,
		},
		{
			Name:   "attendance",
			Port:   server.AttendancePort,
			Target: "http://" + host + ":" + server.AttendancePort,
		},
		{
			Name:   "reimbursement",
			Port:   server.ReimbursementPort,
			Target: "http://" + host + ":" + server.ReimbursementPort,
		},
		{
			Name:   "auth",
			Port:   server.AuthPort,
			Target: "http://" + host + ":" + server.AuthPort,
		},
		{
			Name:   "user",
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

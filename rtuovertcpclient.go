package modbus

// RTUoverTCPClientHandler implements Packager and Transporter interface.
type RTUoverTCPClientHandler struct {
	rtuPackager
	tcpTransporter
}

// NewRTUoverTCPClientHandler allocates a new RTUoverTCPClientHandler.
func NewRTUoverTCPClientHandler(address string) *RTUoverTCPClientHandler {
	handler := &RTUoverTCPClientHandler{}
	handler.Address = address
	return handler
}

// RTUoverTCPClient creates RTU over TCP client with default handler and given connect string.
func RTUoverTCPClient(address string) Client {
	handler := NewRTUoverTCPClientHandler(address)
	return NewClient(handler)
}

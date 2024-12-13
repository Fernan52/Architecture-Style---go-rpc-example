package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Definimos el servicio
type HelloService struct{}

// Método expuesto por el servicio
func (h *HelloService) SayHello(request string, reply *string) error {
	*reply = "Hola, " + request + "!"
	return nil
}

func main() {
	// Registramos el servicio
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		fmt.Println("Error al registrar el servicio:", err)
		return
	}

	// Escuchamos en el puerto
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor escuchando en el puerto 1234...")

	// Aceptamos conexiones entrantes
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

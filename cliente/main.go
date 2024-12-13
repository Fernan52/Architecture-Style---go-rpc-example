package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// Conexión al servidor RPC
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error al conectar con el servidor RPC:", err)
		return
	}
	defer func() {
		// Cerramos la conexión al servidor
		if err := client.Close(); err != nil {
			fmt.Println("Error al cerrar la conexión RPC:", err)
		}
	}()

	// Realizar llamada RPC al método SayHello
	var reply string
	err = client.Call("HelloService.SayHello", "Mundo", &reply)
	if err != nil {
		fmt.Println("Error al llamar al método RPC:", err)
		return
	}

	// Mostrar la respuesta del servidor
	fmt.Println("Respuesta del servidor:", reply)
}

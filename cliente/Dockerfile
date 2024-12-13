# Usamos una imagen base de Go
FROM golang:1.20-alpine

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos go.mod y go.sum para gestionar las dependencias
COPY go.mod go.sum ./

# Instalamos las dependencias de Go
RUN go mod tidy

# Copiamos el resto de los archivos fuente al contenedor
COPY . .

# Comando para ejecutar el cliente
CMD ["go", "run", "main.go"]

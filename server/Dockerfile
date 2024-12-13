# Usamos la imagen base de Go
FROM golang:1.20-alpine

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos el código fuente del servidor al contenedor
COPY . .

# Ejecutamos go mod tidy para gestionar dependencias
RUN go mod tidy

# Exponemos el puerto 1234
EXPOSE 1234

# Comando para ejecutar el servidor
CMD ["go", "run", "main.go"]

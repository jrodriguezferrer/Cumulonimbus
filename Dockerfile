# Imagen base
FROM golang:1.25.4

# Crear directorio de trabajo
WORKDIR /app

# Copiar archivos
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Puerto expuesto
EXPOSE 8080

# Comando de ejecución
CMD ["./main"]
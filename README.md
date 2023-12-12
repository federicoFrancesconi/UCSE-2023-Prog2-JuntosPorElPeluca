# UCSE-2023-Prog2-TPIntegrador
Trabajo Práctico Integrador de la materia Programación II de Ingeniería en Informática en UCSE

## Estructura
Tenemos tres directorios principales en el root del proyecto:
* go: es la API, el backend.
* web: es el frontend con HTML, CSS y JavaScript.
* data: archivos JSON para importar a mongoDB y agilizar las pruebas.

## Instrucciones para levantar el proyecto
1. Abrir una terminal parados en el root, y correr el comando `docker-compose up`.
2. En el explorador de preferencia, ingresar a `localhost:80` para visualizar el frontend.
### Para usar datos de prueba del directorio "data"
1. Abrir MongoDB Compass y crear las colecciones `pedidos`, `productos`, `camiones` y `envios` dentro de la base de datos `empresa` (se crea automáticamente luego de abrir el frontend).
2. En cada colección, importar el archivo .json con el mismo nombre

## Usuarios para tests
### Admin
* Mail: admin@gmail.com
* Contraseña: SoyAdmin123$
### Conductor
* Mail: conductor@gmail.com
* Contraseña: SoyConductor123$
### Operador
* Mail: operador@gmail.com
* Contraseña: SoyConductor123$

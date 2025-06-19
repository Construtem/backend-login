Backend - Login
Microservicio backend encargado de usar firebase para autenticacion de usuario y redirigirlo a su respectivo lugar de trabajo

🛠️ Tecnologías
- Go (Lenguaje)
- Gin (Framework)
- GORM (Libreria de Go para interactuar con bases de datos)
- Postgres

🚀 Funcionalidades
- Conexión con firebase
- Redireccionamiento según el rol del usuario a
    -  inventario.tssw.cl
    -  ventas.tssw.cl

Requisitos
Docker Desktop instalado
Git instalado

Instalación (entorno de desarrollo)
Clonar el repositorio en el directorio deseado:
Desde la terminal debe situarse en el directorio que desee clonar repo (ej: "C:\Users\Admin\Desktop") y ejecutar siguiente comando

**¿Cómo situarse en C:\Users\Admin\Desktop?**
git clone https://github.com/Construtem/backend-login
cd backend-login
Correr aplicación desde directorio creado (ej "C:\Users\Admin\Desktop\backend-login"), ejecutando el siguiente comando:
docker compose up
Luego de ejecutar este comando, su app se encontrará corriendo en el puerto 8080 en "http://localhost:8080"

Contribución
Crea una rama para tu funcionalidad/tarea:
git switch -c feature/<nombre-funcionalidad>
Realiza cambios y haz commit:
git add <archivos-cambiados>
git commit -m "<descripcion pequeña del cambio>"
Pushea tus cambios de la rama:
git push origin feature/<nombre-funcionalidad> 
Crea un Pull Request (PR) a la rama ´develop´ desde GitHub para que sea revisado por otro desarrollador

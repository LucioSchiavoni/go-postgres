# go-postgres

importar gorilla gorm:
"github.com/gorilla/mux"

importar carpetas: 
"github.com/LucioSchiavoni/go-postgres/routes"

para usar variables de entorno:
"github.com/joho/godotenv"


endpoints: 

GET :
"/users"  - todos los usuarios

"/user/{id}"  - usuario por id

POST:
"/user" - crear usuario 
raw:
{
    "username": "",
    "email": "",
    "password":""
}

DELETE:
"/user" - eliminar usuario 
raw:
{
    "ID": 1
}


Posteos de usuarios:

POST:
"/post" -usuario crea post
raw:
{
    "title": "",
    "description": "",
    "userId": 
}

GET:
"/post/3"  - obtener por su id de posteo

GET:
"/description/user_id/description" (cambiar parmetros) - obtener post donde su id de usuario es x y descripcion es x


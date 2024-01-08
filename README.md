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

# go-postgres

> .[!TIP].
> Comando para ejecutar como dev:
air 

> .[!NOTE].
> Importar gorilla gorm:
"github.com/gorilla/mux"

> .[!NOTE].
> Importar carpetas: 
"github.com/LucioSchiavoni/go-postgres/routes"

> .[!NOTE].
> Para usar variables de entorno: 
"github.com/joho/godotenv"

> .[!IMPORTANT].
> Endpoints: 
 
GET :
"/users"  - todos los usuarios

"/user/{id}"  - usuario por id

"/login" - logear usuario con su email y password (devuelve el token)

"/auth" - autorizacion que requiere el token para autenticar usuario (devuelve los datos del usuario)


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


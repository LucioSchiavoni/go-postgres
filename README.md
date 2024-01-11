# Backend con go y postgres

> [!TIP]
>comando para ejecutar como dev:
air 

> [!NOTE]
>importar gorilla gorm:
"github.com/gorilla/mux"

> [!NOTE]
>importar carpetas: 
"github.com/LucioSchiavoni/go-postgres/routes"

> [!NOTE]
>para usar variables de entorno: 
"github.com/joho/godotenv"

> [!IMPORTANT]
>Agregar carpeta temp-images

> [!IMPORTANT]
>Endpoints: 

GET :
"/users"  - todos los usuarios

"/user/{id}"  - usuario por id

"/login" - logear usuario con su email y password (devuelve el token)

"/auth" - autorizacion que requiere el token para autenticar usuario (devuelve los datos del usuario)


POST:
"/user" - crear usuario 
el alta de usuarios se encuentra en formato formdata

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
/usuario/{user_id}/posts  - buscar todos los posts de ese usuario con su user_id = ?

GET:
"/description/user_id/description" (cambiar parmetros) - obtener post donde su id de usuario es x y descripcion es x

GET:
/usuario/id_usuario/title/posts  buscar posts donde el id de usuario es ? y su title es ? 
ejemplo:
"/usuario/19/champions/posts"  

#VER ESTRUCTURA DEBAJO 

Para probar de manera local:

    1 - $git clone https://github.com/manavellam/TestingGorm
        El archivo go.mod instalara automaticamente las librerias utilizadas. Esto facilita la implementacion.

    2 - modificar el app/config.json con la info(User, Pass, DBname) para acceder a una base de datos, donde se creara la tabla "user"

    3 - En postman, colocar los siguiente en el Body para /login/auth:
        
            {
                "Name": "Matias",
                "Password": "1234"
            }

    4 - Probar los endpoints con el Token devuelto. Para agregar un User(/user/add) atencion a los corchetes:
            [	        
	            {
                    "Name": "Nicolas",
                    "Password": "4567"
	            }
            ]

Durante el proyecto, las funciones utilizan siempre que sea posible punteros. De esta forma se evita ocupar la memoria de manera innecesaria y aumenta la velocidad de respuesta.

Durante las consultas, el middleware provisto por gin tonic arrojara alguna info sobre que consulta llego, a que  endpoint y cuanto demoro la consulta.

Estructura del proyecto:

``` bash

TestingGorm
|_App
| |_config.json           Contiene informacion de configuracion. Ejemplo, puerto del server
|_Controllers
| |_LoginUser.go          Responde a /login/auth, loggea usuario dado un par valido de U+P, y que no este loggeado
| |_UserAdd.go            Responde a /user/add, agrega usuario a DB dado un token valido
| |_UserReadAll.go        Responde a /user/read, dado un tkn  valido, devuelve la info en DB users
|_Middleware
| |_IsTknValid.go         Checkea si el token existe, y si es valido
| |_IsUser.go             Checkea si un par U+P es valido
| |_TokenGen.go           Genera un Tkn con tiempo de expiracion para un usuario
|_Models
| |_ClaimsJWT.go          Struct utilizada para Claims a agregarse en un JWT (ExpiresAt, Issuer, etc...)
| |_User.go               Struct utilizada para leer y escribir en DB user con GORM
|_Routes
| |_loginpaths.go         Agrupa endpoints /login/...
| |_routes.go             Configura las routas en el router de GIN GONIC
| |_userpaths.go          Agrupa endpoints /user/...
|_Services
| |_Contains.go           Checks if database contains a pair of User+Pass
| |_DBInit.go             Initiates a  conection with the DB
| |_Insert.go             Inserts a new User into DB (Generico, funciona con cualquier DB y modelo)
| |_Query.go              SELECT * FROM "tablename" (Generico tambien)
|_main.go
```

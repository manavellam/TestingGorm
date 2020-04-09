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
| |_loginpaths.go         Agrupa paths /login/...
| |_routes.go             Configura las routas en el router de GIN GONIC
| |_userpaths.go          Agrupa paths /user/...
|_Services
| |_Contains.go           Checks if database contains a pair of User+Pass
| |_DBInit.go             Initiates a  conection with the DB
| |_Insert.go             Inserts a new User into DB (Generico, funciona con cualquier DB y modelo)
| |_Query.go              SELECT * FROM "tablename" (Generico tambien)
|_main.go
```

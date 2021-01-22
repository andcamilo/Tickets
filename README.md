# Tickets API

#Dockerfile 


 Para ejecutar el contenedor del docker debemos primero asegurarnos tener los permisos correspondientes, si se esta en un equipo linux usar:

  sudo su 

 Luego crear la imagen del mismo con el siguiente comando

  docker build -t main.go .
 
 Una vez hallamos creado la imagen verificamos que exista en el sistema con el siguiente comando
  
 docker image ls
 
 Y cuando veamos que la imagen ya esta creada dentro del sistema utilizamos el siguiente comando para ejecutar el contenedor.
  
  
   docker run -d -p8080:8080  --net="host" --name main.go main.go

 Ahora si podemos proceder a hacerle peticiones al mismo.



# API 


Esta api funciona para la creación, lectura, actualización y eliminación de tickets.

Api la cual puede ser consumida localmente por el puerto 8080 con la siguiente url  http://localhost:8080/
las peticioens se realizan mediante un header content-type application/json
y requiere como parametros de entrada

{ 
	"User": "nombre del usuario",
	"Status": "Estatus del ticker"
}

para cada acciones se utiliza diferentes tipos de peticiones los cuales seran explicadas a continuación:


   CREATE: METHOD = POST http://localhost:8080/tickets/create 

   READ ALL: METHOD = GET http://localhost:8080/tickets

   READ ONE TICKET: METHOD = GET http://localhost:3000/tickets/#

   DELETE: METHOD = DELETE http://localhost:8080/tickets/#

   UPDATE: METHOD = PUT http://localhost:8080/tickets/#  

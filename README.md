# Tickets API

#Dockerfile 




# API 


Esta api funciona para la creación, lectura, actualización y eliminación de tickets.

Api la cual puede ser consumida localmente por el puerto 3000 con la siguiente url  http://localhost:3000/tickets/
las peticioens se realizan mediante un header content-type application/json
y requiere como parametros de entrada

{ 
	"User": "nombre del usuario",
	"Status": "Estatus del ticker"
}

para cada acciones se utiliza diferentes tipos de peticiones los cuales seran explicadas a continuación:

   CREATE: METHOD = POST http://localhost:3000/tickets/create 

   READ ALL: METHOD = GET http://localhost:3000/tickets

   READ ONE TICKET: METHOD = GET http://localhost:3000/tickets/#

   DELETE: METHOD = DELETE http://localhost:3000/tickets/#

   UPDATE: METHOD = PUT http://localhost:3000/tickets/#  

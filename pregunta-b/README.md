## ¿Qué sucede normalmente al hacer range sobre un canal hasta que este se cierra?


a) Se reciben valores sucesivamente hasta que el canal se cierra.
b) Solo se recibe el primer valor por el programa ejecutado.
c) El canal se convierte en un slice.
d) Se crea automáticamente otra goroutine ejecutandose en paralelo.

## go run .\pregunta-b\main.go            
Recibido: Mensaje 1  
Recibido: Mensaje 2  
Recibido: Mensaje 3  
Canal cerrado, terminamos  
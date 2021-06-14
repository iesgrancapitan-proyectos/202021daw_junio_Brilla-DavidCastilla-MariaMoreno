# Módulo Proyecto Integrado del IES Gran Capitán

## Introducción
Este proyecto es nuestra pequeña red social creada por @dcxo y @mariamm99. Se llama **Brilla**, en ella podrán compartir contenido, ya sea en formato texto, imagen, vídeo o combinación de ellas.

Además podrás interactuar con las personas respondiendo a los brillos (comentarios) de otras personas o por interacciones (me gusta, me divierte, me aburre o me entristece).


## Requirimentos.
Solo necesitas docker-compose. 

Para ejecutar el proyecto entero en el modo de desarrollo simplemente hay que irse a la raíz del proyecto y ejecutar el comando `docker-compose up --build` y con esto ya estaría el proyecto ejecutándose en tu maquina local.

En el caso de ejecutar el proyecto entero en el modo de desarrollo volvemos a irnos a la raíz del proyecto y ejecutar el comando `docker-compose -f docker-compose.prod.yml up --build` y con esto ya estaría el proyecto ejecutándose en tu maquina local en ese entorno de producción.


## Herramientas software
 - **Go**: Go es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C, que intenta ser dinámico y con un gran rendimiento. Ha sido desarrollado por Google y es el lenguaje que utilizaremos para el backend. 
 - **ArangoDB**: Es una base de datos multi-modelo combinando los modelos: Clave/Valor, Documentos y Grafos en un solo núcleo desarrollado en C++ y haciendo uso de un lenguaje de consultas AQL (ArangoDB Query Language). Esto le permite realizar consultas entre los diferentes modelos de datos indistintamente.
 - **Svelte**: Svelte es un compilador frontal gratuito y de código abierto. Svelte tiene su propio compilador para convertir el código de la aplicación en JavaScript del lado del cliente en el momento de la compilación. Svelte nos permite unificar el uso de HTML5, SCSS y JavaScript.
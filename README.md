# Módulo Proyecto Integrado del IES Gran Capitán

## Descripción del proyecto
Este proyecto es nuestra pequeña red social creada por @dcxo y @mariamm99. Se llama **Brilla**, en ella podrán compartir contenido, ya sea en formato texto, imagen, vídeo o combinación de ellas.

Además podrás interactuar con las personas respondiendo a los brillos (comentarios) de otras personas o por interacciones (me gusta, me divierte, me aburre o me entristece).


## Instalación
Para instalar este proyecto solo necesitas docker-compose. 

Para ejecutar el proyecto entero en el modo de desarrollo simplemente hay que irse a la raíz del proyecto y ejecutar el comando `docker-compose up --build` y con esto ya estaría el proyecto ejecutándose en tu maquina local.

En el caso de ejecutar el proyecto entero en el modo de desarrollo volvemos a irnos a la raíz del proyecto y ejecutar el comando `docker-compose -f docker-compose.prod.yml up --build` y con esto ya estaría el proyecto ejecutándose en tu maquina local en ese entorno de producción.


## Información sobre cómo usarlo
Hemos intentado crear una red social intuitiva para que sea facil de usar. 

Para comenzar tienes que reguistrarte pulsando el boton de singup. Una vez registrado puedes acceder desde la página principal y te aparecerá el timeline.

En el timeline está el boton para crear contenido y el buscador para buscar brillos o usuarios. En el timelina aparecerá los brillos de forma cronologica inversa segun la fecha de su cración. Apareceran tus brillos, los brillos que has compartido, los brillos de los usuarios que sigues y los brillos compartidos por los usuarios que sigues.

Desde el perfil del resto de usuarios puedes seguir a otros usuarios para ver el contenido que crean o comparten (rebrillar) y desde el tuyo propio puedes editar nombre, foto de perfil, nombre de usuario, contraseña y biografía.

Para mas detalles de como usarlo puedes encontrarlo en el [Manual de usuario](https://github.com/iesgrancapitan-proyectos/202021daw_junio_Brilla-DavidCastilla-MariaMoreno/wiki/Manual_Usuario)


## Autores
David Castilla Ortíz
- github: @dcxo
- twitter: @dcxo

María Moreno Muñoz
- github: @mariamm99
- twitter: @mariamm9961
---
title: "Programmer"
date: 2021-10-29T11:52:55+02:00
draft: true
---

Desde la perspectiva de un programador la idea de herramienta universal puede
sonar disparatada. Estamos acostumbrados a enfrentarnos al software como un
todo que debe resolver el problema; aunque los microservicios esten cambiando
esta tendencia poco a poco.

Pero Overlord no es algo que se pueda "resolver" globalmente. Necesitamos una
aproximación inteligente, modular y extensible. Como un sistema operativo o un
compilador.

Como programador Overlord necesitará de pequeños trozos de código para poder
llevar a cabo sus tareas de recopilación de información. En estos casos cada
trozo de código debe tener dos partes, la consultiva y la ejecutiva.

## Consultando

Cuando creamos código para Overlord este necesita un interfaz para saber que
datos le podemos pedir. Esto se declara en los metadatos de registro del trozo de
código; si por ejemplo necesitamos nuestro codigo provee de pan asi lo 
declararemos.

{{< vid "/programmer/query.mp4" >}}

Las propiedades tienen espacio de nombre (net) es simplemente porque Overlord
puede entender la ip en muchos contextos. Por ejemplo, puede ser una ip de AWS,
la propiedad en ese caso debería llamarse algo parecido a "ip.aws.com". No te
preocupes mucho por el nombre ahora, no necesita una entidad central que asigne
nombres, aunque los nombres los pones tu fijarse en la comunidad puede ser muy
ventajoso.

Ahora que Overlord sabe que puedes ofrecer al sistema te enviará peticiones con
un conjunto de propiedades, por ejemplo, una mascara de red, un nombre de
dominio y otras cosas. Ten en cuenta que recibiras todas las propiedades
posibles, no solo las que te interesen a ti. Si con estas propiedades puedes
recopilar la información, tu código le podrá pasar el nombre completo de la
propiedad que provees, y ya has terminado la parte consultiva.

{{< vid "/programmer/resolved.mp4" >}}

En el caso de que no puedas proveer de esa información porque te faltan
propiedades, deberás devolver que propiedades necesitas para llevar a cabo tu
tarea. Y estas pueden ser muchas, ya que puedes llevar a cabo tu tarea de
diferentes formas.

{{< vid "/programmer/partial.mp4" >}}

Una capacidad diferenciadora de los grafos tradiciones es que el código atómico
de consulta puede gestionar como considere una petición entrante. No se espera
de el que de siempre la misma respuesta, esta puede estar basada en el entorno
que ocupa. Pero si que se espera que de siempre la misma respuesta en el código
ejecutivo. Esto permite tener una gran flexibilidad frente a casos que todavía
no conocemos.

Por ejemplo, en lugar de responder siempre, necesito agua, harina y sal para
hacer pasta, puede que el peticionario lleve productos frescos entre sus
propiedades. El código de consulta puede decidir agregar huevos a la lista de
requisitos para entregar pasta fresca en lugar de pasta seca normal.

También ante las mismas propiedades, pero diferente usuario, podemos requerir
otras cosas. Si entra, por ejemplo, el dueño de la tienda, no le cobraremos
nada por el servicio. Por otro lado si entra un espia (sin permisos en el
sistema objetivo) podemos exigirle que pase antes por comisaría a ver si esta
autorizado a llevarse la receta secreta para la pasta.

## Ejecutando

Cuando Overlord invoca el código ejecutivo te dará, de nuevo y de forma
simétrica al anterior caso, todas las propiedades que posee en ese momento.
Entre estas seguro que está la forma de conseguir la información. Queda en tu
mano decidir que va a pasar, si puedes conseguir la información de más de una
forma.

{{< vid "/programmer/execution.mp4" >}}

Después de ejecutar tu código le entragarás la información a Overlord. Así de
simple (seguramente mediante un API de intercambio por decidir todavía, pero
siempre de la forma más aislada posible para no introducir dependencias en el
código atómico)

# Poninedo el proceso en perspectiva

Este proceso parece bastante sencillo, pero hay que tener en cuenta que la
mayoría de los procesos se basan en entrada y salida de datos; por complejos
que sean. Usando estos pequeños códigos Overlord es capaz de componer
soluciones de mucha complejidad.

En ocasiones te sentirás tentado a crear grandes piezas de código, pero a no
ser que quieras optimizar una ejecución que le esta llevando mucho tiempo a
Overlord te recomendamos manter las cosas lo más simples posible para ti.

# Valor Aportado

Overlord a través de sus comprobaciones puede ser un framework de comprobación
de infinidad de cosas en tu código. No tiene porque usarse a nivel empresarial,
ni es una herramienta que requiera de grandes cantidades de infraestructura.
Queremos mantenerla simple para que se pueda usar en solitario y ofrecer una
capacidad de testing más exaustiva.

Sabemos que hay frameworks de "Property based testing", y seguramente te
encuentres muy cómodo con ellos. Incluso es posible que te oferzcan mejores
abstracciones para ti. Pero ten en cuenta dos cosas:

 * Cuando creas código para Overlord aportas a una comunidad más allá de tu
   caso de testing.
 * Si ese framework es tan bueno ¡puedes integrarlo en Overlord! y aportar así
   a una comunidad más grande.

Intentamos que el esfuezo de integración sea lo menor posible, para que ese
valor aportado tenga un coste bajo.

Además de comprobar controles de seguridad puedes usar las abstracciones de
Overlord como herramientas. Creemos que el grafo dinámico puede adaptarse una
cantidad ingente de casos, incluso en entornos de alta incertidumbre. Siempre
que el tiempo que necesites para obtener una respuesta no sea crítico Overlord
puede ser una solución muy flexible.

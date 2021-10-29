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
datos le podemos dar. Esto se declara en los metadatos de registro del trozo de
código; si por ejemplo nuestro código provee de direcciones IP activas en una
red debemos declarar que la propiedad que devolvemos es "ip.net".

En el porque las propiedades tienen espacio de nombre (net) es simplemente
porque Overlord puede entender la ip en muchos contextos. Por ejemplo, puede
ser una ip de AWS, la propiedad en ese caso debería llamarse algo parecido a
"ip.aws.com". No te preocupes mucho por el nombre ahora, no hay una entidad
central que asigne nombres, aunque los nombres los pones tu fijarse en la
comunidad puede ser muy ventajoso.

Ahora que Overlord sabe que puedes ofrecer al sistema te enviará peticiones con
una conjunto de propiedades, por ejemplo, una mascara de red, un nombre de
dominio y otras cosas. Ten en cuenta que recibiras todas las propiedades
posibles, no solo las que te interesen a ti. Si con estas propiedades puedes
recopilar la información tu código le podrá pasar el nombre completo de la
propiedad que provees, y ya has terminado la parte consultiva.

En el caso de que no puedas proveer de esa información porque te faltan
propiedades, deberás devolver que propiedades necesitas para llevar a cabo tu
tarea. Y estas pueden ser muchas, ya que puedes lleva a cabo tu tarea de
diferentes formas.

Es decisión tuya de elegir solo una forma de proveer la información básandote
en las propiedades que vienen en lugar de varias formas. Cada trozo de código
es una mundo y tu caso puede requerir de cierta rigidez. Esto abre la puerta a
que puedas establecer cierta forma de hacer las cosas ante cierta propiedad;
como por ejemplo, si te viene la dirección de un servidor dns solo requerir los
datos necesarios para comprobar las ip con esa información.

## Ejecutantdo

Cuando Overlord te pide la información te dará, de nuevo y de forma simétrica
al anterior caso, todas las propiedades que posee en ese momento. Entre estas
seguro que está la forma de conseguir la información. Queda en tu mano decidir
que va a pasar si entre estas propiedades puedes conseguir la información de
más de una forma.

Después de ejecutar tu código le entragarás la información a Overlord. Así de
simple.

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

---
title: "Programmer"
date: 2021-10-29T11:52:55+02:00
draft: true
---

Desde la perspectiva de un programador la idea de herramienta universal puede sonar disparatada. Estamos acostumbrados a enfrentarnos al software como un todo que debe resolver el problema; aunque los microservicios estén cambiando esta tendencia poco a poco.

Pero la búsqueda automática de caminos e información auxiliar que necesita Overlord para que pueda desarrollarse de manera independiente no es algo que se pueda "resolver" globalmente. Necesitamos una aproximación inteligente, modular y extensible. Como un sistema operativo o un compilador.

Overlord necesitará de pequeños trozos de código para poder llevar a cabo sus tareas de recopilación de información. En estos casos, cada trozo de código debe tener dos partes, la consultiva y la ejecutiva.

## Consultando

Cuando creamos código para Overlord este necesita un interfaz para saber qué datos le podemos pedir. Esto se declara en los metadatos de registro del trozo de código; si por ejemplo necesitamos que nuestro código provea de pan, así lo declararemos.

{{< vid "/programmer/query.mp4" >}}

Las propiedades tienen espacio de nombre (por ejemplo ip.net) simplemente porque Overlord puede entender la ip en muchos contextos. Por ejemplo, puede ser una ip de AWS, la propiedad en ese caso debería llamarse algo parecido a "ip.aws.com". No te preocupes mucho por el nombre ahora, no necesita una entidad central que asigne nombres. Aunque los nombres los pones tú, fijarse en la comunidad puede ser muy ventajoso.

Ahora que Overlord sabe qué puedes ofrecer, el sistema te enviará peticiones con un conjunto de propiedades, por ejemplo, una máscara de red, un nombre de dominio y otras cosas. Ten en cuenta que recibirás todas las propiedades posibles, no solo las que te interesen a ti. Si con estas propiedades puedes recopilar la información, tu código le podrá pasar el nombre completo de la propiedad que provees, y ya has terminado la parte consultiva.

{{< vid "/programmer/resolved.mp4" >}}

En el caso de que no puedas proveer de esa información porque te faltan parámetros, deberás devolver qué propiedades necesitas para llevar a cabo tu tarea. Y estas pueden ser muchas, ya que puedes llevar a cabo tu tarea de diferentes formas.

{{< vid "/programmer/partial.mp4" >}}

Una capacidad diferenciadora respecto a los grafos tradicionales es que el código atómico de consulta puede gestionar como considerar una petición entrante. No se espera de él que proporcione siempre la misma respuesta, esta puede estar basada en múltiples factores, de hecho por eso es “dinámico”. Se espera del código que tenga siempre la misma salida, pero debe ser flexible en la entrada. Esto permite tener una gran flexibilidad frente a casos que todavía no conocemos.


También ante los mismos parámetros, pero diferente usuario, podemos requerir otras cosas. Si la consulta, por ejemplo, la hace el administrador del sistema no requerimos ningún tipo de permiso adicional. Si por otro lado un usuario anónimo hace la misma consulta podemos requerir un parámetro de autenticación en el sistema.

## Ejecutando

Cuando Overlord invoca el código ejecutivo este recibirá, de nuevo y de forma simétrica al anterior caso, todos los parámetros que posee en ese momento. Entre éstas, seguro que está la forma de conseguir la información. Si puedes conseguir la información de más de una forma, queda en tu mano decidir cómo se ejecutará para conseguir su objetivo.

{{< vid "/programmer/execution.mp4" >}}

Una vez ejecutado el código éste entregará el resultado a Overlord. Así de simple (seguramente mediante un API de intercambio por decidir todavía, pero siempre de la forma más aislada posible para no introducir dependencias en el código atómico)

# Poniendo el proceso en perspectiva

Este proceso parece bastante sencillo, pero utilizando estos pequeños fragmentos de código (que pueden invocar aplicaciones más complejas como SAST), Overlord puede componer soluciones muy complejas.

# Valor Aportado

Overlord a través de sus comprobaciones puede ser un framework de comprobación generalista en tu propio código. No tiene por qué usarse solo a nivel empresarial, ni es una herramienta que requiera de grandes cantidades de infraestructura. Queremos mantenerla simple para que se pueda usar en solitario y ofrecer una capacidad de testing más exhaustiva.

Sabemos que hay frameworks de "Property based testing", y seguramente te encuentres muy cómodo con ellos. Incluso es posible que te ofrezcan mejores abstracciones para ti. Pero ten en cuenta dos cosas:

 * Cuando creas código para Overlord aportas a una comunidad más allá de tu
   caso de testing.
 * Si ese framework es tan bueno ¡puedes integrarlo en Overlord! y aporta así
   a una comunidad más grande.

Intentamos que el esfuerzo de integración sea lo menor posible, para que ese valor aportado tenga un coste bajo.

Además de comprobar controles de seguridad, puedes usar las abstracciones de Overlord como herramientas. Creemos que el grafo dinámico puede adaptarse a una cantidad ingente de casos, especialmente en entornos de alta incertidumbre.


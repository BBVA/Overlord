---
title: "Rules"
date: 2021-10-26T12:14:39+02:00
draft: false
---

En el caso de la ejecución de una regla completa podemos ver el proceso de
abstracción completo desde la regla más general, llegando a un trozo de código
ejecutado.

Este es un ejemplo sencillo y puntual, se reserva para otras ejemplos una regla
fallida, una regla incompleta y otros escenario secundarios. Actualmente la PoC
de Overlord está centrada en la viabilidad por lo que se ciñe a los "Happy
Path".

También queda excluido del ejemplo el log, que debido a su extensión,
resultaría muy complejo de explicar detalladamente. En el log estarán los
resultados parciales de todas las ejecuciones, herramientas y caminos del grafo
dinámico descartados.

# Una regla completa

<style>
.gdoc-expand {
  margin-bottom: 0rem;
  margin-top: 0rem;
}

.blue{
  background-color: #9acdff;
}
</style>

{{< expand "Recomendación / Regla General" >}}

```
=> Rationale
The communications with user must be secure
=> Rule
All ($url) must have (g$tls_version) >= (ch$version)
```

En este punto una regla es una recomendación abstracta de alto nivel. Su
principal función es permitir una idea transversal a todos los proyectos de la
empresa, con esta se puede saber de forma global si las todas las
comunicaciones de la empresa cumplen con este control.

También permite la colaboración entre empresas a un nivel muy general.

En esta regla podemos ver varios conceptos generales como la palabra reservada
*All* que indica que esta regla debe ser cierta para todos los integrantes
siguientes. También podemos ver unas palabras entre parentesis; estos son los
parámetros que permiten instanciar esta regla para un entorno en concreto.

**No te fijes en los detalles funcionales de la regla, según esta progrese por
el sistema entenderás todo el proceso.**

## Actores

Los principales productores de este tipo de reglas serán:

 * Los grupos de trabajo dedicados a seguridad (Ej: OWASP)
 * Las agencias de seguridad y estandarización (Ej: NIST)
 * Los líderes de la industría a nivel de seguridad

Los principales interesados en estas reglas serán:

 * Los CISO y CTO de todas las empresas

{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
The communications with user must be secure
=> Rule
All ($url) must have (g$tls_version) >= (ch$version)
== rule id
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5
```

Esta regla tiene un Hash SHA256 para poder hacer referencia a ella en las
discusiones de la comunidad.

{{< /expand >}}
{{< expand "Instanciación General" "" "">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
```

En este punto de la regla estamos especificando que en nuestra empresa
consideramos segura una versión de TLS de 1.3 o superior. Este bloque permite
establecer los parámetros, que como vimos en el punto anterior tienen un
símbolo "$" y además están entre paréntesis.

Si te fijas este parámetro tiene las letras ch antes del símbolo del "$", esto
le dice a Overlord que ese parámetro solo se puede especificar en la fase de
"Check"; que es donde se establecen los criterios transversales a la
organización.

Este tipo de decisiones suelen ser las que están expresadas como controles de
seguridad. Estos controles idealmente son el resultado de un proceso de treat
modeling.

# Actores

Este tipo de decisiones se toman a nivel transversal en una organización.
Incluuyen no solo al CISO o CTO sino también a los técnicos de seguridad de más
alto nivel y experiencia.

{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779
```

Esta regla tiene un Hash SHA256 para poder hacer referencia a ella en los
sistemas de segimiento de riesgos transversales a la compañía.

{{< /expand >}}
{{< expand "Apuntado" "" "">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
```

Dentro del proceso de testing de riesgos es necesario agregar algunos
parámetros técnicos del proyecto concreto. En este caso estamos agregando un
parámetro que no esta explícitamente en la regla, de forma que el Overlord
priorice los caminos que usan esta información.

Esto es así porque dentro del proceso de comprobación de controles de seguridad
el como se hacen las cosas puede ser relevante en un determinado momento.
Debido a esto el grafo dinámico usará preferentemente los caminos, para obtener
los datos, usando los datos que ya posee primero.

# Actores

En este caso tanto el responsable ténico como su equipo están involucrados en
el proceso. Es necesario inspeccionar los resultados parciales de Overlord para
distinguir de aquellas posibilidades de obtención de información aquellas que
son utilizables en el proyecto actual.

Es de este proceso de revisión de donde se puede consultar que parámetros extra
se pueden proveer para que Overlord compruebe de una u otra forma. 

En este proceso de selección es posible que los responables de auditoría y
seguridad puedan aportar su perspectiva de que nivel de comprobación es
adecuado.

También pueden ver en el log como se ha llevado a cabo la comprobación sin
necesidad de convenirla previamente; de forma que se conserve el principio de
federación de Overlord.

{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
```

En este caso el identificador del proceso de targeting incluye el parámetro
técnico. Este cambiará en el caso de que el servidor dns asociado a este
proceso particular cambie. Overlord refuerza la transparencia siempre que es
posible.

En el caso de que no sea relevante simplemente podemos hacer referencia al
identificador del nivel superior y de esa forma no tendremos notificaciones
cada vez que cambie un detalle del entorno.

{{< /expand >}}
{{< expand "Ejecución" "" "">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
<= cli params
url = [one.com, two.com]
```

Para la ejecución podemos proveer a la regla del resto de parámetros
necesarios. En este caso los nombres de dominio implicados en el test del
control.

# Actores

En este caso los encargados de la ejecución del código de comprobación, ya sea
de forma manual o automatizada. Pueden incluir los siguientes perfiles: 

 * Desarrolladores interesados en mejorar la seguridad del código
 * Auditores que quieran comprobar el estado de un control
 * Expertos en seguridad que quieran comprobar el estados de un control

{{< /expand >}}
{{< expand "+Hash" "" "blue">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
<= cli params
url = [one.com, two.com]
== execution id
26b67743880b9bc6cbdbe66d51abd2d8846822bf7325d2e2bef08c611a6023d5
```

De nuevo el identificador contiene los parámetros de la ejecución de forma que
no se puedan cambiar las condiciones de ejecución sin que este cambie

{{< /expand >}}
{{< expand "Resultado" "" "">}}

```
=> Rationale
The communications with user must be secure  
=> Rule  
All ($url) must have (g$tls_version) >= (ch$version)  
== rule id  
03669a7dfefe38399e193ca8455b7bc92a62b83fdf1991922234acbcde9f04c5  
=> Check Block  
version = 1.3  
== check id  
b125b13bfb29e13c2e4b9ed4d491c1656ccf693c9959b1517908399e82ba8779  
=> Target Block  
dns_ip = 1.2.3.4  
== target id  
20038dd2edf3e8b8673d0e6c5e13b142bf81a88cde9a1c404f4ba3fb75ade4c1
<= cli params
url = [one.com, two.com]
== execution id
26b67743880b9bc6cbdbe66d51abd2d8846822bf7325d2e2bef08c611a6023d5
== Execution path
dns_mask_tool
ssl_tools
rule_check
== Final Rule
All [
  ("one.com" as $url) must have ("1.3" as g$tls_version) >= ("1.3" as ch$version),
  ("two.com" as $url) must have ("1.3" as g$tls_version) >= ("1.3" as ch$version),
  ]
== Result
All Rules OK
== Result id
c04a56d97f7e8e6d506e15c9d37ef45849fdb90a06d6b21d8d41b6516adfc67b
```

En el paso final, usando el grafo dinámico, seguimos los siguientes pasos: 

 * En primer lugar estamos ejecutando código que haciendo uso de
   la herramienta "DNSMask" resolverá las url usando un determinado servidor DNS
   (pasado como parámetro en el bloque "Target Block"). 
 * Luego se usa la herramienta de análisis de ssl y se extrae la propiedad
   buscada (tls_version).
 * Se expande la regla con esta información.
 * Se evalua

De esta forma obtenemos un "All Rules OK". Nuestro control esta funcionando.

Adicionalmente podemos ver que se ha generado un nuevo hash, que incluye datos
de todo lo anterior.

# Actores

El resultado de la ejecución es algo interesa a todos los implicados en el
proceso. En este caso el hash además incluye la hora del sistema para que
evitar que se reutilizen resultados anteriores. O para usar los resultados
cacheados si la caducidad de la regla esta establecida (no en este ejemplo).

Los bloques de código ejecutados son creados por los desarrolladores, si es que
no existen ya. Es una responsabilidad compartida el elegir que trozos de código
se cargan en Overlord.

{{< /expand >}}

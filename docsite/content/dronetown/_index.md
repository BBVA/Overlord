---
title: "Dronetown"
date: 2021-10-28T09:05:40+02:00
draft: false
---

Esta es una metáfora de soporte para ayudar a la compresión del sistema
Overlord, debido a que el nivel de abstracción del mismo puede provocar
incertidumbre en los detalles de implementación.

## Integrantes de la metáfora

Dentro de DroneTown tenemos varios actores que representarán partes de un
hipotético sistema real Overlord. Para entender la funcionalidad debemos obviar
esos sistemas dandolos por ciertos hasta que comprendamos el sistema en su
nucleo.

### DroneTown
La propia ciudad es un integrante de la metáfora, representa el sistema en su
conjunto sin preocuparse de la localización o división del mismo.

Todos los integrantes de DroneTown que están representados como edificio actuan
de forma similar. Cuando les llega un dron procesan secuencialmente su cerebro
hasta que encuentran una orden que entienden. En ese momento llevan a cabo su
tarea. Y si no encuentran una orden que entiendan envian al dron a un sitio por
defecto (que suele ser la torre de control por defecto)

### Dock
Los drones salen del Dock para cumplir sus recados. Estos drones solo tienen
dos ordenes en sus cerebros:

 * Obtener la Información obejetivo
 * Volver al Dock

En el caso de que un dron no tenga objetivo la torre atenderá a su segunda
instrucción y lo enviará de vuelta al Dock.   
Si por el contrario no tiene la instrucción de volver al Dock, el dron
intentantará obtener su objetivo, para después, acabar en el Junkyard lo
consiga o no.

### Junkyard
Es el sitio al que llegan los drones si no tienen un conjunto de ordenes
correctas en el cerebro o si no pueden cumplir su objetivo.

La inspección de estos drones puede revelar los requisitos no conocidos del
objetivo del dron, por lo que vuelven siempre con información aunque no sea la
que se ha pedido.

### Torre
La torre es un edificio estructural de DroneTown, tiene dos
responsabilidades:

 * Registrar todas las direcciones de DroneTown
 * Registrar todas las tiendas de DroneTown

Si la torre se encuentra una instrucción para enviar al dron a una dirección en
concreto, simplemente borra esta instrucción y envia al dron a esa dirección.
Este es un mecanismo común usado por los otros edificios cuando tienen procesos
pendientes.

Si la torre se encuentra una instrucción para obtener un determinado producto,
busca entre sus direcciones el sitio, y lo envía a esa dirección con una
instrucción de comprar.

Es posible que haya más de una tienda que proveea un determinado producto, por
lo que en ese caso la torre fletará tantos drones como tiendas haya para
asegurase de que se exploran todas las posibilidades. Cada uno de estos drones
es un clon completo del que entro en la torre, incluyendo sus pertenencias e
instrucciones.

También es posible que un dron tenga como objetivo un producto que no esta
registrado en la torre. En este caso la torre enviará al dron al Junkyard. La
torre es el único edificio que enviará un dron al Junkyard cuando no sepa como
proceder. El resto de edificios envian al dron a la torre si no saben que
hacer.

### Tienda
Todas las tiendas funcionan de una forma similar. Pueden llevar a cabo tres
operaciones básicamente:
 
 * En el caso de que les llege un dron, y que tengan todo lo necesario, le dan
   el producto que viene buscando y cambian su operación de compra por una de
   compra exitosa.
 * En el caso de que les llege un dron, pero que les falte lo que necesitan, le
   agregaran una instrucción que vuelva a la tienda y luego otra instrucción
   para que consiga lo que le falta.
 * En cualquier caso siempre envian al dron de vuelta a la torre.

## Operaciones especiales

Dentro de Drone Town, y sin romper las reglas de esta, hay algunas tiendas
curiosas que permiten controlar los recados más complejos. En el caso de que
una tienda, por ejemplo, pueda conseguir el producto a partir de varios
derivados, necesitamos decirle al dron que puede conseguir uno u otro. Similar
a un OR lógico. También es posible que una tienda requiera de más de un
producto para llevar a cabo la transacción, por lo que igual manera,
necesitamos decirle al dron que necesitamos un producto y el siguiente. Similar
a un AND.

Para esto hay dos tiendas que se preocupan de entender estas complejidades.

### ANDrés Store
Esta tienda se encarga de los encargos y requisitos que requieren de un AND
lógico.

Cuando llega un dron buscado su producto (el AND) exminan los integrantes de
AND. Por cada uno de los productos exmina si el dron ya los tiene o no. Y en el
caso de no tenerlos le añade una instrucción para comprar ese produto.

Siempre, antes de enviar el dron de vuelta a la torre elimina cualquier rastro
de su actividad para mantener las cosas simples.

### HectOR
En este caso la tienda funciona de forma similar. Inspecciona la mochila del
dron, Si ya tiene alguno de los elementos que necesita, borra el rastro de su
tienda y lo envia a la torre. 

Si no tiene ninguno de los elementos HectOr fletará tantos drones clonados como
elementos haya en la instrucción. De esta forma se asegura de que tendrán el
máximo número de posibilidades para completar el recado con éxito. Esta
operación es similar a la que lleva a cabo la torre cuando hay varias tiendas,
pero en el caso de que haya varios requisitos.

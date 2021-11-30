---
title: "Creador de Reglas"
date: 2021-11-05T12:34:46+01:00
draft: false
---

El proceso de creación de reglas para Overlord es bastante sencillo, pero requiere de entendimiento del DSL específico para reglas. El creador de Reglas es aquel usuario de Overlord que entiende los problemas de seguridad y quiere aportar valor diseñando tests genéricos sobre controles (qué hay que verificar).

En este caso, la comunidad será su principal objetivo, ya que la creación de reglas aporta valor dentro de la misma en forma de consolidación de cómo reconocer que un control está efectivamente funcionando tal como se diseñó, y que tanto el test genérico como la adaptación para un entorno específico pueda ser compartido y contrastado con el resto de la industria facilitando la discusión sobre algo tangible

El valor que aporta a la comunidad es su dominio específico de una plataforma tecnológica concreta y cómo concretar una regla genérica. En muchos de los ejemplos de esta documentación usamos una regla específica de entornos de red por la cual la versión de TLS de un servicio web debe estar por encima de 1.3. Este es un ejemplo perfecto de una regla creada por un experto en seguridad. Esta regla es una instanciación que pertenece a una recomendación general: "Los canales de comunicación con los clientes deben ser seguros".

La tarea del creador de reglas es vital en cuanto a que tiene el conocimiento técnico de seguridad necesario para materializar los tests en entornos genéricos. Entendiendo como entorno general por ejemplo la nube pública y como entorno específico un proyecto en concreto de la nube pública; que sería responsabilidad de los responsables técnicos de cada proyecto concreto. Si bien el Security Risk Manager debe establecer en qué casos deben existir controles, el creador de reglas es aliado necesario que dice qué reglas son necesarias para verificar la existencia y eficacia de los controles en los diferentes entornos.

En muchas ocasiones el creador de reglas no tiene por qué entender los detalles técnicos hasta el último detalle. Por lo que el DSL que define las reglas no debe describir el proceso a nivel de qué pasos son necesarios; ya que sería muy difícil de reutilizar para diferentes entornos, como le pasa hoy en día a los test de controles escritos directamente en código. En lugar de establecer el cómo se comprueban las cosas es más importante saber QUÉ debe comprobarse.

En otros casos el conocimiento técnico es vital para la correcta resolución y la regla debe especificar qué programas y procedimientos se deben llevar a cabo. En este caso las reglas proveen la capacidad de agregar estos detalles técnicos desde un alto nivel de abstracción.

## Parametrizando

Para que las reglas sean realmente reutilizables y útiles para la comunidad en ocasiones es necesario que se puedan ajustar ya en un alto nivel de abstracción. Para esta labor las reglas tienen los parámetros de Objetivo. Estos permiten al creador de reglas que la empresa adapte la regla de seguridad a un determinado nivel de riesgo en función de su análisis de riesgos.

Poniendo un ejemplo sencillo podemos hablar de una regla con el siguiente texto libre (la sintaxis no está cerrada todavía).

```
Every https service must have TLS version at least ($tls_min)
```

En este caso el parámetro "$tls_min" se puede establecer a varios valores que dependen del negocio o caso de uso. Por ejemplo, si tenemos altos requisitos de seguridad podría ser la última versión.

Además si este detalle es estratégico se puede forzar a que sea definido, y que no pueda ser reescrito después, solamente en los primeros pasos de abstracción. Para esto los parámetros técnicos tienen un prefijo que se puede poner antes del símbolo del "$"; esté es tg (de target).

La anterior regla quedaría así:

```
Every https service must have TLS version at least (tg$tls_min)
```

Para obtener mayor información al respecto de las reglas, estas estarán descritas en su propia documentación.


---
title: "Creador de Reglas"
date: 2021-11-05T12:34:46+01:00
draft: true
---

El proceso de creación de reglas para Overlord es bastante sencillo, pero
requiere de entendimiento del DSL específico para reglas. En este caso el
creador de Reglas es aquel usuario de Overlord que entiende los problemas de
seguridad y quiere aportar valor mediante controles genéricos.

En este caso la comunidad será su principal objetivo, ya que la creación de
reglas aporta valor dentro de la misma en forma de reconocimiento e intercambio
de información.

El valor que aporta a la comunidad es su dominio específico de una plataforma
tencológica concreta y a que regla general aplica. En muchos de los ejemplos de
esta documentación usamos una regla específica de entornos de red por la cual
la versión de TLS de un servicio web debe estar por encima de 1.3. Este es un
ejemplo perfecto de una regla creada por un experto en seguridad. Esta regla es
una instanciación que pertenece a una recomendación general: "Los canales de
comunicación con los clientes deben ser seguros".

La tarea del creador de reglas es vital en cuanto a que tiene el conocimiento
técnico de seguridad necesario para materializar los controles en entornos
generales. Entiendo entorno general como la nube pública y el entorno
específico el de un proyecto en concreto de la nube pública; que sería
responsabilidad de los responsables técnicos de cada proyecto concreto. Si bien
el Security Risk Manager debe establecer en que casos hay controles el creador
de reglas es aliado necesario que dice que controles son necesarios en los
diferentes entornos generales.

En muchas ocasiones el creador de reglas no tiene porque entender los detalles
técnicos hasta el máximo detalle. Aunque suele pasar no es necesario para
desempeñar su tarea de seguridad. Por lo que el DSL que define las reglas no
debe describir el proceso a nivel de que pasos son necesarios; ya que sería muy
dificil de reutilizar para diferentes entornos, como le pasa hoy en día a los
test de controles escritos directamente en código. En lugar de establecer el
como se comprueban las cosas es más importante saber QUE debe comprobarse.
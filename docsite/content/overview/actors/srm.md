---
title: "Security Risk Manager"
date: 2021-11-03T11:40:50+01:00
draft: true
---

El Security Risk Manager se centra en el modelado de amenazas, los riesgos y
los controles para mitigarlos. Sin embargo suele perder la última milla de
verificación de que los controles que especificó están o no operativos, sobre
todo en entornos de de Despliegue Contínuo e Infrastructure as Code, por lo que
Overlord debería ser una de las herramientas que más valor le puede aportar.

El tener visión de todos los controles de forma transversal a la compañía es
una tarea abordada por software de gestión de riesgos (ASR&TM: Application
Security Requirements and Threat Modeling o  GRC/IRM). En algunos de los
programas se pueden llegar a comprobar, mediante código integrado en la
herramienta, cuál es el estado de algunos controles típicos.

Pero por ahora ninguna herramienta ofrece la integración y facilidad para poder
comprobar de la forma más automatizada posible, todos los controles de la
organización. La posibilidad de integrar de forma federada todos los sistemas y
herramientas en un solo proceso de ejecución real de la comprobación puede ser
un paso decisivo para la gestión de riesgos contínua.

Las tendencias de desarrollo actuales han planteado un reto tecnológico real a
los Security Risk Manager. La automatización y la proliferación tecnológica han
llevado al límite la demanda que se exige en cuanto a controles de seguridad y
su implantación.

Pese a que Overlord no es un software de gestión de riesgos sus posibilidades
de integración con este son amplias y extensibles. Abriendo la puerta a una
simbiosis que puede cambiar una tendencia típica de este tipo de trabajos; de
comprobar solo lo más esencial en un momento dado, a poder tener una verdadera
visión de conjunto contínua y actualizada en cada cambio.

Overlord le ofrece la posibilidad de gestionar tanto un catálogo genérico de
reglas asociables a cualquier realidad técnica, como una base de código en la
que puede establecer incluso detalles de la comprobación relevantes a nivel de
riesgos.

## Reglas

Las Reglas de Overlord están pensadas para poder ser escritas fácilmente por
cualquier experto en seguridad. No es necesario aprender a programar en ningún
lenguaje, ni tampoco aprender una larga lista de órdenes para establecer como
deben de hacerse las cosas en cada entorno. Definen el qué pero no el cómo,
aunque el cómo estará disponible para su aprobación (eficacia del test) y
auditoría (integridad de lo verificado).

Una regla puede concretarse y adaptarse a cualquier entorno, debido a que
establece el objetivo de comprobación del control de seguridad de forma
abstracta. De esta forma se puede compartir y adaptar.

Por otro lado, tanto el código auxiliar que permita a Overlord obtener los
parámetros intermedios con los que configurar el código del test atómico
(código atómico), como estos mismos tests atómicos ya desarrollados en forma de
productos, herramientas open source o pequeños desarrollos internos pueden ser
integrados en Overlord de forma independiente y desacoplada a las reglas.

Un sistema Overlord con el suficiente código Atómico integrado puede ser capaz
de ejecutar un gran conjunto de nuevas reglas sin necesidad de agregar más
código. Esto es así porque el sistema Overlord intentará buscar el camino por
el código existente hasta el objetivo de la comprobación.

Y si no puede encontrar el camino, puede dar algunas alternativas para que la
cantidad de código a añadir al sistema sea siempre la mínima necesaria, creando
un backlog de elementos a añadir para maximizar la cobertura de comprobaciones
que puede ser priorizado según el riesgo de la no comprobación de los controles
asociados. 

## Código Atómico

Como experto en Riesgos no es necesario saber programar, pero si que es
necesario estar familiarizado con las diferencias entre distintos procesos de
comprobación de controles.

No es lo mismo a nivel de aseguramiento consultar un fichero de configuración
que, por ejemplo, usar una herramienta de análisis de red para comprobar en
sistema real que se cumple una condición o lanzar dicha herramienta desde la
propia red o desde otra.

El código Atómico lleva siempre asociado una descripción de cómo se está
obteniendo la información. Mediante los parámetros de comprobación (Check
Block) se puede forzar a Overlord para que use una determinada pieza de código
en una ejecución; asegurando así que la forma en la que se hacen las cosas
cumple con los requisitos de seguridad en un proyecto en concreto.


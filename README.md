# Balanz - Desafío de Backend en Go

¡Bienvenido al Desafío de Backend en Go! Este desafío está diseñado para evaluar tus habilidades en programación y desarrollo de backend utilizando el lenguaje de programacion Go (Golang).
En este desafío, se espera que diseñes y desarrolles dos endpoints que interactúen con una base de datos PostgreSQL (o SQLite) y emulen la persistencia de archivos en un servicio S3 de AWS.

## Pautas Generales

- Podes utilizar el framework de tu elección (por ejemplo, Gin Gonic, Fiber) para implementar los endpoints de manera eficiente.
- Tenes la libertad de elegir el ORM que prefieras para interactuar con la base de datos PostgreSQL (o SQLite).
- Asegurate de manejar los errores de manera adecuada en tu código.

## Endpoint 1: Parseo y Persistencia de Datos

Diseña un endpoint que consulte data de la siguiente URL:

`https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd`

El endpoint deberá parsear estos datos y persistir en una tabla de una base de datos PostgreSQL (o SQLite) los siguientes campos:

- Symbol
- Name
- Current Price
- Last Updated

## Endpoint 2: Persistencia de Archivos en S3

Diseña otro endpoint que reciba una serie de archivos y emule la carga de ellos en un servicio S3 de AWS. El endpoint debe persistir las keys asociadas a los archivos almacenados en S3 en una tabla junto con un ID proporcionado como parámetro del endpoint.

```
POST /endpoint2?id=b275202c-32ec-4652-998e-d6e97d104142

Headers: Content-Type: multipart/form-data
```

### S3 de AWS

La idea es solamente emular la carga del archivo en S3. Para ello se podría tener una función que retorne la key como un string con el formato `files/*uuid.UUID*-*filename.xxx*`.

---

```
En el caso que no tengas instalado PostgreSQL en tu máquina, en este repositorio hay un docker-compose listo para levantar una instancia del mismo.
```

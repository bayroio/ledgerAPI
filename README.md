# ledgerAPI

#  Installation
* Browse the repository's root
* Build the images 
    - `docker-compose build`
* Start containers 
    - `docker-compose up -d`

After starting containers you can test the Api at:
```url
http://localhost:85/api/

## Listar todos los contenedores:
docker container ls --all

## Listar todos las imágenes:
docker image ls

## Detener y eliminar el contenedor Docker usando su nombre:
docker rm --force linux_tweet_app

## Login en Docker Hub:
docker login

## Publicar la versión de una imagen en Docker Hub:
docker image push $DOCKERID/linux_tweet_app:1.0
docker image push $DOCKERID/linux_tweet_app:2.0

## Para ver la imagen publicada en Docker Hub
https://hub.docker.com/r/<your docker id>/
![spotify](https://e7.pngegg.com/pngimages/158/639/png-clipart-spotify-streaming-media-logo-playlist-spotify-app-icon-logo-music-download.png)
![golang](https://www.clipartmax.com/png/middle/288-2881446_image-result-for-golang-go-programming-language-logo.png)

#Spotify API with Golang

POC utilizando la API de Spotify dentro de Golang

---

## Herramientas

- Golang
- Spotify API
- Gingonic
- Ginkgo
- Gomega

---

## Usabilidad

Dentro del repositorio se encuentra un Makefile que permite correr comandos:
- El primero que se debe correr es el dependencies que permitirá descargar las bibliotecas de Gomega y Ginkgo
- Luego se debe correr el comando run-locally para poder inicializar el proyecto

Una vez inicializado el proyecto, se podrá ver en  algún browser el swagger que permite el acceso a los dos endpoints. La URL será http://localhost:8080/swagger/index.html 

Se visualizarán dos endpoints, el setToken que se utilzia para settear el clienteId, clientSecret y Token de Spotify. Y también el getDiscography que retornará la discografía de la banda que se ingrese

This example will create an IAM user and allow read access to all objects in the S3 bucket `examplebucket`

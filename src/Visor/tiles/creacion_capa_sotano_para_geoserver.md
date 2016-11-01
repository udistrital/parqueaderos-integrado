#Pasos para replicar la creación de una capa raster para geoserver.
* Abrir sotano#.svg con inkscape
* Exportar la imagen a .png con una densidad de píxeles de 50ppp
* Abrir la imagen con qgis y georreferenciarla, existe un plugin llamado Free Hand Raster Georeferencer que permite hacerlo fácilmente.
* Click derecho sobre la capa y guardarla como GTiff con modo de salida Datos Crudos EPSG:4326, WGS 84
* En geoserver crear un workspace llamado parqueaderos, luego crear un store llamado sotano#_tematico
* Luego publicar la nueva capa con EPSG:4326


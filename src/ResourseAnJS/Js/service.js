angular.module("MyFirstApp")
.factory("UrlResource",function($resource){
return $resource("http://192.168.33.10:8080/v1/tipo_propietario/:id",{id: "@id"},{update: {method: "PUT"}})
})

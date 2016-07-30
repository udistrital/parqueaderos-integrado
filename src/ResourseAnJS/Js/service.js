angular.module("MyFirstApp")
.factory("UrlResource",function($resource){
return $resource("http://localhost:8080/v1/tipo_propietario/:id",{id: "@id"},{update: {method: "PUT"}})
})

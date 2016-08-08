/**
 * Created by Shivali on 6/30/15.
 */

angular.module('Propietario')
    .factory('PropietarioServices',["$http","$q",function ($http,$q) {
        return {
            retrievePropietarios:function(){

                var defer=$q.defer();
                $http.get("propietario")
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});
                return defer.promise;
            },
            savePropietario:function(PropietarioObj){

                var defer=$q.defer();
                $http.post("propietario",PropietarioObj)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});
                return defer.promise;
            },
            updatePropietario:function(PropietarioObj){

                var defer=$q.defer();
                $http.put("propietario/"+PropietarioObj.id,PropietarioObj)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});

                return defer.promise;
            },
            deletePropietario:function(id){

                var defer=$q.defer();
                $http.delete("propietario/"+id)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});

                return defer.promise;
            },
            retrievePropietarioById:function(id){

                var defer=$q.defer();
                $http.get("propietario/"+id)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});
                return defer.promise;
            }
        }
    }]);
/**
 * Created by Shivali on 6/30/15.
 */

angular.module('Vehiculo')
    .factory('VehiculoServices',["$http","$q",function ($http,$q) {
        return {
            retrieveVehiculos:function(){

                var defer=$q.defer();
                $http.get("vehiculo")
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});
                return defer.promise;
            },
            saveVehiculo:function(VehiculoObj){

                var defer=$q.defer();
                $http.post("vehiculo",VehiculoObj)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});
                return defer.promise;
            },
            updateVehiculo:function(VehiculoObj){

                var defer=$q.defer();
                $http.put("vehiculo/"+VehiculoObj.id,VehiculoObj)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});

                return defer.promise;
            },
            deleteVehiculo:function(id){

                var defer=$q.defer();
                $http.delete("vehiculo/"+id)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});

                return defer.promise;
            },
            retrieveVehiculoById:function(id){

                var defer=$q.defer();
                $http.get("vehiculo/"+id)
                    .success(function(result){defer.resolve(result)})
                    .error(function(error){defer.reject(error)});
                return defer.promise;
            }
        }
    }]);
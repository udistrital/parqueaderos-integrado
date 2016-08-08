/**
 * Created by Shivali on 6/30/15.
 */

angular.module('Vehiculo')
    .controller('VehiculoCtrl', ["$scope", "$rootScope", "VehiculoServices", function ($scope, $rootScope,VehiculoServices) {

        $scope.currentVehiculo= {};
        $scope.vehiculoFields=[{"type":"input","key":"placa","templateOptions":{"label":"Placa","required":true,"type":"text"}},{"type":"input","key":"id_nfc","templateOptions":{"label":"ID NFC","required":true,"type":"text"}},{"type":"input","key":"id_propietario","templateOptions":{"label":"Propietario","required":true,"type":"text"}}];

        // retrieve Vehiculo to server
        $scope.retrieveVehiculos = (function(){
            VehiculoServices.retrieveVehiculos()
                .then(function(result){
                    $rootScope.vehiculos = result;
                    },function(error){
                        alert(error);
                    })
            })();


        $scope.openAddVehiculo=function(){
            $scope.currentVehiculo= {};
            $('#addVehiculo.modal').modal('show');
        }

        // save Vehiculo to server
        $scope.saveVehiculo = function(){
            VehiculoServices.saveVehiculo($scope.currentVehiculo)
            .then(function(result){
                    $rootScope.vehiculos.push(result);
                },function(error){
                    alert(error);
                })
        }

        $scope.openUpdateVehiculo=function(data){
            $scope.currentVehiculo=angular.copy(data);
            $('#updateVehiculo.modal').modal('show');
        }

        //update data to server
        $scope.updateVehiculo = function(){
            VehiculoServices.updateVehiculo($scope.currentVehiculo)
            .then(function(result){
                for(var key in $rootScope.vehiculos){
                    if(result.id==$rootScope.vehiculos[key].id)
                        $rootScope.vehiculos[key] = result;
                }
                },function(error){
                    alert(error);
                })
        }

        $scope.openDeleteVehiculo=function(data){
            $scope.currentVehiculo=angular.copy(data);
            $('#deleteVehiculo.modal').modal('show');
        }

        //delete data to server
        $scope.deleteVehiculo = function(){
                VehiculoServices.deleteVehiculo($scope.currentVehiculo.id)
                .then(function(result){
                    for(var key in $rootScope.vehiculos){
                        if($scope.currentVehiculo.id==$rootScope.vehiculos[key].id){
                            $rootScope.vehiculos.splice(key,1);
                            break;
                        }
                }
                },function(error){
                    alert(error);
                })
        }

        $scope.emptyVehiculo = function(){
            $scope.currentVehiculo={};
        }
    }]);
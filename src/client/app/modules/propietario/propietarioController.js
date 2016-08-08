/**
 * Created by Shivali on 6/30/15.
 */

angular.module('Propietario')
    .controller('PropietarioCtrl', ["$scope", "$rootScope", "PropietarioServices", function ($scope, $rootScope,PropietarioServices) {

        $scope.currentPropietario= {};
        $scope.propietarioFields=[{"type":"input","key":"documento","templateOptions":{"label":"Documento","required":true,"type":"text"}},{"type":"input","key":"primer_nombre","templateOptions":{"label":"Primer Nombre","required":true,"type":"text"}},{"type":"input","key":"otros_nombres","templateOptions":{"label":"Otros Nombres","required":true,"type":"text"}},{"type":"input","key":"primer_apellido","templateOptions":{"label":"Primer Apellido","required":true,"type":"text"}},{"type":"input","key":"segundo_apellido","templateOptions":{"label":"Segundo Apellido","required":false,"type":"text"}},{"type":"select","key":"id_tipo_propietario","templateOptions":{"label":"Tipo Propietario","required":true,"options":[{"name":"Funcionario","value":"1"},{"name":"Contratista","value":"2"},{"name":"Docente Planta","value":"3"},{"name":"Docente Cátedra","value":"4"},{"name":"Estudiante","value":"5"}]}}];

        // retrieve Propietario to server
        $scope.retrievePropietarios = (function(){
            PropietarioServices.retrievePropietarios()
                .then(function(result){
                    $rootScope.propietarios = result;
                    },function(error){
                        alert(error);
                    })
            })();


        $scope.openAddPropietario=function(){
            $scope.currentPropietario= {};
            $('#addPropietario.modal').modal('show');
        }

        // save Propietario to server
        $scope.savePropietario = function(){
            PropietarioServices.savePropietario($scope.currentPropietario)
            .then(function(result){
                    $rootScope.propietarios.push(result);
                },function(error){
                    alert(error);
                })
        }

        $scope.openUpdatePropietario=function(data){
            $scope.currentPropietario=angular.copy(data);
            $('#updatePropietario.modal').modal('show');
        }

        //update data to server
        $scope.updatePropietario = function(){
            PropietarioServices.updatePropietario($scope.currentPropietario)
            .then(function(result){
                for(var key in $rootScope.propietarios){
                    if(result.id==$rootScope.propietarios[key].id)
                        $rootScope.propietarios[key] = result;
                }
                },function(error){
                    alert(error);
                })
        }

        $scope.openDeletePropietario=function(data){
            $scope.currentPropietario=angular.copy(data);
            $('#deletePropietario.modal').modal('show');
        }

        //delete data to server
        $scope.deletePropietario = function(){
                PropietarioServices.deletePropietario($scope.currentPropietario.id)
                .then(function(result){
                    for(var key in $rootScope.propietarios){
                        if($scope.currentPropietario.id==$rootScope.propietarios[key].id){
                            $rootScope.propietarios.splice(key,1);
                            break;
                        }
                }
                },function(error){
                    alert(error);
                })
        }

        $scope.emptyPropietario = function(){
            $scope.currentPropietario={};
        }
    }]);
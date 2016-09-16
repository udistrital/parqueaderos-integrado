'use strict';

angular.module('myapp')
  .controller('VehiculoController', ['$scope', '$uibModal', 'resolvedVehiculo', 'Vehiculo',
    function($scope, $modal, resolvedVehiculo, Vehiculo) {

      $scope.Vehiculos = resolvedVehiculo;

      $scope.create = function() {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function(id) {
        $scope.Vehiculo = Vehiculo.get({
          id: id
        });
        $scope.open(id);
      };

      $scope.delete = function(id) {
        Vehiculo.delete({
            id: id
          },
          function() {
            $scope.Vehiculos = Vehiculo.query();
          });
      };

      $scope.save = function(id) {
        var VehiculoSave = {
          "Placa": $scope.Vehiculo.Placa,
          "IdNfc": $scope.Vehiculo.IdNfc,
          "IdPropietario": {
            "Id": $scope.Vehiculo.IdPropietario.Id.Id,
            "Documento": "",
            "PrimerNombre": "",
            "OtrosNombres": "",
            "PrimerApellido": "",
            "Segundo Apellido": "",
            "IdTipoPropietario": {
              "Id": null,
              "Tipo": "",
              "Descripcion": ""
            }
          }

        };
        if (id) {
          Vehiculo.update({
              id: id
            }, VehiculoSave,
            function() {
              $scope.Vehiculos = Vehiculo.query();
              $scope.clear();
            });
        } else {
          Vehiculo.save(VehiculoSave,
            function() {
              $scope.Vehiculos = Vehiculo.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function() {
        $scope.Vehiculo = {

          "Placa": "",

          "IdNfc": "",

          "IdPropietario": "",

          "Id": ""
        };
      };

      $scope.open = function(id) {
        var VehiculoSave = $modal.open({
          templateUrl: 'Vehiculo-save.html',
          controller: 'VehiculoSaveController',
          resolve: {
            Vehiculo: function() {
              return $scope.Vehiculo;
            }
          }
        });

        VehiculoSave.result.then(function(entity) {
          $scope.Vehiculo = entity;
          $scope.save(id);
        });
      };
    }
  ])
  .controller('VehiculoSaveController', ['$scope', '$http', '$uibModalInstance', 'Vehiculo', 'CONFIG',
    function($scope, $http, $modalInstance, Vehiculo, CONFIG) {
      $scope.Vehiculo = Vehiculo;
      
      var f = [{}];
      $http.get(CONFIG.WS_URL + '/propietario')
        .success(function(data) {
          data.forEach(function(entry, index) {
            f[index] = {
              Id: entry.Id,
              Nombre: entry.PrimerNombre + ' ' + entry.PrimerApellido
            };
          });
          console.log(f);
          $scope.VehiculosIds = f;
        })
        .error(function(err) {});

      $scope.ok = function() {
        $modalInstance.close($scope.Vehiculo);
      };

      $scope.cancel = function() {
        $modalInstance.dismiss('cancel');
      };
    }
  ]);

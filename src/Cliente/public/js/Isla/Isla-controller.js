'use strict';

angular.module('myapp')
  .controller('IslaController', ['$scope', '$uibModal', 'resolvedIsla', 'Isla',
    function($scope, $modal, resolvedIsla, Isla) {

      $scope.Islas = resolvedIsla;

      $scope.create = function() {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function(id) {
        $scope.Isla = Isla.get({
          id: id
        });
        $scope.open(id);
      };

      $scope.delete = function(id) {
        Isla.delete({
            id: id
          },
          function() {
            $scope.Islas = Isla.query();
          });
      };

      $scope.save = function(id) {
        var IslaSave = {
          "Ocupado": $scope.Isla.Ocupado,
          "IdVehiculo": {
            "Id": $scope.Isla.IdVehiculo.Id.Id,
            "Placa": "",
            "IdNfc": null,
            "IdPropietario": {
              "Id": null,
              "Documento": "",
              "PrimerNombre": "",
              "OtrosNombres": "",
              "PrimerApellido": "",
              "SegundoApellido": "",
              "IdTipoPropietario": {
                "Id": null,
                "Tipo": "",
                "Descripcion": "",
              }
            }
          },
          "Geometria": $scope.Isla.Geometria,
          "IdGrupoIsla": {
            "Id": $scope.Isla.IdGrupoIsla.Id,
            "Geometria": "",
            "IdGrupoPadre": null
          }
        };
        if (id) {
          Isla.update({
              id: id
            }, IslaSave,
            function() {
              $scope.Islas = Isla.query();
              $scope.clear();
            });
        } else {
          Isla.save(IslaSave,
            function() {
              $scope.Islas = Isla.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function() {
        $scope.Isla = {

          "Ocupado": false,

          "IdVehiculo": "",

          "Geometria": "",

          "HoraEntrada": "",

          "HoraSalida": "",

          "Id": ""
        };
      };

      $scope.open = function(id) {
        var IslaSave = $modal.open({
          templateUrl: 'Isla-save.html',
          controller: 'IslaSaveController',
          resolve: {
            Isla: function() {
              return $scope.Isla;
            }
          }
        });

        IslaSave.result.then(function(entity) {
          $scope.Isla = entity;
          $scope.save(id);
        });
      };
    }
  ])
  .controller('IslaSaveController', ['$scope', '$http', '$uibModalInstance', 'Isla',
    function($scope, $http, $modalInstance, Isla) {
      $scope.Isla = Isla;
      var f = [];
      $http.get("v1/vehiculo")
        .success(function(data) {
          data.forEach(function(entry, index) {
            f[index] ={ 
		    Id: entry.Id,
		    IdNfc: entry.IdNfc
	    };
          });
          console.log(f);
	  $scope.IslasVehIds = f;
        })
        .error(function(err) {});

      var f1 = [];
      $http.get("v1/grupo_isla")
        .success(function(data) {
          data.forEach(function(entry, index) {
            f1[index] = entry.Id;
          });
          console.log(f1);
	  $scope.IslasGrIds = f1;
        })
        .error(function(err) {});

      $scope.ok = function() {
        $modalInstance.close($scope.Isla);
      };

      $scope.cancel = function() {
        $modalInstance.dismiss('cancel');
      };
    }
  ]);

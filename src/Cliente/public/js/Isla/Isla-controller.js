'use strict';

angular.module('myapp')
  .controller('IslaController', ['$scope', '$modal', 'resolvedIsla', 'Isla',
    function ($scope, $modal, resolvedIsla, Isla) {
     
      $scope.Islas = resolvedIsla;
     
      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.Isla = Isla.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Isla.delete({id: id},
          function () {
            $scope.Islas = Isla.query();
          });
      };

      $scope.save = function (id) {
	      var IslaSave = {
	"Ocupado":	$scope.Isla.Ocupado,
	"IdVehiculo": {
		"Id":		$scope.Isla.IdVehiculo.Id,
		"Placa":	"",
		"IdNfc":	null,
		"IdPropietario": {
			"Id":			null,
			"Documento":		"",
			"PrimerNombre":		"",
			"OtrosNombres":		"",
			"PrimerApellido":	"",
			"SegundoApellido":	"",
			"IdTipoPropietario": {
				"Id":		null,
				"Tipo":		"",
				"Descripcion":	"",
			}
		}
	},
	"Geometria":	$scope.Isla.Geometria,
	"HoraEntrada":	$scope.Isla.HoraEntrada,
	"HoraSalida":	$scope.Isla.HoraSalida,
	"IdGrupoIsla": {
		"Id":		$scope.Isla.IdGrupoIsla.Id,
		"Geometria":	"",
		"IdGrupoPadre":	null
	}
};
	if (id) {
          Isla.update({id: id}, IslaSave,
            function () {
              $scope.Islas = Isla.query();
              $scope.clear();
            });
        } else {
          Isla.save(IslaSave,
            function () {
              $scope.Islas = Isla.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.Isla = {
          
          "Ocupado": false,
          
          "IdVehiculo": "",
          
          "Geometria": "",
          
          "HoraEntrada": "",
          
          "HoraSalida": "",
          
          "Id": ""
        };
      };

      $scope.open = function (id) {
        var IslaSave = $modal.open({
          templateUrl: 'Isla-save.html',
          controller: 'IslaSaveController',
          resolve: {
            Isla: function () {
              return $scope.Isla;
            }
          }
        });

        IslaSave.result.then(function (entity) {
          $scope.Isla = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('IslaSaveController', ['$scope', '$modalInstance', 'Isla',
    function ($scope, $modalInstance, Isla) {
      $scope.Isla = Isla;

      $scope.ok = function () {
        $modalInstance.close($scope.Isla);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);

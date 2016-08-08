'use strict';

angular.module('myapp')
  .controller('PropietarioController', ['$scope', '$modal', 'resolvedPropietario', 'Propietario',
    function ($scope, $modal, resolvedPropietario, Propietario) {

      $scope.Propietarios = resolvedPropietario;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.Propietario = Propietario.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Propietario.delete({id: id},
          function () {
            $scope.Propietarios = Propietario.query();
          });
      };

      $scope.save = function (id) {
	var PropietarioSave = {
		"Documento":	       $scope.Propietario.Documento,
		"PrimerNombre":        $scope.Propietario.PrimerNombre,
		"OtrosNombres":        $scope.Propietario.OtrosNombres,
		"PrimerApellido":      $scope.Propietario.PrimerApellido,
		"SegundoApellido":     $scope.Propietario.SegundoApellido,
		"IdTipoPropietario":  {
				"Id": 		$scope.Propietario.IdTipoPropietario.Id,
				"Tipo": 	"",
				"Descripcion" : ""
		}
	};
        if (id) {
          Propietario.update({id: id}, PropietarioSave,
            function () {
              $scope.Propietarios = Propietario.query();
              $scope.clear();
            });
        } else {
          Propietario.save(PropietarioSave,
            function () {
              $scope.Propietarios = Propietario.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.Propietario = {
          
          "Documento": "",
          
          "PrimerNombre": "",
          
          "OtrosNombres": "",
          
          "PrimerApellido": "",
          
          "SegundoApellido": "",
          
          "IdTipoPropietario": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var PropietarioSave = $modal.open({
          templateUrl: 'Propietario-save.html',
          controller: 'PropietarioSaveController',
          resolve: {
            Propietario: function () {
              return $scope.Propietario;
            }
          }
        });

        PropietarioSave.result.then(function (entity) {
          $scope.Propietario = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('PropietarioSaveController', ['$scope', '$modalInstance', 'Propietario',
    function ($scope, $modalInstance, Propietario) {
      $scope.Propietario = Propietario;

      

      $scope.ok = function () {
        $modalInstance.close($scope.Propietario);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);

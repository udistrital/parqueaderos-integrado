'use strict';

angular.module('myapp')
  .controller('Tipo_PropietarioController', ['$scope', '$uibModal', 'resolvedTipo_Propietario', 'Tipo_Propietario',
    function($scope, $modal, resolvedTipo_Propietario, Tipo_Propietario) {

      $scope.Tipo_propietarios = resolvedTipo_Propietario;

      $scope.create = function() {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function(id) {
        $scope.Tipo_Propietario = Tipo_Propietario.get({
          id: id
        });
        $scope.open(id);
      };

      $scope.delete = function(id) {
        Tipo_Propietario.delete({
            id: id
          },
          function() {
            $scope.Tipo_propietarios = Tipo_Propietario.query();
          });
      };

      $scope.save = function(id) {
        var Tipo_PropietarioSave = {
          "Tipo": $scope.Tipo_Propietario.Tipo,
          "Descripcion": $scope.Tipo_Propietario.Descripcion
        };
        if (id) {
          Tipo_Propietario.update({
              id: id
            }, Tipo_PropietarioSave,
            function() {
              $scope.Tipo_propietarios = Tipo_Propietario.query();
              $scope.clear();
            });
        } else {
          Tipo_Propietario.save(Tipo_PropietarioSave,
            function() {
              $scope.Tipo_propietarios = Tipo_Propietario.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function() {
        $scope.Tipo_Propietario = {

          "Tipo": "",

          "Descripcion": "",

          "id": ""
        };
      };

      $scope.open = function(id) {
        var Tipo_PropietarioSave = $modal.open({
          templateUrl: 'Tipo_Propietario-save.html',
          controller: 'Tipo_PropietarioSaveController',
          resolve: {
            Tipo_Propietario: function() {
              return $scope.Tipo_Propietario;
            }
          }
        });

        Tipo_PropietarioSave.result.then(function(entity) {
          $scope.Tipo_Propietario = entity;
          $scope.save(id);
        });
      };
    }
  ])
  .controller('Tipo_PropietarioSaveController', ['$scope', '$uibModalInstance', 'Tipo_Propietario',
    function($scope, $modalInstance, Tipo_Propietario) {
      $scope.Tipo_Propietario = Tipo_Propietario;



      $scope.ok = function() {
        $modalInstance.close($scope.Tipo_Propietario);
      };

      $scope.cancel = function() {
        $modalInstance.dismiss('cancel');
      };
    }
  ]);

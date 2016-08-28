'use strict';

angular.module('myapp')
  .controller('Grupo_IslaController', ['$scope', '$modal', 'resolvedGrupo_Isla', 'Grupo_Isla',
    function($scope, $modal, resolvedGrupo_Isla, Grupo_Isla) {

      $scope.Grupo_islas = resolvedGrupo_Isla;

      $scope.create = function() {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function(id) {
        $scope.Grupo_Isla = Grupo_Isla.get({
          id: id
        });
        $scope.open(id);
      };

      $scope.delete = function(id) {
        Grupo_Isla.delete({
            id: id
          },
          function() {
            $scope.Grupo_islas = Grupo_Isla.query();
          });
      };

      $scope.save = function(id) {
        var Grupo_IslaSave = {
          "Geometria": $scope.Grupo_Isla.Geometria,
          "IdGrupoPadre": $scope.Grupo_Isla.IdGrupoPadre
        };
        if (id) {
          Grupo_Isla.update({
              id: id
            }, Grupo_IslaSave,
            function() {
              $scope.Grupo_islas = Grupo_Isla.query();
              $scope.clear();
            });
        } else {
          Grupo_Isla.save(Grupo_IslaSave,
            function() {
              $scope.Grupo_islas = Grupo_Isla.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function() {
        $scope.Grupo_Isla = {

          "Geometria": "",

          "IdGrupoPadre": "",

          "Id": ""
        };
      };

      $scope.open = function(id) {
        var Grupo_IslaSave = $modal.open({
          templateUrl: 'Grupo_Isla-save.html',
          controller: 'Grupo_IslaSaveController',
          resolve: {
            Grupo_Isla: function() {
              return $scope.Grupo_Isla;
            }
          }
        });

        Grupo_IslaSave.result.then(function(entity) {
          $scope.Grupo_Isla = entity;
          $scope.save(id);
        });
      };
    }
  ])
  .controller('Grupo_IslaSaveController', ['$scope', '$modalInstance', 'Grupo_Isla',
    function($scope, $modalInstance, Grupo_Isla) {
      $scope.Grupo_Isla = Grupo_Isla;



      $scope.ok = function() {
        $modalInstance.close($scope.Grupo_Isla);
      };

      $scope.cancel = function() {
        $modalInstance.dismiss('cancel');
      };
    }
  ]);

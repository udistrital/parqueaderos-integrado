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

          "IdGrupoIsla": "",

          "Geometria": "",

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
  .controller('IslaSaveController', ['$scope', '$http', '$uibModalInstance', 'Isla', 'CONFIG',
    function($scope, $http, $modalInstance, Isla, CONFIG) {
      $scope.Isla = Isla;

      var f = [];
      $http.get(CONFIG.WS_URL + '/grupo_isla')
        .success(function(data) {
          data.forEach(function(entry, index) {
            f[index] = entry.Id;
          });
          console.log(f);
	  $scope.IslasGrIds = f;
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

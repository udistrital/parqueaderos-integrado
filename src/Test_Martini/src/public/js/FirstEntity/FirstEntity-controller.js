'use strict';

angular.module('myapp')
  .controller('FirstEntityController', ['$scope', '$modal', 'resolvedFirstEntity', 'FirstEntity',
    function ($scope, $modal, resolvedFirstEntity, FirstEntity) {

      $scope.Firstentities = resolvedFirstEntity;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.FirstEntity = FirstEntity.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        FirstEntity.delete({id: id},
          function () {
            $scope.Firstentities = FirstEntity.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          FirstEntity.update({id: id}, $scope.FirstEntity,
            function () {
              $scope.Firstentities = FirstEntity.query();
              $scope.clear();
            });
        } else {
          FirstEntity.save($scope.FirstEntity,
            function () {
              $scope.Firstentities = FirstEntity.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.FirstEntity = {
          
          "myattr": "",
          
          "myattr1": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var FirstEntitySave = $modal.open({
          templateUrl: 'FirstEntity-save.html',
          controller: 'FirstEntitySaveController',
          resolve: {
            FirstEntity: function () {
              return $scope.FirstEntity;
            }
          }
        });

        FirstEntitySave.result.then(function (entity) {
          $scope.FirstEntity = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('FirstEntitySaveController', ['$scope', '$modalInstance', 'FirstEntity',
    function ($scope, $modalInstance, FirstEntity) {
      $scope.FirstEntity = FirstEntity;

      

      $scope.ok = function () {
        $modalInstance.close($scope.FirstEntity);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);

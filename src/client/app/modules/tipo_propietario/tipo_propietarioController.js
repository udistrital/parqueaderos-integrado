angular.module('Tipo_propietario')
  .controller('Tipo_propietarioCtrl', ["$scope", "$rootScope", "Tipo_propietarioServices", function($scope, $rootScope, Tipo_propietarioServices) {

    $scope.currentTipo_propietario = {};
    $scope.tipo_propietarioFields = [{
      "type": "input",
      "key": "Id",
      "templateOptions": {
        "label": "Id",
        "required": true,
        "type": "number"
      }
    }, {
      "type": "input",
      "key": "Tipo",
      "templateOptions": {
        "label": "Tipo",
        "required": true,
        "type": "text"
      }
    }, {
      "type": "input",
      "key": "Descripcion",
      "templateOptions": {
        "label": "Descripción",
        "required": false,
        "type": "text"
      }
    }];

    // retrieve Tipo_propietario to server
    $scope.retrieveTipo_propietarios = (function() {
      Tipo_propietarioServices.retrieveTipo_propietarios()
        .then(function(result) {
          $rootScope.tipo_propietarios = result;
        }, function(error) {
          alert(error);
        })
    })();


    $scope.openAddTipo_propietario = function() {
      $scope.currentTipo_propietario = {};
      $('#addTipo_propietario.modal').modal('show');
    }

    // save Tipo_propietario to server
    $scope.saveTipo_propietario = function() {
      Tipo_propietarioServices.saveTipo_propietario($scope.currentTipo_propietario)
        .then(function(result) {
          $rootScope.tipo_propietarios.push(result);
        }, function(error) {
          alert(error);
        })
    }

    $scope.openUpdateTipo_propietario = function(data) {
      $scope.currentTipo_propietario = angular.copy(data);
      $('#updateTipo_propietario.modal').modal('show');
    }

    //update data to server
    $scope.updateTipo_propietario = function() {
      Tipo_propietarioServices.updateTipo_propietario($scope.currentTipo_propietario)
        .then(function(result) {
          for (var key in $rootScope.tipo_propietarios) {
            if (result.Id == $rootScope.tipo_propietarios[key].Id)
              $rootScope.tipo_propietarios[key] = result;
          }
        }, function(error) {
          alert(error);
        })
    }

    $scope.openDeleteTipo_propietario = function(data) {
      $scope.currentTipo_propietario = angular.copy(data);
      $('#deleteTipo_propietario.modal').modal('show');
    }

    //delete data to server
    $scope.deleteTipo_propietario = function() {
      Tipo_propietarioServices.deleteTipo_propietario($scope.currentTipo_propietario.Id)
        .then(function(result) {
          for (var key in $rootScope.tipo_propietarios) {
            if ($scope.currentTipo_propietario.Id == $rootScope.tipo_propietarios[key].Id) {
              $rootScope.tipo_propietarios.splice(key, 1);
              break;
            }
          }
        }, function(error) {
          alert(error);
        })
    }

    $scope.emptyTipo_propietario = function() {
      $scope.currentTipo_propietario = {};
    }
  }]);

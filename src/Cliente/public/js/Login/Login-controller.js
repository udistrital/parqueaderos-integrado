'use strict';

angular.module('myapp')
  .controller('LoginController', [
    '$http',
    '$scope',
    '$uibModal',
    'CONFIG',
    '$cookies',
    function(
      $http,
      $scope,
      $modal,
      CONFIG,
      $cookies
    ) {
      //$scope.Login = ["admin", "admin"]

      $scope.enviar = function() {
        console.log($scope.username, $scope.password)
        var completeUrl = CONFIG.WS_URL + '/login?' + "username=" + $scope.username +
          "&password=" + $scope.password //and the same for input 3
        $http({
          method: 'GET',
          url: completeUrl
        }).
        success(function(data, status, headers, config) {
          // here data contains all informations returned by the server
          console.log(data)
          var xsrf = $cookies.get('_xsrf')
          var xsrflist = xsrf.split("|")
          window.xsrf = atob(xsrflist[0])
          console.log(window.xsrf)
          window.location = '#'
        }).
        error(function(data, status, headers, config) {
          //In case your server respond with a 4XX or 5XX error code
          api.showError('Servicio web no está funcionando.')
        })
      }
    }
  ])

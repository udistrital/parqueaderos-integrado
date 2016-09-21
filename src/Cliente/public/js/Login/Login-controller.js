'use strict';

angular.module('myapp')
  .controller('LoginController', ['$http', '$scope', '$uibModal', 'Login',
    function($http, $scope, $modal, Login) {
      $scope.Login = ["admin", "admin"];

      var completeUrl = "http://localhost:8080/v1/login?" + "username="+ $scope.Login[0]
                    + "&password=" + $scope.Login[1] //and the same for input 3
$http({method: 'GET', url: completeUrl}).
success(function(data, status, headers, config) {
   // here data contains all informations returned by the server
   console.log(data);
}).
error(function(data, status, headers, config) {
   //In case your server respond with a 4XX or 5XX error code
});
}]);

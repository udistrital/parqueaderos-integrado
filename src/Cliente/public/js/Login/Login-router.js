'use strict';

angular.module('myapp')
  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/Login', {
        templateUrl: 'views/Login/Login.html',
        controller: 'LoginController',
      })
  }]);

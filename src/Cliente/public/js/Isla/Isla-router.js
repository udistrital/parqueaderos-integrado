'use strict';

angular.module('myapp')
  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/Islas', {
        templateUrl: 'views/Isla/Islas.html',
        controller: 'IslaController',
        resolve: {
          resolvedIsla: ['Isla', function(Isla) {
            return Isla.query();
          }]
        }
      })
  }]);

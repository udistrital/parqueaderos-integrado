'use strict';

angular.module('myapp')
  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/Propietarios', {
        templateUrl: 'views/Propietario/Propietarios.html',
        controller: 'PropietarioController',
        resolve: {
          resolvedPropietario: ['Propietario', function(Propietario) {
            return Propietario.query();
          }]
        }
      })
  }]);

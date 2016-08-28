'use strict';

angular.module('myapp')
  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/Vehiculos', {
        templateUrl: 'views/Vehiculo/Vehiculos.html',
        controller: 'VehiculoController',
        resolve: {
          resolvedVehiculo: ['Vehiculo', function(Vehiculo) {
            return Vehiculo.query();
          }]
        }
      })
  }]);

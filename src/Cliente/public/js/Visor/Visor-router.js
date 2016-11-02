'use strict';

angular.module('myapp')
    .config(['$routeProvider', function($routeProvider) {
        $routeProvider
            .when('/Visor', {
                templateUrl: 'views/Visor/Visor.html',
                //controller: 'VisorController',
                // resolve: {
                //   resolvedVehiculo: ['Vehiculo', function(Vehiculo) {
                //     return Vehiculo.query();
                //   }]
                // }
            })
    }])

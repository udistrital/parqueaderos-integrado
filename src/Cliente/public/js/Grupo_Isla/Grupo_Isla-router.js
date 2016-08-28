'use strict';

angular.module('myapp')
  .config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/Grupo_islas', {
        templateUrl: 'views/Grupo_Isla/Grupo_islas.html',
        controller: 'Grupo_IslaController',
        resolve: {
          resolvedGrupo_Isla: ['Grupo_Isla', function(Grupo_Isla) {
            return Grupo_Isla.query();
          }]
        }
      })
  }]);

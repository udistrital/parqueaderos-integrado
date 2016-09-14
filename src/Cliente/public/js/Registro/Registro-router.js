'use strict';

angular.module('myapp')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/Registros', {
        templateUrl: 'views/Registro/Registros.html',
        controller: 'RegistroController',
        resolve:{
          resolvedRegistro: ['Registro', function (Registro) {
            return Registro.query();
          }]
        }
      })
    }]);

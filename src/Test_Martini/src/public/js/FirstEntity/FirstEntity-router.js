'use strict';

angular.module('myapp')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/Firstentities', {
        templateUrl: 'views/FirstEntity/Firstentities.html',
        controller: 'FirstEntityController',
        resolve:{
          resolvedFirstEntity: ['FirstEntity', function (FirstEntity) {
            return FirstEntity.query();
          }]
        }
      })
    }]);

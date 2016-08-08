'use strict';

angular.module('myapp')
  .factory('Grupo_Isla', ['$resource', function ($resource) {
    return $resource('v1/grupo_isla/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

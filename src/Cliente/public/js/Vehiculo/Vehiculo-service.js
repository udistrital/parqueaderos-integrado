'use strict';

angular.module('myapp')
  .factory('Vehiculo', ['$resource', function ($resource) {
    return $resource('v1/vehiculo/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

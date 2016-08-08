'use strict';

angular.module('myapp')
  .factory('Propietario', ['$resource', function ($resource) {
    return $resource('v1/propietario/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

'use strict';

angular.module('myapp')
  .factory('Isla', ['$resource', function ($resource) {
    return $resource('v1/isla/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

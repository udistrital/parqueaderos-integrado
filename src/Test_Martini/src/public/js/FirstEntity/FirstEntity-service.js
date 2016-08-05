'use strict';

angular.module('myapp')
  .factory('FirstEntity', ['$resource', function ($resource) {
    return $resource('myapp/Firstentities/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

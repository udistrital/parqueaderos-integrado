'use strict';

angular.module('myapp')
  .factory('Propietario', ['$resource', "CONFIG", function($resource) {
    return $resource(CONFIG.WS_URL + '/propietario/:id', {}, {
      'query': {
        method: 'GET',
        isArray: true
      },
      'get': {
        method: 'GET'
      },
      'update': {
        method: 'PUT'
      }
    });
  }]);

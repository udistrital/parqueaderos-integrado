'use strict';

angular.module('myapp')
  .factory('Isla', ['$resource', "CONFIG", function($resource) {
    return $resource(CONFIG.WS_URL + '/isla/:id', {}, {
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

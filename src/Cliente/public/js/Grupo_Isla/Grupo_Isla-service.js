'use strict';

angular.module('myapp')
  .factory('Grupo_Isla', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/grupo_isla/:id', {}, {
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

'use strict';

angular.module('myapp')
  .factory('Propietario', ['$resource', 'CONFIG', function($resource, CONFIG) {
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

'use strict';

angular.module('myapp')
  .factory('Registro', ['$resource', 'CONFIG', function ($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/registro/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

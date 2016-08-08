'use strict';

angular.module('myapp')
  .factory('Tipo_Propietario', ['$resource', function ($resource) {
    return $resource('v1/tipo_propietario/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

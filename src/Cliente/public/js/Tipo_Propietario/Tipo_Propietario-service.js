'use strict';

angular.module('myapp')
  .factory('Tipo_Propietario', ['$resource', "CONFIG", function ($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/tipo_propietario/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);

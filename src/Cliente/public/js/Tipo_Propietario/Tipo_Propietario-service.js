'use strict';

angular.module('myapp')
  .factory('Tipo_Propietario', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/tipo_propietario/:id', {}, {
      'query': {
        method: 'GET',
        headers: { '_xsrf': document.cookie },
        isArray: true
      },
      'get': {
        method: 'GET',
        headers: { 'something': 'anything' }
      },
      'save': {
        method:'POST',
        headers: { 'something': 'anything' }
      },
      'update': {
        method: 'PUT',
        headers: { 'something': 'anything' }
      },
      'delete': {
        method:'DELETE',
        headers: { 'something': 'anything' }
      }
    });
  }]);

'use strict';

angular.module('myapp')
  .factory('Tipo_Propietario', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/tipo_propietario/:id', {}, {
      'query': {
        method: 'GET',
        headers: {
          '_xsrf': document.cookie
        },
        isArray: true,
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'get': {
        method: 'GET',
        headers: {
          'something': 'anything'
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'save': {
        method: 'POST',
        headers: {
          'something': 'anything'
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'update': {
        method: 'PUT',
        headers: {
          'something': 'anything'
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'delete': {
        method: 'DELETE',
        headers: {
          'something': 'anything'
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      }
    });
  }]);

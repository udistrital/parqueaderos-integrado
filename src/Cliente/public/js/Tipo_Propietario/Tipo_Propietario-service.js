'use strict';

angular.module('myapp')
  .factory('Tipo_Propietario', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/tipo_propietario/:id', {}, {
      'query': {
        method: 'GET',
        isArray: true,
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response.data)
          }
        }
      },
      'get': {
        method: 'GET',
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response.data)
          }
        }
      },
      'save': {
        method: 'POST',
        headers: {
          'X-Xsrftoken': window.xsrf
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response.data)
          }
        }
      },
      'update': {
        method: 'PUT',
        headers: {
          'X-Xsrftoken': window.xsrf
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response.data)
          }
        }
      },
      'delete': {
        method: 'DELETE',
        headers: {
          'X-Xsrftoken': window.xsrf
        },
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response.data)
          }
        }
      }
    })
  }])

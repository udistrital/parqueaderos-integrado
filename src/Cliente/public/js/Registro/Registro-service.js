'use strict';

angular.module('myapp')
  .factory('Registro', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/registro/:id', {}, {
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
      'save': { //stackoverflow.com/questions/20584367/how-to-handle-resource-service-errors-in-angularjs
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
      'get': {
        method: 'GET',
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

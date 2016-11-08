'use strict';

angular.module('myapp')
  .factory('Isla', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/isla/:id', {}, {
      'query': {
        method: 'GET',
        isArray: true,
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response)
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
            api.showError(response)
          }
        }
      },
      'get': {
        method: 'GET',
        interceptor: {
          responseError: function(response) {
            console.log(response)
            api.showError(response)
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
            api.showError(response)
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
            api.showError(response)
          }
        }
      }
    })
  }])

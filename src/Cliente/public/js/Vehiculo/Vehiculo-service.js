'use strict';

angular.module('myapp')
  .factory('Vehiculo', ['$resource', 'CONFIG', function($resource, CONFIG) {
    return $resource(CONFIG.WS_URL + '/vehiculo/:id', {}, {
      'query': {
        method: 'GET',
        isArray: true,
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'save': { //stackoverflow.com/questions/20584367/how-to-handle-resource-service-errors-in-angularjs
        method: 'POST',
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'get': {
        method: 'GET',
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      },
      'update': {
        method: 'PUT',
        interceptor: {
          responseError: function(response) {
            console.log(response)
            window.alert(response.data)
          }
        }
      }
    });
  }]);

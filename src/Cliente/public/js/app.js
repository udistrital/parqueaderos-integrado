// Declare app level module which depends on filters, and services
angular.module('myapp', ['ngResource', 'ngRoute', 'ui.bootstrap', 'ui.date'])
  .constant('CONFIG', {
    WS_URL: 'v1',
  })
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/home/home.html',
        controller: 'HomeController'})
      .otherwise({redirectTo: '/'});
  }]);

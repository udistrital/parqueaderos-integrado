// Declare app level module which depends on filters, and services
angular.module('myapp', [
    'ngResource',
    'ngRoute',
    'ngCookies',
    'ui.bootstrap',
    'ui.date'
  ])
  .constant('CONFIG', {
    WS_URL: 'v1',
  })
  .config([
    '$locationProvider',
    '$routeProvider',
    '$resourceProvider',
    function(
      $locationProvider,
      $routeProvider,
      $resourceProvider
    ) {
      $routeProvider
        .when('/', {
          templateUrl: 'views/home/home.html',
          controller: 'HomeController'
        })
        .otherwise({
          redirectTo: '/'
        })
    }
  ])
  .run(['$cookies',
    function($cookies) {
      var xsrf = $cookies.get('_xsrf')
      if (xsrf != undefined) {
        var xsrflist = xsrf.split("|")
        window.xsrf = atob(xsrflist[0])
      }
    }
  ])

var api = {
  showError: function(response) {
    var msj = ''
    if (typeof(response) === 'object') {
      msj = response.data
    } else {
      msj = response
    }

    if (response.status === 401 || response.status === 403) {
      //Permission Deny "Unauthorized", Forbidden,
      window.alert('Necesita permisos para hacer eso, renueve su sesión.')
      window.location = '#/Login'
      var iframe = document.createElement('iframe')
      iframe.src = 'data:text/html;charset=utf-8;base64,' + btoa(msj)
      $('#modal-text').html(iframe)
      $("#myModal").modal('show')
    } else {
      $('#modal-text').html(msj)
      $("#myModal").modal('show')
    }

  },
  showMessage: function(msj) {
    window.alert(msj)
  }
}

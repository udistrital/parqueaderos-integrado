angular.module("Tipo_propietario", []);

angular.module("Propietario", []);

angular.module("Vehiculo", []);

angular.module("circe", ["ui.router", "formly", "formlyBootstrap", "Tipo_propietario", "Propietario", "Vehiculo"])
  .constant("CONFIG", {
    "WS_URL": "http://vagrant:8080/v1",
  })
  .config(["$stateProvider", "$urlRouterProvider", "CONFIG", function($stateProvider, $urlRouterProvider, CONFIG) {
    $stateProvider
      .state("Tipo Propietario", {
        url: "/tipos_propietario",
        templateUrl: "modules/tipo_propietario/views/tipos_propietario.html",
        controller: "Tipo_propietarioCtrl"
      })
      .state("Propietario", {
        url: "/propietarios",
        templateUrl: "modules/propietario/views/propietarios.html",
        controller: "PropietarioCtrl"
      })
      .state("Vehiculo", {
        url: "/vehiculos",
        templateUrl: "modules/vehiculo/views/vehiculos.html",
        controller: "VehiculoCtrl"
      })
    $urlRouterProvider.otherwise("/tipos_propietario");

  }]);

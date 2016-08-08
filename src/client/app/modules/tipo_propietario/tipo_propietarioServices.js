angular.module('Tipo_propietario')
  .constant("CONFIG", {
    "WS_URL": "http://localhost",
  })
  .factory('Tipo_propietarioServices', ["$http", "$q", "CONFIG", function($http, $q, CONFIG) {
    return {
      retrieveTipo_propietarios: function() {

        var defer = $q.defer();
        $http.get(CONFIG.WS_URL + "/tipo_propietario")
          .success(function(result) {
            defer.resolve(result)
          })
          .error(function(error) {
            defer.reject(error)
          });
        return defer.promise;
      },
      saveTipo_propietario: function(Tipo_propietarioObj) {

        var defer = $q.defer();
        $http.post(CONFIG.WS_URL + "/tipo_propietario", Tipo_propietarioObj)
          .success(function(result) {
            defer.resolve(result)
          })
          .error(function(error) {
            defer.reject(error)
          });
        return defer.promise;
      },
      updateTipo_propietario: function(Tipo_propietarioObj) {

        var defer = $q.defer();
        $http.put(CONFIG.WS_URL + "/tipo_propietario/" + Tipo_propietarioObj.Id, Tipo_propietarioObj)
          .success(function(result) {
            defer.resolve(result)
          })
          .error(function(error) {
            defer.reject(error)
          });

        return defer.promise;
      },
      deleteTipo_propietario: function(id) {

        var defer = $q.defer();
        $http.delete(CONFIG.WS_URL + "/tipo_propietario/" + id)
          .success(function(result) {
            defer.resolve(result)
          })
          .error(function(error) {
            defer.reject(error)
          });

        return defer.promise;
      },
      retrieveTipo_propietarioById: function(id) {

        var defer = $q.defer();
        $http.get(CONFIG.WS_URL + "/tipo_propietario/" + id)
          .success(function(result) {
            defer.resolve(result)
          })
          .error(function(error) {
            defer.reject(error)
          });
        return defer.promise;
      }
    }
  }]);

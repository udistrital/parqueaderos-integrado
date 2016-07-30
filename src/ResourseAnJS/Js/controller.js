angular.module("MyFirstApp")
	.controller("FirstController", function($scope, UrlResource){
		$scope.tipo_propietarios = UrlResource.query();
	})
	.controller("SecondController", function($scope, UrlResource, $routeParams){
		$scope.tipo_propietarios = UrlResource.get({id: $routeParams.id});
		$scope.removeReg = function(){
			UrlResource.delete({id: $routeParams.id});
		}
		$scope.modReg = function(){
                        UrlResource.update({id: $routeParams.id}, $scope.tipo_propietario);			
		}	
	})
	.controller("ThirdController",function($scope, UrlResource){
	        $scope.tipo_propietario = {};
                $scope.saveReg = function(){
			UrlResource.save($scope.tipo_propietario); 
			$scope.tipo_propietario = {};
                	}
	});

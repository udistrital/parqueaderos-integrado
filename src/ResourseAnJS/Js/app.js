angular.module("MyFirstApp", ["ngRoute", "ngResource"])
.config(function($routeProvider){
		$routeProvider
			.when("/", {
				controller : "FirstController",
				templateUrl : "Templates/home.html"
				})
			.when("/one/:id",{
					controller : "SecondController",
					templateUrl : "Templates/room.html"
				})
			.when("/new",{
					controller : "ThirdController",
					templateUrl : "Templates/new.html"
				});
});

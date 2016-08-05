/**
 * Created by Shivali on 6/29/15.
 */

    


angular.module("circe", ["ui.router","formly","formlyBootstrap"
    

])
    .config(["$stateProvider", "$urlRouterProvider", function ($stateProvider, $urlRouterProvider) {
        
            $stateProvider
            
         
        
            $urlRouterProvider.otherwise("/");
        
    }]);


angular.module('Eriksson')
.directive("formGroup", function () {
  return {
    restrict: "C",
    //get the controller in the link function
    require: "formGroup",
    link: function ($scope, element, attributes, controller) {
      var errorList = {};
      controller.inputError = function () {
        element.addClass("has-error");
        element.find('.form-control-feedback').hide();
      };
      controller.inputValid = function () {
        element.removeClass("has-error");
        element.addClass("has-success");
        element.find('.form-control-feedback').show();
      };

      if (element.hasClass('has-feedback')) {
        element.append('<span class="glyphicon glyphicon-ok form-control-feedback hide"></span>');
      }
    },
    //the controller is initialized in link function
    controller: function() {return {};}
  };
})

.directive("input", function () {
  return {
    restrict: "E",
    //require controlllers of ngModel and parent directive
    //formGroup
    //they're injected as array parameter of link function
    require: ["?ngModel", "^?formGroup"],
    link: function ($scope, element, attributes, controllers) {
      var modelController = controllers[0];
      var formGroupController = controllers[1];
      if (!modelController || !formGroupController) return;
      var hasBeenVisited = false;
      // // Benötigte CSS-Klasse.
      // Benötigte CSS-Klasse.
      // element.addClass('form-control');
      // check if user has left the field
      element.on("blur", function () {
        $scope.$apply(function () { hasBeenVisited = true; });
      });
      // Watch the validity of the input
      $scope.$watch(function () { return modelController.$invalid && hasBeenVisited; }, function () {
        // $emit messages to the control group
        if (modelController.$invalid && hasBeenVisited) {
          formGroupController.inputError();
        } 
        else if(hasBeenVisited) {
          formGroupController.inputValid();
        }
      });
    }
  };
});

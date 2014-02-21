angular.module('Eriksson').controller('ProjectEditCtrl',
    ['$scope', '$routeParams', 'Projects', 'ProjectFormSchema', function($scope, $routeParams, Projects, ProjectFormSchema) {

  // Projektdaten laden.
  $scope.project = {};

  Projects.get({id: $routeParams.id}, function(p) {
      console.log(p);
    $scope.project = p;
  });

  // Form-Definition:
  $scope.schema = ProjectFormSchema;

  $scope.submitForm = function() {
    // Ist das Formular gültig?
    if ($scope.dynaForm.$valid) {
      console.log($scope.project);
      Projects.update($scope.project, function(data) {
        toastr.success('Erfolgreich gespeichert.');
      },
      function() {
        toastr.error('Fehler beim Speichern.');
      });
    }
  };
}]);

angular.module('Eriksson').controller('ProjectNewCtrl',
['$scope', 'Projects', 'ProjectFormSchema', function($scope, Projects, ProjectFormSchema) {

    // Hier speichert die Form hin.
    $scope.project = {
        OrderDate: new Date(),
        Techs:[]
    };

    // Form-Definition:
    $scope.schema = ProjectFormSchema;

    $scope.submitForm = function() {

        // Ist das Formular gültig?
        if ($scope.dynaForm.$valid) {
            console.log($scope.project);
            Projects.save($scope.project, function(data) {
                toastr.success('Erfolgreich gespeichert.');
            },
            function() {
                toastr.error('Fehler beim Speichern.');
            });
        }
    };
}]);

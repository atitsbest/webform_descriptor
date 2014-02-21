angular.module('Eriksson').controller('ProjectIndexCtrl', ['$scope', '$location', 'Projects', function($scope, $location, Projects) {

    $scope.$table = $('#table_id');

    $scope.$table.dataTable({
        "bPaginate": false,
        "bLengthChange": false,
        "bFilter": true,
        "bSort": true,
        "bInfo": false,
        "bAutoWidth": true,
        "bDeferRender": true,
        "sAjaxSource": "api/projects",
        "aoColumns": [
            {"sTitle": "Projekttitel", "mData": "Name"},
            {"sTitle": "Projktleiter", "mData": "Leader"},
            {"sTitle": "Technologien", "mData": "Techs"},
            {"sTitle": "Bestellvolumen [€]", "mData": "OrderAmount"},
            {"sTitle": "Bestellvolumen [PT]", "mData": "OrderAmountDays"}
        ],
        "sAjaxDataProp": "",
    });

    // Projekt öffnen.
    $scope.$table.on('click', 'tbody > tr', function() {
        $scope.$apply(function() { $location.path('/projects/edit/1234'); });
    });
}]);

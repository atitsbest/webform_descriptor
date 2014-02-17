angular.module('Eriksson').controller('ProjectNewCtrl', ['$scope', 'Projects', function($scope, Projects) {

  $scope.allRisks =['A','B','C','D'];
  $scope.allAccountingModes = ['Fixpreis','Nach Aufwand','Stundenpaket'];
  $scope.allStates = ['beauftragt','abgeschlossen'];
  //
  // Hier speichert die Form hin.
  $scope.project = {
    name: 'Eriksson',
    leader: 'RatAn',
    risk: $scope.allRisks[0],
    accountingMode: $scope.allAccountingModes[0],
    state: $scope.allStates[0],
    orderDate: moment().format("DD.MM.YYYY"),
    techs: {}
  };

  // Form-Definition:
  $scope.schema = {
    name: 'dynaForm',
    'ng-submit': 'submitForm()',
    fields: {
      'name': {
        title: 'Projektname', required: true, // oder: "Das Name muss angegeben werden!"
        minlength: 2, // oder { val: 2, msg: "Der Name ist zu kurz!" }
        maxlength: 30
      },
      'leader': {
        title: 'Projektleiter', required: true,
        pattern: /^[A-Z][a-z]{2}[A-Z][a-z]$/
      },
      'orderAmount': {
        title: 'Auftragssumme', required: true,
        type: 'money'
      },
      'accountingMode': {
        title: 'Verrechnungsmodus',
        type: 'radio',
        values: $scope.allAccountingModes
      },
      'risk': {
        title: 'Risiko-Klasse',
        type: 'select',
        values: $scope.allRisks
      },
      'techs': {
        title: 'Technologien',
        type: 'checkbox',
        values: ['C#','F#','JavaScript','SharePoint','Progress']
      },
      'orderDate': {
        title: 'Bestelldatum', required: true,
        type: 'date'
      }

      /*,
        'orderAmount': {},
        'orderDate': {}*/
    }
  };

  // HAAAAACK!
  setTimeout(function() {
    $('.datepicker').datetimepicker({
      language: 'de',
      pickTime: false
    });
  }, 200);

  $scope.submitForm = function() {
    alert('SUBMIT');

    // Ist das Formular gültig?
    if ($scope.dynaForm.$valid) {
      console.log($scope.project);
      Projects.save($scope.project, function(data) {
        alert('Erfolgreich gespeichert.');
      },
      function() {
        alert('Fehler beim Speichern.');
      });
    }
  };
}]);

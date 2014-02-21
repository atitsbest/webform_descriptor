angular.module('Eriksson').value('ProjectFormSchema', {
    name: 'dynaForm',
    'ng-submit': 'submitForm()',
    fields: {
        'Name': {
            title: 'Projektname', required: true, // oder: "Der Name muss angegeben werden!"
            minlength: 2, // oder { val: 2, msg: "Der Name ist zu kurz!" }
            maxlength: 30
        },
        'Customer': {
            title: 'Kunde', required: true,
            minlength: 2, maxlength: 30
        },
        'Leader': {
            title: 'Projektleiter', required: true,
            pattern: /^[A-Z][a-z]{2}[A-Z][a-z]$/
        },
        'OrderAmount': {
            title: 'Auftragssumme', required: true,
            type: 'money'
        },
        'OrderAmountDays': {
            title: 'Auftragsvolumen PT', required: true,
            type: 'number'
        },
        'AccountingMode': {
            title: 'Verrechnungsmodus',
            type: 'radio',
            values: ['Fixpreis','Nach Aufwand','Stundenpaket']
        },
        'Risk': {
            title: 'Risiko-Klasse',
            type: 'select',
            values: ['A','B','C','D']
        },
        'State': {
            title: 'Projektstatus',
            type: 'radio',
            values: ['beauftragt','abgeschlossen']
        },
        'Techs': {
            title: 'Technologien'
        },
        'OrderDate': {
            title: 'Bestelldatum', required: true,
            type: 'date'
        },
        'BMDOrderNumber': {
            title: 'BMD Bestellnummer',
            maxlength: 30
        },
        'Comment': {
            title: 'Kommentar',
            type: 'textarea',
            maxlength: 1024
        }
    }
});


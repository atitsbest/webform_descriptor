angular.module('Eriksson', ['ngResource'])
.config(function($interpolateProvider) {
    $interpolateProvider.startSymbol('-%');
    $interpolateProvider.endSymbol('%-');
});

$(function() {
  $('#table_id').dataTable({
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
});


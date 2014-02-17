angular.module('Eriksson')
.factory('Projects', ['$resource', function ($resource) {
  return $resource('/api/projects/:id', {}, {
    'query': { method: 'GET', isArray: true},
    'get': { method: 'GET'},
    'update': { method: 'PUT'}
  });
}]);

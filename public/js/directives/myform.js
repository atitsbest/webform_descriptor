angular.module('Eriksson')

.directive('myForm', function ($compile) {
    var _validationKeys = ['required', 'maxlength', 'minlength', 'pattern'],
        inputProviders = { };

    /**
     * Neuen InputProvider registrieren.
     **/
    var _registerInputProvider = function(name, createFn) {
      if (_(inputProviders).has(name)) { throw "Input-Provider für " + name + " bereits registriert!"; }
      inputProviders[name] = createFn;
    };

    /**
     *
     **/
    var _createAttributeString = function(obj, keysToOmit) {
        var filteredObj = _.omit.apply(null, [obj].concat(keysToOmit));

        return _.map(filteredObj, function (v, k) { return k + '="' + v + '"'; })
            .join(' ');
    };

    /**
     *
     **/
    var _createFormElement = function (schema) {
        var formAttrs = _createAttributeString(schema, ['fields']);
        return angular.element('<form ' + formAttrs + ' novalidate></form>');
    };

    /**
     *
     */
    var _createRadioElements = function(fieldName, field, modelName) {
        var radios = _(field.values).map(function (value) {
            return '<div class="radio"><label><input type="radio" name="' + fieldName + '" id="' + fieldName + '" value="' + value + '" ng:model="'+modelName+'.' + fieldName + '">' + value + '</label></div>';
        }).join('');

        return angular.element(radios);
    };
    _registerInputProvider('radio', _createRadioElements);

    /**
     *
     */
    var _createCheckboxElements = function(fieldName, field, modelName) {
        var checkboxes = _(field.values).map(function (value) {
            return '<div class="checkbox"><label><input type="checkbox" name="' + fieldName + '" id="' + fieldName + '" ng:model="'+modelName+'.' + fieldName + '[\''+value+'\']">' + value + '</label></div>';
        }).join('');

        return angular.element(checkboxes);
    };
    _registerInputProvider('checkbox', _createCheckboxElements);

    /**
     *
     **/
    var _createSelectElement = function(fieldName, field, modelName) {
        var html = '<select class="form-control" name="'+fieldName+'" ng:model="'+modelName+'.'+fieldName+'" ng:options="m for m in schema.fields.'+fieldName+'.values"></select>';
        return html;
    };
    _registerInputProvider('select', _createSelectElement);

    /**
     *
     **/
    var _createTextInputElement = function(fieldName, field, modelName) {
        var vals = _.pick.apply(null, [field].concat(_validationKeys)),
            valStr = _(vals).map(function (v, k) {
                return 'ng:' + k + '="' + v + '"';
            }).join(" ");

        return angular.element(
                '<input type="' + (field.type || 'text') + '" class="form-control" id="' + fieldName + '" name="' + fieldName + '"  ng:model="'+modelName+'.' + fieldName + '" ' + valStr + ' />');
    };
    _registerInputProvider('text', _createTextInputElement);

    /**
     * Money Input Provider.
     **/
    _registerInputProvider('money', function(fieldName, field, modelName) {
        var extField = _.extend(field, {type: 'text', pattern: /^\d+(\.\d{1,2})?$/ });
        return _createTextInputElement(fieldName, extField, modelName);
    });

    /**
     * Money Input Provider.
     **/
    _registerInputProvider('date', function(fieldName, field, modelName) {
        var extField = _.extend(field, {type: 'text' }),
            element =  _createTextInputElement(fieldName, extField, modelName);

        return element.addClass('datepicker');
    });

    /**
     *
     **/
    var _createInputElement = function (fieldName, field, modelName) {
        var type = field.type || 'text',
            provider = inputProviders[type];

        if (provider === undefined) { throw "Ungültiger Feldtype ' field.type ' im Formular!"; }
        return provider(fieldName, field, modelName);
    };

    /**
     *
     **/
    var _createLabelElement = function (fieldName, field) {
        return angular.element('<label for="' + fieldName + '">' + (field.title || fieldName) + '</label>');
    };

    /**
     *
     **/
    var _createErrorMessageElements = function (formName, fieldName, field) {
        var vals = _.pick.apply(null, [field].concat(_validationKeys));

        return _(vals).map(function (v, k) {
            var formField = formName + '.' + fieldName,
                valExpr = formField + '.$error.' + k /* && !'+formField+'.$pristine'*/;
            return angular.element('<span class="help-block" ng:show="' + valExpr + '">Fehler ' + k + '</span>');
        });
    };

    return {
        restrict: 'E', // Nur Elemente
        transclude: false, // Inner-Elemente mitnehmen.
        replace: true,
        controller: function ($element, $scope, $attrs) {
            // Form erstellen.
            var schema = $scope[$attrs.schema],
                modelName = $attrs.model,
                form = _createFormElement(schema);

            var fieldNames = _(schema.fields).keys();

            // Alle Felder erstellen.
            _(fieldNames).each(function (fn) {
                var field = schema.fields[fn],
                    group = angular.element('<div class="form-group" ng:class="{\'hasError\':'+schema.name+'.'+fn+'.$invalid}"></div>');

                group.append(_createLabelElement(fn, field));
                group.append(_createInputElement(fn, field, modelName));
                _(_createErrorMessageElements(schema.name, fn, field)).each(function (m) { group.append(m); });
                form.append(group);
            });

            // Html-Elemente der Directive in die Form hängen.
            form.append($element.children());
            // Directive durch die Form ersetzen.
            $element.replaceWith($compile(form)($scope));
            return function preLink($scope, iElement, iAttrs, controller) {};
        }
    };
});

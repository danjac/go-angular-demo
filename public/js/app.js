'use strict';

angular.module('postApp', ['postApp.controllers', 'postApp.services', 'postApp.directives'])
    .constant('urls', {
        api: '/api/:id'
    })
    .config(['$resourceProvider', '$httpProvider', function ($resourceProvider, $httpProvider) {
        // Don't strip trailing slashes from calculated URLs
        $resourceProvider.defaults.stripTrailingSlashes = false;
        $httpProvider.defaults.xsrfCookieName = "csrf_token";
        $httpProvider.defaults.xsrfHeaderName = "X-CSRF-Token";
    }]);

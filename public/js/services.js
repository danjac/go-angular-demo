'use strict';

angular.module('postApp.services', ['ngResource', 'postApp'])
    .service('Post', ['$resource', 'urls', function ($resource, urls) {
        return $resource(urls.api, {id: '@id'});
    }]);

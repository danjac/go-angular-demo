'use strict';

describe('PostCtrl', function () {
    // mocking/DI
    var scope;

    beforeEach(angular.mock.module('postApp'));
    beforeEach(angular.mock.inject(function ($rootScope, $controller) {
        scope = $rootScope.$new();
        $controller('PostCtrl', {$scope: scope});
    }));
    it('should include some empty messages', function () {
        expect(scope.posts).toBe([]);
    });
});

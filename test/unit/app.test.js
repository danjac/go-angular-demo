'use strict';

describe('PostCtrl', function () {
    // mocking/DI
    var scope, $httpBackend;

    beforeEach(angular.mock.module('postApp'));

    beforeEach(angular.mock.inject(function ($rootScope, $controller, _$httpBackend_) {

        $httpBackend = _$httpBackend_;
        $httpBackend.when("GET", "/api/").respond([
            {
                id: 1,
                content: "test1"
            },
            {
                id: 2,
                content: "test2"
            },
            {
                id: 3,
                content: "test3"
            }
        ]);

        $httpBackend.when("POST", "/api/").respond({
            id: 4,
            content: "test4"
        });

        $httpBackend.when("DELETE", "/api/1/").respond("Post deleted");

        scope = $rootScope.$new();
        $controller('PostCtrl', {$scope: scope});

    }));

    it('should add a post', function () {
        $httpBackend.flush();
        scope.newPost.content = "test4";
        scope.addPost();
        expect(scope.posts.length).toBe(4);
        expect(scope.newPost.content).toBe(undefined);
    });

    it('should delete a post', function () {
        $httpBackend.flush();
        scope.deletePost(scope.posts[0]);
        expect(scope.posts.length).toBe(2);
    });

    it('should include some posts', function () {
        $httpBackend.flush();
        expect(scope.posts.length).toBe(3);
    });
});

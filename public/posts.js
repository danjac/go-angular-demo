'use strict';

angular.module('postApp', ['ngResource'])
    .config(['$resourceProvider', function ($resourceProvider) {
        // Don't strip trailing slashes from calculated URLs
        $resourceProvider.defaults.stripTrailingSlashes = false;
    }])
    .service('Post', ['$resource', function ($resource) {
        return $resource("/api/:id", {id: '@id'});
    }])
    .controller('PostCtrl', ['$scope', '$interval', 'Post', function ($scope, $interval, Post) {

        $scope.$watch('content', function () {
            if ($scope.content) {
                var contentLength = $scope.content.length;
                $scope.charsRemaining = 140 - contentLength;
                $scope.contentTooLong = $scope.charsRemaining < 10;
            }
        });

        function getPosts() {
            Post.query().$promise.then(function (posts) {
                $scope.posts = posts;
            });
        }
        $scope.deletePost = function (post) {
            var index = $scope.posts.indexOf(post);
            $scope.posts.splice(index, 1);
            post.$delete();
        };
        $scope.addPost = function () {
            var post = new Post({content: $scope.content});
            $scope.content = "";
            $scope.posts.splice(0, 0, post);
            post.$save();
        };
        getPosts();
        $interval(getPosts, 5000);
    }]);

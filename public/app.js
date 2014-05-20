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

        $scope.newPost = new Post();

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
            $scope.posts.splice(0, 0, $scope.newPost);
            $scope.newPost.$save();
            $scope.newPost = new Post();
        };
        getPosts();
        $interval(getPosts, 5000);
    }])
    .directive('contentLengthTracker', function () {
        return {
            link: function ($scope, element, attrs) {
                var maxLength = parseInt(attrs.maxlength, 10),
                    showWarningAt = maxLength / 10;
                $scope.$watch(
                    function () { return $scope.content; },
                    function () {
                        if ($scope.content) {
                            $scope.charsRemaining = maxLength - $scope.content.length;
                            $scope.showWarning = $scope.charsRemaining <= showWarningAt;
                        }
                    }
                );
            },
            restrict: 'E',
            replace: true,
            scope: {
                content: '=',
                maxlength: "@"
            },
            template: '<div ng-show="content" class="bg-info" ng-class="{\'bg-danger\': showWarning}">{{charsRemaining}} characters remaining</div>'
        };
    });


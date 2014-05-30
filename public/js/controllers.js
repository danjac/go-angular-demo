'use strict';

angular.module('postApp.controllers', ['postApp.services'])
    .controller('PostCtrl', ['$scope', 'Post', function ($scope, Post) {

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
    }]);


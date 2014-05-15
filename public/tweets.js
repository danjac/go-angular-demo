'use strict';

angular.module('tweetApp', ['ngResource'])
    .controller('TweetCtrl', ['$scope', '$interval', '$resource', function ($scope, $interval, $resource) {
        var Tweet = $resource("/api/:id", {id: '@id'});
        function reload() {
            Tweet.query().$promise.then(function (tweets) {
                $scope.tweets = tweets;
            });
        }
        $scope.deleteTweet = function (tweet) {
            tweet.$delete(function () { reload(); });
        };
        $scope.addTweet = function () {
            new Tweet({content: $scope.content}).$save(function () {
                $scope.content = "";
                reload();
            });
        };
        reload();
        $interval(reload, 500);
    }]);

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
            var index = $scope.tweets.indexOf(tweet);
            $scope.tweets.splice(index, 1);
            tweet.$delete();
        };
        $scope.addTweet = function () {
            var tweet = new Tweet({content: $scope.content});
            $scope.content = "";
            $scope.tweets.splice(0, 0, tweet);
            tweet.$save();
        };
        reload();
        $interval(reload, 10000);
    }]);

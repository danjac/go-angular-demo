'use strict';

angular.module('tweetApp', ['ngResource'])
    .service('Tweet', ['$resource', function ($resource) {
        return $resource("/api/:id", {id: '@id'});
    }])
    .controller('TweetCtrl', ['$scope', '$interval', 'Tweet', function ($scope, $interval, Tweet) {
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

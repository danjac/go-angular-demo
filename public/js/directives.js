'use strict';

angular.module('postApp.directives', [])
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


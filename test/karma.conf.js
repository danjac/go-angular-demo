'use strict';

module.exports = function(config){
    config.set({

        basePath : '../',

        files : [
            'public/bower_components/angular/angular.min.js',
            'public/bower_components/angular-resource/angular-resource.min.js',
            'public/js/app.js',
            'public/js/controllers.js',
            'public/js/services.js',
            'public/js/directives.js',
            'test/angular-mocks.js',
            'test/unit/**/*.js'
        ],

        autoWatch : true,

        frameworks: ['jasmine'],

        browsers : ['Firefox'],

        plugins : [
            'karma-chrome-launcher',
            'karma-firefox-launcher',
            'karma-jasmine'
        ],

        junitReporter : {
            outputFile: 'test_out/unit.xml',
            suite: 'unit'
        }

    });
};

The Frontend Service
=======================

The frontend service makes use of backend services and shows information to the
user. This is the part that is a more traditional web app.

# Architecture

The frontend is written in [CoffeeScript](http://coffeescript.org/) and uses
require.js ([AMD](http://requirejs.org/docs/whyamd.html)) for defining modules.

# Getting Started

Here's what you need to do to get the frontend running.

## Install grunt

[Grunt](http://gruntjs.com/) is required for application development.

    $ npm install -g grunt-cli

## Install Dependencies

Install the dependencies using npm. The dependencies are defined in
[package.json](package.json) so you just need to run:

    $ npm install

/*jslint sloppy: true, vars: true, white: true, nomen: true */
/*global module, require, __dirname */

var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res) {
    res.render('index.html', { title: 'Hello World' });
});

module.exports = router;

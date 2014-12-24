/*jslint sloppy: true, vars: true, white: true, nomen: true */
/*global module, require, __dirname */

var express = require('express');
var request = require('request-json');
var async = require('async');

var router = express.Router();


/* GET the games list. */
router.get('/game/', function(req, res) {
    var client = request.newClient(req.app.get('dataapi url'));
    async.parallel([
        function(cb) {
            client.get("game/", function(err, dataRes, games) {
                cb(err, games);
            });
        }
    ],
    function(err, results) {
        // Process the returned data.
        if (err) {
            /* TODO: Better error handling. */
            res.render("500.html", {
                message: err.message,
                error: err
            });
            return;
        }

        var games = results[0];
        res.render('games.html', { games: games});
    });
});

module.exports = router;
